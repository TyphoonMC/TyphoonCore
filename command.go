package typhoon

import (
	"log"
	"strings"
)

type CommandNodeType uint8

const (
	commandNodeTypeRoot = iota
	CommandNodeTypeLiteral
	CommandNodeTypeArgument
)

type CommandSuggestionType string

const (
	CommandSuggestionNone               CommandSuggestionType = ""
	CommandSuggestionAskServer          CommandSuggestionType = "minecraft:ask_server"
	CommandSuggestionAllRecipes         CommandSuggestionType = "minecraft:all_recipes"
	CommandSuggestionAvailableSounds    CommandSuggestionType = "minecraft:available_sounds"
	CommandSuggestionSummonableEntities CommandSuggestionType = "minecraft:summonable_entities"
)

type CommandParser interface {
	GetId() string
	IsMultiple() bool
	IsValid(string) bool
	IsArrayValid([]string) bool
	GetSuggestion() CommandSuggestionType
	Complete(string) []string
	writeProperties(*Player) error
}

type CommandNode struct {
	Type         CommandNodeType
	Execute      func(player *Player, args []string)
	Children     []*CommandNode
	RedirectNode *CommandNode
	Name         string
	Parser       CommandParser
}

func CommandNodeLiteral(
	name string,
	children []*CommandNode,
	execute func(player *Player, args []string)) *CommandNode {
	return &CommandNode{
		CommandNodeTypeLiteral,
		execute,
		children,
		nil,
		name,
		nil,
	}
}

func CommandNodeArgument(
	name string,
	children []*CommandNode,
	parser CommandParser,
	execute func(player *Player, args []string)) *CommandNode {
	return &CommandNode{
		CommandNodeTypeArgument,
		execute,
		children,
		nil,
		name,
		parser,
	}
}

func (c *Core) DeclareCommand(graph *CommandNode) {
	c.rootCommand.Children = append(c.rootCommand.Children, graph)
	c.compileCommands()
}

func (c *Core) onCommand(player *Player, command string) {
	args := strings.Split(command, " ")

	node := &c.rootCommand
	if !c.analyseCommand(player, args, node, 0) {
		m := ChatMessage("Unknown command")
		m.SetColor(&ChatColorRed)
		player.SendMessage(m)
	}
}

func (c *Core) analyseCommand(player *Player, args []string, node *CommandNode, step int) bool {
	if len(args) > step {
		for _, child := range node.Children {
			switch child.Type {
			case CommandNodeTypeLiteral:
				if child.Name == args[step] {
					return c.analyseCommand(player, args, child, step+1)
				}
			case CommandNodeTypeArgument:
				if child.Parser.IsMultiple() {
					if child.Parser.IsArrayValid(args[step:]) {
						return c.analyseCommand(player, args, child, step+1)
					}
				} else {
					if child.Parser.IsValid(args[step]) {
						return c.analyseCommand(player, args, child, step+1)
					}
				}
			}
		}
	}
	if node.Execute != nil {
		node.Execute(player, args[:step])
		return true
	}
	return false
}

func (c *Core) onTabCommand(player *Player, command string) {
	args := strings.Split(command, " ")

	node := &c.rootCommand
	player.WritePacket(&PacketPlayTabComplete{
		c.analyseTabCommand(player, args, node, 0),
	})
}

func commandJoin(args []string, attr string) string {
	str := attr
	if len(args) == 0 {
		str = "/" + str
	}
	return str
}

func (c *Core) analyseTabCommand(player *Player, args []string, node *CommandNode, step int) []string {
	strs := make([]string, 0)
	if len(args) > step {
		for _, child := range node.Children {
			switch child.Type {
			case CommandNodeTypeLiteral:
				if len(args)-1 == step && strings.HasPrefix(child.Name, args[step]) {
					strs = append(strs, commandJoin(args[:len(args)-1], child.Name))
				} else {
					if child.Name == args[step] {
						strs = append(strs, c.analyseTabCommand(player, args, child, step+1)...)
					}
				}
			case CommandNodeTypeArgument:
				if child.Parser.IsMultiple() {
					if child.Parser.IsArrayValid(args[step:]) {
						strs = append(strs, strings.Join(args, " "))
					}
				} else {
					strs = append(strs, child.Parser.Complete(args[step])...)
				}
			}
		}
	}
	return strs
}

type commandNode struct {
	Type         CommandNodeType
	Execute      bool
	Children     []int
	RedirectNode int
	Name         string
	Parser       CommandParser
}

func countNodes(node *CommandNode) int {
	i := 1
	for _, n := range node.Children {
		i += countNodes(n)
	}
	return i
}

func nodeRefs(refs map[*CommandNode]int, node *CommandNode, index int) int {
	refs[node] = index
	i := 1
	for _, n := range node.Children {
		i += nodeRefs(refs, n, index+i)
	}
	return i
}

func compileNode(compile []commandNode, refs map[*CommandNode]int, node *CommandNode) {
	children := make([]int, len(node.Children))
	for i, n := range node.Children {
		compileNode(compile, refs, n)
		children[i] = refs[n]
	}

	redirect := -1
	if node.RedirectNode != nil {
		redirect = refs[node.RedirectNode]
	}

	compile[refs[node]] = commandNode{
		node.Type,
		node.Execute != nil,
		children,
		redirect,
		node.Name,
		node.Parser,
	}
}

func (c *Core) compileCommands() {
	count := countNodes(&c.rootCommand)
	commands := make([]commandNode, count)
	refs := make(map[*CommandNode]int, count)
	nodeRefs(refs, &c.rootCommand, 0)
	compileNode(commands, refs, &c.rootCommand)
	c.compiledCommands = commands
}

func (node *commandNode) flags() uint8 {
	flags := uint8(node.Type)
	if node.Execute {
		flags |= 0x04
	}
	if node.RedirectNode != -1 {
		flags |= 0x08
	}
	if node.Parser != nil && node.Parser.GetSuggestion() != CommandSuggestionNone {
		flags |= 0x10
	}
	return flags
}

func (node *commandNode) writeTo(player *Player) (err error) {
	err = player.WriteUInt8(node.flags())
	if err != nil {
		log.Print(err)
		return
	}
	err = player.WriteVarInt(len(node.Children))
	if err != nil {
		log.Print(err)
		return
	}
	for _, child := range node.Children {
		err = player.WriteVarInt(child)
		if err != nil {
			log.Print(err)
			return
		}
	}
	if node.RedirectNode != -1 {
		err = player.WriteVarInt(node.RedirectNode)
		if err != nil {
			log.Print(err)
			return
		}
	}
	if node.Type == CommandNodeTypeArgument ||
		node.Type == CommandNodeTypeLiteral {
		err = player.WriteString(node.Name)
		if err != nil {
			log.Print(err)
			return
		}
	}
	if node.Type == CommandNodeTypeArgument {
		err = player.WriteString(node.Parser.GetId())
		if err != nil {
			log.Print(err)
			return
		}
		node.Parser.writeProperties(player)
	}
	if node.Parser != nil && node.Parser.GetSuggestion() != CommandSuggestionNone {
		err = player.WriteString(string(node.Parser.GetSuggestion()))
		if err != nil {
			log.Print(err)
			return
		}
	}
	return
}

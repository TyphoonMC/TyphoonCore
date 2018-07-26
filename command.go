package typhoon

import (
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

type CommandParserBool struct{}

func (c *CommandParserBool) GetId() string {
	return "brigadier:bool"
}
func (c *CommandParserBool) IsMultiple() bool {
	return false
}
func (c *CommandParserBool) IsValid(arg string) bool {
	return arg == "true" || arg == "false"
}
func (c *CommandParserBool) IsArrayValid(arg []string) bool {
	return false
}

func (c *Core) DeclareCommand(graph *CommandNode) {
	c.rootCommand.Children = append(c.rootCommand.Children, graph)
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

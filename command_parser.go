package typhoon

import (
	"log"
	"strconv"
	"strings"
)

type OptDouble struct {
	Used  bool
	Value float64
}

type OptFloat struct {
	Used  bool
	Value float32
}

type OptInteger struct {
	Used  bool
	Value int32
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
func (c *CommandParserBool) Complete(arg string) []string {
	ans := make([]string, 0)
	if strings.HasPrefix("true", arg) {
		ans = append(ans, "true")
	}
	if strings.HasPrefix("false", arg) {
		ans = append(ans, "false")
	}
	return ans
}
func (c *CommandParserBool) GetSuggestion() CommandSuggestionType {
	return CommandSuggestionNone
}
func (c *CommandParserBool) writeProperties(player *Player) (err error) {
	return
}

type CommandParserDouble struct {
	Min OptDouble
	Max OptDouble
}

func (c *CommandParserDouble) GetId() string {
	return "brigadier:double"
}
func (c *CommandParserDouble) IsMultiple() bool {
	return false
}
func (c *CommandParserDouble) IsValid(arg string) bool {
	_, err := strconv.ParseFloat(arg, 64)
	return err == nil
}
func (c *CommandParserDouble) IsArrayValid(arg []string) bool {
	return false
}
func (c *CommandParserDouble) Complete(arg string) []string {
	return []string{arg}
}
func (c *CommandParserDouble) GetSuggestion() CommandSuggestionType {
	return CommandSuggestionNone
}
func (c *CommandParserDouble) writeProperties(player *Player) (err error) {
	flags := uint8(0)
	if c.Min.Used {
		flags |= 0x01
	}
	if c.Max.Used {
		flags |= 0x02
	}
	err = player.WriteUInt8(flags)
	if err != nil {
		log.Print(err)
		return
	}
	if c.Min.Used {
		err = player.WriteFloat64(c.Min.Value)
		if err != nil {
			log.Print(err)
			return
		}
	}
	if c.Max.Used {
		err = player.WriteFloat64(c.Max.Value)
		if err != nil {
			log.Print(err)
			return
		}
	}
	return
}

type CommandParserFloat struct {
	Min OptFloat
	Max OptFloat
}

func (c *CommandParserFloat) GetId() string {
	return "brigadier:float"
}
func (c *CommandParserFloat) IsMultiple() bool {
	return false
}
func (c *CommandParserFloat) IsValid(arg string) bool {
	_, err := strconv.ParseFloat(arg, 32)
	return err == nil
}
func (c *CommandParserFloat) IsArrayValid(arg []string) bool {
	return false
}
func (c *CommandParserFloat) Complete(arg string) []string {
	return []string{arg}
}
func (c *CommandParserFloat) GetSuggestion() CommandSuggestionType {
	return CommandSuggestionNone
}
func (c *CommandParserFloat) writeProperties(player *Player) (err error) {
	flags := uint8(0)
	if c.Min.Used {
		flags |= 0x01
	}
	if c.Max.Used {
		flags |= 0x02
	}
	err = player.WriteUInt8(flags)
	if err != nil {
		log.Print(err)
		return
	}
	if c.Min.Used {
		err = player.WriteFloat32(c.Min.Value)
		if err != nil {
			log.Print(err)
			return
		}
	}
	if c.Max.Used {
		err = player.WriteFloat32(c.Max.Value)
		if err != nil {
			log.Print(err)
			return
		}
	}
	return
}

type CommandParserInteger struct {
	Min OptInteger
	Max OptInteger
}

func (c *CommandParserInteger) GetId() string {
	return "brigadier:integer"
}
func (c *CommandParserInteger) IsMultiple() bool {
	return false
}
func (c *CommandParserInteger) IsValid(arg string) bool {
	_, err := strconv.ParseInt(arg, 10, 32)
	return err == nil
}
func (c *CommandParserInteger) IsArrayValid(arg []string) bool {
	return false
}
func (c *CommandParserInteger) Complete(arg string) []string {
	return []string{arg}
}
func (c *CommandParserInteger) GetSuggestion() CommandSuggestionType {
	return CommandSuggestionNone
}
func (c *CommandParserInteger) writeProperties(player *Player) (err error) {
	flags := uint8(0)
	if c.Min.Used {
		flags |= 0x01
	}
	if c.Max.Used {
		flags |= 0x02
	}
	err = player.WriteUInt8(flags)
	if err != nil {
		log.Print(err)
		return
	}
	if c.Min.Used {
		err = player.WriteUInt32(uint32(c.Min.Value))
		if err != nil {
			log.Print(err)
			return
		}
	}
	if c.Max.Used {
		err = player.WriteUInt32(uint32(c.Max.Value))
		if err != nil {
			log.Print(err)
			return
		}
	}
	return
}

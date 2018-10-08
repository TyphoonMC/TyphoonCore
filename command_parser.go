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
	nb, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		return false
	}
	if c.Min.Used && c.Min.Value > nb {
		return false
	}
	if c.Max.Used && c.Max.Value < nb {
		return false
	}
	return true
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
	nb, err := strconv.ParseFloat(arg, 32)
	if err != nil {
		return false
	}
	if c.Min.Used && c.Min.Value > float32(nb) {
		return false
	}
	if c.Max.Used && c.Max.Value < float32(nb) {
		return false
	}
	return true
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
	nb, err := strconv.ParseInt(arg, 10, 32)
	if err != nil {
		return false
	}
	if c.Min.Used && c.Min.Value > int32(nb) {
		return false
	}
	if c.Max.Used && c.Max.Value < int32(nb) {
		return false
	}
	return true
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

type CommandParserStringFormat int

const (
	CommandParserStringFormatSingleWord     CommandParserStringFormat = iota
	commandParserStringFormatQuotablePhrase                           //TODO
	CommandParserStringFormatGreedyPhrase
)

type CommandParserString struct {
	Format CommandParserStringFormat
}

func (c *CommandParserString) GetId() string {
	return "brigadier:string"
}
func (c *CommandParserString) IsMultiple() bool {
	return c.Format != CommandParserStringFormatSingleWord
}
func (c *CommandParserString) IsValid(arg string) bool {
	return true
}
func (c *CommandParserString) IsArrayValid(arg []string) bool {
	return true
}
func (c *CommandParserString) Complete(arg string) []string {
	return []string{arg}
}
func (c *CommandParserString) GetSuggestion() CommandSuggestionType {
	return CommandSuggestionNone
}
func (c *CommandParserString) writeProperties(player *Player) (err error) {
	err = player.WriteVarInt(int(c.Format))
	if err != nil {
		log.Print(err)
		return
	}
	return
}

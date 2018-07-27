package typhoon

import "strings"

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

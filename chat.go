package typhoon

import (
	"encoding/json"
	"log"
)

type ChatColor struct {
	id   byte
	name string
}

func (color *ChatColor) GetId() byte {
	return color.id
}
func (color *ChatColor) GetName() string {
	return color.name
}

var (
	ChatColorBlack      = ChatColor{0, "black"}
	ChatColorDarkBlue   = ChatColor{1, "dark_blue"}
	ChatColorDarkGreen  = ChatColor{2, "dark_green"}
	ChatColorDarkAqua   = ChatColor{3, "dark_aqua"}
	ChatColorDarkRed    = ChatColor{4, "dark_red"}
	ChatColorDarkPurple = ChatColor{5, "dark_purple"}
	ChatColorGold       = ChatColor{6, "gold"}
	ChatColorGray       = ChatColor{7, "gray"}
	ChatColorDarkGray   = ChatColor{8, "dark_gray"}
	ChatColorIndigo     = ChatColor{9, "blue"}
	ChatColorGreen      = ChatColor{10, "green"}
	ChatColorAqua       = ChatColor{11, "aqua"}
	ChatColorRed        = ChatColor{12, "red"}
	ChatColorPink       = ChatColor{13, "light_purple"}
	ChatColorYellow     = ChatColor{14, "yellow"}
	ChatColorWhite      = ChatColor{15, "white"}
)

type ChatAction struct {
	Action string `json:"action"`
	Value  string `json:"value"`
}

type IChatComponent interface {
	JSON() (string, error)
}
type ChatComponent struct {
	Bold          bool             `json:"bold,omitempty"`
	Italic        bool             `json:"italic,omitempty"`
	Underlined    bool             `json:"underlined,omitempty"`
	StrikeThrough bool             `json:"strikethrough,omitempty"`
	Obfuscated    bool             `json:"obfuscated,omitempty"`
	Color         *string          `json:"color,omitempty"`
	Insertion     *string          `json:"insertion,omitempty"`
	ClickEvent    *ChatAction      `json:"clickEvent,omitempty"`
	HoverEvent    *ChatAction      `json:"hoverEvent,omitempty"`
	Extra         []IChatComponent `json:"extra,omitempty"`
}

func (component *ChatComponent) SetBold(bold bool) {
	component.Bold = bold
}
func (component *ChatComponent) SetItalic(italic bool) {
	component.Italic = italic
}
func (component *ChatComponent) SetUnderlined(underlined bool) {
	component.Underlined = underlined
}
func (component *ChatComponent) SetStrikeThrough(strikeThrough bool) {
	component.StrikeThrough = strikeThrough
}
func (component *ChatComponent) SetObfuscated(obfuscated bool) {
	component.Obfuscated = obfuscated
}
func (component *ChatComponent) SetColor(color *ChatColor) {
	component.Color = &color.name
}
func (component *ChatComponent) SetInsertion(insertion string) {
	component.Insertion = &insertion
}
func (component *ChatComponent) SetClickEvent(action *ChatAction) {
	component.ClickEvent = action
}
func (component *ChatComponent) SetHoverEvent(action *ChatAction) {
	component.HoverEvent = action
}
func (component *ChatComponent) AddExtra(extra IChatComponent) {
	component.Extra = append(component.Extra, extra)
}
func (component *ChatComponent) SetExtra(extra []IChatComponent) {
	component.Extra = extra
}
func (component *ChatComponent) JSON() (string, error) {
	b, err := json.Marshal(component)
	return string(b), err
}

type StringChatComponent struct {
	Text string `json:"text"`
	*ChatComponent
}

func (component *StringChatComponent) SetText(text string) {
	component.Text = text
}
func (component *StringChatComponent) JSON() (string, error) {
	b, err := json.Marshal(component)
	return string(b), err
}

func ChatMessage(text string) *StringChatComponent {
	return &StringChatComponent{
		text,
		&ChatComponent{
			false,
			false,
			false,
			false,
			false,
			nil,
			nil,
			nil,
			nil,
			nil,
		}}
}

func (p *Player) SendMessage(message IChatComponent) {
	msg, err := message.JSON()
	if err != nil {
		log.Fatal(err)
		return
	}
	p.WritePacket(&PacketPlayMessage{
		msg,
		CHAT_BOX,
	})
}

func (p *Player) SendRawMessage(message string) {
	p.WritePacket(&PacketPlayMessage{
		message,
		CHAT_BOX,
	})
}

func (p *Player) SendActionBar(message IChatComponent) {
	msg, err := message.JSON()
	if err != nil {
		log.Fatal(err)
		return
	}
	p.WritePacket(&PacketPlayMessage{
		msg,
		ACTION_BAR,
	})
}

func (p *Player) SendRawActionBar(message string) {
	p.WritePacket(&PacketPlayMessage{
		message,
		ACTION_BAR,
	})
}

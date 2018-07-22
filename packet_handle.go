package typhoon

import (
	"fmt"
	"github.com/TyphoonMC/go.uuid"
	"log"
)

type PacketHandshake struct {
	Protocol Protocol
	Address  string
	Port     uint16
	State    State
}

func (packet *PacketHandshake) Read(player *Player, length int) (err error) {
	protocol, err := player.ReadVarInt()
	if err != nil {
		log.Print(err)
		return
	}
	packet.Protocol = Protocol(protocol)
	packet.Address, err = player.ReadStringLimited(config.BufferConfig.HandshakeAddress)
	if err != nil {
		log.Print(err)
		return
	}
	packet.Port, err = player.ReadUInt16()
	if err != nil {
		log.Print(err)
		return
	}
	state, err := player.ReadVarInt()
	if err != nil {
		log.Print(err)
		return
	}
	packet.State = State(state)
	return
}
func (packet *PacketHandshake) Write(player *Player) (err error) {
	return
}
func (packet *PacketHandshake) Handle(player *Player) {
	player.state = packet.State
	player.protocol = packet.Protocol
	player.inaddr.address = packet.Address
	player.inaddr.port = packet.Port
}
func (packet *PacketHandshake) Id() int {
	return 0x00
}

type PacketStatusRequest struct{}

func (packet *PacketStatusRequest) Read(player *Player, length int) (err error) {
	return
}
func (packet *PacketStatusRequest) Write(player *Player) (err error) {
	return
}
func (packet *PacketStatusRequest) Handle(player *Player) {
	protocol := COMPATIBLE_PROTO[0]
	if IsCompatible(player.protocol) {
		protocol = player.protocol
	}

	max_players := config.MaxPlayers
	motd := config.Motd

	if max_players < playersCount && !config.Restricted {
		max_players = playersCount
	}

	response := PacketStatusResponse{
		Response: fmt.Sprintf(`{"version":{"name":"Typhoon","protocol":%d},"players":{"max":%d,"online":%d,"sample":[]},"description":{"text":"%s"},"favicon":"%s","modinfo":{"type":"FML","modList":[]}}`, protocol, max_players, playersCount, JsonEscape(motd), JsonEscape(favicon)),
	}
	player.WritePacket(&response)
}
func (packet *PacketStatusRequest) Id() int {
	return 0x00
}

type PacketStatusResponse struct {
	Response string
}

func (packet *PacketStatusResponse) Read(player *Player, length int) (err error) {
	return
}
func (packet *PacketStatusResponse) Write(player *Player) (err error) {
	err = player.WriteString(packet.Response)
	if err != nil {
		log.Print(err)
		return
	}
	return
}
func (packet *PacketStatusResponse) Handle(player *Player) {}
func (packet *PacketStatusResponse) Id() int {
	return 0x00
}

type PacketStatusPing struct {
	Time uint64
}

func (packet *PacketStatusPing) Read(player *Player, length int) (err error) {
	packet.Time, err = player.ReadUInt64()
	if err != nil {
		log.Print(err)
		return
	}
	return
}
func (packet *PacketStatusPing) Write(player *Player) (err error) {
	err = player.WriteUInt64(packet.Time)
	if err != nil {
		log.Print(err)
		return
	}
	return
}
func (packet *PacketStatusPing) Handle(player *Player) {
	player.WritePacket(packet)
}
func (packet *PacketStatusPing) Id() int {
	return 0x01
}

type PacketLoginStart struct {
	Username string
}

func (packet *PacketLoginStart) Read(player *Player, length int) (err error) {
	packet.Username, err = player.ReadStringLimited(config.BufferConfig.PlayerName)
	if err != nil {
		log.Print(err)
		return
	}
	return
}
func (packet *PacketLoginStart) Write(player *Player) (err error) {
	return
}

var (
	join_game = PacketPlayJoinGame{
		EntityId:     0,
		Gamemode:     SPECTATOR,
		Dimension:    END,
		Difficulty:   NORMAL,
		LevelType:    DEFAULT,
		MaxPlayers:   0xFF,
		ReducedDebug: false,
	}
	position_look = PacketPlayerPositionLook{}
)

func (packet *PacketLoginStart) Handle(player *Player) {
	if !IsCompatible(player.protocol) {
		player.LoginKick("Incompatible version")
		return
	}

	max_players := config.MaxPlayers

	if max_players <= playersCount && config.Restricted {
		player.LoginKick("Server is full")
	}

	player.name = packet.Username

	if config.Compression && player.protocol >= V1_8 {
		setCompression := PacketSetCompression{config.Threshold}
		player.WritePacket(&setCompression)
		player.compression = true
	}

	success := PacketLoginSuccess{
		UUID:     player.uuid,
		Username: player.name,
	}
	player.WritePacket(&success)
	player.state = PLAY
	player.register()

	player.WritePacket(&join_game)
	player.WritePacket(&position_look)

	player.core.CallEvent(&PlayerJoinEvent{
		player,
	})
}
func (packet *PacketLoginStart) Id() int {
	return 0x00
}

type PacketLoginDisconnect struct {
	Component string
}

func (packet *PacketLoginDisconnect) Read(player *Player, length int) (err error) {
	return
}
func (packet *PacketLoginDisconnect) Write(player *Player) (err error) {
	err = player.WriteString(packet.Component)
	if err != nil {
		log.Print(err)
		return
	}
	return
}
func (packet *PacketLoginDisconnect) Handle(player *Player) {}
func (packet *PacketLoginDisconnect) Id() int {
	return 0x00
}

type PacketLoginSuccess struct {
	UUID     string
	Username string
}

func (packet *PacketLoginSuccess) Read(player *Player, length int) (err error) {
	return
}
func (packet *PacketLoginSuccess) Write(player *Player) (err error) {
	err = player.WriteString(packet.UUID)
	if err != nil {
		log.Print(err)
		return
	}
	err = player.WriteString(packet.Username)
	if err != nil {
		log.Print(err)
		return
	}
	return
}
func (packet *PacketLoginSuccess) Handle(player *Player) {}
func (packet *PacketLoginSuccess) Id() int {
	return 0x02
}

type PacketSetCompression struct {
	Threshold int
}

func (packet *PacketSetCompression) Read(player *Player, length int) (err error) {
	return
}
func (packet *PacketSetCompression) Write(player *Player) (err error) {
	err = player.WriteVarInt(packet.Threshold)
	if err != nil {
		log.Print(err)
		return
	}
	return
}
func (packet *PacketSetCompression) Handle(player *Player) {}
func (packet *PacketSetCompression) Id() int {
	return 0x03
}

type PacketPlayChat struct {
	Message string
}

func (packet *PacketPlayChat) Read(player *Player, length int) (err error) {
	packet.Message, err = player.ReadStringLimited(config.BufferConfig.ChatMessage)
	if err != nil {
		log.Print(err)
		return
	}
	return
}
func (packet *PacketPlayChat) Write(player *Player) (err error) {
	return
}
func (packet *PacketPlayChat) Handle(player *Player) {
	if len(packet.Message) > 0 {
		if packet.Message[0] != '/' {
			player.core.CallEvent(&PlayerChatEvent{
				player,
				packet.Message,
			})
		} else {
			//TODO handle commands
		}
	}
}
func (packet *PacketPlayChat) Id() int {
	return 0x02
}

type PacketPlayMessage struct {
	Component string
	Position  ChatPosition
}

func (packet *PacketPlayMessage) Read(player *Player, length int) (err error) {
	return
}
func (packet *PacketPlayMessage) Write(player *Player) (err error) {
	err = player.WriteString(packet.Component)
	if err != nil {
		log.Print(err)
		return
	}
	if player.protocol > V1_7_6 {
		err = player.WriteUInt8(uint8(packet.Position))
		if err != nil {
			log.Print(err)
			return
		}
	}
	return
}
func (packet *PacketPlayMessage) Handle(player *Player) {}
func (packet *PacketPlayMessage) Id() int {
	return 0x0F
}

type PacketBossBar struct {
	UUID     uuid.UUID
	Action   BossBarAction
	Title    string
	Health   float32
	Color    BossBarColor
	Division BossBarDivision
	Flags    uint8
}

func (packet *PacketBossBar) Read(player *Player, length int) (err error) {
	return
}
func (packet *PacketBossBar) Write(player *Player) (err error) {
	err = player.WriteUUID(packet.UUID)
	if err != nil {
		log.Print(err)
		return
	}
	err = player.WriteVarInt(int(packet.Action))
	if err != nil {
		log.Print(err)
		return
	}
	if packet.Action == BOSSBAR_UPDATE_TITLE || packet.Action == BOSSBAR_ADD {
		err = player.WriteString(packet.Title)
		if err != nil {
			log.Print(err)
			return
		}
	}
	if packet.Action == BOSSBAR_UPDATE_HEALTH || packet.Action == BOSSBAR_ADD {
		err = player.WriteFloat32(packet.Health)
		if err != nil {
			log.Print(err)
			return
		}
	}
	if packet.Action == BOSSBAR_UPDATE_STYLE || packet.Action == BOSSBAR_ADD {
		err = player.WriteVarInt(int(packet.Color))
		if err != nil {
			log.Print(err)
			return
		}
		err = player.WriteVarInt(int(packet.Division))
		if err != nil {
			log.Print(err)
			return
		}
	}
	if packet.Action == BOSSBAR_UPDATE_STYLE || packet.Action == BOSSBAR_ADD {
		err = player.WriteUInt8(packet.Flags)
		if err != nil {
			log.Print(err)
			return
		}
	}
	return
}
func (packet *PacketBossBar) Handle(player *Player) {}
func (packet *PacketBossBar) Id() int {
	return 0x0C
}

type PacketPlayPluginMessage struct {
	Channel string
	Data    []byte
}

func (packet *PacketPlayPluginMessage) Read(player *Player, length int) (err error) {
	var read int
	packet.Channel, read, err = player.ReadNStringLimited(20)
	if err != nil {
		log.Print(err)
		return
	}

	dataLength := length - read
	if player.protocol < V1_8 {
		sread, err := player.ReadUInt16()
		if err != nil {
			log.Print(err)
			return err
		}
		dataLength = int(sread)
	}

	packet.Data, err = player.ReadByteArray(dataLength)
	if err != nil {
		log.Print(err)
		return
	}
	return
}
func (packet *PacketPlayPluginMessage) Write(player *Player) (err error) {
	err = player.WriteString(packet.Channel)
	if err != nil {
		log.Print(err)
		return
	}
	if player.protocol < V1_8 {
		err = player.WriteUInt16(uint16(len(packet.Data)))
		if err != nil {
			log.Print(err)
			return err
		}
	}
	err = player.WriteByteArray(packet.Data)
	if err != nil {
		log.Print(err)
		return
	}
	return
}
func (packet *PacketPlayPluginMessage) Handle(player *Player) {
	if packet.Channel == "MC|Brand" || packet.Channel == "minecraft:brand" {
		log.Printf("%s is using %s client", player.name, string(packet.Data))
		player.WritePacket(&PacketPlayPluginMessage{
			packet.Channel,
			[]byte(player.core.brand),
		})
	}
	player.core.CallEvent(&PluginMessageEvent{
		packet.Channel,
		packet.Data,
	})
}
func (packet *PacketPlayPluginMessage) Id() int {
	return 0x18
}

type PacketPlayDisconnect struct {
	Component string
}

func (packet *PacketPlayDisconnect) Read(player *Player, length int) (err error) {
	return
}
func (packet *PacketPlayDisconnect) Write(player *Player) (err error) {
	err = player.WriteString(packet.Component)
	if err != nil {
		log.Print(err)
		return
	}
	return
}
func (packet *PacketPlayDisconnect) Handle(player *Player) {}
func (packet *PacketPlayDisconnect) Id() int {
	return 0x1A
}

type PacketPlayKeepAlive struct {
	Identifier int
}

func (packet *PacketPlayKeepAlive) Read(player *Player, length int) (err error) {
	if player.protocol >= V1_12_2 {
		id, stt := player.ReadUInt64()
		packet.Identifier = int(id)
		err = stt
	} else if player.protocol <= V1_7_6 {
		id, stt := player.ReadUInt32()
		packet.Identifier = int(id)
		err = stt
	} else {
		packet.Identifier, err = player.ReadVarInt()
	}
	if err != nil {
		log.Print(err)
		return
	}
	return
}
func (packet *PacketPlayKeepAlive) Write(player *Player) (err error) {
	if player.protocol >= V1_12_2 {
		err = player.WriteUInt64(uint64(packet.Identifier))
	} else if player.protocol <= V1_7_6 {
		err = player.WriteUInt32(uint32(packet.Identifier))
	} else {
		err = player.WriteVarInt(packet.Identifier)
	}
	if err != nil {
		log.Print(err)
		return
	}
	return
}
func (packet *PacketPlayKeepAlive) Handle(player *Player) {
	if player.protocol > V1_8 {
		if player.keepalive != packet.Identifier {
			player.Kick("Invalid keepalive")
		}
	} else {
		player.keepalive = packet.Identifier
	}
	player.keepalive = 0
}
func (packet *PacketPlayKeepAlive) Id() int {
	return 0x1F
}

type PacketPlayJoinGame struct {
	EntityId     uint32
	Gamemode     Gamemode
	Dimension    Dimension
	Difficulty   Difficulty
	MaxPlayers   uint8
	LevelType    LevelType
	ReducedDebug bool
}

func (packet *PacketPlayJoinGame) Read(player *Player, length int) (err error) {
	return
}
func (packet *PacketPlayJoinGame) Write(player *Player) (err error) {
	if player.protocol <= V1_9 {
		err = player.WriteUInt8(uint8(packet.EntityId))
	} else {
		err = player.WriteUInt32(packet.EntityId)
	}
	if err != nil {
		log.Print(err)
		return
	}
	err = player.WriteUInt8(uint8(packet.Gamemode))
	if err != nil {
		log.Print(err)
		return
	}
	err = player.WriteUInt32(uint32(packet.Dimension))
	if err != nil {
		log.Print(err)
		return
	}
	err = player.WriteUInt8(uint8(packet.Difficulty))
	if err != nil {
		log.Print(err)
		return
	}
	err = player.WriteUInt8(packet.MaxPlayers)
	if err != nil {
		log.Print(err)
		return
	}
	err = player.WriteString(string(packet.LevelType))
	if err != nil {
		log.Print(err)
		return
	}
	if player.protocol > V1_7_6 {
		err = player.WriteBool(packet.ReducedDebug)
		if err != nil {
			log.Print(err)
			return
		}
	}
	return
}
func (packet *PacketPlayJoinGame) Handle(player *Player) {}
func (packet *PacketPlayJoinGame) Id() int {
	return 0x23
}

type PacketPlayerPositionLook struct {
	X          float64
	Y          float64
	Z          float64
	Yaw        float32
	Pitch      float32
	Flags      uint8
	TeleportId int
}

func (packet *PacketPlayerPositionLook) Read(player *Player, length int) (err error) {
	return
}
func (packet *PacketPlayerPositionLook) Write(player *Player) (err error) {
	err = player.WriteFloat64(packet.X)
	if err != nil {
		log.Print(err)
		return
	}
	err = player.WriteFloat64(packet.Y)
	if err != nil {
		log.Print(err)
		return
	}
	err = player.WriteFloat64(packet.Z)
	if err != nil {
		log.Print(err)
		return
	}
	err = player.WriteFloat32(packet.Yaw)
	if err != nil {
		log.Print(err)
		return
	}
	err = player.WriteFloat32(packet.Pitch)
	if err != nil {
		log.Print(err)
		return
	}
	err = player.WriteUInt8(packet.Flags)
	if err != nil {
		log.Print(err)
		return
	}
	if player.protocol > V1_8 {
		err = player.WriteVarInt(packet.TeleportId)
		if err != nil {
			log.Print(err)
			return
		}
	}
	return
}
func (packet *PacketPlayerPositionLook) Handle(player *Player) {}
func (packet *PacketPlayerPositionLook) Id() int {
	return 0x2E
}

type PacketPlayerListHeaderFooter struct {
	Header *string
	Footer *string
}

func (packet *PacketPlayerListHeaderFooter) Read(player *Player, length int) (err error) {
	return
}
func (packet *PacketPlayerListHeaderFooter) Write(player *Player) (err error) {
	var str string
	if packet.Header == nil {
		str = `{"translate":""}`
	} else {
		str = *packet.Header
	}
	err = player.WriteString(str)
	if err != nil {
		log.Print(err)
		return
	}
	if packet.Footer == nil {
		str = `{"translate":""}`
	} else {
		str = *packet.Footer
	}
	err = player.WriteString(str)
	if err != nil {
		log.Print(err)
		return
	}
	return
}
func (packet *PacketPlayerListHeaderFooter) Handle(player *Player) {}
func (packet *PacketPlayerListHeaderFooter) Id() int {
	return 0x47
}

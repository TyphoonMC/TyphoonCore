package typhoon

import (
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/seebs/nbt"
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
		slog.Debug("packet error", "error", err)
		return
	}
	packet.Protocol = Protocol(protocol)
	packet.Address, err = player.ReadStringLimited(config.BufferConfig.HandshakeAddress)
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	packet.Port, err = player.ReadUInt16()
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	state, err := player.ReadVarInt()
	if err != nil {
		slog.Debug("packet error", "error", err)
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
func (packet *PacketHandshake) Id() (int, Protocol) {
	return 0x00, V1_10
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

	count := player.core.playerRegistry.GetPlayerCount()
	if max_players < count && !config.Restricted {
		max_players = count
	}

	response := PacketStatusResponse{
		Response: fmt.Sprintf(`{"version":{"name":"Typhoon","protocol":%d},"players":{"max":%d,"online":%d,"sample":[]},"description":{"text":"%s"},"favicon":"%s","modinfo":{"type":"FML","modList":[]}}`, protocol, max_players, count, JsonEscape(motd), JsonEscape(favicon)),
	}
	player.WritePacket(&response)
}
func (packet *PacketStatusRequest) Id() (int, Protocol) {
	return 0x00, V1_10
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
		slog.Debug("packet error", "error", err)
		return
	}
	return
}
func (packet *PacketStatusResponse) Handle(player *Player) {}
func (packet *PacketStatusResponse) Id() (int, Protocol) {
	return 0x00, V1_10
}

type PacketStatusPing struct {
	Time uint64
}

func (packet *PacketStatusPing) Read(player *Player, length int) (err error) {
	packet.Time, err = player.ReadUInt64()
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	return
}
func (packet *PacketStatusPing) Write(player *Player) (err error) {
	err = player.WriteUInt64(packet.Time)
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	return
}
func (packet *PacketStatusPing) Handle(player *Player) {
	player.WritePacket(packet)
}
func (packet *PacketStatusPing) Id() (int, Protocol) {
	return 0x01, V1_10
}

type PacketLoginStart struct {
	Username string
}

func (packet *PacketLoginStart) Read(player *Player, length int) (err error) {
	packet.Username, err = player.ReadStringLimited(config.BufferConfig.PlayerName)
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	return
}
func (packet *PacketLoginStart) Write(player *Player) (err error) {
	return
}

func (packet *PacketLoginStart) Handle(player *Player) {
	if !IsCompatible(player.protocol) {
		player.Kick("Incompatible version")
		return
	}

	player.name = packet.Username

	// PlayerPreLoginEvent - can deny connection
	preLoginEvent := &PlayerPreLoginEvent{
		Player:  player,
		Name:    packet.Username,
		Address: player.conn.RemoteAddr().String(),
	}
	player.core.CallEvent(preLoginEvent)

	if preLoginEvent.IsCancelled() {
		player.Kick(preLoginEvent.KickReason)
		return
	}

	max_players := config.MaxPlayers
	count := player.core.playerRegistry.GetPlayerCount()
	if max_players <= count && config.Restricted {
		player.Kick("Server is full")
		return
	}

	// PlayerLoginEvent - can deny connection
	loginEvent := &PlayerLoginEvent{
		Player: player,
	}
	player.core.CallEvent(loginEvent)

	if loginEvent.IsCancelled() {
		player.Kick(loginEvent.KickReason)
		return
	}

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

	player.WritePacket(&PacketPlayJoinGame{
		EntityId:            1,
		Gamemode:            player.core.gamemode,
		Dimension:           player.core.world.Dimension,
		HashedSeed:          0,
		Difficulty:          player.core.difficulty,
		LevelType:           DEFAULT,
		MaxPlayers:          0xFF,
		ReducedDebug:        false,
		EnableRespawnScreen: true,
	})

	player.WritePacket(&PacketPlayServerDifficulty{
		player.core.difficulty,
	})

	player.WritePacket(&PacketPlayPlayerAbilities{
		true,
		true,
		true,
		false,
		0.1,
		0.1,
	})

	// Set initial player location
	player.SetLocation(Location{
		X:     player.core.world.Spawn.X,
		Y:     player.core.world.Spawn.Y,
		Z:     player.core.world.Spawn.Z,
		Yaw:   0,
		Pitch: 0,
	})

	player.core.CallEvent(&PlayerPreJoinEvent{
		Player: player,
	})

	player.WritePacket(&PacketPlayerPositionLook{
		player.core.world.Spawn.X,
		player.core.world.Spawn.Y,
		player.core.world.Spawn.Z,
		0,
		0,
		0xFF,
		0,
	})

	player.core.world.SendSpawnChunks(player)

	if player.protocol >= V1_13 {
		player.WritePacket(&PacketPlayDeclareCommands{
			player.core.compiledCommands,
			0,
		})
	}

	player.core.CallEvent(&PlayerJoinEvent{
		Player: player,
	})
}
func (packet *PacketLoginStart) Id() (int, Protocol) {
	return 0x00, V1_10
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
		slog.Debug("packet error", "error", err)
		return
	}
	return
}
func (packet *PacketLoginDisconnect) Handle(player *Player) {}
func (packet *PacketLoginDisconnect) Id() (int, Protocol) {
	return 0x00, V1_10
}

type PacketLoginSuccess struct {
	UUID     uuid.UUID
	Username string
}

func (packet *PacketLoginSuccess) Read(player *Player, length int) (err error) {
	return
}
func (packet *PacketLoginSuccess) Write(player *Player) (err error) {
	if player.protocol >= V1_16 {
		err = player.WriteUUID(packet.UUID)
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
	} else {
		err = player.WriteString(packet.UUID.String())
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
	}
	err = player.WriteString(packet.Username)
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	return
}
func (packet *PacketLoginSuccess) Handle(player *Player) {}
func (packet *PacketLoginSuccess) Id() (int, Protocol) {
	return 0x02, V1_10
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
		slog.Debug("packet error", "error", err)
		return
	}
	return
}
func (packet *PacketSetCompression) Handle(player *Player) {}
func (packet *PacketSetCompression) Id() (int, Protocol) {
	return 0x03, V1_10
}

type PacketPlayChat struct {
	Message string
}

func (packet *PacketPlayChat) Read(player *Player, length int) (err error) {
	packet.Message, err = player.ReadStringLimited(config.BufferConfig.ChatMessage)
	if err != nil {
		slog.Debug("packet error", "error", err)
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
			event := &PlayerChatEvent{
				Player:  player,
				Message: packet.Message,
			}
			player.core.CallEvent(event)
			// Chat broadcast is handled by the event listener
		} else {
			// Command
			cmd := packet.Message[1:]
			event := &PlayerCommandPreprocessEvent{
				Player:  player,
				Command: cmd,
			}
			player.core.CallEvent(event)

			if !event.IsCancelled() {
				player.core.onCommand(player, event.Command)
			}
		}
	}
}
func (packet *PacketPlayChat) Id() (int, Protocol) {
	return 0x02, V1_10
}

type PacketPlayTabComplete struct {
	Matches []string
}

func (packet *PacketPlayTabComplete) Read(player *Player, length int) (err error) {
	return
}
func (packet *PacketPlayTabComplete) Write(player *Player) (err error) {
	err = player.WriteVarInt(len(packet.Matches))
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	for _, s := range packet.Matches {
		err = player.WriteString(s)
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
	}
	return
}
func (packet *PacketPlayTabComplete) Handle(player *Player) {}
func (packet *PacketPlayTabComplete) Id() (int, Protocol) {
	return 0x0E, V1_10
}

type PacketPlayTabCompleteServerbound struct {
	Text          string
	AssumeCommand bool
	Position      Position
}

func (packet *PacketPlayTabCompleteServerbound) Read(player *Player, length int) (err error) {
	packet.Text, err = player.ReadStringLimited(config.BufferConfig.ChatMessage)
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	packet.AssumeCommand, err = player.ReadBool()
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	hasPosition, err := player.ReadBool()
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	if hasPosition {
		packet.Position, err = player.ReadPosition()
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
	}
	return
}
func (packet *PacketPlayTabCompleteServerbound) Write(player *Player) (err error) {
	return
}
func (packet *PacketPlayTabCompleteServerbound) Handle(player *Player) {
	if len(packet.Text) > 0 {
		event := &PlayerTabCompleteEvent{
			Player:      player,
			Buffer:      packet.Text,
			Completions: []string{},
		}
		player.core.CallEvent(event)

		if packet.Text[0] == '/' {
			player.core.onTabCommand(player, packet.Text[1:])
		}
	}
}
func (packet *PacketPlayTabCompleteServerbound) Id() (int, Protocol) {
	return 0x01, V1_10
}

type PacketPlayClientStatus struct {
	Action ClientStatusAction
}

func (packet *PacketPlayClientStatus) Read(player *Player, length int) (err error) {
	act, err := player.ReadVarInt()
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	packet.Action = ClientStatusAction(act)
	return
}
func (packet *PacketPlayClientStatus) Write(player *Player) (err error) {
	return
}
func (packet *PacketPlayClientStatus) Handle(player *Player) {
	return
}
func (packet *PacketPlayClientStatus) Id() (int, Protocol) {
	return 0x03, V1_10
}

type PacketPlayMessage struct {
	Component string
	Position  ChatPosition
	Sender    uuid.UUID
}

func (packet *PacketPlayMessage) Read(player *Player, length int) (err error) {
	return
}
func (packet *PacketPlayMessage) Write(player *Player) (err error) {
	err = player.WriteString(packet.Component)
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	if player.protocol > V1_7_6 {
		err = player.WriteUInt8(uint8(packet.Position))
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
	}
	if player.protocol >= V1_16 {
		err = player.WriteUUID(packet.Sender)
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
	}
	return
}
func (packet *PacketPlayMessage) Handle(player *Player) {}
func (packet *PacketPlayMessage) Id() (int, Protocol) {
	return 0x0F, V1_10
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
		slog.Debug("packet error", "error", err)
		return
	}
	err = player.WriteVarInt(int(packet.Action))
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	if packet.Action == BOSSBAR_UPDATE_TITLE || packet.Action == BOSSBAR_ADD {
		err = player.WriteString(packet.Title)
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
	}
	if packet.Action == BOSSBAR_UPDATE_HEALTH || packet.Action == BOSSBAR_ADD {
		err = player.WriteFloat32(packet.Health)
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
	}
	if packet.Action == BOSSBAR_UPDATE_STYLE || packet.Action == BOSSBAR_ADD {
		err = player.WriteVarInt(int(packet.Color))
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
		err = player.WriteVarInt(int(packet.Division))
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
	}
	if packet.Action == BOSSBAR_UPDATE_STYLE || packet.Action == BOSSBAR_ADD {
		err = player.WriteUInt8(packet.Flags)
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
	}
	return
}
func (packet *PacketBossBar) Handle(player *Player) {}
func (packet *PacketBossBar) Id() (int, Protocol) {
	return 0x0C, V1_10
}

type PacketPlayServerDifficulty struct {
	Difficulty Difficulty
}

func (packet *PacketPlayServerDifficulty) Read(player *Player, length int) (err error) {
	return
}
func (packet *PacketPlayServerDifficulty) Write(player *Player) (err error) {
	err = player.WriteUInt8(uint8(packet.Difficulty))
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	if player.protocol >= V1_14 {
		err = player.WriteBool(false)
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
	}
	return
}
func (packet *PacketPlayServerDifficulty) Handle(player *Player) {}
func (packet *PacketPlayServerDifficulty) Id() (int, Protocol) {
	return 0x0D, V1_10
}

type PacketPlayDeclareCommands struct {
	Nodes     []commandNode
	RootIndex int
}

func (packet *PacketPlayDeclareCommands) Read(player *Player, length int) (err error) {
	return
}
func (packet *PacketPlayDeclareCommands) Write(player *Player) (err error) {
	err = player.WriteVarInt(len(packet.Nodes))
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	for _, n := range packet.Nodes {
		err = (&n).writeTo(player)
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
	}
	err = player.WriteVarInt(packet.RootIndex)
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	return
}
func (packet *PacketPlayDeclareCommands) Handle(player *Player) {}
func (packet *PacketPlayDeclareCommands) Id() (int, Protocol) {
	return 0x11, V1_13
}

type PacketPlayPluginMessage struct {
	Channel string
	Data    []byte
}

func (packet *PacketPlayPluginMessage) Read(player *Player, length int) (err error) {
	var read int
	packet.Channel, read, err = player.ReadNStringLimited(20)
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}

	dataLength := length - read
	if player.protocol < V1_8 {
		sread, err := player.ReadUInt16()
		if err != nil {
			slog.Debug("packet error", "error", err)
			return err
		}
		dataLength = int(sread)
	}

	packet.Data, err = player.ReadByteArray(dataLength)
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	return
}
func (packet *PacketPlayPluginMessage) Write(player *Player) (err error) {
	err = player.WriteString(packet.Channel)
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	if player.protocol < V1_8 {
		err = player.WriteUInt16(uint16(len(packet.Data)))
		if err != nil {
			slog.Debug("packet error", "error", err)
			return err
		}
	}
	err = player.WriteByteArray(packet.Data)
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	return
}
func (packet *PacketPlayPluginMessage) Handle(player *Player) {
	if packet.Channel == "MC|Brand" || packet.Channel == "minecraft:brand" {
		slog.Info("player client brand", "player", player.name, "client", string(packet.Data))
		buff := make([]byte, len(player.core.brand)+1)
		length := binary.PutUvarint(buff, uint64(len(player.core.brand)))
		copy(buff[length:], []byte(player.core.brand))
		/*player.WritePacket(&PacketPlayPluginMessage{
			packet.Channel,
			buff,
		})*/
	}
	player.core.CallEvent(&PluginMessageEvent{
		Player:  player,
		Channel: packet.Channel,
		Data:    packet.Data,
	})
}
func (packet *PacketPlayPluginMessage) Id() (int, Protocol) {
	return 0x18, V1_10
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
		slog.Debug("packet error", "error", err)
		return
	}
	return
}
func (packet *PacketPlayDisconnect) Handle(player *Player) {}
func (packet *PacketPlayDisconnect) Id() (int, Protocol) {
	return 0x1A, V1_10
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
		slog.Debug("packet error", "error", err)
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
		slog.Debug("packet error", "error", err)
		return
	}
	return
}
func (packet *PacketPlayKeepAlive) Handle(player *Player) {
	if player.protocol > V1_8 {
		if player.keepalive != packet.Identifier {
			//player.Kick("Invalid keepalive")
		}
	} else {
		player.keepalive = packet.Identifier
	}
	player.keepalive = 0
}
func (packet *PacketPlayKeepAlive) Id() (int, Protocol) {
	return 0x1F, V1_10
}

type PacketPlayChunkData struct {
	X               int32
	Z               int32
	Dimension       Dimension
	GroundUp        bool
	ServerLightning bool
	Sections        []*ChunkSection
	Biomes          *[]byte
	BlockEntities   []nbt.Compound
}

func (packet *PacketPlayChunkData) Read(player *Player, length int) (err error) {
	return
}
func (packet *PacketPlayChunkData) Write(player *Player) (err error) {
	if player.protocol >= V1_13 {
		return packet.WriteV1_13(player)
	} else if player.protocol >= V1_9 {
		return packet.WriteV1_9(player)
	} else if player.protocol >= V1_8 {
		return packet.WriteV1_8(player)
	}
	return packet.WriteV1_7(player)
}
func (packet *PacketPlayChunkData) WriteV1_13(player *Player) (err error) {
	player.WriteUInt32(uint32(packet.X))
	player.WriteUInt32(uint32(packet.Z))
	player.WriteBool(packet.GroundUp)

	if player.protocol >= V1_16 {
		player.WriteBool(packet.ServerLightning)
	}

	var bitmask uint16 = 0
	for i, s := range packet.Sections {
		if s != nil &&
			!(s.Palette.GetSize() == 1 && s.Palette.RecoverName(0) == "minecraft:air") {
			bitmask |= 1 << i
		}
	}
	player.WriteVarInt(int(bitmask))

	// Heightmaps
	if player.protocol >= V1_14 {
		player.WriteNBTCompound(nbt.Compound{})
	}

	buff := newVarBuffer(16 * 16 * 16 * 16)
	tmp := player.io
	player.io = &ConnReadWrite{
		rdr: tmp.rdr,
		wtr: buff,
	}

	for i, s := range packet.Sections {
		// Check chunk used
		if (bitmask & (1 << i)) == 0 {
			continue
		}

		// Non air blocks
		if player.protocol >= V1_14 {
			player.WriteUInt16(4096)
		}

		// Calculate block bits size
		var bitsPerBlock uint8 = 4
		for s.Palette.GetSize() > 1<<int(bitsPerBlock) {
			bitsPerBlock += 1
		}
		if bitsPerBlock > 8 {
			bitsPerBlock = 13
		}
		player.WriteUInt8(bitsPerBlock)

		individualValueMask := uint64((1 << bitsPerBlock) - 1)

		// Send block palette
		if bitsPerBlock == 13 {
			player.WriteVarInt(0)
		} else {
			player.WriteVarInt(s.Palette.GetSize())

			for _, uid := range s.Palette.GetContent() {
				bId := BLOCK_REGISTRY.GetBlockId(uid, player.protocol)
				player.WriteVarInt(bId)
			}
		}

		// Content length
		length := (16 * 16 * 16) * int(bitsPerBlock) / 64
		player.WriteVarInt(length)

		// Calculate blocks
		data := make([]uint64, length)
		for index, block := range s.Blocks {
			var value uint64
			if bitsPerBlock == 13 {
				//TODO handle global palette
				value = 0
			} else {
				value = uint64(block)
			}

			bitIndex := index * int(bitsPerBlock)
			startLong := bitIndex / 64
			startOffset := bitIndex % 64
			endLong := ((index+1)*int(bitsPerBlock) - 1) / 64

			value &= individualValueMask

			data[startLong] |= value << startOffset

			if startLong != endLong {
				data[endLong] = value >> (64 - startOffset)
			}
		}

		// Write blocks
		for _, d := range data {
			player.WriteUInt64(d)
		}

		if player.protocol < V1_14 {
			// Write lights
			lights := make([]byte, 16*16*16/2)
			for i := range lights {
				lights[i] = 0xFF
			}
			player.WriteByteArray(lights)
			if packet.Dimension == OVERWORLD {
				// Overworld sky light
				player.WriteByteArray(lights)
			}
		}
	}

	if player.protocol >= V1_15 {
		player.io = tmp

		// Add biomes
		if packet.GroundUp {
			for i := 0; i < 1024; i++ {
				player.WriteUInt32(127)
			}
		}

		// Send sections
		player.WriteVarInt(buff.Len())
		player.WriteByteArray(buff.Bytes())
	} else {
		// Add biomes
		if packet.GroundUp {
			for i := 0; i < 256; i++ {
				player.WriteUInt32(127)
			}
		}

		// Send sections
		player.io = tmp
		player.WriteVarInt(buff.Len())
		player.WriteByteArray(buff.Bytes())
	}

	// No entity
	player.WriteVarInt(0)

	return
}
func (packet *PacketPlayChunkData) WriteV1_9(player *Player) (err error) {
	player.WriteUInt32(uint32(packet.X))
	player.WriteUInt32(uint32(packet.Z))
	player.WriteBool(packet.GroundUp)

	var bitmask uint16 = 0
	for i, s := range packet.Sections {
		if s != nil &&
			!(s.Palette.GetSize() == 1 && s.Palette.RecoverName(0) == "minecraft:air") {
			bitmask |= 1 << i
		}
	}
	player.WriteVarInt(int(bitmask))

	buff := newVarBuffer(16 * 16 * 16 * 16)
	tmp := player.io
	player.io = &ConnReadWrite{
		rdr: tmp.rdr,
		wtr: buff,
	}

	for i, s := range packet.Sections {
		// Check chunk used
		if (bitmask & (1 << i)) == 0 {
			continue
		}

		// Calculate block bits size
		var bitsPerBlock uint8 = 4
		for s.Palette.GetSize() > 1<<int(bitsPerBlock) {
			bitsPerBlock += 1
		}
		if bitsPerBlock > 8 {
			bitsPerBlock = 13
		}
		player.WriteUInt8(bitsPerBlock)

		individualValueMask := uint64((1 << bitsPerBlock) - 1)

		// Send block palette
		if bitsPerBlock == 13 {
			player.WriteVarInt(0)
		} else {
			player.WriteVarInt(s.Palette.GetSize())

			for _, uid := range s.Palette.GetContent() {
				bId := BLOCK_REGISTRY.GetBlockId(uid, player.protocol)
				fmt.Println(uid, bId)
				player.WriteVarInt(bId)
			}
		}

		// Content length
		length := (16 * 16 * 16) * int(bitsPerBlock) / 64
		player.WriteVarInt(length)

		// Calculate blocks
		data := make([]uint64, length)
		for index, block := range s.Blocks {
			var value uint64
			if bitsPerBlock == 13 {
				//TODO handle global palette
				value = 0
			} else {
				value = uint64(block)
			}

			bitIndex := index * int(bitsPerBlock)
			startLong := bitIndex / 64
			startOffset := bitIndex % 64
			endLong := ((index+1)*int(bitsPerBlock) - 1) / 64

			value &= individualValueMask

			data[startLong] |= value << startOffset

			if startLong != endLong {
				data[endLong] = value >> (64 - startOffset)
			}
		}

		// Write blocks
		for _, d := range data {
			player.WriteUInt64(d)
		}

		// Write lights
		lights := make([]byte, 16*16*16/2)
		for i := range lights {
			lights[i] = 0xFF
		}
		player.WriteByteArray(lights)
		if packet.Dimension == OVERWORLD {
			// Overworld sky light
			player.WriteByteArray(lights)
		}
	}

	// Add biomes
	if packet.GroundUp {
		biomes := make([]byte, 256)
		player.WriteByteArray(biomes)
	}

	// Send sections
	player.io = tmp
	player.WriteVarInt(buff.Len())
	player.WriteByteArray(buff.Bytes())

	// No entity
	if player.protocol >= V1_9_3 {
		player.WriteVarInt(0)
	}

	return
}
func (packet *PacketPlayChunkData) WriteV1_8(player *Player) (err error) {
	player.WriteUInt32(uint32(packet.X))
	player.WriteUInt32(uint32(packet.Z))
	player.WriteBool(packet.GroundUp)

	var bitmask uint16 = 0
	for i, s := range packet.Sections {
		if s != nil &&
			!(s.Palette.GetSize() == 1 && s.Palette.RecoverName(0) == "minecraft:air") {
			bitmask |= 1 << i
		}
	}
	player.WriteUInt16(bitmask)

	buff := newVarBuffer(16 * 16 * 16 * 16)
	tmp := player.io
	player.io = &ConnReadWrite{
		rdr: tmp.rdr,
		wtr: buff,
	}

	// Write blocks
	for i, s := range packet.Sections {
		// Check chunk used
		if (bitmask & (1 << i)) == 0 {
			continue
		}

		// Calculate blocks
		for _, block := range s.Blocks {
			value := uint16(BLOCK_REGISTRY.GetBlockId(s.Palette.RecoverName(block), player.protocol))
			player.WriteLittleEndianUInt16(value)
		}
	}

	player.WriteVarInt(16)
	player.WriteVarInt(16 * 16 * 16 * 2)

	// Write lights
	for i := range packet.Sections {
		// Check chunk used
		if (bitmask & (1 << i)) == 0 {
			continue
		}

		// Write lights
		lights := make([]byte, 16*16*16/2)
		for i := range lights {
			lights[i] = 0xFF
		}
		player.WriteByteArray(lights)
	}

	// Write sky lights
	if packet.Dimension == OVERWORLD {
		for i := range packet.Sections {
			// Check chunk used
			if (bitmask & (1 << i)) == 0 {
				continue
			}

			// Write sky lights
			lights := make([]byte, 16*16*16/2)
			for i := range lights {
				lights[i] = 0xFF
			}
			player.WriteByteArray(lights)
		}
	}

	// Write biomes
	if packet.GroundUp {
		biomes := make([]byte, 256)
		player.WriteByteArray(biomes)
	}

	// Send sections
	player.io = tmp
	player.WriteVarInt(buff.Len())
	player.WriteByteArray(buff.Bytes())

	return
}
func (packet *PacketPlayChunkData) WriteV1_7(player *Player) (err error) {
	player.WriteUInt32(uint32(packet.X))
	player.WriteUInt32(uint32(packet.Z))
	player.WriteBool(packet.GroundUp)

	var bitmask uint16 = 0
	for i, s := range packet.Sections {
		if s != nil &&
			!(s.Palette.GetSize() == 1 && s.Palette.RecoverName(0) == "minecraft:air") {
			bitmask |= 1 << i
		}
	}
	player.WriteUInt16(bitmask)

	// Add bitmask ???
	player.WriteUInt16(0)

	buff := newVarBuffer(16 * 16 * 16 * 16)
	zwtr := zlib.NewWriter(buff)
	tmp := player.io
	player.io = &ConnReadWrite{
		rdr: tmp.rdr,
		wtr: zwtr,
	}

	// Write blocks
	for i, s := range packet.Sections {
		// Check chunk used
		if (bitmask & (1 << i)) == 0 {
			continue
		}

		// Calculate blocks
		for _, block := range s.Blocks {
			id, _ := BLOCK_REGISTRY.GetLegacyBlockTypeData(s.Palette.RecoverName(block), player.protocol)
			player.WriteUInt8(uint8(id))
		}
	}

	// Write metadata
	for i, s := range packet.Sections {
		// Check chunk used
		if (bitmask & (1 << i)) == 0 {
			continue
		}

		// Calculate blocks data
		for i := 0; i < len(s.Blocks); i += 2 {
			_, data1 := BLOCK_REGISTRY.GetLegacyBlockTypeData(s.Palette.RecoverName(s.Blocks[i]), player.protocol)
			_, data2 := BLOCK_REGISTRY.GetLegacyBlockTypeData(s.Palette.RecoverName(s.Blocks[i+1]), player.protocol)
			player.WriteUInt8(uint8(data1) | (uint8(data2) << 4))
		}
	}

	// Write lights
	for i := range packet.Sections {
		// Check chunk used
		if (bitmask & (1 << i)) == 0 {
			continue
		}

		// Write lights
		lights := make([]byte, 16*16*16/2)
		for i := range lights {
			lights[i] = 0xFF
		}
		player.WriteByteArray(lights)
	}

	// Write sky lights
	if packet.Dimension == OVERWORLD {
		for i := range packet.Sections {
			// Check chunk used
			if (bitmask & (1 << i)) == 0 {
				continue
			}

			// Write sky lights
			lights := make([]byte, 16*16*16/2)
			for i := range lights {
				lights[i] = 0xFF
			}
			player.WriteByteArray(lights)
		}
	}

	// Write biomes
	if packet.GroundUp {
		biomes := make([]byte, 256)
		player.WriteByteArray(biomes)
	}

	// Send sections
	player.io = tmp
	zwtr.Close()
	player.WriteUInt32(uint32(buff.Len()))
	player.WriteByteArray(buff.Bytes())

	return
}
func (packet *PacketPlayChunkData) Handle(player *Player) {
}
func (packet *PacketPlayChunkData) Id() (int, Protocol) {
	return 0x20, V1_10
}

type PacketPlayParticle struct {
	Type         int
	LongDistance bool
	X            float32
	Y            float32
	Z            float32
	OffsetX      float32
	OffsetY      float32
	OffsetZ      float32
	ParticleData float32
	Count        int
	Data         []int
}

func (packet *PacketPlayParticle) Read(player *Player, length int) (err error) {
	return
}
func (packet *PacketPlayParticle) Write(player *Player) (err error) {
	err = player.WriteUInt32(uint32(packet.Type))
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	err = player.WriteBool(packet.LongDistance)
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	err = player.WriteFloat32(packet.X)
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	err = player.WriteFloat32(packet.Y)
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	err = player.WriteFloat32(packet.Z)
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	err = player.WriteFloat32(packet.OffsetX)
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	err = player.WriteFloat32(packet.OffsetY)
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	err = player.WriteFloat32(packet.OffsetZ)
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	err = player.WriteFloat32(packet.ParticleData)
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	err = player.WriteUInt32(uint32(packet.Count))
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	return
}
func (packet *PacketPlayParticle) Handle(player *Player) {}
func (packet *PacketPlayParticle) Id() (int, Protocol) {
	return 0x22, V1_10
}

type PacketPlayJoinGame struct {
	EntityId            uint32
	Gamemode            Gamemode
	Dimension           Dimension
	HashedSeed          uint64
	Difficulty          Difficulty
	MaxPlayers          uint8
	LevelType           LevelType
	ReducedDebug        bool
	EnableRespawnScreen bool
	IsDebug             bool
	IsFlat              bool
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
		slog.Debug("packet error", "error", err)
		return
	}
	err = player.WriteUInt8(uint8(packet.Gamemode))
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	if player.protocol >= V1_16 {
		err = player.WriteUInt8(uint8(packet.Gamemode))
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
	}
	if player.protocol < V1_16 {
		err = player.WriteUInt32(uint32(packet.Dimension.Id))
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
	} else {
		err = player.WriteVarInt(1)
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
		err = player.WriteString(packet.Dimension.Name)
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}

		compound := nbt.Compound{
			"dimension": nbt.MakeCompoundList([]nbt.Compound{
				*packet.Dimension.Entry(),
			}),
		}

		err = player.WriteNBTCompound(compound)
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
		err = player.WriteString(packet.Dimension.String())
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
		err = player.WriteString("minecraft:overworld")
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
	}
	if player.protocol < V1_14 {
		err = player.WriteUInt8(uint8(packet.Difficulty))
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
	}
	if player.protocol >= V1_15 {
		err = player.WriteUInt64(packet.HashedSeed)
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
	}
	err = player.WriteUInt8(packet.MaxPlayers)
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	if player.protocol < V1_16 {
		err = player.WriteString(string(packet.LevelType))
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
	}
	if player.protocol >= V1_14 {
		err = player.WriteVarInt(32)
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
	}
	if player.protocol > V1_7_6 {
		err = player.WriteBool(packet.ReducedDebug)
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
	}
	if player.protocol >= V1_15 {
		err = player.WriteBool(packet.EnableRespawnScreen)
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
	}
	if player.protocol >= V1_16 {
		err = player.WriteBool(packet.IsDebug)
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
		err = player.WriteBool(packet.IsFlat)
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
	}
	return
}
func (packet *PacketPlayJoinGame) Handle(player *Player) {}
func (packet *PacketPlayJoinGame) Id() (int, Protocol) {
	return 0x23, V1_10
}

type PacketPlayPlayerAbilities struct {
	Invulnerable bool
	Fly          bool
	CanFly       bool
	Creative     bool
	FlyingSpeed  float32
	FieldOfView  float32
}

func (packet *PacketPlayPlayerAbilities) Read(player *Player, length int) (err error) {
	return
}
func (packet *PacketPlayPlayerAbilities) Write(player *Player) (err error) {
	var flags uint8 = 0
	if packet.Invulnerable {
		flags |= 0x01
	}
	if packet.Fly {
		flags |= 0x02
	}
	if packet.CanFly {
		flags |= 0x04
	}
	if packet.Creative {
		flags |= 0x08
	}

	err = player.WriteUInt8(flags)
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	err = player.WriteFloat32(packet.FlyingSpeed)
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	err = player.WriteFloat32(packet.FieldOfView)
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	return
}
func (packet *PacketPlayPlayerAbilities) Handle(player *Player) {}
func (packet *PacketPlayPlayerAbilities) Id() (int, Protocol) {
	return 0x2B, V1_10
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
		slog.Debug("packet error", "error", err)
		return
	}
	err = player.WriteFloat64(packet.Y)
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	err = player.WriteFloat64(packet.Z)
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	err = player.WriteFloat32(packet.Yaw)
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	err = player.WriteFloat32(packet.Pitch)
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	if player.protocol == V1_9 {
		// Send on ground
		err = player.WriteBool(true)
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
	} else {
		err = player.WriteUInt8(packet.Flags)
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
		if player.protocol >= V1_9_3 {
			err = player.WriteVarInt(packet.TeleportId)
			if err != nil {
				slog.Debug("packet error", "error", err)
				return
			}
		}
	}
	return
}
func (packet *PacketPlayerPositionLook) Handle(player *Player) {}
func (packet *PacketPlayerPositionLook) Id() (int, Protocol) {
	return 0x2E, V1_10
}

type PacketUpdateHealth struct {
	Health         float32
	Food           int
	FoodSaturation float32
}

func (packet *PacketUpdateHealth) Read(player *Player, length int) (err error) {
	return
}
func (packet *PacketUpdateHealth) Write(player *Player) (err error) {
	err = player.WriteFloat32(packet.Health)
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	err = player.WriteVarInt(packet.Food)
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	err = player.WriteFloat32(packet.FoodSaturation)
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	return
}
func (packet *PacketUpdateHealth) Handle(player *Player) {}
func (packet *PacketUpdateHealth) Id() (int, Protocol) {
	return 0x3E, V1_10
}

type PacketPlaySpawnPosition struct {
	Position Position
}

func (packet *PacketPlaySpawnPosition) Read(player *Player, length int) (err error) {
	return
}
func (packet *PacketPlaySpawnPosition) Write(player *Player) (err error) {
	if player.protocol >= V1_8 {
		err = player.WritePosition(packet.Position)
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
	} else {
		err = player.WriteUInt32(uint32(packet.Position.X))
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
		err = player.WriteUInt32(uint32(packet.Position.Y))
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
		err = player.WriteUInt32(uint32(packet.Position.Z))
		if err != nil {
			slog.Debug("packet error", "error", err)
			return
		}
	}
	return
}
func (packet *PacketPlaySpawnPosition) Handle(player *Player) {}
func (packet *PacketPlaySpawnPosition) Id() (int, Protocol) {
	return 0x43, V1_10
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
		slog.Debug("packet error", "error", err)
		return
	}
	if packet.Footer == nil {
		str = `{"translate":""}`
	} else {
		str = *packet.Footer
	}
	err = player.WriteString(str)
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	return
}
func (packet *PacketPlayerListHeaderFooter) Handle(player *Player) {}
func (packet *PacketPlayerListHeaderFooter) Id() (int, Protocol) {
	return 0x47, V1_10
}

type PacketPlayPlayerPosition struct {
	X        float64
	Y        float64
	Z        float64
	Yaw      float32
	Pitch    float32
	OnGround bool
}

func (packet *PacketPlayPlayerPosition) Read(player *Player, length int) (err error) {
	packet.X, err = player.ReadFloat64()
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	packet.Y, err = player.ReadFloat64()
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	packet.Z, err = player.ReadFloat64()
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	packet.Yaw, err = player.ReadFloat32()
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	packet.Pitch, err = player.ReadFloat32()
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	packet.OnGround, err = player.ReadBool()
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}

	return
}
func (packet *PacketPlayPlayerPosition) Write(player *Player) (err error) {
	return
}
func (packet *PacketPlayPlayerPosition) Handle(player *Player) {
	from := player.GetLocation()
	to := Location{
		X:     packet.X,
		Y:     packet.Y,
		Z:     packet.Z,
		Yaw:   packet.Yaw,
		Pitch: packet.Pitch,
	}

	event := &PlayerMoveEvent{
		Player: player,
		From:   from,
		To:     to,
	}
	player.core.CallEvent(event)

	if !event.IsCancelled() {
		player.SetLocation(to)
	}
}
func (packet *PacketPlayPlayerPosition) Id() (int, Protocol) {
	return 0x0D, V1_10
}

type PacketPlayPlayerLook struct {
	Yaw      float32
	Pitch    float32
	OnGround bool
}

func (packet *PacketPlayPlayerLook) Read(player *Player, length int) (err error) {
	packet.Yaw, err = player.ReadFloat32()
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	packet.Pitch, err = player.ReadFloat32()
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}
	packet.OnGround, err = player.ReadBool()
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}

	return
}
func (packet *PacketPlayPlayerLook) Write(player *Player) (err error) {
	return
}
func (packet *PacketPlayPlayerLook) Handle(player *Player) {
	from := player.GetLocation()
	to := Location{
		X:     from.X,
		Y:     from.Y,
		Z:     from.Z,
		Yaw:   packet.Yaw,
		Pitch: packet.Pitch,
	}

	event := &PlayerMoveEvent{
		Player: player,
		From:   from,
		To:     to,
	}
	player.core.CallEvent(event)

	if !event.IsCancelled() {
		player.SetLocation(to)
		player.SetOnGround(packet.OnGround)
	}
}
func (packet *PacketPlayPlayerLook) Id() (int, Protocol) {
	return 0x0E, V1_10
}

type PacketPlayPlayer struct {
	OnGround bool
}

func (packet *PacketPlayPlayer) Read(player *Player, length int) (err error) {
	packet.OnGround, err = player.ReadBool()
	if err != nil {
		slog.Debug("packet error", "error", err)
		return
	}

	return
}
func (packet *PacketPlayPlayer) Write(player *Player) (err error) {
	return
}
func (packet *PacketPlayPlayer) Handle(player *Player) {
	player.SetOnGround(packet.OnGround)
}
func (packet *PacketPlayPlayer) Id() (int, Protocol) {
	return 0x0F, V1_10
}

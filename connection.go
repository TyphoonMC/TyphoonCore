package typhoon

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	uuid2 "github.com/TyphoonMC/go.uuid"
	"github.com/seebs/nbt"
	"log"
	"net"
	"reflect"
)

type State int8

const (
	HANDSHAKING State = iota
	STATUS
	LOGIN
	PLAY
)

type Gamemode uint8

const (
	SURVIVAL Gamemode = iota
	CREATIVE
	ADVENTURE
	SPECTATOR
)

type Dimension struct {
	Name               string
	Id                 uint8
	Natural            uint8
	FixedTime          *uint64
	AmbientLight       float32
	Shrunk             uint8
	Ultrawarm          uint8
	HasCeiling         uint8
	HasSkylight        uint8
	Infiniburn         string
	LogicalHeight      uint32
	HasRaids           uint8
	RespawnAnchorWorks uint8
	BedWorks           uint8
	PiglinSafe         uint8
}

var (
	NETHER = Dimension{
		"minecraft:the_nether",
		0xFF,
		0,
		uint64Ptr(18000),
		0.1,
		1,
		1,
		1,
		0,
		"",
		256,
		0,
		1,
		0,
		0,
	}
	OVERWORLD = Dimension{
		"minecraft:overworld",
		0x00,
		1,
		nil,
		0,
		0,
		0,
		0,
		1,
		"",
		256,
		0,
		0,
		1,
		1,
	}
	END = Dimension{
		"minecraft:the_end",
		0x01,
		0,
		uint64Ptr(6000),
		0,
		0,
		0,
		0,
		0,
		"",
		0,
		1,
		0,
		0,
		1,
	}
)

func (dimension Dimension) String() string {
	return dimension.Name
}

func (dimension Dimension) Entry() *nbt.Compound {
	element := nbt.Compound{
		"name":                 nbt.String(dimension.String()),
		"natural":              nbt.Byte(dimension.Natural),
		"ambient_light":        nbt.Float(dimension.AmbientLight),
		"shrunk":               nbt.Byte(dimension.Shrunk),
		"ultrawarm":            nbt.Byte(dimension.Ultrawarm),
		"has_ceiling":          nbt.Byte(dimension.HasCeiling),
		"has_skylight":         nbt.Byte(dimension.HasSkylight),
		"infiniburn":           nbt.String(dimension.Infiniburn),
		"logical_height":       nbt.Int(dimension.LogicalHeight),
		"has_raids":            nbt.Byte(dimension.HasRaids),
		"respawn_anchor_works": nbt.Byte(dimension.RespawnAnchorWorks),
		"bed_works":            nbt.Byte(dimension.BedWorks),
		"piglin_safe":          nbt.Byte(dimension.PiglinSafe),
	}

	if dimension.FixedTime != nil {
		element["fixed_time"] = nbt.Long(*dimension.FixedTime)
	}
	return &element
}

type Difficulty uint8

const (
	PEACEFUL Difficulty = iota
	EASY
	NORMAL
	HARD
)

type ChatPosition uint8

const (
	CHAT_BOX ChatPosition = iota
	SYSTEM
	ACTION_BAR
)

type ScoreboardPosition uint8

const (
	LIST ScoreboardPosition = iota
	SIDEBAR
	BELOW_NAME
)

type BossBarAction int

const (
	BOSSBAR_ADD BossBarAction = iota
	BOSSBAR_REMOVE
	BOSSBAR_UPDATE_HEALTH
	BOSSBAR_UPDATE_TITLE
	BOSSBAR_UPDATE_STYLE
	BOSSBAR_UPDATE_FLAGS
)

type BossBarColor int

const (
	BOSSBAR_COLOR_PINK BossBarColor = iota
	BOSSBAR_COLOR_BLUE
	BOSSBAR_COLOR_RED
	BOSSBAR_COLOR_GREEN
	BOSSBAR_COLOR_YELLOW
	BOSSBAR_COLOR_PURPLE
	BOSSBAR_COLOR_WHITE
)

type BossBarDivision int

const (
	BOSSBAR_NODIVISION BossBarDivision = iota
	BOSSBAR_6NOTCHES
	BOSSBAR_10NOTCHES
	BOSSBAR_12NOTCHES
	BOSSBAR_20NOTCHES
)

type LevelType string

const (
	DEFAULT      LevelType = "default"
	FLAT         LevelType = "flat"
	LARGE_BIOMES LevelType = "largeBiomes"
	AMPLIFIED    LevelType = "amplified"
	DEFAULT_1_1  LevelType = "default_1_1"
)

type ClientStatusAction uint8

const (
	PERFORM_RESPAWN ClientStatusAction = iota
	REQUEST_STATS
	OPEN_INVENTORY
)

type Protocol uint16

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

type Location struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

func (loc *Location) ToPosition() *Position {
	return &Position{
		int(loc.X),
		int(loc.Y),
		int(loc.Z),
	}
}

const (
	V1_7_2  Protocol = 4
	V1_7_6  Protocol = 5
	V1_8    Protocol = 47
	V1_9    Protocol = 107
	V1_9_1  Protocol = 108
	V1_9_2  Protocol = 109
	V1_9_3  Protocol = 110
	V1_10   Protocol = 210
	V1_11   Protocol = 315
	V1_11_1 Protocol = 316
	V1_12   Protocol = 335
	V1_12_1 Protocol = 338
	V1_12_2 Protocol = 340
	V1_13   Protocol = 393
	V1_13_1 Protocol = 401
	V1_13_2 Protocol = 404
	V1_14   Protocol = 477
	V1_14_1 Protocol = 480
	V1_14_2 Protocol = 485
	V1_14_3 Protocol = 490
	V1_14_4 Protocol = 498
	V1_15   Protocol = 573
	V1_15_1 Protocol = 575
	V1_15_2 Protocol = 578
	V1_16   Protocol = 727
)

var (
	COMPATIBLE_PROTO = []Protocol{
		V1_7_2, V1_7_6,
		V1_8,
		V1_9, V1_9_1, V1_9_2, V1_9_3,
		V1_10,
		V1_11, V1_11_1,
		V1_12, V1_12_1,
		V1_12_2,
		V1_13,
		V1_13_1,
		V1_13_2,
		V1_14,
		V1_14_1,
		V1_14_2,
		V1_14_3,
		V1_14_4,
		V1_15,
		V1_15_1,
		V1_15_2,
		V1_16,
	}
)

func IsCompatible(proto Protocol) bool {
	for _, x := range COMPATIBLE_PROTO {
		if x == proto {
			return true
		}
	}
	return false
}

func registerProtocol(proto Protocol) {
	COMPATIBLE_PROTO = append(COMPATIBLE_PROTO, proto)
}

type InAddr struct {
	address string
	port    uint16
}

type Player struct {
	core         *Core
	id           int
	conn         net.Conn
	io           *ConnReadWrite
	state        State
	protocol     Protocol
	inaddr       InAddr
	name         string
	uuid         uuid2.UUID
	keepalive    int
	compression  bool
	packetsQueue chan Packet
}

func (player *Player) GetName() string {
	return player.name
}

func (player *Player) GetUUID() uuid2.UUID {
	return player.uuid
}

func (player *Player) ReadPacket() (packet Packet, err error) {
	if !player.compression {
		return player.ReadPacketWithoutCompression()
	} else {
		return player.ReadPacketWithCompression()
	}
}

func (player *Player) ReadPacketWithoutCompression() (packet Packet, err error) {
	length, err := player.ReadVarInt()
	if err != nil {
		return
	}

	id, err := player.ReadVarInt()
	if err != nil {
		return
	}
	if player.state == PLAY {
		id = player.HackServerbound(id)
	}

	packet, err = player.HandlePacket(id, length)
	if err != nil {
		return
	} else if packet != nil {
		if config.Logs {
			log.Printf("#%d -> %d %s %s", player.id, id, reflect.TypeOf(packet), fmt.Sprint(packet))
		}
		packet.Handle(player)
	}
	return
}

func (player *Player) ReadPacketWithCompression() (packet Packet, err error) {
	packetLength, err := player.ReadVarInt()
	if err != nil {
		return
	}

	dataLength, err := player.ReadVarInt()
	if err != nil {
		return
	}
	dataLengthLength := binary.PutUvarint(player.io.buffer[:], uint64(dataLength))

	var id int
	var length int

	if dataLength == 0 {
		id, err = player.ReadVarInt()
		if err != nil {
			return
		}
		idLength := binary.PutUvarint(player.io.buffer[:], uint64(id))
		length = packetLength - dataLengthLength - idLength

		if player.state == PLAY {
			id = player.HackServerbound(id)
		}
		packet, err = player.HandlePacket(id, length)

		if err != nil {
			return
		} else if packet != nil {
			if config.Logs {
				log.Printf("#%d u-> %d %s %s", player.id, id, reflect.TypeOf(packet), fmt.Sprint(packet))
			}
			packet.Handle(player)
		}
	} else {
		var compressed []byte = make([]byte, packetLength-dataLengthLength)
		player.conn.Read(compressed[:])

		r, err := zlib.NewReader(bytes.NewReader(compressed[:]))
		if err != nil {
			return nil, err
		}

		var uncompressed []byte = make([]byte, packetLength-dataLengthLength)
		r.Read(uncompressed[:])
		r.Close()

		tmp := player.io
		rdr := bytes.NewBuffer(uncompressed[:])
		player.io.rdr = rdr

		id, err = player.ReadVarInt()
		if err != nil {
			return nil, err
		}

		if player.state == PLAY {
			id = player.HackServerbound(id)
		}
		packet, err = player.HandlePacket(id, length)
		player.io = tmp

		if err != nil {
			return nil, err
		} else if packet != nil {
			if config.Logs {
				log.Printf("#%d c-> %d %s %s", player.id, id, reflect.TypeOf(packet), fmt.Sprint(packet))
			}
			packet.Handle(player)
		}
	}
	return
}

func (player *Player) WritePacket(packet Packet) (err error) {
	player.packetsQueue <- packet
	return nil
}

func (player *Player) privateWritePacket(packet Packet) (err error) {
	if !player.compression {
		return player.WritePacketWithoutCompression(packet)
	} else {
		return player.WritePacketWithCompression(packet)
	}
}

func (player *Player) WritePacketWithoutCompression(packet Packet) (err error) {
	buff := newVarBuffer(256)
	tmp := player.io
	player.io = &ConnReadWrite{
		rdr: tmp.rdr,
		wtr: buff,
	}

	id, proto := packet.Id()
	if player.state == PLAY {
		id = player.HackClientbound(id, proto)
	}
	if id == -1 {
		return
	}
	player.WriteVarInt(id)
	packet.Write(player)

	ln := newVarBuffer(0)
	player.io.wtr = ln
	player.WriteVarInt(buff.Len())
	player.io = tmp
	player.conn.Write(ln.Bytes())
	player.conn.Write(buff.Bytes())

	if config.Logs {
		log.Printf("#%d <- %d %s %s", player.id, id, reflect.TypeOf(packet), fmt.Sprint(packet))
	}
	return nil
}

func (player *Player) WritePacketWithCompression(packet Packet) (err error) {
	buff := newVarBuffer(256)
	tmp := player.io
	player.io = &ConnReadWrite{
		rdr: tmp.rdr,
		wtr: buff,
	}

	id, proto := packet.Id()
	if player.state == PLAY {
		id = player.HackClientbound(id, proto)
	}
	if id == -1 {
		return
	}
	player.WriteVarInt(id)
	packet.Write(player)

	var rBuff []byte
	var dataLength = 0
	if buff.Len() < config.Threshold {
		rBuff = buff.Bytes()
	} else {
		var b bytes.Buffer
		w := zlib.NewWriter(&b)
		w.Write(buff.Bytes())
		w.Close()
		rBuff = b.Bytes()
		dataLength = len(rBuff)
	}

	buff2 := newVarBuffer(1)
	player.io.wtr = buff2
	player.WriteVarInt(dataLength)
	packetLength := len(rBuff) + buff2.Len()

	buff3 := newVarBuffer(1)
	player.io.wtr = buff3
	player.WriteVarInt(packetLength)

	player.io = tmp
	player.conn.Write(buff3.Bytes())
	player.conn.Write(buff2.Bytes())
	player.conn.Write(rBuff)

	if config.Logs {
		if buff.Len() < config.Threshold {
			log.Println("<-u", id, reflect.TypeOf(packet), packet)
		} else {
			log.Println("<-c", id, reflect.TypeOf(packet), packet)
		}
	}
	return nil
}

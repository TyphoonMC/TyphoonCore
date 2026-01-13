package typhoon

import (
	"bufio"
	"log/slog"
	"math/rand"
	"net"
	"os"
	"reflect"
	"time"

	"github.com/google/uuid"
)

type Core struct {
	connCounter      int
	eventHandlers    map[reflect.Type][]EventCallback
	brand            string
	rootCommand      CommandNode
	compiledCommands []commandNode
	playerRegistry   *PlayerRegistry
	world            *Map
	gamemode         Gamemode
	difficulty       Difficulty
}

func Init() *Core {
	if err := initConfig(); err != nil {
		slog.Error("failed to initialize config", "error", err)
		os.Exit(1)
	}
	initPackets()
	initHacks()
	c := &Core{
		0,
		make(map[reflect.Type][]EventCallback),
		"typhoon",
		CommandNode{
			commandNodeTypeRoot,
			nil,
			nil,
			nil,
			"",
			nil,
		},
		nil,
		newPlayerRegistry(),
		&Map{
			Location{X: 0, Y: 0, Z: 0},
			END,
			[]*Chunk{},
		},
		SPECTATOR,
		PEACEFUL,
	}
	c.compileCommands()
	return c
}

func (c *Core) Start() {
	ln, err := net.Listen("tcp", config.ListenAddress)
	if err != nil {
		slog.Error("failed to start server", "error", err)
		os.Exit(1)
	}
	slog.Info("server launched", "address", config.ListenAddress)
	go c.keepAlive()
	for {
		conn, err := ln.Accept()
		if err != nil {
			slog.Error("failed to accept connection", "error", err)
		} else {
			c.connCounter += 1
			go c.handleConnection(conn, c.connCounter)
		}
	}
}

func (c *Core) SetMap(world *Map) {
	c.world = world
}

func (c *Core) SetGamemode(gamemode Gamemode) {
	c.gamemode = gamemode
}

func (c *Core) SetBrand(brand string) {
	br := make([]byte, len(brand)+1)
	copy(br[:len(brand)], []byte(brand))
	c.brand = string(br)
}

func (c *Core) GetPlayerRegistry() *PlayerRegistry {
	return c.playerRegistry
}

func (c *Core) keepAlive() {
	r := rand.New(rand.NewSource(15768735131534))
	keepalive := &PacketPlayKeepAlive{
		Identifier: 0,
	}
	for {
		c.playerRegistry.ForEachPlayer(func(player *Player) {
			if player.state == PLAY {
				//TODO rework keepalive
				/*if player.keepalive != 0 {
					player.Kick("Timed out")
				}*/

				id := int(r.Int31())
				keepalive.Identifier = id
				player.keepalive = id
				player.WritePacket(keepalive)
			}
		})
		time.Sleep(5000000000)
	}
}

func (c *Core) handleConnection(conn net.Conn, id int) {
	slog.Debug("player connected", "addr", conn.RemoteAddr().String(), "id", id)

	defaultUUID, _ := uuid.Parse("d979912c-bb24-4f23-a6ac-c32985a1e5d3")
	player := &Player{
		core:     c,
		id:       id,
		conn:     conn,
		state:    HANDSHAKING,
		protocol: V1_10,
		io: &ConnReadWrite{
			rdr: bufio.NewReader(conn),
			wtr: bufio.NewWriter(conn),
		},
		inaddr: InAddr{
			"",
			0,
		},
		name:         "",
		uuid:         defaultUUID,
		keepalive:    0,
		compression:  false,
		packetsQueue: make(chan Packet),
	}

	go func() {
		for {
			packet := <-player.packetsQueue
			err := player.privateWritePacket(packet)
			if err != nil {
				break
			}
		}
	}()

	for {
		_, err := player.ReadPacket()
		if err != nil {
			break
		}
	}

	if player.state == PLAY {
		player.core.CallEvent(&PlayerQuitEvent{Player: player})
		player.unregister()
	}
	conn.Close()
	slog.Debug("player disconnected", "addr", conn.RemoteAddr().String(), "id", id)
}

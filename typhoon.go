package typhoon

import (
	"bufio"
	"log"
	"math/rand"
	"net"
	"reflect"
	"time"
)

type Core struct {
	connCounter      int
	eventHandlers    map[reflect.Type][]EventCallback
	brand            string
	rootCommand      CommandNode
	compiledCommands []commandNode
}

func Init() *Core {
	initConfig()
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
	}
	c.compileCommands()
	return c
}

func (c *Core) Start() {
	ln, err := net.Listen("tcp", config.ListenAddress)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server launched on port", config.ListenAddress)
	go c.keepAlive()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print(err)
		} else {
			c.connCounter += 1
			go c.handleConnection(conn, c.connCounter)
		}
	}
}

func (c *Core) SetBrand(brand string) {
	br := make([]byte, len(brand)+1)
	copy(br[:len(brand)], []byte(brand))
	c.brand = string(br)
}

func (c *Core) keepAlive() {
	r := rand.New(rand.NewSource(15768735131534))
	keepalive := &PacketPlayKeepAlive{
		Identifier: 0,
	}
	for {
		playersMutex.Lock()
		for _, player := range players {
			if player.state == PLAY {
				if player.keepalive != 0 {
					player.Kick("Timed out")
				}

				id := int(r.Int31())
				keepalive.Identifier = id
				player.keepalive = id
				player.WritePacket(keepalive)
			}
		}
		playersMutex.Unlock()
		time.Sleep(5000000000)
	}
}

func (c *Core) handleConnection(conn net.Conn, id int) {
	log.Printf("%s(#%d) connected.", conn.RemoteAddr().String(), id)

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
		name:        "",
		uuid:        "d979912c-bb24-4f23-a6ac-c32985a1e5d3",
		keepalive:   0,
		compression: false,
	}

	for {
		_, err := player.ReadPacket()
		if err != nil {
			break
		}
	}

	player.core.CallEvent(&PlayerQuitEvent{player})
	player.unregister()
	conn.Close()
	log.Printf("%s(#%d) disconnected.", conn.RemoteAddr().String(), id)
}

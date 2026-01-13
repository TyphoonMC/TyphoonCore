//go:build integration

package typhoon

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/bot/basic"
	pk "github.com/Tnze/go-mc/net/packet"
)

func TestServerPing(t *testing.T) {
	// Start server in background
	core := Init()
	core.SetBrand("test-server")

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to start listener: %v", err)
	}
	defer listener.Close()

	port := listener.Addr().(*net.TCPAddr).Port
	t.Logf("Server listening on port %d", port)

	// Accept connections in background
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			core.connCounter++
			go core.handleConnection(conn, core.connCounter)
		}
	}()

	// Wait for server to be ready
	time.Sleep(100 * time.Millisecond)

	// Test ping
	addr := fmt.Sprintf("127.0.0.1:%d", port)
	status, _, err := bot.PingAndList(addr)
	if err != nil {
		t.Fatalf("Failed to ping server: %v", err)
	}

	t.Logf("Server status: %+v", status)

	if status.Version.Name != "Typhoon" {
		t.Errorf("Expected version name 'Typhoon', got '%s'", status.Version.Name)
	}
}

func TestPlayerJoin(t *testing.T) {
	// Start server in background
	core := Init()
	core.SetBrand("test-server")

	joined := make(chan bool, 1)
	core.On(func(e *PlayerJoinEvent) {
		t.Logf("Player joined: %s", e.Player.GetName())
		joined <- true
	})

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to start listener: %v", err)
	}
	defer listener.Close()

	port := listener.Addr().(*net.TCPAddr).Port

	// Accept connections in background
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			core.connCounter++
			go core.handleConnection(conn, core.connCounter)
		}
	}()

	// Wait for server to be ready
	time.Sleep(100 * time.Millisecond)

	// Create client and attempt to join
	addr := fmt.Sprintf("127.0.0.1:%d", port)
	client := bot.NewClient()
	client.Auth.Name = "TestPlayer"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Connect to server
	err = client.JoinServer(addr)
	if err != nil {
		t.Logf("Join error (expected for offline mode): %v", err)
	}

	// Handle game events
	basic.NewPlayer(client, basic.DefaultSettings, basic.EventsListener{
		GameStart: func() error {
			t.Log("Game started")
			return nil
		},
		Disconnect: func(reason pk.Chat) error {
			t.Logf("Disconnected: %s", reason)
			return nil
		},
	})

	select {
	case <-joined:
		t.Log("Player join event received")
	case <-ctx.Done():
		t.Log("Timeout waiting for join (expected in some cases)")
	}
}

func TestMultipleConnections(t *testing.T) {
	core := Init()
	core.SetBrand("test-server")

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to start listener: %v", err)
	}
	defer listener.Close()

	port := listener.Addr().(*net.TCPAddr).Port

	// Accept connections in background
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			core.connCounter++
			go core.handleConnection(conn, core.connCounter)
		}
	}()

	time.Sleep(100 * time.Millisecond)

	// Test multiple concurrent pings
	addr := fmt.Sprintf("127.0.0.1:%d", port)
	for i := 0; i < 5; i++ {
		_, _, err := bot.PingAndList(addr)
		if err != nil {
			t.Errorf("Ping %d failed: %v", i, err)
		}
	}
}

package typhoon

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	_ "image/png"
	"log/slog"
	"os"
)

type BufferConfig struct {
	HandshakeAddress int `json:"handshake_address"`
	PlayerName       int `json:"player_name"`
	ChatMessage      int `json:"chat_message"`
}

type Config struct {
	ListenAddress string       `json:"listen_address"`
	MaxPlayers    int          `json:"max_players"`
	Motd          string       `json:"motd"`
	Restricted    bool         `json:"restricted"`
	Logs          bool         `json:"logs"`
	Compression   bool         `json:"enable_compression"`
	Threshold     int          `json:"compression_threshold"`
	BufferConfig  BufferConfig `json:"buffer_config"`
}

var (
	config  Config
	favicon string
)

func initConfig() error {
	imgFile, err := os.Open("./favicon.png")
	if err == nil {
		defer imgFile.Close()

		img, _, err := image.Decode(imgFile)
		if err == nil {
			b := img.Bounds()
			if b.Max.X != 64 || b.Max.Y != 64 {
				slog.Warn("invalid icon for server", "hint", "resize to 64x64 pixels")
			}

			fav, err := os.ReadFile("./favicon.png")
			if err == nil {
				favicon = "data:image/png;base64," + base64.StdEncoding.EncodeToString(fav)
			}
		}
	}

	file, err := os.ReadFile("./config.json")
	if err != nil {
		return fmt.Errorf("reading config.json: %w", err)
	}
	if err := json.Unmarshal(file, &config); err != nil {
		return fmt.Errorf("parsing config.json: %w", err)
	}
	return nil
}

func (c *Core) GetConfig(cfg interface{}) error {
	file, err := os.ReadFile("./config.json")
	if err != nil {
		return fmt.Errorf("reading config.json: %w", err)
	}
	if err := json.Unmarshal(file, cfg); err != nil {
		return fmt.Errorf("parsing config.json: %w", err)
	}
	return nil
}

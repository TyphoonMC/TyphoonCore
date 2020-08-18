package typhoon

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type HackMap struct {
	Clientbound map[string]string `json:"clientbound"`
	Serverbound map[string]string `json:"serverbound"`
}

type HackContent struct {
	Name     string   `json:"name"`
	Protocol Protocol `json:"protocol"`
	Base     Protocol `json:"base"`
	Map      HackMap  `json:"map"`
}

type HackType struct {
	Name    string `json:"name"`
	Version int    `json:"version"`
}

type HackModule struct {
	Type    HackType    `json:"type"`
	Content HackContent `json:"content"`
}

var (
	clientbound = make(map[Protocol]map[int]int)
	serverbound = make(map[Protocol]map[int]int)
)

func initHacks() {
	// Hack 1.8
	clientbound[V1_8] = make(map[int]int)
	clientbound[V1_8][0x00] = 0x0E
	clientbound[V1_8][0x01] = 0x11
	//clientbound[V1_8][0x02] = 0x2C
	clientbound[V1_8][0x03] = 0x0F
	clientbound[V1_8][0x04] = 0x10
	clientbound[V1_8][0x05] = 0x0C
	clientbound[V1_8][0x06] = 0x0B
	clientbound[V1_8][0x07] = 0x37
	clientbound[V1_8][0x08] = 0x25
	clientbound[V1_8][0x09] = 0x35
	clientbound[V1_8][0x0A] = 0x24
	clientbound[V1_8][0x0B] = 0x23
	clientbound[V1_8][0x0C] = -1
	clientbound[V1_8][0x0D] = 0x41
	clientbound[V1_8][0x0E] = 0x3A
	clientbound[V1_8][0x0F] = 0x02
	clientbound[V1_8][0x10] = 0x22
	clientbound[V1_8][0x11] = 0x32
	clientbound[V1_8][0x12] = 0x2E
	clientbound[V1_8][0x13] = 0x2D
	clientbound[V1_8][0x14] = 0x30
	clientbound[V1_8][0x15] = 0x31
	clientbound[V1_8][0x16] = 0x2F
	clientbound[V1_8][0x17] = -1
	clientbound[V1_8][0x18] = 0x3F
	clientbound[V1_8][0x19] = 0x2F
	clientbound[V1_8][0x1A] = 0x40
	clientbound[V1_8][0x1B] = 0x1A
	clientbound[V1_8][0x1C] = 0x27
	clientbound[V1_8][0x1D] = -1
	clientbound[V1_8][0x1E] = 0x2B
	clientbound[V1_8][0x1F] = 0x00
	clientbound[V1_8][0x20] = 0x21
	clientbound[V1_8][0x21] = 0x28
	clientbound[V1_8][0x22] = 0x2A
	clientbound[V1_8][0x23] = 0x01
	clientbound[V1_8][0x24] = 0x34
	clientbound[V1_8][0x25] = 0x15
	clientbound[V1_8][0x26] = 0x17
	clientbound[V1_8][0x27] = 0x16
	clientbound[V1_8][0x28] = 0x14
	clientbound[V1_8][0x29] = -1
	clientbound[V1_8][0x2A] = 0x36
	clientbound[V1_8][0x2B] = 0x39
	clientbound[V1_8][0x2C] = 0x42
	clientbound[V1_8][0x2D] = 0x38
	clientbound[V1_8][0x2E] = 0x08
	clientbound[V1_8][0x2F] = 0x04
	clientbound[V1_8][0x30] = 0x13
	clientbound[V1_8][0x31] = 0x1E
	clientbound[V1_8][0x32] = 0x48
	clientbound[V1_8][0x33] = 0x07
	clientbound[V1_8][0x34] = 0x19
	clientbound[V1_8][0x35] = 0x44
	clientbound[V1_8][0x36] = 0x43
	clientbound[V1_8][0x37] = 0x09
	clientbound[V1_8][0x38] = 0x3D
	clientbound[V1_8][0x39] = 0x1C
	clientbound[V1_8][0x3A] = 0x1B
	clientbound[V1_8][0x3B] = 0x12
	clientbound[V1_8][0x3C] = 0x04
	clientbound[V1_8][0x3D] = 0x1F
	clientbound[V1_8][0x3E] = 0x06
	clientbound[V1_8][0x3F] = 0x3B
	clientbound[V1_8][0x40] = -1
	clientbound[V1_8][0x41] = 0x3E
	clientbound[V1_8][0x42] = 0x3C
	clientbound[V1_8][0x43] = 0x05
	clientbound[V1_8][0x44] = 0x03
	clientbound[V1_8][0x45] = 0x45
	clientbound[V1_8][0x46] = 0x33
	clientbound[V1_8][0x47] = -1
	clientbound[V1_8][0x48] = 0x47
	clientbound[V1_8][0x49] = 0x0D
	clientbound[V1_8][0x4A] = 0x18
	clientbound[V1_8][0x4B] = 0x20
	clientbound[V1_8][0x4C] = 0x1D

	serverbound[V1_8] = make(map[int]int)
	serverbound[V1_8][0x14] = 0x01
	serverbound[V1_8][0x01] = 0x02
	serverbound[V1_8][0x16] = 0x03
	serverbound[V1_8][0x15] = 0x04
	serverbound[V1_8][0x0F] = 0x05
	serverbound[V1_8][0x11] = 0x06
	serverbound[V1_8][0x0E] = 0x07
	serverbound[V1_8][0x0D] = 0x08
	serverbound[V1_8][0x17] = 0x09
	serverbound[V1_8][0x02] = 0x0A
	serverbound[V1_8][0x00] = 0x0B
	serverbound[V1_8][0x04] = 0x0C
	serverbound[V1_8][0x06] = 0x0D
	serverbound[V1_8][0x05] = 0x0E
	serverbound[V1_8][0x03] = 0x0F
	serverbound[V1_8][0x13] = 0x12
	serverbound[V1_8][0x07] = 0x13
	serverbound[V1_8][0x0B] = 0x14
	serverbound[V1_8][0x0C] = 0x15
	serverbound[V1_8][0x19] = 0x16
	serverbound[V1_8][0x09] = 0x17
	serverbound[V1_8][0x10] = 0x18
	serverbound[V1_8][0x12] = 0x19
	serverbound[V1_8][0x0A] = 0x1A
	serverbound[V1_8][0x18] = 0x1B
	serverbound[V1_8][0x08] = 0x1C

	// Hack 1.7.6
	clientbound[V1_7_6] = copyHack(clientbound[V1_8])
	clientbound[V1_7_6][0x0D] = -1

	serverbound[V1_7_6] = serverbound[V1_8]

	// Hack 1.7.2
	clientbound[V1_7_2] = copyHack(clientbound[V1_7_6])

	serverbound[V1_7_2] = serverbound[V1_7_6]

	// Hack 1.12
	clientbound[V1_12] = make(map[int]int)
	clientbound[V1_12][0x28] = 0x25
	for i := 0x25; i <= 0x27; i++ {
		clientbound[V1_12][i] = i + 1
	}
	for i := 0x30; i <= 0x34; i++ {
		clientbound[V1_12][i] = i + 1
	}
	for i := 0x35; i <= 0x49; i++ {
		clientbound[V1_12][i] = i + 2
	}
	for i := 0x4A; i <= 0x4B; i++ {
		clientbound[V1_12][i] = i + 3
	}

	serverbound[V1_12] = make(map[int]int)
	for i := 0x02; i <= 0x16; i++ {
		serverbound[V1_12][i] = i - 1
	}
	serverbound[V1_12][0x18] = 0x16
	for i := 0x1A; i <= 0x20; i++ {
		serverbound[V1_12][i] = i - 3
	}

	// Hack 1.12.1
	clientbound[V1_12_1] = copyHack(clientbound[V1_12])
	for i := 0x2B; i <= 0x4E; i++ {
		clientbound[V1_12_1][lastClientbound(V1_12, i)] = i + 1
	}

	serverbound[V1_12_1] = copyHack(serverbound[V1_12])
	for i := 0x01; i <= 0x11; i++ {
		serverbound[V1_12_1][i] = serverbound[V1_12][i+1]
	}

	// Hack 1.12.2
	clientbound[V1_12_2] = clientbound[V1_12_1]
	serverbound[V1_12_2] = serverbound[V1_12_1]

	// Hack 1.13
	clientbound[V1_13] = copyHack(clientbound[V1_12_2])
	clientbound[V1_13][lastClientbound(V1_12_2, 0x0F)] = 0x0E
	clientbound[V1_13][lastClientbound(V1_12_2, 0x10)] = 0x0F
	clientbound[V1_13][lastClientbound(V1_12_2, 0x0E)] = 0x10
	for i := 0x11; i <= 0x1B; i++ {
		clientbound[V1_13][lastClientbound(V1_12_2, i)] = i + 1
	}
	for i := 0x1C; i <= 0x2E; i++ {
		clientbound[V1_13][lastClientbound(V1_12_2, i)] = i + 2
	}
	for i := 0x2F; i <= 0x48; i++ {
		clientbound[V1_13][lastClientbound(V1_12_2, i)] = i + 3
	}
	for i := 0x49; i <= 0x4F; i++ {
		clientbound[V1_13][lastClientbound(V1_12_2, i)] = i + 4
	}

	serverbound[V1_13] = copyHack(serverbound[V1_12_2])
	serverbound[V1_13][0x05] = serverbound[V1_12_2][0x01]
	for i := 0x06; i <= 0x0A; i++ {
		serverbound[V1_13][i] = serverbound[V1_12_2][i-1]
	}
	for i := 0x0D; i <= 0x14; i++ {
		serverbound[V1_13][i] = serverbound[V1_12_2][i-3]
	}
	for i := 0x16; i <= 0x1B; i++ {
		serverbound[V1_13][i] = serverbound[V1_12_2][i-4]
	}
	for i := 0x1D; i <= 0x1E; i++ {
		serverbound[V1_13][i] = serverbound[V1_12_2][i-5]
	}
	serverbound[V1_13][0x21] = serverbound[V1_12_2][0x1A]
	serverbound[V1_13][0x24] = serverbound[V1_12_2][0x1B]
	for i := 0x26; i <= 0x2A; i++ {
		serverbound[V1_13][i] = serverbound[V1_12_2][i-10]
	}

	// Hack 1.13.1
	clientbound[V1_13_1] = clientbound[V1_13]
	serverbound[V1_13_1] = serverbound[V1_13]

	// Hack 1.13.2
	clientbound[V1_13_2] = clientbound[V1_13]
	serverbound[V1_13_2] = serverbound[V1_13]

	// Hack 1.14
	clientbound[V1_14] = copyHack(clientbound[V1_13_2])
	clientbound[V1_14][lastClientbound(V1_13_2, 0x14)] = 0x2E
	clientbound[V1_14][lastClientbound(V1_13_2, 0x15)] = 0x14
	clientbound[V1_14][lastClientbound(V1_13_2, 0x16)] = 0x15
	clientbound[V1_14][lastClientbound(V1_13_2, 0x17)] = 0x16
	clientbound[V1_14][lastClientbound(V1_13_2, 0x18)] = 0x17
	clientbound[V1_14][lastClientbound(V1_13_2, 0x19)] = 0x18
	clientbound[V1_14][lastClientbound(V1_13_2, 0x1A)] = 0x19
	clientbound[V1_14][lastClientbound(V1_13_2, 0x1B)] = 0x1A
	clientbound[V1_14][lastClientbound(V1_13_2, 0x1C)] = 0x1B
	clientbound[V1_14][lastClientbound(V1_13_2, 0x1D)] = 0x54
	clientbound[V1_14][lastClientbound(V1_13_2, 0x1E)] = 0x1C
	clientbound[V1_14][lastClientbound(V1_13_2, 0x1F)] = 0x1D
	clientbound[V1_14][lastClientbound(V1_13_2, 0x20)] = 0x1E
	clientbound[V1_14][lastClientbound(V1_13_2, 0x21)] = 0x20
	clientbound[V1_14][lastClientbound(V1_13_2, 0x22)] = 0x21
	clientbound[V1_14][lastClientbound(V1_13_2, 0x23)] = 0x22
	clientbound[V1_14][lastClientbound(V1_13_2, 0x24)] = 0x23
	clientbound[V1_14][lastClientbound(V1_13_2, 0x27)] = 0x2B
	clientbound[V1_14][lastClientbound(V1_13_2, 0x2B)] = 0x2C
	clientbound[V1_14][lastClientbound(V1_13_2, 0x2C)] = 0x2D
	clientbound[V1_14][lastClientbound(V1_13_2, 0x2D)] = 0x30
	clientbound[V1_14][lastClientbound(V1_13_2, 0x2E)] = 0x31
	clientbound[V1_14][lastClientbound(V1_13_2, 0x2F)] = 0x32
	clientbound[V1_14][lastClientbound(V1_13_2, 0x30)] = 0x33
	clientbound[V1_14][lastClientbound(V1_13_2, 0x31)] = 0x34
	clientbound[V1_14][lastClientbound(V1_13_2, 0x32)] = 0x35
	clientbound[V1_14][lastClientbound(V1_13_2, 0x34)] = 0x36
	clientbound[V1_14][lastClientbound(V1_13_2, 0x35)] = 0x37
	clientbound[V1_14][lastClientbound(V1_13_2, 0x36)] = 0x38
	clientbound[V1_14][lastClientbound(V1_13_2, 0x37)] = 0x39
	clientbound[V1_14][lastClientbound(V1_13_2, 0x38)] = 0x3A
	clientbound[V1_14][lastClientbound(V1_13_2, 0x39)] = 0x3B
	clientbound[V1_14][lastClientbound(V1_13_2, 0x3A)] = 0x3C
	clientbound[V1_14][lastClientbound(V1_13_2, 0x3B)] = 0x3D
	clientbound[V1_14][lastClientbound(V1_13_2, 0x3C)] = 0x3E
	clientbound[V1_14][lastClientbound(V1_13_2, 0x3D)] = 0x3F
	clientbound[V1_14][lastClientbound(V1_13_2, 0x3E)] = 0x42
	clientbound[V1_14][lastClientbound(V1_13_2, 0x3F)] = 0x43
	clientbound[V1_14][lastClientbound(V1_13_2, 0x40)] = 0x44
	clientbound[V1_14][lastClientbound(V1_13_2, 0x41)] = 0x45
	clientbound[V1_14][lastClientbound(V1_13_2, 0x42)] = 0x46
	clientbound[V1_14][lastClientbound(V1_13_2, 0x43)] = 0x47
	clientbound[V1_14][lastClientbound(V1_13_2, 0x44)] = 0x48
	clientbound[V1_14][lastClientbound(V1_13_2, 0x45)] = 0x49
	clientbound[V1_14][lastClientbound(V1_13_2, 0x46)] = 0x4A
	clientbound[V1_14][lastClientbound(V1_13_2, 0x47)] = 0x4B
	clientbound[V1_14][lastClientbound(V1_13_2, 0x48)] = 0x4C
	clientbound[V1_14][lastClientbound(V1_13_2, 0x49)] = 0x4D
	clientbound[V1_14][lastClientbound(V1_13_2, 0x4A)] = 0x4E
	clientbound[V1_14][lastClientbound(V1_13_2, 0x4B)] = 0x4F
	clientbound[V1_14][lastClientbound(V1_13_2, 0x4C)] = 0x52
	clientbound[V1_14][lastClientbound(V1_13_2, 0x4D)] = 0x51
	clientbound[V1_14][lastClientbound(V1_13_2, 0x4E)] = 0x53
	clientbound[V1_14][lastClientbound(V1_13_2, 0x4F)] = 0x55
	clientbound[V1_14][lastClientbound(V1_13_2, 0x50)] = 0x56
	clientbound[V1_14][lastClientbound(V1_13_2, 0x51)] = 0x57
	clientbound[V1_14][lastClientbound(V1_13_2, 0x52)] = 0x58
	clientbound[V1_14][lastClientbound(V1_13_2, 0x53)] = 0x59
	clientbound[V1_14][lastClientbound(V1_13_2, 0x54)] = 0x5A
	clientbound[V1_14][lastClientbound(V1_13_2, 0x55)] = 0x5B

	serverbound[V1_14] = copyHack(serverbound[V1_13_2])
	serverbound[V1_14][0x03] = serverbound[V1_13_2][0x02]
	serverbound[V1_14][0x04] = serverbound[V1_13_2][0x03]
	serverbound[V1_14][0x05] = serverbound[V1_13_2][0x04]
	serverbound[V1_14][0x06] = serverbound[V1_13_2][0x05]
	serverbound[V1_14][0x07] = serverbound[V1_13_2][0x06]
	serverbound[V1_14][0x08] = serverbound[V1_13_2][0x07]
	serverbound[V1_14][0x09] = serverbound[V1_13_2][0x08]
	serverbound[V1_14][0x0A] = serverbound[V1_13_2][0x09]
	serverbound[V1_14][0x0B] = serverbound[V1_13_2][0x0A]
	serverbound[V1_14][0x0C] = serverbound[V1_13_2][0x0B]
	serverbound[V1_14][0x0D] = serverbound[V1_13_2][0x0C]
	serverbound[V1_14][0x0E] = serverbound[V1_13_2][0x0D]
	serverbound[V1_14][0x0F] = serverbound[V1_13_2][0x0E]
	serverbound[V1_14][0x14] = serverbound[V1_13_2][0x0F]
	serverbound[V1_14][0x11] = serverbound[V1_13_2][0x10]
	serverbound[V1_14][0x12] = serverbound[V1_13_2][0x11]
	serverbound[V1_14][0x13] = serverbound[V1_13_2][0x12]
	serverbound[V1_14][0x15] = serverbound[V1_13_2][0x13]
	serverbound[V1_14][0x16] = serverbound[V1_13_2][0x14]
	serverbound[V1_14][0x17] = serverbound[V1_13_2][0x15]
	serverbound[V1_14][0x18] = serverbound[V1_13_2][0x16]
	serverbound[V1_14][0x19] = serverbound[V1_13_2][0x17]
	serverbound[V1_14][0x1A] = serverbound[V1_13_2][0x18]
	serverbound[V1_14][0x1B] = serverbound[V1_13_2][0x19]
	serverbound[V1_14][0x1C] = serverbound[V1_13_2][0x1A]
	serverbound[V1_14][0x1D] = serverbound[V1_13_2][0x1B]
	serverbound[V1_14][0x1E] = serverbound[V1_13_2][0x1C]
	serverbound[V1_14][0x1F] = serverbound[V1_13_2][0x1D]
	serverbound[V1_14][0x20] = serverbound[V1_13_2][0x1E]
	serverbound[V1_14][0x21] = serverbound[V1_13_2][0x1F]
	serverbound[V1_14][0x22] = serverbound[V1_13_2][0x20]
	serverbound[V1_14][0x23] = serverbound[V1_13_2][0x21]
	serverbound[V1_14][0x24] = serverbound[V1_13_2][0x22]
	serverbound[V1_14][0x25] = serverbound[V1_13_2][0x23]
	serverbound[V1_14][0x26] = serverbound[V1_13_2][0x24]
	serverbound[V1_14][0x28] = serverbound[V1_13_2][0x25]
	serverbound[V1_14][0x29] = serverbound[V1_13_2][0x26]
	serverbound[V1_14][0x2A] = serverbound[V1_13_2][0x27]
	serverbound[V1_14][0x2B] = serverbound[V1_13_2][0x28]
	serverbound[V1_14][0x2C] = serverbound[V1_13_2][0x29]
	serverbound[V1_14][0x2D] = serverbound[V1_13_2][0x2A]

	// Hack 1.14.1
	clientbound[V1_14_1] = clientbound[V1_14]
	serverbound[V1_14_1] = serverbound[V1_14]

	// Hack 1.14.2
	clientbound[V1_14_2] = clientbound[V1_14_1]
	serverbound[V1_14_2] = serverbound[V1_14_1]

	// Hack 1.14.3
	clientbound[V1_14_3] = clientbound[V1_14_2]
	serverbound[V1_14_3] = serverbound[V1_14_2]

	// Hack 1.14.4
	clientbound[V1_14_4] = clientbound[V1_14_3]
	serverbound[V1_14_4] = serverbound[V1_14_3]

	// Hack 1.15
	clientbound[V1_15] = copyHack(clientbound[V1_14_4])
	clientbound[V1_15][lastClientbound(V1_14_4, 0x08)] = 0x09
	clientbound[V1_15][lastClientbound(V1_14_4, 0x09)] = 0x0A
	clientbound[V1_15][lastClientbound(V1_14_4, 0x0A)] = 0x0B
	clientbound[V1_15][lastClientbound(V1_14_4, 0x0B)] = 0x0C
	clientbound[V1_15][lastClientbound(V1_14_4, 0x0C)] = 0x0D
	clientbound[V1_15][lastClientbound(V1_14_4, 0x0D)] = 0x0E
	clientbound[V1_15][lastClientbound(V1_14_4, 0x0E)] = 0x0F
	clientbound[V1_15][lastClientbound(V1_14_4, 0x0F)] = 0x10
	clientbound[V1_15][lastClientbound(V1_14_4, 0x10)] = 0x11
	clientbound[V1_15][lastClientbound(V1_14_4, 0x11)] = 0x12
	clientbound[V1_15][lastClientbound(V1_14_4, 0x12)] = 0x13
	clientbound[V1_15][lastClientbound(V1_14_4, 0x13)] = 0x14
	clientbound[V1_15][lastClientbound(V1_14_4, 0x14)] = 0x15
	clientbound[V1_15][lastClientbound(V1_14_4, 0x15)] = 0x16
	clientbound[V1_15][lastClientbound(V1_14_4, 0x16)] = 0x17
	clientbound[V1_15][lastClientbound(V1_14_4, 0x17)] = 0x18
	clientbound[V1_15][lastClientbound(V1_14_4, 0x18)] = 0x19
	clientbound[V1_15][lastClientbound(V1_14_4, 0x19)] = 0x1A
	clientbound[V1_15][lastClientbound(V1_14_4, 0x1A)] = 0x1B
	clientbound[V1_15][lastClientbound(V1_14_4, 0x1B)] = 0x1C
	clientbound[V1_15][lastClientbound(V1_14_4, 0x1C)] = 0x1D
	clientbound[V1_15][lastClientbound(V1_14_4, 0x1D)] = 0x1E
	clientbound[V1_15][lastClientbound(V1_14_4, 0x1E)] = 0x1F
	clientbound[V1_15][lastClientbound(V1_14_4, 0x1F)] = 0x20
	clientbound[V1_15][lastClientbound(V1_14_4, 0x20)] = 0x21
	clientbound[V1_15][lastClientbound(V1_14_4, 0x21)] = 0x22
	clientbound[V1_15][lastClientbound(V1_14_4, 0x22)] = 0x23
	clientbound[V1_15][lastClientbound(V1_14_4, 0x23)] = 0x24
	clientbound[V1_15][lastClientbound(V1_14_4, 0x24)] = 0x25
	clientbound[V1_15][lastClientbound(V1_14_4, 0x25)] = 0x26
	clientbound[V1_15][lastClientbound(V1_14_4, 0x26)] = 0x27
	clientbound[V1_15][lastClientbound(V1_14_4, 0x27)] = 0x28
	clientbound[V1_15][lastClientbound(V1_14_4, 0x28)] = 0x29
	clientbound[V1_15][lastClientbound(V1_14_4, 0x29)] = 0x2A
	clientbound[V1_15][lastClientbound(V1_14_4, 0x2A)] = 0x2B
	clientbound[V1_15][lastClientbound(V1_14_4, 0x2B)] = 0x2C
	clientbound[V1_15][lastClientbound(V1_14_4, 0x2C)] = 0x2D
	clientbound[V1_15][lastClientbound(V1_14_4, 0x2D)] = 0x2E
	clientbound[V1_15][lastClientbound(V1_14_4, 0x2E)] = 0x2F
	clientbound[V1_15][lastClientbound(V1_14_4, 0x2F)] = 0x30
	clientbound[V1_15][lastClientbound(V1_14_4, 0x30)] = 0x31
	clientbound[V1_15][lastClientbound(V1_14_4, 0x31)] = 0x32
	clientbound[V1_15][lastClientbound(V1_14_4, 0x32)] = 0x33
	clientbound[V1_15][lastClientbound(V1_14_4, 0x33)] = 0x34
	clientbound[V1_15][lastClientbound(V1_14_4, 0x34)] = 0x35
	clientbound[V1_15][lastClientbound(V1_14_4, 0x35)] = 0x36
	clientbound[V1_15][lastClientbound(V1_14_4, 0x36)] = 0x37
	clientbound[V1_15][lastClientbound(V1_14_4, 0x37)] = 0x38
	clientbound[V1_15][lastClientbound(V1_14_4, 0x38)] = 0x39
	clientbound[V1_15][lastClientbound(V1_14_4, 0x39)] = 0x3A
	clientbound[V1_15][lastClientbound(V1_14_4, 0x3A)] = 0x3B
	clientbound[V1_15][lastClientbound(V1_14_4, 0x3B)] = 0x3C
	clientbound[V1_15][lastClientbound(V1_14_4, 0x3C)] = 0x3D
	clientbound[V1_15][lastClientbound(V1_14_4, 0x3D)] = 0x3E
	clientbound[V1_15][lastClientbound(V1_14_4, 0x3E)] = 0x3F
	clientbound[V1_15][lastClientbound(V1_14_4, 0x3F)] = 0x40
	clientbound[V1_15][lastClientbound(V1_14_4, 0x40)] = 0x41
	clientbound[V1_15][lastClientbound(V1_14_4, 0x41)] = 0x42
	clientbound[V1_15][lastClientbound(V1_14_4, 0x42)] = 0x43
	clientbound[V1_15][lastClientbound(V1_14_4, 0x43)] = 0x44
	clientbound[V1_15][lastClientbound(V1_14_4, 0x44)] = 0x45
	clientbound[V1_15][lastClientbound(V1_14_4, 0x45)] = 0x46
	clientbound[V1_15][lastClientbound(V1_14_4, 0x46)] = 0x47
	clientbound[V1_15][lastClientbound(V1_14_4, 0x47)] = 0x48
	clientbound[V1_15][lastClientbound(V1_14_4, 0x48)] = 0x49
	clientbound[V1_15][lastClientbound(V1_14_4, 0x49)] = 0x4A
	clientbound[V1_15][lastClientbound(V1_14_4, 0x4A)] = 0x4B
	clientbound[V1_15][lastClientbound(V1_14_4, 0x4B)] = 0x4C
	clientbound[V1_15][lastClientbound(V1_14_4, 0x4C)] = 0x4D
	clientbound[V1_15][lastClientbound(V1_14_4, 0x4D)] = 0x4E
	clientbound[V1_15][lastClientbound(V1_14_4, 0x4E)] = 0x4F
	clientbound[V1_15][lastClientbound(V1_14_4, 0x4F)] = 0x50
	clientbound[V1_15][lastClientbound(V1_14_4, 0x50)] = 0x51
	clientbound[V1_15][lastClientbound(V1_14_4, 0x51)] = 0x52
	clientbound[V1_15][lastClientbound(V1_14_4, 0x52)] = 0x53
	clientbound[V1_15][lastClientbound(V1_14_4, 0x53)] = 0x54
	clientbound[V1_15][lastClientbound(V1_14_4, 0x54)] = 0x55
	clientbound[V1_15][lastClientbound(V1_14_4, 0x55)] = 0x56
	clientbound[V1_15][lastClientbound(V1_14_4, 0x56)] = 0x57
	clientbound[V1_15][lastClientbound(V1_14_4, 0x57)] = 0x58
	clientbound[V1_15][lastClientbound(V1_14_4, 0x58)] = 0x59
	clientbound[V1_15][lastClientbound(V1_14_4, 0x59)] = 0x5A
	clientbound[V1_15][lastClientbound(V1_14_4, 0x5A)] = 0x5B
	clientbound[V1_15][lastClientbound(V1_14_4, 0x5B)] = 0x5C
	clientbound[V1_15][lastClientbound(V1_14_4, 0x5C)] = 0x08

	serverbound[V1_15] = serverbound[V1_14_4]

	// Hack 1.15.1
	clientbound[V1_15_1] = clientbound[V1_15]
	serverbound[V1_15_1] = serverbound[V1_15]

	// Hack 1.15.2
	clientbound[V1_15_2] = clientbound[V1_15_1]
	serverbound[V1_15_2] = serverbound[V1_15_1]

	// Hack 1.16
	clientbound[V1_16] = copyHack(clientbound[V1_15_2])
	for i := 0x03; i <= 0x5C; i++ {
		clientbound[V1_16][lastClientbound(V1_15_2, i)] = i - 1
	}

	serverbound[V1_16] = copyHack(serverbound[V1_15_2])
	for i := 0x0F; i <= 0x2D; i++ {
		serverbound[V1_16][i+1] = serverbound[V1_15_2][i]
	}

	// Hack 1.16.1
	clientbound[V1_16_1] = clientbound[V1_16]
	serverbound[V1_16_1] = serverbound[V1_16]

	initHackModules()
}

func convUI(i string, v string) (uir int, uvr int, err error) {
	ui, err := strconv.ParseInt(i, 0, 32)
	if err != nil {
		return 0, 0, err
	}
	uv, err := strconv.ParseInt(v, 0, 32)
	if err != nil {
		return 0, 0, err
	}
	return int(ui), int(uv), nil
}

func loadHackModule(module *HackModule) {
	if IsCompatible(module.Content.Base) {
		if clientbound[module.Content.Base] != nil {
			clientbound[module.Content.Protocol] = copyHack(clientbound[module.Content.Base])
		} else {
			clientbound[module.Content.Protocol] = make(map[int]int)
		}
		if serverbound[module.Content.Base] != nil {
			serverbound[module.Content.Protocol] = copyHack(serverbound[module.Content.Base])
		} else {
			serverbound[module.Content.Protocol] = make(map[int]int)
		}

		if module.Content.Base > V1_10 {
			for i, v := range module.Content.Map.Clientbound {
				ui, uv, err := convUI(i, v)
				if err != nil {
					continue
				}
				clientbound[module.Content.Protocol][lastClientbound(module.Content.Base, ui)] = uv
			}
			for i, v := range module.Content.Map.Serverbound {
				ui, uv, err := convUI(i, v)
				if err != nil {
					continue
				}
				serverbound[module.Content.Protocol][uv] = serverbound[module.Content.Base][ui]
			}
		}

		registerProtocol(module.Content.Protocol)
		log.Println("Added", module.Content.Name, "protocol fast support")
	}
}

func loadHackModuleFile(path string) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Can't read file", path)
		return
	}

	var module HackModule
	err = json.Unmarshal(raw, &module)
	if err != nil {
		return
	}

	if module.Type.Name == "protocol-map" && module.Type.Version == 1 {
		loadHackModule(&module)
	}
}

func initHackModules() {
	err := filepath.Walk("modules", func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && strings.HasSuffix(path, ".json") {
			loadHackModuleFile(path)
		}
		return nil
	})
	if err != nil {
		log.Fatal("Can't find modules folder.")
	}
}

func lastClientbound(proto Protocol, i int) int {
	for key, value := range clientbound[proto] {
		if value == i {
			return key
		}
	}
	return i
}

func copyHack(last map[int]int) map[int]int {
	newMap := make(map[int]int)
	for k, v := range last {
		newMap[k] = v
	}
	return newMap
}

func (player *Player) HackServerbound(id int) int {
	_, ok := serverbound[player.protocol]
	if ok {
		if val, ok := serverbound[player.protocol][id]; ok {
			return val
		} else {
			return id
		}
	}
	return id
}

func (player *Player) HackClientbound(id int, protocol Protocol) int {
	//TODO refactor protocol hack
	if protocol != V1_10 {
		return id
	}
	_, ok := clientbound[player.protocol]
	if ok {
		if val, ok := clientbound[player.protocol][id]; ok {
			return val
		} else {
			return id
		}
	}
	return id
}

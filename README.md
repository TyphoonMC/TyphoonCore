# TyphoonCore
## Lightweight minecraft server engine

[![Build Status](https://travis-ci.org/TyphoonMC/TyphoonLimbo.svg?branch=master)](https://travis-ci.org/TyphoonMC/TyphoonCore)
----
### Minecraft protocol support

| Minecraft Version | Protocol Version | Supported |
|-------------------|------------------|-----------|
| 1.7.2 to 1.7.5    | 4                | true      |
| 1.7.6 to 1.7.10   | 5                | true      |
| 1.8 to 1.8.9      | 47               | true      |
| 1.9               | 107              | true      |
| 1.9.1             | 108              | true      |
| 1.9.2             | 109              | true      |
| 1.9.3 to 1.9.4    | 110              | true      |
| 1.10 to 1.10.2    | 210              | true      |
| 1.11              | 315              | true      |
| 1.11.1 to 1.11.2  | 316              | true      |
| 1.12              | 335              | true      |
| 1.12.1            | 338              | true      |
| 1.12.2            | 340              | true      |
| 1.13              | 393              | true      |

#### Snapshot support
TyphoonCore is able to load [TyphoonDep protocol-map modules](https://github.com/TyphoonMC/TyphoonDep/tree/master/protocol-map) to add a partial snapshots support.

All json files are loaded from the "modules" folder in the same directory as the TyphoonLimbo binary.

You can generate a protocol-map module with the [fetcher](https://github.com/TyphoonMC/TyphoonDep/tree/master/protocol-map/fetcher). Just pick the "page" link of your wanted version on [wiki.vg](http://wiki.vg/Protocol_version_numbers) as argument of the fecther and the magic will happen.

### How to use
```shell
go get github.com/TyphoonMC/go.uuid
```

### Example
```go
package main

import (
	"fmt"
	"github.com/TyphoonMC/TyphoonCore"
)

func main() {
	core := typhoon.Init()
	core.SetBrand("exampleserver")

	core.On(func(e *typhoon.PlayerJoinEvent) {
		e.Player.WritePacket(&typhoon.PacketPlayMessage{
			Component: fmt.Sprintf(`{"text":"Welcome %s !"}`, e.Player.GetName()),
			Position:  typhoon.CHAT_BOX,
		})
	})

	core.Start()
}
```

package typhoon

import (
	"bytes"
	"errors"
	"os"

	blocks2 "github.com/TyphoonMC/TyphoonCore/blocks"
	"github.com/seebs/nbt"
)

func LoadSchematic(file string) (*Map, error) {
	dat, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	br := bytes.NewReader(dat)
	tag, _, err := nbt.LoadCompressed(br)
	if err != nil {
		return nil, err
	}

	if tag.Type() != nbt.TypeCompound {
		return nil, errors.New("Unknown nbt type")
	}

	c := tag.(nbt.Compound)

	materials := c["Materials"].(nbt.String)
	if materials.String() != "Alpha" {
		return nil, errors.New("Incompatible schematic version")
	}

	width := int(c["Width"].(nbt.Short))
	height := int(c["Height"].(nbt.Short))
	length := int(c["Length"].(nbt.Short))

	m := &Map{
		Location{X: 0, Y: 0, Z: 0},
		END,
		make([]*Chunk, 0),
	}

	blocks := []int8(c["Blocks"].(nbt.ByteArray))
	data := []int8(c["Data"].(nbt.ByteArray))
	var index int
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			for z := 0; z < length; z++ {
				index = y*width*length + z*width + x
				bigId := blocks2.GetLegacyBlockState(int(uint8(blocks[index])), int(uint8(data[index])))
				name := blocks2.GetNameFromNewBlocks(uint16(4), bigId)
				m.SetBlock(x, y, z, name)
			}
		}
	}
	return m, nil
}

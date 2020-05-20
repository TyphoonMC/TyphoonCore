package typhoon

import (
	"bytes"
	"errors"
	blocks2 "github.com/TyphoonMC/TyphoonCore/blocks"
	"github.com/seebs/nbt"
	"io/ioutil"
)

func LoadSchematic(file string) (*Map, error) {
	dat, err := ioutil.ReadFile(file)
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
		Location{0, 0, 0},
		END,
		make([]*Chunk, 0),
	}

	blocks := []int8(c["Blocks"].(nbt.ByteArray))
	data := []int8(c["Data"].(nbt.ByteArray))
	var index int
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			for z := 0; z < length; z++ {
				index = y * width * length + z * width + x
				bigId := blocks2.GetLegacyFromTypeData(int(uint8(blocks[index])), int(uint8(data[index])))
				m.SetBlock(x, y, z, blocks2.GetNameFromLegacy(int(bigId)))
			}
		}
	}
	return m, nil
}

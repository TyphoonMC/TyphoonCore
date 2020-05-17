package typhoon

import (
	"bytes"
	"errors"
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

	nameMap := make(map[int16]string)
	if c["BlockIds"] != nil {
		//TODO handle this type
		//idList := c["BlockIds"].(nbt.Compound)
	} else if c["SchematicaMapping"] != nil {
		idList := c["SchematicaMapping"].(nbt.Compound)
		for key, value := range idList {
			nameMap[int16(value.(nbt.Short))] = string(key)
		}
	} else {
		return nil, errors.New("Schematic does not contains mapping data")
	}

	mapping := make(map[int16]uint16)
	for key, value := range nameMap {
		mapping[key] = BLOCK_REGISTRY.GetGuid(value)
	}

	m := &Map{
		END,
		make([]*Chunk, 0),
	}

	blocks := []int8(c["Blocks"].(nbt.ByteArray))
	var index int
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			for z := 0; z < length; z++ {
				index = y * width * length + z * width + x
				m.SetBlock(x, y, z, BLOCK_REGISTRY.GetName(mapping[int16(blocks[index])]))
			}
		}
	}
	return m, nil
}

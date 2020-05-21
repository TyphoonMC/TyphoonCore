package blocks

type LegacyBlock struct {
	Id int
	Data int
	Name string
	Protocol uint16
	Fallback *string
}

var legacy = []LegacyBlock{
	{
		0, 0, "minecraft:air",
		0, nil,
	},
	{
		1, 0, "minecraft:stone",
		0, nil,
	},
	{
		1, 1, "minecraft:granite",
		0, nil,
	},
	{
		1, 2, "minecraft:polished_granite",
		0, nil,
	},
	{
		1, 3, "minecraft:diorite",
		0, nil,
	},
	{
		1, 4, "minecraft:polished_diorite",
		0, nil,
	},
	{
		1, 5, "minecraft:andesite",
		0, nil,
	},
	{
		1, 6, "minecraft:polished_andesite",
		0, nil,
	},
	{
		2, 0, "minecraft:grass_block[snowy=false]",
		0, nil,
	},
	{
		3, 0, "minecraft:dirt",
		0, nil,
	},
	{
		3, 1, "minecraft:coarse_dirt",
		0, nil,
	},
	{
		3, 2, "minecraft:podzol[snowy=false]",
		0, nil,
	},
	{
		4, 0, "minecraft:cobblestone",
		0, nil,
	},
	{
		5, 0, "minecraft:oak_planks",
		0, nil,
	},
	{
		5, 1, "minecraft:spruce_planks",
		0, nil,
	},
	{
		5, 2, "minecraft:birch_planks",
		0, nil,
	},
	{
		5, 3, "minecraft:jungle_planks",
		0, nil,
	},
	{
		5, 4, "minecraft:acacia_planks",
		0, nil,
	},
	{
		5, 5, "minecraft:dark_oak_planks",
		0, nil,
	},
	{
		6, 0, "minecraft:oak_sapling[stage=0]",
		0, nil,
	},
	{
		6, 1, "minecraft:spruce_sapling[stage=0]",
		0, nil,
	},
	{
		6, 2, "minecraft:birch_sapling[stage=0]",
		0, nil,
	},
	{
		6, 3, "minecraft:jungle_sapling[stage=0]",
		0, nil,
	},
	{
		6, 4, "minecraft:acacia_sapling[stage=0]",
		0, nil,
	},
	{
		6, 5, "minecraft:dark_oak_sapling[stage=0]",
		0, nil,
	},
	{
		6, 8, "minecraft:oak_sapling[stage=1]",
		0, nil,
	},
	{
		6, 9, "minecraft:spruce_sapling[stage=1]",
		0, nil,
	},
	{
		6, 10, "minecraft:birch_sapling[stage=1]",
		0, nil,
	},
	{
		6, 11, "minecraft:jungle_sapling[stage=1]",
		0, nil,
	},
	{
		6, 12, "minecraft:acacia_sapling[stage=1]",
		0, nil,
	},
	{
		6, 13, "minecraft:dark_oak_sapling[stage=1]",
		0, nil,
	},
	{
		7, 0, "minecraft:bedrock",
		0, nil,
	},
	{
		8, 0, "minecraft:water[level=0]",
		0, nil,
	},
	{
		8, 1, "minecraft:water[level=1]",
		0, nil,
	},
	{
		8, 2, "minecraft:water[level=2]",
		0, nil,
	},
	{
		8, 3, "minecraft:water[level=3]",
		0, nil,
	},
	{
		8, 4, "minecraft:water[level=4]",
		0, nil,
	},
	{
		8, 5, "minecraft:water[level=5]",
		0, nil,
	},
	{
		8, 6, "minecraft:water[level=6]",
		0, nil,
	},
	{
		8, 7, "minecraft:water[level=7]",
		0, nil,
	},
	{
		8, 8, "minecraft:water[level=8]",
		0, nil,
	},
	{
		8, 9, "minecraft:water[level=9]",
		0, nil,
	},
	{
		8, 10, "minecraft:water[level=10]",
		0, nil,
	},
	{
		8, 11, "minecraft:water[level=11]",
		0, nil,
	},
	{
		8, 12, "minecraft:water[level=12]",
		0, nil,
	},
	{
		8, 13, "minecraft:water[level=13]",
		0, nil,
	},
	{
		8, 14, "minecraft:water[level=14]",
		0, nil,
	},
	{
		8, 15, "minecraft:water[level=15]",
		0, nil,
	},
	{
		9, 0, "minecraft:water[level=0]",
		0, nil,
	},
	{
		9, 1, "minecraft:water[level=1]",
		0, nil,
	},
	{
		9, 2, "minecraft:water[level=2]",
		0, nil,
	},
	{
		9, 3, "minecraft:water[level=3]",
		0, nil,
	},
	{
		9, 4, "minecraft:water[level=4]",
		0, nil,
	},
	{
		9, 5, "minecraft:water[level=5]",
		0, nil,
	},
	{
		9, 6, "minecraft:water[level=6]",
		0, nil,
	},
	{
		9, 7, "minecraft:water[level=7]",
		0, nil,
	},
	{
		9, 8, "minecraft:water[level=8]",
		0, nil,
	},
	{
		9, 9, "minecraft:water[level=9]",
		0, nil,
	},
	{
		9, 10, "minecraft:water[level=10]",
		0, nil,
	},
	{
		9, 11, "minecraft:water[level=11]",
		0, nil,
	},
	{
		9, 12, "minecraft:water[level=12]",
		0, nil,
	},
	{
		9, 13, "minecraft:water[level=13]",
		0, nil,
	},
	{
		9, 14, "minecraft:water[level=14]",
		0, nil,
	},
	{
		9, 15, "minecraft:water[level=15]",
		0, nil,
	},
	{
		10, 0, "minecraft:lava[level=0]",
		0, nil,
	},
	{
		10, 1, "minecraft:lava[level=1]",
		0, nil,
	},
	{
		10, 2, "minecraft:lava[level=2]",
		0, nil,
	},
	{
		10, 3, "minecraft:lava[level=3]",
		0, nil,
	},
	{
		10, 4, "minecraft:lava[level=4]",
		0, nil,
	},
	{
		10, 5, "minecraft:lava[level=5]",
		0, nil,
	},
	{
		10, 6, "minecraft:lava[level=6]",
		0, nil,
	},
	{
		10, 7, "minecraft:lava[level=7]",
		0, nil,
	},
	{
		10, 8, "minecraft:lava[level=8]",
		0, nil,
	},
	{
		10, 9, "minecraft:lava[level=9]",
		0, nil,
	},
	{
		10, 10, "minecraft:lava[level=10]",
		0, nil,
	},
	{
		10, 11, "minecraft:lava[level=11]",
		0, nil,
	},
	{
		10, 12, "minecraft:lava[level=12]",
		0, nil,
	},
	{
		10, 13, "minecraft:lava[level=13]",
		0, nil,
	},
	{
		10, 14, "minecraft:lava[level=14]",
		0, nil,
	},
	{
		10, 15, "minecraft:lava[level=15]",
		0, nil,
	},
	{
		11, 0, "minecraft:lava[level=0]",
		0, nil,
	},
	{
		11, 1, "minecraft:lava[level=1]",
		0, nil,
	},
	{
		11, 2, "minecraft:lava[level=2]",
		0, nil,
	},
	{
		11, 3, "minecraft:lava[level=3]",
		0, nil,
	},
	{
		11, 4, "minecraft:lava[level=4]",
		0, nil,
	},
	{
		11, 5, "minecraft:lava[level=5]",
		0, nil,
	},
	{
		11, 6, "minecraft:lava[level=6]",
		0, nil,
	},
	{
		11, 7, "minecraft:lava[level=7]",
		0, nil,
	},
	{
		11, 8, "minecraft:lava[level=8]",
		0, nil,
	},
	{
		11, 9, "minecraft:lava[level=9]",
		0, nil,
	},
	{
		11, 10, "minecraft:lava[level=10]",
		0, nil,
	},
	{
		11, 11, "minecraft:lava[level=11]",
		0, nil,
	},
	{
		11, 12, "minecraft:lava[level=12]",
		0, nil,
	},
	{
		11, 13, "minecraft:lava[level=13]",
		0, nil,
	},
	{
		11, 14, "minecraft:lava[level=14]",
		0, nil,
	},
	{
		11, 15, "minecraft:lava[level=15]",
		0, nil,
	},
	{
		12, 0, "minecraft:sand",
		0, nil,
	},
	{
		12, 1, "minecraft:red_sand",
		0, nil,
	},
	{
		13, 0, "minecraft:gravel",
		0, nil,
	},
	{
		14, 0, "minecraft:gold_ore",
		0, nil,
	},
	{
		15, 0, "minecraft:iron_ore",
		0, nil,
	},
	{
		16, 0, "minecraft:coal_ore",
		0, nil,
	},
	{
		17, 0, "minecraft:oak_log[axis=y]",
		0, nil,
	},
	{
		17, 1, "minecraft:spruce_log[axis=y]",
		0, nil,
	},
	{
		17, 2, "minecraft:birch_log[axis=y]",
		0, nil,
	},
	{
		17, 3, "minecraft:jungle_log[axis=y]",
		0, nil,
	},
	{
		17, 4, "minecraft:oak_log[axis=x]",
		0, nil,
	},
	{
		17, 5, "minecraft:spruce_log[axis=x]",
		0, nil,
	},
	{
		17, 6, "minecraft:birch_log[axis=x]",
		0, nil,
	},
	{
		17, 7, "minecraft:jungle_log[axis=x]",
		0, nil,
	},
	{
		17, 8, "minecraft:oak_log[axis=z]",
		0, nil,
	},
	{
		17, 9, "minecraft:spruce_log[axis=z]",
		0, nil,
	},
	{
		17, 10, "minecraft:birch_log[axis=z]",
		0, nil,
	},
	{
		17, 11, "minecraft:jungle_log[axis=z]",
		0, nil,
	},
	{
		17, 12, "minecraft:oak_wood",
		0, nil,
	},
	{
		17, 13, "minecraft:spruce_wood",
		0, nil,
	},
	{
		17, 14, "minecraft:birch_wood",
		0, nil,
	},
	{
		17, 15, "minecraft:jungle_wood",
		0, nil,
	},
	{
		18, 0, "minecraft:oak_leaves[persistent=false,distance=1]",
		0, nil,
	},
	{
		18, 1, "minecraft:spruce_leaves[persistent=false,distance=1]",
		0, nil,
	},
	{
		18, 2, "minecraft:birch_leaves[persistent=false,distance=1]",
		0, nil,
	},
	{
		18, 3, "minecraft:jungle_leaves[persistent=false,distance=1]",
		0, nil,
	},
	{
		18, 4, "minecraft:oak_leaves[persistent=true,distance=1]",
		0, nil,
	},
	{
		18, 5, "minecraft:spruce_leaves[persistent=true,distance=1]",
		0, nil,
	},
	{
		18, 6, "minecraft:birch_leaves[persistent=true,distance=1]",
		0, nil,
	},
	{
		18, 7, "minecraft:jungle_leaves[persistent=true,distance=1]",
		0, nil,
	},
	{
		18, 8, "minecraft:oak_leaves[persistent=false,distance=1]",
		0, nil,
	},
	{
		18, 9, "minecraft:spruce_leaves[persistent=false,distance=1]",
		0, nil,
	},
	{
		18, 10, "minecraft:birch_leaves[persistent=false,distance=1]",
		0, nil,
	},
	{
		18, 11, "minecraft:jungle_leaves[persistent=false,distance=1]",
		0, nil,
	},
	{
		18, 12, "minecraft:oak_leaves[persistent=true,distance=1]",
		0, nil,
	},
	{
		18, 13, "minecraft:spruce_leaves[persistent=true,distance=1]",
		0, nil,
	},
	{
		18, 14, "minecraft:birch_leaves[persistent=true,distance=1]",
		0, nil,
	},
	{
		18, 15, "minecraft:jungle_leaves[persistent=true,distance=1]",
		0, nil,
	},
	{
		19, 0, "minecraft:sponge",
		0, nil,
	},
	{
		19, 1, "minecraft:wet_sponge",
		0, nil,
	},
	{
		20, 0, "minecraft:glass",
		0, nil,
	},
	{
		21, 0, "minecraft:lapis_ore",
		0, nil,
	},
	{
		22, 0, "minecraft:lapis_block",
		0, nil,
	},
	{
		23, 0, "minecraft:dispenser[triggered=false,facing=down]",
		0, nil,
	},
	{
		23, 1, "minecraft:dispenser[triggered=false,facing=up]",
		0, nil,
	},
	{
		23, 2, "minecraft:dispenser[triggered=false,facing=north]",
		0, nil,
	},
	{
		23, 3, "minecraft:dispenser[triggered=false,facing=south]",
		0, nil,
	},
	{
		23, 4, "minecraft:dispenser[triggered=false,facing=west]",
		0, nil,
	},
	{
		23, 5, "minecraft:dispenser[triggered=false,facing=east]",
		0, nil,
	},
	{
		23, 8, "minecraft:dispenser[triggered=true,facing=down]",
		0, nil,
	},
	{
		23, 9, "minecraft:dispenser[triggered=true,facing=up]",
		0, nil,
	},
	{
		23, 10, "minecraft:dispenser[triggered=true,facing=north]",
		0, nil,
	},
	{
		23, 11, "minecraft:dispenser[triggered=true,facing=south]",
		0, nil,
	},
	{
		23, 12, "minecraft:dispenser[triggered=true,facing=west]",
		0, nil,
	},
	{
		23, 13, "minecraft:dispenser[triggered=true,facing=east]",
		0, nil,
	},
	{
		24, 0, "minecraft:sandstone",
		0, nil,
	},
	{
		24, 1, "minecraft:chiseled_sandstone",
		0, nil,
	},
	{
		24, 2, "minecraft:cut_sandstone",
		0, nil,
	},
	{
		25, 0, "minecraft:note_block",
		0, nil,
	},
	{
		26, 0, "minecraft:red_bed[part=foot,facing=south,occupied=false]",
		0, nil,
	},
	{
		26, 1, "minecraft:red_bed[part=foot,facing=west,occupied=false]",
		0, nil,
	},
	{
		26, 2, "minecraft:red_bed[part=foot,facing=north,occupied=false]",
		0, nil,
	},
	{
		26, 3, "minecraft:red_bed[part=foot,facing=east,occupied=false]",
		0, nil,
	},
	{
		26, 4, "minecraft:red_bed[part=foot,facing=south,occupied=true]",
		0, nil,
	},
	{
		26, 5, "minecraft:red_bed[part=foot,facing=west,occupied=true]",
		0, nil,
	},
	{
		26, 6, "minecraft:red_bed[part=foot,facing=north,occupied=true]",
		0, nil,
	},
	{
		26, 7, "minecraft:red_bed[part=foot,facing=east,occupied=true]",
		0, nil,
	},
	{
		26, 8, "minecraft:red_bed[part=head,facing=south,occupied=false]",
		0, nil,
	},
	{
		26, 9, "minecraft:red_bed[part=head,facing=west,occupied=false]",
		0, nil,
	},
	{
		26, 10, "minecraft:red_bed[part=head,facing=north,occupied=false]",
		0, nil,
	},
	{
		26, 11, "minecraft:red_bed[part=head,facing=east,occupied=false]",
		0, nil,
	},
	{
		26, 12, "minecraft:red_bed[part=head,facing=south,occupied=true]",
		0, nil,
	},
	{
		26, 13, "minecraft:red_bed[part=head,facing=west,occupied=true]",
		0, nil,
	},
	{
		26, 14, "minecraft:red_bed[part=head,facing=north,occupied=true]",
		0, nil,
	},
	{
		26, 15, "minecraft:red_bed[part=head,facing=east,occupied=true]",
		0, nil,
	},
	{
		27, 0, "minecraft:powered_rail[shape=north_south,powered=false]",
		0, nil,
	},
	{
		27, 1, "minecraft:powered_rail[shape=east_west,powered=false]",
		0, nil,
	},
	{
		27, 2, "minecraft:powered_rail[shape=ascending_east,powered=false]",
		0, nil,
	},
	{
		27, 3, "minecraft:powered_rail[shape=ascending_west,powered=false]",
		0, nil,
	},
	{
		27, 4, "minecraft:powered_rail[shape=ascending_north,powered=false]",
		0, nil,
	},
	{
		27, 5, "minecraft:powered_rail[shape=ascending_south,powered=false]",
		0, nil,
	},
	{
		27, 8, "minecraft:powered_rail[shape=north_south,powered=true]",
		0, nil,
	},
	{
		27, 9, "minecraft:powered_rail[shape=east_west,powered=true]",
		0, nil,
	},
	{
		27, 10, "minecraft:powered_rail[shape=ascending_east,powered=true]",
		0, nil,
	},
	{
		27, 11, "minecraft:powered_rail[shape=ascending_west,powered=true]",
		0, nil,
	},
	{
		27, 12, "minecraft:powered_rail[shape=ascending_north,powered=true]",
		0, nil,
	},
	{
		27, 13, "minecraft:powered_rail[shape=ascending_south,powered=true]",
		0, nil,
	},
	{
		28, 0, "minecraft:detector_rail[shape=north_south,powered=false]",
		0, nil,
	},
	{
		28, 1, "minecraft:detector_rail[shape=east_west,powered=false]",
		0, nil,
	},
	{
		28, 2, "minecraft:detector_rail[shape=ascending_east,powered=false]",
		0, nil,
	},
	{
		28, 3, "minecraft:detector_rail[shape=ascending_west,powered=false]",
		0, nil,
	},
	{
		28, 4, "minecraft:detector_rail[shape=ascending_north,powered=false]",
		0, nil,
	},
	{
		28, 5, "minecraft:detector_rail[shape=ascending_south,powered=false]",
		0, nil,
	},
	{
		28, 8, "minecraft:detector_rail[shape=north_south,powered=true]",
		0, nil,
	},
	{
		28, 9, "minecraft:detector_rail[shape=east_west,powered=true]",
		0, nil,
	},
	{
		28, 10, "minecraft:detector_rail[shape=ascending_east,powered=true]",
		0, nil,
	},
	{
		28, 11, "minecraft:detector_rail[shape=ascending_west,powered=true]",
		0, nil,
	},
	{
		28, 12, "minecraft:detector_rail[shape=ascending_north,powered=true]",
		0, nil,
	},
	{
		28, 13, "minecraft:detector_rail[shape=ascending_south,powered=true]",
		0, nil,
	},
	{
		29, 0, "minecraft:sticky_piston[facing=down,extended=false]",
		0, nil,
	},
	{
		29, 1, "minecraft:sticky_piston[facing=up,extended=false]",
		0, nil,
	},
	{
		29, 2, "minecraft:sticky_piston[facing=north,extended=false]",
		0, nil,
	},
	{
		29, 3, "minecraft:sticky_piston[facing=south,extended=false]",
		0, nil,
	},
	{
		29, 4, "minecraft:sticky_piston[facing=west,extended=false]",
		0, nil,
	},
	{
		29, 5, "minecraft:sticky_piston[facing=east,extended=false]",
		0, nil,
	},
	{
		29, 8, "minecraft:sticky_piston[facing=down,extended=true]",
		0, nil,
	},
	{
		29, 9, "minecraft:sticky_piston[facing=up,extended=true]",
		0, nil,
	},
	{
		29, 10, "minecraft:sticky_piston[facing=north,extended=true]",
		0, nil,
	},
	{
		29, 11, "minecraft:sticky_piston[facing=south,extended=true]",
		0, nil,
	},
	{
		29, 12, "minecraft:sticky_piston[facing=west,extended=true]",
		0, nil,
	},
	{
		29, 13, "minecraft:sticky_piston[facing=east,extended=true]",
		0, nil,
	},
	{
		30, 0, "minecraft:cobweb",
		0, nil,
	},
	{
		31, 0, "minecraft:dead_bush",
		0, nil,
	},
	{
		31, 1, "minecraft:grass",
		0, nil,
	},
	{
		31, 2, "minecraft:fern",
		0, nil,
	},
	{
		32, 0, "minecraft:dead_bush",
		0, nil,
	},
	{
		33, 0, "minecraft:piston[facing=down,extended=false]",
		0, nil,
	},
	{
		33, 1, "minecraft:piston[facing=up,extended=false]",
		0, nil,
	},
	{
		33, 2, "minecraft:piston[facing=north,extended=false]",
		0, nil,
	},
	{
		33, 3, "minecraft:piston[facing=south,extended=false]",
		0, nil,
	},
	{
		33, 4, "minecraft:piston[facing=west,extended=false]",
		0, nil,
	},
	{
		33, 5, "minecraft:piston[facing=east,extended=false]",
		0, nil,
	},
	{
		33, 8, "minecraft:piston[facing=down,extended=true]",
		0, nil,
	},
	{
		33, 9, "minecraft:piston[facing=up,extended=true]",
		0, nil,
	},
	{
		33, 10, "minecraft:piston[facing=north,extended=true]",
		0, nil,
	},
	{
		33, 11, "minecraft:piston[facing=south,extended=true]",
		0, nil,
	},
	{
		33, 12, "minecraft:piston[facing=west,extended=true]",
		0, nil,
	},
	{
		33, 13, "minecraft:piston[facing=east,extended=true]",
		0, nil,
	},
	{
		34, 0, "minecraft:piston_head[short=false,facing=down,type=normal]",
		0, nil,
	},
	{
		34, 1, "minecraft:piston_head[short=false,facing=up,type=normal]",
		0, nil,
	},
	{
		34, 2, "minecraft:piston_head[short=false,facing=north,type=normal]",
		0, nil,
	},
	{
		34, 3, "minecraft:piston_head[short=false,facing=south,type=normal]",
		0, nil,
	},
	{
		34, 4, "minecraft:piston_head[short=false,facing=west,type=normal]",
		0, nil,
	},
	{
		34, 5, "minecraft:piston_head[short=false,facing=east,type=normal]",
		0, nil,
	},
	{
		34, 8, "minecraft:piston_head[short=false,facing=down,type=sticky]",
		0, nil,
	},
	{
		34, 9, "minecraft:piston_head[short=false,facing=up,type=sticky]",
		0, nil,
	},
	{
		34, 10, "minecraft:piston_head[short=false,facing=north,type=sticky]",
		0, nil,
	},
	{
		34, 11, "minecraft:piston_head[short=false,facing=south,type=sticky]",
		0, nil,
	},
	{
		34, 12, "minecraft:piston_head[short=false,facing=west,type=sticky]",
		0, nil,
	},
	{
		34, 13, "minecraft:piston_head[short=false,facing=east,type=sticky]",
		0, nil,
	},
	{
		35, 0, "minecraft:white_wool",
		0, nil,
	},
	{
		35, 1, "minecraft:orange_wool",
		0, nil,
	},
	{
		35, 2, "minecraft:magenta_wool",
		0, nil,
	},
	{
		35, 3, "minecraft:light_blue_wool",
		0, nil,
	},
	{
		35, 4, "minecraft:yellow_wool",
		0, nil,
	},
	{
		35, 5, "minecraft:lime_wool",
		0, nil,
	},
	{
		35, 6, "minecraft:pink_wool",
		0, nil,
	},
	{
		35, 7, "minecraft:gray_wool",
		0, nil,
	},
	{
		35, 8, "minecraft:light_gray_wool",
		0, nil,
	},
	{
		35, 9, "minecraft:cyan_wool",
		0, nil,
	},
	{
		35, 10, "minecraft:purple_wool",
		0, nil,
	},
	{
		35, 11, "minecraft:blue_wool",
		0, nil,
	},
	{
		35, 12, "minecraft:brown_wool",
		0, nil,
	},
	{
		35, 13, "minecraft:green_wool",
		0, nil,
	},
	{
		35, 14, "minecraft:red_wool",
		0, nil,
	},
	{
		35, 15, "minecraft:black_wool",
		0, nil,
	},
	{
		36, 0, "minecraft:moving_piston[facing=down,type=normal]",
		0, nil,
	},
	{
		36, 1, "minecraft:moving_piston[facing=up,type=normal]",
		0, nil,
	},
	{
		36, 2, "minecraft:moving_piston[facing=north,type=normal]",
		0, nil,
	},
	{
		36, 3, "minecraft:moving_piston[facing=south,type=normal]",
		0, nil,
	},
	{
		36, 4, "minecraft:moving_piston[facing=west,type=normal]",
		0, nil,
	},
	{
		36, 5, "minecraft:moving_piston[facing=east,type=normal]",
		0, nil,
	},
	{
		36, 8, "minecraft:moving_piston[facing=down,type=sticky]",
		0, nil,
	},
	{
		36, 9, "minecraft:moving_piston[facing=up,type=sticky]",
		0, nil,
	},
	{
		36, 10, "minecraft:moving_piston[facing=north,type=sticky]",
		0, nil,
	},
	{
		36, 11, "minecraft:moving_piston[facing=south,type=sticky]",
		0, nil,
	},
	{
		36, 12, "minecraft:moving_piston[facing=west,type=sticky]",
		0, nil,
	},
	{
		36, 13, "minecraft:moving_piston[facing=east,type=sticky]",
		0, nil,
	},
	{
		37, 0, "minecraft:dandelion",
		0, nil,
	},
	{
		38, 0, "minecraft:poppy",
		0, nil,
	},
	{
		38, 1, "minecraft:blue_orchid",
		0, nil,
	},
	{
		38, 2, "minecraft:allium",
		0, nil,
	},
	{
		38, 3, "minecraft:azure_bluet",
		0, nil,
	},
	{
		38, 4, "minecraft:red_tulip",
		0, nil,
	},
	{
		38, 5, "minecraft:orange_tulip",
		0, nil,
	},
	{
		38, 6, "minecraft:white_tulip",
		0, nil,
	},
	{
		38, 7, "minecraft:pink_tulip",
		0, nil,
	},
	{
		38, 8, "minecraft:oxeye_daisy",
		0, nil,
	},
	{
		39, 0, "minecraft:brown_mushroom",
		0, nil,
	},
	{
		40, 0, "minecraft:red_mushroom",
		0, nil,
	},
	{
		41, 0, "minecraft:gold_block",
		0, nil,
	},
	{
		42, 0, "minecraft:iron_block",
		0, nil,
	},
	{
		43, 0, "minecraft:stone_slab[type=double]",
		0, nil,
	},
	{
		43, 1, "minecraft:sandstone_slab[type=double]",
		0, nil,
	},
	{
		43, 2, "minecraft:petrified_oak_slab[type=double]",
		0, nil,
	},
	{
		43, 3, "minecraft:cobblestone_slab[type=double]",
		0, nil,
	},
	{
		43, 4, "minecraft:brick_slab[type=double]",
		0, nil,
	},
	{
		43, 5, "minecraft:stone_brick_slab[type=double]",
		0, nil,
	},
	{
		43, 6, "minecraft:nether_brick_slab[type=double]",
		0, nil,
	},
	{
		43, 7, "minecraft:quartz_slab[type=double]",
		0, nil,
	},
	{
		43, 8, "minecraft:smooth_stone",
		0, nil,
	},
	{
		43, 9, "minecraft:smooth_sandstone",
		0, nil,
	},
	{
		43, 10, "minecraft:petrified_oak_slab[type=double]",
		0, nil,
	},
	{
		43, 11, "minecraft:cobblestone_slab[type=double]",
		0, nil,
	},
	{
		43, 12, "minecraft:brick_slab[type=double]",
		0, nil,
	},
	{
		43, 13, "minecraft:stone_brick_slab[type=double]",
		0, nil,
	},
	{
		43, 14, "minecraft:nether_brick_slab[type=double]",
		0, nil,
	},
	{
		43, 15, "minecraft:smooth_quartz",
		0, nil,
	},
	{
		44, 0, "minecraft:stone_slab[type=bottom]",
		0, nil,
	},
	{
		44, 1, "minecraft:sandstone_slab[type=bottom]",
		0, nil,
	},
	{
		44, 2, "minecraft:petrified_oak_slab[type=bottom]",
		0, nil,
	},
	{
		44, 3, "minecraft:cobblestone_slab[type=bottom]",
		0, nil,
	},
	{
		44, 4, "minecraft:brick_slab[type=bottom]",
		0, nil,
	},
	{
		44, 5, "minecraft:stone_brick_slab[type=bottom]",
		0, nil,
	},
	{
		44, 6, "minecraft:nether_brick_slab[type=bottom]",
		0, nil,
	},
	{
		44, 7, "minecraft:quartz_slab[type=bottom]",
		0, nil,
	},
	{
		44, 8, "minecraft:stone_slab[type=top]",
		0, nil,
	},
	{
		44, 9, "minecraft:sandstone_slab[type=top]",
		0, nil,
	},
	{
		44, 10, "minecraft:petrified_oak_slab[type=top]",
		0, nil,
	},
	{
		44, 11, "minecraft:cobblestone_slab[type=top]",
		0, nil,
	},
	{
		44, 12, "minecraft:brick_slab[type=top]",
		0, nil,
	},
	{
		44, 13, "minecraft:stone_brick_slab[type=top]",
		0, nil,
	},
	{
		44, 14, "minecraft:nether_brick_slab[type=top]",
		0, nil,
	},
	{
		44, 15, "minecraft:quartz_slab[type=top]",
		0, nil,
	},
	{
		45, 0, "minecraft:bricks",
		0, nil,
	},
	{
		46, 0, "minecraft:tnt[unstable=false]",
		0, nil,
	},
	{
		46, 1, "minecraft:tnt[unstable=true]",
		0, nil,
	},
	{
		47, 0, "minecraft:bookshelf",
		0, nil,
	},
	{
		48, 0, "minecraft:mossy_cobblestone",
		0, nil,
	},
	{
		49, 0, "minecraft:obsidian",
		0, nil,
	},
	{
		50, 1, "minecraft:wall_torch[facing=east]",
		0, nil,
	},
	{
		50, 2, "minecraft:wall_torch[facing=west]",
		0, nil,
	},
	{
		50, 3, "minecraft:wall_torch[facing=south]",
		0, nil,
	},
	{
		50, 4, "minecraft:wall_torch[facing=north]",
		0, nil,
	},
	{
		50, 5, "minecraft:torch",
		0, nil,
	},
	{
		51, 0, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=0]",
		0, nil,
	},
	{
		51, 1, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=1]",
		0, nil,
	},
	{
		51, 2, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=2]",
		0, nil,
	},
	{
		51, 3, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=3]",
		0, nil,
	},
	{
		51, 4, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=4]",
		0, nil,
	},
	{
		51, 5, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=5]",
		0, nil,
	},
	{
		51, 6, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=6]",
		0, nil,
	},
	{
		51, 7, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=7]",
		0, nil,
	},
	{
		51, 8, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=8]",
		0, nil,
	},
	{
		51, 9, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=9]",
		0, nil,
	},
	{
		51, 10, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=10]",
		0, nil,
	},
	{
		51, 11, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=11]",
		0, nil,
	},
	{
		51, 12, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=12]",
		0, nil,
	},
	{
		51, 13, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=13]",
		0, nil,
	},
	{
		51, 14, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=14]",
		0, nil,
	},
	{
		51, 15, "minecraft:fire[east=false,south=false,north=false,west=false,up=false,age=15]",
		0, nil,
	},
	{
		52, 0, "minecraft:spawner",
		0, nil,
	},
	{
		53, 0, "minecraft:oak_stairs[half=bottom,shape=outer_right,facing=east]",
		0, nil,
	},
	{
		53, 1, "minecraft:oak_stairs[half=bottom,shape=outer_right,facing=west]",
		0, nil,
	},
	{
		53, 2, "minecraft:oak_stairs[half=bottom,shape=outer_right,facing=south]",
		0, nil,
	},
	{
		53, 3, "minecraft:oak_stairs[half=bottom,shape=outer_right,facing=north]",
		0, nil,
	},
	{
		53, 4, "minecraft:oak_stairs[half=top,shape=outer_right,facing=east]",
		0, nil,
	},
	{
		53, 5, "minecraft:oak_stairs[half=top,shape=outer_right,facing=west]",
		0, nil,
	},
	{
		53, 6, "minecraft:oak_stairs[half=top,shape=outer_right,facing=south]",
		0, nil,
	},
	{
		53, 7, "minecraft:oak_stairs[half=top,shape=outer_right,facing=north]",
		0, nil,
	},
	{
		54, 2, "minecraft:chest[facing=north,type=single]",
		0, nil,
	},
	{
		54, 3, "minecraft:chest[facing=south,type=single]",
		0, nil,
	},
	{
		54, 4, "minecraft:chest[facing=west,type=single]",
		0, nil,
	},
	{
		54, 5, "minecraft:chest[facing=east,type=single]",
		0, nil,
	},
	{
		55, 0, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=0]",
		0, nil,
	},
	{
		55, 1, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=1]",
		0, nil,
	},
	{
		55, 2, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=2]",
		0, nil,
	},
	{
		55, 3, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=3]",
		0, nil,
	},
	{
		55, 4, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=4]",
		0, nil,
	},
	{
		55, 5, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=5]",
		0, nil,
	},
	{
		55, 6, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=6]",
		0, nil,
	},
	{
		55, 7, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=7]",
		0, nil,
	},
	{
		55, 8, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=8]",
		0, nil,
	},
	{
		55, 9, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=9]",
		0, nil,
	},
	{
		55, 10, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=10]",
		0, nil,
	},
	{
		55, 11, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=11]",
		0, nil,
	},
	{
		55, 12, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=12]",
		0, nil,
	},
	{
		55, 13, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=13]",
		0, nil,
	},
	{
		55, 14, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=14]",
		0, nil,
	},
	{
		55, 15, "minecraft:redstone_wire[east=none,south=none,north=none,west=none,power=15]",
		0, nil,
	},
	{
		56, 0, "minecraft:diamond_ore",
		0, nil,
	},
	{
		57, 0, "minecraft:diamond_block",
		0, nil,
	},
	{
		58, 0, "minecraft:crafting_table",
		0, nil,
	},
	{
		59, 0, "minecraft:wheat[age=0]",
		0, nil,
	},
	{
		59, 1, "minecraft:wheat[age=1]",
		0, nil,
	},
	{
		59, 2, "minecraft:wheat[age=2]",
		0, nil,
	},
	{
		59, 3, "minecraft:wheat[age=3]",
		0, nil,
	},
	{
		59, 4, "minecraft:wheat[age=4]",
		0, nil,
	},
	{
		59, 5, "minecraft:wheat[age=5]",
		0, nil,
	},
	{
		59, 6, "minecraft:wheat[age=6]",
		0, nil,
	},
	{
		59, 7, "minecraft:wheat[age=7]",
		0, nil,
	},
	{
		60, 0, "minecraft:farmland[moisture=0]",
		0, nil,
	},
	{
		60, 1, "minecraft:farmland[moisture=1]",
		0, nil,
	},
	{
		60, 2, "minecraft:farmland[moisture=2]",
		0, nil,
	},
	{
		60, 3, "minecraft:farmland[moisture=3]",
		0, nil,
	},
	{
		60, 4, "minecraft:farmland[moisture=4]",
		0, nil,
	},
	{
		60, 5, "minecraft:farmland[moisture=5]",
		0, nil,
	},
	{
		60, 6, "minecraft:farmland[moisture=6]",
		0, nil,
	},
	{
		60, 7, "minecraft:farmland[moisture=7]",
		0, nil,
	},
	{
		61, 2, "minecraft:furnace[facing=north,lit=false]",
		0, nil,
	},
	{
		61, 3, "minecraft:furnace[facing=south,lit=false]",
		0, nil,
	},
	{
		61, 4, "minecraft:furnace[facing=west,lit=false]",
		0, nil,
	},
	{
		61, 5, "minecraft:furnace[facing=east,lit=false]",
		0, nil,
	},
	{
		62, 2, "minecraft:furnace[facing=north,lit=true]",
		0, nil,
	},
	{
		62, 3, "minecraft:furnace[facing=south,lit=true]",
		0, nil,
	},
	{
		62, 4, "minecraft:furnace[facing=west,lit=true]",
		0, nil,
	},
	{
		62, 5, "minecraft:furnace[facing=east,lit=true]",
		0, nil,
	},
	{
		63, 0, "minecraft:sign[rotation=0]",
		0, nil,
	},
	{
		63, 1, "minecraft:sign[rotation=1]",
		0, nil,
	},
	{
		63, 2, "minecraft:sign[rotation=2]",
		0, nil,
	},
	{
		63, 3, "minecraft:sign[rotation=3]",
		0, nil,
	},
	{
		63, 4, "minecraft:sign[rotation=4]",
		0, nil,
	},
	{
		63, 5, "minecraft:sign[rotation=5]",
		0, nil,
	},
	{
		63, 6, "minecraft:sign[rotation=6]",
		0, nil,
	},
	{
		63, 7, "minecraft:sign[rotation=7]",
		0, nil,
	},
	{
		63, 8, "minecraft:sign[rotation=8]",
		0, nil,
	},
	{
		63, 9, "minecraft:sign[rotation=9]",
		0, nil,
	},
	{
		63, 10, "minecraft:sign[rotation=10]",
		0, nil,
	},
	{
		63, 11, "minecraft:sign[rotation=11]",
		0, nil,
	},
	{
		63, 12, "minecraft:sign[rotation=12]",
		0, nil,
	},
	{
		63, 13, "minecraft:sign[rotation=13]",
		0, nil,
	},
	{
		63, 14, "minecraft:sign[rotation=14]",
		0, nil,
	},
	{
		63, 15, "minecraft:sign[rotation=15]",
		0, nil,
	},
	{
		64, 0, "minecraft:oak_door[hinge=right,half=lower,powered=false,facing=east,open=false]",
		0, nil,
	},
	{
		64, 1, "minecraft:oak_door[hinge=right,half=lower,powered=false,facing=south,open=false]",
		0, nil,
	},
	{
		64, 2, "minecraft:oak_door[hinge=right,half=lower,powered=false,facing=west,open=false]",
		0, nil,
	},
	{
		64, 3, "minecraft:oak_door[hinge=right,half=lower,powered=false,facing=north,open=false]",
		0, nil,
	},
	{
		64, 4, "minecraft:oak_door[hinge=right,half=lower,powered=false,facing=east,open=true]",
		0, nil,
	},
	{
		64, 5, "minecraft:oak_door[hinge=right,half=lower,powered=false,facing=south,open=true]",
		0, nil,
	},
	{
		64, 6, "minecraft:oak_door[hinge=right,half=lower,powered=false,facing=west,open=true]",
		0, nil,
	},
	{
		64, 7, "minecraft:oak_door[hinge=right,half=lower,powered=false,facing=north,open=true]",
		0, nil,
	},
	{
		64, 8, "minecraft:oak_door[hinge=left,half=upper,powered=false,facing=east,open=false]",
		0, nil,
	},
	{
		64, 9, "minecraft:oak_door[hinge=right,half=upper,powered=false,facing=east,open=false]",
		0, nil,
	},
	{
		64, 10, "minecraft:oak_door[hinge=left,half=upper,powered=true,facing=east,open=false]",
		0, nil,
	},
	{
		64, 11, "minecraft:oak_door[hinge=right,half=upper,powered=true,facing=east,open=false]",
		0, nil,
	},
	{
		65, 2, "minecraft:ladder[facing=north]",
		0, nil,
	},
	{
		65, 3, "minecraft:ladder[facing=south]",
		0, nil,
	},
	{
		65, 4, "minecraft:ladder[facing=west]",
		0, nil,
	},
	{
		65, 5, "minecraft:ladder[facing=east]",
		0, nil,
	},
	{
		66, 0, "minecraft:rail[shape=north_south]",
		0, nil,
	},
	{
		66, 1, "minecraft:rail[shape=east_west]",
		0, nil,
	},
	{
		66, 2, "minecraft:rail[shape=ascending_east]",
		0, nil,
	},
	{
		66, 3, "minecraft:rail[shape=ascending_west]",
		0, nil,
	},
	{
		66, 4, "minecraft:rail[shape=ascending_north]",
		0, nil,
	},
	{
		66, 5, "minecraft:rail[shape=ascending_south]",
		0, nil,
	},
	{
		66, 6, "minecraft:rail[shape=south_east]",
		0, nil,
	},
	{
		66, 7, "minecraft:rail[shape=south_west]",
		0, nil,
	},
	{
		66, 8, "minecraft:rail[shape=north_west]",
		0, nil,
	},
	{
		66, 9, "minecraft:rail[shape=north_east]",
		0, nil,
	},
	{
		67, 0, "minecraft:cobblestone_stairs[half=bottom,shape=straight,facing=east]",
		0, nil,
	},
	{
		67, 1, "minecraft:cobblestone_stairs[half=bottom,shape=straight,facing=west]",
		0, nil,
	},
	{
		67, 2, "minecraft:cobblestone_stairs[half=bottom,shape=straight,facing=south]",
		0, nil,
	},
	{
		67, 3, "minecraft:cobblestone_stairs[half=bottom,shape=straight,facing=north]",
		0, nil,
	},
	{
		67, 4, "minecraft:cobblestone_stairs[half=top,shape=straight,facing=east]",
		0, nil,
	},
	{
		67, 5, "minecraft:cobblestone_stairs[half=top,shape=straight,facing=west]",
		0, nil,
	},
	{
		67, 6, "minecraft:cobblestone_stairs[half=top,shape=straight,facing=south]",
		0, nil,
	},
	{
		67, 7, "minecraft:cobblestone_stairs[half=top,shape=straight,facing=north]",
		0, nil,
	},
	{
		68, 2, "minecraft:wall_sign[facing=north]",
		0, nil,
	},
	{
		68, 3, "minecraft:wall_sign[facing=south]",
		0, nil,
	},
	{
		68, 4, "minecraft:wall_sign[facing=west]",
		0, nil,
	},
	{
		68, 5, "minecraft:wall_sign[facing=east]",
		0, nil,
	},
	{
		69, 0, "minecraft:lever[powered=false,facing=north,face=ceiling]",
		0, nil,
	},
	{
		69, 1, "minecraft:lever[powered=false,facing=east,face=wall]",
		0, nil,
	},
	{
		69, 2, "minecraft:lever[powered=false,facing=west,face=wall]",
		0, nil,
	},
	{
		69, 3, "minecraft:lever[powered=false,facing=south,face=wall]",
		0, nil,
	},
	{
		69, 4, "minecraft:lever[powered=false,facing=north,face=wall]",
		0, nil,
	},
	{
		69, 5, "minecraft:lever[powered=false,facing=east,face=floor]",
		0, nil,
	},
	{
		69, 6, "minecraft:lever[powered=false,facing=north,face=floor]",
		0, nil,
	},
	{
		69, 7, "minecraft:lever[powered=false,facing=east,face=ceiling]",
		0, nil,
	},
	{
		69, 8, "minecraft:lever[powered=true,facing=north,face=ceiling]",
		0, nil,
	},
	{
		69, 9, "minecraft:lever[powered=true,facing=east,face=wall]",
		0, nil,
	},
	{
		69, 10, "minecraft:lever[powered=true,facing=west,face=wall]",
		0, nil,
	},
	{
		69, 11, "minecraft:lever[powered=true,facing=south,face=wall]",
		0, nil,
	},
	{
		69, 12, "minecraft:lever[powered=true,facing=north,face=wall]",
		0, nil,
	},
	{
		69, 13, "minecraft:lever[powered=true,facing=east,face=floor]",
		0, nil,
	},
	{
		69, 14, "minecraft:lever[powered=true,facing=north,face=floor]",
		0, nil,
	},
	{
		69, 15, "minecraft:lever[powered=true,facing=east,face=ceiling]",
		0, nil,
	},
	{
		70, 0, "minecraft:stone_pressure_plate[powered=false]",
		0, nil,
	},
	{
		70, 1, "minecraft:stone_pressure_plate[powered=true]",
		0, nil,
	},
	{
		71, 0, "minecraft:iron_door[hinge=right,half=lower,powered=false,facing=east,open=false]",
		0, nil,
	},
	{
		71, 1, "minecraft:iron_door[hinge=right,half=lower,powered=false,facing=south,open=false]",
		0, nil,
	},
	{
		71, 2, "minecraft:iron_door[hinge=right,half=lower,powered=false,facing=west,open=false]",
		0, nil,
	},
	{
		71, 3, "minecraft:iron_door[hinge=right,half=lower,powered=false,facing=north,open=false]",
		0, nil,
	},
	{
		71, 4, "minecraft:iron_door[hinge=right,half=lower,powered=false,facing=east,open=true]",
		0, nil,
	},
	{
		71, 5, "minecraft:iron_door[hinge=right,half=lower,powered=false,facing=south,open=true]",
		0, nil,
	},
	{
		71, 6, "minecraft:iron_door[hinge=right,half=lower,powered=false,facing=west,open=true]",
		0, nil,
	},
	{
		71, 7, "minecraft:iron_door[hinge=right,half=lower,powered=false,facing=north,open=true]",
		0, nil,
	},
	{
		71, 8, "minecraft:iron_door[hinge=left,half=upper,powered=false,facing=east,open=false]",
		0, nil,
	},
	{
		71, 9, "minecraft:iron_door[hinge=right,half=upper,powered=false,facing=east,open=false]",
		0, nil,
	},
	{
		71, 10, "minecraft:iron_door[hinge=left,half=upper,powered=true,facing=east,open=false]",
		0, nil,
	},
	{
		71, 11, "minecraft:iron_door[hinge=right,half=upper,powered=true,facing=east,open=false]",
		0, nil,
	},
	{
		72, 0, "minecraft:oak_pressure_plate[powered=false]",
		0, nil,
	},
	{
		72, 1, "minecraft:oak_pressure_plate[powered=true]",
		0, nil,
	},
	{
		73, 0, "minecraft:redstone_ore[lit=false]",
		0, nil,
	},
	{
		74, 0, "minecraft:redstone_ore[lit=true]",
		0, nil,
	},
	{
		75, 1, "minecraft:redstone_wall_torch[facing=east,lit=false]",
		0, nil,
	},
	{
		75, 2, "minecraft:redstone_wall_torch[facing=west,lit=false]",
		0, nil,
	},
	{
		75, 3, "minecraft:redstone_wall_torch[facing=south,lit=false]",
		0, nil,
	},
	{
		75, 4, "minecraft:redstone_wall_torch[facing=north,lit=false]",
		0, nil,
	},
	{
		75, 5, "minecraft:redstone_torch[lit=false]",
		0, nil,
	},
	{
		76, 1, "minecraft:redstone_wall_torch[facing=east,lit=true]",
		0, nil,
	},
	{
		76, 2, "minecraft:redstone_wall_torch[facing=west,lit=true]",
		0, nil,
	},
	{
		76, 3, "minecraft:redstone_wall_torch[facing=south,lit=true]",
		0, nil,
	},
	{
		76, 4, "minecraft:redstone_wall_torch[facing=north,lit=true]",
		0, nil,
	},
	{
		76, 5, "minecraft:redstone_torch[lit=true]",
		0, nil,
	},
	{
		77, 0, "minecraft:stone_button[powered=false,facing=east,face=ceiling]",
		0, nil,
	},
	{
		77, 1, "minecraft:stone_button[powered=false,facing=east,face=wall]",
		0, nil,
	},
	{
		77, 2, "minecraft:stone_button[powered=false,facing=west,face=wall]",
		0, nil,
	},
	{
		77, 3, "minecraft:stone_button[powered=false,facing=south,face=wall]",
		0, nil,
	},
	{
		77, 4, "minecraft:stone_button[powered=false,facing=north,face=wall]",
		0, nil,
	},
	{
		77, 5, "minecraft:stone_button[powered=false,facing=east,face=floor]",
		0, nil,
	},
	{
		77, 8, "minecraft:stone_button[powered=true,facing=south,face=ceiling]",
		0, nil,
	},
	{
		77, 9, "minecraft:stone_button[powered=true,facing=east,face=wall]",
		0, nil,
	},
	{
		77, 10, "minecraft:stone_button[powered=true,facing=west,face=wall]",
		0, nil,
	},
	{
		77, 11, "minecraft:stone_button[powered=true,facing=south,face=wall]",
		0, nil,
	},
	{
		77, 12, "minecraft:stone_button[powered=true,facing=north,face=wall]",
		0, nil,
	},
	{
		77, 13, "minecraft:stone_button[powered=true,facing=south,face=floor]",
		0, nil,
	},
	{
		78, 0, "minecraft:snow[layers=1]",
		0, nil,
	},
	{
		78, 1, "minecraft:snow[layers=2]",
		0, nil,
	},
	{
		78, 2, "minecraft:snow[layers=3]",
		0, nil,
	},
	{
		78, 3, "minecraft:snow[layers=4]",
		0, nil,
	},
	{
		78, 4, "minecraft:snow[layers=5]",
		0, nil,
	},
	{
		78, 5, "minecraft:snow[layers=6]",
		0, nil,
	},
	{
		78, 6, "minecraft:snow[layers=7]",
		0, nil,
	},
	{
		78, 7, "minecraft:snow[layers=8]",
		0, nil,
	},
	{
		79, 0, "minecraft:ice",
		0, nil,
	},
	{
		80, 0, "minecraft:snow_block",
		0, nil,
	},
	{
		81, 0, "minecraft:cactus[age=0]",
		0, nil,
	},
	{
		81, 1, "minecraft:cactus[age=1]",
		0, nil,
	},
	{
		81, 2, "minecraft:cactus[age=2]",
		0, nil,
	},
	{
		81, 3, "minecraft:cactus[age=3]",
		0, nil,
	},
	{
		81, 4, "minecraft:cactus[age=4]",
		0, nil,
	},
	{
		81, 5, "minecraft:cactus[age=5]",
		0, nil,
	},
	{
		81, 6, "minecraft:cactus[age=6]",
		0, nil,
	},
	{
		81, 7, "minecraft:cactus[age=7]",
		0, nil,
	},
	{
		81, 8, "minecraft:cactus[age=8]",
		0, nil,
	},
	{
		81, 9, "minecraft:cactus[age=9]",
		0, nil,
	},
	{
		81, 10, "minecraft:cactus[age=10]",
		0, nil,
	},
	{
		81, 11, "minecraft:cactus[age=11]",
		0, nil,
	},
	{
		81, 12, "minecraft:cactus[age=12]",
		0, nil,
	},
	{
		81, 13, "minecraft:cactus[age=13]",
		0, nil,
	},
	{
		81, 14, "minecraft:cactus[age=14]",
		0, nil,
	},
	{
		81, 15, "minecraft:cactus[age=15]",
		0, nil,
	},
	{
		82, 0, "minecraft:clay",
		0, nil,
	},
	{
		83, 0, "minecraft:sugar_cane[age=0]",
		0, nil,
	},
	{
		83, 1, "minecraft:sugar_cane[age=1]",
		0, nil,
	},
	{
		83, 2, "minecraft:sugar_cane[age=2]",
		0, nil,
	},
	{
		83, 3, "minecraft:sugar_cane[age=3]",
		0, nil,
	},
	{
		83, 4, "minecraft:sugar_cane[age=4]",
		0, nil,
	},
	{
		83, 5, "minecraft:sugar_cane[age=5]",
		0, nil,
	},
	{
		83, 6, "minecraft:sugar_cane[age=6]",
		0, nil,
	},
	{
		83, 7, "minecraft:sugar_cane[age=7]",
		0, nil,
	},
	{
		83, 8, "minecraft:sugar_cane[age=8]",
		0, nil,
	},
	{
		83, 9, "minecraft:sugar_cane[age=9]",
		0, nil,
	},
	{
		83, 10, "minecraft:sugar_cane[age=10]",
		0, nil,
	},
	{
		83, 11, "minecraft:sugar_cane[age=11]",
		0, nil,
	},
	{
		83, 12, "minecraft:sugar_cane[age=12]",
		0, nil,
	},
	{
		83, 13, "minecraft:sugar_cane[age=13]",
		0, nil,
	},
	{
		83, 14, "minecraft:sugar_cane[age=14]",
		0, nil,
	},
	{
		83, 15, "minecraft:sugar_cane[age=15]",
		0, nil,
	},
	{
		84, 0, "minecraft:jukebox[has_record=false]",
		0, nil,
	},
	{
		84, 1, "minecraft:jukebox[has_record=true]",
		0, nil,
	},
	{
		85, 0, "minecraft:oak_fence[east=false,south=false,north=false,west=false]",
		0, nil,
	},
	{
		86, 0, "minecraft:carved_pumpkin[facing=south]",
		0, nil,
	},
	{
		86, 1, "minecraft:carved_pumpkin[facing=west]",
		0, nil,
	},
	{
		86, 2, "minecraft:carved_pumpkin[facing=north]",
		0, nil,
	},
	{
		86, 3, "minecraft:carved_pumpkin[facing=east]",
		0, nil,
	},
	{
		87, 0, "minecraft:netherrack",
		0, nil,
	},
	{
		88, 0, "minecraft:soul_sand",
		0, nil,
	},
	{
		89, 0, "minecraft:glowstone",
		0, nil,
	},
	{
		90, 1, "minecraft:nether_portal[axis=x]",
		0, nil,
	},
	{
		90, 2, "minecraft:nether_portal[axis=z]",
		0, nil,
	},
	{
		91, 0, "minecraft:jack_o_lantern[facing=south]",
		0, nil,
	},
	{
		91, 1, "minecraft:jack_o_lantern[facing=west]",
		0, nil,
	},
	{
		91, 2, "minecraft:jack_o_lantern[facing=north]",
		0, nil,
	},
	{
		91, 3, "minecraft:jack_o_lantern[facing=east]",
		0, nil,
	},
	{
		92, 0, "minecraft:cake[bites=0]",
		0, nil,
	},
	{
		92, 1, "minecraft:cake[bites=1]",
		0, nil,
	},
	{
		92, 2, "minecraft:cake[bites=2]",
		0, nil,
	},
	{
		92, 3, "minecraft:cake[bites=3]",
		0, nil,
	},
	{
		92, 4, "minecraft:cake[bites=4]",
		0, nil,
	},
	{
		92, 5, "minecraft:cake[bites=5]",
		0, nil,
	},
	{
		92, 6, "minecraft:cake[bites=6]",
		0, nil,
	},
	{
		93, 0, "minecraft:repeater[delay=1,facing=south,locked=false,powered=false]",
		0, nil,
	},
	{
		93, 1, "minecraft:repeater[delay=1,facing=west,locked=false,powered=false]",
		0, nil,
	},
	{
		93, 2, "minecraft:repeater[delay=1,facing=north,locked=false,powered=false]",
		0, nil,
	},
	{
		93, 3, "minecraft:repeater[delay=1,facing=east,locked=false,powered=false]",
		0, nil,
	},
	{
		93, 4, "minecraft:repeater[delay=2,facing=south,locked=false,powered=false]",
		0, nil,
	},
	{
		93, 5, "minecraft:repeater[delay=2,facing=west,locked=false,powered=false]",
		0, nil,
	},
	{
		93, 6, "minecraft:repeater[delay=2,facing=north,locked=false,powered=false]",
		0, nil,
	},
	{
		93, 7, "minecraft:repeater[delay=2,facing=east,locked=false,powered=false]",
		0, nil,
	},
	{
		93, 8, "minecraft:repeater[delay=3,facing=south,locked=false,powered=false]",
		0, nil,
	},
	{
		93, 9, "minecraft:repeater[delay=3,facing=west,locked=false,powered=false]",
		0, nil,
	},
	{
		93, 10, "minecraft:repeater[delay=3,facing=north,locked=false,powered=false]",
		0, nil,
	},
	{
		93, 11, "minecraft:repeater[delay=3,facing=east,locked=false,powered=false]",
		0, nil,
	},
	{
		93, 12, "minecraft:repeater[delay=4,facing=south,locked=false,powered=false]",
		0, nil,
	},
	{
		93, 13, "minecraft:repeater[delay=4,facing=west,locked=false,powered=false]",
		0, nil,
	},
	{
		93, 14, "minecraft:repeater[delay=4,facing=north,locked=false,powered=false]",
		0, nil,
	},
	{
		93, 15, "minecraft:repeater[delay=4,facing=east,locked=false,powered=false]",
		0, nil,
	},
	{
		94, 0, "minecraft:repeater[delay=1,facing=south,locked=false,powered=true]",
		0, nil,
	},
	{
		94, 1, "minecraft:repeater[delay=1,facing=west,locked=false,powered=true]",
		0, nil,
	},
	{
		94, 2, "minecraft:repeater[delay=1,facing=north,locked=false,powered=true]",
		0, nil,
	},
	{
		94, 3, "minecraft:repeater[delay=1,facing=east,locked=false,powered=true]",
		0, nil,
	},
	{
		94, 4, "minecraft:repeater[delay=2,facing=south,locked=false,powered=true]",
		0, nil,
	},
	{
		94, 5, "minecraft:repeater[delay=2,facing=west,locked=false,powered=true]",
		0, nil,
	},
	{
		94, 6, "minecraft:repeater[delay=2,facing=north,locked=false,powered=true]",
		0, nil,
	},
	{
		94, 7, "minecraft:repeater[delay=2,facing=east,locked=false,powered=true]",
		0, nil,
	},
	{
		94, 8, "minecraft:repeater[delay=3,facing=south,locked=false,powered=true]",
		0, nil,
	},
	{
		94, 9, "minecraft:repeater[delay=3,facing=west,locked=false,powered=true]",
		0, nil,
	},
	{
		94, 10, "minecraft:repeater[delay=3,facing=north,locked=false,powered=true]",
		0, nil,
	},
	{
		94, 11, "minecraft:repeater[delay=3,facing=east,locked=false,powered=true]",
		0, nil,
	},
	{
		94, 12, "minecraft:repeater[delay=4,facing=south,locked=false,powered=true]",
		0, nil,
	},
	{
		94, 13, "minecraft:repeater[delay=4,facing=west,locked=false,powered=true]",
		0, nil,
	},
	{
		94, 14, "minecraft:repeater[delay=4,facing=north,locked=false,powered=true]",
		0, nil,
	},
	{
		94, 15, "minecraft:repeater[delay=4,facing=east,locked=false,powered=true]",
		0, nil,
	},
	{
		95, 0, "minecraft:white_stained_glass",
		0, nil,
	},
	{
		95, 1, "minecraft:orange_stained_glass",
		0, nil,
	},
	{
		95, 2, "minecraft:magenta_stained_glass",
		0, nil,
	},
	{
		95, 3, "minecraft:light_blue_stained_glass",
		0, nil,
	},
	{
		95, 4, "minecraft:yellow_stained_glass",
		0, nil,
	},
	{
		95, 5, "minecraft:lime_stained_glass",
		0, nil,
	},
	{
		95, 6, "minecraft:pink_stained_glass",
		0, nil,
	},
	{
		95, 7, "minecraft:gray_stained_glass",
		0, nil,
	},
	{
		95, 8, "minecraft:light_gray_stained_glass",
		0, nil,
	},
	{
		95, 9, "minecraft:cyan_stained_glass",
		0, nil,
	},
	{
		95, 10, "minecraft:purple_stained_glass",
		0, nil,
	},
	{
		95, 11, "minecraft:blue_stained_glass",
		0, nil,
	},
	{
		95, 12, "minecraft:brown_stained_glass",
		0, nil,
	},
	{
		95, 13, "minecraft:green_stained_glass",
		0, nil,
	},
	{
		95, 14, "minecraft:red_stained_glass",
		0, nil,
	},
	{
		95, 15, "minecraft:black_stained_glass",
		0, nil,
	},
	{
		96, 0, "minecraft:oak_trapdoor[half=bottom,facing=north,open=false,powered=false]",
		0, nil,
	},
	{
		96, 1, "minecraft:oak_trapdoor[half=bottom,facing=south,open=false,powered=false]",
		0, nil,
	},
	{
		96, 2, "minecraft:oak_trapdoor[half=bottom,facing=west,open=false,powered=false]",
		0, nil,
	},
	{
		96, 3, "minecraft:oak_trapdoor[half=bottom,facing=east,open=false,powered=false]",
		0, nil,
	},
	{
		96, 4, "minecraft:oak_trapdoor[half=bottom,facing=north,open=true,powered=true]",
		0, nil,
	},
	{
		96, 5, "minecraft:oak_trapdoor[half=bottom,facing=south,open=true,powered=true]",
		0, nil,
	},
	{
		96, 6, "minecraft:oak_trapdoor[half=bottom,facing=west,open=true,powered=true]",
		0, nil,
	},
	{
		96, 7, "minecraft:oak_trapdoor[half=bottom,facing=east,open=true,powered=true]",
		0, nil,
	},
	{
		96, 8, "minecraft:oak_trapdoor[half=top,facing=north,open=false,powered=false]",
		0, nil,
	},
	{
		96, 9, "minecraft:oak_trapdoor[half=top,facing=south,open=false,powered=false]",
		0, nil,
	},
	{
		96, 10, "minecraft:oak_trapdoor[half=top,facing=west,open=false,powered=false]",
		0, nil,
	},
	{
		96, 11, "minecraft:oak_trapdoor[half=top,facing=east,open=false,powered=false]",
		0, nil,
	},
	{
		96, 12, "minecraft:oak_trapdoor[half=top,facing=north,open=true,powered=true]",
		0, nil,
	},
	{
		96, 13, "minecraft:oak_trapdoor[half=top,facing=south,open=true,powered=true]",
		0, nil,
	},
	{
		96, 14, "minecraft:oak_trapdoor[half=top,facing=west,open=true,powered=true]",
		0, nil,
	},
	{
		96, 15, "minecraft:oak_trapdoor[half=top,facing=east,open=true,powered=true]",
		0, nil,
	},
	{
		97, 0, "minecraft:infested_stone",
		0, nil,
	},
	{
		97, 1, "minecraft:infested_cobblestone",
		0, nil,
	},
	{
		97, 2, "minecraft:infested_stone_bricks",
		0, nil,
	},
	{
		97, 3, "minecraft:infested_mossy_stone_bricks",
		0, nil,
	},
	{
		97, 4, "minecraft:infested_cracked_stone_bricks",
		0, nil,
	},
	{
		97, 5, "minecraft:infested_chiseled_stone_bricks",
		0, nil,
	},
	{
		98, 0, "minecraft:stone_bricks",
		0, nil,
	},
	{
		98, 1, "minecraft:mossy_stone_bricks",
		0, nil,
	},
	{
		98, 2, "minecraft:cracked_stone_bricks",
		0, nil,
	},
	{
		98, 3, "minecraft:chiseled_stone_bricks",
		0, nil,
	},
	{
		99, 0, "minecraft:brown_mushroom_block[north=false,east=false,south=false,west=false,up=false,down=false]",
		0, nil,
	},
	{
		99, 1, "minecraft:brown_mushroom_block[north=true,east=false,south=false,west=true,up=true,down=false]",
		0, nil,
	},
	{
		99, 2, "minecraft:brown_mushroom_block[north=true,east=false,south=false,west=false,up=true,down=false]",
		0, nil,
	},
	{
		99, 3, "minecraft:brown_mushroom_block[north=true,east=true,south=false,west=false,up=true,down=false]",
		0, nil,
	},
	{
		99, 4, "minecraft:brown_mushroom_block[north=false,east=false,south=false,west=true,up=true,down=false]",
		0, nil,
	},
	{
		99, 5, "minecraft:brown_mushroom_block[north=false,east=false,south=false,west=false,up=true,down=false]",
		0, nil,
	},
	{
		99, 6, "minecraft:brown_mushroom_block[north=false,east=true,south=false,west=false,up=true,down=false]",
		0, nil,
	},
	{
		99, 7, "minecraft:brown_mushroom_block[north=false,east=false,south=true,west=true,up=true,down=false]",
		0, nil,
	},
	{
		99, 8, "minecraft:brown_mushroom_block[north=false,east=false,south=true,west=false,up=true,down=false]",
		0, nil,
	},
	{
		99, 9, "minecraft:brown_mushroom_block[north=false,east=true,south=true,west=false,up=true,down=false]",
		0, nil,
	},
	{
		99, 10, "minecraft:mushroom_stem[north=true,east=true,south=true,west=true,up=false,down=false]",
		0, nil,
	},
	{
		99, 14, "minecraft:brown_mushroom_block[north=true,east=true,south=true,west=true,up=true,down=true]",
		0, nil,
	},
	{
		99, 15, "minecraft:mushroom_stem[north=true,east=true,south=true,west=true,up=true,down=true]",
		0, nil,
	},
	{
		100, 0, "minecraft:red_mushroom_block[north=false,east=false,south=false,west=false,up=false,down=false]",
		0, nil,
	},
	{
		100, 1, "minecraft:red_mushroom_block[north=true,east=false,south=false,west=true,up=true,down=false]",
		0, nil,
	},
	{
		100, 2, "minecraft:red_mushroom_block[north=true,east=false,south=false,west=false,up=true,down=false]",
		0, nil,
	},
	{
		100, 3, "minecraft:red_mushroom_block[north=true,east=true,south=false,west=false,up=true,down=false]",
		0, nil,
	},
	{
		100, 4, "minecraft:red_mushroom_block[north=false,east=false,south=false,west=true,up=true,down=false]",
		0, nil,
	},
	{
		100, 5, "minecraft:red_mushroom_block[north=false,east=false,south=false,west=false,up=true,down=false]",
		0, nil,
	},
	{
		100, 6, "minecraft:red_mushroom_block[north=false,east=true,south=false,west=false,up=true,down=false]",
		0, nil,
	},
	{
		100, 7, "minecraft:red_mushroom_block[north=false,east=false,south=true,west=true,up=true,down=false]",
		0, nil,
	},
	{
		100, 8, "minecraft:red_mushroom_block[north=false,east=false,south=true,west=false,up=true,down=false]",
		0, nil,
	},
	{
		100, 9, "minecraft:red_mushroom_block[north=false,east=true,south=true,west=false,up=true,down=false]",
		0, nil,
	},
	{
		100, 10, "minecraft:mushroom_stem[north=true,east=true,south=true,west=true,up=false,down=false]",
		0, nil,
	},
	{
		100, 14, "minecraft:red_mushroom_block[north=true,east=true,south=true,west=true,up=true,down=true]",
		0, nil,
	},
	{
		100, 15, "minecraft:mushroom_stem[north=true,east=true,south=true,west=true,up=true,down=true]",
		0, nil,
	},
	{
		101, 0, "minecraft:iron_bars[east=false,south=false,north=false,west=false]",
		0, nil,
	},
	{
		102, 0, "minecraft:glass_pane[east=false,south=false,north=false,west=false]",
		0, nil,
	},
	{
		103, 0, "minecraft:melon",
		0, nil,
	},
	{
		104, 0, "minecraft:pumpkin_stem[age=0]",
		0, nil,
	},
	{
		104, 1, "minecraft:pumpkin_stem[age=1]",
		0, nil,
	},
	{
		104, 2, "minecraft:pumpkin_stem[age=2]",
		0, nil,
	},
	{
		104, 3, "minecraft:pumpkin_stem[age=3]",
		0, nil,
	},
	{
		104, 4, "minecraft:pumpkin_stem[age=4]",
		0, nil,
	},
	{
		104, 5, "minecraft:pumpkin_stem[age=5]",
		0, nil,
	},
	{
		104, 6, "minecraft:pumpkin_stem[age=6]",
		0, nil,
	},
	{
		104, 7, "minecraft:pumpkin_stem[age=7]",
		0, nil,
	},
	{
		105, 0, "minecraft:melon_stem[age=0]",
		0, nil,
	},
	{
		105, 1, "minecraft:melon_stem[age=1]",
		0, nil,
	},
	{
		105, 2, "minecraft:melon_stem[age=2]",
		0, nil,
	},
	{
		105, 3, "minecraft:melon_stem[age=3]",
		0, nil,
	},
	{
		105, 4, "minecraft:melon_stem[age=4]",
		0, nil,
	},
	{
		105, 5, "minecraft:melon_stem[age=5]",
		0, nil,
	},
	{
		105, 6, "minecraft:melon_stem[age=6]",
		0, nil,
	},
	{
		105, 7, "minecraft:melon_stem[age=7]",
		0, nil,
	},
	{
		106, 0, "minecraft:vine[east=false,south=false,north=false,west=false,up=false]",
		0, nil,
	},
	{
		106, 1, "minecraft:vine[east=false,south=true,north=false,west=false,up=false]",
		0, nil,
	},
	{
		106, 2, "minecraft:vine[east=false,south=false,north=false,west=true,up=false]",
		0, nil,
	},
	{
		106, 3, "minecraft:vine[east=false,south=true,north=false,west=true,up=false]",
		0, nil,
	},
	{
		106, 4, "minecraft:vine[east=false,south=false,north=true,west=false,up=false]",
		0, nil,
	},
	{
		106, 5, "minecraft:vine[east=false,south=true,north=true,west=false,up=false]",
		0, nil,
	},
	{
		106, 6, "minecraft:vine[east=false,south=false,north=true,west=true,up=false]",
		0, nil,
	},
	{
		106, 7, "minecraft:vine[east=false,south=true,north=true,west=true,up=false]",
		0, nil,
	},
	{
		106, 8, "minecraft:vine[east=true,south=false,north=false,west=false,up=false]",
		0, nil,
	},
	{
		106, 9, "minecraft:vine[east=true,south=true,north=false,west=false,up=false]",
		0, nil,
	},
	{
		106, 10, "minecraft:vine[east=true,south=false,north=false,west=true,up=false]",
		0, nil,
	},
	{
		106, 11, "minecraft:vine[east=true,south=true,north=false,west=true,up=false]",
		0, nil,
	},
	{
		106, 12, "minecraft:vine[east=true,south=false,north=true,west=false,up=false]",
		0, nil,
	},
	{
		106, 13, "minecraft:vine[east=true,south=true,north=true,west=false,up=false]",
		0, nil,
	},
	{
		106, 14, "minecraft:vine[east=true,south=false,north=true,west=true,up=false]",
		0, nil,
	},
	{
		106, 15, "minecraft:vine[east=true,south=true,north=true,west=true,up=false]",
		0, nil,
	},
	{
		107, 0, "minecraft:oak_fence_gate[in_wall=false,powered=false,facing=south,open=false]",
		0, nil,
	},
	{
		107, 1, "minecraft:oak_fence_gate[in_wall=false,powered=false,facing=west,open=false]",
		0, nil,
	},
	{
		107, 2, "minecraft:oak_fence_gate[in_wall=false,powered=false,facing=north,open=false]",
		0, nil,
	},
	{
		107, 3, "minecraft:oak_fence_gate[in_wall=false,powered=false,facing=east,open=false]",
		0, nil,
	},
	{
		107, 4, "minecraft:oak_fence_gate[in_wall=false,powered=false,facing=south,open=true]",
		0, nil,
	},
	{
		107, 5, "minecraft:oak_fence_gate[in_wall=false,powered=false,facing=west,open=true]",
		0, nil,
	},
	{
		107, 6, "minecraft:oak_fence_gate[in_wall=false,powered=false,facing=north,open=true]",
		0, nil,
	},
	{
		107, 7, "minecraft:oak_fence_gate[in_wall=false,powered=false,facing=east,open=true]",
		0, nil,
	},
	{
		107, 8, "minecraft:oak_fence_gate[in_wall=false,powered=true,facing=south,open=false]",
		0, nil,
	},
	{
		107, 9, "minecraft:oak_fence_gate[in_wall=false,powered=true,facing=west,open=false]",
		0, nil,
	},
	{
		107, 10, "minecraft:oak_fence_gate[in_wall=false,powered=true,facing=north,open=false]",
		0, nil,
	},
	{
		107, 11, "minecraft:oak_fence_gate[in_wall=false,powered=true,facing=east,open=false]",
		0, nil,
	},
	{
		107, 12, "minecraft:oak_fence_gate[in_wall=false,powered=true,facing=south,open=true]",
		0, nil,
	},
	{
		107, 13, "minecraft:oak_fence_gate[in_wall=false,powered=true,facing=west,open=true]",
		0, nil,
	},
	{
		107, 14, "minecraft:oak_fence_gate[in_wall=false,powered=true,facing=north,open=true]",
		0, nil,
	},
	{
		107, 15, "minecraft:oak_fence_gate[in_wall=false,powered=true,facing=east,open=true]",
		0, nil,
	},
	{
		108, 0, "minecraft:brick_stairs[half=bottom,shape=straight,facing=east]",
		0, nil,
	},
	{
		108, 1, "minecraft:brick_stairs[half=bottom,shape=straight,facing=west]",
		0, nil,
	},
	{
		108, 2, "minecraft:brick_stairs[half=bottom,shape=straight,facing=south]",
		0, nil,
	},
	{
		108, 3, "minecraft:brick_stairs[half=bottom,shape=straight,facing=north]",
		0, nil,
	},
	{
		108, 4, "minecraft:brick_stairs[half=top,shape=straight,facing=east]",
		0, nil,
	},
	{
		108, 5, "minecraft:brick_stairs[half=top,shape=straight,facing=west]",
		0, nil,
	},
	{
		108, 6, "minecraft:brick_stairs[half=top,shape=straight,facing=south]",
		0, nil,
	},
	{
		108, 7, "minecraft:brick_stairs[half=top,shape=straight,facing=north]",
		0, nil,
	},
	{
		109, 0, "minecraft:stone_brick_stairs[half=bottom,shape=straight,facing=east]",
		0, nil,
	},
	{
		109, 1, "minecraft:stone_brick_stairs[half=bottom,shape=straight,facing=west]",
		0, nil,
	},
	{
		109, 2, "minecraft:stone_brick_stairs[half=bottom,shape=straight,facing=south]",
		0, nil,
	},
	{
		109, 3, "minecraft:stone_brick_stairs[half=bottom,shape=straight,facing=north]",
		0, nil,
	},
	{
		109, 4, "minecraft:stone_brick_stairs[half=top,shape=straight,facing=east]",
		0, nil,
	},
	{
		109, 5, "minecraft:stone_brick_stairs[half=top,shape=straight,facing=west]",
		0, nil,
	},
	{
		109, 6, "minecraft:stone_brick_stairs[half=top,shape=straight,facing=south]",
		0, nil,
	},
	{
		109, 7, "minecraft:stone_brick_stairs[half=top,shape=straight,facing=north]",
		0, nil,
	},
	{
		110, 0, "minecraft:mycelium[snowy=false]",
		0, nil,
	},
	{
		111, 0, "minecraft:lily_pad",
		0, nil,
	},
	{
		112, 0, "minecraft:nether_bricks",
		0, nil,
	},
	{
		113, 0, "minecraft:nether_brick_fence[east=false,south=false,north=false,west=false]",
		0, nil,
	},
	{
		114, 0, "minecraft:nether_brick_stairs[half=bottom,shape=straight,facing=east]",
		0, nil,
	},
	{
		114, 1, "minecraft:nether_brick_stairs[half=bottom,shape=straight,facing=west]",
		0, nil,
	},
	{
		114, 2, "minecraft:nether_brick_stairs[half=bottom,shape=straight,facing=south]",
		0, nil,
	},
	{
		114, 3, "minecraft:nether_brick_stairs[half=bottom,shape=straight,facing=north]",
		0, nil,
	},
	{
		114, 4, "minecraft:nether_brick_stairs[half=top,shape=straight,facing=east]",
		0, nil,
	},
	{
		114, 5, "minecraft:nether_brick_stairs[half=top,shape=straight,facing=west]",
		0, nil,
	},
	{
		114, 6, "minecraft:nether_brick_stairs[half=top,shape=straight,facing=south]",
		0, nil,
	},
	{
		114, 7, "minecraft:nether_brick_stairs[half=top,shape=straight,facing=north]",
		0, nil,
	},
	{
		115, 0, "minecraft:nether_wart[age=0]",
		0, nil,
	},
	{
		115, 1, "minecraft:nether_wart[age=1]",
		0, nil,
	},
	{
		115, 2, "minecraft:nether_wart[age=2]",
		0, nil,
	},
	{
		115, 3, "minecraft:nether_wart[age=3]",
		0, nil,
	},
	{
		116, 0, "minecraft:enchanting_table",
		0, nil,
	},
	{
		117, 0, "minecraft:brewing_stand[has_bottle_0=false,has_bottle_1=false,has_bottle_2=false]",
		0, nil,
	},
	{
		117, 1, "minecraft:brewing_stand[has_bottle_0=true,has_bottle_1=false,has_bottle_2=false]",
		0, nil,
	},
	{
		117, 2, "minecraft:brewing_stand[has_bottle_0=false,has_bottle_1=true,has_bottle_2=false]",
		0, nil,
	},
	{
		117, 3, "minecraft:brewing_stand[has_bottle_0=true,has_bottle_1=true,has_bottle_2=false]",
		0, nil,
	},
	{
		117, 4, "minecraft:brewing_stand[has_bottle_0=false,has_bottle_1=false,has_bottle_2=true]",
		0, nil,
	},
	{
		117, 5, "minecraft:brewing_stand[has_bottle_0=true,has_bottle_1=false,has_bottle_2=true]",
		0, nil,
	},
	{
		117, 6, "minecraft:brewing_stand[has_bottle_0=false,has_bottle_1=true,has_bottle_2=true]",
		0, nil,
	},
	{
		117, 7, "minecraft:brewing_stand[has_bottle_0=true,has_bottle_1=true,has_bottle_2=true]",
		0, nil,
	},
	{
		118, 0, "minecraft:cauldron[level=0]",
		0, nil,
	},
	{
		118, 1, "minecraft:cauldron[level=1]",
		0, nil,
	},
	{
		118, 2, "minecraft:cauldron[level=2]",
		0, nil,
	},
	{
		118, 3, "minecraft:cauldron[level=3]",
		0, nil,
	},
	{
		119, 0, "minecraft:end_portal",
		0, nil,
	},
	{
		120, 0, "minecraft:end_portal_frame[eye=false,facing=south]",
		0, nil,
	},
	{
		120, 1, "minecraft:end_portal_frame[eye=false,facing=west]",
		0, nil,
	},
	{
		120, 2, "minecraft:end_portal_frame[eye=false,facing=north]",
		0, nil,
	},
	{
		120, 3, "minecraft:end_portal_frame[eye=false,facing=east]",
		0, nil,
	},
	{
		120, 4, "minecraft:end_portal_frame[eye=true,facing=south]",
		0, nil,
	},
	{
		120, 5, "minecraft:end_portal_frame[eye=true,facing=west]",
		0, nil,
	},
	{
		120, 6, "minecraft:end_portal_frame[eye=true,facing=north]",
		0, nil,
	},
	{
		120, 7, "minecraft:end_portal_frame[eye=true,facing=east]",
		0, nil,
	},
	{
		121, 0, "minecraft:end_stone",
		0, nil,
	},
	{
		122, 0, "minecraft:dragon_egg",
		0, nil,
	},
	{
		123, 0, "minecraft:redstone_lamp[lit=false]",
		0, nil,
	},
	{
		124, 0, "minecraft:redstone_lamp[lit=true]",
		0, nil,
	},
	{
		125, 0, "minecraft:oak_slab[type=double]",
		0, nil,
	},
	{
		125, 1, "minecraft:spruce_slab[type=double]",
		0, nil,
	},
	{
		125, 2, "minecraft:birch_slab[type=double]",
		0, nil,
	},
	{
		125, 3, "minecraft:jungle_slab[type=double]",
		0, nil,
	},
	{
		125, 4, "minecraft:acacia_slab[type=double]",
		0, nil,
	},
	{
		125, 5, "minecraft:dark_oak_slab[type=double]",
		0, nil,
	},
	{
		126, 0, "minecraft:oak_slab[type=bottom]",
		0, nil,
	},
	{
		126, 1, "minecraft:spruce_slab[type=bottom]",
		0, nil,
	},
	{
		126, 2, "minecraft:birch_slab[type=bottom]",
		0, nil,
	},
	{
		126, 3, "minecraft:jungle_slab[type=bottom]",
		0, nil,
	},
	{
		126, 4, "minecraft:acacia_slab[type=bottom]",
		0, nil,
	},
	{
		126, 5, "minecraft:dark_oak_slab[type=bottom]",
		0, nil,
	},
	{
		126, 8, "minecraft:oak_slab[type=top]",
		0, nil,
	},
	{
		126, 9, "minecraft:spruce_slab[type=top]",
		0, nil,
	},
	{
		126, 10, "minecraft:birch_slab[type=top]",
		0, nil,
	},
	{
		126, 11, "minecraft:jungle_slab[type=top]",
		0, nil,
	},
	{
		126, 12, "minecraft:acacia_slab[type=top]",
		0, nil,
	},
	{
		126, 13, "minecraft:dark_oak_slab[type=top]",
		0, nil,
	},
	{
		127, 0, "minecraft:cocoa[facing=south,age=0]",
		0, nil,
	},
	{
		127, 1, "minecraft:cocoa[facing=west,age=0]",
		0, nil,
	},
	{
		127, 2, "minecraft:cocoa[facing=north,age=0]",
		0, nil,
	},
	{
		127, 3, "minecraft:cocoa[facing=east,age=0]",
		0, nil,
	},
	{
		127, 4, "minecraft:cocoa[facing=south,age=1]",
		0, nil,
	},
	{
		127, 5, "minecraft:cocoa[facing=west,age=1]",
		0, nil,
	},
	{
		127, 6, "minecraft:cocoa[facing=north,age=1]",
		0, nil,
	},
	{
		127, 7, "minecraft:cocoa[facing=east,age=1]",
		0, nil,
	},
	{
		127, 8, "minecraft:cocoa[facing=south,age=2]",
		0, nil,
	},
	{
		127, 9, "minecraft:cocoa[facing=west,age=2]",
		0, nil,
	},
	{
		127, 10, "minecraft:cocoa[facing=north,age=2]",
		0, nil,
	},
	{
		127, 11, "minecraft:cocoa[facing=east,age=2]",
		0, nil,
	},
	{
		128, 0, "minecraft:sandstone_stairs[half=bottom,shape=straight,facing=east]",
		0, nil,
	},
	{
		128, 1, "minecraft:sandstone_stairs[half=bottom,shape=straight,facing=west]",
		0, nil,
	},
	{
		128, 2, "minecraft:sandstone_stairs[half=bottom,shape=straight,facing=south]",
		0, nil,
	},
	{
		128, 3, "minecraft:sandstone_stairs[half=bottom,shape=straight,facing=north]",
		0, nil,
	},
	{
		128, 4, "minecraft:sandstone_stairs[half=top,shape=straight,facing=east]",
		0, nil,
	},
	{
		128, 5, "minecraft:sandstone_stairs[half=top,shape=straight,facing=west]",
		0, nil,
	},
	{
		128, 6, "minecraft:sandstone_stairs[half=top,shape=straight,facing=south]",
		0, nil,
	},
	{
		128, 7, "minecraft:sandstone_stairs[half=top,shape=straight,facing=north]",
		0, nil,
	},
	{
		129, 0, "minecraft:emerald_ore",
		0, nil,
	},
	{
		130, 2, "minecraft:ender_chest[facing=north]",
		0, nil,
	},
	{
		130, 3, "minecraft:ender_chest[facing=south]",
		0, nil,
	},
	{
		130, 4, "minecraft:ender_chest[facing=west]",
		0, nil,
	},
	{
		130, 5, "minecraft:ender_chest[facing=east]",
		0, nil,
	},
	{
		131, 0, "minecraft:tripwire_hook[powered=false,attached=false,facing=south]",
		0, nil,
	},
	{
		131, 1, "minecraft:tripwire_hook[powered=false,attached=false,facing=west]",
		0, nil,
	},
	{
		131, 2, "minecraft:tripwire_hook[powered=false,attached=false,facing=north]",
		0, nil,
	},
	{
		131, 3, "minecraft:tripwire_hook[powered=false,attached=false,facing=east]",
		0, nil,
	},
	{
		131, 4, "minecraft:tripwire_hook[powered=false,attached=true,facing=south]",
		0, nil,
	},
	{
		131, 5, "minecraft:tripwire_hook[powered=false,attached=true,facing=west]",
		0, nil,
	},
	{
		131, 6, "minecraft:tripwire_hook[powered=false,attached=true,facing=north]",
		0, nil,
	},
	{
		131, 7, "minecraft:tripwire_hook[powered=false,attached=true,facing=east]",
		0, nil,
	},
	{
		131, 8, "minecraft:tripwire_hook[powered=true,attached=false,facing=south]",
		0, nil,
	},
	{
		131, 9, "minecraft:tripwire_hook[powered=true,attached=false,facing=west]",
		0, nil,
	},
	{
		131, 10, "minecraft:tripwire_hook[powered=true,attached=false,facing=north]",
		0, nil,
	},
	{
		131, 11, "minecraft:tripwire_hook[powered=true,attached=false,facing=east]",
		0, nil,
	},
	{
		131, 12, "minecraft:tripwire_hook[powered=true,attached=true,facing=south]",
		0, nil,
	},
	{
		131, 13, "minecraft:tripwire_hook[powered=true,attached=true,facing=west]",
		0, nil,
	},
	{
		131, 14, "minecraft:tripwire_hook[powered=true,attached=true,facing=north]",
		0, nil,
	},
	{
		131, 15, "minecraft:tripwire_hook[powered=true,attached=true,facing=east]",
		0, nil,
	},
	{
		132, 0, "minecraft:tripwire[disarmed=false,east=false,powered=false,south=false,north=false,west=false,attached=false]",
		0, nil,
	},
	{
		132, 1, "minecraft:tripwire[disarmed=false,east=false,powered=true,south=false,north=false,west=false,attached=false]",
		0, nil,
	},
	{
		132, 4, "minecraft:tripwire[disarmed=false,east=false,powered=false,south=false,north=false,west=false,attached=true]",
		0, nil,
	},
	{
		132, 5, "minecraft:tripwire[disarmed=false,east=false,powered=true,south=false,north=false,west=false,attached=true]",
		0, nil,
	},
	{
		132, 8, "minecraft:tripwire[disarmed=true,east=false,powered=false,south=false,north=false,west=false,attached=false]",
		0, nil,
	},
	{
		132, 9, "minecraft:tripwire[disarmed=true,east=false,powered=true,south=false,north=false,west=false,attached=false]",
		0, nil,
	},
	{
		132, 12, "minecraft:tripwire[disarmed=true,east=false,powered=false,south=false,north=false,west=false,attached=true]",
		0, nil,
	},
	{
		132, 13, "minecraft:tripwire[disarmed=true,east=false,powered=true,south=false,north=false,west=false,attached=true]",
		0, nil,
	},
	{
		133, 0, "minecraft:emerald_block",
		0, nil,
	},
	{
		134, 0, "minecraft:spruce_stairs[half=bottom,shape=straight,facing=east]",
		0, nil,
	},
	{
		134, 1, "minecraft:spruce_stairs[half=bottom,shape=straight,facing=west]",
		0, nil,
	},
	{
		134, 2, "minecraft:spruce_stairs[half=bottom,shape=straight,facing=south]",
		0, nil,
	},
	{
		134, 3, "minecraft:spruce_stairs[half=bottom,shape=straight,facing=north]",
		0, nil,
	},
	{
		134, 4, "minecraft:spruce_stairs[half=top,shape=straight,facing=east]",
		0, nil,
	},
	{
		134, 5, "minecraft:spruce_stairs[half=top,shape=straight,facing=west]",
		0, nil,
	},
	{
		134, 6, "minecraft:spruce_stairs[half=top,shape=straight,facing=south]",
		0, nil,
	},
	{
		134, 7, "minecraft:spruce_stairs[half=top,shape=straight,facing=north]",
		0, nil,
	},
	{
		135, 0, "minecraft:birch_stairs[half=bottom,shape=straight,facing=east]",
		0, nil,
	},
	{
		135, 1, "minecraft:birch_stairs[half=bottom,shape=straight,facing=west]",
		0, nil,
	},
	{
		135, 2, "minecraft:birch_stairs[half=bottom,shape=straight,facing=south]",
		0, nil,
	},
	{
		135, 3, "minecraft:birch_stairs[half=bottom,shape=straight,facing=north]",
		0, nil,
	},
	{
		135, 4, "minecraft:birch_stairs[half=top,shape=straight,facing=east]",
		0, nil,
	},
	{
		135, 5, "minecraft:birch_stairs[half=top,shape=straight,facing=west]",
		0, nil,
	},
	{
		135, 6, "minecraft:birch_stairs[half=top,shape=straight,facing=south]",
		0, nil,
	},
	{
		135, 7, "minecraft:birch_stairs[half=top,shape=straight,facing=north]",
		0, nil,
	},
	{
		136, 0, "minecraft:jungle_stairs[half=bottom,shape=straight,facing=east]",
		0, nil,
	},
	{
		136, 1, "minecraft:jungle_stairs[half=bottom,shape=straight,facing=west]",
		0, nil,
	},
	{
		136, 2, "minecraft:jungle_stairs[half=bottom,shape=straight,facing=south]",
		0, nil,
	},
	{
		136, 3, "minecraft:jungle_stairs[half=bottom,shape=straight,facing=north]",
		0, nil,
	},
	{
		136, 4, "minecraft:jungle_stairs[half=top,shape=straight,facing=east]",
		0, nil,
	},
	{
		136, 5, "minecraft:jungle_stairs[half=top,shape=straight,facing=west]",
		0, nil,
	},
	{
		136, 6, "minecraft:jungle_stairs[half=top,shape=straight,facing=south]",
		0, nil,
	},
	{
		136, 7, "minecraft:jungle_stairs[half=top,shape=straight,facing=north]",
		0, nil,
	},
	{
		137, 0, "minecraft:command_block[conditional=false,facing=down]",
		0, nil,
	},
	{
		137, 1, "minecraft:command_block[conditional=false,facing=up]",
		0, nil,
	},
	{
		137, 2, "minecraft:command_block[conditional=false,facing=north]",
		0, nil,
	},
	{
		137, 3, "minecraft:command_block[conditional=false,facing=south]",
		0, nil,
	},
	{
		137, 4, "minecraft:command_block[conditional=false,facing=west]",
		0, nil,
	},
	{
		137, 5, "minecraft:command_block[conditional=false,facing=east]",
		0, nil,
	},
	{
		137, 8, "minecraft:command_block[conditional=true,facing=down]",
		0, nil,
	},
	{
		137, 9, "minecraft:command_block[conditional=true,facing=up]",
		0, nil,
	},
	{
		137, 10, "minecraft:command_block[conditional=true,facing=north]",
		0, nil,
	},
	{
		137, 11, "minecraft:command_block[conditional=true,facing=south]",
		0, nil,
	},
	{
		137, 12, "minecraft:command_block[conditional=true,facing=west]",
		0, nil,
	},
	{
		137, 13, "minecraft:command_block[conditional=true,facing=east]",
		0, nil,
	},
	{
		138, 0, "minecraft:beacon",
		0, nil,
	},
	{
		139, 0, "minecraft:cobblestone_wall[east=false,south=false,north=false,west=false,up=false]",
		0, nil,
	},
	{
		139, 1, "minecraft:mossy_cobblestone_wall[east=false,south=false,north=false,west=false,up=false]",
		0, nil,
	},
	{
		140, 0, "minecraft:flower_pot",
		0, nil,
	},
	{
		140, 1, "minecraft:potted_poppy",
		0, nil,
	},
	{
		140, 2, "minecraft:potted_dandelion",
		0, nil,
	},
	{
		140, 3, "minecraft:potted_oak_sapling",
		0, nil,
	},
	{
		140, 4, "minecraft:potted_spruce_sapling",
		0, nil,
	},
	{
		140, 5, "minecraft:potted_birch_sapling",
		0, nil,
	},
	{
		140, 6, "minecraft:potted_jungle_sapling",
		0, nil,
	},
	{
		140, 7, "minecraft:potted_red_mushroom",
		0, nil,
	},
	{
		140, 8, "minecraft:potted_brown_mushroom",
		0, nil,
	},
	{
		140, 9, "minecraft:potted_cactus",
		0, nil,
	},
	{
		140, 10, "minecraft:potted_dead_bush",
		0, nil,
	},
	{
		140, 11, "minecraft:potted_fern",
		0, nil,
	},
	{
		140, 12, "minecraft:potted_acacia_sapling",
		0, nil,
	},
	{
		140, 13, "minecraft:potted_dark_oak_sapling",
		0, nil,
	},
	{
		140, 14, "minecraft:potted_blue_orchid",
		0, nil,
	},
	{
		140, 15, "minecraft:potted_allium",
		0, nil,
	},
	{
		141, 0, "minecraft:carrots[age=0]",
		0, nil,
	},
	{
		141, 1, "minecraft:carrots[age=1]",
		0, nil,
	},
	{
		141, 2, "minecraft:carrots[age=2]",
		0, nil,
	},
	{
		141, 3, "minecraft:carrots[age=3]",
		0, nil,
	},
	{
		141, 4, "minecraft:carrots[age=4]",
		0, nil,
	},
	{
		141, 5, "minecraft:carrots[age=5]",
		0, nil,
	},
	{
		141, 6, "minecraft:carrots[age=6]",
		0, nil,
	},
	{
		141, 7, "minecraft:carrots[age=7]",
		0, nil,
	},
	{
		142, 0, "minecraft:potatoes[age=0]",
		0, nil,
	},
	{
		142, 1, "minecraft:potatoes[age=1]",
		0, nil,
	},
	{
		142, 2, "minecraft:potatoes[age=2]",
		0, nil,
	},
	{
		142, 3, "minecraft:potatoes[age=3]",
		0, nil,
	},
	{
		142, 4, "minecraft:potatoes[age=4]",
		0, nil,
	},
	{
		142, 5, "minecraft:potatoes[age=5]",
		0, nil,
	},
	{
		142, 6, "minecraft:potatoes[age=6]",
		0, nil,
	},
	{
		142, 7, "minecraft:potatoes[age=7]",
		0, nil,
	},
	{
		143, 0, "minecraft:oak_button[powered=false,facing=east,face=ceiling]",
		0, nil,
	},
	{
		143, 1, "minecraft:oak_button[powered=false,facing=east,face=wall]",
		0, nil,
	},
	{
		143, 2, "minecraft:oak_button[powered=false,facing=west,face=wall]",
		0, nil,
	},
	{
		143, 3, "minecraft:oak_button[powered=false,facing=south,face=wall]",
		0, nil,
	},
	{
		143, 4, "minecraft:oak_button[powered=false,facing=north,face=wall]",
		0, nil,
	},
	{
		143, 5, "minecraft:oak_button[powered=false,facing=east,face=floor]",
		0, nil,
	},
	{
		143, 8, "minecraft:oak_button[powered=true,facing=south,face=ceiling]",
		0, nil,
	},
	{
		143, 9, "minecraft:oak_button[powered=true,facing=east,face=wall]",
		0, nil,
	},
	{
		143, 10, "minecraft:oak_button[powered=true,facing=west,face=wall]",
		0, nil,
	},
	{
		143, 11, "minecraft:oak_button[powered=true,facing=south,face=wall]",
		0, nil,
	},
	{
		143, 12, "minecraft:oak_button[powered=true,facing=north,face=wall]",
		0, nil,
	},
	{
		143, 13, "minecraft:oak_button[powered=true,facing=south,face=floor]",
		0, nil,
	},
	{
		144, 0, "minecraft:skeleton_skull[rotation=0]",
		0, nil,
	},
	{
		144, 1, "minecraft:skeleton_skull[rotation=4]",
		0, nil,
	},
	{
		144, 2, "minecraft:skeleton_wall_skull[facing=north]",
		0, nil,
	},
	{
		144, 3, "minecraft:skeleton_wall_skull[facing=south]",
		0, nil,
	},
	{
		144, 4, "minecraft:skeleton_wall_skull[facing=west]",
		0, nil,
	},
	{
		144, 5, "minecraft:skeleton_wall_skull[facing=east]",
		0, nil,
	},
	{
		144, 8, "minecraft:skeleton_skull[rotation=8]",
		0, nil,
	},
	{
		144, 9, "minecraft:skeleton_skull[rotation=12]",
		0, nil,
	},
	{
		144, 10, "minecraft:skeleton_wall_skull[facing=north]",
		0, nil,
	},
	{
		144, 11, "minecraft:skeleton_wall_skull[facing=south]",
		0, nil,
	},
	{
		144, 12, "minecraft:skeleton_wall_skull[facing=west]",
		0, nil,
	},
	{
		144, 13, "minecraft:skeleton_wall_skull[facing=east]",
		0, nil,
	},
	{
		145, 0, "minecraft:anvil[facing=south]",
		0, nil,
	},
	{
		145, 1, "minecraft:anvil[facing=west]",
		0, nil,
	},
	{
		145, 2, "minecraft:anvil[facing=north]",
		0, nil,
	},
	{
		145, 3, "minecraft:anvil[facing=east]",
		0, nil,
	},
	{
		145, 4, "minecraft:chipped_anvil[facing=south]",
		0, nil,
	},
	{
		145, 5, "minecraft:chipped_anvil[facing=west]",
		0, nil,
	},
	{
		145, 6, "minecraft:chipped_anvil[facing=north]",
		0, nil,
	},
	{
		145, 7, "minecraft:chipped_anvil[facing=east]",
		0, nil,
	},
	{
		145, 8, "minecraft:damaged_anvil[facing=south]",
		0, nil,
	},
	{
		145, 9, "minecraft:damaged_anvil[facing=west]",
		0, nil,
	},
	{
		145, 10, "minecraft:damaged_anvil[facing=north]",
		0, nil,
	},
	{
		145, 11, "minecraft:damaged_anvil[facing=east]",
		0, nil,
	},
	{
		146, 2, "minecraft:trapped_chest[facing=north,type=single]",
		0, nil,
	},
	{
		146, 3, "minecraft:trapped_chest[facing=south,type=single]",
		0, nil,
	},
	{
		146, 4, "minecraft:trapped_chest[facing=west,type=single]",
		0, nil,
	},
	{
		146, 5, "minecraft:trapped_chest[facing=east,type=single]",
		0, nil,
	},
	{
		147, 0, "minecraft:light_weighted_pressure_plate[power=0]",
		0, nil,
	},
	{
		147, 1, "minecraft:light_weighted_pressure_plate[power=1]",
		0, nil,
	},
	{
		147, 2, "minecraft:light_weighted_pressure_plate[power=2]",
		0, nil,
	},
	{
		147, 3, "minecraft:light_weighted_pressure_plate[power=3]",
		0, nil,
	},
	{
		147, 4, "minecraft:light_weighted_pressure_plate[power=4]",
		0, nil,
	},
	{
		147, 5, "minecraft:light_weighted_pressure_plate[power=5]",
		0, nil,
	},
	{
		147, 6, "minecraft:light_weighted_pressure_plate[power=6]",
		0, nil,
	},
	{
		147, 7, "minecraft:light_weighted_pressure_plate[power=7]",
		0, nil,
	},
	{
		147, 8, "minecraft:light_weighted_pressure_plate[power=8]",
		0, nil,
	},
	{
		147, 9, "minecraft:light_weighted_pressure_plate[power=9]",
		0, nil,
	},
	{
		147, 10, "minecraft:light_weighted_pressure_plate[power=10]",
		0, nil,
	},
	{
		147, 11, "minecraft:light_weighted_pressure_plate[power=11]",
		0, nil,
	},
	{
		147, 12, "minecraft:light_weighted_pressure_plate[power=12]",
		0, nil,
	},
	{
		147, 13, "minecraft:light_weighted_pressure_plate[power=13]",
		0, nil,
	},
	{
		147, 14, "minecraft:light_weighted_pressure_plate[power=14]",
		0, nil,
	},
	{
		147, 15, "minecraft:light_weighted_pressure_plate[power=15]",
		0, nil,
	},
	{
		148, 0, "minecraft:heavy_weighted_pressure_plate[power=0]",
		0, nil,
	},
	{
		148, 1, "minecraft:heavy_weighted_pressure_plate[power=1]",
		0, nil,
	},
	{
		148, 2, "minecraft:heavy_weighted_pressure_plate[power=2]",
		0, nil,
	},
	{
		148, 3, "minecraft:heavy_weighted_pressure_plate[power=3]",
		0, nil,
	},
	{
		148, 4, "minecraft:heavy_weighted_pressure_plate[power=4]",
		0, nil,
	},
	{
		148, 5, "minecraft:heavy_weighted_pressure_plate[power=5]",
		0, nil,
	},
	{
		148, 6, "minecraft:heavy_weighted_pressure_plate[power=6]",
		0, nil,
	},
	{
		148, 7, "minecraft:heavy_weighted_pressure_plate[power=7]",
		0, nil,
	},
	{
		148, 8, "minecraft:heavy_weighted_pressure_plate[power=8]",
		0, nil,
	},
	{
		148, 9, "minecraft:heavy_weighted_pressure_plate[power=9]",
		0, nil,
	},
	{
		148, 10, "minecraft:heavy_weighted_pressure_plate[power=10]",
		0, nil,
	},
	{
		148, 11, "minecraft:heavy_weighted_pressure_plate[power=11]",
		0, nil,
	},
	{
		148, 12, "minecraft:heavy_weighted_pressure_plate[power=12]",
		0, nil,
	},
	{
		148, 13, "minecraft:heavy_weighted_pressure_plate[power=13]",
		0, nil,
	},
	{
		148, 14, "minecraft:heavy_weighted_pressure_plate[power=14]",
		0, nil,
	},
	{
		148, 15, "minecraft:heavy_weighted_pressure_plate[power=15]",
		0, nil,
	},
	{
		149, 0, "minecraft:comparator[mode=compare,powered=false,facing=south]",
		0, nil,
	},
	{
		149, 1, "minecraft:comparator[mode=compare,powered=false,facing=west]",
		0, nil,
	},
	{
		149, 2, "minecraft:comparator[mode=compare,powered=false,facing=north]",
		0, nil,
	},
	{
		149, 3, "minecraft:comparator[mode=compare,powered=false,facing=east]",
		0, nil,
	},
	{
		149, 4, "minecraft:comparator[mode=subtract,powered=false,facing=south]",
		0, nil,
	},
	{
		149, 5, "minecraft:comparator[mode=subtract,powered=false,facing=west]",
		0, nil,
	},
	{
		149, 6, "minecraft:comparator[mode=subtract,powered=false,facing=north]",
		0, nil,
	},
	{
		149, 7, "minecraft:comparator[mode=subtract,powered=false,facing=east]",
		0, nil,
	},
	{
		149, 8, "minecraft:comparator[mode=compare,powered=false,facing=south]",
		0, nil,
	},
	{
		149, 9, "minecraft:comparator[mode=compare,powered=false,facing=west]",
		0, nil,
	},
	{
		149, 10, "minecraft:comparator[mode=compare,powered=false,facing=north]",
		0, nil,
	},
	{
		149, 11, "minecraft:comparator[mode=compare,powered=false,facing=east]",
		0, nil,
	},
	{
		149, 12, "minecraft:comparator[mode=subtract,powered=false,facing=south]",
		0, nil,
	},
	{
		149, 13, "minecraft:comparator[mode=subtract,powered=false,facing=west]",
		0, nil,
	},
	{
		149, 14, "minecraft:comparator[mode=subtract,powered=false,facing=north]",
		0, nil,
	},
	{
		149, 15, "minecraft:comparator[mode=subtract,powered=false,facing=east]",
		0, nil,
	},
	{
		150, 0, "minecraft:comparator[mode=compare,powered=true,facing=south]",
		0, nil,
	},
	{
		150, 1, "minecraft:comparator[mode=compare,powered=true,facing=west]",
		0, nil,
	},
	{
		150, 2, "minecraft:comparator[mode=compare,powered=true,facing=north]",
		0, nil,
	},
	{
		150, 3, "minecraft:comparator[mode=compare,powered=true,facing=east]",
		0, nil,
	},
	{
		150, 4, "minecraft:comparator[mode=subtract,powered=true,facing=south]",
		0, nil,
	},
	{
		150, 5, "minecraft:comparator[mode=subtract,powered=true,facing=west]",
		0, nil,
	},
	{
		150, 6, "minecraft:comparator[mode=subtract,powered=true,facing=north]",
		0, nil,
	},
	{
		150, 7, "minecraft:comparator[mode=subtract,powered=true,facing=east]",
		0, nil,
	},
	{
		150, 8, "minecraft:comparator[mode=compare,powered=true,facing=south]",
		0, nil,
	},
	{
		150, 9, "minecraft:comparator[mode=compare,powered=true,facing=west]",
		0, nil,
	},
	{
		150, 10, "minecraft:comparator[mode=compare,powered=true,facing=north]",
		0, nil,
	},
	{
		150, 11, "minecraft:comparator[mode=compare,powered=true,facing=east]",
		0, nil,
	},
	{
		150, 12, "minecraft:comparator[mode=subtract,powered=true,facing=south]",
		0, nil,
	},
	{
		150, 13, "minecraft:comparator[mode=subtract,powered=true,facing=west]",
		0, nil,
	},
	{
		150, 14, "minecraft:comparator[mode=subtract,powered=true,facing=north]",
		0, nil,
	},
	{
		150, 15, "minecraft:comparator[mode=subtract,powered=true,facing=east]",
		0, nil,
	},
	{
		151, 0, "minecraft:daylight_detector[inverted=false,power=0]",
		0, nil,
	},
	{
		151, 1, "minecraft:daylight_detector[inverted=false,power=1]",
		0, nil,
	},
	{
		151, 2, "minecraft:daylight_detector[inverted=false,power=2]",
		0, nil,
	},
	{
		151, 3, "minecraft:daylight_detector[inverted=false,power=3]",
		0, nil,
	},
	{
		151, 4, "minecraft:daylight_detector[inverted=false,power=4]",
		0, nil,
	},
	{
		151, 5, "minecraft:daylight_detector[inverted=false,power=5]",
		0, nil,
	},
	{
		151, 6, "minecraft:daylight_detector[inverted=false,power=6]",
		0, nil,
	},
	{
		151, 7, "minecraft:daylight_detector[inverted=false,power=7]",
		0, nil,
	},
	{
		151, 8, "minecraft:daylight_detector[inverted=false,power=8]",
		0, nil,
	},
	{
		151, 9, "minecraft:daylight_detector[inverted=false,power=9]",
		0, nil,
	},
	{
		151, 10, "minecraft:daylight_detector[inverted=false,power=10]",
		0, nil,
	},
	{
		151, 11, "minecraft:daylight_detector[inverted=false,power=11]",
		0, nil,
	},
	{
		151, 12, "minecraft:daylight_detector[inverted=false,power=12]",
		0, nil,
	},
	{
		151, 13, "minecraft:daylight_detector[inverted=false,power=13]",
		0, nil,
	},
	{
		151, 14, "minecraft:daylight_detector[inverted=false,power=14]",
		0, nil,
	},
	{
		151, 15, "minecraft:daylight_detector[inverted=false,power=15]",
		0, nil,
	},
	{
		152, 0, "minecraft:redstone_block",
		0, nil,
	},
	{
		153, 0, "minecraft:nether_quartz_ore",
		0, nil,
	},
	{
		154, 0, "minecraft:hopper[facing=down,enabled=true]",
		0, nil,
	},
	{
		154, 2, "minecraft:hopper[facing=north,enabled=true]",
		0, nil,
	},
	{
		154, 3, "minecraft:hopper[facing=south,enabled=true]",
		0, nil,
	},
	{
		154, 4, "minecraft:hopper[facing=west,enabled=true]",
		0, nil,
	},
	{
		154, 5, "minecraft:hopper[facing=east,enabled=true]",
		0, nil,
	},
	{
		154, 8, "minecraft:hopper[facing=down,enabled=false]",
		0, nil,
	},
	{
		154, 10, "minecraft:hopper[facing=north,enabled=false]",
		0, nil,
	},
	{
		154, 11, "minecraft:hopper[facing=south,enabled=false]",
		0, nil,
	},
	{
		154, 12, "minecraft:hopper[facing=west,enabled=false]",
		0, nil,
	},
	{
		154, 13, "minecraft:hopper[facing=east,enabled=false]",
		0, nil,
	},
	{
		155, 0, "minecraft:quartz_block",
		0, nil,
	},
	{
		155, 1, "minecraft:chiseled_quartz_block",
		0, nil,
	},
	{
		155, 2, "minecraft:quartz_pillar[axis=y]",
		0, nil,
	},
	{
		155, 3, "minecraft:quartz_pillar[axis=x]",
		0, nil,
	},
	{
		155, 4, "minecraft:quartz_pillar[axis=z]",
		0, nil,
	},
	{
		155, 6, "minecraft:quartz_pillar[axis=x]",
		0, nil,
	},
	{
		155, 10, "minecraft:quartz_pillar[axis=z]",
		0, nil,
	},
	{
		156, 0, "minecraft:quartz_stairs[half=bottom,shape=straight,facing=east]",
		0, nil,
	},
	{
		156, 1, "minecraft:quartz_stairs[half=bottom,shape=straight,facing=west]",
		0, nil,
	},
	{
		156, 2, "minecraft:quartz_stairs[half=bottom,shape=straight,facing=south]",
		0, nil,
	},
	{
		156, 3, "minecraft:quartz_stairs[half=bottom,shape=straight,facing=north]",
		0, nil,
	},
	{
		156, 4, "minecraft:quartz_stairs[half=top,shape=straight,facing=east]",
		0, nil,
	},
	{
		156, 5, "minecraft:quartz_stairs[half=top,shape=straight,facing=west]",
		0, nil,
	},
	{
		156, 6, "minecraft:quartz_stairs[half=top,shape=straight,facing=south]",
		0, nil,
	},
	{
		156, 7, "minecraft:quartz_stairs[half=top,shape=straight,facing=north]",
		0, nil,
	},
	{
		157, 0, "minecraft:activator_rail[shape=north_south,powered=false]",
		0, nil,
	},
	{
		157, 1, "minecraft:activator_rail[shape=east_west,powered=false]",
		0, nil,
	},
	{
		157, 2, "minecraft:activator_rail[shape=ascending_east,powered=false]",
		0, nil,
	},
	{
		157, 3, "minecraft:activator_rail[shape=ascending_west,powered=false]",
		0, nil,
	},
	{
		157, 4, "minecraft:activator_rail[shape=ascending_north,powered=false]",
		0, nil,
	},
	{
		157, 5, "minecraft:activator_rail[shape=ascending_south,powered=false]",
		0, nil,
	},
	{
		157, 8, "minecraft:activator_rail[shape=north_south,powered=true]",
		0, nil,
	},
	{
		157, 9, "minecraft:activator_rail[shape=east_west,powered=true]",
		0, nil,
	},
	{
		157, 10, "minecraft:activator_rail[shape=ascending_east,powered=true]",
		0, nil,
	},
	{
		157, 11, "minecraft:activator_rail[shape=ascending_west,powered=true]",
		0, nil,
	},
	{
		157, 12, "minecraft:activator_rail[shape=ascending_north,powered=true]",
		0, nil,
	},
	{
		157, 13, "minecraft:activator_rail[shape=ascending_south,powered=true]",
		0, nil,
	},
	{
		158, 0, "minecraft:dropper[triggered=false,facing=down]",
		0, nil,
	},
	{
		158, 1, "minecraft:dropper[triggered=false,facing=up]",
		0, nil,
	},
	{
		158, 2, "minecraft:dropper[triggered=false,facing=north]",
		0, nil,
	},
	{
		158, 3, "minecraft:dropper[triggered=false,facing=south]",
		0, nil,
	},
	{
		158, 4, "minecraft:dropper[triggered=false,facing=west]",
		0, nil,
	},
	{
		158, 5, "minecraft:dropper[triggered=false,facing=east]",
		0, nil,
	},
	{
		158, 8, "minecraft:dropper[triggered=true,facing=down]",
		0, nil,
	},
	{
		158, 9, "minecraft:dropper[triggered=true,facing=up]",
		0, nil,
	},
	{
		158, 10, "minecraft:dropper[triggered=true,facing=north]",
		0, nil,
	},
	{
		158, 11, "minecraft:dropper[triggered=true,facing=south]",
		0, nil,
	},
	{
		158, 12, "minecraft:dropper[triggered=true,facing=west]",
		0, nil,
	},
	{
		158, 13, "minecraft:dropper[triggered=true,facing=east]",
		0, nil,
	},
	{
		159, 0, "minecraft:white_terracotta",
		0, nil,
	},
	{
		159, 1, "minecraft:orange_terracotta",
		0, nil,
	},
	{
		159, 2, "minecraft:magenta_terracotta",
		0, nil,
	},
	{
		159, 3, "minecraft:light_blue_terracotta",
		0, nil,
	},
	{
		159, 4, "minecraft:yellow_terracotta",
		0, nil,
	},
	{
		159, 5, "minecraft:lime_terracotta",
		0, nil,
	},
	{
		159, 6, "minecraft:pink_terracotta",
		0, nil,
	},
	{
		159, 7, "minecraft:gray_terracotta",
		0, nil,
	},
	{
		159, 8, "minecraft:light_gray_terracotta",
		0, nil,
	},
	{
		159, 9, "minecraft:cyan_terracotta",
		0, nil,
	},
	{
		159, 10, "minecraft:purple_terracotta",
		0, nil,
	},
	{
		159, 11, "minecraft:blue_terracotta",
		0, nil,
	},
	{
		159, 12, "minecraft:brown_terracotta",
		0, nil,
	},
	{
		159, 13, "minecraft:green_terracotta",
		0, nil,
	},
	{
		159, 14, "minecraft:red_terracotta",
		0, nil,
	},
	{
		159, 15, "minecraft:black_terracotta",
		0, nil,
	},
	{
		160, 0, "minecraft:white_stained_glass_pane[east=false,south=false,north=false,west=false]",
		0, nil,
	},
	{
		160, 1, "minecraft:orange_stained_glass_pane[east=false,south=false,north=false,west=false]",
		0, nil,
	},
	{
		160, 2, "minecraft:magenta_stained_glass_pane[east=false,south=false,north=false,west=false]",
		0, nil,
	},
	{
		160, 3, "minecraft:light_blue_stained_glass_pane[east=false,south=false,north=false,west=false]",
		0, nil,
	},
	{
		160, 4, "minecraft:yellow_stained_glass_pane[east=false,south=false,north=false,west=false]",
		0, nil,
	},
	{
		160, 5, "minecraft:lime_stained_glass_pane[east=false,south=false,north=false,west=false]",
		0, nil,
	},
	{
		160, 6, "minecraft:pink_stained_glass_pane[east=false,south=false,north=false,west=false]",
		0, nil,
	},
	{
		160, 7, "minecraft:gray_stained_glass_pane[east=false,south=false,north=false,west=false]",
		0, nil,
	},
	{
		160, 8, "minecraft:light_gray_stained_glass_pane[east=false,south=false,north=false,west=false]",
		0, nil,
	},
	{
		160, 9, "minecraft:cyan_stained_glass_pane[east=false,south=false,north=false,west=false]",
		0, nil,
	},
	{
		160, 10, "minecraft:purple_stained_glass_pane[east=false,south=false,north=false,west=false]",
		0, nil,
	},
	{
		160, 11, "minecraft:blue_stained_glass_pane[east=false,south=false,north=false,west=false]",
		0, nil,
	},
	{
		160, 12, "minecraft:brown_stained_glass_pane[east=false,south=false,north=false,west=false]",
		0, nil,
	},
	{
		160, 13, "minecraft:green_stained_glass_pane[east=false,south=false,north=false,west=false]",
		0, nil,
	},
	{
		160, 14, "minecraft:red_stained_glass_pane[east=false,south=false,north=false,west=false]",
		0, nil,
	},
	{
		160, 15, "minecraft:black_stained_glass_pane[east=false,south=false,north=false,west=false]",
		0, nil,
	},
	{
		161, 0, "minecraft:acacia_leaves[persistent=false,distance=1]",
		0, nil,
	},
	{
		161, 1, "minecraft:dark_oak_leaves[persistent=false,distance=1]",
		0, nil,
	},
	{
		161, 4, "minecraft:acacia_leaves[persistent=true,distance=1]",
		0, nil,
	},
	{
		161, 5, "minecraft:dark_oak_leaves[persistent=true,distance=1]",
		0, nil,
	},
	{
		161, 8, "minecraft:acacia_leaves[persistent=false,distance=1]",
		0, nil,
	},
	{
		161, 9, "minecraft:dark_oak_leaves[persistent=false,distance=1]",
		0, nil,
	},
	{
		161, 12, "minecraft:acacia_leaves[persistent=true,distance=1]",
		0, nil,
	},
	{
		161, 13, "minecraft:dark_oak_leaves[persistent=true,distance=1]",
		0, nil,
	},
	{
		162, 0, "minecraft:acacia_log[axis=y]",
		0, nil,
	},
	{
		162, 1, "minecraft:dark_oak_log[axis=y]",
		0, nil,
	},
	{
		162, 4, "minecraft:acacia_log[axis=x]",
		0, nil,
	},
	{
		162, 5, "minecraft:dark_oak_log[axis=x]",
		0, nil,
	},
	{
		162, 8, "minecraft:acacia_log[axis=z]",
		0, nil,
	},
	{
		162, 9, "minecraft:dark_oak_log[axis=z]",
		0, nil,
	},
	{
		162, 12, "minecraft:acacia_wood",
		0, nil,
	},
	{
		162, 13, "minecraft:dark_oak_wood",
		0, nil,
	},
	{
		163, 0, "minecraft:acacia_stairs[half=bottom,shape=straight,facing=east]",
		0, nil,
	},
	{
		163, 1, "minecraft:acacia_stairs[half=bottom,shape=straight,facing=west]",
		0, nil,
	},
	{
		163, 2, "minecraft:acacia_stairs[half=bottom,shape=straight,facing=south]",
		0, nil,
	},
	{
		163, 3, "minecraft:acacia_stairs[half=bottom,shape=straight,facing=north]",
		0, nil,
	},
	{
		163, 4, "minecraft:acacia_stairs[half=top,shape=straight,facing=east]",
		0, nil,
	},
	{
		163, 5, "minecraft:acacia_stairs[half=top,shape=straight,facing=west]",
		0, nil,
	},
	{
		163, 6, "minecraft:acacia_stairs[half=top,shape=straight,facing=south]",
		0, nil,
	},
	{
		163, 7, "minecraft:acacia_stairs[half=top,shape=straight,facing=north]",
		0, nil,
	},
	{
		164, 0, "minecraft:dark_oak_stairs[half=bottom,shape=straight,facing=east]",
		0, nil,
	},
	{
		164, 1, "minecraft:dark_oak_stairs[half=bottom,shape=straight,facing=west]",
		0, nil,
	},
	{
		164, 2, "minecraft:dark_oak_stairs[half=bottom,shape=straight,facing=south]",
		0, nil,
	},
	{
		164, 3, "minecraft:dark_oak_stairs[half=bottom,shape=straight,facing=north]",
		0, nil,
	},
	{
		164, 4, "minecraft:dark_oak_stairs[half=top,shape=straight,facing=east]",
		0, nil,
	},
	{
		164, 5, "minecraft:dark_oak_stairs[half=top,shape=straight,facing=west]",
		0, nil,
	},
	{
		164, 6, "minecraft:dark_oak_stairs[half=top,shape=straight,facing=south]",
		0, nil,
	},
	{
		164, 7, "minecraft:dark_oak_stairs[half=top,shape=straight,facing=north]",
		0, nil,
	},
	{
		165, 0, "minecraft:slime_block",
		0, nil,
	},
	{
		166, 0, "minecraft:barrier",
		0, nil,
	},
	{
		167, 0, "minecraft:iron_trapdoor[half=bottom,facing=north,open=false]",
		0, nil,
	},
	{
		167, 1, "minecraft:iron_trapdoor[half=bottom,facing=south,open=false]",
		0, nil,
	},
	{
		167, 2, "minecraft:iron_trapdoor[half=bottom,facing=west,open=false]",
		0, nil,
	},
	{
		167, 3, "minecraft:iron_trapdoor[half=bottom,facing=east,open=false]",
		0, nil,
	},
	{
		167, 4, "minecraft:iron_trapdoor[half=bottom,facing=north,open=true]",
		0, nil,
	},
	{
		167, 5, "minecraft:iron_trapdoor[half=bottom,facing=south,open=true]",
		0, nil,
	},
	{
		167, 6, "minecraft:iron_trapdoor[half=bottom,facing=west,open=true]",
		0, nil,
	},
	{
		167, 7, "minecraft:iron_trapdoor[half=bottom,facing=east,open=true]",
		0, nil,
	},
	{
		167, 8, "minecraft:iron_trapdoor[half=top,facing=north,open=false]",
		0, nil,
	},
	{
		167, 9, "minecraft:iron_trapdoor[half=top,facing=south,open=false]",
		0, nil,
	},
	{
		167, 10, "minecraft:iron_trapdoor[half=top,facing=west,open=false]",
		0, nil,
	},
	{
		167, 11, "minecraft:iron_trapdoor[half=top,facing=east,open=false]",
		0, nil,
	},
	{
		167, 12, "minecraft:iron_trapdoor[half=top,facing=north,open=true]",
		0, nil,
	},
	{
		167, 13, "minecraft:iron_trapdoor[half=top,facing=south,open=true]",
		0, nil,
	},
	{
		167, 14, "minecraft:iron_trapdoor[half=top,facing=west,open=true]",
		0, nil,
	},
	{
		167, 15, "minecraft:iron_trapdoor[half=top,facing=east,open=true]",
		0, nil,
	},
	{
		168, 0, "minecraft:prismarine",
		0, nil,
	},
	{
		168, 1, "minecraft:prismarine_bricks",
		0, nil,
	},
	{
		168, 2, "minecraft:dark_prismarine",
		0, nil,
	},
	{
		169, 0, "minecraft:sea_lantern",
		0, nil,
	},
	{
		170, 0, "minecraft:hay_block[axis=y]",
		0, nil,
	},
	{
		170, 4, "minecraft:hay_block[axis=x]",
		0, nil,
	},
	{
		170, 8, "minecraft:hay_block[axis=z]",
		0, nil,
	},
	{
		171, 0, "minecraft:white_carpet",
		0, nil,
	},
	{
		171, 1, "minecraft:orange_carpet",
		0, nil,
	},
	{
		171, 2, "minecraft:magenta_carpet",
		0, nil,
	},
	{
		171, 3, "minecraft:light_blue_carpet",
		0, nil,
	},
	{
		171, 4, "minecraft:yellow_carpet",
		0, nil,
	},
	{
		171, 5, "minecraft:lime_carpet",
		0, nil,
	},
	{
		171, 6, "minecraft:pink_carpet",
		0, nil,
	},
	{
		171, 7, "minecraft:gray_carpet",
		0, nil,
	},
	{
		171, 8, "minecraft:light_gray_carpet",
		0, nil,
	},
	{
		171, 9, "minecraft:cyan_carpet",
		0, nil,
	},
	{
		171, 10, "minecraft:purple_carpet",
		0, nil,
	},
	{
		171, 11, "minecraft:blue_carpet",
		0, nil,
	},
	{
		171, 12, "minecraft:brown_carpet",
		0, nil,
	},
	{
		171, 13, "minecraft:green_carpet",
		0, nil,
	},
	{
		171, 14, "minecraft:red_carpet",
		0, nil,
	},
	{
		171, 15, "minecraft:black_carpet",
		0, nil,
	},
	{
		172, 0, "minecraft:terracotta",
		0, nil,
	},
	{
		173, 0, "minecraft:coal_block",
		0, nil,
	},
	{
		174, 0, "minecraft:packed_ice",
		0, nil,
	},
	{
		175, 0, "minecraft:sunflower[half=lower]",
		0, nil,
	},
	{
		175, 1, "minecraft:lilac[half=lower]",
		0, nil,
	},
	{
		175, 2, "minecraft:tall_grass[half=lower]",
		0, nil,
	},
	{
		175, 3, "minecraft:large_fern[half=lower]",
		0, nil,
	},
	{
		175, 4, "minecraft:rose_bush[half=lower]",
		0, nil,
	},
	{
		175, 5, "minecraft:peony[half=lower]",
		0, nil,
	},
	{
		175, 8, "minecraft:sunflower[half=upper]",
		0, nil,
	},
	{
		175, 9, "minecraft:lilac[half=upper]",
		0, nil,
	},
	{
		175, 10, "minecraft:tall_grass[half=upper]",
		0, nil,
	},
	{
		175, 11, "minecraft:large_fern[half=upper]",
		0, nil,
	},
	{
		175, 12, "minecraft:rose_bush[half=upper]",
		0, nil,
	},
	{
		175, 13, "minecraft:peony[half=upper]",
		0, nil,
	},
	{
		176, 0, "minecraft:white_banner[rotation=0]",
		0, nil,
	},
	{
		176, 1, "minecraft:white_banner[rotation=1]",
		0, nil,
	},
	{
		176, 2, "minecraft:white_banner[rotation=2]",
		0, nil,
	},
	{
		176, 3, "minecraft:white_banner[rotation=3]",
		0, nil,
	},
	{
		176, 4, "minecraft:white_banner[rotation=4]",
		0, nil,
	},
	{
		176, 5, "minecraft:white_banner[rotation=5]",
		0, nil,
	},
	{
		176, 6, "minecraft:white_banner[rotation=6]",
		0, nil,
	},
	{
		176, 7, "minecraft:white_banner[rotation=7]",
		0, nil,
	},
	{
		176, 8, "minecraft:white_banner[rotation=8]",
		0, nil,
	},
	{
		176, 9, "minecraft:white_banner[rotation=9]",
		0, nil,
	},
	{
		176, 10, "minecraft:white_banner[rotation=10]",
		0, nil,
	},
	{
		176, 11, "minecraft:white_banner[rotation=11]",
		0, nil,
	},
	{
		176, 12, "minecraft:white_banner[rotation=12]",
		0, nil,
	},
	{
		176, 13, "minecraft:white_banner[rotation=13]",
		0, nil,
	},
	{
		176, 14, "minecraft:white_banner[rotation=14]",
		0, nil,
	},
	{
		176, 15, "minecraft:white_banner[rotation=15]",
		0, nil,
	},
	{
		177, 2, "minecraft:white_wall_banner[facing=north]",
		0, nil,
	},
	{
		177, 3, "minecraft:white_wall_banner[facing=south]",
		0, nil,
	},
	{
		177, 4, "minecraft:white_wall_banner[facing=west]",
		0, nil,
	},
	{
		177, 5, "minecraft:white_wall_banner[facing=east]",
		0, nil,
	},
	{
		178, 0, "minecraft:daylight_detector[inverted=true,power=0]",
		0, nil,
	},
	{
		178, 1, "minecraft:daylight_detector[inverted=true,power=1]",
		0, nil,
	},
	{
		178, 2, "minecraft:daylight_detector[inverted=true,power=2]",
		0, nil,
	},
	{
		178, 3, "minecraft:daylight_detector[inverted=true,power=3]",
		0, nil,
	},
	{
		178, 4, "minecraft:daylight_detector[inverted=true,power=4]",
		0, nil,
	},
	{
		178, 5, "minecraft:daylight_detector[inverted=true,power=5]",
		0, nil,
	},
	{
		178, 6, "minecraft:daylight_detector[inverted=true,power=6]",
		0, nil,
	},
	{
		178, 7, "minecraft:daylight_detector[inverted=true,power=7]",
		0, nil,
	},
	{
		178, 8, "minecraft:daylight_detector[inverted=true,power=8]",
		0, nil,
	},
	{
		178, 9, "minecraft:daylight_detector[inverted=true,power=9]",
		0, nil,
	},
	{
		178, 10, "minecraft:daylight_detector[inverted=true,power=10]",
		0, nil,
	},
	{
		178, 11, "minecraft:daylight_detector[inverted=true,power=11]",
		0, nil,
	},
	{
		178, 12, "minecraft:daylight_detector[inverted=true,power=12]",
		0, nil,
	},
	{
		178, 13, "minecraft:daylight_detector[inverted=true,power=13]",
		0, nil,
	},
	{
		178, 14, "minecraft:daylight_detector[inverted=true,power=14]",
		0, nil,
	},
	{
		178, 15, "minecraft:daylight_detector[inverted=true,power=15]",
		0, nil,
	},
	{
		179, 0, "minecraft:red_sandstone",
		0, nil,
	},
	{
		179, 1, "minecraft:chiseled_red_sandstone",
		0, nil,
	},
	{
		179, 2, "minecraft:cut_red_sandstone",
		0, nil,
	},
	{
		180, 0, "minecraft:red_sandstone_stairs[half=bottom,shape=straight,facing=east]",
		0, nil,
	},
	{
		180, 1, "minecraft:red_sandstone_stairs[half=bottom,shape=straight,facing=west]",
		0, nil,
	},
	{
		180, 2, "minecraft:red_sandstone_stairs[half=bottom,shape=straight,facing=south]",
		0, nil,
	},
	{
		180, 3, "minecraft:red_sandstone_stairs[half=bottom,shape=straight,facing=north]",
		0, nil,
	},
	{
		180, 4, "minecraft:red_sandstone_stairs[half=top,shape=straight,facing=east]",
		0, nil,
	},
	{
		180, 5, "minecraft:red_sandstone_stairs[half=top,shape=straight,facing=west]",
		0, nil,
	},
	{
		180, 6, "minecraft:red_sandstone_stairs[half=top,shape=straight,facing=south]",
		0, nil,
	},
	{
		180, 7, "minecraft:red_sandstone_stairs[half=top,shape=straight,facing=north]",
		0, nil,
	},
	{
		181, 0, "minecraft:red_sandstone_slab[type=double]",
		0, nil,
	},
	{
		181, 8, "minecraft:smooth_red_sandstone",
		0, nil,
	},
	{
		182, 0, "minecraft:red_sandstone_slab[type=bottom]",
		0, nil,
	},
	{
		182, 8, "minecraft:red_sandstone_slab[type=top]",
		0, nil,
	},
	{
		183, 0, "minecraft:spruce_fence_gate[in_wall=false,powered=false,facing=south,open=false]",
		0, nil,
	},
	{
		183, 1, "minecraft:spruce_fence_gate[in_wall=false,powered=false,facing=west,open=false]",
		0, nil,
	},
	{
		183, 2, "minecraft:spruce_fence_gate[in_wall=false,powered=false,facing=north,open=false]",
		0, nil,
	},
	{
		183, 3, "minecraft:spruce_fence_gate[in_wall=false,powered=false,facing=east,open=false]",
		0, nil,
	},
	{
		183, 4, "minecraft:spruce_fence_gate[in_wall=false,powered=false,facing=south,open=true]",
		0, nil,
	},
	{
		183, 5, "minecraft:spruce_fence_gate[in_wall=false,powered=false,facing=west,open=true]",
		0, nil,
	},
	{
		183, 6, "minecraft:spruce_fence_gate[in_wall=false,powered=false,facing=north,open=true]",
		0, nil,
	},
	{
		183, 7, "minecraft:spruce_fence_gate[in_wall=false,powered=false,facing=east,open=true]",
		0, nil,
	},
	{
		183, 8, "minecraft:spruce_fence_gate[in_wall=false,powered=true,facing=south,open=false]",
		0, nil,
	},
	{
		183, 9, "minecraft:spruce_fence_gate[in_wall=false,powered=true,facing=west,open=false]",
		0, nil,
	},
	{
		183, 10, "minecraft:spruce_fence_gate[in_wall=false,powered=true,facing=north,open=false]",
		0, nil,
	},
	{
		183, 11, "minecraft:spruce_fence_gate[in_wall=false,powered=true,facing=east,open=false]",
		0, nil,
	},
	{
		183, 12, "minecraft:spruce_fence_gate[in_wall=false,powered=true,facing=south,open=true]",
		0, nil,
	},
	{
		183, 13, "minecraft:spruce_fence_gate[in_wall=false,powered=true,facing=west,open=true]",
		0, nil,
	},
	{
		183, 14, "minecraft:spruce_fence_gate[in_wall=false,powered=true,facing=north,open=true]",
		0, nil,
	},
	{
		183, 15, "minecraft:spruce_fence_gate[in_wall=false,powered=true,facing=east,open=true]",
		0, nil,
	},
	{
		184, 0, "minecraft:birch_fence_gate[in_wall=false,powered=false,facing=south,open=false]",
		0, nil,
	},
	{
		184, 1, "minecraft:birch_fence_gate[in_wall=false,powered=false,facing=west,open=false]",
		0, nil,
	},
	{
		184, 2, "minecraft:birch_fence_gate[in_wall=false,powered=false,facing=north,open=false]",
		0, nil,
	},
	{
		184, 3, "minecraft:birch_fence_gate[in_wall=false,powered=false,facing=east,open=false]",
		0, nil,
	},
	{
		184, 4, "minecraft:birch_fence_gate[in_wall=false,powered=false,facing=south,open=true]",
		0, nil,
	},
	{
		184, 5, "minecraft:birch_fence_gate[in_wall=false,powered=false,facing=west,open=true]",
		0, nil,
	},
	{
		184, 6, "minecraft:birch_fence_gate[in_wall=false,powered=false,facing=north,open=true]",
		0, nil,
	},
	{
		184, 7, "minecraft:birch_fence_gate[in_wall=false,powered=false,facing=east,open=true]",
		0, nil,
	},
	{
		184, 8, "minecraft:birch_fence_gate[in_wall=false,powered=true,facing=south,open=false]",
		0, nil,
	},
	{
		184, 9, "minecraft:birch_fence_gate[in_wall=false,powered=true,facing=west,open=false]",
		0, nil,
	},
	{
		184, 10, "minecraft:birch_fence_gate[in_wall=false,powered=true,facing=north,open=false]",
		0, nil,
	},
	{
		184, 11, "minecraft:birch_fence_gate[in_wall=false,powered=true,facing=east,open=false]",
		0, nil,
	},
	{
		184, 12, "minecraft:birch_fence_gate[in_wall=false,powered=true,facing=south,open=true]",
		0, nil,
	},
	{
		184, 13, "minecraft:birch_fence_gate[in_wall=false,powered=true,facing=west,open=true]",
		0, nil,
	},
	{
		184, 14, "minecraft:birch_fence_gate[in_wall=false,powered=true,facing=north,open=true]",
		0, nil,
	},
	{
		184, 15, "minecraft:birch_fence_gate[in_wall=false,powered=true,facing=east,open=true]",
		0, nil,
	},
	{
		185, 0, "minecraft:jungle_fence_gate[in_wall=false,powered=false,facing=south,open=false]",
		0, nil,
	},
	{
		185, 1, "minecraft:jungle_fence_gate[in_wall=false,powered=false,facing=west,open=false]",
		0, nil,
	},
	{
		185, 2, "minecraft:jungle_fence_gate[in_wall=false,powered=false,facing=north,open=false]",
		0, nil,
	},
	{
		185, 3, "minecraft:jungle_fence_gate[in_wall=false,powered=false,facing=east,open=false]",
		0, nil,
	},
	{
		185, 4, "minecraft:jungle_fence_gate[in_wall=false,powered=false,facing=south,open=true]",
		0, nil,
	},
	{
		185, 5, "minecraft:jungle_fence_gate[in_wall=false,powered=false,facing=west,open=true]",
		0, nil,
	},
	{
		185, 6, "minecraft:jungle_fence_gate[in_wall=false,powered=false,facing=north,open=true]",
		0, nil,
	},
	{
		185, 7, "minecraft:jungle_fence_gate[in_wall=false,powered=false,facing=east,open=true]",
		0, nil,
	},
	{
		185, 8, "minecraft:jungle_fence_gate[in_wall=false,powered=true,facing=south,open=false]",
		0, nil,
	},
	{
		185, 9, "minecraft:jungle_fence_gate[in_wall=false,powered=true,facing=west,open=false]",
		0, nil,
	},
	{
		185, 10, "minecraft:jungle_fence_gate[in_wall=false,powered=true,facing=north,open=false]",
		0, nil,
	},
	{
		185, 11, "minecraft:jungle_fence_gate[in_wall=false,powered=true,facing=east,open=false]",
		0, nil,
	},
	{
		185, 12, "minecraft:jungle_fence_gate[in_wall=false,powered=true,facing=south,open=true]",
		0, nil,
	},
	{
		185, 13, "minecraft:jungle_fence_gate[in_wall=false,powered=true,facing=west,open=true]",
		0, nil,
	},
	{
		185, 14, "minecraft:jungle_fence_gate[in_wall=false,powered=true,facing=north,open=true]",
		0, nil,
	},
	{
		185, 15, "minecraft:jungle_fence_gate[in_wall=false,powered=true,facing=east,open=true]",
		0, nil,
	},
	{
		186, 0, "minecraft:dark_oak_fence_gate[in_wall=false,powered=false,facing=south,open=false]",
		0, nil,
	},
	{
		186, 1, "minecraft:dark_oak_fence_gate[in_wall=false,powered=false,facing=west,open=false]",
		0, nil,
	},
	{
		186, 2, "minecraft:dark_oak_fence_gate[in_wall=false,powered=false,facing=north,open=false]",
		0, nil,
	},
	{
		186, 3, "minecraft:dark_oak_fence_gate[in_wall=false,powered=false,facing=east,open=false]",
		0, nil,
	},
	{
		186, 4, "minecraft:dark_oak_fence_gate[in_wall=false,powered=false,facing=south,open=true]",
		0, nil,
	},
	{
		186, 5, "minecraft:dark_oak_fence_gate[in_wall=false,powered=false,facing=west,open=true]",
		0, nil,
	},
	{
		186, 6, "minecraft:dark_oak_fence_gate[in_wall=false,powered=false,facing=north,open=true]",
		0, nil,
	},
	{
		186, 7, "minecraft:dark_oak_fence_gate[in_wall=false,powered=false,facing=east,open=true]",
		0, nil,
	},
	{
		186, 8, "minecraft:dark_oak_fence_gate[in_wall=false,powered=true,facing=south,open=false]",
		0, nil,
	},
	{
		186, 9, "minecraft:dark_oak_fence_gate[in_wall=false,powered=true,facing=west,open=false]",
		0, nil,
	},
	{
		186, 10, "minecraft:dark_oak_fence_gate[in_wall=false,powered=true,facing=north,open=false]",
		0, nil,
	},
	{
		186, 11, "minecraft:dark_oak_fence_gate[in_wall=false,powered=true,facing=east,open=false]",
		0, nil,
	},
	{
		186, 12, "minecraft:dark_oak_fence_gate[in_wall=false,powered=true,facing=south,open=true]",
		0, nil,
	},
	{
		186, 13, "minecraft:dark_oak_fence_gate[in_wall=false,powered=true,facing=west,open=true]",
		0, nil,
	},
	{
		186, 14, "minecraft:dark_oak_fence_gate[in_wall=false,powered=true,facing=north,open=true]",
		0, nil,
	},
	{
		186, 15, "minecraft:dark_oak_fence_gate[in_wall=false,powered=true,facing=east,open=true]",
		0, nil,
	},
	{
		187, 0, "minecraft:acacia_fence_gate[in_wall=false,powered=false,facing=south,open=false]",
		0, nil,
	},
	{
		187, 1, "minecraft:acacia_fence_gate[in_wall=false,powered=false,facing=west,open=false]",
		0, nil,
	},
	{
		187, 2, "minecraft:acacia_fence_gate[in_wall=false,powered=false,facing=north,open=false]",
		0, nil,
	},
	{
		187, 3, "minecraft:acacia_fence_gate[in_wall=false,powered=false,facing=east,open=false]",
		0, nil,
	},
	{
		187, 4, "minecraft:acacia_fence_gate[in_wall=false,powered=false,facing=south,open=true]",
		0, nil,
	},
	{
		187, 5, "minecraft:acacia_fence_gate[in_wall=false,powered=false,facing=west,open=true]",
		0, nil,
	},
	{
		187, 6, "minecraft:acacia_fence_gate[in_wall=false,powered=false,facing=north,open=true]",
		0, nil,
	},
	{
		187, 7, "minecraft:acacia_fence_gate[in_wall=false,powered=false,facing=east,open=true]",
		0, nil,
	},
	{
		187, 8, "minecraft:acacia_fence_gate[in_wall=false,powered=true,facing=south,open=false]",
		0, nil,
	},
	{
		187, 9, "minecraft:acacia_fence_gate[in_wall=false,powered=true,facing=west,open=false]",
		0, nil,
	},
	{
		187, 10, "minecraft:acacia_fence_gate[in_wall=false,powered=true,facing=north,open=false]",
		0, nil,
	},
	{
		187, 11, "minecraft:acacia_fence_gate[in_wall=false,powered=true,facing=east,open=false]",
		0, nil,
	},
	{
		187, 12, "minecraft:acacia_fence_gate[in_wall=false,powered=true,facing=south,open=true]",
		0, nil,
	},
	{
		187, 13, "minecraft:acacia_fence_gate[in_wall=false,powered=true,facing=west,open=true]",
		0, nil,
	},
	{
		187, 14, "minecraft:acacia_fence_gate[in_wall=false,powered=true,facing=north,open=true]",
		0, nil,
	},
	{
		187, 15, "minecraft:acacia_fence_gate[in_wall=false,powered=true,facing=east,open=true]",
		0, nil,
	},
	{
		188, 0, "minecraft:spruce_fence[east=false,south=false,north=false,west=false]",
		0, nil,
	},
	{
		189, 0, "minecraft:birch_fence[east=false,south=false,north=false,west=false]",
		0, nil,
	},
	{
		190, 0, "minecraft:jungle_fence[east=false,south=false,north=false,west=false]",
		0, nil,
	},
	{
		191, 0, "minecraft:dark_oak_fence[east=false,south=false,north=false,west=false]",
		0, nil,
	},
	{
		192, 0, "minecraft:acacia_fence[east=false,south=false,north=false,west=false]",
		0, nil,
	},
	{
		193, 0, "minecraft:spruce_door[hinge=right,half=lower,powered=false,facing=east,open=false]",
		0, nil,
	},
	{
		193, 1, "minecraft:spruce_door[hinge=right,half=lower,powered=false,facing=south,open=false]",
		0, nil,
	},
	{
		193, 2, "minecraft:spruce_door[hinge=right,half=lower,powered=false,facing=west,open=false]",
		0, nil,
	},
	{
		193, 3, "minecraft:spruce_door[hinge=right,half=lower,powered=false,facing=north,open=false]",
		0, nil,
	},
	{
		193, 4, "minecraft:spruce_door[hinge=right,half=lower,powered=false,facing=east,open=true]",
		0, nil,
	},
	{
		193, 5, "minecraft:spruce_door[hinge=right,half=lower,powered=false,facing=south,open=true]",
		0, nil,
	},
	{
		193, 6, "minecraft:spruce_door[hinge=right,half=lower,powered=false,facing=west,open=true]",
		0, nil,
	},
	{
		193, 7, "minecraft:spruce_door[hinge=right,half=lower,powered=false,facing=north,open=true]",
		0, nil,
	},
	{
		193, 8, "minecraft:spruce_door[hinge=left,half=upper,powered=false,facing=east,open=false]",
		0, nil,
	},
	{
		193, 9, "minecraft:spruce_door[hinge=right,half=upper,powered=false,facing=east,open=false]",
		0, nil,
	},
	{
		193, 10, "minecraft:spruce_door[hinge=left,half=upper,powered=true,facing=east,open=false]",
		0, nil,
	},
	{
		193, 11, "minecraft:spruce_door[hinge=right,half=upper,powered=true,facing=east,open=false]",
		0, nil,
	},
	{
		194, 0, "minecraft:birch_door[hinge=right,half=lower,powered=false,facing=east,open=false]",
		0, nil,
	},
	{
		194, 1, "minecraft:birch_door[hinge=right,half=lower,powered=false,facing=south,open=false]",
		0, nil,
	},
	{
		194, 2, "minecraft:birch_door[hinge=right,half=lower,powered=false,facing=west,open=false]",
		0, nil,
	},
	{
		194, 3, "minecraft:birch_door[hinge=right,half=lower,powered=false,facing=north,open=false]",
		0, nil,
	},
	{
		194, 4, "minecraft:birch_door[hinge=right,half=lower,powered=false,facing=east,open=true]",
		0, nil,
	},
	{
		194, 5, "minecraft:birch_door[hinge=right,half=lower,powered=false,facing=south,open=true]",
		0, nil,
	},
	{
		194, 6, "minecraft:birch_door[hinge=right,half=lower,powered=false,facing=west,open=true]",
		0, nil,
	},
	{
		194, 7, "minecraft:birch_door[hinge=right,half=lower,powered=false,facing=north,open=true]",
		0, nil,
	},
	{
		194, 8, "minecraft:birch_door[hinge=left,half=upper,powered=false,facing=east,open=false]",
		0, nil,
	},
	{
		194, 9, "minecraft:birch_door[hinge=right,half=upper,powered=false,facing=east,open=false]",
		0, nil,
	},
	{
		194, 10, "minecraft:birch_door[hinge=left,half=upper,powered=true,facing=east,open=false]",
		0, nil,
	},
	{
		194, 11, "minecraft:birch_door[hinge=right,half=upper,powered=true,facing=east,open=false]",
		0, nil,
	},
	{
		195, 0, "minecraft:jungle_door[hinge=right,half=lower,powered=false,facing=east,open=false]",
		0, nil,
	},
	{
		195, 1, "minecraft:jungle_door[hinge=right,half=lower,powered=false,facing=south,open=false]",
		0, nil,
	},
	{
		195, 2, "minecraft:jungle_door[hinge=right,half=lower,powered=false,facing=west,open=false]",
		0, nil,
	},
	{
		195, 3, "minecraft:jungle_door[hinge=right,half=lower,powered=false,facing=north,open=false]",
		0, nil,
	},
	{
		195, 4, "minecraft:jungle_door[hinge=right,half=lower,powered=false,facing=east,open=true]",
		0, nil,
	},
	{
		195, 5, "minecraft:jungle_door[hinge=right,half=lower,powered=false,facing=south,open=true]",
		0, nil,
	},
	{
		195, 6, "minecraft:jungle_door[hinge=right,half=lower,powered=false,facing=west,open=true]",
		0, nil,
	},
	{
		195, 7, "minecraft:jungle_door[hinge=right,half=lower,powered=false,facing=north,open=true]",
		0, nil,
	},
	{
		195, 8, "minecraft:jungle_door[hinge=left,half=upper,powered=false,facing=east,open=false]",
		0, nil,
	},
	{
		195, 9, "minecraft:jungle_door[hinge=right,half=upper,powered=false,facing=east,open=false]",
		0, nil,
	},
	{
		195, 10, "minecraft:jungle_door[hinge=left,half=upper,powered=true,facing=east,open=false]",
		0, nil,
	},
	{
		195, 11, "minecraft:jungle_door[hinge=right,half=upper,powered=true,facing=east,open=false]",
		0, nil,
	},
	{
		196, 0, "minecraft:acacia_door[hinge=right,half=lower,powered=false,facing=east,open=false]",
		0, nil,
	},
	{
		196, 1, "minecraft:acacia_door[hinge=right,half=lower,powered=false,facing=south,open=false]",
		0, nil,
	},
	{
		196, 2, "minecraft:acacia_door[hinge=right,half=lower,powered=false,facing=west,open=false]",
		0, nil,
	},
	{
		196, 3, "minecraft:acacia_door[hinge=right,half=lower,powered=false,facing=north,open=false]",
		0, nil,
	},
	{
		196, 4, "minecraft:acacia_door[hinge=right,half=lower,powered=false,facing=east,open=true]",
		0, nil,
	},
	{
		196, 5, "minecraft:acacia_door[hinge=right,half=lower,powered=false,facing=south,open=true]",
		0, nil,
	},
	{
		196, 6, "minecraft:acacia_door[hinge=right,half=lower,powered=false,facing=west,open=true]",
		0, nil,
	},
	{
		196, 7, "minecraft:acacia_door[hinge=right,half=lower,powered=false,facing=north,open=true]",
		0, nil,
	},
	{
		196, 8, "minecraft:acacia_door[hinge=left,half=upper,powered=false,facing=east,open=false]",
		0, nil,
	},
	{
		196, 9, "minecraft:acacia_door[hinge=right,half=upper,powered=false,facing=east,open=false]",
		0, nil,
	},
	{
		196, 10, "minecraft:acacia_door[hinge=left,half=upper,powered=true,facing=east,open=false]",
		0, nil,
	},
	{
		196, 11, "minecraft:acacia_door[hinge=right,half=upper,powered=true,facing=east,open=false]",
		0, nil,
	},
	{
		197, 0, "minecraft:dark_oak_door[hinge=right,half=lower,powered=false,facing=east,open=false]",
		0, nil,
	},
	{
		197, 1, "minecraft:dark_oak_door[hinge=right,half=lower,powered=false,facing=south,open=false]",
		0, nil,
	},
	{
		197, 2, "minecraft:dark_oak_door[hinge=right,half=lower,powered=false,facing=west,open=false]",
		0, nil,
	},
	{
		197, 3, "minecraft:dark_oak_door[hinge=right,half=lower,powered=false,facing=north,open=false]",
		0, nil,
	},
	{
		197, 4, "minecraft:dark_oak_door[hinge=right,half=lower,powered=false,facing=east,open=true]",
		0, nil,
	},
	{
		197, 5, "minecraft:dark_oak_door[hinge=right,half=lower,powered=false,facing=south,open=true]",
		0, nil,
	},
	{
		197, 6, "minecraft:dark_oak_door[hinge=right,half=lower,powered=false,facing=west,open=true]",
		0, nil,
	},
	{
		197, 7, "minecraft:dark_oak_door[hinge=right,half=lower,powered=false,facing=north,open=true]",
		0, nil,
	},
	{
		197, 8, "minecraft:dark_oak_door[hinge=left,half=upper,powered=false,facing=east,open=false]",
		0, nil,
	},
	{
		197, 9, "minecraft:dark_oak_door[hinge=right,half=upper,powered=false,facing=east,open=false]",
		0, nil,
	},
	{
		197, 10, "minecraft:dark_oak_door[hinge=left,half=upper,powered=true,facing=east,open=false]",
		0, nil,
	},
	{
		197, 11, "minecraft:dark_oak_door[hinge=right,half=upper,powered=true,facing=east,open=false]",
		0, nil,
	},
	{
		198, 0, "minecraft:end_rod[facing=down]",
		0, nil,
	},
	{
		198, 1, "minecraft:end_rod[facing=up]",
		0, nil,
	},
	{
		198, 2, "minecraft:end_rod[facing=north]",
		0, nil,
	},
	{
		198, 3, "minecraft:end_rod[facing=south]",
		0, nil,
	},
	{
		198, 4, "minecraft:end_rod[facing=west]",
		0, nil,
	},
	{
		198, 5, "minecraft:end_rod[facing=east]",
		0, nil,
	},
	{
		199, 0, "minecraft:chorus_plant[east=false,south=false,north=false,west=false,up=false,down=false]",
		0, nil,
	},
	{
		200, 0, "minecraft:chorus_flower[age=0]",
		0, nil,
	},
	{
		200, 1, "minecraft:chorus_flower[age=1]",
		0, nil,
	},
	{
		200, 2, "minecraft:chorus_flower[age=2]",
		0, nil,
	},
	{
		200, 3, "minecraft:chorus_flower[age=3]",
		0, nil,
	},
	{
		200, 4, "minecraft:chorus_flower[age=4]",
		0, nil,
	},
	{
		200, 5, "minecraft:chorus_flower[age=5]",
		0, nil,
	},
	{
		201, 0, "minecraft:purpur_block",
		107, fallback("minecraft:quartz_block"),
	},
	{
		202, 0, "minecraft:purpur_pillar[axis=y]",
		107, fallback("minecraft:quartz_pillar[axis=y]"),
	},
	{
		202, 4, "minecraft:purpur_pillar[axis=x]",
		107, fallback("minecraft:quartz_pillar[axis=x]"),
	},
	{
		202, 8, "minecraft:purpur_pillar[axis=z]",
		107, fallback("minecraft:quartz_pillar[axis=z]"),
	},
	{
		203, 0, "minecraft:purpur_stairs[half=bottom,shape=straight,facing=east]",
		107, fallback("minecraft:quartz_stairs[half=bottom,shape=straight,facing=east]"),
	},
	{
		203, 1, "minecraft:purpur_stairs[half=bottom,shape=straight,facing=west]",
		107, fallback("minecraft:quartz_stairs[half=bottom,shape=straight,facing=west]"),
	},
	{
		203, 2, "minecraft:purpur_stairs[half=bottom,shape=straight,facing=south]",
		107, fallback("minecraft:quartz_stairs[half=bottom,shape=straight,facing=south]"),
	},
	{
		203, 3, "minecraft:purpur_stairs[half=bottom,shape=straight,facing=north]",
		107, fallback("minecraft:quartz_stairs[half=bottom,shape=straight,facing=north]"),
	},
	{
		203, 4, "minecraft:purpur_stairs[half=top,shape=straight,facing=east]",
		107, fallback("minecraft:quartz_stairs[half=top,shape=straight,facing=east]"),
	},
	{
		203, 5, "minecraft:purpur_stairs[half=top,shape=straight,facing=west]",
		107, fallback("minecraft:quartz"),
	},
	{
		203, 6, "minecraft:purpur_stairs[half=top,shape=straight,facing=south]",
		107, fallback("minecraft:quartz_stairs[half=top,shape=straight,facing=west]"),
	},
	{
		203, 7, "minecraft:purpur_stairs[half=top,shape=straight,facing=north]",
		107, fallback("minecraft:quartz_stairs[half=top,shape=straight,facing=north]"),
	},
	{
		204, 0, "minecraft:purpur_slab[type=double]",
		107, fallback("minecraft:quartz_slab[type=double]"),
	},
	{
		205, 0, "minecraft:purpur_slab[type=bottom]",
		107, fallback("minecraft:quartz_slab[type=bottom]"),
	},
	{
		205, 8, "minecraft:purpur_slab[type=top]",
		107, fallback("minecraft:quartz_slab[type=top]"),
	},
	{
		206, 0, "minecraft:end_stone_bricks",
		0, nil,
	},
	{
		207, 0, "minecraft:beetroots[age=0]",
		0, nil,
	},
	{
		207, 1, "minecraft:beetroots[age=1]",
		0, nil,
	},
	{
		207, 2, "minecraft:beetroots[age=2]",
		0, nil,
	},
	{
		207, 3, "minecraft:beetroots[age=3]",
		0, nil,
	},
	{
		208, 0, "minecraft:grass_path",
		107, fallback("minecraft:grass"),
	},
	{
		209, 0, "minecraft:end_gateway",
		0, nil,
	},
	{
		210, 0, "minecraft:repeating_command_block[conditional=false,facing=down]",
		0, nil,
	},
	{
		210, 1, "minecraft:repeating_command_block[conditional=false,facing=up]",
		0, nil,
	},
	{
		210, 2, "minecraft:repeating_command_block[conditional=false,facing=north]",
		0, nil,
	},
	{
		210, 3, "minecraft:repeating_command_block[conditional=false,facing=south]",
		0, nil,
	},
	{
		210, 4, "minecraft:repeating_command_block[conditional=false,facing=west]",
		0, nil,
	},
	{
		210, 5, "minecraft:repeating_command_block[conditional=false,facing=east]",
		0, nil,
	},
	{
		210, 8, "minecraft:repeating_command_block[conditional=true,facing=down]",
		0, nil,
	},
	{
		210, 9, "minecraft:repeating_command_block[conditional=true,facing=up]",
		0, nil,
	},
	{
		210, 10, "minecraft:repeating_command_block[conditional=true,facing=north]",
		0, nil,
	},
	{
		210, 11, "minecraft:repeating_command_block[conditional=true,facing=south]",
		0, nil,
	},
	{
		210, 12, "minecraft:repeating_command_block[conditional=true,facing=west]",
		0, nil,
	},
	{
		210, 13, "minecraft:repeating_command_block[conditional=true,facing=east]",
		0, nil,
	},
	{
		211, 0, "minecraft:chain_command_block[conditional=false,facing=down]",
		0, nil,
	},
	{
		211, 1, "minecraft:chain_command_block[conditional=false,facing=up]",
		0, nil,
	},
	{
		211, 2, "minecraft:chain_command_block[conditional=false,facing=north]",
		0, nil,
	},
	{
		211, 3, "minecraft:chain_command_block[conditional=false,facing=south]",
		0, nil,
	},
	{
		211, 4, "minecraft:chain_command_block[conditional=false,facing=west]",
		0, nil,
	},
	{
		211, 5, "minecraft:chain_command_block[conditional=false,facing=east]",
		0, nil,
	},
	{
		211, 8, "minecraft:chain_command_block[conditional=true,facing=down]",
		0, nil,
	},
	{
		211, 9, "minecraft:chain_command_block[conditional=true,facing=up]",
		0, nil,
	},
	{
		211, 10, "minecraft:chain_command_block[conditional=true,facing=north]",
		0, nil,
	},
	{
		211, 11, "minecraft:chain_command_block[conditional=true,facing=south]",
		0, nil,
	},
	{
		211, 12, "minecraft:chain_command_block[conditional=true,facing=west]",
		0, nil,
	},
	{
		211, 13, "minecraft:chain_command_block[conditional=true,facing=east]",
		0, nil,
	},
	{
		212, 0, "minecraft:frosted_ice[age=0]",
		107, fallback("minecraft:ice"),
	},
	{
		212, 1, "minecraft:frosted_ice[age=1]",
		107, fallback("minecraft:ice"),
	},
	{
		212, 2, "minecraft:frosted_ice[age=2]",
		107, fallback("minecraft:packed_ice"),
	},
	{
		212, 3, "minecraft:frosted_ice[age=3]",
		107, fallback("minecraft:packed_ice"),
	},
	{
		213, 0, "minecraft:magma_block",
		0, nil,
	},
	{
		214, 0, "minecraft:nether_wart_block",
		0, nil,
	},
	{
		215, 0, "minecraft:red_nether_bricks",
		0, nil,
	},
	{
		216, 0, "minecraft:bone_block[axis=y]",
		0, nil,
	},
	{
		216, 4, "minecraft:bone_block[axis=x]",
		0, nil,
	},
	{
		216, 8, "minecraft:bone_block[axis=z]",
		0, nil,
	},
	{
		217, 0, "minecraft:structure_void",
		0, nil,
	},
	{
		218, 0, "minecraft:observer[powered=false,facing=down]",
		0, nil,
	},
	{
		218, 1, "minecraft:observer[powered=false,facing=up]",
		0, nil,
	},
	{
		218, 2, "minecraft:observer[powered=false,facing=north]",
		0, nil,
	},
	{
		218, 3, "minecraft:observer[powered=false,facing=south]",
		0, nil,
	},
	{
		218, 4, "minecraft:observer[powered=false,facing=west]",
		0, nil,
	},
	{
		218, 5, "minecraft:observer[powered=false,facing=east]",
		0, nil,
	},
	{
		218, 8, "minecraft:observer[powered=true,facing=down]",
		0, nil,
	},
	{
		218, 9, "minecraft:observer[powered=true,facing=up]",
		0, nil,
	},
	{
		218, 10, "minecraft:observer[powered=true,facing=north]",
		0, nil,
	},
	{
		218, 11, "minecraft:observer[powered=true,facing=south]",
		0, nil,
	},
	{
		218, 12, "minecraft:observer[powered=true,facing=west]",
		0, nil,
	},
	{
		218, 13, "minecraft:observer[powered=true,facing=east]",
		0, nil,
	},
	{
		219, 0, "minecraft:white_shulker_box[facing=down]",
		0, nil,
	},
	{
		219, 1, "minecraft:white_shulker_box[facing=up]",
		0, nil,
	},
	{
		219, 2, "minecraft:white_shulker_box[facing=north]",
		0, nil,
	},
	{
		219, 3, "minecraft:white_shulker_box[facing=south]",
		0, nil,
	},
	{
		219, 4, "minecraft:white_shulker_box[facing=west]",
		0, nil,
	},
	{
		219, 5, "minecraft:white_shulker_box[facing=east]",
		0, nil,
	},
	{
		220, 0, "minecraft:orange_shulker_box[facing=down]",
		0, nil,
	},
	{
		220, 1, "minecraft:orange_shulker_box[facing=up]",
		0, nil,
	},
	{
		220, 2, "minecraft:orange_shulker_box[facing=north]",
		0, nil,
	},
	{
		220, 3, "minecraft:orange_shulker_box[facing=south]",
		0, nil,
	},
	{
		220, 4, "minecraft:orange_shulker_box[facing=west]",
		0, nil,
	},
	{
		220, 5, "minecraft:orange_shulker_box[facing=east]",
		0, nil,
	},
	{
		221, 0, "minecraft:magenta_shulker_box[facing=down]",
		0, nil,
	},
	{
		221, 1, "minecraft:magenta_shulker_box[facing=up]",
		0, nil,
	},
	{
		221, 2, "minecraft:magenta_shulker_box[facing=north]",
		0, nil,
	},
	{
		221, 3, "minecraft:magenta_shulker_box[facing=south]",
		0, nil,
	},
	{
		221, 4, "minecraft:magenta_shulker_box[facing=west]",
		0, nil,
	},
	{
		221, 5, "minecraft:magenta_shulker_box[facing=east]",
		0, nil,
	},
	{
		222, 0, "minecraft:light_blue_shulker_box[facing=down]",
		0, nil,
	},
	{
		222, 1, "minecraft:light_blue_shulker_box[facing=up]",
		0, nil,
	},
	{
		222, 2, "minecraft:light_blue_shulker_box[facing=north]",
		0, nil,
	},
	{
		222, 3, "minecraft:light_blue_shulker_box[facing=south]",
		0, nil,
	},
	{
		222, 4, "minecraft:light_blue_shulker_box[facing=west]",
		0, nil,
	},
	{
		222, 5, "minecraft:light_blue_shulker_box[facing=east]",
		0, nil,
	},
	{
		223, 0, "minecraft:yellow_shulker_box[facing=down]",
		0, nil,
	},
	{
		223, 1, "minecraft:yellow_shulker_box[facing=up]",
		0, nil,
	},
	{
		223, 2, "minecraft:yellow_shulker_box[facing=north]",
		0, nil,
	},
	{
		223, 3, "minecraft:yellow_shulker_box[facing=south]",
		0, nil,
	},
	{
		223, 4, "minecraft:yellow_shulker_box[facing=west]",
		0, nil,
	},
	{
		223, 5, "minecraft:yellow_shulker_box[facing=east]",
		0, nil,
	},
	{
		224, 0, "minecraft:lime_shulker_box[facing=down]",
		0, nil,
	},
	{
		224, 1, "minecraft:lime_shulker_box[facing=up]",
		0, nil,
	},
	{
		224, 2, "minecraft:lime_shulker_box[facing=north]",
		0, nil,
	},
	{
		224, 3, "minecraft:lime_shulker_box[facing=south]",
		0, nil,
	},
	{
		224, 4, "minecraft:lime_shulker_box[facing=west]",
		0, nil,
	},
	{
		224, 5, "minecraft:lime_shulker_box[facing=east]",
		0, nil,
	},
	{
		225, 0, "minecraft:pink_shulker_box[facing=down]",
		0, nil,
	},
	{
		225, 1, "minecraft:pink_shulker_box[facing=up]",
		0, nil,
	},
	{
		225, 2, "minecraft:pink_shulker_box[facing=north]",
		0, nil,
	},
	{
		225, 3, "minecraft:pink_shulker_box[facing=south]",
		0, nil,
	},
	{
		225, 4, "minecraft:pink_shulker_box[facing=west]",
		0, nil,
	},
	{
		225, 5, "minecraft:pink_shulker_box[facing=east]",
		0, nil,
	},
	{
		226, 0, "minecraft:gray_shulker_box[facing=down]",
		0, nil,
	},
	{
		226, 1, "minecraft:gray_shulker_box[facing=up]",
		0, nil,
	},
	{
		226, 2, "minecraft:gray_shulker_box[facing=north]",
		0, nil,
	},
	{
		226, 3, "minecraft:gray_shulker_box[facing=south]",
		0, nil,
	},
	{
		226, 4, "minecraft:gray_shulker_box[facing=west]",
		0, nil,
	},
	{
		226, 5, "minecraft:gray_shulker_box[facing=east]",
		0, nil,
	},
	{
		227, 0, "minecraft:light_gray_shulker_box[facing=down]",
		0, nil,
	},
	{
		227, 1, "minecraft:light_gray_shulker_box[facing=up]",
		0, nil,
	},
	{
		227, 2, "minecraft:light_gray_shulker_box[facing=north]",
		0, nil,
	},
	{
		227, 3, "minecraft:light_gray_shulker_box[facing=south]",
		0, nil,
	},
	{
		227, 4, "minecraft:light_gray_shulker_box[facing=west]",
		0, nil,
	},
	{
		227, 5, "minecraft:light_gray_shulker_box[facing=east]",
		0, nil,
	},
	{
		228, 0, "minecraft:cyan_shulker_box[facing=down]",
		0, nil,
	},
	{
		228, 1, "minecraft:cyan_shulker_box[facing=up]",
		0, nil,
	},
	{
		228, 2, "minecraft:cyan_shulker_box[facing=north]",
		0, nil,
	},
	{
		228, 3, "minecraft:cyan_shulker_box[facing=south]",
		0, nil,
	},
	{
		228, 4, "minecraft:cyan_shulker_box[facing=west]",
		0, nil,
	},
	{
		228, 5, "minecraft:cyan_shulker_box[facing=east]",
		0, nil,
	},
	{
		229, 0, "minecraft:purple_shulker_box[facing=down]",
		0, nil,
	},
	{
		229, 1, "minecraft:purple_shulker_box[facing=up]",
		0, nil,
	},
	{
		229, 2, "minecraft:purple_shulker_box[facing=north]",
		0, nil,
	},
	{
		229, 3, "minecraft:purple_shulker_box[facing=south]",
		0, nil,
	},
	{
		229, 4, "minecraft:purple_shulker_box[facing=west]",
		0, nil,
	},
	{
		229, 5, "minecraft:purple_shulker_box[facing=east]",
		0, nil,
	},
	{
		230, 0, "minecraft:blue_shulker_box[facing=down]",
		0, nil,
	},
	{
		230, 1, "minecraft:blue_shulker_box[facing=up]",
		0, nil,
	},
	{
		230, 2, "minecraft:blue_shulker_box[facing=north]",
		0, nil,
	},
	{
		230, 3, "minecraft:blue_shulker_box[facing=south]",
		0, nil,
	},
	{
		230, 4, "minecraft:blue_shulker_box[facing=west]",
		0, nil,
	},
	{
		230, 5, "minecraft:blue_shulker_box[facing=east]",
		0, nil,
	},
	{
		231, 0, "minecraft:brown_shulker_box[facing=down]",
		0, nil,
	},
	{
		231, 1, "minecraft:brown_shulker_box[facing=up]",
		0, nil,
	},
	{
		231, 2, "minecraft:brown_shulker_box[facing=north]",
		0, nil,
	},
	{
		231, 3, "minecraft:brown_shulker_box[facing=south]",
		0, nil,
	},
	{
		231, 4, "minecraft:brown_shulker_box[facing=west]",
		0, nil,
	},
	{
		231, 5, "minecraft:brown_shulker_box[facing=east]",
		0, nil,
	},
	{
		232, 0, "minecraft:green_shulker_box[facing=down]",
		0, nil,
	},
	{
		232, 1, "minecraft:green_shulker_box[facing=up]",
		0, nil,
	},
	{
		232, 2, "minecraft:green_shulker_box[facing=north]",
		0, nil,
	},
	{
		232, 3, "minecraft:green_shulker_box[facing=south]",
		0, nil,
	},
	{
		232, 4, "minecraft:green_shulker_box[facing=west]",
		0, nil,
	},
	{
		232, 5, "minecraft:green_shulker_box[facing=east]",
		0, nil,
	},
	{
		233, 0, "minecraft:red_shulker_box[facing=down]",
		0, nil,
	},
	{
		233, 1, "minecraft:red_shulker_box[facing=up]",
		0, nil,
	},
	{
		233, 2, "minecraft:red_shulker_box[facing=north]",
		0, nil,
	},
	{
		233, 3, "minecraft:red_shulker_box[facing=south]",
		0, nil,
	},
	{
		233, 4, "minecraft:red_shulker_box[facing=west]",
		0, nil,
	},
	{
		233, 5, "minecraft:red_shulker_box[facing=east]",
		0, nil,
	},
	{
		234, 0, "minecraft:black_shulker_box[facing=down]",
		0, nil,
	},
	{
		234, 1, "minecraft:black_shulker_box[facing=up]",
		0, nil,
	},
	{
		234, 2, "minecraft:black_shulker_box[facing=north]",
		0, nil,
	},
	{
		234, 3, "minecraft:black_shulker_box[facing=south]",
		0, nil,
	},
	{
		234, 4, "minecraft:black_shulker_box[facing=west]",
		0, nil,
	},
	{
		234, 5, "minecraft:black_shulker_box[facing=east]",
		0, nil,
	},
	{
		235, 0, "minecraft:white_glazed_terracotta[facing=south]",
		0, nil,
	},
	{
		235, 1, "minecraft:white_glazed_terracotta[facing=west]",
		0, nil,
	},
	{
		235, 2, "minecraft:white_glazed_terracotta[facing=north]",
		0, nil,
	},
	{
		235, 3, "minecraft:white_glazed_terracotta[facing=east]",
		0, nil,
	},
	{
		236, 0, "minecraft:orange_glazed_terracotta[facing=south]",
		0, nil,
	},
	{
		236, 1, "minecraft:orange_glazed_terracotta[facing=west]",
		0, nil,
	},
	{
		236, 2, "minecraft:orange_glazed_terracotta[facing=north]",
		0, nil,
	},
	{
		236, 3, "minecraft:orange_glazed_terracotta[facing=east]",
		0, nil,
	},
	{
		237, 0, "minecraft:magenta_glazed_terracotta[facing=south]",
		0, nil,
	},
	{
		237, 1, "minecraft:magenta_glazed_terracotta[facing=west]",
		0, nil,
	},
	{
		237, 2, "minecraft:magenta_glazed_terracotta[facing=north]",
		0, nil,
	},
	{
		237, 3, "minecraft:magenta_glazed_terracotta[facing=east]",
		0, nil,
	},
	{
		238, 0, "minecraft:light_blue_glazed_terracotta[facing=south]",
		0, nil,
	},
	{
		238, 1, "minecraft:light_blue_glazed_terracotta[facing=west]",
		0, nil,
	},
	{
		238, 2, "minecraft:light_blue_glazed_terracotta[facing=north]",
		0, nil,
	},
	{
		238, 3, "minecraft:light_blue_glazed_terracotta[facing=east]",
		0, nil,
	},
	{
		239, 0, "minecraft:yellow_glazed_terracotta[facing=south]",
		0, nil,
	},
	{
		239, 1, "minecraft:yellow_glazed_terracotta[facing=west]",
		0, nil,
	},
	{
		239, 2, "minecraft:yellow_glazed_terracotta[facing=north]",
		0, nil,
	},
	{
		239, 3, "minecraft:yellow_glazed_terracotta[facing=east]",
		0, nil,
	},
	{
		240, 0, "minecraft:lime_glazed_terracotta[facing=south]",
		0, nil,
	},
	{
		240, 1, "minecraft:lime_glazed_terracotta[facing=west]",
		0, nil,
	},
	{
		240, 2, "minecraft:lime_glazed_terracotta[facing=north]",
		0, nil,
	},
	{
		240, 3, "minecraft:lime_glazed_terracotta[facing=east]",
		0, nil,
	},
	{
		241, 0, "minecraft:pink_glazed_terracotta[facing=south]",
		0, nil,
	},
	{
		241, 1, "minecraft:pink_glazed_terracotta[facing=west]",
		0, nil,
	},
	{
		241, 2, "minecraft:pink_glazed_terracotta[facing=north]",
		0, nil,
	},
	{
		241, 3, "minecraft:pink_glazed_terracotta[facing=east]",
		0, nil,
	},
	{
		242, 0, "minecraft:gray_glazed_terracotta[facing=south]",
		0, nil,
	},
	{
		242, 1, "minecraft:gray_glazed_terracotta[facing=west]",
		0, nil,
	},
	{
		242, 2, "minecraft:gray_glazed_terracotta[facing=north]",
		0, nil,
	},
	{
		242, 3, "minecraft:gray_glazed_terracotta[facing=east]",
		0, nil,
	},
	{
		243, 0, "minecraft:light_gray_glazed_terracotta[facing=south]",
		0, nil,
	},
	{
		243, 1, "minecraft:light_gray_glazed_terracotta[facing=west]",
		0, nil,
	},
	{
		243, 2, "minecraft:light_gray_glazed_terracotta[facing=north]",
		0, nil,
	},
	{
		243, 3, "minecraft:light_gray_glazed_terracotta[facing=east]",
		0, nil,
	},
	{
		244, 0, "minecraft:cyan_glazed_terracotta[facing=south]",
		0, nil,
	},
	{
		244, 1, "minecraft:cyan_glazed_terracotta[facing=west]",
		0, nil,
	},
	{
		244, 2, "minecraft:cyan_glazed_terracotta[facing=north]",
		0, nil,
	},
	{
		244, 3, "minecraft:cyan_glazed_terracotta[facing=east]",
		0, nil,
	},
	{
		245, 0, "minecraft:purple_glazed_terracotta[facing=south]",
		0, nil,
	},
	{
		245, 1, "minecraft:purple_glazed_terracotta[facing=west]",
		0, nil,
	},
	{
		245, 2, "minecraft:purple_glazed_terracotta[facing=north]",
		0, nil,
	},
	{
		245, 3, "minecraft:purple_glazed_terracotta[facing=east]",
		0, nil,
	},
	{
		246, 0, "minecraft:blue_glazed_terracotta[facing=south]",
		0, nil,
	},
	{
		246, 1, "minecraft:blue_glazed_terracotta[facing=west]",
		0, nil,
	},
	{
		246, 2, "minecraft:blue_glazed_terracotta[facing=north]",
		0, nil,
	},
	{
		246, 3, "minecraft:blue_glazed_terracotta[facing=east]",
		0, nil,
	},
	{
		247, 0, "minecraft:brown_glazed_terracotta[facing=south]",
		0, nil,
	},
	{
		247, 1, "minecraft:brown_glazed_terracotta[facing=west]",
		0, nil,
	},
	{
		247, 2, "minecraft:brown_glazed_terracotta[facing=north]",
		0, nil,
	},
	{
		247, 3, "minecraft:brown_glazed_terracotta[facing=east]",
		0, nil,
	},
	{
		248, 0, "minecraft:green_glazed_terracotta[facing=south]",
		0, nil,
	},
	{
		248, 1, "minecraft:green_glazed_terracotta[facing=west]",
		0, nil,
	},
	{
		248, 2, "minecraft:green_glazed_terracotta[facing=north]",
		0, nil,
	},
	{
		248, 3, "minecraft:green_glazed_terracotta[facing=east]",
		0, nil,
	},
	{
		249, 0, "minecraft:red_glazed_terracotta[facing=south]",
		0, nil,
	},
	{
		249, 1, "minecraft:red_glazed_terracotta[facing=west]",
		0, nil,
	},
	{
		249, 2, "minecraft:red_glazed_terracotta[facing=north]",
		0, nil,
	},
	{
		249, 3, "minecraft:red_glazed_terracotta[facing=east]",
		0, nil,
	},
	{
		250, 0, "minecraft:black_glazed_terracotta[facing=south]",
		0, nil,
	},
	{
		250, 1, "minecraft:black_glazed_terracotta[facing=west]",
		0, nil,
	},
	{
		250, 2, "minecraft:black_glazed_terracotta[facing=north]",
		0, nil,
	},
	{
		250, 3, "minecraft:black_glazed_terracotta[facing=east]",
		0, nil,
	},
	{
		251, 0, "minecraft:white_concrete",
		0, nil,
	},
	{
		251, 1, "minecraft:orange_concrete",
		0, nil,
	},
	{
		251, 2, "minecraft:magenta_concrete",
		0, nil,
	},
	{
		251, 3, "minecraft:light_blue_concrete",
		0, nil,
	},
	{
		251, 4, "minecraft:yellow_concrete",
		0, nil,
	},
	{
		251, 5, "minecraft:lime_concrete",
		0, nil,
	},
	{
		251, 6, "minecraft:pink_concrete",
		0, nil,
	},
	{
		251, 7, "minecraft:gray_concrete",
		0, nil,
	},
	{
		251, 8, "minecraft:light_gray_concrete",
		0, nil,
	},
	{
		251, 9, "minecraft:cyan_concrete",
		0, nil,
	},
	{
		251, 10, "minecraft:purple_concrete",
		0, nil,
	},
	{
		251, 11, "minecraft:blue_concrete",
		0, nil,
	},
	{
		251, 12, "minecraft:brown_concrete",
		0, nil,
	},
	{
		251, 13, "minecraft:green_concrete",
		0, nil,
	},
	{
		251, 14, "minecraft:red_concrete",
		0, nil,
	},
	{
		251, 15, "minecraft:black_concrete",
		0, nil,
	},
	{
		252, 0, "minecraft:white_concrete_powder",
		0, nil,
	},
	{
		252, 1, "minecraft:orange_concrete_powder",
		0, nil,
	},
	{
		252, 2, "minecraft:magenta_concrete_powder",
		0, nil,
	},
	{
		252, 3, "minecraft:light_blue_concrete_powder",
		0, nil,
	},
	{
		252, 4, "minecraft:yellow_concrete_powder",
		0, nil,
	},
	{
		252, 5, "minecraft:lime_concrete_powder",
		0, nil,
	},
	{
		252, 6, "minecraft:pink_concrete_powder",
		0, nil,
	},
	{
		252, 7, "minecraft:gray_concrete_powder",
		0, nil,
	},
	{
		252, 8, "minecraft:light_gray_concrete_powder",
		0, nil,
	},
	{
		252, 9, "minecraft:cyan_concrete_powder",
		0, nil,
	},
	{
		252, 10, "minecraft:purple_concrete_powder",
		0, nil,
	},
	{
		252, 11, "minecraft:blue_concrete_powder",
		0, nil,
	},
	{
		252, 12, "minecraft:brown_concrete_powder",
		0, nil,
	},
	{
		252, 13, "minecraft:green_concrete_powder",
		0, nil,
	},
	{
		252, 14, "minecraft:red_concrete_powder",
		0, nil,
	},
	{
		252, 15, "minecraft:black_concrete_powder",
		0, nil,
	},
	{
		255, 0, "minecraft:structure_block[mode=save]",
		0, nil,
	},
	{
		255, 1, "minecraft:structure_block[mode=load]",
		0, nil,
	},
	{
		255, 2, "minecraft:structure_block[mode=corner]",
		0, nil,
	},
	{
		255, 3, "minecraft:structure_block[mode=data]",
		0, nil,
	},
}

var legacyFromInt, legacyFromString = generateLegacyBlockMapping()

func generateLegacyBlockMapping() (map[int]*LegacyBlock, map[string]*LegacyBlock) {
	legacyFromInt := make(map[int]*LegacyBlock)
	legacyFromString := make(map[string]*LegacyBlock)
	for i, v := range legacy {
		bigId := v.GetBlockState()
		legacyFromInt[bigId] = &legacy[i]
		legacyFromString[v.Name] = &legacy[i]
	}

	return legacyFromInt, legacyFromString
}

func fallback(name string) *string {
	return &name
}

func GetLegacyFromState(legacy int) *LegacyBlock {
	if val, ok := legacyFromInt[legacy]; ok {
		return val
	} else {
		return legacyFromString["minecraft:air"]
	}
}

func GetLegacyFromName(name string) *LegacyBlock {
	if val, ok := legacyFromString[name]; ok {
		return val
	} else {
		return legacyFromString["minecraft:air"]
	}
}

func GetLegacyMapping() map[int]*LegacyBlock {
	return legacyFromInt
}

func (block *LegacyBlock) GetBlockState() int {
	return block.Id << 4 | (block.Data & 0xF)
}

func GetLegacyBlockState(id int, data int) int {
	return id << 4 | (data & 0xF)
}

func GetTypeDataFromLegacy(legacy int) (int, int) {
	return legacy >> 4, legacy & 0xF
}
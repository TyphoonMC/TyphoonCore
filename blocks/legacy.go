package blocks

func GetLegacyBlockState(id int, data int) int {
	return id << 4 | (data & 0xF)
}

func GetTypeDataFromLegacy(legacy int) (int, int) {
	return legacy >> 4, legacy & 0xF
}
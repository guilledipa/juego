package assets

func AssetBytes(name string) []byte {
	a := map[string][]byte{
		"bg_green":          bgGreenBytes,
		"bg_wood":           bgWoodBytes,
		"curtains":          curtainBytes,
		"curtains_straight": curtainStraightBytes,
	}
	return a[name]
}

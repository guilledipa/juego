package assets

func AssetBytes(name string) []byte {
	a := map[string][]byte{
		"bg_green":         bgGreenBytes,
		"bg_wood":          bgWoodBytes,
		"curtains":         curtainBytes,
		"curtain_straight": curtainStraightBytes,
	}
	return a[name]
}

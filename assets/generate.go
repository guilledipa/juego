//go:generate file2byteslice -input ./PNG/Stall/bg_green.png -output bg_green.go -package assets -var bgGreenBytes
//go:generate file2byteslice -input ./PNG/Stall/curtain.png -output curtains.go -package assets -var curtainBytes
//go:generate file2byteslice -input ./PNG/Stall/curtain_straight.png -output curtains_straight.go -package assets -var curtainStraightBytes
//go:generate file2byteslice -input ./PNG/Stall/bg_wood.png -output bg_wood.go -package assets -var bgWoodBytes
//go:generate file2byteslice -input ./stall.png -output stall.go -package assets -var stallBytes

package assets

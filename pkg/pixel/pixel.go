package pixel

type Pixel struct {
	R, G, B, A uint32
}

const MaxAlpha = 1<<16 - 1

var TrueBlack = Pixel{R: 0, G: 0, B: 0, A: MaxAlpha}

func (p Pixel) Is(r, g, b, a uint32) bool {
	return p.R == r && p.G == g && p.B == b && p.A == a
}

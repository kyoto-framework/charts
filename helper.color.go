package ktc

import (
	"strconv"

	"github.com/wcharczuk/go-chart/v2/drawing"
)

func colorChartFromUint(s [4]uint8) drawing.Color {
	return drawing.Color{
		R: s[0],
		G: s[1],
		B: s[2],
		A: s[3],
	}
}

func parseHex(hex string) uint8 {
	v, _ := strconv.ParseInt(hex, 16, 16)
	return uint8(v)
}

func ColorUintFromHex(hex string) [4]uint8 {
	if len(hex) == 3 { // RGB, 1 symbol each
		return [4]uint8{
			parseHex(string(hex[0])) * 0x11,
			parseHex(string(hex[1])) * 0x11,
			parseHex(string(hex[2])) * 0x11,
			255,
		}
	} else if len(hex) == 4 { // RGBA, 1 symbol each
		return [4]uint8{
			parseHex(string(hex[0])) * 0x11,
			parseHex(string(hex[1])) * 0x11,
			parseHex(string(hex[2])) * 0x11,
			parseHex(string(hex[3])) * 0x11,
		}
	} else if len(hex) == 6 { // RGB, 2 symbols each
		return [4]uint8{
			parseHex(string(hex[0:2])),
			parseHex(string(hex[2:4])),
			parseHex(string(hex[4:6])),
			255,
		}
	} else if len(hex) == 8 { // RGBA, 2 symbols each
		return [4]uint8{
			parseHex(string(hex[0:2])),
			parseHex(string(hex[2:4])),
			parseHex(string(hex[4:6])),
			parseHex(string(hex[6:8])),
		}
	} else {
		return [4]uint8{}
	}
}

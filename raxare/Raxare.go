package raxare

import (
	"fmt"
	"gorate/glicko"
)

func GetRaxare(r RatingSystem, p glicko.Player) int {
	midGlick := (r.GetTopPlayer().Rating + r.GetBottomPlayer().Rating)/2
	midPlayer := glicko.NewPlayer()
	midPlayer.Rating = midGlick
	midPlayer.Deviation = 350/2
	return int(p.GetChanceToBeat(midPlayer) * 20_000)
}

func RDStr(p glicko.Player) string {
	devC := float64(p.Deviation) / 350
	switch {
	case devC < 0.2:
			return "++"
	case devC < 0.3:
			return "+"
	case devC < 0.6:
			return "?"
	default:
		return "??"
	}
}

func GetPlayerStringRepr(r RatingSystem, p glicko.Player) string {
	rax := GetRaxare(r, p)
	devC := float64(p.Deviation) / 350
	retStr := "RXE: "
	if devC >= 0.5 {
		retStr += "???"
	} else {
		retStr += fmt.Sprint(rax)
	}
	retStr += "   GLK: "
	retStr += fmt.Sprint(p.Rating)
	retStr += RDStr(p)
	return retStr
}
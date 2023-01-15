package raxare

import (
	"gorate/glicko"
	"math"
)

type RatingSystem interface {
	GetTopPlayer() glicko.Player
	GetBottomPlayer() glicko.Player
	GetDefaultPlayer() glicko.Player
}

type SimpleRatingSystem struct {
	Players []glicko.Player
}

func (s SimpleRatingSystem) GetTopPlayer() glicko.Player {
	top := 0
	var plr glicko.Player
	for _, p := range s.Players {
		if p.Rating > top {
			top = p.Rating
			plr = p
		}
	}
	return plr
}


func (s SimpleRatingSystem) GetBottomPlayer() glicko.Player {
	top := int(math.Inf(1))
	var plr glicko.Player
	for _, p := range s.Players {
		if p.Rating < top {
			top = p.Rating
			plr = p
		}
	}
	return plr
}


func (s SimpleRatingSystem) GetDefaultPlayer() glicko.Player {
	return glicko.NewPlayer()
}

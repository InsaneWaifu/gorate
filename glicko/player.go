package glicko

import (
	"math"
)

type Player struct {
	Rating    int
	Deviation int
}

type Game struct {
	Opponent Player
	Result Result
}

type Result float64

const (
	Win  Result = 1
	Draw Result = 0.5
	Loss Result = 0
)

func NewPlayer() Player {
	return Player{Rating: 1500, Deviation: 350}
}

func (p Player) UpdateRating(timeSinceLastGame float64, games []Game) Player {
	// Factor in time to RD
	p.Deviation = int(
		math.Min(
			float64(p.Deviation) + timeSinceLastGame,
			350,
		),
	)

	var dSqrStart float64 = math.Pow(q, 2)
	var dSqrTimesBy float64
	var ratDiffTimesBy float64

	for _, g := range games {
		var Rj float64 = float64(g.Opponent.Rating)
		var RDj float64 = float64(g.Opponent.Deviation)
		var gRDj float64 = 1.0 /
			math.Sqrt(
				1 + ((3 * math.Pow(q, 2) * math.Pow(RDj, 2))/math.Pow(math.Pi, 2)),
			)
		var essr float64 = 1 / 
			(1 + math.Pow(10, (
				(gRDj * (float64(p.Rating) - Rj)) / -400)))


		dSqrTimesBy += math.Pow(gRDj, 2) * essr * (1 - essr)
		ratDiffTimesBy += gRDj * (float64(g.Result) - essr)
	}

	var dSqr = math.Pow(dSqrStart * dSqrTimesBy, -1)

	var ratDiffStart float64 = q / ((1/math.Pow(float64(p.Deviation), 2)) + (1/dSqr))
	var ratDiff float64 = ratDiffStart * ratDiffTimesBy
	var newRat int = p.Rating + int(ratDiff)
	
	var RDsqr = math.Pow(float64(p.Deviation), 2)
	var newRDSqr float64 = math.Pow((1/RDsqr) + (1/dSqr), -1)
	var newRD int = int(math.Sqrt(newRDSqr))

	p.Rating = newRat
	p.Deviation = newRD

	return p
}

func (p Player) GetConfidenceInterval() (int, int) {
	dScale := int(1.96 * float64(p.Deviation))
	return p.Rating - dScale, p.Rating + dScale
}

func (p Player) GetChanceToBeat(opp Player) float64 {
	var gArg float64 = math.Sqrt(math.Pow(float64(p.Deviation), 2) + math.Pow(float64(opp.Deviation), 2))
	var g float64 = 1.0 /
			math.Sqrt(
				1 + ((3 * math.Pow(q, 2) * math.Pow(gArg, 2))/math.Pow(math.Pi, 2)),
			)
	var e float64 = 1
	e /= 1 + math.Pow(10, (-g * (float64(p.Rating) - float64(opp.Rating)))/400)
	return e
}
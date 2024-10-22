// Copyright 2020 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Model representing the instantaneous score of a match.

package game

const (
	AutonBottomShelfPoints   = 5
	AutonTopShelfPoints      = 10
	TeleopBottomShelfPoints  = 3
	TeleopTopShelfPoints     = 7
	TeleopGoldenCubePoints   = 5
	EndgameParkPoints        = 2
	EndgameHamperBasePoints  = 5
	EndgameHamperStackPoints = 10
	FoulPoints               = 5
	TechFoulPoints           = 15
)

// would be const if go supported const arrays
var AutonTaxiPoints = [...]int{0, 3, 7}

type Score struct {
	Taxi         [2]AutonTaxiStatus
	Shelf        Shelf
	Hamper       int
	Parked       [2]bool
	GoldenCube   bool
	OppFouls     int
	OppTechFouls int
}

type AutonTaxiStatus int

const (
	AutonTaxiNone AutonTaxiStatus = iota
	AutonTaxiPartial
	AutonTaxiFull
)

type Shelf struct {
	AutonTopShelfCubes     int
	AutonBottomShelfCubes  int
	TeleopTopShelfCubes    int
	TeleopBottomShelfCubes int
}

func (shelf *Shelf) Equals(other *Shelf) bool {
	if shelf.AutonBottomShelfCubes != other.AutonBottomShelfCubes ||
		shelf.AutonTopShelfCubes != other.AutonTopShelfCubes ||
		shelf.TeleopBottomShelfCubes != other.TeleopBottomShelfCubes ||
		shelf.TeleopTopShelfCubes != other.TeleopTopShelfCubes {
		return false
	}
	return true
}

func (score *Score) AutoPoints() int {
	points := AutonTaxiPoints[score.Taxi[0]] + AutonTaxiPoints[score.Taxi[1]]
	points += score.Shelf.AutonBottomShelfCubes*AutonBottomShelfPoints + score.Shelf.AutonTopShelfCubes*AutonTopShelfPoints
	return points
}

func (score *Score) TeleopPoints() int {
	points := score.Shelf.TeleopBottomShelfCubes*TeleopBottomShelfPoints + score.Shelf.TeleopTopShelfCubes*TeleopTopShelfPoints
	if score.GoldenCube {
		points += TeleopGoldenCubePoints
	}
	return points
}

func (score *Score) EndgamePoints() int {
	points := 0
	if score.Hamper > 0 {
		points += EndgameHamperBasePoints
		points += (score.Hamper - 1) * EndgameHamperStackPoints
	}
	return points
}

func (score *Score) OppPenalties() int {
	return FoulPoints*score.OppFouls + TechFoulPoints*score.OppTechFouls
}

func (score *Score) TotalPoints() int {
	return score.AutoPoints() + score.TeleopPoints() + score.EndgamePoints() + score.OppPenalties()
}

// Calculates and returns the summary fields used for ranking and display.
func (score *Score) Summarize() *ScoreSummary {
	summary := new(ScoreSummary)

	summary.AutoPoints = score.AutoPoints()
	summary.TeleopPoints = score.TeleopPoints()
	summary.EndgamePoints = score.EndgamePoints()
	summary.OppPenalties = score.OppPenalties()
	summary.Score = summary.AutoPoints + summary.TeleopPoints + summary.EndgamePoints + summary.OppPenalties

	return summary
}

// Returns true if and only if all fields of the two scores are equal.
func (score *Score) Equals(other *Score) bool {
	if score.Taxi != other.Taxi ||
		!score.Shelf.Equals(&other.Shelf) ||
		score.Hamper != other.Hamper ||
		score.Parked != other.Parked ||
		score.GoldenCube != other.GoldenCube ||
		score.OppFouls != other.OppFouls ||
		score.OppTechFouls != other.OppTechFouls {
		return false
	}

	return true
}

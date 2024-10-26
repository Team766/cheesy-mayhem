// Copyright 2017 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Helper methods for use in tests in this package and others.

package game

func TestScore1() *Score {
	return &Score{
		Taxi: [2]AutonTaxiStatus{AutonTaxiPartial, AutonTaxiFull},
		Shelf: Shelf{
			AutonBottomShelfCubes:  2,
			AutonTopShelfCubes:     1,
			TeleopBottomShelfCubes: 4,
			TeleopTopShelfCubes:    4,
		},
		Hamper:       4,
		Park:         [2]bool{true, false},
		GoldenCube:   false,
		OppFouls:     0,
		OppTechFouls: 2,
	}
}

func TestScore2() *Score {
	return &Score{
		Taxi: [2]AutonTaxiStatus{AutonTaxiNone, AutonTaxiPartial},
		Shelf: Shelf{
			AutonBottomShelfCubes:  0,
			AutonTopShelfCubes:     2,
			TeleopBottomShelfCubes: 2,
			TeleopTopShelfCubes:    1,
		},
		Hamper:       2,
		Park:         [2]bool{false, false},
		GoldenCube:   true,
		OppFouls:     1,
		OppTechFouls: 0,
	}
}

func TestRanking1() *Ranking {
	return &Ranking{254, 1, 0, RankingFields{20, 625, 90, 554, 0.254, 3, 2, 1, 10}}
}

func TestRanking2() *Ranking {
	return &Ranking{1114, 2, 1, RankingFields{18, 700, 625, 90, 0.1114, 1, 3, 2, 10}}
}

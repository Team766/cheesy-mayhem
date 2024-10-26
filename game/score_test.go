// Copyright 2017 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)

package game

import (
	"fmt"
	"testing"

	"encoding/json"

	"github.com/stretchr/testify/assert"
)

func TestScoreSummary(t *testing.T) {
	redScore := TestScore1()
	blueScore := TestScore2()

	redSummary := redScore.Summarize()
	assert.Equal(t, 30, redSummary.AutoPoints)
	assert.Equal(t, 40, redSummary.TeleopPoints)
	assert.Equal(t, 35, redSummary.EndgamePoints)
	assert.Equal(t, 30, redSummary.OppPenalties)

	blueSummary := blueScore.Summarize()
	assert.Equal(t, 23, blueSummary.AutoPoints)
	assert.Equal(t, 18, blueSummary.TeleopPoints)
	assert.Equal(t, 15, blueSummary.EndgamePoints)
	assert.Equal(t, 5, blueSummary.OppPenalties)
}

func TestScoreEquals(t *testing.T) {
	score1 := TestScore1()
	score2 := TestScore1()
	assert.True(t, score1.Equals(score2))
	assert.True(t, score2.Equals(score1))

	score3 := TestScore2()
	assert.False(t, score1.Equals(score3))
	assert.False(t, score3.Equals(score1))

	score2 = TestScore1()
	score2.GoldenCube = true
	assert.False(t, score1.Equals(score2))
	assert.False(t, score2.Equals(score1))

	score2 = TestScore1()
	score2.OppFouls = 2
	assert.False(t, score1.Equals(score2))
	assert.False(t, score2.Equals(score1))

	score2 = TestScore1()
	score2.Taxi = [2]AutonTaxiStatus{AutonTaxiFull, AutonTaxiFull}
	assert.False(t, score1.Equals(score2))
	assert.False(t, score2.Equals(score1))
}

func TestScoreJson(t *testing.T) {
	score := TestScore1()
	json, err := json.Marshal(score)
	if err != nil {
		fmt.Println(json)
	}
	assert.Equal(t,
		"{\"Taxi\":[1,2],\"Shelf\":{\"AutonTopShelfCubes\":1,\"AutonBottomShelfCubes\":2,\"TeleopTopShelfCubes\":4,\"TeleopBottomShelfCubes\":4},\"Hamper\":4,\"Parked\":[true,false],\"GoldenCube\":false,\"OppFouls\":0,\"OppTechFouls\":2}",
		string(json))
}

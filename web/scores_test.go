// Copyright 2020 Team 254. All Rights Reserved.
// Author: kenschenke@gmail.com (Ken Schenke)

package web

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/Team254/cheesy-arena-lite/field"
	"github.com/Team254/cheesy-arena-lite/game"
	"github.com/stretchr/testify/assert"
)

func TestGetScores(t *testing.T) {
	web := setupTestWeb(t)

	score1 := game.TestScore1()
	score2 := game.TestScore2()

	web.arena.RedScore = score1
	web.arena.BlueScore = score2

	recorder := web.getHttpResponse("/api/scores")
	assert.Equal(t, 200, recorder.Code)

	var reqScores jsonScore
	json.Unmarshal(recorder.Body.Bytes(), &reqScores)
	assert.Equal(t, int(score1.Taxi[0]), reqScores.Red.Taxi[0])
	assert.Equal(t, int(score1.Taxi[1]), reqScores.Red.Taxi[1])
	assert.Equal(t, score1.Shelf.AutonBottomShelfCubes, reqScores.Red.Shelf.AutonBottomShelf)
	assert.Equal(t, score1.Shelf.AutonTopShelfCubes, reqScores.Red.Shelf.AutonTopShelf)
	assert.Equal(t, score1.Shelf.TeleopBottomShelfCubes, reqScores.Red.Shelf.TeleopBottomShelf)
	assert.Equal(t, score1.Shelf.TeleopTopShelfCubes, reqScores.Red.Shelf.TeleopBottomShelf)
	assert.Equal(t, score1.Hamper, *reqScores.Red.Hamper)
	assert.Equal(t, score1.GoldenCube, *reqScores.Red.GoldenCube)
	assert.Equal(t, score1.Park[0], reqScores.Red.Park[0])
	assert.Equal(t, score1.Park[1], reqScores.Red.Park[1])
	assert.Equal(t, score1.Fouls, *reqScores.Red.Foul)
	assert.Equal(t, score1.TechFouls, *reqScores.Red.TechFoul)

	assert.Equal(t, int(score2.Taxi[0]), reqScores.Blue.Taxi[0])
	assert.Equal(t, int(score2.Taxi[1]), reqScores.Blue.Taxi[1])
	assert.Equal(t, score2.Shelf.AutonBottomShelfCubes, reqScores.Blue.Shelf.AutonBottomShelf)
	assert.Equal(t, score2.Shelf.AutonTopShelfCubes, reqScores.Blue.Shelf.AutonTopShelf)
	assert.Equal(t, score2.Shelf.TeleopBottomShelfCubes, reqScores.Blue.Shelf.TeleopBottomShelf)
	assert.Equal(t, score2.Shelf.TeleopTopShelfCubes, reqScores.Blue.Shelf.TeleopTopShelf)
	assert.Equal(t, score2.Hamper, *reqScores.Blue.Hamper)
	assert.Equal(t, score2.GoldenCube, *reqScores.Blue.GoldenCube)
	assert.Equal(t, score2.Park[0], reqScores.Blue.Park[0])
	assert.Equal(t, score2.Park[1], reqScores.Blue.Park[1])
	assert.Equal(t, score2.Fouls, *reqScores.Blue.Foul)
	assert.Equal(t, score2.TechFouls, *reqScores.Blue.TechFoul)
}

func TestPatchScores(t *testing.T) {
	web := setupTestWeb(t)
	var recorder *httptest.ResponseRecorder

	web.arena.MatchState = field.PreMatch
	recorder = web.patchHttpResponse("/api/scores", "{}")
	assert.Equal(t, 400, recorder.Code)
	assert.Equal(t, "Score cannot be updated in this match state\n", recorder.Body.String())

	score1 := game.TestScore1()
	score2 := game.TestScore2()

	web.arena.RedScore = score1
	web.arena.BlueScore = score2

	web.arena.MatchState = field.PostMatch
	recorder = web.patchHttpResponse("/api/scores",
		"{\"red\":{\"taxi\": [2,2], \"hamper\": 2, \"foul\": 4, \"tech_foul\": 1}}")
	assert.Equal(t, 200, recorder.Code)

	assert.Equal(t, 34, web.arena.RedScore.AutoPoints())
	assert.Equal(t, 40, web.arena.RedScore.TeleopPoints())
	assert.Equal(t, 17, web.arena.RedScore.EndgamePoints())
	assert.Equal(t, 30, web.arena.RedScore.Penalties())
	assert.Equal(t, 23, web.arena.BlueScore.AutoPoints())
	assert.Equal(t, 18, web.arena.BlueScore.TeleopPoints())
	assert.Equal(t, 15, web.arena.BlueScore.EndgamePoints())
	assert.Equal(t, 20, web.arena.BlueScore.Penalties())

	recorder = web.patchHttpResponse("/api/scores",
		"{\"blue\":{\"shelf\": {\"auton_top\": 4, \"teleop_bottom\": 1 }, \"golden_cube\": true}}")
	assert.Equal(t, 200, recorder.Code)

	assert.Equal(t, 34, web.arena.RedScore.AutoPoints())
	assert.Equal(t, 40, web.arena.RedScore.TeleopPoints())
	assert.Equal(t, 17, web.arena.RedScore.EndgamePoints())
	assert.Equal(t, 30, web.arena.RedScore.Penalties())
	assert.Equal(t, 43, web.arena.BlueScore.AutoPoints())
	assert.Equal(t, 8, web.arena.BlueScore.TeleopPoints())
	assert.Equal(t, 15, web.arena.BlueScore.EndgamePoints())
	assert.Equal(t, 20, web.arena.BlueScore.Penalties())
}

func TestPutScores(t *testing.T) {
	web := setupTestWeb(t)
	var recorder *httptest.ResponseRecorder

	web.arena.MatchState = field.PreMatch
	recorder = web.putHttpResponse("/api/scores", "{}")
	assert.Equal(t, 400, recorder.Code)
	assert.Equal(t, "Score cannot be updated in this match state\n", recorder.Body.String())

	score1 := game.TestScore1()
	score2 := game.TestScore2()

	web.arena.RedScore = score1
	web.arena.BlueScore = score2

	web.arena.MatchState = field.PostMatch
	recorder = web.putHttpResponse("/api/scores",
		"{\"blue\":{\"taxi\": [2,2], \"hamper\": 2}}")
	assert.Equal(t, 200, recorder.Code)

	assert.Equal(t, 0, web.arena.RedScore.AutoPoints())
	assert.Equal(t, 0, web.arena.RedScore.TeleopPoints())
	assert.Equal(t, 0, web.arena.RedScore.EndgamePoints())
	assert.Equal(t, 14, web.arena.BlueScore.AutoPoints())
	assert.Equal(t, 0, web.arena.BlueScore.TeleopPoints())
	assert.Equal(t, 15, web.arena.BlueScore.EndgamePoints())
}

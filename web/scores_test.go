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

	web.arena.RedScore.Taxi[0] = score1.Taxi[0]
	web.arena.RedScore.Taxi[1] = score1.Taxi[1]
	web.arena.RedScore.Shelf.AutonBottomShelfCubes = score1.Shelf.AutonBottomShelfCubes
	web.arena.RedScore.Shelf.AutonTopShelfCubes = score1.Shelf.AutonTopShelfCubes
	web.arena.RedScore.Shelf.TeleopBottomShelfCubes = score1.Shelf.TeleopBottomShelfCubes
	web.arena.RedScore.Shelf.TeleopTopShelfCubes = score1.Shelf.TeleopTopShelfCubes
	web.arena.RedScore.Hamper = score1.Hamper
	web.arena.RedScore.Park = score1.Park
	web.arena.RedScore.GoldenCube = score1.GoldenCube
	web.arena.RedScore.OppFouls = score1.OppFouls
	web.arena.RedScore.OppTechFouls = score1.OppTechFouls

	web.arena.BlueScore.Taxi[0] = score2.Taxi[0]
	web.arena.BlueScore.Taxi[1] = score2.Taxi[1]
	web.arena.BlueScore.Shelf.AutonBottomShelfCubes = score2.Shelf.AutonBottomShelfCubes
	web.arena.BlueScore.Shelf.AutonTopShelfCubes = score2.Shelf.AutonTopShelfCubes
	web.arena.BlueScore.Shelf.TeleopBottomShelfCubes = score2.Shelf.TeleopBottomShelfCubes
	web.arena.BlueScore.Shelf.TeleopTopShelfCubes = score2.Shelf.TeleopTopShelfCubes
	web.arena.BlueScore.Hamper = score2.Hamper
	web.arena.BlueScore.Park = score2.Park
	web.arena.BlueScore.GoldenCube = score2.GoldenCube
	web.arena.BlueScore.OppFouls = score2.OppFouls
	web.arena.BlueScore.OppTechFouls = score2.OppTechFouls

	recorder := web.getHttpResponse("/api/scores")
	assert.Equal(t, 200, recorder.Code)

	var reqScores jsonScore
	json.Unmarshal(recorder.Body.Bytes(), &reqScores)
	// assert.Equal(t, "", recorder.Body.String())
	assert.Equal(t, int(score1.Taxi[0]), reqScores.Red.Taxi[0])
	assert.Equal(t, int(score1.Taxi[1]), reqScores.Red.Taxi[1])
	assert.Equal(t, score1.Shelf.AutonBottomShelfCubes, reqScores.Red.Shelf.AutonBottomShelf)
	assert.Equal(t, score1.Shelf.AutonTopShelfCubes, reqScores.Red.Shelf.AutonTopShelf)
	assert.Equal(t, score1.Shelf.TeleopBottomShelfCubes, reqScores.Red.Shelf.TeleopBottomShelf)
	assert.Equal(t, score1.Shelf.TeleopTopShelfCubes, reqScores.Red.Shelf.TeleopBottomShelf)
	assert.Equal(t, score1.Hamper, reqScores.Red.Hamper)
	assert.Equal(t, score1.GoldenCube, reqScores.Red.GoldenCube)
	assert.Equal(t, score1.Park[0], reqScores.Red.Park[0])
	assert.Equal(t, score1.Park[1], reqScores.Red.Park[1])
	assert.Equal(t, score1.OppFouls, reqScores.Blue.Foul)
	assert.Equal(t, score1.OppTechFouls, reqScores.Blue.TechFoul)

	assert.Equal(t, int(score2.Taxi[0]), reqScores.Blue.Taxi[0])
	assert.Equal(t, int(score2.Taxi[1]), reqScores.Blue.Taxi[1])
	assert.Equal(t, score2.Shelf.AutonBottomShelfCubes, reqScores.Blue.Shelf.AutonBottomShelf)
	assert.Equal(t, score2.Shelf.AutonTopShelfCubes, reqScores.Blue.Shelf.AutonTopShelf)
	assert.Equal(t, score2.Shelf.TeleopBottomShelfCubes, reqScores.Blue.Shelf.TeleopBottomShelf)
	assert.Equal(t, score2.Shelf.TeleopTopShelfCubes, reqScores.Blue.Shelf.TeleopTopShelf)
	assert.Equal(t, score2.Hamper, reqScores.Blue.Hamper)
	assert.Equal(t, score2.GoldenCube, reqScores.Blue.GoldenCube)
	assert.Equal(t, score2.Park[0], reqScores.Blue.Park[0])
	assert.Equal(t, score2.Park[1], reqScores.Blue.Park[1])
	assert.Equal(t, score2.OppFouls, reqScores.Red.Foul)
	assert.Equal(t, score2.OppTechFouls, reqScores.Red.TechFoul)
}

func TestPatchScores(t *testing.T) {
	web := setupTestWeb(t)
	var recorder *httptest.ResponseRecorder

	web.arena.MatchState = field.PreMatch
	recorder = web.patchHttpResponse("/api/scores", "{}")
	assert.Equal(t, 400, recorder.Code)
	assert.Equal(t, "Score cannot be updated in this match state\n", recorder.Body.String())

	// FIXME: update these tests after adding support for PATCH
	// score1 := game.TestScore1()
	// score2 := game.TestScore2()
	// web.arena.RedScore.AutoPoints = score1.AutoPoints
	// web.arena.RedScore.TeleopPoints = score1.TeleopPoints
	// web.arena.RedScore.EndgamePoints = score1.EndgamePoints
	// web.arena.BlueScore.AutoPoints = score2.AutoPoints
	// web.arena.BlueScore.TeleopPoints = score2.TeleopPoints
	// web.arena.BlueScore.EndgamePoints = score2.EndgamePoints

	web.arena.MatchState = field.PostMatch
	recorder = web.patchHttpResponse("/api/scores",
		"{\"red\":{\"auto\":5,\"teleop\":10,\"endgame\":15}}")
	assert.Equal(t, 200, recorder.Code)

	// assert.Equal(t, score1.AutoPoints+5, web.arena.RedScore.AutoPoints)
	// assert.Equal(t, score1.TeleopPoints+10, web.arena.RedScore.TeleopPoints)
	// assert.Equal(t, score1.EndgamePoints+15, web.arena.RedScore.EndgamePoints)
	// assert.Equal(t, score2.AutoPoints, web.arena.BlueScore.AutoPoints)
	// assert.Equal(t, score2.TeleopPoints, web.arena.BlueScore.TeleopPoints)
	// assert.Equal(t, score2.EndgamePoints, web.arena.BlueScore.EndgamePoints)

	recorder = web.patchHttpResponse("/api/scores",
		"{\"blue\":{\"auto\":-5,\"teleop\":-10,\"endgame\":-15}}")
	assert.Equal(t, 200, recorder.Code)

	// assert.Equal(t, score1.AutoPoints+5, web.arena.RedScore.AutoPoints)
	// assert.Equal(t, score1.TeleopPoints+10, web.arena.RedScore.TeleopPoints)
	// assert.Equal(t, score1.EndgamePoints+15, web.arena.RedScore.EndgamePoints)
	// assert.Equal(t, score2.AutoPoints-5, web.arena.BlueScore.AutoPoints)
	// assert.Equal(t, score2.TeleopPoints-10, web.arena.BlueScore.TeleopPoints)
	// assert.Equal(t, score2.EndgamePoints-15, web.arena.BlueScore.EndgamePoints)
}

func TestPutScores(t *testing.T) {
	web := setupTestWeb(t)
	var recorder *httptest.ResponseRecorder

	web.arena.MatchState = field.PreMatch
	recorder = web.putHttpResponse("/api/scores", "{}")
	assert.Equal(t, 400, recorder.Code)
	assert.Equal(t, "Score cannot be updated in this match state\n", recorder.Body.String())

	// FIXME: update these tests after adding support for PUT
	// score1 := game.TestScore1()
	// score2 := game.TestScore2()
	// web.arena.RedScore.AutoPoints = score1.AutoPoints
	// web.arena.RedScore.TeleopPoints = score1.TeleopPoints
	// web.arena.RedScore.EndgamePoints = score1.EndgamePoints
	// web.arena.BlueScore.AutoPoints = score2.AutoPoints
	// web.arena.BlueScore.TeleopPoints = score2.TeleopPoints
	// web.arena.BlueScore.EndgamePoints = score2.EndgamePoints

	web.arena.MatchState = field.PostMatch
	recorder = web.putHttpResponse("/api/scores",
		"{\"red\":{\"auto\":5,\"teleop\":10,\"endgame\":15}}")
	assert.Equal(t, 200, recorder.Code)

	// assert.Equal(t, 5, web.arena.RedScore.AutoPoints)
	// assert.Equal(t, 10, web.arena.RedScore.TeleopPoints)
	// assert.Equal(t, 15, web.arena.RedScore.EndgamePoints)
	// assert.Equal(t, 0, web.arena.BlueScore.AutoPoints)
	// assert.Equal(t, 0, web.arena.BlueScore.TeleopPoints)
	// assert.Equal(t, 0, web.arena.BlueScore.EndgamePoints)

	recorder = web.putHttpResponse("/api/scores",
		"{\"blue\":{\"auto\":5,\"teleop\":10,\"endgame\":15}}")
	assert.Equal(t, 200, recorder.Code)

	// assert.Equal(t, 0, web.arena.RedScore.AutoPoints)
	// assert.Equal(t, 0, web.arena.RedScore.TeleopPoints)
	// assert.Equal(t, 0, web.arena.RedScore.EndgamePoints)
	// assert.Equal(t, 5, web.arena.BlueScore.AutoPoints)
	// assert.Equal(t, 10, web.arena.BlueScore.TeleopPoints)
	// assert.Equal(t, 15, web.arena.BlueScore.EndgamePoints)
}

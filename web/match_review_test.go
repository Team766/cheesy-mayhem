// Copyright 2014 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)

package web

import (
	"fmt"
	"testing"

	"github.com/Team254/cheesy-arena-lite/game"
	"github.com/Team254/cheesy-arena-lite/model"
	"github.com/Team254/cheesy-arena-lite/tournament"
	"github.com/stretchr/testify/assert"
)

func TestMatchReview(t *testing.T) {
	web := setupTestWeb(t)

	match1 := model.Match{Type: "practice", DisplayName: "1", Status: game.RedWonMatch}
	match2 := model.Match{Type: "practice", DisplayName: "2"}
	match3 := model.Match{Type: "qualification", DisplayName: "1", Status: game.BlueWonMatch}
	match4 := model.Match{Type: "elimination", DisplayName: "SF1-1", Status: game.TieMatch}
	match5 := model.Match{Type: "elimination", DisplayName: "SF1-2"}
	web.arena.Database.CreateMatch(&match1)
	web.arena.Database.CreateMatch(&match2)
	web.arena.Database.CreateMatch(&match3)
	web.arena.Database.CreateMatch(&match4)
	web.arena.Database.CreateMatch(&match5)

	// Check that all matches are listed on the page.
	recorder := web.getHttpResponse("/match_review")
	assert.Equal(t, 200, recorder.Code)
	assert.Contains(t, recorder.Body.String(), ">P1<")
	assert.Contains(t, recorder.Body.String(), ">P2<")
	assert.Contains(t, recorder.Body.String(), ">Q1<")
	assert.Contains(t, recorder.Body.String(), ">SF1-1<")
	assert.Contains(t, recorder.Body.String(), ">SF1-2<")
}

func TestMatchReviewEditExistingResult(t *testing.T) {
	web := setupTestWeb(t)

	match := model.Match{Type: "elimination", DisplayName: "QF4-3", Status: game.RedWonMatch, Red1: 1001,
		Red2: 1002, Blue1: 1004, Blue2: 1005, ElimRedAlliance: 1, ElimBlueAlliance: 2}
	assert.Nil(t, web.arena.Database.CreateMatch(&match))
	matchResult := model.BuildTestMatchResult(match.Id, 1)
	matchResult.MatchType = match.Type
	assert.Nil(t, web.arena.Database.CreateMatchResult(matchResult))
	fmt.Println(matchResult)

	tournament.CreateTestAlliances(web.arena.Database, 2)
	web.arena.EventSettings.NumElimAlliances = 2
	web.arena.CreatePlayoffBracket()

	recorder := web.getHttpResponse("/match_review")
	assert.Equal(t, 200, recorder.Code)
	assert.Contains(t, recorder.Body.String(), ">QF4-3<")
	assert.Contains(t, recorder.Body.String(), ">127<") // The red score
	assert.Contains(t, recorder.Body.String(), ">61<")  // The blue score

	// Check response for non-existent match.
	recorder = web.getHttpResponse(fmt.Sprintf("/match_review/%d/edit", 12345))
	assert.Equal(t, 500, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "No such match")

	recorder = web.getHttpResponse(fmt.Sprintf("/match_review/%d/edit", match.Id))
	assert.Equal(t, 200, recorder.Code)
	assert.Contains(t, recorder.Body.String(), " QF4-3 ")

	// Update the score to something else.
	postBody := fmt.Sprintf(
		"matchResultJson={\"MatchId\":%d,\"RedScore\":{\"Taxi\":[0,2],\"Hamper\":4},"+
			"\"BlueScore\":{\"Taxi\":[1,1],\"Hamper\":2}}",
		match.Id,
	)
	recorder = web.postHttpResponse(fmt.Sprintf("/match_review/%d/edit", match.Id), postBody)
	assert.Equal(t, 303, recorder.Code, recorder.Body.String())

	// Check for the updated scores back on the match list page.
	recorder = web.getHttpResponse("/match_review")
	assert.Equal(t, 200, recorder.Code)
	assert.Contains(t, recorder.Body.String(), ">QF4-3<")
	assert.Contains(t, recorder.Body.String(), ">42<") // The red score
	assert.Contains(t, recorder.Body.String(), ">21<") // The blue score
}

func TestMatchReviewCreateNewResult(t *testing.T) {
	web := setupTestWeb(t)

	match := model.Match{Type: "elimination", DisplayName: "QF4-3", Status: game.RedWonMatch, Red1: 1001,
		Red2: 1002, Red3: 1003, Blue1: 1004, Blue2: 1005, Blue3: 1006, ElimRedAlliance: 1, ElimBlueAlliance: 2}
	web.arena.Database.CreateMatch(&match)
	tournament.CreateTestAlliances(web.arena.Database, 2)
	web.arena.EventSettings.NumElimAlliances = 2
	web.arena.CreatePlayoffBracket()

	recorder := web.getHttpResponse("/match_review")
	assert.Equal(t, 200, recorder.Code)
	assert.Contains(t, recorder.Body.String(), ">QF4-3<")
	assert.NotContains(t, recorder.Body.String(), ">71<") // The red score
	assert.NotContains(t, recorder.Body.String(), ">72<") // The blue score

	recorder = web.getHttpResponse(fmt.Sprintf("/match_review/%d/edit", match.Id))
	assert.Equal(t, 200, recorder.Code)
	assert.Contains(t, recorder.Body.String(), " QF4-3 ")

	// Update the score to something else.
	// Update the score to something else.
	postBody := fmt.Sprintf(
		"matchResultJson={\"MatchId\":%d,\"RedScore\":{\"Taxi\":[0,2],\"Hamper\":4},"+
			"\"BlueScore\":{\"Taxi\":[1,1],\"Hamper\":2}}",
		match.Id,
	)
	recorder = web.postHttpResponse(fmt.Sprintf("/match_review/%d/edit", match.Id), postBody)
	assert.Equal(t, 303, recorder.Code, recorder.Body.String())

	// Check for the updated scores back on the match list page.
	recorder = web.getHttpResponse("/match_review")
	assert.Equal(t, 200, recorder.Code)
	assert.Contains(t, recorder.Body.String(), ">QF4-3<")
	assert.Contains(t, recorder.Body.String(), ">42<") // The red score
	assert.Contains(t, recorder.Body.String(), ">21<") // The blue score
}

func TestMatchReviewEditCurrentMatch(t *testing.T) {
	web := setupTestWeb(t)

	match := model.Match{
		Type:        "qualification",
		DisplayName: "352",
		Red1:        1001,
		Red2:        1002,
		Red3:        1003,
		Blue1:       1004,
		Blue2:       1005,
		Blue3:       1006,
	}
	web.arena.Database.CreateMatch(&match)
	web.arena.LoadMatch(&match)
	assert.Equal(t, match, *web.arena.CurrentMatch)

	recorder := web.getHttpResponse("/match_review/current/edit")
	assert.Equal(t, 200, recorder.Code)
	assert.Contains(t, recorder.Body.String(), " 352 ")

	postBody := fmt.Sprintf(
		"matchResultJson={\"MatchId\":%d,\"RedScore\":{\"Taxi\":[0,2],\"Hamper\":4},"+
			"\"BlueScore\":{\"Taxi\":[1,1],\"Hamper\":2}}",
		match.Id,
	)
	recorder = web.postHttpResponse("/match_review/current/edit", postBody)
	assert.Equal(t, 303, recorder.Code, recorder.Body.String())
	assert.Equal(t, "/match_play", recorder.Header().Get("Location"))

	// Check that the persisted match is still unedited and that the realtime scores have been updated instead.
	match2, _ := web.arena.Database.GetMatchById(match.Id)
	assert.Equal(t, game.MatchNotPlayed, match2.Status)
	assert.Equal(t, 7, web.arena.RedScore.AutoPoints())
	assert.Equal(t, 0, web.arena.RedScore.TeleopPoints())
	assert.Equal(t, 35, web.arena.RedScore.EndgamePoints())
	assert.Equal(t, 6, web.arena.BlueScore.AutoPoints())
	assert.Equal(t, 0, web.arena.BlueScore.TeleopPoints())
	assert.Equal(t, 15, web.arena.BlueScore.EndgamePoints())
}

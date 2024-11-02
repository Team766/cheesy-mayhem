// Copyright 2020 Team 254. All Rights Reserved.
// Author: kenschenke@gmail.com (Ken Schenke)
//
// Web handlers for handling realtime scores API.

/*

FIXME: update all of this.

API Docs

JSON Schema:

{
	FIXME
}

GET http://10.0.100.5/api/scores

Returns current score.

PUT http://10.0.100.5/api/scores

Sets the current scores from the request body. All
parts are optional. Anything missing is set to zero.

Example:

{
	FIXME
}

Red teleop and endgame are set to zero as well as all blue scores.

PATCH http://10.0.100.5/api/scores

Adds or subtracts the current scores from the request
body. All parts are optional. Scores missing from the
request body are left untouched.

Example:

{
	FIXME
}

FIXME
10 is added to red auto. Red teleop and endgame are left untouched.
5 is subtracted from blue teleop. Blue auto and endgame are left untouched.

*/

package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Team254/cheesy-arena-lite/field"
	"github.com/Team254/cheesy-arena-lite/game"
)

type jsonShelf struct {
	AutonBottomShelf  int `json:"auton_bottom"`
	AutonTopShelf     int `json:"auton_top"`
	TeleopBottomShelf int `json:"teleop_bottom"`
	TeleopTopShelf    int `json:"teleop_top"`
}

type jsonAllianceScore struct {
	Taxi       [2]int    `json:"taxi"`
	Shelf      jsonShelf `json:"shelf"`
	Hamper     int       `json:"hamper"`
	Park       [2]bool   `json:"park"`
	GoldenCube bool      `json:"golden_cube"`
	Foul       int       `json:"foul"`
	TechFoul   int       `json:"tech_foul"`
}

type jsonScore struct {
	Red  jsonAllianceScore `json:"red"`
	Blue jsonAllianceScore `json:"blue"`
}

func getJsonForScore(score *game.Score) jsonAllianceScore {
	return jsonAllianceScore{
		Taxi: [2]int{int(score.Taxi[0]), int(score.Taxi[1])},
		Shelf: jsonShelf{
			AutonBottomShelf:  score.Shelf.AutonBottomShelfCubes,
			AutonTopShelf:     score.Shelf.AutonTopShelfCubes,
			TeleopBottomShelf: score.Shelf.TeleopBottomShelfCubes,
			TeleopTopShelf:    score.Shelf.TeleopTopShelfCubes,
		},
		Hamper:     score.Hamper,
		Park:       [2]bool{score.Park[0], score.Park[1]},
		GoldenCube: score.GoldenCube,
		Foul:       score.Fouls,
		TechFoul:   score.TechFouls,
	}
}

func getTaxiFromJsonField(taxi [2]int) [2]game.AutonTaxiStatus {
	return [2]game.AutonTaxiStatus{game.AutonTaxiStatus(taxi[0]), game.AutonTaxiStatus(taxi[1])}
}

func getShelfFromJsonField(shelf jsonShelf) game.Shelf {
	return game.Shelf{
		AutonBottomShelfCubes:  shelf.AutonBottomShelf,
		AutonTopShelfCubes:     shelf.AutonTopShelf,
		TeleopBottomShelfCubes: shelf.TeleopBottomShelf,
		TeleopTopShelfCubes:    shelf.TeleopTopShelf,
	}
}

func updateScoreFromJson(json jsonAllianceScore, scoreMap map[string]interface{}, score *game.Score) {
	fmt.Println(json)
	if _, ok := scoreMap["taxi"]; ok {
		score.Taxi = getTaxiFromJsonField(json.Taxi)
	}

	if _, ok := scoreMap["shelf"]; ok {
		score.Shelf = getShelfFromJsonField(json.Shelf)
	}

	if _, ok := scoreMap["golden_cube"]; ok {
		score.GoldenCube = json.GoldenCube
	}

	if _, ok := scoreMap["hamper"]; ok {
		score.Hamper = json.Hamper
	}

	if _, ok := scoreMap["park"]; ok {
		score.Park = json.Park
	}

	// TODO: add support for penalties
}

func (web *Web) getScoresHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(jsonScore{
		Red:  getJsonForScore(web.arena.RedScore),
		Blue: getJsonForScore(web.arena.BlueScore),
	})
}

func (web *Web) setScoresHandler(w http.ResponseWriter, r *http.Request) {
	if web.arena.MatchState == field.PreMatch || web.arena.MatchState == field.TimeoutActive ||
		web.arena.MatchState == field.PostTimeout {
		http.Error(w, "Score cannot be updated in this match state", http.StatusBadRequest)
		return
	}

	var scores jsonScore
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleWebErr(w, err)
		return
	}

	var scoresMap map[string]interface{}

	// ick - is there a way to parse into the struct - and tell which fields are explicitly filled?
	json.Unmarshal(reqBody, &scoresMap)
	json.Unmarshal(reqBody, &scores)

	if r.Method == "PUT" {
		web.arena.RedScore = new(game.Score)
		web.arena.BlueScore = new(game.Score)
	}

	// FIXME: update this logic
	if red, ok := scoresMap["red"].(map[string]interface{}); ok {
		updateScoreFromJson(scores.Red, red, web.arena.RedScore)
	}

	if blue, ok := scoresMap["blue"].(map[string]interface{}); ok {
		updateScoreFromJson(scores.Blue, blue, web.arena.BlueScore)
	}

	// TODO: return current scores?
	web.arena.RealtimeScoreNotifier.Notify()
}

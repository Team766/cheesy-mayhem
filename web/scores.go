// Copyright 2020 Team 254. All Rights Reserved.
// Author: kenschenke@gmail.com (Ken Schenke)
//
// Web handlers for handling realtime scores API.

/*

API Docs

JSON Schema:

  “red”: { “taxi”: [int,int],
          “shelf”: {
            “auton_bottom”: int,
            “auton_top”: int,
            “teleop_bottom”: int,
            “teleop_top”: int
          },
          “hamper”: int,
          “park”: [bool,bool],
          “golden_cube”: bool
          “foul”: int,
          “tech_foul”, int },
  “blue”: { ... }

(taxi values: NONE = 0, PARTIAL = 1, FULL = 2)

GET http://10.0.100.5/api/scores

Returns current score.

PUT http://10.0.100.5/api/scores

Sets the current scores from the request body. All
parts are optional. Anything missing is set to zero.

Example:

{
	  “red”: { “taxi”: [2, 1],
          “shelf”: {
            “auton_bottom”: 1,
            “auton_top”: 2,
            “teleop_bottom”: 2,
            “teleop_top”: 3,
          }
}

Red hamper, park, golden_cube, and fouls are set to zero as well as all blue scores.

PATCH http://10.0.100.5/api/scores

Sets just the provided portion of the score.  Fields that are provided override existing
values; the rest are left untouched.

Example:

{
	  "blue”: {
          “shelf”: {
            “auton_bottom”: 1,
            “auton_top”: 2,
            “teleop_bottom”: 2,
            “teleop_top”: 3,
          }},
	  "red": {
		"hamper": 2,
	  }

The blue shelf status and red hamper count are updated.  The blue taxi, hamper, park, golden_cube, and fouls,
and red taxi, shelf, park, golden_cube, and fouls are left untouched.
*/

package web

import (
	"encoding/json"
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
	Taxi       *[2]int    `json:"taxi"`
	Shelf      *jsonShelf `json:"shelf"`
	Hamper     *int       `json:"hamper"`
	Park       *[2]bool   `json:"park"`
	GoldenCube *bool      `json:"golden_cube"`
	Foul       *int       `json:"foul"`
	TechFoul   *int       `json:"tech_foul"`
}

type jsonScore struct {
	Red  *jsonAllianceScore `json:"red"`
	Blue *jsonAllianceScore `json:"blue"`
}

func getJsonForScore(score *game.Score) *jsonAllianceScore {
	return &jsonAllianceScore{
		Taxi: &[2]int{int(score.Taxi[0]), int(score.Taxi[1])},
		Shelf: &jsonShelf{
			AutonBottomShelf:  score.Shelf.AutonBottomShelfCubes,
			AutonTopShelf:     score.Shelf.AutonTopShelfCubes,
			TeleopBottomShelf: score.Shelf.TeleopBottomShelfCubes,
			TeleopTopShelf:    score.Shelf.TeleopTopShelfCubes,
		},
		Hamper:     &score.Hamper,
		Park:       &[2]bool{score.Park[0], score.Park[1]},
		GoldenCube: &score.GoldenCube,
		Foul:       &score.Fouls,
		TechFoul:   &score.TechFouls,
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

func updateScoreFromJson(json jsonAllianceScore, score *game.Score) {
	if json.Taxi != nil {
		score.Taxi = getTaxiFromJsonField(*json.Taxi)
	}

	if json.Shelf != nil {
		score.Shelf = getShelfFromJsonField(*json.Shelf)
	}

	if json.GoldenCube != nil {
		score.GoldenCube = *json.GoldenCube
	}

	if json.Hamper != nil {
		score.Hamper = *json.Hamper
	}

	if json.Park != nil {
		score.Park = *json.Park
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

	json.Unmarshal(reqBody, &scores)

	if r.Method == "PUT" {
		web.arena.RedScore = new(game.Score)
		web.arena.BlueScore = new(game.Score)
	}

	if scores.Red != nil {
		updateScoreFromJson(*scores.Red, web.arena.RedScore)
	}

	if scores.Blue != nil {
		updateScoreFromJson(*scores.Blue, web.arena.BlueScore)
	}

	// TODO: return current scores?
	web.arena.RealtimeScoreNotifier.Notify()
}

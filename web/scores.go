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

func (web *Web) getScoresHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(jsonScore{
		Red: jsonAllianceScore{
			Taxi: [2]int{int(web.arena.RedScore.Taxi[0]), int(web.arena.RedScore.Taxi[1])},
			Shelf: jsonShelf{
				AutonBottomShelf:  web.arena.RedScore.Shelf.AutonBottomShelfCubes,
				AutonTopShelf:     web.arena.RedScore.Shelf.AutonTopShelfCubes,
				TeleopBottomShelf: web.arena.RedScore.Shelf.TeleopBottomShelfCubes,
				TeleopTopShelf:    web.arena.RedScore.Shelf.TeleopTopShelfCubes,
			},
			Hamper:     web.arena.RedScore.Hamper,
			Park:       [2]bool{web.arena.RedScore.Park[0], web.arena.RedScore.Park[1]},
			GoldenCube: web.arena.RedScore.GoldenCube,
			Foul:       web.arena.BlueScore.OppFouls,
			TechFoul:   web.arena.BlueScore.OppTechFouls,
		},
		Blue: jsonAllianceScore{
			Taxi: [2]int{int(web.arena.BlueScore.Taxi[0]), int(web.arena.BlueScore.Taxi[1])},
			Shelf: jsonShelf{
				AutonBottomShelf:  web.arena.BlueScore.Shelf.AutonBottomShelfCubes,
				AutonTopShelf:     web.arena.BlueScore.Shelf.AutonTopShelfCubes,
				TeleopBottomShelf: web.arena.BlueScore.Shelf.TeleopBottomShelfCubes,
				TeleopTopShelf:    web.arena.BlueScore.Shelf.TeleopTopShelfCubes,
			},
			Hamper:     web.arena.BlueScore.Hamper,
			Park:       [2]bool{web.arena.BlueScore.Park[0], web.arena.BlueScore.Park[1]},
			GoldenCube: web.arena.BlueScore.GoldenCube,
			Foul:       web.arena.RedScore.OppFouls,
			TechFoul:   web.arena.RedScore.OppTechFouls,
		},
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
		// TODO: remove printlns
		// TODO: refactor below so code can be reused across red, blue portions of json
		fmt.Println(red)
		fmt.Println("Updating red score")
		if _, ok := red["taxi"]; ok {
			fmt.Println("Updating taxi status")
			web.arena.RedScore.Taxi = [2]game.AutonTaxiStatus{game.AutonTaxiStatus(scores.Red.Taxi[0]), game.AutonTaxiStatus(scores.Red.Taxi[1])}
		}

		if _, ok := red["shelf"]; ok {
			fmt.Println("Updating shelf status")

			web.arena.RedScore.Shelf = game.Shelf{
				AutonBottomShelfCubes:  scores.Red.Shelf.AutonBottomShelf,
				AutonTopShelfCubes:     scores.Red.Shelf.AutonTopShelf,
				TeleopBottomShelfCubes: scores.Red.Shelf.TeleopBottomShelf,
				TeleopTopShelfCubes:    scores.Red.Shelf.TeleopTopShelf,
			}
		}

		if _, ok := red["golden_cube"]; ok {
			fmt.Println("Updating golden_cube status")

			web.arena.RedScore.GoldenCube = scores.Red.GoldenCube
		}

		if _, ok := red["hamper"]; ok {
			fmt.Println("Updating hamper status")

			web.arena.RedScore.Hamper = scores.Red.Hamper
		}

		if _, ok := red["park"]; ok {
			fmt.Println("Updating park status")

			web.arena.RedScore.Park = scores.Red.Park
		}

		// TODO: add support for penalties
	}

	if _, ok := scoresMap["blue"]; ok {
	}

	web.arena.RealtimeScoreNotifier.Notify()
}

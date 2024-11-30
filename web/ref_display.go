package web

import (
	"fmt"
	"net/http"

	"github.com/Team254/cheesy-arena-lite/model"
	"github.com/Team254/cheesy-arena-lite/websocket"
	"github.com/gorilla/mux"
)

// renders the ref UI for an individual alliance
// each alliance will typically have two individual refs, with each ref scoring non-overlapping
// field elements.  through a websockets integration, each ref will be able to see score updates
// submitted by other refs in realtime.
func (web *Web) refDisplayHandler(w http.ResponseWriter, r *http.Request) {
	if !web.userIsExpected(w, r, []string{adminUser, refUser}) {
		return
	}

	vars := mux.Vars(r)
	alliance := vars["alliance"]
	if alliance != "red" && alliance != "blue" {
		handleWebErr(w, fmt.Errorf("Invalid alliance: '%s'", alliance))
		return
	}
	// if !web.enforceDisplayConfiguration(w, r, map[string]string{"background": "#0f0", "reversed": "false",
	// 	"overlayLocation": "bottom"}) {
	// 	return
	// }

	template, err := web.parseFiles("templates/ref_display.html", "templates/base.html")
	if err != nil {
		handleWebErr(w, err)
		return
	}

	shelfLocations := [2]string{"Bottom", "Top"}

	data := struct {
		*model.EventSettings
		Match          *model.Match
		Alliance       string
		ShelfLocations [2]string
	}{web.arena.EventSettings, web.arena.CurrentMatch, alliance, shelfLocations}
	err = template.ExecuteTemplate(w, "base_no_navbar", data)
	if err != nil {
		handleWebErr(w, err)
		return
	}
}

// The websocket endpoint for the ref display client to receive status updates.
func (web *Web) refDisplayWebsocketHandler(w http.ResponseWriter, r *http.Request) {
	if !web.userIsExpected(w, r, []string{adminUser, refUser}) {
		return
	}

	ws, err := websocket.NewWebsocket(w, r)
	if err != nil {
		handleWebErr(w, err)
		return
	}
	defer ws.Close()

	// Subscribe the websocket to the notifiers whose messages will be passed on to the client.
	ws.HandleNotifiers(web.arena.MatchLoadNotifier, web.arena.MatchTimingNotifier, web.arena.MatchTimeNotifier, web.arena.RealtimeScoreNotifier)
}

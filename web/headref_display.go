package web

import (
	"net/http"

	"github.com/Team254/cheesy-arena-lite/model"
	"github.com/Team254/cheesy-arena-lite/websocket"
)

// renders the head ref UI, allowing the head ref to view and edit scores and fouls for both alliances.
// through a websockets integration, the head ref will be able to see score updates submitted by other
// refs in realtime.
func (web *Web) headrefDisplayHandler(w http.ResponseWriter, r *http.Request) {
	if !web.userIsAdmin(w, r) {
		return
	}

	// if !web.enforceDisplayConfiguration(w, r, map[string]string{"background": "#0f0", "reversed": "false",
	// 	"overlayLocation": "bottom"}) {
	// 	return
	// }

	template, err := web.parseFiles("templates/headref_display.html", "templates/base.html")
	if err != nil {
		handleWebErr(w, err)
		return
	}

	alliances := [2]string{"red", "blue"}
	shelfLocations := [2]string{"Bottom", "Top"}

	data := struct {
		*model.EventSettings
		Match          *model.Match
		Alliances      [2]string
		ShelfLocations [2]string
	}{web.arena.EventSettings, web.arena.CurrentMatch, alliances, shelfLocations}

	err = template.ExecuteTemplate(w, "base_no_navbar", data)
	if err != nil {
		handleWebErr(w, err)
		return
	}
}

// The websocket endpoint for the ref display client to receive status updates.
func (web *Web) headrefDisplayWebsocketHandler(w http.ResponseWriter, r *http.Request) {
	if !web.userIsAdmin(w, r) {
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

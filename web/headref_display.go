package web

import (
	"net/http"

	"github.com/Team254/cheesy-arena-lite/model"
	"github.com/Team254/cheesy-arena-lite/websocket"
)

// Renders the audience display to be chroma keyed over the video feed.
func (web *Web) headrefDisplayHandler(w http.ResponseWriter, r *http.Request) {
	// if !web.enforceDisplayConfiguration(w, r, map[string]string{"background": "#0f0", "reversed": "false",
	// 	"overlayLocation": "bottom"}) {
	// 	return
	// }

	template, err := web.parseFiles("templates/headref_display.html", "templates/base.html")
	if err != nil {
		handleWebErr(w, err)
		return
	}

	alliances := [2]string { "red", "blue" }

	data := struct {
		*model.EventSettings
		Match    *model.Match
		Alliances [2]string
	}{web.arena.EventSettings, web.arena.CurrentMatch, alliances}
	err = template.ExecuteTemplate(w, "base_no_navbar", data)
	if err != nil {
		handleWebErr(w, err)
		return
	}
}

// The websocket endpoint for the ref display client to receive status updates.
func (web *Web) headrefDisplayWebsocketHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.NewWebsocket(w, r)
	if err != nil {
		handleWebErr(w, err)
		return
	}
	defer ws.Close()

	// Subscribe the websocket to the notifiers whose messages will be passed on to the client.
	ws.HandleNotifiers(web.arena.MatchLoadNotifier, web.arena.MatchTimingNotifier, web.arena.MatchTimeNotifier, web.arena.RealtimeScoreNotifier)
}

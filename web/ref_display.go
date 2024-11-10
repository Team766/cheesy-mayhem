package web

import (
	"fmt"
	"net/http"

	"github.com/Team254/cheesy-arena-lite/model"
	"github.com/gorilla/mux"
)

// Renders the audience display to be chroma keyed over the video feed.
func (web *Web) refDisplayHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	alliance := vars["alliance"]
	if alliance != "red" && alliance != "blue" {
		handleWebErr(w, fmt.Errorf("Invalid alliance: '%s'. (%s)", alliance, r))
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

	data := struct {
		*model.EventSettings
		alliance string
	}{web.arena.EventSettings, alliance}
	err = template.ExecuteTemplate(w, "base_no_navbar", data)
	if err != nil {
		handleWebErr(w, err)
		return
	}

}

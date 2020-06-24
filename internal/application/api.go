/*
#######
##                                      _                __
##       ___ _______ _  ___ ___        (_)__  ___ ____  / /____
##      / _ `/ __/  ' \/ -_) _ \  _   / / _ \/ _ `/ _ \/ __/ -_)
##      \_,_/_/ /_/_/_/\__/_//_/ (_) /_/_//_/\_,_/ .__/\__/\__/
##                                              /_/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package application

import (
	"fmt"
	"html"
	"net/http"

	"github.com/arnumina/armen.core/pkg/component"
)

func (a *Application) setAPI(cm component.Manager) {
	a.router = cm.Server().NewAppRouter(a.Name())

	a.router.HandleFunc(
		"/hello",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		},
	).Methods("GET")
}

/*
######################################################################################################## @(°_°)@ #######
*/

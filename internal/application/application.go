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
	"github.com/arnumina/armen.core/pkg/component"
	"github.com/arnumina/armen.core/pkg/util"
	"github.com/arnumina/armen.core/pkg/value"
	"github.com/gorilla/mux"

	"github.com/arnumina/armen.inapte/internal/resources"
)

type (
	// Application AFAIRE
	Application struct {
		logger    component.Logger
		resources *resources.Resources
		router    *mux.Router
	}
)

// New AFAIRE
func New(name string) *Application {
	return &Application{
		resources: resources.New(name),
	}
}

// Name AFAIRE
func (a *Application) Name() string {
	return a.resources.AppName()
}

// Build AFAIRE
func (a *Application) Build(cm component.Manager, cfg *value.Value) error {
	a.logger = cm.Logger().Clone(util.LoggerPrefix(a.Name(), util.NewUUID()))

	if err := a.subscribe(cm.Bus()); err != nil {
		return err
	}

	if err := a.resources.Build(cm, cfg); err != nil {
		return err
	}

	a.setAPI(cm)

	return nil
}

// Close AFAIRE
func (a *Application) Close() {
	a.resources.Close()
}

/*
######################################################################################################## @(°_°)@ #######
*/

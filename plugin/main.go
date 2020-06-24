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

package main

import (
	"time"

	"github.com/arnumina/armen.core/pkg/util"

	"github.com/arnumina/armen.inapte/internal/application"
)

var (
	_version string
	_builtAt string
)

type wrapper struct {
	*application.Application
}

// PluginName AFAIRE
func (w *wrapper) PluginName() string {
	return "armen.inapte"
}

// PluginVersion AFAIRE
func (w *wrapper) PluginVersion() string {
	return _version
}

// PluginBuiltAt AFAIRE
func (w *wrapper) PluginBuiltAt() time.Time {
	return util.UnixToTime(_builtAt)
}

// Export AFAIRE
func Export() interface{} {
	return &wrapper{
		Application: application.New("inapte"),
	}
}

func main() {
	_ = Export() // avoid linter errors ////////////////////////////////////////////////////////////////////////////////
}

/*
######################################################################################################## @(°_°)@ #######
*/

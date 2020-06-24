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

package resources

import (
	"github.com/arnumina/armen.core/pkg/component"
	"github.com/arnumina/armen.core/pkg/jw"
	"github.com/arnumina/armen.core/pkg/message"
	"github.com/arnumina/armen.core/pkg/value"
)

// Resources AFAIRE
type Resources struct {
	appName string
	channel chan<- *message.Message
	model   component.Model
}

// New AFAIRE
func New(appName string) *Resources {
	return &Resources{
		appName: appName,
	}
}

// Build AFAIRE
func (r *Resources) Build(cm component.Manager, _ *value.Value) error {
	r.channel = cm.Bus().NewPublisher(r.AppName(), 1, 1)
	r.model = cm.Model()

	return nil
}

// AppName AFAIRE
func (r *Resources) AppName() string {
	return r.appName
}

// MaybeInsertJob AFAIRE
func (r *Resources) MaybeInsertJob(job *jw.Job) (bool, error) {
	return r.model.MaybeInsertJob(job)
}

// InsertWorkflow AFAIRE
func (r *Resources) InsertWorkflow(wf *jw.Workflow) error {
	return r.model.InsertWorkflow(wf)
}

// Close AFAIRE
func (r *Resources) Close() {
	close(r.channel)
}

/*
######################################################################################################## @(°_°)@ #######
*/

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
	"github.com/arnumina/armen.core/pkg/failure"
	"github.com/arnumina/armen.core/pkg/jw"
	"github.com/arnumina/armen.core/pkg/message"

	"github.com/arnumina/armen.inapte/internal/jobs/nop"
	"github.com/arnumina/armen.inapte/internal/jobs/test"
)

const (
	_typeJobNop  = "nop"
	_typeJobTest = "test"
)

func (a *Application) createJob(
	id, n, t, o string, c *string, e jw.Exclusivity, g *string, p jw.Priority, em *string) *jw.Job {
	return jw.NewJob(id, n, a.Name(), t, o, c, e, g, p, em)
}

func (a *Application) maybeInsertJob(job *jw.Job) (bool, error) {
	return a.resources.MaybeInsertJob(job)
}

func (a *Application) createJobTest(msg *message.Message) {
	var c = _typeJobTest
	job := a.createJob(msg.ID, _typeJobTest, _typeJobTest, msg.Publisher, &c, jw.Itself, nil, jw.Low, nil)
	_, _ = a.maybeInsertJob(job)
}

// RunJob AFAIRE
func (a *Application) RunJob(job *jw.Job, logger component.Logger) *jw.Result {
	var jwr jw.Runner

	switch job.Type {
	case _typeJobNop:
		jwr = nop.New(job, logger, a.resources)
	case _typeJobTest:
		jwr = test.New(job, logger, a.resources)
	default:
		return jw.NewResult(
			failure.New(nil).Msg("the type of this job is not valid"), /////////////////////////////////////////////////
			0,
		)
	}

	return jwr.Run()
}

/*
######################################################################################################## @(°_°)@ #######
*/

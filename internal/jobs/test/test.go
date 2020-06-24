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

package test

import (
	"github.com/arnumina/armen.core/pkg/failure"

	"github.com/arnumina/armen.core/pkg/component"
	"github.com/arnumina/armen.core/pkg/jw"
	"github.com/arnumina/armen.core/pkg/util"
	"github.com/arnumina/armen.inapte/internal/resources"
)

type (
	// JobTest AFAIRE
	JobTest struct {
		*jw.Job
		logger component.Logger
		res    *resources.Resources
	}
)

// New AFAIRE
func New(job *jw.Job, logger component.Logger, res *resources.Resources) *JobTest {
	return &JobTest{
		Job:    job,
		logger: logger,
		res:    res,
	}
}

// Run AFAIRE
func (j *JobTest) Run() *jw.Result {
	var steps = map[string]*jw.Step{
		"alpha": {
			Application: "inapte",
			Type:        "nop",
			Exclusivity: "no",
			Group:       nil,
			Config:      nil,
			Next: map[string]interface{}{
				"default": "beta",
			},
		},
		"beta": {
			Application: "inapte",
			Type:        "nop",
			Exclusivity: "no",
			Group:       nil,
			Config: map[string]interface{}{
				"day": 19,
			},
			Next: map[string]interface{}{
				"default": "gamma",
			},
		},
		"gamma": {
			Application: "inapte",
			Type:        "nop",
			Exclusivity: "no",
			Group:       nil,
			Config:      nil,
			Next:        nil,
		},
	}

	wf := jw.NewWorkflow(util.NewUUID(), "abg", "Workflow de test", "test", jw.None, nil, "alpha", steps)

	if err := j.res.InsertWorkflow(wf); err != nil {
		return jw.NewResult(failure.New(err), 0)
	}

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/

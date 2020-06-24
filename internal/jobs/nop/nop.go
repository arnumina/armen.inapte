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

package nop

import (
	"time"

	"github.com/arnumina/armen.core/pkg/component"
	"github.com/arnumina/armen.core/pkg/jw"
	"github.com/arnumina/armen.inapte/internal/resources"
)

type (
	// JobNop AFAIRE
	JobNop struct {
		*jw.Job
		logger component.Logger
		res    *resources.Resources
	}
)

// New AFAIRE
func New(job *jw.Job, logger component.Logger, res *resources.Resources) *JobNop {
	return &JobNop{
		Job:    job,
		logger: logger,
		res:    res,
	}
}

// Run AFAIRE
func (j *JobNop) Run() *jw.Result {
	j.Data[j.ID] = time.Now()
	j.Data["_"+j.Name] = j.ID

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/

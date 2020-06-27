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
	"github.com/arnumina/armen.core/pkg/message"
)

func (a *Application) msgHandler(msg *message.Message) {
	a.logger.Info( /////////////////////////////////////////////////////////////////////////////////////////////////////
		"Consume message",
		"id", msg.ID,
		"topic", msg.Topic,
		"publisher", msg.Publisher,
	)

	switch msg.Topic {
	case "create.job.inapte.test":
		_, _ = a.maybeInsertJob(a.createJobTest(msg))
	default:
		a.logger.Error( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"The topic of this message is not valid",
			"id", msg.ID,
			"topic", msg.Topic,
			"publisher", msg.Publisher,
		)
	}
}

func (a *Application) subscribe(bus component.Bus) error {
	if err := bus.Subscribe(
		a.msgHandler,
		"create[.]job[.]inapte[.].+",
	); err != nil {
		return err
	}

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/

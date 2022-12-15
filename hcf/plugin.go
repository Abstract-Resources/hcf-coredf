package hcf

import (
	"github.com/aabstractt/hcf-core/hcf/datasource"
	"github.com/aabstractt/hcf-core/hcf/profile"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	plugin *HCF = nil

	scoreboardTicker *time.Ticker = nil
)

type HCF struct {
	server *server.Server
	logger *logrus.Logger
}

func NewPlugin(srv *server.Server, logger *logrus.Logger) {
	plugin = &HCF{
		server: srv,
		logger: logger,
	}

	go func() {
		scoreboardTicker = time.NewTicker(time.Millisecond * 50)

		for {
			select {
			case <-scoreboardTicker.C:
				for _, profileVar := range profile.All() {
					profileVar.UpdateScoreboard()
				}
			}
		}
	}()

	for srv.Accept(func(player *player.Player) {
		logger.Infof("Successfully connected %v", player.Name())

		go func() {
			logger.Infof("Fetching the " + player.Name() + "'s profile stored on " + datasource.GetCurrentDataSource().GetName() + "!")

			profile.RegisterNewProfile(player, logger, datasource.GetCurrentDataSource().LoadProfileStorage(player.XUID()))
		}()

		//_ = profile.RegisterNewProfile(player, logger)

		/*p.RegisterHandler(handlers.NewHandleQuit(p.GetXuid()))
		p.RegisterHandler(handlers.NewHandleChat(p.GetXuid()))*/
	}) {}
}

func Plugin() *HCF {
	return plugin
}

func ScoreboardTicker() *time.Ticker  {
	return scoreboardTicker
}
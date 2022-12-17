package hcf

import (
	"github.com/aabstractt/hcf-core/hcf/command/faction"
	"github.com/aabstractt/hcf-core/hcf/datasource"
	"github.com/aabstractt/hcf-core/hcf/profile"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
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

	cmd.Register(cmd.New("faction", "Factions management", []string{"f"}, faction.CreateArgument{}))

	for srv.Accept(func(player *player.Player) {
		logger.Infof("Successfully connected %v", player.Name())

		go func() {
			logger.Infof("Fetching the " + player.Name() + "'s profile stored on " + datasource.GetCurrentDataSource().GetName() + "!")

			profile.RegisterNewProfile(player, logger, datasource.GetCurrentDataSource().LoadProfileStorage(player.XUID()))
		}()

		//_ = profile.RegisterNewProfile(player, logger)

		/*p.RegisterHandler(handlers.NewHandleQuit(p.XUID()))
		p.RegisterHandler(handlers.NewHandleChat(p.XUID()))*/
	}) {}
}

func Plugin() *HCF {
	return plugin
}

func Server() *server.Server {
	return plugin.server
}

func ScoreboardTicker() *time.Ticker  {
	return scoreboardTicker
}
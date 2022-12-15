package hcf

import (
	"github.com/aabstractt/hcf-core/hcf/datasource"
	"github.com/aabstractt/hcf-core/hcf/profile"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/sirupsen/logrus"
)

var (
	plugin *HCFCore = nil

	dataSource datasource.DataSource = nil
)

type HCFCore struct {

	server *server.Server
	logger *logrus.Logger
}

func Initialize(srv *server.Server, logger *logrus.Logger, dSource datasource.DataSource) {
	plugin = &HCFCore{
		server: srv,
		logger: logger,
	}

	dataSource = dSource

	for srv.Accept(func(player *player.Player) {
		logger.Infof("Successfully connected %v", player.Name())

		go func() {
			profile.RegisterNewProfile(player, logger, dataSource.FetchProfile(player.XUID(), player.Name()))
		}()

		//_ = profile.RegisterNewProfile(player, logger)

		/*p.RegisterHandler(handlers.NewHandleQuit(p.GetXuid()))
		p.RegisterHandler(handlers.NewHandleChat(p.GetXuid()))*/
	}) {}
}

func Plugin() *HCFCore {
	return plugin
}

func DataSource() datasource.DataSource {
	return dataSource
}
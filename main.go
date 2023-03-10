package main

import (
	"fmt"
	factionCommand "github.com/aabstractt/hcf-core/hcf/command/faction"
	"github.com/aabstractt/hcf-core/hcf/config"
	"github.com/aabstractt/hcf-core/hcf/datasource"
	"github.com/aabstractt/hcf-core/hcf/faction"
	"github.com/aabstractt/hcf-core/hcf/profile"
	"github.com/aabstractt/hcf-core/hcf/utils"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var scoreboardTicker *time.Ticker = nil

func main() {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{ForceColors: true}
	log.Level = logrus.DebugLevel

	chat.Global.Subscribe(chat.StdoutSubscriber{})

	conf, err := readConfig(log)
	if err != nil {
		log.Fatalf("error reading conf file: %v", err)
	}

	srvConf, err := readServerConfig()
	if err != nil {
		log.Fatalf("error reading conf file: %v", err)

		return
	}

	config.SetConfig(&srvConf)

	log.Warn("Server config was created! Please restart the server to modify that")

	chat.Global.Subscribe(chat.StdoutSubscriber{})

	srv := conf.New()
	srv.CloseOnProgramEnd()
	srv.Listen()

	utils.SetServer(srv)
	datasource.NewDataSource(log)
	go faction.RegisterFactionsStored()

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

	cmd.Register(cmd.New("faction", "Factions management", []string{"f"}, factionCommand.CreateArgument{}))

	for srv.Accept(func(player *player.Player) {
		log.Infof("Successfully connected %v", player.Name())

		go func() {
			log.Infof("Fetching the " + player.Name() + "'s profile stored on " + datasource.GetCurrentDataSource().GetName() + "!")

			profile.RegisterNewProfile(player, log, datasource.GetCurrentDataSource().LoadProfileStorage(player.XUID()))
		}()
	}) {}

	scoreboardTicker.Stop()

	// Flush all profiles stored into cache and save that on the db provider
	profile.Close()
}

// readConfig reads the configuration from the config.toml file, or creates the
// file if it does not yet exist.
func readConfig(log server.Logger) (server.Config, error) {
	c := server.DefaultConfig()
	var zero server.Config
	if _, err := os.Stat("config.toml"); os.IsNotExist(err) {
		data, err := toml.Marshal(c)
		if err != nil {
			return zero, fmt.Errorf("encode default config: %v", err)
		}
		if err := os.WriteFile("config.toml", data, 0644); err != nil {
			return zero, fmt.Errorf("create default config: %v", err)
		}
		return zero, nil
	}
	data, err := os.ReadFile("config.toml")
	if err != nil {
		return zero, fmt.Errorf("read config: %v", err)
	}

	err = toml.Unmarshal(data, &c)
	if err != nil {
		return zero, fmt.Errorf("decode config: %v", err)
	}

	return c.Config(log)
}

func readServerConfig() (config.ServerConfig, error) {
	c := config.DefaultConfig()
	var zero config.ServerConfig
	if _, err := os.Stat("server_config.toml"); os.IsNotExist(err) {
		data, err := toml.Marshal(c)
		if err != nil {
			return zero, fmt.Errorf("encode default config: %v", err)
		}
		if err := os.WriteFile("server_config.toml", data, 0644); err != nil {
			return zero, fmt.Errorf("create default config: %v", err)
		}
		return zero, nil
	}
	data, err := os.ReadFile("server_config.toml")
	if err != nil {
		return zero, fmt.Errorf("read config: %v", err)
	}

	err = toml.Unmarshal(data, &c)
	if err != nil {
		return zero, fmt.Errorf("decode config: %v", err)
	}

	return c, nil
}
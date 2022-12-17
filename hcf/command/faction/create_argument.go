package faction

import (
	"github.com/aabstractt/hcf-core/hcf/config"
	"github.com/aabstractt/hcf-core/hcf/faction"
	"github.com/aabstractt/hcf-core/hcf/profile"
	"github.com/aabstractt/hcf-core/hcf/utils"
	"github.com/aabstractt/hcf-core/hcf/utils/chat"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/google/uuid"
)

type CreateArgument struct {
	Create cmd.SubCommand `cmd:"create"`

	FactionName string `cmd:"name"`
}

func (clazz CreateArgument) Run(source cmd.Source, _ *cmd.Output) {
	sender, allowed := source.(*player.Player)

	if !allowed {
		sender.Message("Run this command in-game")

		return
	}

	pf := profile.GetIfLoaded(sender.XUID())
	if pf == nil {
		sender.Message("Run this command in-game")

		return
	}

	if faction.GetProfileFaction(pf) != nil {
		sender.Message(chat.YOU_ALREADY_IN_FACTION.Build())

		return
	}

	f := faction.GetFaction(clazz.FactionName)
	if f != nil {
		sender.Message(chat.FACTION_ALREADY_EXISTS.Build(f.Name()))

		return
	}

	srvConf := config.DefaultConfig()

	f = faction.NewFaction(
		uuid.New().String(),
		clazz.FactionName,
		sender.XUID(),
		srvConf.Factions.DTR,
		0,
		0,
		srvConf.Factions.Balance,
		srvConf.Factions.Points,
	)
	f.ForceSave()

	faction.RegisterNewFaction(f)
	faction.JoinFaction(pf, f, 0) // TODO: Change this and use an enum var

	message := chat.ReplacePlaceHolders("PLAYER_FACTION_CREATED", map[string]string{
		"player": sender.Name(),
		"faction": f.Name(),
	})
	for _, p := range utils.Server().Players() {
		p.Message(message)
	}
}
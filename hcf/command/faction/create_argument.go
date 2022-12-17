package faction

import (
	"github.com/aabstractt/hcf-core/hcf/faction"
	"github.com/aabstractt/hcf-core/hcf/profile"
	"github.com/aabstractt/hcf-core/hcf/utils"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

type CreateArgument struct {
	Create cmd.SubCommand `cmd:"create"`

	FactionName string `cmd:"name"`
}

func (clazz CreateArgument) Run(source cmd.Source, _ *cmd.Output) {
	sender, allowed := source.(*player.Player)

	if !allowed || profile.GetIfLoaded(sender.XUID()) == nil {
		sender.Message("Run this command in-game")

		return
	}

	if faction.GetPlayerFaction(sender.Name()) != nil {
		sender.Message(utils.ReplacePlaceHolders("YOU_ALREADY_IN_FACTION"))

		return
	}

	f := faction.GetFaction(clazz.FactionName)
	if f != nil {
		sender.Message(utils.ReplacePlaceHolders("FACTION_ALREADY_EXISTS", clazz.FactionName))

		return
	}

	// TODO: Create a new faction and register that
}
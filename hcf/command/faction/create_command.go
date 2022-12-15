package faction

import (
	"github.com/aabstractt/hcf-core/hcf/profile"
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

	/*kit.RegisterNewKit(clazz.KitName, sender.Inventory().Items(), sender.Armour().Items(), true)

	sender.Message("Kit " + clazz.KitName + " was successfully created!")*/
}
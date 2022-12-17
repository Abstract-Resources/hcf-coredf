package profile

import (
	"fmt"
	"github.com/aabstractt/hcf-core/hcf/datasource"
	"github.com/aabstractt/hcf-core/hcf/datasource/storage"
	"github.com/aabstractt/hcf-core/hcf/utils"
	"github.com/df-mc/dragonfly/server/player"
	scoreboard2 "github.com/df-mc/dragonfly/server/player/scoreboard"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/maps"
	"reflect"
	"time"
)

var (
	profiles = make(map[string]*Profile)
)

type Profile struct {
	player *player.Player

	xuid string
	name string

	factionId string
	factionRole int
	kills int
	deaths int
	balance int

	logger *logrus.Logger
	joinedAt time.Time

	handlerMethods map[string]reflect.Method
}

func RegisterNewProfile(player *player.Player, logger *logrus.Logger, profileStorage *storage.ProfileStorage) *Profile {
	xuid := ""
	name := ""

	if player != nil {
		xuid = player.XUID()
		name = player.Name()
	}

	joinedBefore := profileStorage != nil

	if !joinedBefore {
		profileStorage = storage.NewProfileStorage(xuid, name, "", 0, 0, 0, 0)
	}

	profile := &Profile{
		player:         player,
		xuid:           xuid,
		name:           name,

		factionId:   profileStorage.FactionId,
		factionRole: profileStorage.FactionRole,
		kills:       profileStorage.Kills,
		deaths:      profileStorage.Deaths,
		balance:     profileStorage.Balance,

		logger:         logger,
		joinedAt: time.Now(),
		handlerMethods: make(map[string]reflect.Method),
	}

	if !joinedBefore {
		profile.Save(profileStorage, false)
	}

	profiles[xuid] = profile

	if player != nil {
		//player.Handle(NewPlayerHandler(profile))
	}

	return profile
}

func GetIfLoaded(xuid string) *Profile {
	return profiles[xuid]
}

func FlushProfile(xuid string) {
	profile := GetIfLoaded(xuid)
	if profile == nil {
		return
	}

	profile.kills += 5

	profile.Save(nil, true)

	profile.Logger().Info(profile.Name() + " was flushed successfully!")

	delete(profiles, xuid)
}

func All() []*Profile {
	return maps.Values(profiles)
}

func Close() {
	for xuid := range profiles {
		FlushProfile(xuid)
	}
}

func (profile Profile) UpdateScoreboard()  {
	scoreboard := scoreboard2.New(text.Colourf("<green><bold>HCF"))

	_, err := scoreboard.WriteString(fmt.Sprintf(utils.Colour("&e\n&b&lClaim: &r&aSpawn\n&a&lTime: &r&c%v\n&a\n&7mc.serverhcf.net"), fmt.Sprintf("%.1f", time.Since(profile.joinedAt).Seconds())))
	if err != nil {
		return
	}

	scoreboard.RemovePadding()
	profile.Player().SendScoreboard(scoreboard)
}

func (profile Profile) Logger() *logrus.Logger {
	return profile.logger
}

func (profile Profile) Player() *player.Player {
	return profile.player
}

func (profile Profile) XUID() string {
	return profile.xuid
}

func (profile Profile) Name() string {
	return profile.name
}

func (profile Profile) FactionId() (string, bool) {
	return profile.factionId, len(profile.factionId) > 0
}

// SetFactionId To allow use nil values
func (profile Profile) SetFactionId(factionId string) {
	profile.factionId = factionId
}

func (profile Profile) FactionRole() int {
	return profile.factionRole
}

func (profile Profile) SetFactionRole(factionRole int) {
	profile.factionRole = factionRole
}

func (profile Profile) Kills() int {
	return profile.kills
}

func (profile Profile) SetKills(kills int) {
	profile.kills = kills
}

func (profile Profile) Deaths() int {
	return profile.deaths
}

func (profile Profile) SetDeaths(deaths int) {
	profile.deaths = deaths
}

func (profile Profile) Balance() int {
	return profile.balance
}

func (profile Profile) SetBalance(balance int) {
	profile.balance = balance
}

func (profile Profile) Save(profileStorage *storage.ProfileStorage, joinedBefore bool) {
	if profileStorage == nil {
		profileStorage = storage.NewProfileStorage(
			profile.xuid,
			profile.name,
			profile.factionId,
			profile.factionRole,
			profile.kills,
			profile.deaths,
			profile.balance,
		)
	}

	// Execute this on other Thread to prevent lag spike on the Main thread!
	go datasource.GetCurrentDataSource().SaveProfileStorage(*profileStorage, joinedBefore)
}
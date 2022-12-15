package profile

import (
	"fmt"
	"github.com/aabstractt/hcf-core/hcf/datasource"
	"github.com/aabstractt/hcf-core/hcf/profile/storage"
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

	factionId *string
	factionRole int
	kills int
	deaths int
	balance int

	logger *logrus.Logger
	joinedAt time.Time

	handlerMethods map[string]reflect.Method
}

func RegisterNewProfile(player *player.Player, logger *logrus.Logger, profileData *storage.ProfileStorage) *Profile {
	xuid := ""
	name := ""

	if player != nil {
		xuid = player.XUID()
		name = player.Name()
	}

	joinedBefore := profileData != nil

	if !joinedBefore {
		profileData = storage.NewProfileStorage(xuid, name, nil, 0, 0, 0, 0)
	}

	profile := &Profile{
		player:         player,
		xuid:           xuid,
		name:           name,

		factionId: profileData.FactionId(),
		factionRole: profileData.FactionRole(),
		kills: profileData.Kills(),
		deaths: profileData.Deaths(),
		balance: profileData.Balance(),

		logger:         logger,
		joinedAt: time.Now(),
		handlerMethods: make(map[string]reflect.Method),
	}
	
	if !joinedBefore {
		profile.PushDataSource(profileData)
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

	profile.PushDataSource(nil)

	profile.Logger().Info(profile.GetName() + " was flushed successfully!")

	delete(profiles, xuid)
}

func All() []*Profile {
	return maps.Values(profiles)
}

func Close() {
	for _, profile := range All() {
		profile.PushDataSource(nil)

		delete(profiles, profile.GetXuid())
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

func (profile Profile) GetXuid() string {
	return profile.xuid
}

func (profile Profile) GetName() string {
	return profile.name
}

func (profile Profile) GetFactionId() string {
	return *profile.factionId // To disallow change the value i think x d
}

// SetFactionId To allow use nil values
func (profile Profile) SetFactionId(factionId *string) {
	profile.factionId = factionId
}

func (profile Profile) GetFactionRole() int {
	return profile.factionRole
}

func (profile Profile) SetFactionRole(factionRole int) {
	profile.factionRole = factionRole
}

func (profile Profile) GetKills() int {
	return profile.kills
}

func (profile Profile) SetKills(kills int) {
	profile.kills = kills
}

func (profile Profile) GetDeaths() int {
	return profile.deaths
}

func (profile Profile) SetDeaths(deaths int) {
	profile.deaths = deaths
}

func (profile Profile) GetBalance() int {
	return profile.balance
}

func (profile Profile) SetBalance(balance int) {
	profile.balance = balance
}

func (profile Profile) PushDataSource(profileStorage *storage.ProfileStorage) {
	fmt.Printf("Pushing " + profile.name + " into " + datasource.GetCurrentDataSource().GetName())

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
	go datasource.GetCurrentDataSource().SaveProfileStorage(*profileStorage)
}
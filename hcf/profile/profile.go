package profile

import (
	"github.com/aabstractt/hcf-core/hcf/datasource"
	"github.com/aabstractt/hcf-core/hcf/profile/storage"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/sirupsen/logrus"
	"reflect"
	"sync"
)

var (
	profiles sync.Map
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
		handlerMethods: make(map[string]reflect.Method),
	}
	profile.PushDataSource(profileData)

	profiles.Store(xuid, profile)

	if player != nil {
		//player.Handle(NewPlayerHandler(profile))
	}

	return profile
}

func GetIfLoaded(xuid string) *Profile {
	profileVar, exists := profiles.Load(xuid)

	if exists {
		return profileVar.(*Profile)
	}

	return nil
}

func FlushProfile(xuid string) {
	profile := GetIfLoaded(xuid)
	if profile == nil {
		return
	}

	profile.Logger().Info(profile.GetName() + " was flushed successfully!")

	profiles.Delete(xuid)
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
	go datasource.GetCurrentDataSource().PushProfileStorage(*profileStorage)
}
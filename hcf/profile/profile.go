package profile

import (
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

func RegisterNewProfile(player *player.Player, logger *logrus.Logger, profileData *ProfileData) *Profile {
	xuid := ""
	name := ""

	if player != nil {
		xuid = player.XUID()
		name = player.Name()
	}

	joinedBefore := profileData != nil

	if !joinedBefore {
		profileData = NewProfileData(xuid, name, nil, 0, 0, 0, 0)
	}

	profile := &Profile{
		player:         player,
		xuid:           xuid,
		name:           name,

		factionId: profileData.factionId,
		factionRole: profileData.factionRole,
		kills: profileData.kills,
		deaths: profileData.deaths,
		balance: profileData.balance,

		logger:         logger,
		handlerMethods: make(map[string]reflect.Method),
	}
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
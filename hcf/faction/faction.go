package faction

import (
	"github.com/aabstractt/hcf-core/hcf/datasource"
	"github.com/aabstractt/hcf-core/hcf/datasource/storage"
	"github.com/aabstractt/hcf-core/hcf/profile"
	"github.com/aabstractt/hcf-core/hcf/utils"
	"strings"
)

var (
	factions = map[string]*Faction{}
	factionsId = map[string]string{}
)

type Faction struct {

	id string
	name string
	leaderXuid string

	deathsUntilRaidable float32
	regenCooldown int
	lastDtrUpdate int
	balance int
	points int
}

func (faction Faction) Id() string {
	return faction.id
}

func (faction Faction) Name() string {
	return faction.name
}

func (faction Faction) ForceSave() {
	// TODO: Save the faction storage into the data source provider using async method
	go datasource.GetCurrentDataSource().SaveFactionStorage(storage.FactionStorage{})
}

func NewFaction(id string, name string, leaderXuid string, deathsUntilRaidable float32, regenCooldown int, lastDtrUpdate int, balance int, points int) *Faction {
	return &Faction{
		id:                  id,
		name:                name,
		leaderXuid:          leaderXuid,
		deathsUntilRaidable: deathsUntilRaidable,
		regenCooldown:       regenCooldown,
		lastDtrUpdate:       lastDtrUpdate,
		balance:             balance,
		points:              points,
	}
}

func GetPlayerFaction(name string) *Faction {
	player, exists := utils.Server().PlayerByName(name)

	if player == nil || !exists {
		return nil
	}

	return GetProfileFaction(profile.GetIfLoaded(player.XUID()))
}

func GetProfileFaction(pf *profile.Profile) *Faction {
	if pf == nil {
		return nil
	}

	factionId, exists := pf.FactionId()
	if !exists {
		return nil
	}

	return factions[factionId]
}

func GetFaction(factionName string) *Faction {
	factionId, present := factionsId[strings.ToLower(factionName)]

	if present {
		return factions[factionId]
	}

	return factions[factionName]
}

func RegisterFactionsStored() {
	for _, factionStorage := range datasource.GetCurrentDataSource().LoadFactionsStored() {
		RegisterNewFaction(&Faction{
			id: factionStorage.Id(),
			name: factionStorage.Name(),
		})
	}
}

func RegisterNewFaction(f *Faction) {
	factions[f.Id()] = f
	factionsId[strings.ToLower(f.Name())] = f.Id()
}

func JoinFaction(pf *profile.Profile, f *Faction, factionRole int) {
	pf.SetFactionId(f.Id())
	pf.SetFactionRole(factionRole)

	pf.Save(nil)

	// TODO: Register faction member
}
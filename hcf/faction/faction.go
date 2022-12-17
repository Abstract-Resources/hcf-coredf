package faction

import (
	"github.com/aabstractt/hcf-core/hcf"
	"github.com/aabstractt/hcf-core/hcf/datasource"
	"github.com/aabstractt/hcf-core/hcf/datasource/storage"
	"github.com/aabstractt/hcf-core/hcf/profile"
)

var (
	factions = map[string]*Faction{}
	factionsId = map[string]string{}
)

type Faction struct {

	id string
	name string
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

func GetPlayerFaction(name string) *Faction {
	player, exists := hcf.Server().PlayerByName(name)

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
	factionId, present := factionsId[factionName]

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
	factionsId[f.Name()] = f.Id()
}
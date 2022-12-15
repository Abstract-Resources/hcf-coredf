package storage

type ProfileStorage struct {

	xuid string
	name string

	factionId *string
	factionRole int
	kills int
	deaths int
	balance int
}

func (profileStorage ProfileStorage) FactionId() *string {
	return profileStorage.factionId
}

func (profileStorage ProfileStorage) FactionRole() int {
	return profileStorage.factionRole
}

func (profileStorage ProfileStorage) Kills() int {
	return profileStorage.kills
}

func (profileStorage ProfileStorage) Deaths() int {
	return profileStorage.deaths
}

func (profileStorage ProfileStorage) Balance() int {
	return profileStorage.balance
}

func NewProfileStorage(xuid string, name string, factionId *string, factionRole int, kills int, deaths int, balance int) *ProfileStorage {
	return &ProfileStorage{
		xuid: xuid,
		name: name,

		factionId: factionId,
		factionRole: factionRole,
		kills: kills,
		deaths: deaths,
		balance: balance,
	}
}
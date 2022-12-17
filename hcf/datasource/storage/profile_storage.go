package storage

type ProfileStorage struct {
	Xuid string		`json:"xuid"`
	Name string 	`json:"name"`

	FactionId   string	`json:"faction_id" bson:"faction_id,omitempty"`
	FactionRole int		`json:"faction_role"`
	Kills       int
	Deaths int
	Balance int
}

func NewProfileStorage(xuid string, name string, factionId string, factionRole int, kills int, deaths int, balance int) *ProfileStorage {
	return &ProfileStorage{
		Xuid: xuid,
		Name: name,

		FactionId:   factionId,
		FactionRole: factionRole,
		Kills:       kills,
		Deaths:      deaths,
		Balance:     balance,
	}
}
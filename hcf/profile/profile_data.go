package profile

type ProfileData struct {

	xuid string
	name string

	factionId *string
	factionRole int
	kills int
	deaths int
	balance int
}

func NewProfileData(xuid string, name string, factionId *string, factionRole int, kills int, deaths int, balance int) *ProfileData {
	return &ProfileData{
		xuid: xuid,
		name: name,

		factionId: factionId,
		factionRole: factionRole,
		kills: kills,
		deaths: deaths,
		balance: balance,
	}
}
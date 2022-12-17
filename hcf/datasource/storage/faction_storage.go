package storage

type FactionStorage struct {

	id string
	name string
}

func (factionStorage FactionStorage) Id() string {
	return factionStorage.id
}

func (factionStorage FactionStorage) Name() string {
	return factionStorage.name
}
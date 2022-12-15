package datasource

import (
	"github.com/aabstractt/hcf-core/hcf/profile/storage"
)

type MongoDBDataSource struct {
	DataSource

	address string
	username string
	password string
	dbname string
}

func (dataSource MongoDBDataSource) GetName() string {
	return "MongoDB"
}

func (dataSource MongoDBDataSource) PushProfileStorage(profileStorage storage.ProfileStorage) {

}

func (dataSource MongoDBDataSource) FetchProfileStorage(xuid string, name string) *storage.ProfileStorage {
	return nil
}

func NewMongoDB(address string, username string, password string, dbname string) *MongoDBDataSource {
	return &MongoDBDataSource{
		address:    address,
		username:   username,
		password:   password,
		dbname:     dbname,
	}
}
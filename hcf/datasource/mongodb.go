package datasource

import "github.com/aabstractt/hcf-core/hcf/profile"

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

func (dataSource MongoDBDataSource) StoreProfile(profileData profile.ProfileData) {

}

func (dataSource MongoDBDataSource) FetchProfile(xuid string, name string) *profile.ProfileData {
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
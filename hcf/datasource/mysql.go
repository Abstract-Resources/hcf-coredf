package datasource

import "github.com/aabstractt/hcf-core/hcf/profile"

type MySQLDataSource struct {
	DataSource

	address string
	username string
	password string
	dbname string
}

func (dataSource MySQLDataSource) GetName() string {
	return "MySQL"
}

func (dataSource MySQLDataSource) StoreProfile(profileData profile.ProfileData) {

}

func (dataSource MySQLDataSource) FetchProfile(xuid string, name string) *profile.ProfileData {
	return nil
}

func NewMySQL(address string, username string, password string, dbname string) *MySQLDataSource {
	return &MySQLDataSource{
		address:    address,
		username:   username,
		password:   password,
		dbname:     dbname,
	}
}
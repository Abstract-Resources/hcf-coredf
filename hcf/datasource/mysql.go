package datasource

import (
	"github.com/aabstractt/hcf-core/hcf/profile/storage"
)

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

func (dataSource MySQLDataSource) PushProfileStorage(profileStorage storage.ProfileStorage) {

}

func (dataSource MySQLDataSource) FetchProfileStorage(xuid string, name string) *storage.ProfileStorage {
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
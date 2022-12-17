package datasource

import (
	"github.com/aabstractt/hcf-core/hcf/datasource/storage"
	"github.com/sirupsen/logrus"
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

func (dataSource MySQLDataSource) Initialize(log *logrus.Logger) bool {
	return true
}

func (dataSource MySQLDataSource) SaveProfileStorage(profileStorage storage.ProfileStorage) {

}

func (dataSource MySQLDataSource) LoadProfileStorage(xuid string) *storage.ProfileStorage {
	return nil
}

func (dataSource MySQLDataSource) SaveFactionStorage(factionStorage storage.FactionStorage) {

}

func (dataSource MySQLDataSource) LoadFactionStorage(factionId string) *storage.FactionStorage {
	return nil
}

func (dataSource MySQLDataSource) LoadFactionsStored() []storage.FactionStorage {
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
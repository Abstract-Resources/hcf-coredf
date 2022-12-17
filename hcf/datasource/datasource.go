package datasource

import (
	"github.com/aabstractt/hcf-core/hcf/config"
	"github.com/aabstractt/hcf-core/hcf/datasource/storage"
	"github.com/sirupsen/logrus"
	"strings"
)

var dataSource DataSource = nil

type DataSource interface {

	GetName() string

	Initialize(log *logrus.Logger) bool

	SaveProfileStorage(profileStorage storage.ProfileStorage)

	LoadProfileStorage(xuid string) *storage.ProfileStorage

	SaveFactionStorage(factionStorage storage.FactionStorage)

	LoadFactionStorage(factionId string) *storage.FactionStorage

	LoadFactionsStored() []storage.FactionStorage
}

func GetCurrentDataSource() DataSource {
	if dataSource == nil {
		panic("Cannot get data source without initialize that")
	}

	return dataSource
}

func NewDataSource(log *logrus.Logger) {
	provider := config.DefaultConfig().Provider

	if strings.ToLower(provider.ProviderName) == "mongodb" {
		dataSource = NewMongoDB(provider.Address, provider.Username, provider.Password, provider.Dbname)
	} else if strings.ToLower(provider.ProviderName) == "mysql" {
		dataSource = NewMySQL(provider.Address, provider.Username, provider.Password, provider.Dbname)
	} else {
		panic("Please provide a valid type Data Source")

		return
	}

	if !dataSource.Initialize(log) {
		log.Fatal("An error occurred while triad initialize the database!")

		return
	}

	log.Info("Successfully initialized '" + dataSource.GetName() + "' as database provider")
}
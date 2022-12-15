package datasource

import (
	"github.com/aabstractt/hcf-core/hcf/config"
	"github.com/aabstractt/hcf-core/hcf/profile/storage"
	"strings"
)

var (
	dataSource DataSource = nil
)

type DataSource interface {

	GetName() string

	PushProfileStorage(profileData storage.ProfileStorage)

	FetchProfileStorage(xuid string, name string) *storage.ProfileStorage
}

func GetCurrentDataSource() DataSource {
	if dataSource == nil {
		panic("Cannot get data source without initialize that")
	}

	return dataSource
}

func NewDataSource(conf *config.ServerConfig) {
	if conf == nil {
		return
	}

	provider := conf.Provider

	if strings.ToLower(provider.ProviderName) == "mongodb" {
		dataSource = NewMongoDB(provider.Address, provider.Username, provider.Password, provider.Dbname)
	} else if strings.ToLower(provider.ProviderName) == "mysql" {
		dataSource = NewMySQL(provider.Address, provider.Username, provider.Password, provider.Dbname)
	} else {
		panic("Please provide a valid type Data Source")
	}
}
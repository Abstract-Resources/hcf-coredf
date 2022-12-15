package datasource

import (
	"github.com/aabstractt/hcf-core/hcf/profile/storage"
)

type DataSource interface {

	GetName() string

	PushProfileStorage(profileData storage.ProfileStorage)

	FetchProfileStorage(xuid string, name string) *storage.ProfileStorage
}
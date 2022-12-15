package datasource

import (
	"github.com/aabstractt/hcf-core/hcf/profile"
)

type DataSource interface {

	GetName() string

	StoreProfile(profileData profile.ProfileData)

	FetchProfile(xuid string, name string) *profile.ProfileData
}
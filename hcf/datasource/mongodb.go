package datasource

import (
	"context"
	"fmt"
	"github.com/aabstractt/hcf-core/hcf/datasource/storage"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	logger *logrus.Logger

	profilesCollection *mongo.Collection
)

type MongoDBDataSource struct {
	DataSource

	address string
	username string
	password string
	dbname string
}

func (dataSource MongoDBDataSource) Initialize(log *logrus.Logger) bool {
	logger = log

	//clientOptions := options.Client().ApplyURI("mongodb+srv://xavier:xavier123@@cluster0.y05ycq4.mongodb.net/test")
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@%s", dataSource.username, dataSource.password, dataSource.address))

	localClient, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		logger.Fatal(err)

		return false
	}

	err = localClient.Ping(context.TODO(), nil)

	if err != nil {
		logger.Fatal(err)

		return false
	}

	client = localClient

	profilesCollection = client.Database(dataSource.dbname).Collection("profiles")

	return true
}

func (dataSource MongoDBDataSource) GetName() string {
	return "MongoDB"
}

func (dataSource MongoDBDataSource) SaveProfileStorage(profileStorage storage.ProfileStorage) {
	if profilesCollection == nil {
		return
	}

	_, err := profilesCollection.UpdateOne(
		context.TODO(),
		bson.D{{"xuid", profileStorage.Xuid}},
		bson.D{{"$set", profileStorage}},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		logger.Fatal("[S/Profiles] ", err)
	}
}

func (dataSource MongoDBDataSource) LoadProfileStorage(xuid string) *storage.ProfileStorage {
	if profilesCollection == nil {
		return nil
	}

	raw, err := profilesCollection.FindOne(context.TODO(), bson.D{{"xuid", xuid}}).DecodeBytes()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}

		logger.Fatal("[L/Profiles] ", err)

		return nil
	}

	return storage.NewProfileStorage(
		xuid,
		raw.Lookup("name").String(),
		raw.Lookup("factionid").String(),
		int(raw.Lookup("factionrole").Int32()),
		int(raw.Lookup("kills").Int32()),
		int(raw.Lookup("deaths").Int32()),
		int(raw.Lookup("balance").Int32()),
	)
}

func (dataSource MongoDBDataSource) SaveFactionStorage(factionStorage storage.FactionStorage) {

}

func (dataSource MongoDBDataSource) LoadFactionStorage(factionId string) *storage.FactionStorage {
	return nil
}

func (dataSource MongoDBDataSource) LoadFactionsStored() []storage.FactionStorage {
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
package ryaas

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertLapangan(db string, lapangan Lapangan) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("lapangan").InsertOne(context.TODO(), lapangan)
	if err != nil {
		fmt.Printf("InsertLapangan: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertKategori(db string, kategori Kategori) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("kategori").InsertOne(context.TODO(), kategori)
	if err != nil {
		fmt.Printf("InsertKategori: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertKontak(db string, kontak Kontak) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("kontak").InsertOne(context.TODO(), kontak)
	if err != nil {
		fmt.Printf("InsertKontak: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertBank(db string, bank Bank) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("bank").InsertOne(context.TODO(), bank)
	if err != nil {
		fmt.Printf("InsertBank: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertDiskon(db string, diskon Diskon) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("diskon").InsertOne(context.TODO(), diskon)
	if err != nil {
		fmt.Printf("InsertDiskon: %v\n", err)
	}
	return insertResult.InsertedID
}

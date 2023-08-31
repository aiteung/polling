package polling

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aiteung/module/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetKandidat(db *mongo.Database) (data []Kandidat, err error) {
	kandidat := db.Collection("kandidat")
	filter := bson.M{} // Empty filter to get all data
	cur, err := kandidat.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.Background(), &data); err != nil {
		return nil, err
	}

	if len(data) < 1 {
		return nil, errors.New("Data Tidak Ada")
	}
	return data, nil
}

func GetNamaAndNomorKandidat(db *mongo.Database) (data []KandidatInfo, err error) {
	kandidat := db.Collection("kandidat")
	filter := bson.M{} // Empty filter to get all data

	// Proyeksi untuk memilih hanya field yang diperlukan
	projection := bson.M{"nama_kandidat": 1, "nomor_kandidat": 2}

	cur, err := kandidat.Find(context.TODO(), filter, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.Background(), &data); err != nil {
		return nil, err
	}

	if len(data) < 1 {
		return nil, errors.New("Data Tidak Ada")
	}
	return data, nil
}

func GetAnggota(db *mongo.Database) (data []Anggota, err error) {
	kandidat := db.Collection("anggota")
	filter := bson.M{} // Empty filter to get all data
	cur, err := kandidat.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.Background(), &data); err != nil {
		return nil, err
	}

	if len(data) < 1 {
		return nil, errors.New("Data Tidak Ada")
	}
	return data, nil
}

func GetAnggotaFromPhoneNumber(mongoconn *mongo.Database, phone_number string) (anggota Anggota) {
	data := mongoconn.Collection("anggota")
	filter := bson.M{"phone_number": phone_number}
	err := data.FindOne(context.TODO(), filter).Decode(&anggota)
	if err != nil {
		fmt.Printf("GetAnggotaFromPhoneNumber: %v\n", err)
	}
	return anggota
}

func GetKandidatFromPhoneNumber(mongoconn *mongo.Database, phone_number string) (kandidat Kandidat) {
	data := mongoconn.Collection("kandidat")
	filter := bson.M{"phone_number": phone_number}
	err := data.FindOne(context.TODO(), filter).Decode(&kandidat)
	if err != nil {
		fmt.Printf("GetAnggotaFromPhoneNumber: %v\n", err)
	}
	return kandidat
}

func GetPollingFromPhoneNumber(mongoconn *mongo.Database, phone_number string) (polling Polling) {
	poll := mongoconn.Collection("polling")
	filter := bson.M{"phone_number": phone_number}
	err := poll.FindOne(context.TODO(), filter).Decode(&poll)
	if err != nil {
		fmt.Printf("GetPollingTodayFromPhoneNumber: %v\n", err)
	}
	return polling
}

func InsertPolling(Pesan model.IteungMessage, Keterangan string, mongoconn *mongo.Database) (InsertedID interface{}) {
	insertResult, err := mongoconn.Collection("presensi").InsertOne(context.TODO(), fillStructPolling(Pesan, Keterangan, mongoconn))
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func ConvertTimestampToJkt(waktu time.Time) time.Time {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return waktu.In(loc)
}

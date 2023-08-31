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

func GetKandidatByNomorUrut(mongoconn *mongo.Database, NomorKandidat string) (selectedKandidat Kandidat) {
	kandidatCollection := mongoconn.Collection("kandidat")
	filter := bson.M{"nomor_kandidat": NomorKandidat}

	err := kandidatCollection.FindOne(context.TODO(), filter).Decode(&selectedKandidat)
	if err != nil {
		fmt.Printf("GetKandidatByNomorUrut: %v\n", err)
	}

	return selectedKandidat
}

func GetOnePresensi(NomorKandidat, db *mongo.Database) (data Kandidat) {
	kandidat := db.Collection("kandidat")
	filter := bson.M{"nomor_kandidat": NomorKandidat}
	err := kandidat.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("Data Tidak Ada : %v\n", err)
	}
	return data
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

func InsertPolling(Pesan model.IteungMessage, Keterangan string, NomorKandidat string, mongoconn *mongo.Database) (InsertedID interface{}) {
	insertResult, err := mongoconn.Collection("polling").InsertOne(context.TODO(), fillStructPolling(Pesan, Keterangan, NomorKandidat, mongoconn))
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func ConvertTimestampToJkt(waktu time.Time) time.Time {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return waktu.In(loc)
}

func HandleUserInput(Pesan model.IteungMessage, mongoconn *mongo.Database, userInput int) (reply string) {
	switch userInput {
	case 1, 2:
		kandidat := GetKandidatByIndex(mongoconn, userInput-1) // Convert user input to array index
		anggota := GetAnggotaFromPhoneNumber(mongoconn, Pesan.Phone_number)
		id := InsertPolling(Pesan, "polling", kandidat.NomorKandidat, mongoconn)
		reply = MessagePolling(anggota, kandidat, id)
	default:
		reply = "Nomor urut kandidat tidak valid. Mohon pilih 1 atau 2."
	}
	return
}
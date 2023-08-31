package polling

import (
	"reflect"
	"time"

	"github.com/aiteung/module/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func Handler(Pesan model.IteungMessage, mongoconn *mongo.Database) (reply string) {
	kandidat, _ := GetKandidat(mongoconn)
	if kandidat != nil {
		reply = PollingHandler(Pesan, mongoconn)
	} else {
		reply = NoPollingHandler(Pesan, mongoconn)
	}
	return
}

func NoPollingHandler(Pesan model.IteungMessage, mongoconn *mongo.Database) (reply string) {
	return
}

func PollingHandler(Pesan model.IteungMessage, mongoconn *mongo.Database) (reply string) {
	anggota := GetAnggotaFromPhoneNumber(mongoconn, Pesan.Phone_number)
	alreadypolling := GetPollingFromPhoneNumber(mongoconn, Pesan.Phone_number)
	kandidat := GetKandidatFromPhoneNumber(mongoconn, Pesan.Phone_number)
	if !reflect.ValueOf(alreadypolling).IsZero() {
		reply = MessageSudahPolling(anggota, kandidat)
	} else {
		id := InsertPolling(Pesan, "polling", mongoconn)
		reply = MessagePolling(anggota, kandidat, id)
	}
	return
}

func fillStructPolling(Pesan model.IteungMessage, Keterangan string, mongoconn *mongo.Database) (polling Polling) {
	polling.PhoneNumber = Pesan.Phone_number
	polling.Keterangan = Keterangan
	polling.Datetime = ConvertTimestampToJkt(time.Now())
	polling.Keterangan = Keterangan
	polling.Anggota = GetAnggotaFromPhoneNumber(mongoconn, Pesan.Phone_number)
	return polling
}

package polling

import (
	"reflect"
	"time"

	"github.com/aiteung/module/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func Handler(Pesan model.IteungMessage, mongoconn *mongo.Database) (reply string) {
	reply = ListKandidatMessage(mongoconn)
	return
}

func PollingHandler(Pesan model.IteungMessage, mongoconn *mongo.Database, selectedCandidate int) (reply string) {
	anggota := GetAnggotaFromPhoneNumber(mongoconn, Pesan.Phone_number)
	alreadypolling := GetPollingFromPhoneNumber(mongoconn, Pesan.Phone_number)

	if !reflect.ValueOf(alreadypolling).IsZero() {
		reply = "Anda sudah melakukan polling sebelumnya."
	} else {
		// Memanggil HandleUserInput dengan nomor kandidat yang dipilih
		reply = HandleUserInput(Pesan, mongoconn, selectedCandidate+1) // +1 karena indeks dimulai dari 0

		// Hanya jika pemanggilan HandleUserInput berhasil, baru lanjutkan ke InsertPolling
		if reply == "Terima kasih atas polling Anda!" {
			id := InsertPolling(Pesan, "polling", GetKandidatByIndex(mongoconn, selectedCandidate).NomorKandidat, mongoconn)
			selectedKandidat := GetKandidatByIndex(mongoconn, selectedCandidate)
			reply = MessagePolling(anggota, selectedKandidat, id)
		}
	}
	return
}

func GetKandidatByIndex(mongoconn *mongo.Database, index int) (kandidat Kandidat) {
	kandidatInfo, err := GetNamaAndNomorKandidat(mongoconn)
	if err != nil || index < 0 || index >= len(kandidatInfo) {
		// Handle error
		return Kandidat{} // or an appropriate default value
	}
	return Kandidat{NamaKandidat: kandidatInfo[index].NamaKandidat, NomorKandidat: kandidatInfo[index].NomorKandidat}
}

func fillStructPolling(Pesan model.IteungMessage, Keterangan string, NomorKandidat string, mongoconn *mongo.Database) (polling Polling) {
	polling.PhoneNumber = Pesan.Phone_number
	polling.Keterangan = Keterangan
	polling.Datetime = ConvertTimestampToJkt(time.Now())
	polling.Keterangan = Keterangan
	polling.Anggota = GetAnggotaFromPhoneNumber(mongoconn, Pesan.Phone_number)

	// Get data kandidat based on the provided candidate number (NomorKandidat)
	kandidatData := GetKandidatByNomorUrut(mongoconn, NomorKandidat) // Replace this with your actual function to get candidate data

	// Fill the candidate information in the polling struct
	polling.Kandidat = kandidatData

	return polling
}

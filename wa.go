package polling

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

func MessagePolling(anggota Anggota, kandidat Kandidat, id interface{}) string {
	msg := "*Polling*\n"
	msg = msg + "Hai kak _*" + anggota.NamaAnggota + "*_,\ndengan nomor telepon *" + anggota.PhoneNumber + "*, \nmakasih ya sudah melakukan polling ketua IAI\nKandidat yang kakak pilih: _*" + kandidat.NamaKandidat + "*_\nNomor urut kandidat : *" + kandidat.NomorKandidat + "*\nTerima kasih kakak...\n"
	msg = msg + fmt.Sprintf("ID Polling : %v", id)
	return msg
}

func MessageSudahPolling(anggota Anggota, kandidat Kandidat) string {
	msg := "*Polling*\n"
	msg = msg + "Hai kak _*" + anggota.NamaAnggota + "*_,\ndengan nomor telepon *" + anggota.PhoneNumber + "*, \nKakak sudah melakukan polling\nKandidat yang kakak pilih: _*" + kandidat.NamaKandidat + "*_\nTerima kasih kakak...\n"
	return msg
}

func ListKandidatMessage(mongoconn *mongo.Database) string {
	kandidatInfo, err := GetNamaAndNomorKandidat(mongoconn)
	if err != nil {
		// Handle error
		return "Maaf, terjadi kesalahan."
	}

	msg := "Silahkan pilih kandidat yang menurut kakak memenuhi untuk menjadi ketua:\n"
	for idx, kandidat := range kandidatInfo {
		msg += fmt.Sprintf("%d. %s\n", idx+1, kandidat.NamaKandidat)
	}

	msg += "Gunakan hak suara anda sebaik-baiknya, terima kasih."
	return msg
}

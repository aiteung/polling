package polling

import (
	"fmt"
	"strings"

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
		return "Maaf, terjadi kesalahan saat mengambil data kandidat."
	}

	var messageBuilder strings.Builder
	messageBuilder.WriteString("Pilih kandidat dengan mengklik salah satu tautan di bawah ini:\n\n")

	for idx, kandidat := range kandidatInfo {
		link := fmt.Sprintf("[Kandidat %s](https://api.whatsapp.com/send?phone=628112000279&text=Iteung+pilih+calon+kandidat+%d)", kandidat.NamaKandidat, idx+1)
		messageBuilder.WriteString(fmt.Sprintf("%d. %s\n", idx+1, link))
	}

	return messageBuilder.String()
}

func MintaQRCode() string {
	msg := "Silahkan klik link dibawah ini\n"
	return msg + "https://getqr.github.io/#https://wa.me/628112000279?text=Iteung%20minta%20kandidat%20ketua%20iai"
}

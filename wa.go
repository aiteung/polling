package polling

import (
	"fmt"
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

package polling

import (
	// "fmt"
	"fmt"
	"os"
	"testing"

	"github.com/aiteung/atdb"
	"github.com/aiteung/module/model"
)

var MongoInfo = atdb.DBInfo{
	DBString: os.Getenv("MONGOSTRING"),
	DBName:   "iai",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

var Pesan = model.IteungMessage{
	Phone_number: "6289522910966",
}

// func TestPolling(t *testing.T) {
// 	cek, err := GetNamaAndNomorKandidat(MongoConn) // Menggunakan MongoConn yang telah Anda inisialisasi sebelumnya
// 	if err != nil {
// 		t.Errorf("Error while getting kandidat: %v", err)
// 		return
// 	}

// 	fmt.Println(cek)
// }

// func TestPolling(t *testing.T) {
// 	// Panggil fungsi pengujian
// 	result := ListKandidatMessage(MongoConn)
// 	fmt.Println(result)
// }

// func TestPolling(t *testing.T) {
// 	// Panggil fungsi pengujian
// 	result := Handler(Pesan, MongoConn)
// 	fmt.Println(result)
// }

// func TestPolling(t *testing.T) {
// 	// Simulasikan pemilihan nomor urut
// 	userInput := 3 // Ganti dengan nomor urut yang sesuai (1 atau 2)

// 	// Panggil fungsi untuk memproses pemilihan
// 	reply := HandleUserInput(Pesan, MongoConn, userInput)

// 	// Buat asersi untuk memeriksa hasil balasan
// 	expectedReply := "..." // Ganti dengan balasan yang diharapkan
// 	if reply != expectedReply {
// 		t.Errorf("Expected reply: %s, Got reply: %s", expectedReply, reply)
// 	}
// }

func TestPolling(t *testing.T) {
	// Panggil fungsi untuk memproses pemilihan
	reply := ListKandidatMessage(MongoConn)
	fmt.Println(reply)

	// Buat asersi untuk memeriksa hasil balasan
	// expectedReply := "..." // Ganti dengan balasan yang diharapkan
	// if reply != expectedReply {
	// 	t.Errorf("Expected reply: %s, Got reply: %s", expectedReply, reply)
	// }
}

// func TestPolling(t *testing.T) {
// 	cek, err := GetAnggota(MongoConn) // Menggunakan MongoConn yang telah Anda inisialisasi sebelumnya
// 	if err != nil {
// 		t.Errorf("Error while getting kandidat: %v", err)
// 		return
// 	}

// 	fmt.Println(cek)
// }

// func TestPolling(t *testing.T) {
// 	phonnumber := "6289522910966"
// 	cek := GetAnggotaFromPhoneNumber(MongoConn, phonnumber)
// 	fmt.Println(cek)
// }

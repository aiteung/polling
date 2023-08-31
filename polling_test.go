package polling

import (
	"fmt"
	"os"
	"testing"

	"github.com/aiteung/atdb"
)

var MongoInfo = atdb.DBInfo{
	DBString: os.Getenv("MONGOSTRING"),
	DBName:   "iai",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func TestPolling(t *testing.T) {
	cek, err := GetNamaAndNomorKandidat(MongoConn) // Menggunakan MongoConn yang telah Anda inisialisasi sebelumnya
	if err != nil {
		t.Errorf("Error while getting kandidat: %v", err)
		return
	}

	fmt.Println(cek)
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

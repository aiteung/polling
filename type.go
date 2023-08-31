package polling

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Kandidat struct {
	Id            primitive.ObjectID `bson:"_id" json:"_id"`
	NamaKandidat  string             `bson:"nama_kandidat" json:"nama_kandidat"`
	NomorKandidat string             `bson:"nomor_kandidat" json:"nomor_kandidat"`
	PhoneNumber   string             `bson:"phone_number" json:"phone_number"`
}

type KandidatInfo struct {
	NamaKandidat  string `bson:"nama_kandidat" json:"nama_kandidat"`
	NomorKandidat string `bson:"nomor_kandidat" json:"nomor_kandidat"`
}

type Anggota struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	NamaAnggota string             `bson:"nama_anggota" json:"nama_anggota"`
	PhoneNumber string             `bson:"phone_number" json:"phone_number"`
}

type Polling struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Datetime    time.Time          `bson:"datetime" json:"datetime"`
	PhoneNumber string             `bson:"phone_number" json:"phone_number"`
	Keterangan  string             `bson:"keterangan" json:"keterangan"`
	Kandidat    Kandidat           `bson:"kandidat" json:"kandidat"`
	Anggota     Anggota            `bson:"anggota" json:"anggota"`
}

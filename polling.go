package polling

import (
	"github.com/aiteung/module/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func Handler(Pesan model.IteungMessage, mongoconn *mongo.Database) (reply string) {
	return "Oke"
}

package ryaas

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Lapangan struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama  string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Harga string             `bson:"harga,omitempty" json:"harga,omitempty"`
}

type Kategori struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama     string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Turnamen string             `bson:"turnamen,omitempty" json:"turnamen,omitempty"`
}

type Kontak struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
}

type Bank struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_Bank string             `bson:"nama_bank,omitempty" json:"nama_bank,omitempty"`
	Atas_Nama string             `bson:"atas_nama,omitempty" json:"atas_nama,omitempty"`
}

type Diskon struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Harga string             `bson:"harga,omitempty" json:"harga,omitempty"`
}

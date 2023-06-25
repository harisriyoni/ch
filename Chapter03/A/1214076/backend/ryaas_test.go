package ryaas

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertLapangan(t *testing.T) {
	dbname := "lapangan"
	lapangan := Lapangan{
		ID:    primitive.NewObjectID(),
		Nama:  "Lapangan A",
		Harga: "100000",
	}
	insertedID := InsertLapangan(dbname, lapangan)
	if insertedID == nil {
		t.Error("Failed to insert user")
	}
}

func TestInsertKategori(t *testing.T) {
	dbname := "kategori"
	kategori := Kategori{
		ID:       primitive.NewObjectID(),
		Nama:     "Turnamen",
		Turnamen: "turnamen",
	}
	insertedID := InsertKategori(dbname, kategori)
	if insertedID == nil {
		t.Error("Failed to insert surat")
	}
}

func TestInsertKontak(t *testing.T) {
	dbname := "kontak"
	kontak := Kontak{
		ID:           primitive.NewObjectID(),
		Nama:         "WAWAN",
		Phone_number: "62876686833",
	}
	insertedID := InsertKontak(dbname, kontak)
	if insertedID == nil {
		t.Error("Failed to insert lokasi")
	}
}

func TestInsertBank(t *testing.T) {
	dbname := "bank"
	bank := Bank{
		ID:        primitive.NewObjectID(),
		Nama_Bank: "BNI",
		Atas_Nama: "Microtron",
	}
	insertedID := InsertBank(dbname, bank)
	if insertedID == nil {
		t.Error("Failed to insert bank")
	}
}

func TestInsertDiskon(t *testing.T) {
	dbname := "diskon"
	diskon := Diskon{
		ID:    primitive.NewObjectID(),
		Harga: "35000",
	}
	insertedID := InsertDiskon(dbname, diskon)
	if insertedID == nil {
		t.Error("Failed to insert lokasi")
	}
}

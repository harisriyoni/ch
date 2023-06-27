package iyasbackend

import (
	"fmt"
	"testing"
)

// func TestInsertLapangan(t *testing.T) {
// 	dbname := "lapangan"
// 	lapangan := Lapangan{
// 		ID:    primitive.NewObjectID(),
// 		Nama:  "Lapangan A",
// 		Harga: "100000",
// 	}
// 	insertedID := InsertLapangan(dbname, lapangan)
// 	if insertedID == nil {
// 		t.Error("Failed to insert user")
// 	}
// }

// func TestInsertKategori(t *testing.T) {
// 	dbname := "kategori"
// 	kategori := Kategori{
// 		ID:       primitive.NewObjectID(),
// 		Nama:     "Turnamen",
// 		Turnamen: "turnamen",
// 	}
// 	insertedID := InsertKategori(dbname, kategori)
// 	if insertedID == nil {
// 		t.Error("Failed to insert surat")
// 	}
// }

// func TestInsertKontak(t *testing.T) {
// 	dbname := "kontak"
// 	kontak := Kontak{
// 		ID:           primitive.NewObjectID(),
// 		Nama:         "WAWAN",
// 		Phone_number: "62876686833",
// 	}
// 	insertedID := InsertKontak(dbname, kontak)
// 	if insertedID == nil {
// 		t.Error("Failed to insert lokasi")
// 	}
// }

// func TestInsertBank(t *testing.T) {
// 	dbname := "bank"
// 	bank := Bank{
// 		ID:        primitive.NewObjectID(),
// 		Nama_Bank: "BNI",
// 		Atas_Nama: "Microtron",
// 	}
// 	insertedID := InsertBank(dbname, bank)
// 	if insertedID == nil {
// 		t.Error("Failed to insert bank")
// 	}
// }

// func TestInsertDiskon(t *testing.T) {
// 	dbname := "diskon"
// 	diskon := Diskon{
// 		ID:    primitive.NewObjectID(),
// 		Harga: "35000",
// 	}
// 	insertedID := InsertDiskon(dbname, diskon)
// 	if insertedID == nil {
// 		t.Error("Failed to insert lokasi")
// 	}
// }

func TestGetDataLapangan(t *testing.T) {
	stats := "Lapangan A"
	data := GetDataLapangan(stats)
	fmt.Println(data)
}

func TestGetDataKategori(t *testing.T) {
	stats := "Futsal"
	data := GetDataKategori(stats)
	fmt.Println(data)
}

func TestGetDataKontak(t *testing.T) {
	stats := "35000"
	data := GetDataKontak(stats)
	fmt.Println(data)
}

func TestGetDataBank(t *testing.T) {
	stats := "Microtron"
	data := GetDataBank(stats)
	fmt.Println(data)
}
func TestGetDataDiskon(t *testing.T) {
	stats := "0821234123"
	data := GetDataDiskon(stats)
	fmt.Println(data)
}

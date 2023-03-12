package nizar

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertDataTagihan(t *testing.T) {
	dbname := "tagihandb"
	datatagihan:= DataTagihan{
		ID:      primitive.NewObjectID(),
		Nama_Mahasiswa:    "Nizar Abdul Kholiq",
		Program_studi :   "Teknik Informatika",
		Jumlah: "200.000",
		Status: "Belum",
		
	}
	insertedID := InsertDataTagihan(dbname, datatagihan)
	if insertedID == nil {
		t.Error("Failed to insert billing")
	}
}

func TestInsertDataSudah(t *testing.T) {
	dbname := "tagihandb"
	datasudah := DataSudah{
		ID:       primitive.NewObjectID(),
		Nama_Mahasiswa: "Zaya wijaya",
		Status: "Sudah",
	}
	insertedID := InsertDataSudah(dbname, datasudah)
	if insertedID == nil {
		t.Error("Failed to insert user")
	}
}


func TestInsertDataBelum(t *testing.T) {
	dbname := "tagihandb"
	databelum := DataBelum{
		ID:       primitive.NewObjectID(),
		Nama_Mahasiswa: "Surya Azi",
		Status: "Belum",
	}
	insertedID := InsertDataBelum(dbname, databelum)
	if insertedID == nil {
		t.Error("Failed to insert user")
	}
}
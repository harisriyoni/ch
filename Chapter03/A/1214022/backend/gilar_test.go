package gilar

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertDashboard(t *testing.T) {
	dbname := "tablerps"
	dashboard := Dashboard{
		ID:         primitive.NewObjectID(),
		Username:   "Patricia Semklo",
		Email:      "patricia.semklo@app.com",
		Location:   "Denmark",
		Orders:     "6",
		Lastorders: "#904300",
		Totalspent: "$3,455.00",
		Sks:        "12",
	}
	insertedID := InsertDashboard(dbname, dashboard)
	if insertedID == nil {
		t.Error("Failed to insert dashboard")
	}
}

func TestInsertMahasiswa(t *testing.T) {
	dbname := "tablerps"
	mahasiswa := Mahasiswa{
		Nama_mhs:   "John Doe",
		Email_mhs:  "johndoe@flex.co",
		Gambar_mhs: "image.jpg",
	}
	insertedID := InsertMahasiswa(dbname, mahasiswa)
	if insertedID == nil {
		t.Error("Failed to insert mahasiswa")
	}
}

func TestInsertDosen(t *testing.T) {
	dbname := "tablerps"
	dosen := Dosen{
		Nama:      "Macauley Herring",
		Jabatan:   "DOSEN",
		Noted:     "Dance is the hidden language of the soul.",
		Img_dosen: "image.jpg",
	}
	insertedID := InsertDosen(dbname, dosen)
	if insertedID == nil {
		t.Error("Failed to insert dosen")
	}
}

func TestInsertAbout(t *testing.T) {
	dbname := "tablerps"
	about := About{
		Pertanyaan: "Apakah kurikulum ini menyulitkan hidup?",
		Jawaban:    "jelas ini akan membuat jiwamu bergetar sampai ke bulan",
	}
	insertedID := InsertAbout(dbname, about)
	if insertedID == nil {
		t.Error("Failed to insert about")
	}
}
func TestInsertContacus(t *testing.T) {
	dbname := "tablerps"
	contacus := Contacus{
		Phone_number: "0822126722",
		Email:        "kucingterbang@gmail.crot",
	}
	insertedID := InsertContacus(dbname, contacus)
	if insertedID == nil {
		t.Error("Failed to insert team")
	}
}

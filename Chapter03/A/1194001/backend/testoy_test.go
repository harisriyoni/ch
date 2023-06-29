package main

import (
	"fmt"
	"testing"
)

func TestInsertPresensi(t *testing.T) {
	phonenumber := "6811110023231"
	checkin := "masuk"
	biodata := Mahasiswa{
		Nama:         "ujang",
		Phone_number: "6284564562",
		Prodi:        "S1 TEKNIK TUMBUHAN",
		Kelas:        "2A",
	}
	hasil := InsertPresensi(phonenumber, checkin, biodata)
	fmt.Println(hasil)

}

func TestGetMahasiswaFromPhoneNumber(t *testing.T) {
	phonenumber := "6811110023231"
	biodata := GetMahasiswaFromPhoneNumber(phonenumber)
	fmt.Println(biodata)
}

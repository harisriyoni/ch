package faisal

import (
	"fmt"
	"testing"
)

func TestInsertProfile(t *testing.T) {
	pendidikan := "S1"
	username := "Fluxy"
	checkin := "masuk"
	biodata := Profile{
		FullName : "Fitrana Soleh",
		Email : "fitrana12@gmail.com",
	}
	hasil:=InsertOneDoc(username , checkin , biodata)
	fmt.Println(hasil)
	fmt.Println(pendidikan)
	

}

func TestGetProfileByUsername(t *testing.T) {
	username := "Fitrana "
	biodata:=GetProfileByUsername(username)
	fmt.Println(biodata)
}
package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Rissa")
	if result != "Belajar Membuat Unit Test Rissa" {

		//panic("Failed unit test")
		//tidak disarankan pake panic karna jika 1 gagal maka semua akan berhenti

		//t.Fail()
		//fail() menggagalkan unit test namun akan melanjutkan unit tes. tapi semua tetap dianggap gagal

		//failNow() jika ada 1 gagal maka akan gagal unit test tsb
		t.Error("Ini error pake Error")
	}

	fmt.Println("Ini jalan kalau error")

	// untuk running unit test menggunakan terminal
	// go test
	// go test -v (melihat function apa saja yg di test
	// go test -v -run TestHelloWorld (nama file)
	// go test ./... (running 1 package / 1 folder)
}

func TestHelloWorldMarissa(t *testing.T) {
	result := HelloWorld("Rissa")
	if result != "Belajar Membuat Unit Test Rissa" {

		//panic("Failed unit test")
		//tidak disarankan pake panic karna jika 1 gagal maka semua akan berhenti

		//t.Fail()
		//fail() menggagalkan unit test namun akan melanjutkan unit tes. tapi semua tetap dianggap gagal
		//t.FailNow()
		//failNow() jika ada 1 gagal maka akan gagal unit test tsb

		t.Fatal("Ini eror ya")
		//Fatal() dia sama kaya Fail namun dia bisa diberi error message
	}
	fmt.Println("Ini gak jalan waktu error")

	// untuk running unit test menggunakan terminal
	// go test
	// go test -v (melihat function apa saja yg di test
	// go test -v -run TestHelloWorld (nama file)
	// go test ./... (running 1 package / 1 folder)
}

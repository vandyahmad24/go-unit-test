package helper

import "testing"

//untuk membuat file unit test hariss diakhir _test
//hello_world_test.go
//Aturan Function Harus diawali Test

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("ahmad")
	if result != "Hello ahmad" {
		t.Error("Harus Hello ahmad")
	}
//	tidakboleh return value
}

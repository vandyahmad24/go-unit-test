package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)


//benchmart
func BenchmarkHelloWorld(b *testing.B){
	for i:=0; i<b.N; i++{
		HelloWorld("vandy")
	}
}
//mapping benchmartk
func BenchmarkTable(b *testing.B){
	benchmart := []struct{
		name string
		request string
	}{
		{"Vandy","Vandy"},{"Ahmad","Ahmad"},{"Misry","Misry"},{"Ar Razy","Ar Razy"},
	}
	for _, v := range benchmart{
		b.Run(v.name, func(b *testing.B) {
			for i:=1; i<b.N; i++{
				HelloWorld(v.request)
			}
		})
	}
}


//untuk membuat file unit test hariss diakhir _test
//hello_world_test.go
//Aturan Function Harus diawali Test
func TestMain(m *testing.M){
	//sebelum unit test
	fmt.Println("Sebelum unit test")
	m.Run()
	fmt.Println("Setelah unit test")
}
//mapping

//unit sub test
func TestSubTest(t *testing.T){
	t.Run("Vandy", func(t *testing.T) {
		result := HelloWorld("misry")
		require.Equal(t, "Hello misry",result, "Hasil Harus Hello misry")
	})
	t.Run("Ahmad", func(t *testing.T) {
		result := HelloWorld("ar razy")
		require.Equal(t, "Hello ar razy",result, "Hasil Harus Hello ar razy")
	})
}


func TestHelloWorld(t *testing.T) {
	result := HelloWorld("vandy")
	require.Equal(t, "Hello vandy",result, "Hasil Harus Hello vandy")
	fmt.Println("Test diekseksui")
//	tidakboleh return value
}


func TestHelloWorldAhmad(t *testing.T) {
	result := HelloWorld("ahmad")
	assert.Equal(t, "Hello ahmad",result, "Hasil Harus Hello ahmad")
	fmt.Println("Test diekseksui")
	//	tidakboleh return value
}



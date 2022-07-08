package library

import "fmt"

type mahasiswa struct {
	nama string
	umur int
}

func Panggil() {
	var x1, x2, x3 mahasiswa
	x1.nama = "Agus"
	x1.umur = 25
	x2.nama = "Koalski"
	x2.umur = 23
	x3.nama = "Rendo"
	x3.umur = 19

	x1.metode()
	x2.metode()
	x3.metode()
}

func (x mahasiswa) metode() {
	fmt.Println("nama :", x.nama, "umur :", x.umur)
}

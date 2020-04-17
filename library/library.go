package library

import (
	"fmt"
)

type Mahasiswa struct {
	Nama string
	Umur int
}

func (m *Mahasiswa) TampilMahasiswa() {
	fmt.Println("Nama :", m.Nama)
	fmt.Println("Umur :", m.Umur)
}

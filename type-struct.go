package main

import "fmt"

func main() {

	type Mahasiswa struct {
		nama  string
		umur  int
		alamat string
	}

    var mhs1 Mahasiswa
    mhs1.nama = "Andi"
    mhs1.umur = 20
    mhs1.alamat = "Jalan Merdeka No. 10"

	fmt.Println(mhs1.nama)
}

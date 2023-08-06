package main

import "fmt"

func main() {
	type NoKTP int
	type Married bool

	type Mahasiswa struct {
		nama string
		umur int
		alamat string
	}

	var nik NoKTP = 1234123
	var married Married = true

	var mhswa Mahasiswa
	mhswa.nama = "Madam"
	mhswa.umur = 123

	fmt.Println(nik)
	fmt.Println(married)
	fmt.Println(mhswa)
}
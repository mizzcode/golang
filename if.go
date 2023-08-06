package main

func main() {
	name := "Mizz"

	if name == "Mizz" {
		println("Hello", name)
	} else if name == "Jani" {
		println("Hello", name)
	} else if name == "Jul" {
		println("Hello", name)
	} else {
		println("Kamu siapa")
	}

	// length := len(name)

	// if length > 5 {
	// 	println("Nama Terlalu Panjang")
	// } else {
	// 	println("Berhasil Daftar")
	// }

	// shorthand
	if length := len(name); length > 5 {
		println("Nama Terlalu Panjang")
	} else {
		println("Berhasil Daftar")
	}

}
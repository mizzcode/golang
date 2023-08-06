package main

func main() {
	name := "Salman"

	switch name {
		case "Mizz":
            println("Hello Mizz!")
		case "Jani":
			println("Hello Jani!")
		case "Salman":
			println("Hello Salman!")
        default:
			println("Kamu sopo")
	}

	// shorthand switch
	// switch length := len(name); length > 5 {
	// 	case true:
    //         println("Nama Terlalu Panjang!")
    //     case false:
    //         println("Nama Sudah Benar!")
	// }

	// shorthand switch
	length := len(name)
	switch {
		case length > 10:
            println("Nama Terlalu Panjang!")
        case length > 5:
            println("Nama Lumayan Panjang!")
		default:
			println("Nama Sudah Benar!")
	}
}
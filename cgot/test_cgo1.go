package main
import "C"

func main() {

	v := 42
	C.printint(C.int(v))

}







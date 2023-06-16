package variables

import "fmt"

func MuestroEnteros() {
	var intComun int
	intDe32 := int32(10)
	intDe64 := int64(100)

	fmt.Println("int comun = ", intComun)
	fmt.Println("int de 32 bits = ", intDe32)
	fmt.Println("int de 64 bits = ", intDe64)
}

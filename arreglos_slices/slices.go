package arreglos_slices

import "fmt"

var tablaS []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 98, 21, 642, 452, 45, 67, 36, 26, 30}

func MuestroSlice() {
	fmt.Println(tablaS)

	procion := arreglo[3:]  // Slice creado desde un vector, desde la posición 3
	procion2 := arreglo[:5] //Slice creado desde un verctor, de la posicion 0 a la 5
	procion3 := arreglo[6:7]

	fmt.Println(procion)
	fmt.Println(procion2)
	fmt.Println(procion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0)

	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("\nLargo %d, Capacidad %d ", len(nums), cap(nums))
	fmt.Println("")
}

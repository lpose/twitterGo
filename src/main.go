package main

import "fmt"

func sumar(x int, y int) int {
	return x + y
}

func main() {

	/*
			var n1 int
			n1 = 4
			n2 := 5
			fmt.Print("Hola, mundo\n")
			fmt.Println("Mi número favorito es: ", rand.Intn(10))
			fmt.Println("La suma de n1 + n2 es: ", sumar(n1, n2))


				for i := 0; i < 10; i++ {
					fmt.Println("Mi número favorito es: ", rand.Intn(10), "\n")
				}


			//Array

			array := [4]int{100, 200, 300, 400}
			fmt.Println(array)

			array[3] = 2000
			fmt.Println("Array[3]: ", array[3])

			//Slices

			sliceTest := make([]int, 3)
			fmt.Println(sliceTest)
			sliceTest[0] = 0
			sliceTest[1] = 1
			sliceTest[2] = 2
			fmt.Println("Slice sin agregados: ", sliceTest)
			sliceTest = append(sliceTest, 3)

			fmt.Println("Slice con agregados: ", sliceTest)



		i, err := strconv.Atoi("asd")

		if err != nil {
			fmt.Printf("No se puede convertir el número: %v\n", err)
			return
		} else {
			fmt.Println(i)
			return
		}
	*/

	a := 4
	p := &a
	suma(p)
	fmt.Println(a)

}

func suma(x *int) *int {
	*x = (*x) * 2
	return x
}

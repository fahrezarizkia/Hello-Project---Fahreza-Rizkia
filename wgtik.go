package main

import "fmt"

func main() {
	var n,jum,rata2 int

	for i := 1; i <= 3; i++ {
		fmt.printf("Quiz %d: ", i)
		fmt.scanln(&n)
		jum += n
	}
	
	rata2 =  jum / 3
	fmt.Println("Rata-rata: ", rata2)

	if rata2 >= 60 {
		fmt.Println("lulus")
	} else {
		fmt.Println("tidak lulus")
	}
}

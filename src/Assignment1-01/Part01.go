package main

import "fmt"

func main() {
	fmt.Println("Enter the dimensions.")
	var rowNo int
	var colNo int
	fmt.Print("Number of Rows: ")
	fmt.Scanf("%d\n", &rowNo)
	fmt.Print("Number of Columns:")
	fmt.Scanf("%d\n", &colNo)
	// allocate composed 2d array
	values := make([][]int, rowNo)
	for i := range values {
		values[i] = make([]int, colNo)
	}
	fmt.Print("Enter the elements (0 or 1):\n")

	var userInput int
	for i := 0; i < rowNo; i++ {
		for j := 0; j < colNo; j++ {
			//for 1 == 1 {
				fmt.Print("[", i, "][", j, "]: ")
				fmt.Scanf("%d\n", &userInput)
				//if userInput != 1 && userInput != 0 {
				//	fmt.Println("The input value is not valid. Please enter 1 or 0.")
				//} else {
					values[i][j] = userInput
				//	break
				//}
			//}
		}
	}

	// These are the rows.
	//row1 := []int{1,1,0,0,0}
	//row2 := []int{1,1,0,0,0}
	//row3 := []int{0,0,1,0,0}
	//row4 := []int{0,0,0,1,1}

	// Append each row to the two-dimensional slice.
	//values = append(values, row1)
	//values = append(values, row2)
	//values = append(values, row3)
	//values = append(values, row4)

	//rowNo = 4

	//colNo = 5

	fmt.Println("\nThe array entered is: ")
	for i := 0; i < rowNo; i++ {
		for j := 0; j < colNo; j++ {
			fmt.Print("\t", values[i][j])
		}
		fmt.Print("\n")
	}

	var count int
	for i := 0; i < rowNo; i++ {
		for j := 0; j < colNo; j++ {
			if values[i][j] == 1 {
				count = count + 1
				//fmt.Println("count: ")
				//fmt.Println(count)
				checkIsland(i, j, values)
			}
		}
	}

	fmt.Println("\nNo. of islands = ", count)
}

func checkIsland(a int, b int, array [][]int) {
	//fmt.Print(a)
	//fmt.Println(b)

	if a < 0 || a >= len(array) || b < 0 || b >= len(array[a]) || array[a][b] != 1 {
		return
	}
	if b < 0 || b >= len(array[a]) {
		return
	}

	array[a][b] = -1
	checkIsland(a, b-1, array)
	checkIsland(a, b+1, array)
	checkIsland(a-1, b, array)
	checkIsland(a+1, b, array)
}

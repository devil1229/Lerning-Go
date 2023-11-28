// Golang Program to Print Pascal's Triangle

package main

import "fmt"

func printPascals(){

	var rows int
    temp := 1
	fmt.Println("Enter the number of rows: ")
	fmt.Scan(&rows)

	for i := 0; i < rows ; i++ {
		
		for j:= 1; j <= rows -i; j++ {
			fmt.Print(" ")
		}

		for k:= 0; k <= i; k++ {

			if k == 0 || i == 0 {
				temp = 1
			} else {
                temp = temp*(i-k + 1)/k
			}

			fmt.Printf(" %v", temp)
		}

		fmt.Println("")
	}
}


// Golang Program to Add Two Matrix Using Multi-dimensional Arrays

func addTwoMatrix() {
	var matrix1[100][100] int
	var matrix2[100][100] int
	var sum[100][100] int
	var row,col int
	fmt.Print("Enter number of rows: ")
	fmt.Scanln(&row)
	fmt.Print("Enter number of cols: ")
	fmt.Scanln(&col)

	fmt.Println()
	fmt.Println("========== Matrix1 =============")
	fmt.Println()
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("Enter the element for Matrix1 %d %d :",i+1,j+1)
			fmt.Scanln(&matrix1[i][j])
		}
	}

	fmt.Println()
	fmt.Println("========== Matrix2 =============")
	fmt.Println()

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("Enter the element for Matrix2 %d %d :",i+1,j+1)
			fmt.Scanln(&matrix2[i][j])
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			sum[i][j] = matrix1[i][j]+matrix2[i][j]
		}
	}

	fmt.Println()
	fmt.Println("========== Sum of Matrix =============")
	fmt.Println()

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf(" %d ",sum[i][j])
			if(j==col-1){
				fmt.Println("")
			}
		}
	}
}

//GOLANG Program to Generate Fibonacci Sequence Up to a Certain Number

func fibbonacci() {
	firstTerm := 0
	secondTerm :=1
	nextTerm := 0

	var n int
	fmt.Println("Please enter the number of terms: ")
	fmt.Scan(&n)
	fmt.Print("Fibonacci Series :")
	for i := 0; i<n; i++ {
		if i == 0{
			fmt.Print(" ", firstTerm)
			continue
		}
		if i == 1{
			fmt.Print(" ", secondTerm)
			continue
		}

		nextTerm = firstTerm + secondTerm
		firstTerm = secondTerm
		secondTerm = nextTerm
		fmt.Print(" ", nextTerm)
	}
}
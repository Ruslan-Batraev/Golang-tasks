package main

import (
	//"bufio"
	"fmt"
	//"os"
)

func main() {
	//task1()
	//sum()
	//evenOrOdd()
	//factorial()
	//workWithString("Hello World")
	//findMax()
	//isPalindrom("шалаш")
	//sumArr()
	//generateFibonacciNumbers(8)
	getUniqueElements([]int{1, 2, 2, 3, 4, 4, 5, 1})
}

func task1() {
	fmt.Println("Hello, World!")

}

func sum() {
	var num1, num2 int
	fmt.Printf("Input two numbers: ")
	fmt.Scanln(&num1, &num2)
	fmt.Printf("%d + %d = %d\n", num1, num2, num1+num2)
}

func evenOrOdd() {
	var num int
	fmt.Printf("Input number: ")
	fmt.Scanln(&num)
	if num%2 == 0 {
		fmt.Println("Четное")
	} else {
		fmt.Println("Нечётное")
	}
}

func factorial() {
	var num int
	fact := 1

	fmt.Print("Input number: ")
	fmt.Scanln(&num)

	if num < 0 {
		fmt.Println("Факториал не определён для отрицательных чисел.")
		return
	}

	if num == 0 {
		fmt.Println("Факториал 0 равен 1.")
		return
	}

	for i := 1; i <= num; i++ {

		fact *= i
	}
	fmt.Printf("Факториал числа %d равен %d\n", num, fact)
}

func workWithString(str string) string {
	//var str string
	//fmt.Println("Input the string: ")
	//scanner := bufio.NewScanner(os.Stdin)

	//if scanner.Scan() {
	//	str = scanner.Text()
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	str = string(runes)
	//}
	fmt.Println(str)
	return str
}

func findMax() {
	max := 0
	numbers := [7]int{1, 6, 3, 3, 6, 7, 2}
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	fmt.Println(max)
}

func isPalindrom(str string) bool {
	fmt.Println(str)
	strTest := workWithString(str)
	fmt.Println("---------------------")
	if strTest == str {
		fmt.Println("палиндромом")
		return true
	}
	fmt.Println("не палиндромом")
	return false
}

func sumArr() {
	sum := 0
	arr := [...]int{0, 1, 6, 4, 3, 5, 7, 3}
	for _, value := range arr {
		sum += value
	}
	fmt.Println(sum)
}

func generateFibonacciNumbers(n int) []int {
	if n <= 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}

	fibSequence := make([]int, n)
	fibSequence[0] = 0
	fibSequence[1] = 1

	for i := 2; i < n; i++ {
		fibSequence[i] = fibSequence[i-1] + fibSequence[i-2]
	}
	fmt.Println(fibSequence)
	return fibSequence
}

func getUniqueElements(arr []int) []int {
	encountered := make(map[int]bool)

	uniqueElements := []int{}

	for _, element := range arr {
		if !encountered[element] {
			encountered[element] = true
			uniqueElements = append(uniqueElements, element)
		}
	}
	fmt.Println(uniqueElements)
	return uniqueElements
}

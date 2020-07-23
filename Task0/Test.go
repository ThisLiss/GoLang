package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"math"
	"strings"
	"golang.org/x/tour/pic"
)

//Упражнение на замыкания
func fibonacci() func() int {
	first, second := 0, 1
	toPrint := 0
	return func() int {
		toPrint = first
		first, second = second, first + second
		return toPrint
	}
}

//Упражнение на срезы
func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dy)

	for i := range res {
		res[i] = make([]uint8, dx)
	}

	for i := range res {
		for j := range res[i] {
			res[i][j] = uint8(i * j)
		}
	}

	return res
}

//Упражнение на циклы
func Sqrt(x float64) float64 {
	z := 1.0
	old := 0.0
	for i:=0; i <= 100; i++	{
		old = z
		z = z - (z*z - x)/(2*z)
		if Equals(old,z) {
			break
		}
	}
	return z
}

//Вспомогательная функция для упражнения по циклам
func Equals(numA, numB float64) bool {
	delta := math.Abs(numA - numB)
	if delta < 1e-10 {
		return true
	}
	return false
}

//Упражнение на карты
func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordCount := make(map[string]int)
	for i := range words {
		wordCount[words[i]]++
	}

	return wordCount
}

func main() {
	//Запуск упражнения на срезы
	pic.Show(Pic)

	//Запуск упражнения на циклы
	fmt.Println(Sqrt(101))

	//Запуск упражнения на замыкания
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	//Запуск упражнения на карты
	wc.Test(WordCount)
}

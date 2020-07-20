package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	testText1 := "I ate a donut. Then I ate another donut."

	randomBigSlice := getRandomSlice(300)

	testSlice1 := []int{1,2,3,4,5,6,7,8,9,10}

	testSlice2 := testSlice1[3:8]

	fmt.Println("Map 1 Есть текст, надо посчитать сколько раз каждое слова встречается.")
	fmt.Println(map1(testText1))
	fmt.Println()

	fmt.Println("Map 2 Есть очень большой массив или slice целых чисел, надо сказать какие числа в нем упоминаются хоть по разу.")
	//Если нужно посмотреть на слайс, который передаём
	//fmt.Println("Рандомный слайлс: ", randomBigSlice)
	fmt.Println("Мап по рандомному слайсу: ",map2(randomBigSlice))
	fmt.Println()

	fmt.Println("Map 3 Есть два больших массива чисел, надо найти, какие числа упоминаются в обоих")
	fmt.Println(map3(testSlice1, testSlice2))
	fmt.Println()


	fmt.Println("Map 4 Сделать Фибоначчи с мемоизацией")
	returnNum := map4()
	fmt.Println(returnNum(7))
	fmt.Println(returnNum(1))
	fmt.Println(returnNum(2))
	fmt.Println(returnNum(6))
	fmt.Println(returnNum(0))

}

func getRandomSlice(len int) []int {
	slice := make([]int, len)
	for i := 0; i < len; i++ {
		slice[i] = rand.Intn(100)
	}

	return slice
}

func map1(s string) map[string]int {
	words := strings.Fields(s)
	wordCount := make(map[string]int)
	for i := range words {
		wordCount[words[i]]++
	}

	return wordCount
}



func map2(slice []int) map[int]int {
	nums := make(map[int]int)
	for i := range slice {
		nums[slice[i]]++
	}

	return nums
}

func map3(slice1 []int, slice2 []int) map[int]int {
	wordCount := make(map[int]int)
	for _, value := range slice1 {
		for _, value2 := range slice2 {
			if value == value2 {
				wordCount[value]++
			}
		}
	}

	return wordCount
}

func map4() func(n int) int {
	fibMap := make(map[int]int)
	fibMap[0] = 0
	fibMap[1]++
	first, second := -1, 1
	i := 0
	return func(n int) int {
		if fibMap[n] == 0 {
			for i <= n{
				newN := first + second
				first, second = second, first + second
				fibMap[i] = newN
				i++
			}
		}
		return fibMap[n]
	}
}

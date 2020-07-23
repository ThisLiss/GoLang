package main

import (
	"fmt"
	"sort"
)

func main() {
	mainSlice := getSlice( 10)

	fmt.Println("Начальный слайлс")
	fmt.Println(mainSlice)
	fmt.Println()

	slice1(mainSlice)

	slice2(mainSlice)

	slice3(mainSlice)

	fmt.Println("Slice 4 Взять последнее число slice, вернуть его пользователю, а из slice этот элемент удалить")
	fmt.Println(slice4(&mainSlice))
	fmt.Println()

	fmt.Println("Slice 5 Взять первое число slice, вернуть его пользователю, а из slice этот элемент удалить")
	fmt.Println(slice5(&mainSlice))
	fmt.Println()

	fmt.Println("Slice 6 Взять i-е число slice, вернуть его пользователю, а из slice этот элемент удалить. Число i передает пользователь в функцию")
	fmt.Println(slice6(4,&mainSlice))
	fmt.Println()

	fmt.Println("Slice 7 Объединить два slice и вернуть новый со всеми элементами первого и второго")
	fmt.Println(slice7(mainSlice[0:3], getSlice(5)))
	fmt.Println()

	slice8(getSlice(5), getSlice(3))

	secondSlice := getSlice(10)

	slice9(secondSlice)

	slice10(3,secondSlice)

	slice11(secondSlice)

	slice12(3, secondSlice)

	fmt.Println("Slice 13 - Вернуть пользователю копию переданного slice")
	fmt.Println(slice13(secondSlice))
	fmt.Println()

	slice14(secondSlice)

	slice15(secondSlice)
}

func getSlice(len int) []int {
	slice := make([]int, len)
	for i := 0; i < len; i++{
		slice[i] = i
	}
	return slice
}

func slice1(slice []int){
	addedSlice := getSlice(len(slice))
	for i := range addedSlice{
		addedSlice[i]++
	}
	fmt.Println("Slice 1 - К каждому элементу []int прибавить 1 ")
	fmt.Println(addedSlice)
	fmt.Println()
}

func slice2(slice []int){
	slice = append(slice,  5)
	fmt.Println("Slice 2 - Добавить в конец slice число 5 ")
	fmt.Println(slice)
	fmt.Println()
}

func slice3(slice []int){
	addedSlice := []int{5}
	slice = append(addedSlice, slice...)
	fmt.Println("Slice 3 - Добавить в начало slice число 5 ")
	fmt.Println(slice)
	fmt.Println()
}

func slice4(slice *[]int) int{
	deletedValue := (*slice)[len(*slice)-1]
	*slice = (*slice)[:len(*slice)-1]
	return deletedValue
}

func slice5(slice *[]int) int{
	deletedValue := (*slice)[0]
	*slice = (*slice)[1:]
	return deletedValue
}

func slice6(delIndex int, slice *[]int) int{
	deletedValue := (*slice)[delIndex]
	*slice = append((*slice)[:delIndex],(*slice)[delIndex+1:]...)
	return deletedValue
}

func slice7(slice1 []int, slice2 []int) []int{
	return append(slice1, slice2...)
}

func slice8(slice1 []int, slice2 []int){
	fmt.Println("Slice 8 Из первого slice удалить все числа, которые есть во втором")
	fmt.Println("Начальные слайлы: ", slice1, slice2)
	for _, value1 := range slice2 {
		for j, value2 := range slice1 {
			if value1 == value2 {
				slice1 = append((slice1)[:j], (slice1)[j+1:]...)
			}
		}
	}
	fmt.Println("Конечный: ", slice1)
	fmt.Println()
}

func slice9(slice []int) {
	slice = append(slice[1:], slice[0])
	fmt.Println("Slice 9 - Сдвинуть все элементы slice на 1 влево. Нулевой становится последним, первый - нулевым, последний - предпоследним")
	fmt.Println(slice)
	fmt.Println()
}

func slice10(i int, slice []int) {
	slice = append(slice[i:], slice[:i]...)
	fmt.Println("Slice 10 - Тоже, но сдвиг на заданное пользователем i")
	fmt.Println(slice)
	fmt.Println()
}

func slice11(slice []int) {
	slice = append(slice[len(slice)-1:], slice[:len(slice)-1]...)
	fmt.Println("Slice 11 - Тоже, что 9, но сдвиг вправо")
	fmt.Println(slice)
	fmt.Println()
}

func slice12(i int, slice []int) {
	slice = append(slice[len(slice)-i:], slice[:len(slice)-i]...)
	fmt.Println("Slice 12 - Тоже, но сдвиг на i")
	fmt.Println(slice)
	fmt.Println()
}

func slice13(slice []int) []int{
	sliceCopy := make([]int, len(slice))
	copy(sliceCopy,slice)
	return sliceCopy
}

func slice14(slice []int){
	for i:= 0; i < len(slice)-1; i += 2{
		temp := slice[i]
		slice[i] = slice[i+1]
		slice[i+1] = temp
	}
	fmt.Println("Slice 14 - В slice поменять все четные с ближайшими нечетными индексами. 0 и 1, 2 и 3, 4 и 5...")
	fmt.Println(slice)
	fmt.Println()
}

func slice15(slice []int){
	lexicoSlice := []string{"Евгений", "Вова", "Анастасия", "Михаэль"}
	fmt.Println("Slice 15 - Упорядочить slice в порядке: прямом, обратном, лексикографическом.")
	sort.Sort(sort.IntSlice(slice))
	fmt.Println("Прямой упорядоченный слайс", slice)
	sort.Sort(sort.Reverse(sort.IntSlice(slice)))
	fmt.Println("Обратно упорядоченный слайс", slice)
	sort.Strings(lexicoSlice)
	fmt.Println("Лексикографически упорядоченный слайс строк", lexicoSlice)
}
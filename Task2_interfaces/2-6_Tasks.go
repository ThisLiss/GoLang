package main

import (
	"fmt"
	"reflect"
	"sort"
)

type SavedOrders struct {
	orders []string
	price float32
}


func addProduct(prod *map[string]float32, name string, price float32){
	(*prod)[name] = price
}

func deleteProduct(prod *map[string]float32, name string){
	delete(*prod, name)
}

func changePrice(prod *map[string]float32, name string, price float32){
	_, ok := (*prod)[name]
	if ok {
		addProduct(prod, name, price)
	}
}

func changeName(prod *map[string]float32, name string, changedName string)  {
	(*prod)[changedName] = (*prod)[name]
	delete(*prod,name)
}

func calcOrderPrice(prod map[string]float32, order []string) float32{
	var price float32

	for i := range order {
		price+= prod[order[i]]
	}

	return price
}

func calcOrderPriceSaved() func (prod map[string]float32, order []string) float32{
	var count = 0
	var savedOrders = make([]SavedOrders, 1000)

	return func(prod map[string]float32, order []string) float32 {

		var price float32
		var isExists = false

		for j := range savedOrders {
			if reflect.DeepEqual(savedOrders[j].orders,order) {
				isExists = true
				price = savedOrders[j].price
			}
		}

		if !isExists {
			fmt.Println("Считаем...")
			price = 0.0
			for i := range order {
				price+= prod[order[i]]
			}
			savedOrders[count].orders = order
			savedOrders[count].price = price
			count++
		} else {
			fmt.Println("Мы считали, мы считали и уж очень мы устали! Не считаем!")
		}


		return price
	}
}

func sortAccounts(accs map[string]float64){
	keys := make([]string, 0, len(accs))
	values := make([]float64, 0, len(accs))
	for k, v := range accs {
		keys = append(keys, k)
		values = append(values, v)
	}

	sort.Strings(keys)
	fmt.Println("\nСортировка в алфавитном порядке: ")
	for i := range keys {
		fmt.Printf("%v ", keys[i])
	}

	sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	fmt.Println("\nСортировка в обратном порядке: ")
	for i := range keys {
		fmt.Printf("%v ", keys[i])
	}

	fmt.Println("\nПо количеству денег по убыванию: ")
	sort.Sort(sort.Reverse(sort.Float64Slice(values)))
	for i := range values {
		fmt.Printf("%.0f ", values[i])
	}
	fmt.Println()
}

func main() {

	//2: Мапа с товарами. Написать методы добавления, удаления, изменения цены товара, изменения имени товара.
	products := make(map[string]float32)

	addProduct(&products, "Бананы", 59.99)
	addProduct(&products, "Мандарины", 105)
	addProduct(&products, "Дошик", 14.99)
	addProduct(&products, "Аниме", 999.99)
	fmt.Println(products)

	deleteProduct(&products, "Аниме")
	fmt.Println(products)

	changePrice(&products, "Мандарины", 99.99)
	fmt.Println(products)

	changeName(&products, "Дошик", "Дошик куриный")
	fmt.Println(products)

	//3: Пользователь даёт список товаров, программа должна по map с наименованиями товаров посчитать сумму заказа.
	order := []string{"Бананы", "Мандарины"}
	fmt.Println()
	fmt.Printf("Цена товаров: %v\n", calcOrderPrice(products, order))


	//4: Сделать 1е, но у нас приходит несколько сотен таких списков заказов и мы хотим запоминать уже посчитанные заказы, чтобы если встречается такой же, то сразу говорить его цену без расчёта.
	order2 := []string{"Бананы", "Мандарины", "Дошик куриный"}
	order3 := []string{"Бананы", "Дошик куриный"}
	saved := calcOrderPriceSaved()
	fmt.Println()
	fmt.Printf("Цена товаров: %v\n", saved(products, order))
	fmt.Printf("Цена товаров: %v\n", saved(products, order2))
	fmt.Printf("Цена товаров: %v\n", saved(products, order))
	fmt.Printf("Цена товаров: %v\n", saved(products, order3))

	//5: Сделать пользовательские аккаунты со счетом типа "вася: 300р, петя: 30000000р".
	accounts := map[string]float64{
		"Вася":300,
		"Петя":30000000,
	}

	//6: Есть map аккаунтов и счетов, как описано в 3. Надо вывести ее в отсортированном виде с сортировкой: по имени в алфавитном порядке, по имени в обратном порядке, по количеству денег по убыванию.
	sortAccounts(accounts)
}




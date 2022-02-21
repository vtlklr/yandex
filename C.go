package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Offer struct {
	Market_sku int    `json:"market_sku"`
	Offer_id   string `json:"offer_id"`
	Price      int    `json:"price"`
}
type lessFunc func(p1, p2 *Offer) bool

type Offers struct {
	Offers []Offer `json:"offers"`
	less   []lessFunc
}

func (ms *Offers) Sort(Offers []Offer) {
	ms.Offers = Offers
	sort.Sort(ms)
}

// OrderedBy возвращает Offers, который выполняет сортировку с использованием меньших функций по порядку.
// Вызов его метода Sort для сортировки данных.
func OrderedBy(less ...lessFunc) *Offers {
	return &Offers{
		less: less,
	}
}

// Len является частью sort.Interface.
func (ms *Offers) Len() int {
	return len(ms.Offers)
}

// Своп - это часть sort.Interface.
func (ms *Offers) Swap(i, j int) {
	ms.Offers[i], ms.Offers[j] = ms.Offers[j], ms.Offers[i]
}

// Меньше - это часть sort.Interface. Это реализуется путем зацикливания по
// меньше функций, пока не будет найдено сравнение, которое различает
// два элемента (один меньше другого). Обратите внимание, что он может вызывать
// меньше функций в два раза за вызов. Мы могли бы изменить функции, чтобы вернуть
// -1, 0, 1 и уменьшим количество вызовов для большей эффективности:
// упражнение для читателя.
func (ms *Offers) Less(i, j int) bool {
	p, q := &ms.Offers[i], &ms.Offers[j]
	// Пробуем все, кроме последнего сравнения.
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			// p <q, значит, у нас есть решение.
			return true
		case less(q, p):
			// p> q, значит, у нас есть решение.
			return false
		}
		// p == q; попробуйте следующее сравнение.
	}
	// Все сравнения здесь говорят "равно", поэтому просто возвращаем все
	// итоговые сравнительные отчеты.
	return ms.less[k](p, q)
}
func Scan1() string {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода: ", err)
	}
	return str
}
func main() {
	//{"offers": [{"offer_id": "offer1", "market_sku": 13290123, "price": 1499},{"offer_id": "offer2", "market_sku": 500000, "price": 499}]}
	//{"offers": [{"offer_id": "offer3", "market_sku": 66666666, "price": 1499}]}
	m := Scan1()
	m = strings.TrimSpace(m)
	n, err := strconv.Atoi(m)
	if err != nil {
		fmt.Println(err)
	}

	of := make([]Offers, n)
	var s string
	for i := 0; i < n; i++ {
		//читаем строку, чистим
		s = strings.TrimSpace(Scan1())
		//преобразуем строку в срез байт
		data := []byte(s)
		//декодируем срез байт
		if err := json.Unmarshal(data, &of[i]); err != nil {
			fmt.Println(err)
			return
		}

	}
	var all Offers
	var one Offer

	for i := 0; i < n; i++ {
		l := len(of[i].Offers)
		for j := 0; j < l; j++ {
			one = of[i].Offers[j]
			all.Offers = append(all.Offers, one)
		}

	}
	price := func(c1, c2 *Offer) bool {
		return c1.Price < c2.Price
	}
	marketsku := func(c1, c2 *Offer) bool {
		return c1.Market_sku < c2.Market_sku
	}
	OrderedBy(price, marketsku).Sort(all.Offers)

	d, _ := json.Marshal(all)
	fmt.Println(string(d))
}

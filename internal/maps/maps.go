package maps

import (
	"fmt"
	"sort"
)

type User struct {
	ID     int
	Name   string
	City   string
	Active bool
	Tags   []string
}

type Product struct {
	SKU      string
	Quantity int
	Price    int
}

func CountWords(words []string) map[string]int {
	// TODO: посчитать, сколько раз встречается каждое слово.
	// Вернуть новую готовую к записи map. Регистр и пустая строка имеют значение.
	m := make(map[string]int)
	for _, val := range words {
		m[val] = m[val] + 1
	}
	return m
}

func GetOrDefault(values map[string]int, key string, fallback int) int {
	// TODO: вернуть значение существующего ключа или fallback, если ключ отсутствует.
	// Нулевое значение по существующему ключу должно возвращаться как обычное значение.
	_, ok := values[key]
	if ok {
		return values[key]
	}
	return fallback
}

func DeleteAndReport(values map[string]int, key string) bool {
	// TODO: удалить ключ и сообщить, существовал ли он до вызова.
	// Для отсутствующего ключа, пустой и nil map вернуть false без panic.
	_, ok := values[key]
	if ok != true || values == nil || len(values) == 0 {
		return false
	}
	delete(values, key)
	return true
}

func MergeCounters(left, right map[string]int) map[string]int {
	// TODO: объединить два счётчика в новой map, складывая значения общих ключей.
	// Исходные map после вызова должны остаться без изменений.
	sum := make(map[string]int)
	for i, val := range left {
		sum[i] = sum[i] + val
	}
	for i, val := range right {
		sum[i] = sum[i] + val
	}
	return sum
}

func KeysByValueAtLeast(values map[string]int, min int) []string {
	// TODO: вернуть ключи, значения которых не меньше min.
	// Результат должен иметь стабильный алфавитный порядок.
	key := make([]string, 0)
	for i, val := range values {
		if val >= min {
			key = append(key, i)
		}
	}
	sort.Strings(key)
	return key
}

func MaxKeyByValue(values map[string]int) (string, bool) {
	// TODO: найти ключ с максимальным значением.
	// Для пустого входа вернуть "", false; при равенстве выбрать меньший ключ по алфавиту.
	var key int
	var str string
	zn := false
	if values == nil || len(values) == 0 {
		return "", false
	}

	for i, val := range values {
		if zn {
			if key < val {
				key = val
				str = i
			} else if key == val {
				if str > i {
					str = i
				}
			}
		} else {
			key = val
			str = i
			zn = true
		}

	}
	return str, true

}

func BuildUserIndex(users []User) map[int]User {
	// TODO: построить индекс пользователей по ID.
	// Если один ID встречается несколько раз, в результате должен остаться последний пользователь.
	if users == nil || len(users) == 0 {
		return map[int]User{}
	}
	itog := make(map[int]User, 0)
	for i, val := range users {
		itog[val.ID] = users[i]
	}
	return itog
}

func GroupActiveUsersByCity(users []User) map[string][]User {
	// TODO: сгруппировать только активных пользователей по городу.
	// Порядок пользователей внутри каждого города должен совпадать с исходным слайсом.

	if users == nil || len(users) == 0 {
		return map[string][]User{}
	}
	itog := make(map[string][]User, 0)
	for i, val := range users {
		if val.Active == true {
			itog[val.City] = append(itog[val.City], users[i])
		}
	}
	return itog
}

func CountTags(users []User) map[string]int {
	// TODO: посчитать все появления тегов у всех пользователей.
	// Повторяющийся тег у одного пользователя также считается отдельным появлением.
	if users == nil || len(users) == 0 {
		return map[string]int{}
	}
	itog := make(map[string]int, 0)
	for _, val := range users {
		for _, b := range val.Tags {
			itog[b]++
		}
	}
	return itog
}

func BuildInventory(products []Product) map[string]Product {
	// TODO: построить склад по SKU.
	// При повторном SKU оставить последнюю запись из входного слайса.
	if products == nil || len(products) == 0 {
		return map[string]Product{}
	}
	itog := make(map[string]Product, 0)
	for i, val := range products {
		itog[val.SKU] = products[i]
	}
	return itog
}

func ReserveStock(inventory map[string]Product, sku string, count int) bool {
	// TODO: уменьшить остаток существующего товара на положительное count, если товара хватает.
	// При неуспехе вернуть false и не менять склад.
	st := false
	if inventory == nil || len(inventory) == 0 || count <= 0 {
		return false
	}
	s := Product{}
	for i, val := range inventory {
		s = val
		if s.SKU == sku {
			if s.Quantity >= count {
				s.Quantity -= count
				inventory[i] = s
				st = true
				fmt.Print(s.Quantity)
				return st
			}
		}
	}
	return false
}

func Restock(inventory map[string]*Product, sku string, count int) bool {
	// TODO: увеличить остаток товара, хранящегося в map как указатель.
	// Ключ должен существовать, указатель не должен быть nil, count должен быть положительным.
	st := false
	if inventory == nil || len(inventory) == 0 || count <= 0 {
		return false
	}
	//var s *Product
	for i, val := range inventory {
		if val == nil {
			return false
		}
		if i != sku {
			return false
		}
		//s = val
		if inventory[i].SKU == sku {
			inventory[i].Quantity += count
			//inventory[i] = s
			st = true
			//fmt.Print(s.Quantity)
			return st
		}
	}
	return false
}

func LowStockSKUs(inventory map[string]Product, limit int) []string {
	// TODO: вернуть SKU товаров с остатком не больше limit.
	// Результат должен иметь стабильный алфавитный порядок.
	s := []string{}
	if inventory == nil || len(inventory) == 0 {
		return s
	}
	for i, val := range inventory {
		if val.Quantity <= limit {
			s = append(s, i)
		}
	}
	sort.Strings(s)
	return s
}

func InventoryValue(inventory map[string]Product) int {
	// TODO: посчитать общую стоимость склада как сумму Quantity * Price.
	// Пустой и nil склад имеют стоимость 0.
	var s int
	if inventory == nil || len(inventory) == 0 {
		return 0
	}
	for _, val := range inventory {
		s = s + val.Quantity*val.Price
	}
	return s
}

func ApplyPriceUpdates(inventory map[string]Product, updates map[string]int) int {
	// TODO: применить положительные новые цены только к существующим товарам.
	// Вернуть число реально обновлённых товаров; неизвестные SKU и неположительные цены пропустить.
	count := 0
	if inventory == nil || len(inventory) == 0 || updates == nil || len(updates) == 0 {
		return 0
	}
	for i, val := range updates {
		s := inventory[i]
		if val > 0 && s.SKU == i {
			s.Price = val
			inventory[i] = s
			count++
		}

	}
	return count
}

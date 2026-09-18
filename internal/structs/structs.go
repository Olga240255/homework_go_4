package structs

import (
	"fmt"
	"slices"
	"strconv"
	"unsafe"
)

type User struct {
	ID         int
	Name       string
	Active     bool
	LoginCount int
	Tags       []string
}

func UpdateUserName(users map[int]User, id int, name string) bool {
	// TODO: изменить имя существующего пользователя, хранящегося в map как значение.
	// Если пользователя нет, вернуть false и не добавлять новый ключ.
	u, ok := users[id]
	if ok != true {
		return false
	}
	u.Name = name
	users[id] = u
	return true
}

func ActivateUser(users map[int]User, id int) bool {
	// TODO: сделать существующего пользователя активным.
	// Если пользователя нет, вернуть false.
	u, ok := users[id]
	if ok != true {
		return false
	}
	u.Active = true
	users[id] = u
	return true
}

func IncrementLoginCount(users map[int]User, id int) bool {
	// TODO: увеличить LoginCount существующего пользователя на единицу.
	// Если пользователя нет, вернуть false.
	u, ok := users[id]
	if ok != true {
		return false
	}
	u.LoginCount++
	users[id] = u
	return true
}

func CopyUsers(users map[int]User) map[int]User {
	// TODO: вернуть независимую копию map и вложенных слайсов Tags.
	// Изменения любой части результата не должны менять исходные данные.
	if users == nil {
		return map[int]User{}
	}
	s := make(map[int]User, len(users))
	for key, user := range users {
		// Копируем сам слайс Tags, чтобы разорвать связь с исходным
		var copiedTags []string
		if user.Tags != nil {
			copiedTags = make([]string, len(user.Tags))
			copy(copiedTags, user.Tags)
		}

		// Создаем копию структуры с новым слайсом
		userCopy := user
		userCopy.Tags = copiedTags

		// Сохраняем в новую карту
		s[key] = userCopy
	}
	return s
}

func BuildPointerIndex(users []User) map[int]*User {
	// TODO: построить индекс ID -> *User из независимых копий входных пользователей.
	// Указатели разных ID не должны вести к одному объекту или менять исходный слайс.
	if users == nil {
		return make(map[int]*User)
	}
	u := slices.Clone(users)
	m := make(map[int]*User)
	for i, val := range u {
		m[val.ID] = &u[i]
	}
	return m
}

func RenameThroughPointer(users map[int]*User, id int, name string) bool {
	// TODO: изменить имя объекта по существующему ненулевому указателю.
	// Для отсутствующего ключа или nil-значения вернуть false.
	user, ok := users[id]
	if ok == false || user == nil {
		return false
	}
	users[id].Name = name
	return true
}

func (u User) DisplayName() string {
	// TODO: вернуть строку "ID:Name".
	return fmt.Sprintf("%d:%s", u.ID, u.Name)

}

type Admin struct {
	User
	Permissions []string
}

func (a Admin) DisplayName() string {
	// TODO: вернуть строку "admin:ID:Name", перекрывая одноимённый метод встроенного User.
	return fmt.Sprintf("admin:%d:%s", a.ID, a.Name)
}

func (a Admin) HasPermission(permission string) bool {
	// TODO: проверить точное наличие permission в списке администратора.
	// Регистр, пробелы и повторения не нормализуются.
	for _, val := range a.Permissions {
		if val == permission {
			return true
		}
	}
	return false
}

func RenameEmbeddedUser(admin *Admin, name string) {
	// TODO: изменить Name у встроенного User.
	// Для nil admin функция должна безопасно завершиться.
	if admin == nil {
		return
	}
	admin.Name = name
}

type Account struct {
	ID      int
	Balance int
}

type PremiumAccount struct {
	Account
	Bonus int
}

func (p PremiumAccount) Total() int {
	// TODO: вернуть сумму основного баланса и бонуса.
	return p.Balance + p.Bonus
}

func PremiumLabel(p PremiumAccount) string {
	// TODO: вернуть строку "ID:Balance total=Total".
	return strconv.Itoa(p.ID) + ":" + strconv.Itoa(p.Balance) + " total=" + strconv.Itoa(p.Total())
}

type BadLayout struct {
	Active   bool
	Amount   int64
	Verified bool
	Retries  int32
}

type GoodLayout struct {
	Amount   int64
	Retries  int32
	Active   bool
	Verified bool
}

func UserValueSize(user User) uintptr {
	// TODO: вернуть фактический размер переданного значения User в байтах.
	return unsafe.Sizeof(user)
}

func UserPointerSize(user *User) uintptr {
	// TODO: вернуть фактический размер переменной-указателя на User в байтах.
	// Результат не зависит от того, nil указатель или нет.
	return unsafe.Sizeof(user)
}

func LayoutSizes() (bad, good, saved uintptr) {
	// TODO: вернуть размеры BadLayout и GoodLayout, затем число сэкономленных байт.
	return unsafe.Sizeof(BadLayout{}), unsafe.Sizeof(GoodLayout{}), unsafe.Sizeof(BadLayout{}) - unsafe.Sizeof(GoodLayout{})
}

var _ uintptr = unsafe.Sizeof(User{})

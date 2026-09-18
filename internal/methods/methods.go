package methods

import (
	"errors"
	"strconv"
)

var (
	ErrNilAccount        = errors.New("nil account")
	ErrNilDestination    = errors.New("nil destination")
	ErrSameAccount       = errors.New("source and destination are the same")
	ErrInactive          = errors.New("inactive account")
	ErrInvalidAmount     = errors.New("invalid amount")
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrNilCart           = errors.New("nil cart")
	ErrInvalidSKU        = errors.New("invalid sku")
	ErrCartLimit         = errors.New("cart limit exceeded")
)

type Account struct {
	ID       int
	Owner    string
	Balance  int
	Active   bool
	Currency string
}

func (a Account) Label() string {
	// TODO: вернуть строку "#ID Owner: Balance Currency" с одним пробелом после двоеточия.
	return "#" + strconv.Itoa(a.ID) + " " + a.Owner + ": " + strconv.Itoa(a.Balance) + " " + a.Currency
}

func (a Account) CanWithdraw(amount int) bool {
	// TODO: сообщить, можно ли списать amount без нарушения состояния счёта.
	// Счёт должен быть активным, сумма — положительной, денег должно хватать.
	if a.Active == true && amount > 0 && a.Balance >= amount {
		return true
	}
	return false

}

func (a Account) Renamed(owner string) Account {
	// TODO: вернуть копию счёта с новым владельцем.
	// Исходное значение Account не должно изменяться.
	b := a
	b.Owner = owner
	return b
}

func (a *Account) Rename(owner string) error {
	// TODO: изменить владельца исходного счёта.
	// Для nil receiver вернуть ErrNilAccount и не паниковать.
	if a == nil {
		return ErrNilAccount
	}
	a.Owner = owner
	return nil
}

func (a *Account) Deposit(amount int) error {
	// TODO: пополнить активный счёт на положительную сумму.
	// При ошибке вернуть подходящую sentinel-ошибку и сохранить прежний баланс.
	if a == nil {
		return ErrNilAccount
	}
	if a.Active == false {
		return ErrInactive
	}
	if amount <= 0 {
		return ErrInvalidAmount
	}
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount int) error {
	// TODO: списать положительную сумму с активного счёта, если денег достаточно.
	// При ошибке вернуть подходящую sentinel-ошибку и сохранить прежний баланс.
	if a == nil {
		return ErrNilAccount
	}
	if a.Active == false {
		return ErrInactive
	}
	if amount <= 0 {
		return ErrInvalidAmount
	}
	if a.Balance < amount {
		return ErrInsufficientFunds
	}
	a.Balance -= amount
	return nil
}

func (a *Account) TransferTo(destination *Account, amount int) error {
	// TODO: атомарно перевести сумму между двумя разными активными счетами.
	// При любой ошибке оба баланса должны остаться прежними.
	if a == nil {
		return ErrNilAccount
	}
	if destination == nil {
		return ErrNilDestination
	}
	if a.Active == false || destination.Active == false {
		return ErrInactive
	}
	if amount <= 0 {
		return ErrInvalidAmount
	}
	if a.Balance < amount {
		return ErrInsufficientFunds
	}
	if a == destination {
		return ErrSameAccount
	}
	a.Balance -= amount
	destination.Balance += amount
	return nil
}

func (a *Account) Activate() error {
	// TODO: сделать счёт активным.
	// Для nil receiver вернуть ErrNilAccount.
	if a == nil {
		return ErrNilAccount
	}
	a.Active = true
	return nil
}

func (a *Account) Deactivate() error {
	// TODO: сделать счёт неактивным.
	// Для nil receiver вернуть ErrNilAccount.
	if a == nil {
		return ErrNilAccount
	}
	a.Active = false
	return nil
}

func (a Account) Snapshot() Account {
	// TODO: вернуть независимую копию текущего значения счёта.
	var b Account
	b = a
	return b
}

func (a *Account) Reset() {
	// TODO: сбросить все поля существующего счёта в zero value.
	// Вызов для nil receiver должен быть безопасным.
	if a != nil {
		var b Account
		*a = b
	}

}

func (a *Account) ApplyBonus(percent int) error {
	// TODO: увеличить баланс активного счёта на указанный неотрицательный процент.
	// Дробная часть от целочисленного расчёта отбрасывается; при ошибке баланс не менять.

	if a == nil {
		return ErrNilAccount
	}
	if a.Active == false {
		return ErrInactive
	}
	if percent < 0 {
		return ErrInvalidAmount
	}
	a.Balance = a.Balance * (100 + percent) / 100
	return nil
}

func (a Account) SameOwner(other Account) bool {
	// TODO: сравнить владельцев двух счетов точным сравнением строк.
	return a.Owner == other.Owner
}

type Cart struct {
	Items    map[string]int
	MaxItems int
}

func (c *Cart) Add(sku string, count int) error {
	// TODO: добавить положительное количество товара с непустым SKU.
	// Общее число единиц после добавления не должно превышать MaxItems.
	if c == nil {
		return ErrNilCart
	}
	if sku == "" {
		return ErrInvalidSKU
	}
	if count <= 0 {
		return ErrInvalidAmount
	}
	if c.Items == nil {
		c.Items = make(map[string]int, 0)
	}
	sum := 0
	for _, val := range c.Items {
		sum += val
	}
	if sum+count > c.MaxItems {
		return ErrCartLimit
	} else {
		c.Items[sku] += count
	}
	return nil

}

func (c Cart) TotalItems() int {
	// TODO: вернуть сумму количеств всех товаров корзины.
	sum := 0
	for _, val := range c.Items {
		sum += val
	}
	return sum
}

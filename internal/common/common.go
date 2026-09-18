package common

import (
	"fmt"
	"strconv"

	"github.com/rinatkh/homework_go_4/internal/methods"
)

type Operation struct {
	AccountID int
	Kind      string
	Amount    int
}

type LedgerResult struct {
	Balances map[int]int
	Errors   []string
}

func ApplyOperations(accounts map[int]*methods.Account, operations []Operation) LedgerResult {
	// TODO: последовательно применить deposit/withdraw к счетам и продолжать после ошибок.
	// В Balances вернуть итоговые балансы ненулевых счетов, в Errors — сообщения в порядке операций.
	var LR LedgerResult
	LR.Balances = make(map[int]int)
	LR.Errors = make([]string, 0)
	//operations nil
	if operations == nil {
	} else {
		for _, val := range operations {
			r, ok := accounts[val.AccountID]
			if ok == false || r == nil {
				LR.Errors = append(LR.Errors, fmt.Sprintf("account %d: not found", val.AccountID))
			} else {
				switch val.Kind {
				case "deposit":
					{
						e := accounts[val.AccountID].Deposit(val.Amount)
						if e != nil {
							LR.Errors = append(LR.Errors, "account "+strconv.Itoa(val.AccountID)+": "+e.Error())
						}
						LR.Balances[val.AccountID] = accounts[val.AccountID].Balance
					}
				case "withdraw":
					{
						e := accounts[val.AccountID].Withdraw(val.Amount)
						if e != nil {
							LR.Errors = append(LR.Errors, "account "+strconv.Itoa(val.AccountID)+": "+e.Error())
						}
						LR.Balances[val.AccountID] = accounts[val.AccountID].Balance
					}
				default:
					{
						LR.Errors = append(LR.Errors, "account "+strconv.Itoa(val.AccountID)+": unknown operation \""+val.Kind+"\"")
						LR.Balances[val.AccountID] = accounts[val.AccountID].Balance
					}
				}
			}
		}

	}
	return LR
}

func BuildAccountIndex(accounts []methods.Account) map[int]*methods.Account {
	// TODO: построить индекс по ID из независимых копий счетов.
	// Разные ключи не должны вести к одной переменной, изменения индекса не должны менять входной слайс.
	caccounts := make([]methods.Account, 0)
	caccounts = accounts

	result := make(map[int]*methods.Account)
	if accounts == nil || len(accounts) == 0 {
		return result
	}
	for _, val := range caccounts {
		result[val.ID] = &val
	}
	return result
}

//`ApplyUserEvents(users map[int]string, active map[int]bool,
// events []Event) (map[int]string, map[int]bool)` |
// Две map состояния и события. |
// Непустое `Name` обновляет имя, `Active != nil` обновляет статус.
// Пустые поля события не стирают прежние данные.
// События применяются по порядку, новые ID разрешены.
// Если входная map не `nil`, изменить и вернуть тот же объект;
// если `nil` — создать новую map. |

type Event struct {
	UserID int
	Name   string
	Active *bool
}

func ApplyUserEvents(
	users map[int]string,
	active map[int]bool,
	events []Event,
) (map[int]string, map[int]bool) {
	if users == nil {
		users = make(map[int]string)
	}
	if active == nil {
		active = make(map[int]bool)
	}
	if events == nil {
		return users, active
	}
	for _, val := range events {
		if val.Name != "" {
			users[val.UserID] = val.Name
		}
		if val.Active != nil {
			active[val.UserID] = *val.Active
		}

	}
	// TODO: применить события к состоянию пользователей и вернуть обе map.
	// Непустое Name обновляет имя, ненулевой Active обновляет статус; nil map должны поддерживаться.
	return users, active
}

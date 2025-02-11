package banks

var (
	deposits = make(chan int)
	balance  = make(chan int)
)

func Deposits(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balance
}

func teller() {
	var balances int
	for {
		select {
		case amount := <-deposits:
			balances += amount
		case balance <- balances:
		}
	}
}

func init() {
	go teller()
}

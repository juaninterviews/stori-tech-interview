package uploader

type IBalanceManager interface {
	InsertBalance()
	GetBalance()
}

type BalanceManager struct{}

func NewBalanceManager() *BalanceManager {
	return &BalanceManager{}
}

func (m *BalanceManager) InsertBalance() {
	//TODO: Implement
}

func (m *BalanceManager) GetBalance() {
	//TODO: Implement
}

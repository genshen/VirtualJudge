package accounts

type PojAccountInterface struct {

}

func (pi PojAccountInterface)LoginAccount(accountIndex int8)error {
	println(accountIndex)
	return nil
}
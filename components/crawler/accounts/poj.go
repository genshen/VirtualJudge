package accounts

import "gensh.me/VirtualJudge/components/crawler/utils"

type PojAccountInterface struct {

}
//todo
func (pi PojAccountInterface)LoginAccount(*Account) error {
	println("1")
	return nil
}

//make sure accountIndex is safe!
func (pi PojAccountInterface)LoginAccountByIndex(accountIndex uint) error {
	return pi.LoginAccount(&OJs[utils.POJ - 1].Accounts[accountIndex])
}

//todo select the minimal task account
func (pi PojAccountInterface)GetAvailableAccount() (uint, *Account) {

	return 0, &OJs[utils.POJ - 1].Accounts[0]
}
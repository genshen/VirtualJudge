package status

import "gensh.me/VirtualJudge/components/crawler/accounts"

type StatusInterface interface {
	FetchStatus(ai accounts.AccountInterface, si StatusInterface, accountIndex uint, runId string) (*TaskResult, error)
}

//remember to follow the order in crawler/utils/values.go->const
var statusInterfaces = []StatusInterface{new(POJStatusInterface)}

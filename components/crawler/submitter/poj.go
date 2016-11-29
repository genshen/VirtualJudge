package submitter

import (
	"gensh.me/VirtualJudge/components/crawler/accounts"
)

type POJSubmitInterface struct {

}

func (r POJSubmitInterface)RemoteSubmit(account *accounts.Account, problemId string, language int8, code string) error {

	return nil
}

//convert language type to poj language type
//0:G++,1:GCC,2:JAVA,3:Pascal,4:C+,5:C,6:Fortran
func (r POJSubmitInterface) GetLanguageType(language int8) int8 {
	return language
}

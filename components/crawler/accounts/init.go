package accounts

import (
	"os"
	"log"
	"github.com/astaxie/beego"
	"io/ioutil"
	"encoding/json"
	"errors"
)

type AccountInterface interface {
	LoginAccount(*Account) error
	LoginAccountByIndex(accountIndex uint) error
	GetAvailableAccount() (uint,*Account)
}

type Account struct {
	Username string
	Password string
	Tasks    uint
	Session  string
}

type OJ struct {
	Name             string
	Enable           bool
	AccountInterface AccountInterface
	Accounts         []Account
}

var OJs []OJ

func init() {
	config_path := beego.AppConfig.String("oj_acounts_path")
	f, err := os.Open(config_path)
	if err != nil {
		log.Fatal(err)
		return
	}
	config, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
		return
	}

	var oj_config map[string]struct {
		Enable   bool `json:"enable"`
		Name     string `json:"name"`
		Accounts [] struct {
			Id       string  `json:"id"`
			Password string  `json:"password"`
		}
	}

	if err := json.Unmarshal(config, &oj_config); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}

	//edit here if you want to add more ojs,notice the index with crawler/utils/values.go->const
	// (eg:OJ_NAMES[utils.POJ-1] == "POJ")
	var OJ_NAMES = []string{"POJ", "HOJ"}
	var OJ_ACCOUNT_INTERFACE = []AccountInterface{new(PojAccountInterface)}
	for index, v := range OJ_NAMES {
		if ojValue, ok := oj_config[v]; ok &&len(ojValue.Accounts) != 0 {
			//add this oj information to variable OJs
			var oj = OJ{Enable:true, Name:v, AccountInterface:OJ_ACCOUNT_INTERFACE[index],
				Accounts:[]Account{}}
			for _, v := range ojValue.Accounts {
				//add all account to one oj
				oj.Accounts = append(oj.Accounts, Account{Username:v.Id, Password:v.Password, Tasks:0, Session:""})
			}
			OJs = append(OJs, oj)
		} else {
			OJs = append(OJs, OJ{Enable:false, Name:v, Accounts:[]Account{}})
		}
	}
}

/*
ojType:match with crawler/utils/values.go->const
all interface function will be visited by this interface,so don't always judge the array index in interface function
*/
func GetInterfaceByOjType(ojType int) (AccountInterface, error) {
	if index := ojType - 1; index < len(OJs) && index >= 0 && OJs[index].Enable {
		return OJs[index].AccountInterface, nil
		//if accountIndex < int8(len(OJs[ojType - 1].Accounts)) && accountIndex >= 0 {
		//	return aif.LoginAccount(accountIndex)
		//}
	}
	return nil, errors.New("no oj resource found") // no account found
}
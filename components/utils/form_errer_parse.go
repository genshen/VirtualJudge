package utils

import (
	"github.com/astaxie/beego/validation"
	"encoding/json"
	"fmt"
)

type SimpleJsonResponse struct {
	Status   int
	Error    interface{}
	Addition interface{}
}

type Err struct {
	Name    string
	Message string
}

type Field struct {
	Value  string
	Errors []Err
}

func NewInstant(Errors []*validation.Error, f map[string]string) map[string]Field {
	var fields = make(map[string]Field)
	var F Field
	var ok bool
	for _, err := range Errors {
		if F, ok = fields[err.Key]; !ok {
			//not exists, add
			F = Field{Value:f[err.Key]}
		}
		F.Errors = append(F.Errors, Err{err.Key, err.Message})
		fields[err.Key] = F
	}
	return fields
}

func NewInstantToByte(Errors []*validation.Error, f map[string]string) []byte {
	b, err := json.Marshal(NewInstant(Errors, f))
	if err != nil {
		fmt.Println("json err:", err) //todo err return
	}
	return b
}

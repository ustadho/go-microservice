package models

import "fmt"

type Flags struct {
	Ip   string
	Port string
}

var flagObj *Flags

func (f *Flags) GetApplicationUrl() (*string, *ErrorDetail) {
	f, err := GetFlags()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s:%s", f.Ip, f.Port)
	return &url, nil
}

func NewFlags(ip, port string) *Flags {
	if flagObj == nil {
		flagObj = &Flags{
			Ip:   ip,
			Port: port,
		}
	}
	return flagObj
}

func GetFlags() (*Flags, *ErrorDetail) {
	if flagObj == nil {
		return nil, &ErrorDetail{
			ErrorType:    ErrorTypeFatal,
			ErrorMessage: "Flags not set",
		}
	}
	return flagObj, nil
}

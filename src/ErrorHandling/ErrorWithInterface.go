package main

import "fmt"

type error interface {
	Error() string
}

type PassWordError struct {
	Len         int
	RequiredLen int
}

func (err PassWordError) Error() string {
	return "암호 길이가 짧습니다"
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		return PassWordError{len(password), 8}
	}
	return nil
}

func main() {

	err := RegisterAccount("myID", "myPw")
	if err != nil {
		if errInfo, ok := err.(PassWordError); ok {
			fmt.Println(errInfo)
		}
	} else {
		fmt.Println("회원가입됐습니다")
	}

}

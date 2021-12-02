package main

import "fmt"

func main() {
	code := "001RMnFa1CHicC0X3HHa116Z8v1RMnF3"
	wxInfo, err := LoginWXSmall(code)
	if err != nil {
		fmt.Printf("LoginWXSmall err:%s", err.Error())
		return
	}

	encryptData := ""
	iv := ""
	m, err := DecryptWXOpenData(wxInfo.Sessionkey, encryptData, iv)
	if err != nil {
		fmt.Printf("DecryptWXOpenData err:%s", err.Error())
		return
	}
	fmt.Printf("m:%+", m)
}

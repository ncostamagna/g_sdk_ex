package main

import (
	"errors"
	"fmt"
	"os"

	userSdk "github.com/ncostamagna/g_sdk_ex/user"
)

func main() {
	userTrans := userSdk.NewHttpClient("http://localhost:8081", "")

	user, err := userTrans.Get("769fc490-9981-4710-b9f5-0f6557e619a4")
	if err != nil {
		if errors.As(err, &userSdk.ErrNotFound{}) {
			fmt.Println("Not found:", err.Error())
			os.Exit(1)
		}
		fmt.Println("Internal Server Error:", err.Error())
		os.Exit(1)
	}

	fmt.Println(user)

}

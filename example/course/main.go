package main

import (
	"errors"
	"fmt"
	"os"

	couseSdk "github.com/ncostamagna/g_sdk_ex/course"
)

func main() {
	courseTrans := couseSdk.NewHttpClient("http://localhost:8082", "")

	course, err := courseTrans.Get("ef5ca563-70e4-408b-8802-def77371b7bf")
	if err != nil {
		if errors.As(err, &couseSdk.ErrNotFound{}) {
			fmt.Println("Not found:", err.Error())
			os.Exit(1)
		}
		fmt.Println("Internal Server Error:", err.Error())
		os.Exit(1)
	}

	fmt.Println(course)

}

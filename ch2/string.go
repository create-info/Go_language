package ch2

import (
	"fmt"
	"strings"
)

func StringPrefix()  {
	fmt.Println(strings.HasPrefix("/serviceCode/airpay/mesgType/verify", "/serviceCode/airpay/mesgType1"))
}

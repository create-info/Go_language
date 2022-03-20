package ch2_test

import (
	"Go_language/ch2"
	"fmt"
	"testing"
)

func TestGetDateTimeWithTimeZone(t *testing.T) {
	timeWithTimeZone := ch2.GetDateTimeWithTimeZone()
	fmt.Println(timeWithTimeZone)
}

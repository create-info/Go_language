package ch2_test

import (
	"fmt"
	"testing"
)

func TestGetDateTimeWithTimeZone(t *testing.T) {
	timeWithTimeZone := Go_language.GetDateTimeWithTimeZone()
	fmt.Println(timeWithTimeZone)
}
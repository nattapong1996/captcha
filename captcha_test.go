package captcha

import (
	"fmt"
	"testing"
)

func TestCaptchaCase1111to1plusOne(t *testing.T)  {
	result := Captcha(1,1,1,1)
	experd := "1+one"
	if result == experd {
		fmt.Println(result)
	}
}

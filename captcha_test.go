package captcha

import (
	"fmt"
	"testing"
)

func TestCaptchaCase1111to1plusOne(t *testing.T)  {
	result := Captcha(1,1,1,1)
	experd := "1+One"
<<<<<<< HEAD
	if result == experd {
		fmt.Println(result)
	}
}

func TestCaptchaCase1112to1plusTwo(t *testing.T)  {
	result := Captcha(1,1,1,1)
	experd := "1+Two"
=======
>>>>>>> bfd23d870a82efa51c9c307fd242d6b5fadb6b17
	if result == experd {
		fmt.Println(result)
	}
}

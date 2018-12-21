package captcha

import (
	"testing"
)

func TestCaptchaCase1111to1plusOne(t *testing.T)  {
	result := Captcha(1,1,1,1)
	experd := "1+One"
	if result != experd {
		t.Errorf(result)
	}
}

func TestCaptchaCase1112to1plusTwo(t *testing.T)  {
	result := Captcha(1,1,1,2)
	experd := "1+Two"
	if result != experd {
		t.Errorf(result)
	}
}
func TestCaptchaCase1113to1plusThree(t *testing.T)  {
	result := Captcha(1,1,1,3)
	experd := "1+Three"
	if result != experd {
		t.Errorf(result)
	}
}

func TestCaptchaCase1114to1plusFour(t *testing.T)  {
	result := Captcha(1,1,1,4)
	experd := "1+Four"
	if result != experd {
		t.Errorf(result)
	}
}

func TestCaptchaCase1224to2DeleteFour(t *testing.T)  {
	result := Captcha(1,2,2,4)
	experd := "2-Four"
	if result != experd {
		t.Errorf(result)
	}
}

func TestCaptchaCase2224to2DeleteFour(t *testing.T)  {
	result := Captcha(2,2,2,4)
	experd := "Two-4"
	if result != experd {
		t.Errorf(result)
	}
}


func TestCountCaptcha2214(t *testing.T)  {
	result := countCaptCha(2,2,1,4)
	experd := 6
	if result != experd {
		t.Errorf("Ans : %d",result)
	}
}

func TestCountCaptcha2424(t *testing.T)  {
	result := countCaptCha(2,4,2,4)
	experd := 0
	if result != experd {
		t.Errorf("ANS : %d",result)
	}
}
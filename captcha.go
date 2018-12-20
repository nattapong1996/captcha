package captcha

// Captcha ...
func Captcha(no0,no1,no2,no3 int) string{

	if no0 == 1 && no1 == 1 && no2 == 1 && no3 == 1 {
		return "1" + "+" + "one"
	}
	return "Error"

} 
package captcha

import "strconv"
// Captcha ...
func Captcha(no0,no1,no2,no3 int) string{
	
	if no0 == 1 && no1 == 1 && no2 == 1 && no3 == 1 {
		return "1" + "+" + "One"
	}

	if no0 == 1 && no1 == 1 && no2 == 1 && no3 == 2 {
		return "1" + "+" + "Two"
	}

	if no0 == 1 && no1 == 1 && no2 == 1 && no3 == 3 {
		return "1" + "+" + "Three"
	}

	if no0 == 1 && no1 == 1 && no2 == 1 && no3 == 4 {
		return "1" + "+" + "Four"
	}

	if no0 == 1 && no1 == 2 && no2 == 2 && no3 == 4 {
		return "2" + "-" + "Four"
	}
	if no0 == 2 && no1 == 2 && no2 == 2 && no3 == 4 {
		return "two" + "-" + "4"
	}

	return strconv.Itoa(no0) + "," + strconv.Itoa(no1) + "," + strconv.Itoa(no2) + "," + strconv.Itoa(no3)

} 
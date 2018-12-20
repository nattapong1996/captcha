package captcha

import "strconv"
// Captcha ...
func Captcha(no0,no1,no2,no3 int) string{

	opeRand := [10]string{"One","Two","Three","Four","Five","Six","Seven","Eight","์Nine"}
	opeRator := [3]string{"+","-","*"}
	var result = ""

	if no0 == 1 {

		result = strconv.Itoa(no1) + opeRator[no2-1] + opeRand[no3-1]

	}else if no0 == 2 {

		result = opeRand[no1-1] + opeRator[no2-1] + strconv.Itoa(no3)

	} else {

		result = strconv.Itoa(no0) + "," + strconv.Itoa(no1) + "," + strconv.Itoa(no2) + "," + strconv.Itoa(no3)

	}

	return result

} 
package captcha

import "strconv"

// type rusultCaptcha struct{
// 	result string
// 	amount int
// }

// Captcha ...
func Captcha(no0,no1,no2,no3 int) string{

	opeRand := [10]string{"","One","Two","Three","Four","Five","Six","Seven","Eight","Nine"}
	opeRator := [4]string{"","+","-","*"}
	var result = ""

	if no0 >= 1 && no0 <= 2 && no1 >= 1 && no1 <= 9 && no2 >= 1 && no2 <= 3 && no3 >= 1 && no3 <= 9 {
		if no0 == 1 {
			result = strconv.Itoa(no1) + opeRator[no2] + opeRand[no3]
		} else if no0 == 2 {
			result = opeRand[no1] + opeRator[no2] + strconv.Itoa(no3)
		}
	} else { 
		result = strconv.Itoa(no0)+","+strconv.Itoa(no1)+","+strconv.Itoa(no2)+","+strconv.Itoa(no3)
	}

	return result

}

func countCaptCha(no0,no1,no2,no3 int) int {
	amount := 0
	switch no2 {
		case 1	:	amount = no1 + no3
		case 2	:	amount = no1 - no3
		case 3	:	amount = no1 * no3
		default	:	amount = 0
	}
	return amount
}
package creon

import ole "github.com/hspan/go-ole"
// 아래는 VARIANT로 리턴한 값을 각 타입으로 변환하여 리턴하는 함수들이다.
// []string, int, float, string 타입을 지원하는 함수가 포함됨
// 다른 타입은 go-ole 패키지의 variant.go 의 Value함수를 참조할 것

//RetSS - VARIANT를 []string 타입으로 변환하여 리턴한다.
//입력 : *ole.VARIANT
//리턴 : []string
func RetSS(r *ole.VARIANT) (ret []string) {
	data := r.ToArray().ToValueArray()
	for _, value := range data {
		strValue := value.(string)
		ret = append(ret, strValue)
	}
	return 
}

//RetInt 정수 반환 int32로 반환된 값을 int로 변환하여 리턴함
func RetInt(r *ole.VARIANT) (ret int) {
	ret = int(int32(r.Val))
	return
}


//RetStr 문자열 반환
func RetStr(r *ole.VARIANT) (ret string) {
	ret = r.ToString()
	return
}

//RetBool 참/거짓 반환
func RetBool(r *ole.VARIANT) (ret bool) {
	ret = r.Val != 0
	return
}

//RetLong 롱타입(int64)
func RetLong(r * ole.VARIANT) (ret int64) {
	ret = int64(r.Val)
	return
}
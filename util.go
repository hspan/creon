package creon

import ole "github.com/hspan/go-ole"
// 아래는 VARIANT로 리턴한 값을 각 타입으로 변환하여 리턴하는 함수들이다.
// []string, int, float, string 타입을 지원하는 함수가 포함됨

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

//RetInt 정수 반환
func RetInt(r *ole.VARIANT) (ret int) {
	ret = int(r.Value().(int32))
	return
}


//RetStr 문자열 반환
func RetStr(r *ole.VARIANT) (ret string) {
	ret = r.Value().(string)
	return
}

//RetBool 참/거짓 반환
func RetBool(r *ole.VARIANT) (ret bool) {
	ret = r.Val != 0
	return
}
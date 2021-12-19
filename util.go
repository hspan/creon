package creon

import (
	ole "github.com/go-ole/go-ole"
	"fmt"
	"unsafe"
	"reflect"
)

// 아래는 VARIANT로 리턴한 값을 각 타입으로 변환하여 리턴하는 함수들이다.
// []string, int, float, string 타입을 지원하는 함수가 포함됨
// 다른 타입은 go-ole 패키지의 variant.go 의 Value함수를 참조할 것

//RetSS - VARIANT를 []string 타입으로 변환하여 리턴한다.
//입력 : *ole.VARIANT
//리턴 : []string
func RetSS(r *ole.VARIANT) (ret []string) {
	a:= r.ToArray().ToValueArray()
	for _, v := range a {
		ret = append(ret, v.(string))
	}
	return
}

//RetInt (long) 값을 int로 반환
//32bit로 작동하므로 int = int32임
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

//RetLong long long 를 int64로 반환함
func RetLong(r *ole.VARIANT) (ret int64) {
	ret = int64(r.Val)
	return
}

// Int64 - 각 종 정수형값을 int64로 반환함
// 주의사항 - ulonglong의 경우 오류가 있을 수 도 있음(매우 큰 값(2^63 이상의 값))
//          - 반환값이 int가 아님을 주의
func Int64(r *ole.VARIANT) (ret int64) {
	switch r.VT {
	case ole.VT_I2, ole.VT_UI2, ole.VT_I4, ole.VT_UI4, ole.VT_I8, ole.VT_UI8:
		ret = r.Val
	default:
		panic(fmt.Errorf("반환된 값은 %s입니다. 확인 후 프로그램을 수정하시기 바랍니다.", Check_VT(r)))
	}
	return
}

func Check_VT(r *ole.VARIANT) (ret string) {
	switch r.VT {
	case ole.VT_I2:
		ret = "Int16(short)"
	case ole.VT_I4:
		ret = "Int32(long)"
	case ole.VT_I8:
		ret = "float32(long long)"
	case ole.VT_R4:
		ret = "float32(float)"
	case ole.VT_R8:
		ret = "float64(double)"
	case ole.VT_BSTR:
		ret = "string"
	case (ole.VT_ARRAY | ole.VT_VARIANT):
		// ole.VARIANT의 배열
		a := r.ToArray().ToValueArray()
		if reflect.TypeOf(a).Elem().Kind() == reflect.String {
			ret = "[]string(string list)"
		}
	}
	return
}

// Int32 - 각 종 정수형값을 int32로 반환함
func Int32(r *ole.VARIANT) (ret int32) {
	switch r.VT {
	case ole.VT_I2, ole.VT_UI2, ole.VT_I4, ole.VT_UI4:
		ret = int32(r.Val)
	default:
		panic(fmt.Errorf("반환된 값은 %s입니다. 확인 후 프로그램을 수정하시기 바랍니다.", Check_VT(r)))
	}
	return
}

// Int - 각 종 정수형값을 int로 반환함 (32bit)
func Int(r *ole.VARIANT) (ret int) {
	ret = int(Int32(r))
	return
}

// Float32 : 반환형이 float일때
func Float32(r *ole.VARIANT) (ret float32) {
	switch r.VT {
	case ole.VT_R4:
		ret = *(*float32)(unsafe.Pointer(&r.Val))
	default:
		panic(fmt.Errorf("반환된 값은 %s입니다. 확인 후 프로그램을 수정하시기 바랍니다.", Check_VT(r)))
	}
	return
}

// Float64 : 반환형이 double일때 사용
func Float64(r *ole.VARIANT) (ret float64) {
	switch r.VT {
	case ole.VT_R8:
		ret = *(*float64)(unsafe.Pointer(&r.Val))
	default:
		panic(fmt.Errorf("반환된 값은 %s입니다. 확인 후 프로그램을 수정하시기 바랍니다.", Check_VT(r)))
	}
	return
}

// String : 문자열 반환
func String(r *ole.VARIANT) (ret string) {
	switch r.VT {
	case ole.VT_BSTR:
		ret = r.ToString()
	default:
		panic(fmt.Errorf("반환된 값은 %s입니다. 확인 후 프로그램을 수정하시기 바랍니다.", Check_VT(r)))
	}
	return
}

// StrSlice : 문자열 슬라이스 반환
func StrSlice(r *ole.VARIANT) (ret []string) {
	a := r.ToArray().ToValueArray()
	for _, v := range a {
		ret = append(ret, v.(string))
	}
	return
}

//Bool 참/거짓 반환
func Bool(r *ole.VARIANT) (ret bool) {
	ret = r.Val != 0
	return
}
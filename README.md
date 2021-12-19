# creon
=====
이 패키지는 대신증권의 creonplus(cybosplus) API를 golang에서 사용할 수 있도록 한 것입니다.


github.com/ippoeyeslhw/cpgo 를 수정하여 만들었습니다.
## 수정내용
1. coclass들 중 자주 사용하는 일부는 별도로 정의 하였습니다.
CpClass만 정의 된 상태에서는 coclass에 없는 함수를 호출하더라도 실행전에는 에러를 확인할 수 없어 CpClass를 임베딩하여 선언하고 각각의 coclass별로 함수를 선언합니다.
2. 공통적으로 사용하는 setValue 등은 CpClass에 선언하여 coclass들도 사용할 수 있도록 했습니다.

3. cpgo 패키지에 없는 coclass의 함수들을 추가할 예정입니다.

4. CpClass 구조체의 obj 가 외부에서 참조할 수 없도록 되어 있어 Obj로 수정하여 외부에서 호출할 수 있도록 수정하여 패키지 내에 정의되지 않은 함수도 사용할 수 있도록 하였습니다.
 
v0에서는 하나의 오브젝트에 있는 함수를 다 추가하면 마이너 버전을 올릴 예정입니다. - 주식 거래만 
(선물이나 elw등은 제외함)
선물 등은 추후 v1에나 추가할 예정입니다.

## 버전별
### v0
#### v0.0.0 - CpCodeMgr 일부
#### V0.1.0 - CpCodeMgr + CpCybos
#### V0.2.1 - CpStockMst, CpStockCode, CpTdUtil 추가
<br>
작업을 시작한 단계로 정상적인 작동은 보장하지 않습니다.
<p>

# 필요한것들
 * 윈도우 운영체제
 * 사이보스플러스
 * golang (32bits)
 * [go-ole패키지](https://github.com/go-ole/go-ole)
 * 사이보스플러스 [도움말(비공식)](http://cybosplus.github.io/)

# 설치

```
go get github.com/hspan/creon

```

# 사용법
## 프로그램 시작
```
err := ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED)
if err != nil {panic(err)}
defer ole.CoUninitialize()
```
위와 같은 내용을 creon plus를 사용하는 최상위 함수(보통 main)에 호출해야 합니다.
ole.CoInitialize(0)의 경우에는 에러가 발생하는 경우가 있었으며 멀티스레딩을 사용하지 못하는 것이므로 위와 같이 멀티스레딩이 가능하도록 호출하는 것이 좋을 것 같습니다.

## 객체 생성
1. 패키지에 정의된 coclass
```
codeMgr := &creon.CpCodeMgr{}
codeMgr.Create()
```
위와 같이 &creon.coclass명{}으로 선언하고 Create 함수로 생성합니다.

2. 패키지에 없는 coclass
```
object := &creon.CpClass{}
object.Create("CpTrade.CpTd0311")
```
(CpTd0311는 이 패키지 내에 정의 되어 있지만 위와 같은 방식으로 선언하여 사용하는 것도 가능합니다.)

## 정의되지 않은 함수를 사용할 경우
예) 해외 전종목 코드 조회
```
// 객체 생성
CpUsCode := &creon.CpClass{}
CpUsCode.Create("CpUtil.CpUsCode")

// 함수 호출
list := creon.StrSlice(GetUsCodeList(CpUsCode, 1))
```

```
// GetUsCodeList 호출 함수
func GetUsCodeList(c *creon.CpClass, USTYPE int) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetUsCodeList", USTYPE)
}
```

## 리턴된 인터페이스를 각 타입으로 변환함수
RetSS - []string으로 변환
RetInt - int32로 반환된 값을 int로 다시 변환하여 반환
RetStr - string 반환
RetBool - 불 값 반환
RetLong - Long형태의 정수를 int64로 반환함
위 5개는 추후 버전이 올라가면 삭제할 예정입니다.  아래의 함수를 사용하십시오.

크레온 반환값에 따른 함수선택
short. long : Int, Int32
short. long, longlong : Int64
float : Float32
double : Float64
bool : Bool
string : String
string배열 : StrSlice

## 주의사항
creon plus(cybos plus)는 32bit입니다.  이를 사용하기 위해서는 32bit로 컴파일 하여야 합니다.

컴퓨터에 설치된 go가 64bit일 경우 32bit로 크로스 컴파일 하여야 합니다.
windows의 명령프롬프트에서는
```
set goarch=386
```

파워쉘의 경우에는 
```
$env:GOARCH="386"
```
한 후 컴파일 하십시오.

64비트로 컴파일 한 후 실행할 경우 
```
panic: 클래스가 등록되지 않았습니다.
```
에러가 발생합니다.

## char
입력할 변수의 형이 char일 경우 golang의 Rune (작은 따옴표로 묶인 문자)를 사용해야 함

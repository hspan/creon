# creon
=====
이 패키지는 대신증권의 creonplus(cybosplus) API를 golang에서 사용할 수 있도록 한 것입니다.


github.com/ippoeyeslhw/cpgo 를 수정하여 만들었습니다.
## 수정내용
1. coclass들을 별도로 정의 하였습니다.
CpClass만 정의 된 상태에서는 coclass에 없는 함수를 호출하더라도 실행전에는 에러를 확인할 수 없어 CpClass를 임베딩하여 선언하고 각각의 coclass별로 함수를 선언합니다.
2. 공통적으로 사용하는 setValue 등은 CpClass에 선언하여 coclass들도 사용할 수 있도록 합니다.

3. cpgo 패키지에 없는 coclass의 함수들을 추가할 예정입니다.


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
위와 같은 내용을 creon plus를 사용하는 최상위 함수(보통 amin)에 호출해야 합니다.
ole.CoInitialize(0)의 경우에는 에러가 발생하는 경우가 있었으며 멀티스레딩을 사용하지 못하는 것이므로 위와 같이 멀티스레딩이 가능하도록 호출하는 것이 좋을 것 같습니다.

## 객체 생성
```
codeMgr := &creon.CpCodeMgr{}
codeMgr.Create()
```
위와 같이 &creon.coclass명{}으로 선언하고 Create 함수로 생성합니다.

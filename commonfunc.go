package creon

import (
	ole "github.com/hspan/go-ole"
	"github.com/hspan/go-ole/oleutil"
	"golang.org/x/sys/windows"
	"fmt"
)

func IsUserAnAdmin() (bool, error) {
	shell32 := windows.NewLazySystemDLL("Shell32.dll")
	defer windows.FreeLibrary(windows.Handle(shell32.Handle()))

	isUserAnAdminProc := shell32.NewProc("IsUserAnAdmin")
	ret, _, winError := isUserAnAdminProc.Call()

	if winError != windows.NTE_OP_OK {
		return false, fmt.Errorf("IsUserAnAdmin returns error code %d", winError)
	}
	if ret == 0 {
		return false, nil
	}
	return true, nil
}

// 사이보스플러스 Property Getter 메서드입니다.
// Continue 프로퍼티 값을 얻어옵니다.
func (c *CpClass) GetContinue() (r *ole.VARIANT) {
	return oleutil.MustGetProperty(c.obj, "Continue")
}

// 사이보스플러스 SetInputValue 메서드 Wrapper
func (c *CpClass) SetInputValue(typ int, val interface{}) {
	_ = oleutil.MustCallMethod(c.obj, "SetInputValue", typ, val)
}

// 사이보스플러스 BlockRequest 메서드 Wrapper
func (c *CpClass) BlockRequest() (result *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "BlockRequest")
}

// 사이보스플러스 BlockRequest2 메서드 Wrapper
func (c *CpClass) BlockRequest2(option int) (result *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "BlockRequest2", option)
}

// 사이보스플러스 BlockRequest 메서드 Wrapper
func (c *CpClass) Request() {
	if c.evnt == nil {
		panic("err")
	}
	_ = oleutil.MustCallMethod(c.obj, "Request")
}

// 사이보스플러스 Subscribe 메서드 Wrapper
func (c *CpClass) Subscribe() {
	if c.evnt == nil {
		panic("err")
	}
	_ = oleutil.MustCallMethod(c.obj, "Subscribe")
}

// 사이보스플러스 SubscribeLastest 메서드 Wrapper
func (c *CpClass) SubscribeLastest() {
	if c.evnt == nil {
		panic("err")
	}
	_ = oleutil.MustCallMethod(c.obj, "SubscribeLastest")
}

// 사이보스플러스 Unsubscribe 메서드 Wrapper
func (c *CpClass) Unsubscribe() {
	_ = oleutil.MustCallMethod(c.obj, "Unsubscribe")
}

// 사이보스플러스 GetHeaderValue 메서드 Wrapper
func (c *CpClass) GetHeaderValue(typ int) (result *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetHeaderValue", typ)
}

// 사이보스플러스 GetDataValue  메서드 Wrapper
func (c *CpClass) GetDataValue(typ int, idx int) (result *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetDataValue", typ, idx)
}

// 사이보스플러스 GetDibStatus 메서드 Wrapper
func (c *CpClass) GetDibStatus() (result *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetDibStatus")
}

// 사이보스플러스 GetDibMsg1 메서드 Wrapper
func (c *CpClass) GetDibMsg1() (result *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetDibMsg1")
}


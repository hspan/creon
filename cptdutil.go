package creon

import (
	ole "github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

type CpTdUtil struct {
	CpClass
}

func (c *CpTdUtil) Create() {
	c.CpClass.Create("CpTrade.CpTdUtil")
}

// property
func (c *CpTdUtil) AccountNumber() (r *ole.VARIANT) {
	return oleutil.MustGetProperty(c.obj, "AccountNumber")
}

//method
// 함수 인자를 전달받도록 되어 있으나 0으로만 사용하므로 인자 없이 사용함
func (c *CpTdUtil) TradeInit() (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "TradeInit", 0)
}

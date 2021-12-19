package creon

import (
	ole "github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

type CpCybos struct {
	CpClass
}

func (c *CpCybos) Create() {
	c.CpClass.Create("CpUtil.CpCybos")
}

//Property
func (c *CpClass) IsConnect() (r *ole.VARIANT) {
	return oleutil.MustGetProperty(c.Obj, "IsConnect")
}

func (c *CpClass) ServerType() (r *ole.VARIANT) {
	return oleutil.MustGetProperty(c.Obj, "ServerType")
}

func (c *CpClass) LimitRequestRemainTime() (r *ole.VARIANT) {
	return oleutil.MustGetProperty(c.Obj, "LimitRequestRemainTime")
}

//Method
func (c *CpCybos) GetLimitRemainCount(limitType int) {
	_ = oleutil.MustCallMethod(c.Obj, "GetLimitRemainCount", limitType)
}

func (c *CpClass) PlusDisconnect() {
	_ = oleutil.MustCallMethod(c.Obj, "PlusDisconnect")
}
package creon

import (
	ole "github.com/hspan/go-ole"
	"github.com/hspan/go-ole/oleutil"
)

type CpCybos struct {
	CpClass
}

func (c *CpCybos) Create() {
	c.CpClass.Create("CpUtil.CpCybos")
}

//Property
func (c *CpClass) IsConnect() (r *ole.VARIANT) {
	return oleutil.MustGetProperty(c.obj, "IsConnect")
}

func (c *CpClass) ServerType() (r *ole.VARIANT) {
	return oleutil.MustGetProperty(c.obj, "ServerType")
}

func (c *CpClass) LimitRequestRemainTime() (r *ole.VARIANT) {
	return oleutil.MustGetProperty(c.obj, "LimitRequestRemainTime")
}

//Method
func (c *CpCybos) GetLimitRemainCount(limitType int) {
	_ = oleutil.MustCallMethod(c.obj, "GetLimitRemainCount", limitType)
}

func (c *CpClass) PlusDisconnect() {
	_ = oleutil.MustCallMethod(c.obj, "PlusDisconnect")
}

//Event
// OnDisConnect
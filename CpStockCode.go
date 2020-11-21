package creon

import (
	ole "github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

type CpStockCode struct {
	CpClass
}

func (c *CpStockCode) Create() {
	c.CpClass.Create("CpUtil.CpStockCode")
}

func (c *CpStockCode) CodeToName(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "CodeToName", code)
}

func (c *CpStockCode) NameToCode(name string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "NameToCode", name)
}

func (c *CpStockCode) CodeToFullCode(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "CodeToFullCode", code)
}

func (c *CpStockCode) FullCodeToName(fullcode string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "FullCodeToName", fullcode)
}

func (c *CpStockCode) FullCodeToCode(fullcode string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "FullCodeToCode", fullcode)
}

func (c *CpStockCode) CodeToIndex(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "CodeToIndex", code)
}

func (c *CpStockCode) GetCount() (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetCount", code)
}

func (c *CpStockCode) GetData(typ int, idx int) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetData", typ, idx)
}

func (c *CpStockCode) GetPriceUnit(code string, basePrice int64, directionUp bool) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetPriceUnit", code, basePrice, directionUp)
}

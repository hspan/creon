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
	return oleutil.MustCallMethod(c.Obj, "CodeToName", code)
}

func (c *CpStockCode) NameToCode(name string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "NameToCode", name)
}

func (c *CpStockCode) CodeToFullCode(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "CodeToFullCode", code)
}

func (c *CpStockCode) FullCodeToName(fullcode string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "FullCodeToName", fullcode)
}

func (c *CpStockCode) FullCodeToCode(fullcode string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "FullCodeToCode", fullcode)
}

func (c *CpStockCode) CodeToIndex(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "CodeToIndex", code)
}

func (c *CpStockCode) GetCount() (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetCount")
}

func (c *CpStockCode) GetData(typ int, idx int) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetData", typ, idx)
}

func (c *CpStockCode) GetPriceUnit(code string, basePrice int64, directionUp bool) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetPriceUnit", code, basePrice, directionUp)
}

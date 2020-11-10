package creon

import (
	ole "github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

func (c *CpClass) CodeToName(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "CodeToName", code)
}

func (c *CpClass) GetStockMarketKind(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockMarketKind", code)
}

func (c *CpClass) GetStockSectionKind(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockSectionKind", code)
}

func (c *CpClass) GetStockListedDate(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockListedDate", code)
}

func (c *CpClass) GetStockYdClosePrice(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockYdClosePrice", code)
}

func (c *CpClass) IsSPAC(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "IsSPAC", code)
}

func (c *CpClass) GetStockListByMarket(market int) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockListByMarket", market)
}


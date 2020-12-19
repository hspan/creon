package creon

import (
	ole "github.com/hspan/go-ole"
	"github.com/hspan/go-ole/oleutil"
)

type CpCodeMgr struct {
	CpClass
}

func (c *CpCodeMgr) Create() {
	c.CpClass.Create("CpUtil.CpCodeMgr")
}

func (c *CpCodeMgr) CodeToName(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "CodeToName", code)
}

func (c *CpCodeMgr) GetStockMarginRate(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockMarginRate", code)
}

func (c *CpCodeMgr) GetStockMemeMin(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockMemeMin", code)
}

func (c *CpCodeMgr) GetStockMarketKind(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockMarketKind", code)
}

func (c *CpCodeMgr) GetStockIndustryCode(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockIndustryCode", code)
}

func (c *CpCodeMgr) IsSPAC(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "IsSPAC", code)
}


func (c *CpCodeMgr) GetStockControlKind (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockControlKind", code)
}

func (c *CpCodeMgr) GetStockSupervisionKind (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockSupervisionKind", code)
}

func (c *CpCodeMgr) GetStockStatusKind (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockStatusKind", code)
}

func (c *CpCodeMgr) GetStockCapital (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockCapital", code)
}

func (c *CpCodeMgr) GetStockFiscalMonth (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockFiscalMonth", code)
}

func (c *CpCodeMgr) GetStockGroupCode (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockGroupCode", code)
}

func (c *CpCodeMgr) GetStockKospi200Kind (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockKospi200Kind", code)
}

func (c *CpCodeMgr) GetStockSectionKind(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockSectionKind", code)
}

func (c *CpCodeMgr) GetStockLacKind (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockLacKind", code)
}

func (c *CpCodeMgr) GetStockListedDate (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockListedDate", code)
}

func (c *CpCodeMgr) GetStockMaxPrice (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockMaxPrice", code)
}

func (c *CpCodeMgr) GetStockMinPrice (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockMinPrice", code)
}

func (c *CpCodeMgr) GetStockParPrice (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockParPrice", code)
}

func (c *CpCodeMgr) GetStockStdPrice (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockStdPrice", code)
}

func (c *CpCodeMgr) GetStockYdOpenPrice (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockYdOpenPrice", code)
}

func (c *CpCodeMgr) GetStockYdHighPrice (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockYdHighPrice", code)
}

func (c *CpCodeMgr) GetStockYdLowPrice (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockYdLowPrice", code)
}

func (c *CpCodeMgr) GetStockYdClosePrice (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockYdClosePrice", code)
}

func (c *CpCodeMgr) IsStockCreditEnable(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "IsStockCreditEnable", code)
}

func (c *CpCodeMgr) GetStockParPriceChageType (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockParPriceChageType", code)
}



func (c *CpCodeMgr) GetMiniFutureList() (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetMiniFutureList")
}

func (c *CpCodeMgr) GetMiniOptionList() (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetMiniOptionList")
}

func (c *CpCodeMgr) ReLoadPortData() (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "ReLoadPortData")
}

func (c *CpCodeMgr) GetStockElwBasketCodeList(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockElwBasketCodeList", code)
}

func (c *CpCodeMgr) GetStockElwBasketCompList(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockElwBasketCompList", code)
}

func (c *CpCodeMgr) GetStockListByMarket(marketcode string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockListByMarket", marketcode)
}

func (c *CpCodeMgr) GetGroupCodeList(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetGroupCodeList", code)
}

func (c *CpCodeMgr) GetGroupName (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetGroupName", code)
}

func (c *CpCodeMgr) GetIndustryList () (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetIndustryList")
}

func (c *CpCodeMgr) GetIndustryName (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetIndustryName", code)
}

func (c *CpCodeMgr) GetMemberList ( ) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetMemberList")
}

func (c *CpCodeMgr) GetMemberName (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetMemberName", code)
}

func (c *CpCodeMgr) GetMarketStartTime ()  (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetMarketStartTime")
}

func (c *CpCodeMgr) GetMarketEndTime ()  (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetMarketEndTime")
}

func (c *CpCodeMgr) IsFrnMember(code string)  (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "IsFrnMember", code)
}

func (c *CpCodeMgr) GetTickUnit (code string)  (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetTickUnit", code)
}

func (c *CpCodeMgr) GetTickValue (code string)  (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetTickValue", code)
}

func (c *CpCodeMgr) OvFutGetAllCodeList( )  (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "OvFutGetAllCodeList")
}

func (c *CpCodeMgr) OvFutGetExchList( )  (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "OvFutGetExchList")
}

func (c *CpCodeMgr) OvFutCodeToName(code string)  (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "OvFutCodeToName", code)
}

func (c *CpCodeMgr) OvFutGetExchCode(code string)  (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "OvFutGetExchCode", code)
}

func (c *CpCodeMgr) OvFutGetLastTradeDate(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "OvFutGetLastTradeDate", code)
}

func (c *CpCodeMgr) OvFutGetProdCode(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "OvFutGetProdCode", code)
}

func (c *CpCodeMgr) GetStartTime(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStartTime", code)
}

func (c *CpCodeMgr) GetEndTime(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetEndTime", code)
}

func (c *CpCodeMgr) IsTradeCondition(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "IsTradeCondition", code)
}

func (c *CpCodeMgr) GetStockFutureList( ) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockFutureList")
}

func (c *CpCodeMgr) GetStockFutureBaseList( ) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockFutureBaseList")
}

func (c *CpCodeMgr) GetStockFutureListByBaseCode(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockFutureListByBaseCode", code)
}

func (c *CpCodeMgr) GetStockFutureBaseCode(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.obj, "GetStockFutureBaseCode", code)
}
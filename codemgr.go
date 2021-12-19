package creon
// codeMgr 함수
import (
	ole "github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

type CpCodeMgr struct {
	CpClass
}

func (c *CpCodeMgr) Create() {
	c.CpClass.Create("CpUtil.CpCodeMgr")
}

func (c *CpCodeMgr) CodeToName(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "CodeToName", code)
}

func (c *CpCodeMgr) GetStockMarginRate(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetStockMarginRate", code)
}

func (c *CpCodeMgr) GetStockMemeMin(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetStockMemeMin", code)
}

func (c *CpCodeMgr) GetStockMarketKind(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetStockMarketKind", code)
}

func (c *CpCodeMgr) GetStockIndustryCode(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetStockIndustryCode", code)
}

func (c *CpCodeMgr) IsSPAC(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "IsSPAC", code)
}

func (c *CpCodeMgr) GetStockControlKind(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetStockControlKind", code)
}

func (c *CpCodeMgr) GetOverHeating(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetOverHeating", code)
}


func (c *CpCodeMgr) IsStockArrgSby(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "IsStockArrgSby", code)
}

func (c *CpCodeMgr) IsStockIoi(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "IsStockIoi", code)
}

func (c *CpCodeMgr) GetStockSupervisionKind (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetStockSupervisionKind", code)
}

func (c *CpCodeMgr) GetStockStatusKind (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetStockStatusKind", code)
}

func (c *CpCodeMgr) GetStockCapital (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetStockCapital", code)
}

func (c *CpCodeMgr) GetStockFiscalMonth (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetStockFiscalMonth", code)
}

func (c *CpCodeMgr) GetStockGroupCode (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetStockGroupCode", code)
}

func (c *CpCodeMgr) GetStockKospi200Kind (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetStockKospi200Kind", code)
}

func (c *CpCodeMgr) GetStockSectionKind(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetStockSectionKind", code)
}

func (c *CpCodeMgr) GetStockLacKind (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetStockLacKind", code)
}

func (c *CpCodeMgr) GetStockListedDate (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetStockListedDate", code)
}

func (c *CpCodeMgr) GetStockMaxPrice (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetStockMaxPrice", code)
}

func (c *CpCodeMgr) GetStockMinPrice (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetStockMinPrice", code)
}

func (c *CpCodeMgr) GetStockParPrice (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetStockParPrice", code)
}

func (c *CpCodeMgr) GetStockStdPrice (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetStockStdPrice", code)
}

func (c *CpCodeMgr) GetStockYdOpenPrice (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetStockYdOpenPrice", code)
}

func (c *CpCodeMgr) GetStockYdHighPrice (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetStockYdHighPrice", code)
}

func (c *CpCodeMgr) GetStockYdLowPrice (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetStockYdLowPrice", code)
}

func (c *CpCodeMgr) GetStockYdClosePrice (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetStockYdClosePrice", code)
}

func (c *CpCodeMgr) IsStockCreditEnable(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "IsStockCreditEnable", code)
}

func (c *CpCodeMgr) GetStockParPriceChageType (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetStockParPriceChageType", code)
}

func (c *CpCodeMgr) GetMiniFutureList() (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetMiniFutureList")
}

func (c *CpCodeMgr) GetMiniOptionList() (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetMiniOptionList")
}

func (c *CpCodeMgr) ReLoadPortData() (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "ReLoadPortData")
}

func (c *CpCodeMgr) IsBigListingStock(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "IsBigListingStock", code)
}

func (c *CpCodeMgr) GetStockElwBasketCodeList(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetStockElwBasketCodeList", code)
}

func (c *CpCodeMgr) GetStockElwBasketCompList(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetStockElwBasketCompList", code)
}

func (c *CpCodeMgr) GetStockListByMarket(marketcode int) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetStockListByMarket", marketcode)
}

func (c *CpCodeMgr) GetGroupCodeList(code int) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetGroupCodeList", code)
}

func (c *CpCodeMgr) GetGroupName (code int) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetGroupName", code)
}

func (c *CpCodeMgr) GetIndustryList () (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetIndustryList")
}

func (c *CpCodeMgr) GetIndustryName (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetIndustryName", code)
}

func (c *CpCodeMgr) GetMemberList ( ) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetMemberList")
}

func (c *CpCodeMgr) GetMemberName (code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetMemberName", code)
}

func (c *CpCodeMgr) GetMarketStartTime ()  (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetMarketStartTime")
}

func (c *CpCodeMgr) GetMarketEndTime ()  (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetMarketEndTime")
}

func (c *CpCodeMgr) IsFrnMember(code string)  (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "IsFrnMember", code)
}

func (c *CpCodeMgr) GetTickUnit (code string)  (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetTickUnit", code)
}

func (c *CpCodeMgr) GetTickValue (code string)  (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetTickValue", code)
}

func (c *CpCodeMgr) OvFutGetAllCodeList( )  (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "OvFutGetAllCodeList")
}

func (c *CpCodeMgr) OvFutGetExchList( )  (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "OvFutGetExchList")
}

func (c *CpCodeMgr) OvFutCodeToName(code string)  (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "OvFutCodeToName", code)
}

func (c *CpCodeMgr) OvFutGetExchCode(code string)  (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "OvFutGetExchCode", code)
}

func (c *CpCodeMgr) OvFutGetLastTradeDate(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "OvFutGetLastTradeDate", code)
}

func (c *CpCodeMgr) OvFutGetProdCode(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "OvFutGetProdCode", code)
}

func (c *CpCodeMgr) GetStartTime(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetStartTime", code)
}

func (c *CpCodeMgr) GetEndTime(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetEndTime", code)
}

func (c *CpCodeMgr) IsTradeCondition(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "IsTradeCondition", code)
}

func (c *CpCodeMgr) GetStockFutureList( ) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetStockFutureList")
}

func (c *CpCodeMgr) GetStockFutureBaseList( ) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetStockFutureBaseList")
}

func (c *CpCodeMgr) GetStockFutureListByBaseCode(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetStockFutureListByBaseCode", code)
}

func (c *CpCodeMgr) GetStockFutureBaseCode(code string) (r *ole.VARIANT) {
	return oleutil.MustCallMethod(c.Obj, "GetStockFutureBaseCode", code)
}
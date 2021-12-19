package creon

type StockChart struct {
	CpClass
}

func (c *StockChart) Create() {
	c.CpClass.Create("CpSysDib.StockChart")
}

type StockMst struct {
	CpClass
}

func (c *StockMst) Create() {
	c.CpClass.Create("DsCbo1.StockMst")
}

type StockMstM struct {
	CpClass
}

func (c *StockMstM) Create() {
	c.CpClass.Create("DsCbo1.StockMstM")
}

type StockMst2 struct {
	CpClass
}

func (c *StockMst2) Create() {
	c.CpClass.Create("DsCbo1.StockMst2")
}

type CpTd6033 struct {
	CpClass
}

func (c *CpTd6033) Create() {
	c.CpClass.Create("CpTrade.CpTd6033")
}

type MarketEye struct {
	CpClass
}

func (c *MarketEye) Create() {
	c.CpClass.Create("CpSysDib.MarketEye")
}
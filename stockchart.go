package creon

type StockChart struct {
	CpClass
}

func (c *StockChart) Create() {
	c.CpClass.Create("CpSysDib.StockChart")
}
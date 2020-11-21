package creon

type StockMst struct {
	CpClass
}

func (c *StockMst) Create() {
	c.CpClass.Create("DsCbo1.StockMst")
}
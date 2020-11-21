package creon

import (
	ole "github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)
type StockMst struct {
	CpClass
}

func (c *StockMst) Create() {
	c.CpClass.Create("DsCbo1.StockMst")
}
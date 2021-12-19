package creon

import (
	"errors"
	"fmt"
)
type CpTd0311 struct {
	CpClass
}

func (c *CpTd0311) Create() {
	c.CpClass.Create("CpTrade.CpTd0311")
}

func (c *CpTd0311) Order_Base(acc, accFlag, code string,  vol int64, price int64) {
	c.SetInputValue(1, acc)
	c.SetInputValue(2, accFlag)
	c.SetInputValue(3, code)
	c.SetInputValue(4, vol)
	c.SetInputValue(5, price)
}

func (c *CpTd0311) Sell_Base(acc, accFlag, code string,  vol int64, price int64) {
	c.SetInputValue(0, "1")
	c.Order_Base(acc, accFlag, code, vol, price)
}

func (c *CpTd0311)Sell(acc, accFlag, code string,  vol int64, price int64, order_cond string, hoga_cond string) (err error) {
	c.Sell_Base(acc, accFlag, code, vol, price)
	c.SetInputValue(6, order_cond)
	c.SetInputValue(7, hoga_cond)
	err = c.Check_Order_Status()
	return
}

func (c *CpTd0311) Sell_at_Price(acc, accFlag, code string,  vol int64, price int64) (err error) {
	c.Sell_Base(acc, accFlag, code, vol, price)
	c.SetInputValue(6, "0")
	c.SetInputValue(7, "01")
	err = c.Check_Order_Status()
	return
}

func (c *CpTd0311) Sell_at_Price_with_Ioc(acc, accFlag, code string,  vol int64, price int64) (err error) {
	c.Sell_Base(acc, accFlag, code, vol, price)
	c.SetInputValue(6, "1")
	c.SetInputValue(7, "01")
	err = c.Check_Order_Status()
	return
}

func (c *CpTd0311) Sell_at_Price_with_Fok(acc, accFlag, code string,  vol int64, price int64) (err error) {
	c.Sell_Base(acc, accFlag, code, vol, price)
	c.SetInputValue(6, "2")
	c.SetInputValue(7, "01")
	err = c.Check_Order_Status()
	return
}

func (c *CpTd0311) Buy_Base(acc, accFlag, code string,  vol int64, price int64) {
	c.SetInputValue(0, "2")
	c.Order_Base(acc, accFlag, code, vol, price)
}

func (c *CpTd0311)Buy(acc, accFlag, code string,  vol int64, price int64, order_cond string, hoga_cond string) (err error) {
	c.Sell_Base(acc, accFlag, code, vol, price)
	c.SetInputValue(6, order_cond)
	c.SetInputValue(7, hoga_cond)
	err = c.Check_Order_Status()
	return
}

func (c *CpTd0311) Buy_at_Price(acc, accFlag, code string,  vol int64, price int64) (err error) {
	c.Buy_Base(acc, accFlag, code, vol, price)
	c.SetInputValue(6, "0")
	c.SetInputValue(7, "01")
	err = c.Check_Order_Status()
	return
}

func (c *CpTd0311) Buy_at_Price_with_Ioc(acc, accFlag, code string,  vol int64, price int64) (err error) {
	c.Buy_Base(acc, accFlag, code, vol, price)
	c.SetInputValue(6, "1")
	c.SetInputValue(7, "01")
	err = c.Check_Order_Status()
	return
}

func (c *CpTd0311) Buy_at_Price_with_Fok(acc, accFlag, code string,  vol int64, price int64) (err error) {
	c.Buy_Base(acc, accFlag, code, vol, price)
	c.SetInputValue(6, "2")
	c.SetInputValue(7, "01")
	err = c.Check_Order_Status()
	return
}

func (c * CpTd0311) Check_Order_Status() (err error) {
	ret := Int(c.BlockRequest())
	if ret != 0 {
		err_txt := fmt.Sprintf("주문 오류 #%d", ret)
		err = errors.New(err_txt)
		return
	}
	rqStatus := Int(c.GetDibStatus())
	errMsg := String(c.GetDibMsg1())
	if rqStatus != 0 {
		err_txt := fmt.Sprintf("주문실패 : %d - %s", rqStatus, errMsg)
		err = errors.New(err_txt)
		return
	}
	err = nil
	return
}
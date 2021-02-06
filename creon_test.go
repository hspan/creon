package creon

import (
	"testing"
	ole "github.com/hspan/go-ole"
	"sort"
)

var boolean  map[bool]string
func TestCoInitializeEx(t *testing.T) {
	ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED)
	boolean = make(map[bool]string)
	boolean[true] = "true"
	boolean[false] = "false"
}

func TestCodeMgr(t *testing.T) {
	codeMgr := CpCodeMgr{}
	codeMgr.Create()
	defer codeMgr.Release()
	list := RetSS(codeMgr.GetStockListByMarket(1))
	sort.Strings(list)
	if list[0] != "A000020" {
		t.Errorf("error!! need A000020 but %s return", list[0])
	}
	Name := RetStr(codeMgr.CodeToName("A000020"))
	if Name != "동화약품" {
		t.Errorf("error!! need 동화약품 but %s return", Name)
	}
	market := RetInt(codeMgr.GetStockMarketKind("A000020"))
	if market != 1 {
		t.Errorf("error!! need 1 but %d return", market)
	}
	section := RetInt(codeMgr.GetStockSectionKind("A000020"))
	if market != 1 {
		t.Errorf("error!! need 1(CPC_KSE_SECTION_KIND_ST) but %d return", section)
	}
	ok := RetBool(codeMgr.IsSPAC("A000020"))
	if ok != false {
		t.Errorf("error!! need false but %s return", boolean[ok])
	}
	ldate := RetLong(codeMgr.GetStockListedDate("A000020"))
	if ldate != 19760324 {
		t.Errorf("error!! need false but %d return", ldate)
	}
}

func TestCybos(t *testing.T) {
	cybos := CpCybos{}
	cybos.Create()
	if ! RetBool(cybos.IsConnect()) {
		t.Error("연결에러")
	}
}
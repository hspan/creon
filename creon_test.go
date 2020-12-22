package creon

import (
	"testing"
	ole "github.com/hspan/go-ole"
	"fmt"
	"sort"
)

func TestCoInitializeEx(t *testing.T) {
	ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED)
}

func RetSS(r *ole.VARIANT) (ret []string) {
	data := r.ToArray().ToValueArray()
	for _, value := range data {
		strValue := fmt.Sprintf("%s", value)
		ret = append(ret, strValue)
	}
	return 
}
func TestCodeMgr(t *testing.T) {
	ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED)
	codeMgr := CpCodeMgr{}
	codeMgr.Create()
	defer codeMgr.Release()
	list := RetSS(codeMgr.GetStockListByMarket(1))
	sort.Strings(list)
	if list[0] != "A000020" {
		t.Errorf("error!! need A000020 but %s return", list[0])
	}
}
package creon

import (
	ole "github.com/go-ole/go-ole"
	"syscall"
	"unsafe"
	"strings"
	"fmt"
)

var (
	modoleaut32, _ = syscall.LoadDLL("oleaut32.dll")
	procSafeArrayCreateVector, _      = modoleaut32.FindProc("SafeArrayCreateVector")
	procSafeArrayPutElement, _        = modoleaut32.FindProc("SafeArrayPutElement")
)

// DISPPARAMS are the arguments that passed to methods or property.
type DISPPARAMS struct {
	rgvarg            uintptr
	rgdispidNamedArgs uintptr
	cArgs             uint32
	cNamedArgs        uint32
}
// EXCEPINFO defines exception info.
type EXCEPINFO struct {
	wCode             uint16
	wReserved         uint16
	bstrSource        *uint16
	bstrDescription   *uint16
	bstrHelpFile      *uint16
	dwHelpContext     uint32
	pvReserved        uintptr
	pfnDeferredFillIn uintptr
	scode             uint32
}
// Error implements error interface and returns error string.
func (e EXCEPINFO) Error() string {
	if e.bstrDescription != nil {
		return strings.TrimSpace(ole.BstrToString(e.bstrDescription))
	}

	src := "Unknown"
	if e.bstrSource != nil {
		src = ole.BstrToString(e.bstrSource)
	}

	code := e.scode
	if e.wCode != 0 {
		code = uint32(e.wCode)
	}

	return fmt.Sprintf("%v: %#x", src, code)
}

// safeArrayCreateVector creates SafeArray.
//
// AKA: SafeArrayCreateVector in Windows API.
func safeArrayCreateVector(variantType ole.VT, lowerBound int32, length uint32) (safearray *ole.SafeArray, err error) {
	sa, _, err := procSafeArrayCreateVector.Call(
		uintptr(variantType),
		uintptr(lowerBound),
		uintptr(length))
	safearray = (*ole.SafeArray)(unsafe.Pointer(sa))
	return
}

// convertHresultToError converts syscall to error, if call is unsuccessful.
func convertHresultToError(hr uintptr, r2 uintptr, ignore error) (err error) {
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

// safeArrayPutElement stores the data element at the specified location in the
// array.
//
// AKA: SafeArrayPutElement in Windows API.
func safeArrayPutElement(safearray *ole.SafeArray, index int64, element uintptr) (err error) {
	err = convertHresultToError(
		procSafeArrayPutElement.Call(
			uintptr(unsafe.Pointer(safearray)),
			uintptr(unsafe.Pointer(&index)),
			uintptr(unsafe.Pointer(element))))
	return
}

func safeArrayFromInt32Slice(slice []int32) *ole.SafeArray {
	array, _ := safeArrayCreateVector(ole.VT_I4, 0, uint32(len(slice)))

	if array == nil {
		panic("Could not convert []int32 to SAFEARRAY")
	}
	// SysAllocStringLen(s)
	for i, v := range slice {
		safeArrayPutElement(array, int64(i), uintptr(unsafe.Pointer(&v)))
	}
	return array
}

func invoke(v *CpClass, dispid int32, dispatch int16, params ...interface{}) (result *ole.VARIANT, err error) {
	var dispparams DISPPARAMS
	disp := v.Obj
	if dispatch&ole.DISPATCH_PROPERTYPUT != 0 {
		dispnames := [1]int32{ole.DISPID_PROPERTYPUT}
		dispparams.rgdispidNamedArgs = uintptr(unsafe.Pointer(&dispnames[0]))
		dispparams.cNamedArgs = 1
	}
	var vargs []ole.VARIANT
	if len(params) > 0 {
		vargs = make([]ole.VARIANT, len(params))
		ole.VariantInit(&vargs[1])
		vargs[1] = ole.NewVariant(ole.VT_I4, int64(params[0].(int)))
		ole.VariantInit(&vargs[0])
		safeByteArray := safeArrayFromInt32Slice(params[1].([]int32))
		vargs[0] = ole.NewVariant(ole.VT_ARRAY|ole.VT_I4, int64(uintptr(unsafe.Pointer(safeByteArray))))
		defer ole.VariantClear(&vargs[0])	
		dispparams.rgvarg = uintptr(unsafe.Pointer(&vargs[0]))
		dispparams.cArgs = uint32(len(params))
	}

	result = new(ole.VARIANT)
	var excepInfo EXCEPINFO
	ole.VariantInit(result)
	hr, _, _ := syscall.Syscall9(
		disp.VTable().Invoke,
		9,
		uintptr(unsafe.Pointer(disp)),
		uintptr(dispid),
		uintptr(unsafe.Pointer(ole.IID_NULL)),
		uintptr(ole.GetUserDefaultLCID()),
		uintptr(dispatch),
		uintptr(unsafe.Pointer(&dispparams)),
		uintptr(unsafe.Pointer(result)),
		uintptr(unsafe.Pointer(&excepInfo)),
		0)
	if hr != 0 {
		err = ole.NewErrorWithSubError(hr, ole.BstrToString(excepInfo.bstrDescription), excepInfo)
	}
	for _, varg := range vargs {
		if varg.VT == ole.VT_BSTR && varg.Val != 0 {
			ole.SysFreeString(((*int16)(unsafe.Pointer(uintptr(varg.Val)))))
		}
		/*
			if varg.VT == (ole.VT_BSTR|ole.VT_BYREF) && varg.Val != 0 {
				*(params[n].(*string)) = LpOleStrToString((*uint16)(unsafe.Pointer(uintptr(varg.Val))))
				println(*(params[n].(*string)))
				fmt.Fprintln(os.Stderr, *(params[n].(*string)))
			}
		*/
	}
	return
}

func (v *CpClass) Invoke(dispid int32, dispatch int16, params ...interface{}) (result *ole.VARIANT, err error) {
	result, err = invoke(v, dispid, dispatch, params...)
	return
}

func getIDsOfName(disp *ole.IDispatch, names []string) (dispid []int32, err error) {
	wnames := make([]*uint16, len(names))
	for i := 0; i < len(names); i++ {
		wnames[i] = syscall.StringToUTF16Ptr(names[i])
	}
	dispid = make([]int32, len(names))
	namelen := uint32(len(names))
	hr, _, _ := syscall.Syscall6(
		disp.VTable().GetIDsOfNames,
		6,
		uintptr(unsafe.Pointer(disp)),
		uintptr(unsafe.Pointer(ole.IID_NULL)),
		uintptr(unsafe.Pointer(&wnames[0])),
		uintptr(namelen),
		uintptr(ole.GetUserDefaultLCID()),
		uintptr(unsafe.Pointer(&dispid[0])))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *CpClass) GetIDsOfName(names []string) (dispid []int32, err error) {
	dispid, err = getIDsOfName(v.Obj, names)
	return
}

// GetSingleIDOfName is a helper that returns single display ID for IDispatch name.
//
// This replaces the common pattern of attempting to get a single name from the list of available
// IDs. It gives the first ID, if it is available.
func (v *CpClass) GetSingleIDOfName(name string) (displayID int32, err error) {
	var displayIDs []int32
	displayIDs, err = v.GetIDsOfName([]string{name})
	if err != nil {
		return
	}
	displayID = displayIDs[0]
	return
}

// InvokeWithOptionalArgs accepts arguments as an array, works like Invoke.
//
// Accepts name and will attempt to retrieve Display ID to pass to Invoke.
//
// Passing params as an array is a workaround that could be fixed in later versions of Go that
// prevent passing empty params. During testing it was discovered that this is an acceptable way of
// getting around not being able to pass params normally.
func (v *CpClass) InvokeWithOptionalArgs(name string, dispatch int16, params []interface{}) (result *ole.VARIANT, err error) {
	displayID, err := v.GetSingleIDOfName(name)
	if err != nil {
		return
	}

	if len(params) < 1 {
		result, err = v.Invoke(displayID, dispatch)
	} else {
		result, err = v.Invoke(displayID, dispatch, params...)
	}

	return
}


// CallMethod calls method on IDispatch with parameters.
func CallMethod(cp *CpClass, name string, params ...interface{}) (result *ole.VARIANT, err error) {
	return cp.InvokeWithOptionalArgs(name, ole.DISPATCH_METHOD, params)
}

// MustCallMethod calls method on IDispatch with parameters or panics.
func MustCallMethod(cp *CpClass, name string, params ...interface{}) (result *ole.VARIANT) {
	r, err := CallMethod(cp, name, params...)
	if err != nil {
		panic(err.Error())
	}
	return r
}
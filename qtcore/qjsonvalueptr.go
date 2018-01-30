package qtcore

// /usr/include/qt/QtCore/qjsonvalue.h
// #include <qjsonvalue.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
type QJsonValuePtr struct {
	*qtrt.CObject
}

func (this *QJsonValuePtr) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QJsonValuePtr) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQJsonValuePtrFromPointer(cthis unsafe.Pointer) *QJsonValuePtr {
	return &QJsonValuePtr{&qtrt.CObject{cthis}}
}
func (*QJsonValuePtr) NewFromPointer(cthis unsafe.Pointer) *QJsonValuePtr {
	return NewQJsonValuePtrFromPointer(cthis)
}

// /usr/include/qt/QtCore/qjsonvalue.h:226
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QJsonValuePtr(const QJsonValue &)
func NewQJsonValuePtr(val *QJsonValue) *QJsonValuePtr {
	var convArg0 = val.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonValuePtrC2ERK10QJsonValue", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonValuePtrFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

//  body block end

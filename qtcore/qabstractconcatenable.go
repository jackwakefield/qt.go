package qtcore

// /usr/include/qt/QtCore/qstringbuilder.h
// #include <qstringbuilder.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
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
type QAbstractConcatenable struct {
	*qtrt.CObject
}

func (this *QAbstractConcatenable) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAbstractConcatenable) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQAbstractConcatenableFromPointer(cthis unsafe.Pointer) *QAbstractConcatenable {
	return &QAbstractConcatenable{&qtrt.CObject{cthis}}
}
func (*QAbstractConcatenable) NewFromPointer(cthis unsafe.Pointer) *QAbstractConcatenable {
	return NewQAbstractConcatenableFromPointer(cthis)
}

// /usr/include/qt/QtCore/qstringbuilder.h:61
// index:0
// Protected static Visibility=Default Availability=Available
// [-2] void convertFromAscii(const char *, int, QChar *&)
func (this *QAbstractConcatenable) ConvertFromAscii(a string, len int, out *QChar) {
	var convArg0 = qtrt.CString(a)
	defer qtrt.FreeMem(convArg0)
	var convArg2 = out.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QAbstractConcatenable16convertFromAsciiEPKciRP5QChar", ffiqt.FFI_TYPE_POINTER, convArg0, len, convArg2)
	gopp.ErrPrint(err, rv)
}
func QAbstractConcatenable_ConvertFromAscii(a string, len int, out *QChar) {
	var nilthis *QAbstractConcatenable
	nilthis.ConvertFromAscii(a, len, out)
}

// /usr/include/qt/QtCore/qstringbuilder.h:62
// index:1
// Protected static inline Visibility=Default Availability=Available
// [-2] void convertFromAscii(char, QChar *&)
func (this *QAbstractConcatenable) ConvertFromAscii_1(a byte, out *QChar) {
	var convArg1 = out.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QAbstractConcatenable16convertFromAsciiEcRP5QChar", ffiqt.FFI_TYPE_POINTER, a, convArg1)
	gopp.ErrPrint(err, rv)
}
func QAbstractConcatenable_ConvertFromAscii_1(a byte, out *QChar) {
	var nilthis *QAbstractConcatenable
	nilthis.ConvertFromAscii_1(a, out)
}

// /usr/include/qt/QtCore/qstringbuilder.h:66
// index:0
// Protected static Visibility=Default Availability=Available
// [-2] void appendLatin1To(const char *, int, QChar *)
func (this *QAbstractConcatenable) AppendLatin1To(a string, len int, out *QChar /*777 QChar **/) {
	var convArg0 = qtrt.CString(a)
	defer qtrt.FreeMem(convArg0)
	var convArg2 = out.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QAbstractConcatenable14appendLatin1ToEPKciP5QChar", ffiqt.FFI_TYPE_POINTER, convArg0, len, convArg2)
	gopp.ErrPrint(err, rv)
}
func QAbstractConcatenable_AppendLatin1To(a string, len int, out *QChar /*777 QChar **/) {
	var nilthis *QAbstractConcatenable
	nilthis.AppendLatin1To(a, len, out)
}

//  body block end

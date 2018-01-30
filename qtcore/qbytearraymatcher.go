package qtcore

// /usr/include/qt/QtCore/qbytearraymatcher.h
// #include <qbytearraymatcher.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
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
type QByteArrayMatcher struct {
	*qtrt.CObject
}

func (this *QByteArrayMatcher) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QByteArrayMatcher) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQByteArrayMatcherFromPointer(cthis unsafe.Pointer) *QByteArrayMatcher {
	return &QByteArrayMatcher{&qtrt.CObject{cthis}}
}
func (*QByteArrayMatcher) NewFromPointer(cthis unsafe.Pointer) *QByteArrayMatcher {
	return NewQByteArrayMatcherFromPointer(cthis)
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:53
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QByteArrayMatcher()
func NewQByteArrayMatcher() *QByteArrayMatcher {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QByteArrayMatcherC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQByteArrayMatcherFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:54
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QByteArrayMatcher(const QByteArray &)
func NewQByteArrayMatcher_1(pattern *QByteArray) *QByteArrayMatcher {
	var convArg0 = pattern.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QByteArrayMatcherC2ERK10QByteArray", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQByteArrayMatcherFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:55
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QByteArrayMatcher(const char *, int)
func NewQByteArrayMatcher_2(pattern string, length int) *QByteArrayMatcher {
	var convArg0 = qtrt.CString(pattern)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QByteArrayMatcherC2EPKci", ffiqt.FFI_TYPE_POINTER, convArg0, length)
	gopp.ErrPrint(err, rv)
	gothis := NewQByteArrayMatcherFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QByteArrayMatcher()
func DeleteQByteArrayMatcher(*QByteArrayMatcher) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QByteArrayMatcherD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPattern(const QByteArray &)
func (this *QByteArrayMatcher) SetPattern(pattern *QByteArray) {
	var convArg0 = pattern.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QByteArrayMatcher10setPatternERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:63
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexIn(const QByteArray &, int)
func (this *QByteArrayMatcher) IndexIn(ba *QByteArray, from int) int {
	var convArg0 = ba.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QByteArrayMatcher7indexInERK10QByteArrayi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:64
// index:1
// Public Visibility=Default Availability=Available
// [4] int indexIn(const char *, int, int)
func (this *QByteArrayMatcher) IndexIn_1(str string, len int, from int) int {
	var convArg0 = qtrt.CString(str)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QByteArrayMatcher7indexInEPKcii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len, from)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:65
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray pattern()
func (this *QByteArrayMatcher) Pattern() *QByteArray /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QByteArrayMatcher7patternEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end

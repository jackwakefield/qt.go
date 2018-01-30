package qtnetwork

// /usr/include/qt/QtNetwork/qdnslookup.h
// #include <qdnslookup.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
type QDnsServiceRecord struct {
	*qtrt.CObject
}

func (this *QDnsServiceRecord) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QDnsServiceRecord) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQDnsServiceRecordFromPointer(cthis unsafe.Pointer) *QDnsServiceRecord {
	return &QDnsServiceRecord{&qtrt.CObject{cthis}}
}
func (*QDnsServiceRecord) NewFromPointer(cthis unsafe.Pointer) *QDnsServiceRecord {
	return NewQDnsServiceRecordFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDnsServiceRecord()
func NewQDnsServiceRecord() *QDnsServiceRecord {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDnsServiceRecordC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQDnsServiceRecordFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtNetwork/qdnslookup.h:142
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QDnsServiceRecord()
func DeleteQDnsServiceRecord(*QDnsServiceRecord) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDnsServiceRecordD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:144
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QDnsServiceRecord &)
func (this *QDnsServiceRecord) Swap(other *QDnsServiceRecord) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDnsServiceRecord4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:146
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name()
func (this *QDnsServiceRecord) Name() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QDnsServiceRecord4nameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qdnslookup.h:147
// index:0
// Public Visibility=Default Availability=Available
// [2] quint16 port()
func (this *QDnsServiceRecord) Port() uint16 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QDnsServiceRecord4portEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint16(rv) // 222
}

// /usr/include/qt/QtNetwork/qdnslookup.h:148
// index:0
// Public Visibility=Default Availability=Available
// [2] quint16 priority()
func (this *QDnsServiceRecord) Priority() uint16 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QDnsServiceRecord8priorityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint16(rv) // 222
}

// /usr/include/qt/QtNetwork/qdnslookup.h:149
// index:0
// Public Visibility=Default Availability=Available
// [8] QString target()
func (this *QDnsServiceRecord) Target() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QDnsServiceRecord6targetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qdnslookup.h:150
// index:0
// Public Visibility=Default Availability=Available
// [4] quint32 timeToLive()
func (this *QDnsServiceRecord) TimeToLive() uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QDnsServiceRecord10timeToLiveEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qdnslookup.h:151
// index:0
// Public Visibility=Default Availability=Available
// [2] quint16 weight()
func (this *QDnsServiceRecord) Weight() uint16 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QDnsServiceRecord6weightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint16(rv) // 222
}

//  body block end

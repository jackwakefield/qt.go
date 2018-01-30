package qtnetwork

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h
// #include <qabstractnetworkcache.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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
type QAbstractNetworkCache struct {
	*qtcore.QObject
}

func (this *QAbstractNetworkCache) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QAbstractNetworkCache) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQAbstractNetworkCacheFromPointer(cthis unsafe.Pointer) *QAbstractNetworkCache {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QAbstractNetworkCache{bcthis0}
}
func (*QAbstractNetworkCache) NewFromPointer(cthis unsafe.Pointer) *QAbstractNetworkCache {
	return NewQAbstractNetworkCacheFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:116
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QAbstractNetworkCache) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QAbstractNetworkCache10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:119
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractNetworkCache()
func DeleteQAbstractNetworkCache(*QAbstractNetworkCache) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QAbstractNetworkCacheD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:121
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QNetworkCacheMetaData metaData(const QUrl &)
func (this *QAbstractNetworkCache) MetaData(url *qtcore.QUrl) *QNetworkCacheMetaData /*123*/ {
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QAbstractNetworkCache8metaDataERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQNetworkCacheMetaDataFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:122
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void updateMetaData(const QNetworkCacheMetaData &)
func (this *QAbstractNetworkCache) UpdateMetaData(metaData *QNetworkCacheMetaData) {
	var convArg0 = metaData.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QAbstractNetworkCache14updateMetaDataERK21QNetworkCacheMetaData", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:123
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QIODevice * data(const QUrl &)
func (this *QAbstractNetworkCache) Data(url *qtcore.QUrl) *qtcore.QIODevice /*777 QIODevice **/ {
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QAbstractNetworkCache4dataERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:124
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool remove(const QUrl &)
func (this *QAbstractNetworkCache) Remove(url *qtcore.QUrl) bool {
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QAbstractNetworkCache6removeERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:125
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] qint64 cacheSize()
func (this *QAbstractNetworkCache) CacheSize() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QAbstractNetworkCache9cacheSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:127
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QIODevice * prepare(const QNetworkCacheMetaData &)
func (this *QAbstractNetworkCache) Prepare(metaData *QNetworkCacheMetaData) *qtcore.QIODevice /*777 QIODevice **/ {
	var convArg0 = metaData.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QAbstractNetworkCache7prepareERK21QNetworkCacheMetaData", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:131
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void clear()
func (this *QAbstractNetworkCache) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QAbstractNetworkCache5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:134
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QAbstractNetworkCache(QObject *)
func NewQAbstractNetworkCache(parent *qtcore.QObject /*777 QObject **/) *QAbstractNetworkCache {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QAbstractNetworkCacheC1EP7QObject", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractNetworkCacheFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

//  body block end

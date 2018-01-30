package qtgui

// /usr/include/qt/QtGui/qpolygon.h
// #include <qpolygon.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
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
type QPolygonF struct {
	*qtrt.CObject
}

func (this *QPolygonF) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPolygonF) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQPolygonFFromPointer(cthis unsafe.Pointer) *QPolygonF {
	return &QPolygonF{&qtrt.CObject{cthis}}
}
func (*QPolygonF) NewFromPointer(cthis unsafe.Pointer) *QPolygonF {
	return NewQPolygonFFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpolygon.h:144
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QPolygonF()
func NewQPolygonF() *QPolygonF {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPolygonFC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qpolygon.h:146
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QPolygonF(int)
func NewQPolygonF_1(size int) *QPolygonF {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPolygonFC2Ei", ffiqt.FFI_TYPE_POINTER, size)
	gopp.ErrPrint(err, rv)
	gothis := NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qpolygon.h:151
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QPolygonF(const QRectF &)
func NewQPolygonF_2(r *qtcore.QRectF) *QPolygonF {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPolygonFC2ERK6QRectF", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qpolygon.h:152
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QPolygonF(const QPolygon &)
func NewQPolygonF_3(a *QPolygon) *QPolygonF {
	var convArg0 = a.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPolygonFC2ERK8QPolygon", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qpolygon.h:145
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void ~QPolygonF()
func DeleteQPolygonF(*QPolygonF) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPolygonFD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:159
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QPolygonF &)
func (this *QPolygonF) Swap(other *QPolygonF) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPolygonF4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:163
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void translate(qreal, qreal)
func (this *QPolygonF) Translate(dx float64, dy float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPolygonF9translateEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:164
// index:1
// Public Visibility=Default Availability=Available
// [-2] void translate(const QPointF &)
func (this *QPolygonF) Translate_1(offset *qtcore.QPointF) {
	var convArg0 = offset.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPolygonF9translateERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:166
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPolygonF translated(qreal, qreal)
func (this *QPolygonF) Translated(dx float64, dy float64) *QPolygonF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPolygonF10translatedEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:167
// index:1
// Public Visibility=Default Availability=Available
// [8] QPolygonF translated(const QPointF &)
func (this *QPolygonF) Translated_1(offset *qtcore.QPointF) *QPolygonF /*123*/ {
	var convArg0 = offset.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPolygonF10translatedERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:169
// index:0
// Public Visibility=Default Availability=Available
// [8] QPolygon toPolygon()
func (this *QPolygonF) ToPolygon() *QPolygon /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPolygonF9toPolygonEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:171
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isClosed()
func (this *QPolygonF) IsClosed() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPolygonF8isClosedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpolygon.h:173
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF boundingRect()
func (this *QPolygonF) BoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPolygonF12boundingRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:175
// index:0
// Public Visibility=Default Availability=Available
// [1] bool containsPoint(const QPointF &, Qt::FillRule)
func (this *QPolygonF) ContainsPoint(pt *qtcore.QPointF, fillRule int) bool {
	var convArg0 = pt.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPolygonF13containsPointERK7QPointFN2Qt8FillRuleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, fillRule)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpolygon.h:177
// index:0
// Public Visibility=Default Availability=Available
// [8] QPolygonF united(const QPolygonF &)
func (this *QPolygonF) United(r *QPolygonF) *QPolygonF /*123*/ {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPolygonF6unitedERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:178
// index:0
// Public Visibility=Default Availability=Available
// [8] QPolygonF intersected(const QPolygonF &)
func (this *QPolygonF) Intersected(r *QPolygonF) *QPolygonF /*123*/ {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPolygonF11intersectedERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:179
// index:0
// Public Visibility=Default Availability=Available
// [8] QPolygonF subtracted(const QPolygonF &)
func (this *QPolygonF) Subtracted(r *QPolygonF) *QPolygonF /*123*/ {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPolygonF10subtractedERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:181
// index:0
// Public Visibility=Default Availability=Available
// [1] bool intersects(const QPolygonF &)
func (this *QPolygonF) Intersects(r *QPolygonF) bool {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPolygonF10intersectsERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end

package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
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
type QHelpEvent struct {
	*qtcore.QEvent
}

func (this *QHelpEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QHelpEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQHelpEventFromPointer(cthis unsafe.Pointer) *QHelpEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QHelpEvent{bcthis0}
}
func (*QHelpEvent) NewFromPointer(cthis unsafe.Pointer) *QHelpEvent {
	return NewQHelpEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:680
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QHelpEvent(enum QEvent::Type, const QPoint &, const QPoint &)
func NewQHelpEvent(type_ int, pos *qtcore.QPoint, globalPos *qtcore.QPoint) *QHelpEvent {
	var convArg1 = pos.GetCthis()
	var convArg2 = globalPos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QHelpEventC2EN6QEvent4TypeERK6QPointS4_", ffiqt.FFI_TYPE_POINTER, type_, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQHelpEventFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:681
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QHelpEvent()
func DeleteQHelpEvent(*QHelpEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QHelpEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:683
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int x()
func (this *QHelpEvent) X() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QHelpEvent1xEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:684
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int y()
func (this *QHelpEvent) Y() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QHelpEvent1yEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:685
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int globalX()
func (this *QHelpEvent) GlobalX() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QHelpEvent7globalXEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:686
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int globalY()
func (this *QHelpEvent) GlobalY() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QHelpEvent7globalYEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:688
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QPoint & pos()
func (this *QHelpEvent) Pos() *qtcore.QPoint {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QHelpEvent3posEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:689
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QPoint & globalPos()
func (this *QHelpEvent) GlobalPos() *qtcore.QPoint {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QHelpEvent9globalPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

//  body block end

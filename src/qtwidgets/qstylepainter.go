//  header block begin
// /usr/include/qt/QtWidgets/qstylepainter.h
// #include <qstylepainter.h>
// #include <QtWidgets>
package qtwidgets

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
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QStylePainter struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qstylepainter.h:54
// index:0
// inline
// void QStylePainter()
func NewQStylePainter() *QStylePainter {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStylePainterC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QStylePainter{cthis}
}

// /usr/include/qt/QtWidgets/qstylepainter.h:55
// index:1
// inline
// void QStylePainter(class QWidget *)
func NewQStylePainter_1(w unsafe.Pointer) *QStylePainter {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStylePainterC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, w)
	gopp.ErrPrint(err, rv)
	return &QStylePainter{cthis}
}

// /usr/include/qt/QtWidgets/qstylepainter.h:56
// index:2
// inline
// void QStylePainter(class QPaintDevice *, class QWidget *)
func NewQStylePainter_2(pd unsafe.Pointer, w unsafe.Pointer) *QStylePainter {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStylePainterC2EP12QPaintDeviceP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, pd, w)
	gopp.ErrPrint(err, rv)
	return &QStylePainter{cthis}
}

// /usr/include/qt/QtWidgets/qstylepainter.h:57
// index:0
// inline
// bool begin(class QWidget *)
func (this *QStylePainter) Begin(w unsafe.Pointer) {
	// 0: (, QWidget * w), (w)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStylePainter5beginEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstylepainter.h:58
// index:1
// inline
// bool begin(class QPaintDevice *, class QWidget *)
func (this *QStylePainter) Begin_1(pd unsafe.Pointer, w unsafe.Pointer) {
	// 1: (, QPaintDevice * pd, QWidget * w), (pd, w)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStylePainter5beginEP12QPaintDeviceP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, pd, w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstylepainter.h:64
// index:0
// inline
// void drawPrimitive(class QStyle::PrimitiveElement, const class QStyleOption &)
func (this *QStylePainter) DrawPrimitive(pe int, opt unsafe.Pointer) {
	// 0: (, QStyle::PrimitiveElement pe, const QStyleOption & opt), (&pe, opt)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStylePainter13drawPrimitiveEN6QStyle16PrimitiveElementERK12QStyleOption", ffiqt.FFI_TYPE_VOID, this.cthis, &pe, opt)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstylepainter.h:65
// index:0
// inline
// void drawControl(class QStyle::ControlElement, const class QStyleOption &)
func (this *QStylePainter) DrawControl(ce int, opt unsafe.Pointer) {
	// 0: (, QStyle::ControlElement ce, const QStyleOption & opt), (&ce, opt)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStylePainter11drawControlEN6QStyle14ControlElementERK12QStyleOption", ffiqt.FFI_TYPE_VOID, this.cthis, &ce, opt)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstylepainter.h:66
// index:0
// inline
// void drawComplexControl(class QStyle::ComplexControl, const class QStyleOptionComplex &)
func (this *QStylePainter) DrawComplexControl(cc int, opt unsafe.Pointer) {
	// 0: (, QStyle::ComplexControl cc, const QStyleOptionComplex & opt), (&cc, opt)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStylePainter18drawComplexControlEN6QStyle14ComplexControlERK19QStyleOptionComplex", ffiqt.FFI_TYPE_VOID, this.cthis, &cc, opt)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstylepainter.h:67
// index:0
// inline
// void drawItemText(const class QRect &, int, const class QPalette &, _Bool, const class QString &, class QPalette::ColorRole)
func (this *QStylePainter) DrawItemText(r unsafe.Pointer, flags int, pal unsafe.Pointer, enabled bool, text unsafe.Pointer, textRole int) {
	// 0: (, const QRect & r, int flags, const QPalette & pal, bool enabled, const QString & text, QPalette::ColorRole textRole), (r, &flags, pal, &enabled, text, &textRole)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStylePainter12drawItemTextERK5QRectiRK8QPalettebRK7QStringNS3_9ColorRoleE", ffiqt.FFI_TYPE_VOID, this.cthis, r, &flags, pal, &enabled, text, &textRole)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstylepainter.h:69
// index:0
// inline
// void drawItemPixmap(const class QRect &, int, const class QPixmap &)
func (this *QStylePainter) DrawItemPixmap(r unsafe.Pointer, flags int, pixmap unsafe.Pointer) {
	// 0: (, const QRect & r, int flags, const QPixmap & pixmap), (r, &flags, pixmap)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStylePainter14drawItemPixmapERK5QRectiRK7QPixmap", ffiqt.FFI_TYPE_VOID, this.cthis, r, &flags, pixmap)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstylepainter.h:70
// index:0
// inline
// QStyle * style()
func (this *QStylePainter) Style() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QStylePainter5styleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
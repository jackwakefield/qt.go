package qtgui

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h
// #include <qabstracttextdocumentlayout.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
type QTextObjectInterface struct {
	*qtrt.CObject
}

func (this *QTextObjectInterface) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func NewQTextObjectInterfaceFromPointer(cthis unsafe.Pointer) *QTextObjectInterface {
	return &QTextObjectInterface{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:141
// index:0
// Public virtual
// void ~QTextObjectInterface()
func DeleteQTextObjectInterface(*QTextObjectInterface) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QTextObjectInterfaceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:142
// index:0
// Public pure virtual
// QSizeF intrinsicSize(class QTextDocument *, int, const class QTextFormat &)
func (this *QTextObjectInterface) IntrinsicSize(doc *QTextDocument /*444 QTextDocument **/, posInDocument int, format *QTextFormat) *qtcore.QSizeF /*123*/ {
	var convArg0 = doc.GetCthis()
	var convArg2 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QTextObjectInterface13intrinsicSizeEP13QTextDocumentiRK11QTextFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &posInDocument, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:143
// index:0
// Public pure virtual
// void drawObject(class QPainter *, const class QRectF &, class QTextDocument *, int, const class QTextFormat &)
func (this *QTextObjectInterface) DrawObject(painter *QPainter /*444 QPainter **/, rect *qtcore.QRectF, doc *QTextDocument /*444 QTextDocument **/, posInDocument int, format *QTextFormat) {
	var convArg0 = painter.GetCthis()
	var convArg1 = rect.GetCthis()
	var convArg2 = doc.GetCthis()
	var convArg4 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QTextObjectInterface10drawObjectEP8QPainterRK6QRectFP13QTextDocumentiRK11QTextFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, &posInDocument, convArg4)
	gopp.ErrPrint(err, rv)
}

//  body block end
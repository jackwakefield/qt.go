package qtgui

// /usr/include/qt/QtGui/qinputmethod.h
// #include <qinputmethod.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
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
type QInputMethod struct {
	*qtcore.QObject
}

func (this *QInputMethod) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QInputMethod) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQInputMethodFromPointer(cthis unsafe.Pointer) *QInputMethod {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QInputMethod{bcthis0}
}
func (*QInputMethod) NewFromPointer(cthis unsafe.Pointer) *QInputMethod {
	return NewQInputMethodFromPointer(cthis)
}

// /usr/include/qt/QtGui/qinputmethod.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QInputMethod) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qinputmethod.h:68
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform inputItemTransform()
func (this *QInputMethod) InputItemTransform() *QTransform /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod18inputItemTransformEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qinputmethod.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInputItemTransform(const QTransform &)
func (this *QInputMethod) SetInputItemTransform(transform *QTransform) {
	var convArg0 = transform.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod21setInputItemTransformERK10QTransform", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:71
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF inputItemRectangle()
func (this *QInputMethod) InputItemRectangle() *qtcore.QRectF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod18inputItemRectangleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qinputmethod.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInputItemRectangle(const QRectF &)
func (this *QInputMethod) SetInputItemRectangle(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod21setInputItemRectangleERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:75
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF cursorRectangle()
func (this *QInputMethod) CursorRectangle() *qtcore.QRectF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod15cursorRectangleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qinputmethod.h:76
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF anchorRectangle()
func (this *QInputMethod) AnchorRectangle() *qtcore.QRectF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod15anchorRectangleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qinputmethod.h:79
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF keyboardRectangle()
func (this *QInputMethod) KeyboardRectangle() *qtcore.QRectF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod17keyboardRectangleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qinputmethod.h:81
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF inputItemClipRectangle()
func (this *QInputMethod) InputItemClipRectangle() *qtcore.QRectF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod22inputItemClipRectangleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qinputmethod.h:89
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isVisible()
func (this *QInputMethod) IsVisible() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod9isVisibleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qinputmethod.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVisible(_Bool)
func (this *QInputMethod) SetVisible(visible bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod10setVisibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:92
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAnimating()
func (this *QInputMethod) IsAnimating() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod11isAnimatingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qinputmethod.h:94
// index:0
// Public Visibility=Default Availability=Available
// [8] QLocale locale()
func (this *QInputMethod) Locale() *qtcore.QLocale /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod6localeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQLocaleFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qinputmethod.h:95
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::LayoutDirection inputDirection()
func (this *QInputMethod) InputDirection() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod14inputDirectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:97
// index:0
// Public static Visibility=Default Availability=Available
// [16] QVariant queryFocusObject(Qt::InputMethodQuery, QVariant)
func (this *QInputMethod) QueryFocusObject(query int, argument *qtcore.QVariant /*123*/) *qtcore.QVariant /*123*/ {
	var convArg1 = argument.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod16queryFocusObjectEN2Qt16InputMethodQueryE8QVariant", ffiqt.FFI_TYPE_POINTER, query, convArg1)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QInputMethod_QueryFocusObject(query int, argument *qtcore.QVariant /*123*/) *qtcore.QVariant /*123*/ {
	var nilthis *QInputMethod
	rv := nilthis.QueryFocusObject(query, argument)
	return rv
}

// /usr/include/qt/QtGui/qinputmethod.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void show()
func (this *QInputMethod) Show() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod4showEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void hide()
func (this *QInputMethod) Hide() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod4hideEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reset()
func (this *QInputMethod) Reset() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod5resetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void commit()
func (this *QInputMethod) Commit() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod6commitEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void invokeAction(enum QInputMethod::Action, int)
func (this *QInputMethod) InvokeAction(a int, cursorPosition int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod12invokeActionENS_6ActionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), a, cursorPosition)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cursorRectangleChanged()
func (this *QInputMethod) CursorRectangleChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod22cursorRectangleChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void anchorRectangleChanged()
func (this *QInputMethod) AnchorRectangleChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod22anchorRectangleChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void keyboardRectangleChanged()
func (this *QInputMethod) KeyboardRectangleChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod24keyboardRectangleChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void inputItemClipRectangleChanged()
func (this *QInputMethod) InputItemClipRectangleChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod29inputItemClipRectangleChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void visibleChanged()
func (this *QInputMethod) VisibleChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod14visibleChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void animatingChanged()
func (this *QInputMethod) AnimatingChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod16animatingChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void localeChanged()
func (this *QInputMethod) LocaleChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod13localeChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void inputDirectionChanged(Qt::LayoutDirection)
func (this *QInputMethod) InputDirectionChanged(newDirection int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod21inputDirectionChangedEN2Qt15LayoutDirectionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), newDirection)
	gopp.ErrPrint(err, rv)
}

type QInputMethod__Action = int

const QInputMethod__Click QInputMethod__Action = 0
const QInputMethod__ContextMenu QInputMethod__Action = 1

//  body block end

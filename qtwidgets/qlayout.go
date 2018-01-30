package qtwidgets

// /usr/include/qt/QtWidgets/qlayout.h
// #include <qlayout.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
type QLayout struct {
	*qtcore.QObject
	*QLayoutItem
}

func (this *QLayout) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QLayout) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
	this.QLayoutItem = NewQLayoutItemFromPointer(cthis)
}
func NewQLayoutFromPointer(cthis unsafe.Pointer) *QLayout {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	bcthis1 := NewQLayoutItemFromPointer(cthis)
	return &QLayout{bcthis0, bcthis1}
}
func (*QLayout) NewFromPointer(cthis unsafe.Pointer) *QLayout {
	return NewQLayoutFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qlayout.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QLayout) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qlayout.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QLayout(QWidget *)
func NewQLayout(parent *QWidget /*777 QWidget **/) *QLayout {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayoutC1EP7QWidget", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qlayout.h:81
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QLayout()
func NewQLayout_1() *QLayout {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayoutC1Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qlayout.h:82
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QLayout()
func DeleteQLayout(*QLayout) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayoutD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:84
// index:0
// Public Visibility=Default Availability=Available
// [4] int margin()
func (this *QLayout) Margin() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout6marginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlayout.h:85
// index:0
// Public Visibility=Default Availability=Available
// [4] int spacing()
func (this *QLayout) Spacing() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout7spacingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlayout.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMargin(int)
func (this *QLayout) SetMargin(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout9setMarginEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSpacing(int)
func (this *QLayout) SetSpacing(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout10setSpacingEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContentsMargins(int, int, int, int)
func (this *QLayout) SetContentsMargins(left int, top int, right int, bottom int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout18setContentsMarginsEiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), left, top, right, bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:91
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setContentsMargins(const QMargins &)
func (this *QLayout) SetContentsMargins_1(margins *qtcore.QMargins) {
	var convArg0 = margins.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout18setContentsMarginsERK8QMargins", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getContentsMargins(int *, int *, int *, int *)
func (this *QLayout) GetContentsMargins(left unsafe.Pointer /*666*/, top unsafe.Pointer /*666*/, right unsafe.Pointer /*666*/, bottom unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout18getContentsMarginsEPiS0_S0_S0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &left, &top, &right, &bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:93
// index:0
// Public Visibility=Default Availability=Available
// [16] QMargins contentsMargins()
func (this *QLayout) ContentsMargins() *qtcore.QMargins /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout15contentsMarginsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMarginsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qlayout.h:94
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect contentsRect()
func (this *QLayout) ContentsRect() *qtcore.QRect /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout12contentsRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qlayout.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSizeConstraint(enum QLayout::SizeConstraint)
func (this *QLayout) SetSizeConstraint(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout17setSizeConstraintENS_14SizeConstraintE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:101
// index:0
// Public Visibility=Default Availability=Available
// [4] QLayout::SizeConstraint sizeConstraint()
func (this *QLayout) SizeConstraint() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout14sizeConstraintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMenuBar(QWidget *)
func (this *QLayout) SetMenuBar(w *QWidget /*777 QWidget **/) {
	var convArg0 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout10setMenuBarEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * menuBar()
func (this *QLayout) MenuBar() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout7menuBarEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qlayout.h:105
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * parentWidget()
func (this *QLayout) ParentWidget() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout12parentWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qlayout.h:107
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void invalidate()
func (this *QLayout) Invalidate() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout10invalidateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:108
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QRect geometry()
func (this *QLayout) Geometry() *qtcore.QRect /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout8geometryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qlayout.h:109
// index:0
// Public Visibility=Default Availability=Available
// [1] bool activate()
func (this *QLayout) Activate() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout8activateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlayout.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void update()
func (this *QLayout) Update() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout6updateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addWidget(QWidget *)
func (this *QLayout) AddWidget(w *QWidget /*777 QWidget **/) {
	var convArg0 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout9addWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:113
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void addItem(QLayoutItem *)
func (this *QLayout) AddItem(arg0 *QLayoutItem /*777 QLayoutItem **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout7addItemEP11QLayoutItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeWidget(QWidget *)
func (this *QLayout) RemoveWidget(w *QWidget /*777 QWidget **/) {
	var convArg0 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout12removeWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeItem(QLayoutItem *)
func (this *QLayout) RemoveItem(arg0 *QLayoutItem /*777 QLayoutItem **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout10removeItemEP11QLayoutItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:118
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::Orientations expandingDirections()
func (this *QLayout) ExpandingDirections() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout19expandingDirectionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:119
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSize()
func (this *QLayout) MinimumSize() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout11minimumSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qlayout.h:120
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize maximumSize()
func (this *QLayout) MaximumSize() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout11maximumSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qlayout.h:121
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setGeometry(const QRect &)
func (this *QLayout) SetGeometry(arg0 *qtcore.QRect) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout11setGeometryERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:122
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QLayoutItem * itemAt(int)
func (this *QLayout) ItemAt(index int) *QLayoutItem /*777 QLayoutItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout6itemAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qlayout.h:123
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QLayoutItem * takeAt(int)
func (this *QLayout) TakeAt(index int) *QLayoutItem /*777 QLayoutItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout6takeAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qlayout.h:124
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int indexOf(QWidget *)
func (this *QLayout) IndexOf(arg0 *QWidget /*777 QWidget **/) int {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout7indexOfEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlayout.h:125
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int count()
func (this *QLayout) Count() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlayout.h:126
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isEmpty()
func (this *QLayout) IsEmpty() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlayout.h:127
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QSizePolicy::ControlTypes controlTypes()
func (this *QLayout) ControlTypes() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout12controlTypesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:132
// index:0
// Public Visibility=Default Availability=Available
// [4] int totalHeightForWidth(int)
func (this *QLayout) TotalHeightForWidth(w int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout19totalHeightForWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), w)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlayout.h:133
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize totalMinimumSize()
func (this *QLayout) TotalMinimumSize() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout16totalMinimumSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qlayout.h:134
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize totalMaximumSize()
func (this *QLayout) TotalMaximumSize() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout16totalMaximumSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qlayout.h:135
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize totalSizeHint()
func (this *QLayout) TotalSizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout13totalSizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qlayout.h:136
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QLayout * layout()
func (this *QLayout) Layout() *QLayout /*777 QLayout **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout6layoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQLayoutFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qlayout.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEnabled(_Bool)
func (this *QLayout) SetEnabled(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout10setEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:139
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEnabled()
func (this *QLayout) IsEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout9isEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlayout.h:142
// index:0
// Public static Visibility=Default Availability=Available
// [8] QSize closestAcceptableSize(const QWidget *, const QSize &)
func (this *QLayout) ClosestAcceptableSize(w *QWidget /*777 const QWidget **/, s *qtcore.QSize) *qtcore.QSize /*123*/ {
	var convArg0 = w.GetCthis()
	var convArg1 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout21closestAcceptableSizeEPK7QWidgetRK5QSize", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QLayout_ClosestAcceptableSize(w *QWidget /*777 const QWidget **/, s *qtcore.QSize) *qtcore.QSize /*123*/ {
	var nilthis *QLayout
	rv := nilthis.ClosestAcceptableSize(w, s)
	return rv
}

// /usr/include/qt/QtWidgets/qlayout.h:145
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void widgetEvent(QEvent *)
func (this *QLayout) WidgetEvent(arg0 *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout11widgetEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:146
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void childEvent(QChildEvent *)
func (this *QLayout) ChildEvent(e *qtcore.QChildEvent /*777 QChildEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout10childEventEP11QChildEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:147
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void addChildLayout(QLayout *)
func (this *QLayout) AddChildLayout(l *QLayout /*777 QLayout **/) {
	var convArg0 = l.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout14addChildLayoutEPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:148
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void addChildWidget(QWidget *)
func (this *QLayout) AddChildWidget(w *QWidget /*777 QWidget **/) {
	var convArg0 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout14addChildWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:149
// index:0
// Protected Visibility=Default Availability=Available
// [1] bool adoptLayout(QLayout *)
func (this *QLayout) AdoptLayout(layout *QLayout /*777 QLayout **/) bool {
	var convArg0 = layout.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout11adoptLayoutEPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlayout.h:151
// index:0
// Protected Visibility=Default Availability=Available
// [16] QRect alignmentRect(const QRect &)
func (this *QLayout) AlignmentRect(arg0 *qtcore.QRect) *qtcore.QRect /*123*/ {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout13alignmentRectERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

type QLayout__SizeConstraint = int

const QLayout__SetDefaultConstraint QLayout__SizeConstraint = 0
const QLayout__SetNoConstraint QLayout__SizeConstraint = 1
const QLayout__SetMinimumSize QLayout__SizeConstraint = 2
const QLayout__SetFixedSize QLayout__SizeConstraint = 3
const QLayout__SetMaximumSize QLayout__SizeConstraint = 4
const QLayout__SetMinAndMaxSize QLayout__SizeConstraint = 5

//  body block end

package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h
// #include <qgraphicslayoutitem.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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
type QGraphicsLayoutItem struct {
	*qtrt.CObject
}

func (this *QGraphicsLayoutItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QGraphicsLayoutItem) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQGraphicsLayoutItemFromPointer(cthis unsafe.Pointer) *QGraphicsLayoutItem {
	return &QGraphicsLayoutItem{&qtrt.CObject{cthis}}
}
func (*QGraphicsLayoutItem) NewFromPointer(cthis unsafe.Pointer) *QGraphicsLayoutItem {
	return NewQGraphicsLayoutItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsLayoutItem(QGraphicsLayoutItem *, _Bool)
func NewQGraphicsLayoutItem(parent *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/, isLayout bool) *QGraphicsLayoutItem {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItemC1EPS_b", ffiqt.FFI_TYPE_POINTER, convArg0, isLayout)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsLayoutItemFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsLayoutItem()
func DeleteQGraphicsLayoutItem(*QGraphicsLayoutItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSizePolicy(const QSizePolicy &)
func (this *QGraphicsLayoutItem) SetSizePolicy(policy *QSizePolicy) {
	var convArg0 = policy.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem13setSizePolicyERK11QSizePolicy", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:61
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setSizePolicy(QSizePolicy::Policy, QSizePolicy::Policy, QSizePolicy::ControlType)
func (this *QGraphicsLayoutItem) SetSizePolicy_1(hPolicy int, vPolicy int, controlType int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem13setSizePolicyEN11QSizePolicy6PolicyES1_NS0_11ControlTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), hPolicy, vPolicy, controlType)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:62
// index:0
// Public Visibility=Default Availability=Available
// [4] QSizePolicy sizePolicy()
func (this *QGraphicsLayoutItem) SizePolicy() *QSizePolicy /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem10sizePolicyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSizePolicyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumSize(const QSizeF &)
func (this *QGraphicsLayoutItem) SetMinimumSize(size *qtcore.QSizeF) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem14setMinimumSizeERK6QSizeF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:65
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setMinimumSize(qreal, qreal)
func (this *QGraphicsLayoutItem) SetMinimumSize_1(w float64, h float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem14setMinimumSizeEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:66
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF minimumSize()
func (this *QGraphicsLayoutItem) MinimumSize() *qtcore.QSizeF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem11minimumSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumWidth(qreal)
func (this *QGraphicsLayoutItem) SetMinimumWidth(width float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem15setMinimumWidthEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:68
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal minimumWidth()
func (this *QGraphicsLayoutItem) MinimumWidth() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem12minimumWidthEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumHeight(qreal)
func (this *QGraphicsLayoutItem) SetMinimumHeight(height float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem16setMinimumHeightEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), height)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:70
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal minimumHeight()
func (this *QGraphicsLayoutItem) MinimumHeight() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem13minimumHeightEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPreferredSize(const QSizeF &)
func (this *QGraphicsLayoutItem) SetPreferredSize(size *qtcore.QSizeF) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem16setPreferredSizeERK6QSizeF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:73
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setPreferredSize(qreal, qreal)
func (this *QGraphicsLayoutItem) SetPreferredSize_1(w float64, h float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem16setPreferredSizeEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:74
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF preferredSize()
func (this *QGraphicsLayoutItem) PreferredSize() *qtcore.QSizeF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem13preferredSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPreferredWidth(qreal)
func (this *QGraphicsLayoutItem) SetPreferredWidth(width float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem17setPreferredWidthEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:76
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal preferredWidth()
func (this *QGraphicsLayoutItem) PreferredWidth() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem14preferredWidthEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPreferredHeight(qreal)
func (this *QGraphicsLayoutItem) SetPreferredHeight(height float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem18setPreferredHeightEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), height)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:78
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal preferredHeight()
func (this *QGraphicsLayoutItem) PreferredHeight() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem15preferredHeightEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumSize(const QSizeF &)
func (this *QGraphicsLayoutItem) SetMaximumSize(size *qtcore.QSizeF) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem14setMaximumSizeERK6QSizeF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:81
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setMaximumSize(qreal, qreal)
func (this *QGraphicsLayoutItem) SetMaximumSize_1(w float64, h float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem14setMaximumSizeEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:82
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF maximumSize()
func (this *QGraphicsLayoutItem) MaximumSize() *qtcore.QSizeF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem11maximumSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumWidth(qreal)
func (this *QGraphicsLayoutItem) SetMaximumWidth(width float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem15setMaximumWidthEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:84
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal maximumWidth()
func (this *QGraphicsLayoutItem) MaximumWidth() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem12maximumWidthEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumHeight(qreal)
func (this *QGraphicsLayoutItem) SetMaximumHeight(height float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem16setMaximumHeightEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), height)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:86
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal maximumHeight()
func (this *QGraphicsLayoutItem) MaximumHeight() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem13maximumHeightEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setGeometry(const QRectF &)
func (this *QGraphicsLayoutItem) SetGeometry(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem11setGeometryERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:89
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF geometry()
func (this *QGraphicsLayoutItem) Geometry() *qtcore.QRectF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem8geometryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:90
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void getContentsMargins(qreal *, qreal *, qreal *, qreal *)
func (this *QGraphicsLayoutItem) GetContentsMargins(left unsafe.Pointer /*666*/, top unsafe.Pointer /*666*/, right unsafe.Pointer /*666*/, bottom unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem18getContentsMarginsEPdS0_S0_S0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &left, &top, &right, &bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:91
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF contentsRect()
func (this *QGraphicsLayoutItem) ContentsRect() *qtcore.QRectF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem12contentsRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:93
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF effectiveSizeHint(Qt::SizeHint, const QSizeF &)
func (this *QGraphicsLayoutItem) EffectiveSizeHint(which int, constraint *qtcore.QSizeF) *qtcore.QSizeF /*123*/ {
	var convArg1 = constraint.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem17effectiveSizeHintEN2Qt8SizeHintERK6QSizeF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:95
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void updateGeometry()
func (this *QGraphicsLayoutItem) UpdateGeometry() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem14updateGeometryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsLayoutItem * parentLayoutItem()
func (this *QGraphicsLayoutItem) ParentLayoutItem() *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem16parentLayoutItemEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setParentLayoutItem(QGraphicsLayoutItem *)
func (this *QGraphicsLayoutItem) SetParentLayoutItem(parent *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem19setParentLayoutItemEPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:100
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isLayout()
func (this *QGraphicsLayoutItem) IsLayout() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem8isLayoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:101
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsItem * graphicsItem()
func (this *QGraphicsLayoutItem) GraphicsItem() *QGraphicsItem /*777 QGraphicsItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem12graphicsItemEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:102
// index:0
// Public Visibility=Default Availability=Available
// [1] bool ownedByLayout()
func (this *QGraphicsLayoutItem) OwnedByLayout() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem13ownedByLayoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:105
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setGraphicsItem(QGraphicsItem *)
func (this *QGraphicsLayoutItem) SetGraphicsItem(item *QGraphicsItem /*777 QGraphicsItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem15setGraphicsItemEP13QGraphicsItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:106
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setOwnedByLayout(_Bool)
func (this *QGraphicsLayoutItem) SetOwnedByLayout(ownedByLayout bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem16setOwnedByLayoutEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ownedByLayout)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:109
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [16] QSizeF sizeHint(Qt::SizeHint, const QSizeF &)
func (this *QGraphicsLayoutItem) SizeHint(which int, constraint *qtcore.QSizeF) *qtcore.QSizeF /*123*/ {
	var convArg1 = constraint.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem8sizeHintEN2Qt8SizeHintERK6QSizeF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end

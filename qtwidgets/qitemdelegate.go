package qtwidgets

// /usr/include/qt/QtWidgets/qitemdelegate.h
// #include <qitemdelegate.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 55
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
type QItemDelegate struct {
	*QAbstractItemDelegate
}

func (this *QItemDelegate) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractItemDelegate.GetCthis()
	}
}
func (this *QItemDelegate) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractItemDelegate = NewQAbstractItemDelegateFromPointer(cthis)
}
func NewQItemDelegateFromPointer(cthis unsafe.Pointer) *QItemDelegate {
	bcthis0 := NewQAbstractItemDelegateFromPointer(cthis)
	return &QItemDelegate{bcthis0}
}
func (*QItemDelegate) NewFromPointer(cthis unsafe.Pointer) *QItemDelegate {
	return NewQItemDelegateFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QItemDelegate) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QItemDelegate(QObject *)
func NewQItemDelegate(parent *qtcore.QObject /*777 QObject **/) *QItemDelegate {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QItemDelegateC2EP7QObject", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQItemDelegateFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QItemDelegate()
func DeleteQItemDelegate(*QItemDelegate) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QItemDelegateD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:65
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasClipping()
func (this *QItemDelegate) HasClipping() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate11hasClippingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setClipping(_Bool)
func (this *QItemDelegate) SetClipping(clip bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QItemDelegate11setClippingEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), clip)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QStyleOptionViewItem &, const QModelIndex &)
func (this *QItemDelegate) Paint(painter *qtgui.QPainter /*777 QPainter **/, option *QStyleOptionViewItem, index *qtcore.QModelIndex) {
	var convArg0 = painter.GetCthis()
	var convArg1 = option.GetCthis()
	var convArg2 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate5paintEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint(const QStyleOptionViewItem &, const QModelIndex &)
func (this *QItemDelegate) SizeHint(option *QStyleOptionViewItem, index *qtcore.QModelIndex) *qtcore.QSize /*123*/ {
	var convArg0 = option.GetCthis()
	var convArg1 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate8sizeHintERK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QWidget * createEditor(QWidget *, const QStyleOptionViewItem &, const QModelIndex &)
func (this *QItemDelegate) CreateEditor(parent *QWidget /*777 QWidget **/, option *QStyleOptionViewItem, index *qtcore.QModelIndex) *QWidget /*777 QWidget **/ {
	var convArg0 = parent.GetCthis()
	var convArg1 = option.GetCthis()
	var convArg2 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate12createEditorEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setEditorData(QWidget *, const QModelIndex &)
func (this *QItemDelegate) SetEditorData(editor *QWidget /*777 QWidget **/, index *qtcore.QModelIndex) {
	var convArg0 = editor.GetCthis()
	var convArg1 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate13setEditorDataEP7QWidgetRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:81
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setModelData(QWidget *, QAbstractItemModel *, const QModelIndex &)
func (this *QItemDelegate) SetModelData(editor *QWidget /*777 QWidget **/, model *qtcore.QAbstractItemModel /*777 QAbstractItemModel **/, index *qtcore.QModelIndex) {
	var convArg0 = editor.GetCthis()
	var convArg1 = model.GetCthis()
	var convArg2 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate12setModelDataEP7QWidgetP18QAbstractItemModelRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:83
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void updateEditorGeometry(QWidget *, const QStyleOptionViewItem &, const QModelIndex &)
func (this *QItemDelegate) UpdateEditorGeometry(editor *QWidget /*777 QWidget **/, option *QStyleOptionViewItem, index *qtcore.QModelIndex) {
	var convArg0 = editor.GetCthis()
	var convArg1 = option.GetCthis()
	var convArg2 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate20updateEditorGeometryEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] QItemEditorFactory * itemEditorFactory()
func (this *QItemDelegate) ItemEditorFactory() *QItemEditorFactory /*777 QItemEditorFactory **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate17itemEditorFactoryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQItemEditorFactoryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemEditorFactory(QItemEditorFactory *)
func (this *QItemDelegate) SetItemEditorFactory(factory *QItemEditorFactory /*777 QItemEditorFactory **/) {
	var convArg0 = factory.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QItemDelegate20setItemEditorFactoryEP18QItemEditorFactory", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:92
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void drawDisplay(QPainter *, const QStyleOptionViewItem &, const QRect &, const QString &)
func (this *QItemDelegate) DrawDisplay(painter *qtgui.QPainter /*777 QPainter **/, option *QStyleOptionViewItem, rect *qtcore.QRect, text *qtcore.QString) {
	var convArg0 = painter.GetCthis()
	var convArg1 = option.GetCthis()
	var convArg2 = rect.GetCthis()
	var convArg3 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate11drawDisplayEP8QPainterRK20QStyleOptionViewItemRK5QRectRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:94
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void drawDecoration(QPainter *, const QStyleOptionViewItem &, const QRect &, const QPixmap &)
func (this *QItemDelegate) DrawDecoration(painter *qtgui.QPainter /*777 QPainter **/, option *QStyleOptionViewItem, rect *qtcore.QRect, pixmap *qtgui.QPixmap) {
	var convArg0 = painter.GetCthis()
	var convArg1 = option.GetCthis()
	var convArg2 = rect.GetCthis()
	var convArg3 = pixmap.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate14drawDecorationEP8QPainterRK20QStyleOptionViewItemRK5QRectRK7QPixmap", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:96
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void drawFocus(QPainter *, const QStyleOptionViewItem &, const QRect &)
func (this *QItemDelegate) DrawFocus(painter *qtgui.QPainter /*777 QPainter **/, option *QStyleOptionViewItem, rect *qtcore.QRect) {
	var convArg0 = painter.GetCthis()
	var convArg1 = option.GetCthis()
	var convArg2 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate9drawFocusEP8QPainterRK20QStyleOptionViewItemRK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:98
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void drawCheck(QPainter *, const QStyleOptionViewItem &, const QRect &, Qt::CheckState)
func (this *QItemDelegate) DrawCheck(painter *qtgui.QPainter /*777 QPainter **/, option *QStyleOptionViewItem, rect *qtcore.QRect, state int) {
	var convArg0 = painter.GetCthis()
	var convArg1 = option.GetCthis()
	var convArg2 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate9drawCheckEP8QPainterRK20QStyleOptionViewItemRK5QRectN2Qt10CheckStateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:100
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void drawBackground(QPainter *, const QStyleOptionViewItem &, const QModelIndex &)
func (this *QItemDelegate) DrawBackground(painter *qtgui.QPainter /*777 QPainter **/, option *QStyleOptionViewItem, index *qtcore.QModelIndex) {
	var convArg0 = painter.GetCthis()
	var convArg1 = option.GetCthis()
	var convArg2 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate14drawBackgroundEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:103
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void doLayout(const QStyleOptionViewItem &, QRect *, QRect *, QRect *, _Bool)
func (this *QItemDelegate) DoLayout(option *QStyleOptionViewItem, checkRect *qtcore.QRect /*777 QRect **/, iconRect *qtcore.QRect /*777 QRect **/, textRect *qtcore.QRect /*777 QRect **/, hint bool) {
	var convArg0 = option.GetCthis()
	var convArg1 = checkRect.GetCthis()
	var convArg2 = iconRect.GetCthis()
	var convArg3 = textRect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate8doLayoutERK20QStyleOptionViewItemP5QRectS4_S4_b", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:106
// index:0
// Protected Visibility=Default Availability=Available
// [16] QRect rect(const QStyleOptionViewItem &, const QModelIndex &, int)
func (this *QItemDelegate) Rect(option *QStyleOptionViewItem, index *qtcore.QModelIndex, role int) *qtcore.QRect /*123*/ {
	var convArg0 = option.GetCthis()
	var convArg1 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate4rectERK20QStyleOptionViewItemRK11QModelIndexi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, role)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:108
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)
func (this *QItemDelegate) EventFilter(object *qtcore.QObject /*777 QObject **/, event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = object.GetCthis()
	var convArg1 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QItemDelegate11eventFilterEP7QObjectP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:109
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool editorEvent(QEvent *, QAbstractItemModel *, const QStyleOptionViewItem &, const QModelIndex &)
func (this *QItemDelegate) EditorEvent(event *qtcore.QEvent /*777 QEvent **/, model *qtcore.QAbstractItemModel /*777 QAbstractItemModel **/, option *QStyleOptionViewItem, index *qtcore.QModelIndex) bool {
	var convArg0 = event.GetCthis()
	var convArg1 = model.GetCthis()
	var convArg2 = option.GetCthis()
	var convArg3 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QItemDelegate11editorEventEP6QEventP18QAbstractItemModelRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:112
// index:0
// Protected Visibility=Default Availability=Available
// [192] QStyleOptionViewItem setOptions(const QModelIndex &, const QStyleOptionViewItem &)
func (this *QItemDelegate) SetOptions(index *qtcore.QModelIndex, option *QStyleOptionViewItem) *QStyleOptionViewItem /*123*/ {
	var convArg0 = index.GetCthis()
	var convArg1 = option.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate10setOptionsERK11QModelIndexRK20QStyleOptionViewItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStyleOptionViewItemFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:115
// index:0
// Protected Visibility=Default Availability=Available
// [32] QPixmap decoration(const QStyleOptionViewItem &, const QVariant &)
func (this *QItemDelegate) Decoration(option *QStyleOptionViewItem, variant *qtcore.QVariant) *qtgui.QPixmap /*123*/ {
	var convArg0 = option.GetCthis()
	var convArg1 = variant.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate10decorationERK20QStyleOptionViewItemRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:116
// index:0
// Protected Visibility=Default Availability=Available
// [8] QPixmap * selected(const QPixmap &, const QPalette &, _Bool)
func (this *QItemDelegate) Selected(pixmap *qtgui.QPixmap, palette *qtgui.QPalette, enabled bool) *qtgui.QPixmap /*777 QPixmap **/ {
	var convArg0 = pixmap.GetCthis()
	var convArg1 = palette.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate8selectedERK7QPixmapRK8QPaletteb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, enabled)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:118
// index:0
// Protected Visibility=Default Availability=Available
// [16] QRect doCheck(const QStyleOptionViewItem &, const QRect &, const QVariant &)
func (this *QItemDelegate) DoCheck(option *QStyleOptionViewItem, bounding *qtcore.QRect, variant *qtcore.QVariant) *qtcore.QRect /*123*/ {
	var convArg0 = option.GetCthis()
	var convArg1 = bounding.GetCthis()
	var convArg2 = variant.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate7doCheckERK20QStyleOptionViewItemRK5QRectRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:120
// index:0
// Protected Visibility=Default Availability=Available
// [16] QRect textRectangle(QPainter *, const QRect &, const QFont &, const QString &)
func (this *QItemDelegate) TextRectangle(painter *qtgui.QPainter /*777 QPainter **/, rect *qtcore.QRect, font *qtgui.QFont, text *qtcore.QString) *qtcore.QRect /*123*/ {
	var convArg0 = painter.GetCthis()
	var convArg1 = rect.GetCthis()
	var convArg2 = font.GetCthis()
	var convArg3 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate13textRectangleEP8QPainterRK5QRectRK5QFontRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end

package qtgui

// /usr/include/qt/QtGui/qpaintdevice.h
// #include <qpaintdevice.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 54
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
type QPaintDevice struct {
	*qtrt.CObject
}

func (this *QPaintDevice) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func NewQPaintDeviceFromPointer(cthis unsafe.Pointer) *QPaintDevice {
	return &QPaintDevice{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qpaintdevice.h:72
// index:0
// Public virtual
// void ~QPaintDevice()
func DeleteQPaintDevice(*QPaintDevice) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintDeviceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevice.h:74
// index:0
// Public virtual
// int devType()
func (this *QPaintDevice) DevType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice7devTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qpaintdevice.h:75
// index:0
// Public
// bool paintingActive()
func (this *QPaintDevice) PaintingActive() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice14paintingActiveEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpaintdevice.h:76
// index:0
// Public pure virtual
// QPaintEngine * paintEngine()
func (this *QPaintDevice) PaintEngine() *QPaintEngine /*444 QPaintEngine **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice11paintEngineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPaintEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qpaintdevice.h:78
// index:0
// Public inline
// int width()
func (this *QPaintDevice) Width() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice5widthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qpaintdevice.h:79
// index:0
// Public inline
// int height()
func (this *QPaintDevice) Height() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice6heightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qpaintdevice.h:80
// index:0
// Public inline
// int widthMM()
func (this *QPaintDevice) WidthMM() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice7widthMMEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qpaintdevice.h:81
// index:0
// Public inline
// int heightMM()
func (this *QPaintDevice) HeightMM() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice8heightMMEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qpaintdevice.h:82
// index:0
// Public inline
// int logicalDpiX()
func (this *QPaintDevice) LogicalDpiX() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice11logicalDpiXEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qpaintdevice.h:83
// index:0
// Public inline
// int logicalDpiY()
func (this *QPaintDevice) LogicalDpiY() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice11logicalDpiYEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qpaintdevice.h:84
// index:0
// Public inline
// int physicalDpiX()
func (this *QPaintDevice) PhysicalDpiX() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice12physicalDpiXEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qpaintdevice.h:85
// index:0
// Public inline
// int physicalDpiY()
func (this *QPaintDevice) PhysicalDpiY() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice12physicalDpiYEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qpaintdevice.h:86
// index:0
// Public inline
// int devicePixelRatio()
func (this *QPaintDevice) DevicePixelRatio() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice16devicePixelRatioEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qpaintdevice.h:87
// index:0
// Public inline
// qreal devicePixelRatioF()
func (this *QPaintDevice) DevicePixelRatioF() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice17devicePixelRatioFEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qpaintdevice.h:88
// index:0
// Public inline
// int colorCount()
func (this *QPaintDevice) ColorCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice10colorCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qpaintdevice.h:89
// index:0
// Public inline
// int depth()
func (this *QPaintDevice) Depth() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice5depthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qpaintdevice.h:91
// index:0
// Public static inline
// qreal devicePixelRatioFScale()
func (this *QPaintDevice) DevicePixelRatioFScale() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintDevice22devicePixelRatioFScaleEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return float64(rv) // 222
}
func QPaintDevice_DevicePixelRatioFScale() float64 {
	var nilthis *QPaintDevice
	rv := nilthis.DevicePixelRatioFScale()
	return rv
}

// /usr/include/qt/QtGui/qpaintdevice.h:93
// index:0
// Protected
// void QPaintDevice()
func NewQPaintDevice() *QPaintDevice {
	cthis := qtrt.Calloc(1, 256) // 24
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintDeviceC1Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQPaintDeviceFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpaintdevice.h:101
// index:1
// Private inline
// void QPaintDevice(const class QPaintDevice &)
func NewQPaintDevice_1(arg0 *QPaintDevice) *QPaintDevice {
	cthis := qtrt.Calloc(1, 256) // 24
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintDeviceC1ERKS_", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPaintDeviceFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpaintdevice.h:94
// index:0
// Protected virtual
// int metric(enum QPaintDevice::PaintDeviceMetric)
func (this *QPaintDevice) Metric(metric int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice6metricENS_17PaintDeviceMetricE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &metric)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qpaintdevice.h:95
// index:0
// Protected virtual
// void initPainter(class QPainter *)
func (this *QPaintDevice) InitPainter(painter *QPainter /*444 QPainter **/) {
	var convArg0 = painter.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice11initPainterEP8QPainter", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevice.h:96
// index:0
// Protected virtual
// QPaintDevice * redirected(class QPoint *)
func (this *QPaintDevice) Redirected(offset *qtcore.QPoint /*444 QPoint **/) *QPaintDevice /*444 QPaintDevice **/ {
	var convArg0 = offset.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice10redirectedEP6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPaintDeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qpaintdevice.h:97
// index:0
// Protected virtual
// QPainter * sharedPainter()
func (this *QPaintDevice) SharedPainter() *QPainter /*444 QPainter **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice13sharedPainterEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPainterFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

//  body block end
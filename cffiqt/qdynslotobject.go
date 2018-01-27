package ffiqt

/*
extern void callbackAllQDynSlotObject(void*, int, int, void**, char*, int, int*, void*);
*/
import "C"
import (
	"gopp"
	"log"
	"unsafe"

	"github.com/therecipe/qt"
)

// 也许放在qtrt包中更好，但在测试时放在了ffiqt包中，暂时放在这
type QDynSlotObject struct {
	cthis unsafe.Pointer

	//state
	sigsrc unsafe.Pointer
	sigobj interface{} // 保存类型信息
}

func (this *QDynSlotObject) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return this.cthis
}

var _QDynSlotObj *QDynSlotObject

func NewQDynSlotObject(signalName string, argc int) *QDynSlotObject {
	this := &QDynSlotObject{}
	// void *fnptr, char* name, int argc, int*argtys, void* cbptr
	name_ := C.CString(signalName)
	argc_ := int(argc)
	argtys_ := (*int)(nil)
	var cbptr_ unsafe.Pointer
	fnptr_ := unsafe.Pointer(C.callbackAllQDynSlotObject)

	rv, err := InvokeQtFunc6("QDynSlotObject_new", FFI_TYPE_POINTER,
		fnptr_, name_, argc_, argtys_, cbptr_)
	gopp.ErrPrint(err, rv)

	this.cthis = unsafe.Pointer(uintptr(rv))
	_QDynSlotObj = this
	return this
}

func DeleteQDynSlotObject(o *QDynSlotObject) {
	rv, err := InvokeQtFunc6("QDynSlotObject_delete", FFI_TYPE_VOID, o.cthis)
	gopp.ErrPrint(err, rv)
}

type CObjectIF interface {
	GetCthis() unsafe.Pointer
	SetCthis(unsafe.Pointer)
}

func (this *QDynSlotObject) Connect(cobj CObjectIF, signame string, f interface{}) {
	cptr := cobj.GetCthis()
	gopp.NilPrint(cptr, "cptr nil")

	if cptr != nil {
		// 相当于初始化信号开关
		if !qt.ExistsSignal(cptr, signame) {
			// C.QMessageBox_ConnectButtonClicked(ptr.Pointer())
			this.ConnectSwitch(cptr, signame, true, cobj)
		}

		if signal := qt.LendSignal(cptr, signame); signal != nil {
			qt.ConnectSignal(cptr, signame, func(argvals /* **C.uchar*/ unsafe.Pointer, sigobj interface{}) {
				signal.(func(unsafe.Pointer, interface{}))(argvals, sigobj)
				callbackInvoke(f, argvals, sigobj)
				if false {
					signal.(func())()
					f.(func())()
				}
			})
		} else {
			qt.ConnectSignal(cptr, signame, func(argvals /* **C.uchar*/ unsafe.Pointer, sigobj interface{}) {
				callbackInvoke(f, argvals, sigobj)
			})
		}
	}
}

func (this *QDynSlotObject) ConnectSwitch(src unsafe.Pointer, signame string, on bool, cobj CObjectIF) {
	signamep := QSIGNAL(signame)
	signame_ := C.CString(signamep)
	log.Println(signame_, signamep, on)

	// 相当于初始化信号开关
	var subDynSlot *QDynSlotObject
	if !qt.ExistsSignal(src, signamep) {
		subDynSlot = NewQDynSlotObject(signamep, int(456))
		subDynSlot.sigsrc = src
		subDynSlot.sigobj = cobj
	}
	if signal := qt.LendSignal(src, signamep); signal != nil {
		subDynSlot = signal.(*QDynSlotObject)
	} else {
		qt.ConnectSignal(subDynSlot.cthis, signamep, subDynSlot)
	}

	if on {
		rv, err := InvokeQtFunc6("QDynSlotObject_connect_switch", FFI_TYPE_VOID,
			src, signame_, subDynSlot.cthis, int(1))
		gopp.ErrPrint(err, rv)
	} else {
		rv, err := InvokeQtFunc6("QDynSlotObject_connect_switch", FFI_TYPE_VOID,
			src, signame_, subDynSlot.cthis, int(0))
		gopp.ErrPrint(err, rv)
	}
}

func (this *QDynSlotObject) Connect_test(to *QDynSlotObject) {
	rv, err := InvokeQtFunc6("QDynSlotObject_connect_test", FFI_TYPE_VOID, this.cthis, to.cthis)
	gopp.ErrPrint(err, rv)
}

func (this *QDynSlotObject) Trigger_test() {
	rv, err := InvokeQtFunc6("QDynSlotObject_trigger_test", FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}
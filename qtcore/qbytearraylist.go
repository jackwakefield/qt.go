package qtcore

// /usr/include/qt/QtCore/qbytearray.h
// #include <qbytearray.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end

//  body block begin
type QByteArrayList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QByteArrayList) Operator_equal_0() *QByteArrayList {
	// QByteArrayList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QByteArrayList) Operator_equal_1() *QByteArrayList {
	// QByteArrayList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QByteArrayList) Swap_0() {
	// QByteArrayList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QByteArrayList) Operator_equal_equal_0() bool {
	// QByteArrayList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QByteArrayList) Operator_not_equal_0() bool {
	// QByteArrayList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QByteArrayList) Size_0() int {
	// QByteArrayList_size_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QByteArrayList) Detach_0() {
	// QByteArrayList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QByteArrayList) DetachShared_0() {
	// QByteArrayList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QByteArrayList) IsDetached_0() bool {
	// QByteArrayList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QByteArrayList) SetSharable_0() {
	// QByteArrayList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QByteArrayList) IsSharedWith_0() bool {
	// QByteArrayList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QByteArrayList) IsEmpty_0() bool {
	// QByteArrayList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QByteArrayList) Clear_0() {
	// QByteArrayList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QByteArrayList) At_0() *QByteArray {
	// QByteArrayList_at_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QByteArray{}
}

// const T & operator[](int)
func (this *QByteArrayList) Operator_get_index_0() *QByteArray {
	// QByteArrayList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QByteArray{}
}

// T & operator[](int)
func (this *QByteArrayList) Operator_get_index_1() *QByteArray {
	// QByteArrayList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QByteArray{}
}

// void reserve(int)
func (this *QByteArrayList) Reserve_0() {
	// QByteArrayList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QByteArrayList) Append_0() {
	// QByteArrayList_append_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QByteArrayList) Append_1() {
	// QByteArrayList_append_1()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QByteArrayList) Prepend_0() {
	// QByteArrayList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QByteArrayList) Insert_0() {
	// QByteArrayList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QByteArrayList) Replace_0() {
	// QByteArrayList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QByteArrayList) RemoveAt_0() {
	// QByteArrayList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QByteArrayList) RemoveAll_0() int {
	// QByteArrayList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QByteArrayList) RemoveOne_0() bool {
	// QByteArrayList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QByteArrayList) TakeAt_0() *QByteArray {
	// QByteArrayList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QByteArray{}
}

// T takeFirst()
func (this *QByteArrayList) TakeFirst_0() *QByteArray {
	// QByteArrayList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QByteArray{}
}

// T takeLast()
func (this *QByteArrayList) TakeLast_0() *QByteArray {
	// QByteArrayList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QByteArray{}
}

// void move(int, int)
func (this *QByteArrayList) Move_0() {
	// QByteArrayList_move_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QByteArrayList) Swap_1() {
	// QByteArrayList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QByteArrayList) IndexOf_0() int {
	// QByteArrayList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QByteArrayList) LastIndexOf_0() int {
	// QByteArrayList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QByteArrayList) Contains_0() bool {
	// QByteArrayList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QByteArrayList) Count_0() int {
	// QByteArrayList_count_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QByteArrayList) Begin_0() {
	// QByteArrayList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QByteArrayList) Begin_1() {
	// QByteArrayList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QByteArrayList) Cbegin_0() {
	// QByteArrayList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QByteArrayList) ConstBegin_0() {
	// QByteArrayList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QByteArrayList) End_0() {
	// QByteArrayList_end_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QByteArrayList) End_1() {
	// QByteArrayList_end_1()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QByteArrayList) Cend_0() {
	// QByteArrayList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QByteArrayList) ConstEnd_0() {
	// QByteArrayList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QByteArrayList) Rbegin_0() {
	// QByteArrayList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QByteArrayList) Rend_0() {
	// QByteArrayList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QByteArrayList) Rbegin_1() {
	// QByteArrayList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QByteArrayList) Rend_1() {
	// QByteArrayList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QByteArrayList) Crbegin_0() {
	// QByteArrayList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QByteArrayList) Crend_0() {
	// QByteArrayList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QByteArrayList) Insert_1() {
	// QByteArrayList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QByteArrayList) Erase_0() {
	// QByteArrayList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QByteArrayList) Erase_1() {
	// QByteArrayList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QByteArrayList) Count_1() int {
	// QByteArrayList_count_1()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QByteArrayList) Length_0() int {
	// QByteArrayList_length_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QByteArrayList) First_0() *QByteArray {
	// QByteArrayList_first_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QByteArray{}
}

// const T & constFirst()
func (this *QByteArrayList) ConstFirst_0() *QByteArray {
	// QByteArrayList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QByteArray{}
}

// const T & first()
func (this *QByteArrayList) First_1() *QByteArray {
	// QByteArrayList_first_1()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QByteArray{}
}

// T & last()
func (this *QByteArrayList) Last_0() *QByteArray {
	// QByteArrayList_last_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QByteArray{}
}

// const T & last()
func (this *QByteArrayList) Last_1() *QByteArray {
	// QByteArrayList_last_1()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QByteArray{}
}

// const T & constLast()
func (this *QByteArrayList) ConstLast_0() *QByteArray {
	// QByteArrayList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QByteArray{}
}

// void removeFirst()
func (this *QByteArrayList) RemoveFirst_0() {
	// QByteArrayList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QByteArrayList) RemoveLast_0() {
	// QByteArrayList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QByteArrayList) StartsWith_0() bool {
	// QByteArrayList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QByteArrayList) EndsWith_0() bool {
	// QByteArrayList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QByteArrayList) Mid_0() *QByteArrayList {
	// QByteArrayList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QByteArrayList) Value_0() *QByteArray {
	// QByteArrayList_value_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QByteArray{}
}

// T value(int, const T &)
func (this *QByteArrayList) Value_1() *QByteArray {
	// QByteArrayList_value_1()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QByteArray{}
}

// void push_back(const T &)
func (this *QByteArrayList) Push_back_0() {
	// QByteArrayList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QByteArrayList) Push_front_0() {
	// QByteArrayList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QByteArrayList) Front_0() *QByteArray {
	// QByteArrayList_front_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QByteArray{}
}

// const T & front()
func (this *QByteArrayList) Front_1() *QByteArray {
	// QByteArrayList_front_1()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QByteArray{}
}

// T & back()
func (this *QByteArrayList) Back_0() *QByteArray {
	// QByteArrayList_back_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QByteArray{}
}

// const T & back()
func (this *QByteArrayList) Back_1() *QByteArray {
	// QByteArrayList_back_1()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QByteArray{}
}

// void pop_front()
func (this *QByteArrayList) Pop_front_0() {
	// QByteArrayList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QByteArrayList) Pop_back_0() {
	// QByteArrayList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QByteArrayList) Empty_0() bool {
	// QByteArrayList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QByteArrayList) Operator_add_equal_0() *QByteArrayList {
	// QByteArrayList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QByteArrayList) Operator_add_0() *QByteArrayList {
	// QByteArrayList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QByteArrayList) Operator_add_equal_1() *QByteArrayList {
	// QByteArrayList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QByteArrayList) Operator_left_shift_0() *QByteArrayList {
	// QByteArrayList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QByteArrayList) Operator_left_shift_1() *QByteArrayList {
	// QByteArrayList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QByteArrayList) ToVector_0() {
	// QByteArrayList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QByteArrayList) ToSet_0() {
	// QByteArrayList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QByteArrayList) FromVector_0() *QByteArrayList {
	// QByteArrayList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QByteArrayList) FromSet_0() *QByteArrayList {
	// QByteArrayList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QByteArrayList) FromStdList_0() *QByteArrayList {
	// QByteArrayList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QByteArrayList) ToStdList_0() {
	// QByteArrayList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QByteArrayList) Detach_helper_grow_0() {
	// QByteArrayList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QByteArrayList) Detach_helper_0() {
	// QByteArrayList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QByteArrayList) Detach_helper_1() {
	// QByteArrayList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QByteArrayList) Dealloc_0() {
	// QByteArrayList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QByteArrayList) Node_construct_0() {
	// QByteArrayList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QByteArrayList) Node_destruct_0() {
	// QByteArrayList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QByteArrayList) Node_copy_0() {
	// QByteArrayList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QByteArrayList) Node_destruct_1() {
	// QByteArrayList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QByteArrayList) IsValidIterator_0() bool {
	// QByteArrayList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QByteArrayList) Op_eq_impl_0() bool {
	// QByteArrayList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QByteArrayList) Op_eq_impl_1() bool {
	// QByteArrayList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QByteArrayList) Contains_impl_0() bool {
	// QByteArrayList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QByteArrayList) Contains_impl_1() bool {
	// QByteArrayList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QByteArrayList) Count_impl_0() int {
	// QByteArrayList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QByteArrayList) Count_impl_1() int {
	// QByteArrayList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QByteArrayList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end

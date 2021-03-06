package qtcore

// /usr/include/qt/QtCore/qstring.h
// #include <qstring.h>
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
type QVariantHash struct {
	*qtrt.CObject
}

// QHash::Node * concrete(QHashData::Node *)
func (this *QVariantHash) Concrete_0() {
	// QVariantHash_concrete_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_concrete_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int alignOfNode()
func (this *QVariantHash) AlignOfNode_0() int {
	// QVariantHash_alignOfNode_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_alignOfNode_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QHash<K, V> & operator=(const QHash<K, V> &)
func (this *QVariantHash) Operator_equal_0() {
	// QVariantHash_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QHash<K, V> & operator=(QHash<K, V> &&)
func (this *QVariantHash) Operator_equal_1() {
	// QVariantHash_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(QHash<K, V> &)
func (this *QVariantHash) Swap_0() {
	// QVariantHash_swap_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QHash<K, V> &)
func (this *QVariantHash) Operator_equal_equal_0() bool {
	// QVariantHash_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QHash<K, V> &)
func (this *QVariantHash) Operator_not_equal_0() bool {
	// QVariantHash_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QVariantHash) Size_0() int {
	// QVariantHash_size_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool isEmpty()
func (this *QVariantHash) IsEmpty_0() bool {
	// QVariantHash_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int capacity()
func (this *QVariantHash) Capacity_0() int {
	// QVariantHash_capacity_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_capacity_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void reserve(int)
func (this *QVariantHash) Reserve_0() {
	// QVariantHash_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void squeeze()
func (this *QVariantHash) Squeeze_0() {
	// QVariantHash_squeeze_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_squeeze_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach()
func (this *QVariantHash) Detach_0() {
	// QVariantHash_detach_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QVariantHash) IsDetached_0() bool {
	// QVariantHash_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QVariantHash) SetSharable_0() {
	// QVariantHash_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QHash<K, V> &)
func (this *QVariantHash) IsSharedWith_0() bool {
	// QVariantHash_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QVariantHash) Clear_0() {
	// QVariantHash_clear_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int remove(const Key &)
func (this *QVariantHash) Remove_0() int {
	// QVariantHash_remove_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_remove_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T take(const Key &)
func (this *QVariantHash) Take_0() *QVariant {
	// QVariantHash_take_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_take_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// bool contains(const Key &)
func (this *QVariantHash) Contains_0() bool {
	// QVariantHash_contains_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// const Key key(const T &)
func (this *QVariantHash) Key_0() {
	// QVariantHash_key_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_key_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const Key key(const T &, const Key &)
func (this *QVariantHash) Key_1() {
	// QVariantHash_key_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_key_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T value(const Key &)
func (this *QVariantHash) Value_0() *QVariant {
	// QVariantHash_value_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// const T value(const Key &, const T &)
func (this *QVariantHash) Value_1() *QVariant {
	// QVariantHash_value_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// T & operator[](const Key &)
func (this *QVariantHash) Operator_get_index_0() *QVariant {
	// QVariantHash_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// const T operator[](const Key &)
func (this *QVariantHash) Operator_get_index_1() *QVariant {
	// QVariantHash_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// QList<Key> uniqueKeys()
func (this *QVariantHash) UniqueKeys_0() {
	// QVariantHash_uniqueKeys_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_uniqueKeys_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<Key> keys()
func (this *QVariantHash) Keys_0() {
	// QVariantHash_keys_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_keys_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<Key> keys(const T &)
func (this *QVariantHash) Keys_1() {
	// QVariantHash_keys_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_keys_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> values()
func (this *QVariantHash) Values_0() {
	// QVariantHash_values_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_values_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> values(const Key &)
func (this *QVariantHash) Values_1() {
	// QVariantHash_values_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_values_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count(const Key &)
func (this *QVariantHash) Count_0() int {
	// QVariantHash_count_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QHash::iterator begin()
func (this *QVariantHash) Begin_0() {
	// QVariantHash_begin_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QHash::const_iterator begin()
func (this *QVariantHash) Begin_1() {
	// QVariantHash_begin_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QHash::const_iterator cbegin()
func (this *QVariantHash) Cbegin_0() {
	// QVariantHash_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QHash::const_iterator constBegin()
func (this *QVariantHash) ConstBegin_0() {
	// QVariantHash_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QHash::iterator end()
func (this *QVariantHash) End_0() {
	// QVariantHash_end_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QHash::const_iterator end()
func (this *QVariantHash) End_1() {
	// QVariantHash_end_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QHash::const_iterator cend()
func (this *QVariantHash) Cend_0() {
	// QVariantHash_cend_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QHash::const_iterator constEnd()
func (this *QVariantHash) ConstEnd_0() {
	// QVariantHash_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QHash::key_iterator keyBegin()
func (this *QVariantHash) KeyBegin_0() {
	// QVariantHash_keyBegin_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_keyBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QHash::key_iterator keyEnd()
func (this *QVariantHash) KeyEnd_0() {
	// QVariantHash_keyEnd_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_keyEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QHash::key_value_iterator keyValueBegin()
func (this *QVariantHash) KeyValueBegin_0() {
	// QVariantHash_keyValueBegin_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_keyValueBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QHash::key_value_iterator keyValueEnd()
func (this *QVariantHash) KeyValueEnd_0() {
	// QVariantHash_keyValueEnd_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_keyValueEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QHash::const_key_value_iterator keyValueBegin()
func (this *QVariantHash) KeyValueBegin_1() {
	// QVariantHash_keyValueBegin_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_keyValueBegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QHash::const_key_value_iterator constKeyValueBegin()
func (this *QVariantHash) ConstKeyValueBegin_0() {
	// QVariantHash_constKeyValueBegin_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_constKeyValueBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QHash::const_key_value_iterator keyValueEnd()
func (this *QVariantHash) KeyValueEnd_1() {
	// QVariantHash_keyValueEnd_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_keyValueEnd_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QHash::const_key_value_iterator constKeyValueEnd()
func (this *QVariantHash) ConstKeyValueEnd_0() {
	// QVariantHash_constKeyValueEnd_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_constKeyValueEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QPair<QHash::iterator, QHash::iterator> equal_range(const Key &)
func (this *QVariantHash) Equal_range_0() {
	// QVariantHash_equal_range_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_equal_range_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QPair<QHash::const_iterator, QHash::const_iterator> equal_range(const Key &)
func (this *QVariantHash) Equal_range_1() {
	// QVariantHash_equal_range_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_equal_range_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QHash::iterator erase(QHash::iterator)
func (this *QVariantHash) Erase_0() {
	// QVariantHash_erase_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QHash::iterator erase(QHash::const_iterator)
func (this *QVariantHash) Erase_1() {
	// QVariantHash_erase_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QVariantHash) Count_1() int {
	// QVariantHash_count_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QHash::iterator find(const Key &)
func (this *QVariantHash) Find_0() {
	// QVariantHash_find_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_find_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QHash::const_iterator find(const Key &)
func (this *QVariantHash) Find_1() {
	// QVariantHash_find_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_find_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QHash::const_iterator constFind(const Key &)
func (this *QVariantHash) ConstFind_0() {
	// QVariantHash_constFind_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_constFind_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QHash::iterator insert(const Key &, const T &)
func (this *QVariantHash) Insert_0() {
	// QVariantHash_insert_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QHash::iterator insertMulti(const Key &, const T &)
func (this *QVariantHash) InsertMulti_0() {
	// QVariantHash_insertMulti_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_insertMulti_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QHash<K, V> & unite(const QHash<K, V> &)
func (this *QVariantHash) Unite_0() {
	// QVariantHash_unite_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_unite_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QVariantHash) Empty_0() bool {
	// QVariantHash_empty_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void detach_helper()
func (this *QVariantHash) Detach_helper_0() {
	// QVariantHash_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void freeData(QHashData *)
func (this *QVariantHash) FreeData_0() {
	// QVariantHash_freeData_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_freeData_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QHash::Node ** findNode(const Key &, uint *)
func (this *QVariantHash) FindNode_0() {
	// QVariantHash_findNode_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_findNode_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QHash::Node ** findNode(const Key &, uint)
func (this *QVariantHash) FindNode_1() {
	// QVariantHash_findNode_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_findNode_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QHash::Node * createNode(uint, const Key &, const T &, QHash::Node **)
func (this *QVariantHash) CreateNode_0() {
	// QVariantHash_createNode_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_createNode_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void deleteNode(QHash::Node *)
func (this *QVariantHash) DeleteNode_0() {
	// QVariantHash_deleteNode_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_deleteNode_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void deleteNode2(QHashData::Node *)
func (this *QVariantHash) DeleteNode2_0() {
	// QVariantHash_deleteNode2_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_deleteNode2_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void duplicateNode(QHashData::Node *, void *)
func (this *QVariantHash) DuplicateNode_0() {
	// QVariantHash_duplicateNode_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_duplicateNode_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QHash::iterator &)
func (this *QVariantHash) IsValidIterator_0() bool {
	// QVariantHash_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isValidIterator(const QHash::const_iterator &)
func (this *QVariantHash) IsValidIterator_1() bool {
	// QVariantHash_isValidIterator_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_isValidIterator_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isValidNode(QHashData::Node *)
func (this *QVariantHash) IsValidNode_0() bool {
	// QVariantHash_isValidNode_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantHash_isValidNode_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

//  body block end

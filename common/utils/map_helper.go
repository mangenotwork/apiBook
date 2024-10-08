package utils

import (
	"apiBook/common/log"
	"errors"
	"math/rand"
	"reflect"
	"sort"
	"sync"
	"time"
)

// orderMap 固定顺序map
type orderMap[K, V comparable] struct {
	mux     sync.Mutex // TODO 优化 使用读写锁
	data    map[K]V
	keyList []K // TODO 优化 使用链表
	size    int
}

// OrderMap ues: OrderMap[K, V]()
func OrderMap[K, V comparable]() *orderMap[K, V] {
	obj := &orderMap[K, V]{
		mux:     sync.Mutex{},
		data:    make(map[K]V),
		keyList: make([]K, 0),
		size:    0,
	}
	return obj
}

func (m *orderMap[K, V]) Add(key K, value V) *orderMap[K, V] {
	m.mux.Lock()
	defer m.mux.Unlock()
	if _, ok := m.data[key]; ok {
		m.data[key] = value
		return m
	}
	m.keyList = append(m.keyList, key)
	m.size++
	m.data[key] = value
	return m
}

func (m *orderMap[K, V]) Get(key K) V {
	m.mux.Lock()
	defer m.mux.Unlock()
	return m.data[key]
}

func (m *orderMap[K, V]) Del(key K) *orderMap[K, V] {
	m.mux.Lock()
	defer m.mux.Unlock()
	if _, ok := m.data[key]; ok {
		delete(m.data, key)
		for i := 0; i < m.size; i++ {
			if m.keyList[i] == key {
				m.keyList = append(m.keyList[:i], m.keyList[i+1:]...)
				m.size--
				return m
			}
		}
	}
	return m
}

func (m *orderMap[K, V]) Len() int {
	return m.size
}

func (m *orderMap[K, V]) KeyList() []K {
	return m.keyList
}

func (m *orderMap[K, V]) AddMap(data map[K]V) *orderMap[K, V] {
	for k, v := range data {
		m.Add(k, v)
	}
	return m
}

func (m *orderMap[K, V]) Range(f func(k K, v V)) *orderMap[K, V] {
	for i := 0; i < m.size; i++ {
		f(m.keyList[i], m.data[m.keyList[i]])
	}
	return m
}

// RangeAt Range 遍历map含顺序id
func (m *orderMap[K, V]) RangeAt(f func(id int, k K, v V)) *orderMap[K, V] {
	for i := 0; i < m.size; i++ {
		f(i, m.keyList[i], m.data[m.keyList[i]])
	}
	return m
}

// CheckValue 查看map是否存在指定的值
func (m *orderMap[K, V]) CheckValue(value V) bool {
	m.mux.Lock()
	defer m.mux.Unlock()
	for i := 0; i < m.size; i++ {
		if m.data[m.keyList[i]] == value {
			return true
		}
	}
	return false
}

// Reverse map反序
func (m *orderMap[K, V]) Reverse() *orderMap[K, V] {
	for i, j := 0, len(m.keyList)-1; i < j; i, j = i+1, j-1 {
		m.keyList[i], m.keyList[j] = m.keyList[j], m.keyList[i]
	}
	return m
}

func (m *orderMap[K, V]) Json() (string, error) {
	return AnyToJson(m.data)
}

func (m *orderMap[K, V]) DebugPrint() {
	m.RangeAt(func(id int, k K, v V) {
		log.DebugF("item:%d key:%v value:%v", id, k, v)
	})
}

// Insert 插入值指定位置
func (m *orderMap[K, V]) Insert(k K, v V, position int) error {
	m.Add(k, v)
	return m.Move(k, position)
}

// Move 值移动指定位置操作
func (m *orderMap[K, V]) Move(k K, position int) error {
	if position >= m.size {
		return errors.New("position >= map len")
	}
	has := false
	for i := 0; i < m.size; i++ {
		if m.keyList[i] == k {
			m.keyList = append(m.keyList[0:i], m.keyList[i+1:]...)
			has = true
			break
		}
	}
	if has {
		m.keyList = append(m.keyList[:position+1], m.keyList[position:]...)
		m.keyList[position] = k
	}
	return nil
}

func (m *orderMap[K, V]) GetAtPosition(position int) (K, V, error) {
	if position >= m.size {
		var (
			k K
			v V
		)
		return k, v, errors.New("position >= map len")
	}
	k := m.keyList[position]
	v := m.data[k]
	return k, v, nil
}

// Pop 首位读取并移除
func (m *orderMap[K, V]) Pop() (K, V, error) {
	if m.size < 1 {
		var (
			k K
			v V
		)
		return k, v, errors.New("map size is 0")
	}
	k := m.keyList[0]
	v := m.data[k]
	m.Del(k)
	return k, v, nil
}

// BackPop 末尾读取并移除
func (m *orderMap[K, V]) BackPop() (K, V, error) {
	if m.size < 1 {
		var (
			k K
			v V
		)
		return k, v, errors.New("map size is 0")
	}
	k := m.keyList[m.size-1]
	v := m.data[k]
	m.Del(k)
	return k, v, nil
}

// Shuffle 洗牌
func (m *orderMap[K, V]) Shuffle() {
	if m.size <= 1 {
		return
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(m.size, func(i, j int) {
		m.keyList[i], m.keyList[j] = m.keyList[j], m.keyList[i]
	})
}

// SortDesc 根据k排序 desc
func (m *orderMap[K, V]) SortDesc() {
	sort.Slice(m.keyList, func(i, j int) bool {
		if i > j {
			return true
		}
		return false
	})
}

// SortAsc 根据k排序 asc
func (m *orderMap[K, V]) SortAsc() {
	sort.Slice(m.keyList, func(i, j int) bool {
		if i < j {
			return true
		}
		return false
	})
}

func (m *orderMap[K, V]) CopyMap() map[K]V {
	temp := make(map[K]V, m.size)
	m.Range(func(k K, v V) {
		temp[k] = v
	})
	return temp
}

type Set map[any]struct{}

func NewSet() Set {
	return make(Set)
}

func (s Set) Has(key any) bool {
	_, ok := s[key]
	return ok
}

func (s Set) Add(key any) {
	s[key] = struct{}{}
}

func (s Set) Delete(key any) {
	delete(s, key)
}

func (s Set) DebugPrint() {
	for k := range s {
		log.Debug(k)
	}
}

type Stack[K comparable] struct {
	data map[int]K
}

func NewStack[K comparable]() *Stack[K] {
	return &Stack[K]{
		data: make(map[int]K),
	}
}

func (s *Stack[K]) Push(data K) {
	s.data[len(s.data)] = data
}

func (s *Stack[K]) Pop() K {
	if len(s.data) < 1 {
		var k K
		return k
	}
	l := len(s.data) - 1
	k := s.data[l]
	delete(s.data, l)
	return k
}

func (s *Stack[K]) DebugPrint() {
	for i := 0; i < len(s.data); i++ {
		log.Debug(s.data)
	}
}

func MapCopy[K, V comparable](data map[K]V) (copy map[K]V) {
	copy = make(map[K]V, len(data))
	for k, v := range data {
		copy[k] = v
	}
	return
}

func MapMergeCopy[K, V comparable](src ...map[K]V) (copy map[K]V) {
	copy = make(map[K]V)
	for _, m := range src {
		for k, v := range m {
			copy[k] = v
		}
	}
	return
}

// Map2Slice Eg: {"K1": "v1", "K2": "v2"} => ["K1", "v1", "K2", "v2"]
func Map2Slice(data any) []any {
	var (
		reflectValue = reflect.ValueOf(data)
		reflectKind  = reflectValue.Kind()
	)
	for reflectKind == reflect.Ptr {
		reflectValue = reflectValue.Elem()
		reflectKind = reflectValue.Kind()
	}
	switch reflectKind {
	case reflect.Map:
		array := make([]any, 0)
		for _, key := range reflectValue.MapKeys() {
			array = append(array, key.Interface())
			array = append(array, reflectValue.MapIndex(key).Interface())
		}
		return array
	default:

	}
	return nil
}

// Slice2Map ["K1", "v1", "K2", "v2"] => {"K1": "v1", "K2": "v2"}
// ["K1", "v1", "K2"]       => nil
func Slice2Map(slice any) map[any]any {
	var (
		reflectValue = reflect.ValueOf(slice)
		reflectKind  = reflectValue.Kind()
	)
	for reflectKind == reflect.Ptr {
		reflectValue = reflectValue.Elem()
		reflectKind = reflectValue.Kind()
	}
	switch reflectKind {
	case reflect.Slice, reflect.Array:
		length := reflectValue.Len()
		if length%2 != 0 {
			return nil
		}
		data := make(map[any]any)
		for i := 0; i < reflectValue.Len(); i += 2 {
			data[reflectValue.Index(i).Interface()] = reflectValue.Index(i + 1).Interface()
		}
		return data
	default:

	}
	return nil
}

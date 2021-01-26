/*
 * Collection utility of LifeCycle Struct
 *
 * Generated by: Go Streamer
 * Author
 */

package cios

import (
	"fmt"
	"math"
	"reflect"
	"sort"
)

type LifeCycleStream []LifeCycle

func LifeCycleStreamOf(arg ...LifeCycle) LifeCycleStream {
	return arg
}
func LifeCycleStreamFrom(arg []LifeCycle) LifeCycleStream {
	return arg
}
func CreateLifeCycleStream(arg ...LifeCycle) *LifeCycleStream {
	tmp := LifeCycleStreamOf(arg...)
	return &tmp
}
func GenerateLifeCycleStream(arg []LifeCycle) *LifeCycleStream {
	tmp := LifeCycleStreamFrom(arg)
	return &tmp
}

func (self *LifeCycleStream) Add(arg LifeCycle) *LifeCycleStream {
	return self.AddAll(arg)
}
func (self *LifeCycleStream) AddAll(arg ...LifeCycle) *LifeCycleStream {
	*self = append(*self, arg...)
	return self
}
func (self *LifeCycleStream) AddSafe(arg *LifeCycle) *LifeCycleStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *LifeCycleStream) Aggregate(fn func(LifeCycle, LifeCycle) LifeCycle) *LifeCycleStream {
	result := LifeCycleStreamOf()
	self.ForEach(func(v LifeCycle, i int) {
		if i == 0 {
			result.Add(fn(LifeCycle{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *LifeCycleStream) AllMatch(fn func(LifeCycle, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *LifeCycleStream) AnyMatch(fn func(LifeCycle, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *LifeCycleStream) Clone() *LifeCycleStream {
	temp := make([]LifeCycle, self.Len())
	copy(temp, *self)
	return (*LifeCycleStream)(&temp)
}
func (self *LifeCycleStream) Copy() *LifeCycleStream {
	return self.Clone()
}
func (self *LifeCycleStream) Concat(arg []LifeCycle) *LifeCycleStream {
	return self.AddAll(arg...)
}
func (self *LifeCycleStream) Contains(arg LifeCycle) bool {
	return self.FindIndex(func(_arg LifeCycle, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *LifeCycleStream) Clean() *LifeCycleStream {
	*self = LifeCycleStreamOf()
	return self
}
func (self *LifeCycleStream) Delete(index int) *LifeCycleStream {
	return self.DeleteRange(index, index)
}
func (self *LifeCycleStream) DeleteRange(startIndex, endIndex int) *LifeCycleStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *LifeCycleStream) Distinct() *LifeCycleStream {
	caches := map[string]bool{}
	result := LifeCycleStreamOf()
	for _, v := range *self {
		key := fmt.Sprintf("%+v", v)
		if f, ok := caches[key]; ok {
			if !f {
				result = append(result, v)
			}
		} else if caches[key] = true; !f {
			result = append(result, v)
		}

	}
	*self = result
	return self
}
func (self *LifeCycleStream) Each(fn func(LifeCycle)) *LifeCycleStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *LifeCycleStream) EachRight(fn func(LifeCycle)) *LifeCycleStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *LifeCycleStream) Equals(arr []LifeCycle) bool {
	if (*self == nil) != (arr == nil) || len(*self) != len(arr) {
		return false
	}
	for i := range *self {
		if !reflect.DeepEqual((*self)[i], arr[i]) {
			return false
		}
	}
	return true
}
func (self *LifeCycleStream) Filter(fn func(LifeCycle, int) bool) *LifeCycleStream {
	result := LifeCycleStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *LifeCycleStream) FilterSlim(fn func(LifeCycle, int) bool) *LifeCycleStream {
	result := LifeCycleStreamOf()
	caches := map[string]bool{}
	for i, v := range *self {
		key := fmt.Sprintf("%+v", v)
		if f, ok := caches[key]; ok {
			if f {
				result.Add(v)
			}
		} else if caches[key] = fn(v, i); caches[key] {
			result.Add(v)

		}
	}
	*self = result
	return self
}
func (self *LifeCycleStream) Find(fn func(LifeCycle, int) bool) *LifeCycle {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *LifeCycleStream) FindOr(fn func(LifeCycle, int) bool, or LifeCycle) LifeCycle {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *LifeCycleStream) FindIndex(fn func(LifeCycle, int) bool) int {
	if self == nil {
		return -1
	}
	for i, v := range *self {
		if fn(v, i) {
			return i
		}
	}
	return -1
}
func (self *LifeCycleStream) First() *LifeCycle {
	return self.Get(0)
}
func (self *LifeCycleStream) FirstOr(arg LifeCycle) LifeCycle {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *LifeCycleStream) ForEach(fn func(LifeCycle, int)) *LifeCycleStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *LifeCycleStream) ForEachRight(fn func(LifeCycle, int)) *LifeCycleStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *LifeCycleStream) GroupBy(fn func(LifeCycle, int) string) map[string][]LifeCycle {
	m := map[string][]LifeCycle{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *LifeCycleStream) GroupByValues(fn func(LifeCycle, int) string) [][]LifeCycle {
	var tmp [][]LifeCycle
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *LifeCycleStream) IndexOf(arg LifeCycle) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *LifeCycleStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *LifeCycleStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *LifeCycleStream) Last() *LifeCycle {
	return self.Get(self.Len() - 1)
}
func (self *LifeCycleStream) LastOr(arg LifeCycle) LifeCycle {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *LifeCycleStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *LifeCycleStream) Limit(limit int) *LifeCycleStream {
	self.Slice(0, limit)
	return self
}

func (self *LifeCycleStream) Map(fn func(LifeCycle, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *LifeCycleStream) Map2Int(fn func(LifeCycle, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *LifeCycleStream) Map2Int32(fn func(LifeCycle, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *LifeCycleStream) Map2Int64(fn func(LifeCycle, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *LifeCycleStream) Map2Float32(fn func(LifeCycle, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *LifeCycleStream) Map2Float64(fn func(LifeCycle, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *LifeCycleStream) Map2Bool(fn func(LifeCycle, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *LifeCycleStream) Map2Bytes(fn func(LifeCycle, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *LifeCycleStream) Map2String(fn func(LifeCycle, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *LifeCycleStream) Max(fn func(LifeCycle, int) float64) *LifeCycle {
	f := self.Get(0)
	if f == nil {
		return nil
	}
	m := fn(*f, 0)
	index := 0
	for i := 1; i < self.Len(); i++ {
		v := fn(*self.Get(i), i)
		m = math.Max(m, v)
		if m == v {
			index = i
		}
	}
	return self.Get(index)
}
func (self *LifeCycleStream) Min(fn func(LifeCycle, int) float64) *LifeCycle {
	f := self.Get(0)
	if f == nil {
		return nil
	}
	m := fn(*f, 0)
	index := 0
	for i := 1; i < self.Len(); i++ {
		v := fn(*self.Get(i), i)
		m = math.Min(m, v)
		if m == v {
			index = i
		}
	}
	return self.Get(index)
}
func (self *LifeCycleStream) NoneMatch(fn func(LifeCycle, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *LifeCycleStream) Get(index int) *LifeCycle {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *LifeCycleStream) GetOr(index int, arg LifeCycle) LifeCycle {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *LifeCycleStream) Peek(fn func(*LifeCycle, int)) *LifeCycleStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *LifeCycleStream) Reduce(fn func(LifeCycle, LifeCycle, int) LifeCycle) *LifeCycleStream {
	return self.ReduceInit(fn, LifeCycle{})
}
func (self *LifeCycleStream) ReduceInit(fn func(LifeCycle, LifeCycle, int) LifeCycle, initialValue LifeCycle) *LifeCycleStream {
	result := LifeCycleStreamOf()
	self.ForEach(func(v LifeCycle, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *LifeCycleStream) ReduceInterface(fn func(interface{}, LifeCycle, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(LifeCycle{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *LifeCycleStream) ReduceString(fn func(string, LifeCycle, int) string) []string {
	result := []string{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn("", v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *LifeCycleStream) ReduceInt(fn func(int, LifeCycle, int) int) []int {
	result := []int{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *LifeCycleStream) ReduceInt32(fn func(int32, LifeCycle, int) int32) []int32 {
	result := []int32{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *LifeCycleStream) ReduceInt64(fn func(int64, LifeCycle, int) int64) []int64 {
	result := []int64{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *LifeCycleStream) ReduceFloat32(fn func(float32, LifeCycle, int) float32) []float32 {
	result := []float32{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0.0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *LifeCycleStream) ReduceFloat64(fn func(float64, LifeCycle, int) float64) []float64 {
	result := []float64{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0.0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *LifeCycleStream) ReduceBool(fn func(bool, LifeCycle, int) bool) []bool {
	result := []bool{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(false, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *LifeCycleStream) Reverse() *LifeCycleStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *LifeCycleStream) Replace(fn func(LifeCycle, int) LifeCycle) *LifeCycleStream {
	return self.ForEach(func(v LifeCycle, i int) { self.Set(i, fn(v, i)) })
}
func (self *LifeCycleStream) Select(fn func(LifeCycle) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *LifeCycleStream) Set(index int, val LifeCycle) *LifeCycleStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *LifeCycleStream) Skip(skip int) *LifeCycleStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *LifeCycleStream) SkippingEach(fn func(LifeCycle, int) int) *LifeCycleStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *LifeCycleStream) Slice(startIndex, n int) *LifeCycleStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []LifeCycle{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *LifeCycleStream) Sort(fn func(i, j int) bool) *LifeCycleStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *LifeCycleStream) Tail() *LifeCycle {
	return self.Last()
}
func (self *LifeCycleStream) TailOr(arg LifeCycle) LifeCycle {
	return self.LastOr(arg)
}
func (self *LifeCycleStream) ToList() []LifeCycle {
	return self.Val()
}
func (self *LifeCycleStream) Unique() *LifeCycleStream {
	return self.Distinct()
}
func (self *LifeCycleStream) Val() []LifeCycle {
	if self == nil {
		return []LifeCycle{}
	}
	return *self.Copy()
}
func (self *LifeCycleStream) While(fn func(LifeCycle, int) bool) *LifeCycleStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *LifeCycleStream) Where(fn func(LifeCycle) bool) *LifeCycleStream {
	result := LifeCycleStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *LifeCycleStream) WhereSlim(fn func(LifeCycle) bool) *LifeCycleStream {
	result := LifeCycleStreamOf()
	caches := map[string]bool{}
	for _, v := range *self {
		key := fmt.Sprintf("%+v", v)
		if f, ok := caches[key]; ok {
			if f {
				result.Add(v)
			}
		} else if caches[key] = fn(v); caches[key] {
			result.Add(v)
		}
	}
	*self = result
	return self
}
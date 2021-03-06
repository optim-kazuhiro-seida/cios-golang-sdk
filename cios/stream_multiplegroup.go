/*
 * Collection utility of MultipleGroup Struct
 *
 * Generated by: Go Streamer
 */

package cios

import (
	"fmt"
	"math"
	"reflect"
	"sort"
)

type MultipleGroupStream []MultipleGroup

func MultipleGroupStreamOf(arg ...MultipleGroup) MultipleGroupStream {
	return arg
}
func MultipleGroupStreamFrom(arg []MultipleGroup) MultipleGroupStream {
	return arg
}
func CreateMultipleGroupStream(arg ...MultipleGroup) *MultipleGroupStream {
	tmp := MultipleGroupStreamOf(arg...)
	return &tmp
}
func GenerateMultipleGroupStream(arg []MultipleGroup) *MultipleGroupStream {
	tmp := MultipleGroupStreamFrom(arg)
	return &tmp
}

func (self *MultipleGroupStream) Add(arg MultipleGroup) *MultipleGroupStream {
	return self.AddAll(arg)
}
func (self *MultipleGroupStream) AddAll(arg ...MultipleGroup) *MultipleGroupStream {
	*self = append(*self, arg...)
	return self
}
func (self *MultipleGroupStream) AddSafe(arg *MultipleGroup) *MultipleGroupStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *MultipleGroupStream) Aggregate(fn func(MultipleGroup, MultipleGroup) MultipleGroup) *MultipleGroupStream {
	result := MultipleGroupStreamOf()
	self.ForEach(func(v MultipleGroup, i int) {
		if i == 0 {
			result.Add(fn(MultipleGroup{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *MultipleGroupStream) AllMatch(fn func(MultipleGroup, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *MultipleGroupStream) AnyMatch(fn func(MultipleGroup, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *MultipleGroupStream) Clone() *MultipleGroupStream {
	temp := make([]MultipleGroup, self.Len())
	copy(temp, *self)
	return (*MultipleGroupStream)(&temp)
}
func (self *MultipleGroupStream) Copy() *MultipleGroupStream {
	return self.Clone()
}
func (self *MultipleGroupStream) Concat(arg []MultipleGroup) *MultipleGroupStream {
	return self.AddAll(arg...)
}
func (self *MultipleGroupStream) Contains(arg MultipleGroup) bool {
	return self.FindIndex(func(_arg MultipleGroup, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *MultipleGroupStream) Clean() *MultipleGroupStream {
	*self = MultipleGroupStreamOf()
	return self
}
func (self *MultipleGroupStream) Delete(index int) *MultipleGroupStream {
	return self.DeleteRange(index, index)
}
func (self *MultipleGroupStream) DeleteRange(startIndex, endIndex int) *MultipleGroupStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *MultipleGroupStream) Distinct() *MultipleGroupStream {
	caches := map[string]bool{}
	result := MultipleGroupStreamOf()
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
func (self *MultipleGroupStream) Each(fn func(MultipleGroup)) *MultipleGroupStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *MultipleGroupStream) EachRight(fn func(MultipleGroup)) *MultipleGroupStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *MultipleGroupStream) Equals(arr []MultipleGroup) bool {
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
func (self *MultipleGroupStream) Filter(fn func(MultipleGroup, int) bool) *MultipleGroupStream {
	result := MultipleGroupStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *MultipleGroupStream) FilterSlim(fn func(MultipleGroup, int) bool) *MultipleGroupStream {
	result := MultipleGroupStreamOf()
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
func (self *MultipleGroupStream) Find(fn func(MultipleGroup, int) bool) *MultipleGroup {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *MultipleGroupStream) FindOr(fn func(MultipleGroup, int) bool, or MultipleGroup) MultipleGroup {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *MultipleGroupStream) FindIndex(fn func(MultipleGroup, int) bool) int {
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
func (self *MultipleGroupStream) First() *MultipleGroup {
	return self.Get(0)
}
func (self *MultipleGroupStream) FirstOr(arg MultipleGroup) MultipleGroup {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *MultipleGroupStream) ForEach(fn func(MultipleGroup, int)) *MultipleGroupStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *MultipleGroupStream) ForEachRight(fn func(MultipleGroup, int)) *MultipleGroupStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *MultipleGroupStream) GroupBy(fn func(MultipleGroup, int) string) map[string][]MultipleGroup {
	m := map[string][]MultipleGroup{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *MultipleGroupStream) GroupByValues(fn func(MultipleGroup, int) string) [][]MultipleGroup {
	var tmp [][]MultipleGroup
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *MultipleGroupStream) IndexOf(arg MultipleGroup) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *MultipleGroupStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *MultipleGroupStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *MultipleGroupStream) Last() *MultipleGroup {
	return self.Get(self.Len() - 1)
}
func (self *MultipleGroupStream) LastOr(arg MultipleGroup) MultipleGroup {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *MultipleGroupStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *MultipleGroupStream) Limit(limit int) *MultipleGroupStream {
	self.Slice(0, limit)
	return self
}

func (self *MultipleGroupStream) Map(fn func(MultipleGroup, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleGroupStream) Map2Int(fn func(MultipleGroup, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleGroupStream) Map2Int32(fn func(MultipleGroup, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleGroupStream) Map2Int64(fn func(MultipleGroup, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleGroupStream) Map2Float32(fn func(MultipleGroup, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleGroupStream) Map2Float64(fn func(MultipleGroup, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleGroupStream) Map2Bool(fn func(MultipleGroup, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleGroupStream) Map2Bytes(fn func(MultipleGroup, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleGroupStream) Map2String(fn func(MultipleGroup, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleGroupStream) Max(fn func(MultipleGroup, int) float64) *MultipleGroup {
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
func (self *MultipleGroupStream) Min(fn func(MultipleGroup, int) float64) *MultipleGroup {
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
func (self *MultipleGroupStream) NoneMatch(fn func(MultipleGroup, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *MultipleGroupStream) Get(index int) *MultipleGroup {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *MultipleGroupStream) GetOr(index int, arg MultipleGroup) MultipleGroup {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *MultipleGroupStream) Peek(fn func(*MultipleGroup, int)) *MultipleGroupStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *MultipleGroupStream) Reduce(fn func(MultipleGroup, MultipleGroup, int) MultipleGroup) *MultipleGroupStream {
	return self.ReduceInit(fn, MultipleGroup{})
}
func (self *MultipleGroupStream) ReduceInit(fn func(MultipleGroup, MultipleGroup, int) MultipleGroup, initialValue MultipleGroup) *MultipleGroupStream {
	result := MultipleGroupStreamOf()
	self.ForEach(func(v MultipleGroup, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *MultipleGroupStream) ReduceInterface(fn func(interface{}, MultipleGroup, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(MultipleGroup{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *MultipleGroupStream) ReduceString(fn func(string, MultipleGroup, int) string) []string {
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
func (self *MultipleGroupStream) ReduceInt(fn func(int, MultipleGroup, int) int) []int {
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
func (self *MultipleGroupStream) ReduceInt32(fn func(int32, MultipleGroup, int) int32) []int32 {
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
func (self *MultipleGroupStream) ReduceInt64(fn func(int64, MultipleGroup, int) int64) []int64 {
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
func (self *MultipleGroupStream) ReduceFloat32(fn func(float32, MultipleGroup, int) float32) []float32 {
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
func (self *MultipleGroupStream) ReduceFloat64(fn func(float64, MultipleGroup, int) float64) []float64 {
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
func (self *MultipleGroupStream) ReduceBool(fn func(bool, MultipleGroup, int) bool) []bool {
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
func (self *MultipleGroupStream) Reverse() *MultipleGroupStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *MultipleGroupStream) Replace(fn func(MultipleGroup, int) MultipleGroup) *MultipleGroupStream {
	return self.ForEach(func(v MultipleGroup, i int) { self.Set(i, fn(v, i)) })
}
func (self *MultipleGroupStream) Select(fn func(MultipleGroup) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *MultipleGroupStream) Set(index int, val MultipleGroup) *MultipleGroupStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *MultipleGroupStream) Skip(skip int) *MultipleGroupStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *MultipleGroupStream) SkippingEach(fn func(MultipleGroup, int) int) *MultipleGroupStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *MultipleGroupStream) Slice(startIndex, n int) *MultipleGroupStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []MultipleGroup{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *MultipleGroupStream) Sort(fn func(i, j int) bool) *MultipleGroupStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *MultipleGroupStream) Tail() *MultipleGroup {
	return self.Last()
}
func (self *MultipleGroupStream) TailOr(arg MultipleGroup) MultipleGroup {
	return self.LastOr(arg)
}
func (self *MultipleGroupStream) ToList() []MultipleGroup {
	return self.Val()
}
func (self *MultipleGroupStream) Unique() *MultipleGroupStream {
	return self.Distinct()
}
func (self *MultipleGroupStream) Val() []MultipleGroup {
	if self == nil {
		return []MultipleGroup{}
	}
	return *self.Copy()
}
func (self *MultipleGroupStream) While(fn func(MultipleGroup, int) bool) *MultipleGroupStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *MultipleGroupStream) Where(fn func(MultipleGroup) bool) *MultipleGroupStream {
	result := MultipleGroupStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *MultipleGroupStream) WhereSlim(fn func(MultipleGroup) bool) *MultipleGroupStream {
	result := MultipleGroupStreamOf()
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

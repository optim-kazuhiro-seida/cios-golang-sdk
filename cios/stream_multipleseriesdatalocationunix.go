/*
 * Collection utility of MultipleSeriesDataLocationUnix Struct
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

type MultipleSeriesDataLocationUnixStream []MultipleSeriesDataLocationUnix

func MultipleSeriesDataLocationUnixStreamOf(arg ...MultipleSeriesDataLocationUnix) MultipleSeriesDataLocationUnixStream {
	return arg
}
func MultipleSeriesDataLocationUnixStreamFrom(arg []MultipleSeriesDataLocationUnix) MultipleSeriesDataLocationUnixStream {
	return arg
}
func CreateMultipleSeriesDataLocationUnixStream(arg ...MultipleSeriesDataLocationUnix) *MultipleSeriesDataLocationUnixStream {
	tmp := MultipleSeriesDataLocationUnixStreamOf(arg...)
	return &tmp
}
func GenerateMultipleSeriesDataLocationUnixStream(arg []MultipleSeriesDataLocationUnix) *MultipleSeriesDataLocationUnixStream {
	tmp := MultipleSeriesDataLocationUnixStreamFrom(arg)
	return &tmp
}

func (self *MultipleSeriesDataLocationUnixStream) Add(arg MultipleSeriesDataLocationUnix) *MultipleSeriesDataLocationUnixStream {
	return self.AddAll(arg)
}
func (self *MultipleSeriesDataLocationUnixStream) AddAll(arg ...MultipleSeriesDataLocationUnix) *MultipleSeriesDataLocationUnixStream {
	*self = append(*self, arg...)
	return self
}
func (self *MultipleSeriesDataLocationUnixStream) AddSafe(arg *MultipleSeriesDataLocationUnix) *MultipleSeriesDataLocationUnixStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *MultipleSeriesDataLocationUnixStream) Aggregate(fn func(MultipleSeriesDataLocationUnix, MultipleSeriesDataLocationUnix) MultipleSeriesDataLocationUnix) *MultipleSeriesDataLocationUnixStream {
	result := MultipleSeriesDataLocationUnixStreamOf()
	self.ForEach(func(v MultipleSeriesDataLocationUnix, i int) {
		if i == 0 {
			result.Add(fn(MultipleSeriesDataLocationUnix{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *MultipleSeriesDataLocationUnixStream) AllMatch(fn func(MultipleSeriesDataLocationUnix, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *MultipleSeriesDataLocationUnixStream) AnyMatch(fn func(MultipleSeriesDataLocationUnix, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *MultipleSeriesDataLocationUnixStream) Clone() *MultipleSeriesDataLocationUnixStream {
	temp := make([]MultipleSeriesDataLocationUnix, self.Len())
	copy(temp, *self)
	return (*MultipleSeriesDataLocationUnixStream)(&temp)
}
func (self *MultipleSeriesDataLocationUnixStream) Copy() *MultipleSeriesDataLocationUnixStream {
	return self.Clone()
}
func (self *MultipleSeriesDataLocationUnixStream) Concat(arg []MultipleSeriesDataLocationUnix) *MultipleSeriesDataLocationUnixStream {
	return self.AddAll(arg...)
}
func (self *MultipleSeriesDataLocationUnixStream) Contains(arg MultipleSeriesDataLocationUnix) bool {
	return self.FindIndex(func(_arg MultipleSeriesDataLocationUnix, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *MultipleSeriesDataLocationUnixStream) Clean() *MultipleSeriesDataLocationUnixStream {
	*self = MultipleSeriesDataLocationUnixStreamOf()
	return self
}
func (self *MultipleSeriesDataLocationUnixStream) Delete(index int) *MultipleSeriesDataLocationUnixStream {
	return self.DeleteRange(index, index)
}
func (self *MultipleSeriesDataLocationUnixStream) DeleteRange(startIndex, endIndex int) *MultipleSeriesDataLocationUnixStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *MultipleSeriesDataLocationUnixStream) Distinct() *MultipleSeriesDataLocationUnixStream {
	caches := map[string]bool{}
	result := MultipleSeriesDataLocationUnixStreamOf()
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
func (self *MultipleSeriesDataLocationUnixStream) Each(fn func(MultipleSeriesDataLocationUnix)) *MultipleSeriesDataLocationUnixStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *MultipleSeriesDataLocationUnixStream) EachRight(fn func(MultipleSeriesDataLocationUnix)) *MultipleSeriesDataLocationUnixStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *MultipleSeriesDataLocationUnixStream) Equals(arr []MultipleSeriesDataLocationUnix) bool {
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
func (self *MultipleSeriesDataLocationUnixStream) Filter(fn func(MultipleSeriesDataLocationUnix, int) bool) *MultipleSeriesDataLocationUnixStream {
	result := MultipleSeriesDataLocationUnixStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *MultipleSeriesDataLocationUnixStream) FilterSlim(fn func(MultipleSeriesDataLocationUnix, int) bool) *MultipleSeriesDataLocationUnixStream {
	result := MultipleSeriesDataLocationUnixStreamOf()
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
func (self *MultipleSeriesDataLocationUnixStream) Find(fn func(MultipleSeriesDataLocationUnix, int) bool) *MultipleSeriesDataLocationUnix {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *MultipleSeriesDataLocationUnixStream) FindOr(fn func(MultipleSeriesDataLocationUnix, int) bool, or MultipleSeriesDataLocationUnix) MultipleSeriesDataLocationUnix {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *MultipleSeriesDataLocationUnixStream) FindIndex(fn func(MultipleSeriesDataLocationUnix, int) bool) int {
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
func (self *MultipleSeriesDataLocationUnixStream) First() *MultipleSeriesDataLocationUnix {
	return self.Get(0)
}
func (self *MultipleSeriesDataLocationUnixStream) FirstOr(arg MultipleSeriesDataLocationUnix) MultipleSeriesDataLocationUnix {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *MultipleSeriesDataLocationUnixStream) ForEach(fn func(MultipleSeriesDataLocationUnix, int)) *MultipleSeriesDataLocationUnixStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *MultipleSeriesDataLocationUnixStream) ForEachRight(fn func(MultipleSeriesDataLocationUnix, int)) *MultipleSeriesDataLocationUnixStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *MultipleSeriesDataLocationUnixStream) GroupBy(fn func(MultipleSeriesDataLocationUnix, int) string) map[string][]MultipleSeriesDataLocationUnix {
	m := map[string][]MultipleSeriesDataLocationUnix{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *MultipleSeriesDataLocationUnixStream) GroupByValues(fn func(MultipleSeriesDataLocationUnix, int) string) [][]MultipleSeriesDataLocationUnix {
	var tmp [][]MultipleSeriesDataLocationUnix
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *MultipleSeriesDataLocationUnixStream) IndexOf(arg MultipleSeriesDataLocationUnix) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *MultipleSeriesDataLocationUnixStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *MultipleSeriesDataLocationUnixStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *MultipleSeriesDataLocationUnixStream) Last() *MultipleSeriesDataLocationUnix {
	return self.Get(self.Len() - 1)
}
func (self *MultipleSeriesDataLocationUnixStream) LastOr(arg MultipleSeriesDataLocationUnix) MultipleSeriesDataLocationUnix {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *MultipleSeriesDataLocationUnixStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *MultipleSeriesDataLocationUnixStream) Limit(limit int) *MultipleSeriesDataLocationUnixStream {
	self.Slice(0, limit)
	return self
}

func (self *MultipleSeriesDataLocationUnixStream) Map(fn func(MultipleSeriesDataLocationUnix, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleSeriesDataLocationUnixStream) Map2Int(fn func(MultipleSeriesDataLocationUnix, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleSeriesDataLocationUnixStream) Map2Int32(fn func(MultipleSeriesDataLocationUnix, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleSeriesDataLocationUnixStream) Map2Int64(fn func(MultipleSeriesDataLocationUnix, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleSeriesDataLocationUnixStream) Map2Float32(fn func(MultipleSeriesDataLocationUnix, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleSeriesDataLocationUnixStream) Map2Float64(fn func(MultipleSeriesDataLocationUnix, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleSeriesDataLocationUnixStream) Map2Bool(fn func(MultipleSeriesDataLocationUnix, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleSeriesDataLocationUnixStream) Map2Bytes(fn func(MultipleSeriesDataLocationUnix, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleSeriesDataLocationUnixStream) Map2String(fn func(MultipleSeriesDataLocationUnix, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleSeriesDataLocationUnixStream) Max(fn func(MultipleSeriesDataLocationUnix, int) float64) *MultipleSeriesDataLocationUnix {
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
func (self *MultipleSeriesDataLocationUnixStream) Min(fn func(MultipleSeriesDataLocationUnix, int) float64) *MultipleSeriesDataLocationUnix {
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
func (self *MultipleSeriesDataLocationUnixStream) NoneMatch(fn func(MultipleSeriesDataLocationUnix, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *MultipleSeriesDataLocationUnixStream) Get(index int) *MultipleSeriesDataLocationUnix {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *MultipleSeriesDataLocationUnixStream) GetOr(index int, arg MultipleSeriesDataLocationUnix) MultipleSeriesDataLocationUnix {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *MultipleSeriesDataLocationUnixStream) Peek(fn func(*MultipleSeriesDataLocationUnix, int)) *MultipleSeriesDataLocationUnixStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *MultipleSeriesDataLocationUnixStream) Reduce(fn func(MultipleSeriesDataLocationUnix, MultipleSeriesDataLocationUnix, int) MultipleSeriesDataLocationUnix) *MultipleSeriesDataLocationUnixStream {
	return self.ReduceInit(fn, MultipleSeriesDataLocationUnix{})
}
func (self *MultipleSeriesDataLocationUnixStream) ReduceInit(fn func(MultipleSeriesDataLocationUnix, MultipleSeriesDataLocationUnix, int) MultipleSeriesDataLocationUnix, initialValue MultipleSeriesDataLocationUnix) *MultipleSeriesDataLocationUnixStream {
	result := MultipleSeriesDataLocationUnixStreamOf()
	self.ForEach(func(v MultipleSeriesDataLocationUnix, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *MultipleSeriesDataLocationUnixStream) ReduceInterface(fn func(interface{}, MultipleSeriesDataLocationUnix, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(MultipleSeriesDataLocationUnix{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *MultipleSeriesDataLocationUnixStream) ReduceString(fn func(string, MultipleSeriesDataLocationUnix, int) string) []string {
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
func (self *MultipleSeriesDataLocationUnixStream) ReduceInt(fn func(int, MultipleSeriesDataLocationUnix, int) int) []int {
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
func (self *MultipleSeriesDataLocationUnixStream) ReduceInt32(fn func(int32, MultipleSeriesDataLocationUnix, int) int32) []int32 {
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
func (self *MultipleSeriesDataLocationUnixStream) ReduceInt64(fn func(int64, MultipleSeriesDataLocationUnix, int) int64) []int64 {
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
func (self *MultipleSeriesDataLocationUnixStream) ReduceFloat32(fn func(float32, MultipleSeriesDataLocationUnix, int) float32) []float32 {
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
func (self *MultipleSeriesDataLocationUnixStream) ReduceFloat64(fn func(float64, MultipleSeriesDataLocationUnix, int) float64) []float64 {
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
func (self *MultipleSeriesDataLocationUnixStream) ReduceBool(fn func(bool, MultipleSeriesDataLocationUnix, int) bool) []bool {
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
func (self *MultipleSeriesDataLocationUnixStream) Reverse() *MultipleSeriesDataLocationUnixStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *MultipleSeriesDataLocationUnixStream) Replace(fn func(MultipleSeriesDataLocationUnix, int) MultipleSeriesDataLocationUnix) *MultipleSeriesDataLocationUnixStream {
	return self.ForEach(func(v MultipleSeriesDataLocationUnix, i int) { self.Set(i, fn(v, i)) })
}
func (self *MultipleSeriesDataLocationUnixStream) Select(fn func(MultipleSeriesDataLocationUnix) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *MultipleSeriesDataLocationUnixStream) Set(index int, val MultipleSeriesDataLocationUnix) *MultipleSeriesDataLocationUnixStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *MultipleSeriesDataLocationUnixStream) Skip(skip int) *MultipleSeriesDataLocationUnixStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *MultipleSeriesDataLocationUnixStream) SkippingEach(fn func(MultipleSeriesDataLocationUnix, int) int) *MultipleSeriesDataLocationUnixStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *MultipleSeriesDataLocationUnixStream) Slice(startIndex, n int) *MultipleSeriesDataLocationUnixStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []MultipleSeriesDataLocationUnix{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *MultipleSeriesDataLocationUnixStream) Sort(fn func(i, j int) bool) *MultipleSeriesDataLocationUnixStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *MultipleSeriesDataLocationUnixStream) Tail() *MultipleSeriesDataLocationUnix {
	return self.Last()
}
func (self *MultipleSeriesDataLocationUnixStream) TailOr(arg MultipleSeriesDataLocationUnix) MultipleSeriesDataLocationUnix {
	return self.LastOr(arg)
}
func (self *MultipleSeriesDataLocationUnixStream) ToList() []MultipleSeriesDataLocationUnix {
	return self.Val()
}
func (self *MultipleSeriesDataLocationUnixStream) Unique() *MultipleSeriesDataLocationUnixStream {
	return self.Distinct()
}
func (self *MultipleSeriesDataLocationUnixStream) Val() []MultipleSeriesDataLocationUnix {
	if self == nil {
		return []MultipleSeriesDataLocationUnix{}
	}
	return *self.Copy()
}
func (self *MultipleSeriesDataLocationUnixStream) While(fn func(MultipleSeriesDataLocationUnix, int) bool) *MultipleSeriesDataLocationUnixStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *MultipleSeriesDataLocationUnixStream) Where(fn func(MultipleSeriesDataLocationUnix) bool) *MultipleSeriesDataLocationUnixStream {
	result := MultipleSeriesDataLocationUnixStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *MultipleSeriesDataLocationUnixStream) WhereSlim(fn func(MultipleSeriesDataLocationUnix) bool) *MultipleSeriesDataLocationUnixStream {
	result := MultipleSeriesDataLocationUnixStreamOf()
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

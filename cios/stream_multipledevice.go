/*
 * Collection utility of MultipleDevice Struct
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

type MultipleDeviceStream []MultipleDevice

func MultipleDeviceStreamOf(arg ...MultipleDevice) MultipleDeviceStream {
	return arg
}
func MultipleDeviceStreamFrom(arg []MultipleDevice) MultipleDeviceStream {
	return arg
}
func CreateMultipleDeviceStream(arg ...MultipleDevice) *MultipleDeviceStream {
	tmp := MultipleDeviceStreamOf(arg...)
	return &tmp
}
func GenerateMultipleDeviceStream(arg []MultipleDevice) *MultipleDeviceStream {
	tmp := MultipleDeviceStreamFrom(arg)
	return &tmp
}

func (self *MultipleDeviceStream) Add(arg MultipleDevice) *MultipleDeviceStream {
	return self.AddAll(arg)
}
func (self *MultipleDeviceStream) AddAll(arg ...MultipleDevice) *MultipleDeviceStream {
	*self = append(*self, arg...)
	return self
}
func (self *MultipleDeviceStream) AddSafe(arg *MultipleDevice) *MultipleDeviceStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *MultipleDeviceStream) Aggregate(fn func(MultipleDevice, MultipleDevice) MultipleDevice) *MultipleDeviceStream {
	result := MultipleDeviceStreamOf()
	self.ForEach(func(v MultipleDevice, i int) {
		if i == 0 {
			result.Add(fn(MultipleDevice{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *MultipleDeviceStream) AllMatch(fn func(MultipleDevice, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *MultipleDeviceStream) AnyMatch(fn func(MultipleDevice, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *MultipleDeviceStream) Clone() *MultipleDeviceStream {
	temp := make([]MultipleDevice, self.Len())
	copy(temp, *self)
	return (*MultipleDeviceStream)(&temp)
}
func (self *MultipleDeviceStream) Copy() *MultipleDeviceStream {
	return self.Clone()
}
func (self *MultipleDeviceStream) Concat(arg []MultipleDevice) *MultipleDeviceStream {
	return self.AddAll(arg...)
}
func (self *MultipleDeviceStream) Contains(arg MultipleDevice) bool {
	return self.FindIndex(func(_arg MultipleDevice, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *MultipleDeviceStream) Clean() *MultipleDeviceStream {
	*self = MultipleDeviceStreamOf()
	return self
}
func (self *MultipleDeviceStream) Delete(index int) *MultipleDeviceStream {
	return self.DeleteRange(index, index)
}
func (self *MultipleDeviceStream) DeleteRange(startIndex, endIndex int) *MultipleDeviceStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *MultipleDeviceStream) Distinct() *MultipleDeviceStream {
	caches := map[string]bool{}
	result := MultipleDeviceStreamOf()
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
func (self *MultipleDeviceStream) Each(fn func(MultipleDevice)) *MultipleDeviceStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *MultipleDeviceStream) EachRight(fn func(MultipleDevice)) *MultipleDeviceStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *MultipleDeviceStream) Equals(arr []MultipleDevice) bool {
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
func (self *MultipleDeviceStream) Filter(fn func(MultipleDevice, int) bool) *MultipleDeviceStream {
	result := MultipleDeviceStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *MultipleDeviceStream) FilterSlim(fn func(MultipleDevice, int) bool) *MultipleDeviceStream {
	result := MultipleDeviceStreamOf()
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
func (self *MultipleDeviceStream) Find(fn func(MultipleDevice, int) bool) *MultipleDevice {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *MultipleDeviceStream) FindOr(fn func(MultipleDevice, int) bool, or MultipleDevice) MultipleDevice {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *MultipleDeviceStream) FindIndex(fn func(MultipleDevice, int) bool) int {
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
func (self *MultipleDeviceStream) First() *MultipleDevice {
	return self.Get(0)
}
func (self *MultipleDeviceStream) FirstOr(arg MultipleDevice) MultipleDevice {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *MultipleDeviceStream) ForEach(fn func(MultipleDevice, int)) *MultipleDeviceStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *MultipleDeviceStream) ForEachRight(fn func(MultipleDevice, int)) *MultipleDeviceStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *MultipleDeviceStream) GroupBy(fn func(MultipleDevice, int) string) map[string][]MultipleDevice {
	m := map[string][]MultipleDevice{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *MultipleDeviceStream) GroupByValues(fn func(MultipleDevice, int) string) [][]MultipleDevice {
	var tmp [][]MultipleDevice
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *MultipleDeviceStream) IndexOf(arg MultipleDevice) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *MultipleDeviceStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *MultipleDeviceStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *MultipleDeviceStream) Last() *MultipleDevice {
	return self.Get(self.Len() - 1)
}
func (self *MultipleDeviceStream) LastOr(arg MultipleDevice) MultipleDevice {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *MultipleDeviceStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *MultipleDeviceStream) Limit(limit int) *MultipleDeviceStream {
	self.Slice(0, limit)
	return self
}

func (self *MultipleDeviceStream) Map(fn func(MultipleDevice, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDeviceStream) Map2Int(fn func(MultipleDevice, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDeviceStream) Map2Int32(fn func(MultipleDevice, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDeviceStream) Map2Int64(fn func(MultipleDevice, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDeviceStream) Map2Float32(fn func(MultipleDevice, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDeviceStream) Map2Float64(fn func(MultipleDevice, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDeviceStream) Map2Bool(fn func(MultipleDevice, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDeviceStream) Map2Bytes(fn func(MultipleDevice, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDeviceStream) Map2String(fn func(MultipleDevice, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDeviceStream) Max(fn func(MultipleDevice, int) float64) *MultipleDevice {
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
func (self *MultipleDeviceStream) Min(fn func(MultipleDevice, int) float64) *MultipleDevice {
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
func (self *MultipleDeviceStream) NoneMatch(fn func(MultipleDevice, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *MultipleDeviceStream) Get(index int) *MultipleDevice {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *MultipleDeviceStream) GetOr(index int, arg MultipleDevice) MultipleDevice {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *MultipleDeviceStream) Peek(fn func(*MultipleDevice, int)) *MultipleDeviceStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *MultipleDeviceStream) Reduce(fn func(MultipleDevice, MultipleDevice, int) MultipleDevice) *MultipleDeviceStream {
	return self.ReduceInit(fn, MultipleDevice{})
}
func (self *MultipleDeviceStream) ReduceInit(fn func(MultipleDevice, MultipleDevice, int) MultipleDevice, initialValue MultipleDevice) *MultipleDeviceStream {
	result := MultipleDeviceStreamOf()
	self.ForEach(func(v MultipleDevice, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *MultipleDeviceStream) ReduceInterface(fn func(interface{}, MultipleDevice, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(MultipleDevice{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *MultipleDeviceStream) ReduceString(fn func(string, MultipleDevice, int) string) []string {
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
func (self *MultipleDeviceStream) ReduceInt(fn func(int, MultipleDevice, int) int) []int {
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
func (self *MultipleDeviceStream) ReduceInt32(fn func(int32, MultipleDevice, int) int32) []int32 {
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
func (self *MultipleDeviceStream) ReduceInt64(fn func(int64, MultipleDevice, int) int64) []int64 {
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
func (self *MultipleDeviceStream) ReduceFloat32(fn func(float32, MultipleDevice, int) float32) []float32 {
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
func (self *MultipleDeviceStream) ReduceFloat64(fn func(float64, MultipleDevice, int) float64) []float64 {
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
func (self *MultipleDeviceStream) ReduceBool(fn func(bool, MultipleDevice, int) bool) []bool {
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
func (self *MultipleDeviceStream) Reverse() *MultipleDeviceStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *MultipleDeviceStream) Replace(fn func(MultipleDevice, int) MultipleDevice) *MultipleDeviceStream {
	return self.ForEach(func(v MultipleDevice, i int) { self.Set(i, fn(v, i)) })
}
func (self *MultipleDeviceStream) Select(fn func(MultipleDevice) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *MultipleDeviceStream) Set(index int, val MultipleDevice) *MultipleDeviceStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *MultipleDeviceStream) Skip(skip int) *MultipleDeviceStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *MultipleDeviceStream) SkippingEach(fn func(MultipleDevice, int) int) *MultipleDeviceStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *MultipleDeviceStream) Slice(startIndex, n int) *MultipleDeviceStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []MultipleDevice{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *MultipleDeviceStream) Sort(fn func(i, j int) bool) *MultipleDeviceStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *MultipleDeviceStream) Tail() *MultipleDevice {
	return self.Last()
}
func (self *MultipleDeviceStream) TailOr(arg MultipleDevice) MultipleDevice {
	return self.LastOr(arg)
}
func (self *MultipleDeviceStream) ToList() []MultipleDevice {
	return self.Val()
}
func (self *MultipleDeviceStream) Unique() *MultipleDeviceStream {
	return self.Distinct()
}
func (self *MultipleDeviceStream) Val() []MultipleDevice {
	if self == nil {
		return []MultipleDevice{}
	}
	return *self.Copy()
}
func (self *MultipleDeviceStream) While(fn func(MultipleDevice, int) bool) *MultipleDeviceStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *MultipleDeviceStream) Where(fn func(MultipleDevice) bool) *MultipleDeviceStream {
	result := MultipleDeviceStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *MultipleDeviceStream) WhereSlim(fn func(MultipleDevice) bool) *MultipleDeviceStream {
	result := MultipleDeviceStreamOf()
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
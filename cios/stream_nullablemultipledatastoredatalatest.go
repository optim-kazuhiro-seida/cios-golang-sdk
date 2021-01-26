/*
 * Collection utility of NullableMultipleDataStoreDataLatest Struct
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

type NullableMultipleDataStoreDataLatestStream []NullableMultipleDataStoreDataLatest

func NullableMultipleDataStoreDataLatestStreamOf(arg ...NullableMultipleDataStoreDataLatest) NullableMultipleDataStoreDataLatestStream {
	return arg
}
func NullableMultipleDataStoreDataLatestStreamFrom(arg []NullableMultipleDataStoreDataLatest) NullableMultipleDataStoreDataLatestStream {
	return arg
}
func CreateNullableMultipleDataStoreDataLatestStream(arg ...NullableMultipleDataStoreDataLatest) *NullableMultipleDataStoreDataLatestStream {
	tmp := NullableMultipleDataStoreDataLatestStreamOf(arg...)
	return &tmp
}
func GenerateNullableMultipleDataStoreDataLatestStream(arg []NullableMultipleDataStoreDataLatest) *NullableMultipleDataStoreDataLatestStream {
	tmp := NullableMultipleDataStoreDataLatestStreamFrom(arg)
	return &tmp
}

func (self *NullableMultipleDataStoreDataLatestStream) Add(arg NullableMultipleDataStoreDataLatest) *NullableMultipleDataStoreDataLatestStream {
	return self.AddAll(arg)
}
func (self *NullableMultipleDataStoreDataLatestStream) AddAll(arg ...NullableMultipleDataStoreDataLatest) *NullableMultipleDataStoreDataLatestStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableMultipleDataStoreDataLatestStream) AddSafe(arg *NullableMultipleDataStoreDataLatest) *NullableMultipleDataStoreDataLatestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableMultipleDataStoreDataLatestStream) Aggregate(fn func(NullableMultipleDataStoreDataLatest, NullableMultipleDataStoreDataLatest) NullableMultipleDataStoreDataLatest) *NullableMultipleDataStoreDataLatestStream {
	result := NullableMultipleDataStoreDataLatestStreamOf()
	self.ForEach(func(v NullableMultipleDataStoreDataLatest, i int) {
		if i == 0 {
			result.Add(fn(NullableMultipleDataStoreDataLatest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableMultipleDataStoreDataLatestStream) AllMatch(fn func(NullableMultipleDataStoreDataLatest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableMultipleDataStoreDataLatestStream) AnyMatch(fn func(NullableMultipleDataStoreDataLatest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableMultipleDataStoreDataLatestStream) Clone() *NullableMultipleDataStoreDataLatestStream {
	temp := make([]NullableMultipleDataStoreDataLatest, self.Len())
	copy(temp, *self)
	return (*NullableMultipleDataStoreDataLatestStream)(&temp)
}
func (self *NullableMultipleDataStoreDataLatestStream) Copy() *NullableMultipleDataStoreDataLatestStream {
	return self.Clone()
}
func (self *NullableMultipleDataStoreDataLatestStream) Concat(arg []NullableMultipleDataStoreDataLatest) *NullableMultipleDataStoreDataLatestStream {
	return self.AddAll(arg...)
}
func (self *NullableMultipleDataStoreDataLatestStream) Contains(arg NullableMultipleDataStoreDataLatest) bool {
	return self.FindIndex(func(_arg NullableMultipleDataStoreDataLatest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableMultipleDataStoreDataLatestStream) Clean() *NullableMultipleDataStoreDataLatestStream {
	*self = NullableMultipleDataStoreDataLatestStreamOf()
	return self
}
func (self *NullableMultipleDataStoreDataLatestStream) Delete(index int) *NullableMultipleDataStoreDataLatestStream {
	return self.DeleteRange(index, index)
}
func (self *NullableMultipleDataStoreDataLatestStream) DeleteRange(startIndex, endIndex int) *NullableMultipleDataStoreDataLatestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableMultipleDataStoreDataLatestStream) Distinct() *NullableMultipleDataStoreDataLatestStream {
	caches := map[string]bool{}
	result := NullableMultipleDataStoreDataLatestStreamOf()
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
func (self *NullableMultipleDataStoreDataLatestStream) Each(fn func(NullableMultipleDataStoreDataLatest)) *NullableMultipleDataStoreDataLatestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableMultipleDataStoreDataLatestStream) EachRight(fn func(NullableMultipleDataStoreDataLatest)) *NullableMultipleDataStoreDataLatestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableMultipleDataStoreDataLatestStream) Equals(arr []NullableMultipleDataStoreDataLatest) bool {
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
func (self *NullableMultipleDataStoreDataLatestStream) Filter(fn func(NullableMultipleDataStoreDataLatest, int) bool) *NullableMultipleDataStoreDataLatestStream {
	result := NullableMultipleDataStoreDataLatestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMultipleDataStoreDataLatestStream) FilterSlim(fn func(NullableMultipleDataStoreDataLatest, int) bool) *NullableMultipleDataStoreDataLatestStream {
	result := NullableMultipleDataStoreDataLatestStreamOf()
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
func (self *NullableMultipleDataStoreDataLatestStream) Find(fn func(NullableMultipleDataStoreDataLatest, int) bool) *NullableMultipleDataStoreDataLatest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableMultipleDataStoreDataLatestStream) FindOr(fn func(NullableMultipleDataStoreDataLatest, int) bool, or NullableMultipleDataStoreDataLatest) NullableMultipleDataStoreDataLatest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableMultipleDataStoreDataLatestStream) FindIndex(fn func(NullableMultipleDataStoreDataLatest, int) bool) int {
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
func (self *NullableMultipleDataStoreDataLatestStream) First() *NullableMultipleDataStoreDataLatest {
	return self.Get(0)
}
func (self *NullableMultipleDataStoreDataLatestStream) FirstOr(arg NullableMultipleDataStoreDataLatest) NullableMultipleDataStoreDataLatest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleDataStoreDataLatestStream) ForEach(fn func(NullableMultipleDataStoreDataLatest, int)) *NullableMultipleDataStoreDataLatestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableMultipleDataStoreDataLatestStream) ForEachRight(fn func(NullableMultipleDataStoreDataLatest, int)) *NullableMultipleDataStoreDataLatestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableMultipleDataStoreDataLatestStream) GroupBy(fn func(NullableMultipleDataStoreDataLatest, int) string) map[string][]NullableMultipleDataStoreDataLatest {
	m := map[string][]NullableMultipleDataStoreDataLatest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableMultipleDataStoreDataLatestStream) GroupByValues(fn func(NullableMultipleDataStoreDataLatest, int) string) [][]NullableMultipleDataStoreDataLatest {
	var tmp [][]NullableMultipleDataStoreDataLatest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableMultipleDataStoreDataLatestStream) IndexOf(arg NullableMultipleDataStoreDataLatest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableMultipleDataStoreDataLatestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableMultipleDataStoreDataLatestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableMultipleDataStoreDataLatestStream) Last() *NullableMultipleDataStoreDataLatest {
	return self.Get(self.Len() - 1)
}
func (self *NullableMultipleDataStoreDataLatestStream) LastOr(arg NullableMultipleDataStoreDataLatest) NullableMultipleDataStoreDataLatest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleDataStoreDataLatestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableMultipleDataStoreDataLatestStream) Limit(limit int) *NullableMultipleDataStoreDataLatestStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableMultipleDataStoreDataLatestStream) Map(fn func(NullableMultipleDataStoreDataLatest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDataStoreDataLatestStream) Map2Int(fn func(NullableMultipleDataStoreDataLatest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDataStoreDataLatestStream) Map2Int32(fn func(NullableMultipleDataStoreDataLatest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDataStoreDataLatestStream) Map2Int64(fn func(NullableMultipleDataStoreDataLatest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDataStoreDataLatestStream) Map2Float32(fn func(NullableMultipleDataStoreDataLatest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDataStoreDataLatestStream) Map2Float64(fn func(NullableMultipleDataStoreDataLatest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDataStoreDataLatestStream) Map2Bool(fn func(NullableMultipleDataStoreDataLatest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDataStoreDataLatestStream) Map2Bytes(fn func(NullableMultipleDataStoreDataLatest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDataStoreDataLatestStream) Map2String(fn func(NullableMultipleDataStoreDataLatest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDataStoreDataLatestStream) Max(fn func(NullableMultipleDataStoreDataLatest, int) float64) *NullableMultipleDataStoreDataLatest {
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
func (self *NullableMultipleDataStoreDataLatestStream) Min(fn func(NullableMultipleDataStoreDataLatest, int) float64) *NullableMultipleDataStoreDataLatest {
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
func (self *NullableMultipleDataStoreDataLatestStream) NoneMatch(fn func(NullableMultipleDataStoreDataLatest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableMultipleDataStoreDataLatestStream) Get(index int) *NullableMultipleDataStoreDataLatest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableMultipleDataStoreDataLatestStream) GetOr(index int, arg NullableMultipleDataStoreDataLatest) NullableMultipleDataStoreDataLatest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleDataStoreDataLatestStream) Peek(fn func(*NullableMultipleDataStoreDataLatest, int)) *NullableMultipleDataStoreDataLatestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *NullableMultipleDataStoreDataLatestStream) Reduce(fn func(NullableMultipleDataStoreDataLatest, NullableMultipleDataStoreDataLatest, int) NullableMultipleDataStoreDataLatest) *NullableMultipleDataStoreDataLatestStream {
	return self.ReduceInit(fn, NullableMultipleDataStoreDataLatest{})
}
func (self *NullableMultipleDataStoreDataLatestStream) ReduceInit(fn func(NullableMultipleDataStoreDataLatest, NullableMultipleDataStoreDataLatest, int) NullableMultipleDataStoreDataLatest, initialValue NullableMultipleDataStoreDataLatest) *NullableMultipleDataStoreDataLatestStream {
	result := NullableMultipleDataStoreDataLatestStreamOf()
	self.ForEach(func(v NullableMultipleDataStoreDataLatest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableMultipleDataStoreDataLatestStream) ReduceInterface(fn func(interface{}, NullableMultipleDataStoreDataLatest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableMultipleDataStoreDataLatest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleDataStoreDataLatestStream) ReduceString(fn func(string, NullableMultipleDataStoreDataLatest, int) string) []string {
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
func (self *NullableMultipleDataStoreDataLatestStream) ReduceInt(fn func(int, NullableMultipleDataStoreDataLatest, int) int) []int {
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
func (self *NullableMultipleDataStoreDataLatestStream) ReduceInt32(fn func(int32, NullableMultipleDataStoreDataLatest, int) int32) []int32 {
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
func (self *NullableMultipleDataStoreDataLatestStream) ReduceInt64(fn func(int64, NullableMultipleDataStoreDataLatest, int) int64) []int64 {
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
func (self *NullableMultipleDataStoreDataLatestStream) ReduceFloat32(fn func(float32, NullableMultipleDataStoreDataLatest, int) float32) []float32 {
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
func (self *NullableMultipleDataStoreDataLatestStream) ReduceFloat64(fn func(float64, NullableMultipleDataStoreDataLatest, int) float64) []float64 {
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
func (self *NullableMultipleDataStoreDataLatestStream) ReduceBool(fn func(bool, NullableMultipleDataStoreDataLatest, int) bool) []bool {
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
func (self *NullableMultipleDataStoreDataLatestStream) Reverse() *NullableMultipleDataStoreDataLatestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableMultipleDataStoreDataLatestStream) Replace(fn func(NullableMultipleDataStoreDataLatest, int) NullableMultipleDataStoreDataLatest) *NullableMultipleDataStoreDataLatestStream {
	return self.ForEach(func(v NullableMultipleDataStoreDataLatest, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableMultipleDataStoreDataLatestStream) Select(fn func(NullableMultipleDataStoreDataLatest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableMultipleDataStoreDataLatestStream) Set(index int, val NullableMultipleDataStoreDataLatest) *NullableMultipleDataStoreDataLatestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableMultipleDataStoreDataLatestStream) Skip(skip int) *NullableMultipleDataStoreDataLatestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableMultipleDataStoreDataLatestStream) SkippingEach(fn func(NullableMultipleDataStoreDataLatest, int) int) *NullableMultipleDataStoreDataLatestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableMultipleDataStoreDataLatestStream) Slice(startIndex, n int) *NullableMultipleDataStoreDataLatestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableMultipleDataStoreDataLatest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableMultipleDataStoreDataLatestStream) Sort(fn func(i, j int) bool) *NullableMultipleDataStoreDataLatestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableMultipleDataStoreDataLatestStream) Tail() *NullableMultipleDataStoreDataLatest {
	return self.Last()
}
func (self *NullableMultipleDataStoreDataLatestStream) TailOr(arg NullableMultipleDataStoreDataLatest) NullableMultipleDataStoreDataLatest {
	return self.LastOr(arg)
}
func (self *NullableMultipleDataStoreDataLatestStream) ToList() []NullableMultipleDataStoreDataLatest {
	return self.Val()
}
func (self *NullableMultipleDataStoreDataLatestStream) Unique() *NullableMultipleDataStoreDataLatestStream {
	return self.Distinct()
}
func (self *NullableMultipleDataStoreDataLatestStream) Val() []NullableMultipleDataStoreDataLatest {
	if self == nil {
		return []NullableMultipleDataStoreDataLatest{}
	}
	return *self.Copy()
}
func (self *NullableMultipleDataStoreDataLatestStream) While(fn func(NullableMultipleDataStoreDataLatest, int) bool) *NullableMultipleDataStoreDataLatestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableMultipleDataStoreDataLatestStream) Where(fn func(NullableMultipleDataStoreDataLatest) bool) *NullableMultipleDataStoreDataLatestStream {
	result := NullableMultipleDataStoreDataLatestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMultipleDataStoreDataLatestStream) WhereSlim(fn func(NullableMultipleDataStoreDataLatest) bool) *NullableMultipleDataStoreDataLatestStream {
	result := NullableMultipleDataStoreDataLatestStreamOf()
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
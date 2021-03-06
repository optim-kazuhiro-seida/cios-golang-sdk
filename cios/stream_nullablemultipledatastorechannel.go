/*
 * Collection utility of NullableMultipleDataStoreChannel Struct
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

type NullableMultipleDataStoreChannelStream []NullableMultipleDataStoreChannel

func NullableMultipleDataStoreChannelStreamOf(arg ...NullableMultipleDataStoreChannel) NullableMultipleDataStoreChannelStream {
	return arg
}
func NullableMultipleDataStoreChannelStreamFrom(arg []NullableMultipleDataStoreChannel) NullableMultipleDataStoreChannelStream {
	return arg
}
func CreateNullableMultipleDataStoreChannelStream(arg ...NullableMultipleDataStoreChannel) *NullableMultipleDataStoreChannelStream {
	tmp := NullableMultipleDataStoreChannelStreamOf(arg...)
	return &tmp
}
func GenerateNullableMultipleDataStoreChannelStream(arg []NullableMultipleDataStoreChannel) *NullableMultipleDataStoreChannelStream {
	tmp := NullableMultipleDataStoreChannelStreamFrom(arg)
	return &tmp
}

func (self *NullableMultipleDataStoreChannelStream) Add(arg NullableMultipleDataStoreChannel) *NullableMultipleDataStoreChannelStream {
	return self.AddAll(arg)
}
func (self *NullableMultipleDataStoreChannelStream) AddAll(arg ...NullableMultipleDataStoreChannel) *NullableMultipleDataStoreChannelStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableMultipleDataStoreChannelStream) AddSafe(arg *NullableMultipleDataStoreChannel) *NullableMultipleDataStoreChannelStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableMultipleDataStoreChannelStream) Aggregate(fn func(NullableMultipleDataStoreChannel, NullableMultipleDataStoreChannel) NullableMultipleDataStoreChannel) *NullableMultipleDataStoreChannelStream {
	result := NullableMultipleDataStoreChannelStreamOf()
	self.ForEach(func(v NullableMultipleDataStoreChannel, i int) {
		if i == 0 {
			result.Add(fn(NullableMultipleDataStoreChannel{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableMultipleDataStoreChannelStream) AllMatch(fn func(NullableMultipleDataStoreChannel, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableMultipleDataStoreChannelStream) AnyMatch(fn func(NullableMultipleDataStoreChannel, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableMultipleDataStoreChannelStream) Clone() *NullableMultipleDataStoreChannelStream {
	temp := make([]NullableMultipleDataStoreChannel, self.Len())
	copy(temp, *self)
	return (*NullableMultipleDataStoreChannelStream)(&temp)
}
func (self *NullableMultipleDataStoreChannelStream) Copy() *NullableMultipleDataStoreChannelStream {
	return self.Clone()
}
func (self *NullableMultipleDataStoreChannelStream) Concat(arg []NullableMultipleDataStoreChannel) *NullableMultipleDataStoreChannelStream {
	return self.AddAll(arg...)
}
func (self *NullableMultipleDataStoreChannelStream) Contains(arg NullableMultipleDataStoreChannel) bool {
	return self.FindIndex(func(_arg NullableMultipleDataStoreChannel, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableMultipleDataStoreChannelStream) Clean() *NullableMultipleDataStoreChannelStream {
	*self = NullableMultipleDataStoreChannelStreamOf()
	return self
}
func (self *NullableMultipleDataStoreChannelStream) Delete(index int) *NullableMultipleDataStoreChannelStream {
	return self.DeleteRange(index, index)
}
func (self *NullableMultipleDataStoreChannelStream) DeleteRange(startIndex, endIndex int) *NullableMultipleDataStoreChannelStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableMultipleDataStoreChannelStream) Distinct() *NullableMultipleDataStoreChannelStream {
	caches := map[string]bool{}
	result := NullableMultipleDataStoreChannelStreamOf()
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
func (self *NullableMultipleDataStoreChannelStream) Each(fn func(NullableMultipleDataStoreChannel)) *NullableMultipleDataStoreChannelStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableMultipleDataStoreChannelStream) EachRight(fn func(NullableMultipleDataStoreChannel)) *NullableMultipleDataStoreChannelStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableMultipleDataStoreChannelStream) Equals(arr []NullableMultipleDataStoreChannel) bool {
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
func (self *NullableMultipleDataStoreChannelStream) Filter(fn func(NullableMultipleDataStoreChannel, int) bool) *NullableMultipleDataStoreChannelStream {
	result := NullableMultipleDataStoreChannelStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMultipleDataStoreChannelStream) FilterSlim(fn func(NullableMultipleDataStoreChannel, int) bool) *NullableMultipleDataStoreChannelStream {
	result := NullableMultipleDataStoreChannelStreamOf()
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
func (self *NullableMultipleDataStoreChannelStream) Find(fn func(NullableMultipleDataStoreChannel, int) bool) *NullableMultipleDataStoreChannel {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableMultipleDataStoreChannelStream) FindOr(fn func(NullableMultipleDataStoreChannel, int) bool, or NullableMultipleDataStoreChannel) NullableMultipleDataStoreChannel {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableMultipleDataStoreChannelStream) FindIndex(fn func(NullableMultipleDataStoreChannel, int) bool) int {
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
func (self *NullableMultipleDataStoreChannelStream) First() *NullableMultipleDataStoreChannel {
	return self.Get(0)
}
func (self *NullableMultipleDataStoreChannelStream) FirstOr(arg NullableMultipleDataStoreChannel) NullableMultipleDataStoreChannel {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleDataStoreChannelStream) ForEach(fn func(NullableMultipleDataStoreChannel, int)) *NullableMultipleDataStoreChannelStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableMultipleDataStoreChannelStream) ForEachRight(fn func(NullableMultipleDataStoreChannel, int)) *NullableMultipleDataStoreChannelStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableMultipleDataStoreChannelStream) GroupBy(fn func(NullableMultipleDataStoreChannel, int) string) map[string][]NullableMultipleDataStoreChannel {
	m := map[string][]NullableMultipleDataStoreChannel{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableMultipleDataStoreChannelStream) GroupByValues(fn func(NullableMultipleDataStoreChannel, int) string) [][]NullableMultipleDataStoreChannel {
	var tmp [][]NullableMultipleDataStoreChannel
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableMultipleDataStoreChannelStream) IndexOf(arg NullableMultipleDataStoreChannel) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableMultipleDataStoreChannelStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableMultipleDataStoreChannelStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableMultipleDataStoreChannelStream) Last() *NullableMultipleDataStoreChannel {
	return self.Get(self.Len() - 1)
}
func (self *NullableMultipleDataStoreChannelStream) LastOr(arg NullableMultipleDataStoreChannel) NullableMultipleDataStoreChannel {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleDataStoreChannelStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableMultipleDataStoreChannelStream) Limit(limit int) *NullableMultipleDataStoreChannelStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableMultipleDataStoreChannelStream) Map(fn func(NullableMultipleDataStoreChannel, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDataStoreChannelStream) Map2Int(fn func(NullableMultipleDataStoreChannel, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDataStoreChannelStream) Map2Int32(fn func(NullableMultipleDataStoreChannel, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDataStoreChannelStream) Map2Int64(fn func(NullableMultipleDataStoreChannel, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDataStoreChannelStream) Map2Float32(fn func(NullableMultipleDataStoreChannel, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDataStoreChannelStream) Map2Float64(fn func(NullableMultipleDataStoreChannel, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDataStoreChannelStream) Map2Bool(fn func(NullableMultipleDataStoreChannel, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDataStoreChannelStream) Map2Bytes(fn func(NullableMultipleDataStoreChannel, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDataStoreChannelStream) Map2String(fn func(NullableMultipleDataStoreChannel, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDataStoreChannelStream) Max(fn func(NullableMultipleDataStoreChannel, int) float64) *NullableMultipleDataStoreChannel {
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
func (self *NullableMultipleDataStoreChannelStream) Min(fn func(NullableMultipleDataStoreChannel, int) float64) *NullableMultipleDataStoreChannel {
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
func (self *NullableMultipleDataStoreChannelStream) NoneMatch(fn func(NullableMultipleDataStoreChannel, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableMultipleDataStoreChannelStream) Get(index int) *NullableMultipleDataStoreChannel {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableMultipleDataStoreChannelStream) GetOr(index int, arg NullableMultipleDataStoreChannel) NullableMultipleDataStoreChannel {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleDataStoreChannelStream) Peek(fn func(*NullableMultipleDataStoreChannel, int)) *NullableMultipleDataStoreChannelStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableMultipleDataStoreChannelStream) Reduce(fn func(NullableMultipleDataStoreChannel, NullableMultipleDataStoreChannel, int) NullableMultipleDataStoreChannel) *NullableMultipleDataStoreChannelStream {
	return self.ReduceInit(fn, NullableMultipleDataStoreChannel{})
}
func (self *NullableMultipleDataStoreChannelStream) ReduceInit(fn func(NullableMultipleDataStoreChannel, NullableMultipleDataStoreChannel, int) NullableMultipleDataStoreChannel, initialValue NullableMultipleDataStoreChannel) *NullableMultipleDataStoreChannelStream {
	result := NullableMultipleDataStoreChannelStreamOf()
	self.ForEach(func(v NullableMultipleDataStoreChannel, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableMultipleDataStoreChannelStream) ReduceInterface(fn func(interface{}, NullableMultipleDataStoreChannel, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableMultipleDataStoreChannel{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleDataStoreChannelStream) ReduceString(fn func(string, NullableMultipleDataStoreChannel, int) string) []string {
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
func (self *NullableMultipleDataStoreChannelStream) ReduceInt(fn func(int, NullableMultipleDataStoreChannel, int) int) []int {
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
func (self *NullableMultipleDataStoreChannelStream) ReduceInt32(fn func(int32, NullableMultipleDataStoreChannel, int) int32) []int32 {
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
func (self *NullableMultipleDataStoreChannelStream) ReduceInt64(fn func(int64, NullableMultipleDataStoreChannel, int) int64) []int64 {
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
func (self *NullableMultipleDataStoreChannelStream) ReduceFloat32(fn func(float32, NullableMultipleDataStoreChannel, int) float32) []float32 {
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
func (self *NullableMultipleDataStoreChannelStream) ReduceFloat64(fn func(float64, NullableMultipleDataStoreChannel, int) float64) []float64 {
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
func (self *NullableMultipleDataStoreChannelStream) ReduceBool(fn func(bool, NullableMultipleDataStoreChannel, int) bool) []bool {
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
func (self *NullableMultipleDataStoreChannelStream) Reverse() *NullableMultipleDataStoreChannelStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableMultipleDataStoreChannelStream) Replace(fn func(NullableMultipleDataStoreChannel, int) NullableMultipleDataStoreChannel) *NullableMultipleDataStoreChannelStream {
	return self.ForEach(func(v NullableMultipleDataStoreChannel, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableMultipleDataStoreChannelStream) Select(fn func(NullableMultipleDataStoreChannel) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableMultipleDataStoreChannelStream) Set(index int, val NullableMultipleDataStoreChannel) *NullableMultipleDataStoreChannelStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableMultipleDataStoreChannelStream) Skip(skip int) *NullableMultipleDataStoreChannelStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableMultipleDataStoreChannelStream) SkippingEach(fn func(NullableMultipleDataStoreChannel, int) int) *NullableMultipleDataStoreChannelStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableMultipleDataStoreChannelStream) Slice(startIndex, n int) *NullableMultipleDataStoreChannelStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableMultipleDataStoreChannel{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableMultipleDataStoreChannelStream) Sort(fn func(i, j int) bool) *NullableMultipleDataStoreChannelStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableMultipleDataStoreChannelStream) Tail() *NullableMultipleDataStoreChannel {
	return self.Last()
}
func (self *NullableMultipleDataStoreChannelStream) TailOr(arg NullableMultipleDataStoreChannel) NullableMultipleDataStoreChannel {
	return self.LastOr(arg)
}
func (self *NullableMultipleDataStoreChannelStream) ToList() []NullableMultipleDataStoreChannel {
	return self.Val()
}
func (self *NullableMultipleDataStoreChannelStream) Unique() *NullableMultipleDataStoreChannelStream {
	return self.Distinct()
}
func (self *NullableMultipleDataStoreChannelStream) Val() []NullableMultipleDataStoreChannel {
	if self == nil {
		return []NullableMultipleDataStoreChannel{}
	}
	return *self.Copy()
}
func (self *NullableMultipleDataStoreChannelStream) While(fn func(NullableMultipleDataStoreChannel, int) bool) *NullableMultipleDataStoreChannelStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableMultipleDataStoreChannelStream) Where(fn func(NullableMultipleDataStoreChannel) bool) *NullableMultipleDataStoreChannelStream {
	result := NullableMultipleDataStoreChannelStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMultipleDataStoreChannelStream) WhereSlim(fn func(NullableMultipleDataStoreChannel) bool) *NullableMultipleDataStoreChannelStream {
	result := NullableMultipleDataStoreChannelStreamOf()
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

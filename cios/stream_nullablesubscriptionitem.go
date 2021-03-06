/*
 * Collection utility of NullableSubscriptionItem Struct
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

type NullableSubscriptionItemStream []NullableSubscriptionItem

func NullableSubscriptionItemStreamOf(arg ...NullableSubscriptionItem) NullableSubscriptionItemStream {
	return arg
}
func NullableSubscriptionItemStreamFrom(arg []NullableSubscriptionItem) NullableSubscriptionItemStream {
	return arg
}
func CreateNullableSubscriptionItemStream(arg ...NullableSubscriptionItem) *NullableSubscriptionItemStream {
	tmp := NullableSubscriptionItemStreamOf(arg...)
	return &tmp
}
func GenerateNullableSubscriptionItemStream(arg []NullableSubscriptionItem) *NullableSubscriptionItemStream {
	tmp := NullableSubscriptionItemStreamFrom(arg)
	return &tmp
}

func (self *NullableSubscriptionItemStream) Add(arg NullableSubscriptionItem) *NullableSubscriptionItemStream {
	return self.AddAll(arg)
}
func (self *NullableSubscriptionItemStream) AddAll(arg ...NullableSubscriptionItem) *NullableSubscriptionItemStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableSubscriptionItemStream) AddSafe(arg *NullableSubscriptionItem) *NullableSubscriptionItemStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableSubscriptionItemStream) Aggregate(fn func(NullableSubscriptionItem, NullableSubscriptionItem) NullableSubscriptionItem) *NullableSubscriptionItemStream {
	result := NullableSubscriptionItemStreamOf()
	self.ForEach(func(v NullableSubscriptionItem, i int) {
		if i == 0 {
			result.Add(fn(NullableSubscriptionItem{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableSubscriptionItemStream) AllMatch(fn func(NullableSubscriptionItem, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableSubscriptionItemStream) AnyMatch(fn func(NullableSubscriptionItem, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableSubscriptionItemStream) Clone() *NullableSubscriptionItemStream {
	temp := make([]NullableSubscriptionItem, self.Len())
	copy(temp, *self)
	return (*NullableSubscriptionItemStream)(&temp)
}
func (self *NullableSubscriptionItemStream) Copy() *NullableSubscriptionItemStream {
	return self.Clone()
}
func (self *NullableSubscriptionItemStream) Concat(arg []NullableSubscriptionItem) *NullableSubscriptionItemStream {
	return self.AddAll(arg...)
}
func (self *NullableSubscriptionItemStream) Contains(arg NullableSubscriptionItem) bool {
	return self.FindIndex(func(_arg NullableSubscriptionItem, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableSubscriptionItemStream) Clean() *NullableSubscriptionItemStream {
	*self = NullableSubscriptionItemStreamOf()
	return self
}
func (self *NullableSubscriptionItemStream) Delete(index int) *NullableSubscriptionItemStream {
	return self.DeleteRange(index, index)
}
func (self *NullableSubscriptionItemStream) DeleteRange(startIndex, endIndex int) *NullableSubscriptionItemStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableSubscriptionItemStream) Distinct() *NullableSubscriptionItemStream {
	caches := map[string]bool{}
	result := NullableSubscriptionItemStreamOf()
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
func (self *NullableSubscriptionItemStream) Each(fn func(NullableSubscriptionItem)) *NullableSubscriptionItemStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableSubscriptionItemStream) EachRight(fn func(NullableSubscriptionItem)) *NullableSubscriptionItemStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableSubscriptionItemStream) Equals(arr []NullableSubscriptionItem) bool {
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
func (self *NullableSubscriptionItemStream) Filter(fn func(NullableSubscriptionItem, int) bool) *NullableSubscriptionItemStream {
	result := NullableSubscriptionItemStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSubscriptionItemStream) FilterSlim(fn func(NullableSubscriptionItem, int) bool) *NullableSubscriptionItemStream {
	result := NullableSubscriptionItemStreamOf()
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
func (self *NullableSubscriptionItemStream) Find(fn func(NullableSubscriptionItem, int) bool) *NullableSubscriptionItem {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableSubscriptionItemStream) FindOr(fn func(NullableSubscriptionItem, int) bool, or NullableSubscriptionItem) NullableSubscriptionItem {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableSubscriptionItemStream) FindIndex(fn func(NullableSubscriptionItem, int) bool) int {
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
func (self *NullableSubscriptionItemStream) First() *NullableSubscriptionItem {
	return self.Get(0)
}
func (self *NullableSubscriptionItemStream) FirstOr(arg NullableSubscriptionItem) NullableSubscriptionItem {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSubscriptionItemStream) ForEach(fn func(NullableSubscriptionItem, int)) *NullableSubscriptionItemStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableSubscriptionItemStream) ForEachRight(fn func(NullableSubscriptionItem, int)) *NullableSubscriptionItemStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableSubscriptionItemStream) GroupBy(fn func(NullableSubscriptionItem, int) string) map[string][]NullableSubscriptionItem {
	m := map[string][]NullableSubscriptionItem{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableSubscriptionItemStream) GroupByValues(fn func(NullableSubscriptionItem, int) string) [][]NullableSubscriptionItem {
	var tmp [][]NullableSubscriptionItem
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableSubscriptionItemStream) IndexOf(arg NullableSubscriptionItem) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableSubscriptionItemStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableSubscriptionItemStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableSubscriptionItemStream) Last() *NullableSubscriptionItem {
	return self.Get(self.Len() - 1)
}
func (self *NullableSubscriptionItemStream) LastOr(arg NullableSubscriptionItem) NullableSubscriptionItem {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSubscriptionItemStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableSubscriptionItemStream) Limit(limit int) *NullableSubscriptionItemStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableSubscriptionItemStream) Map(fn func(NullableSubscriptionItem, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSubscriptionItemStream) Map2Int(fn func(NullableSubscriptionItem, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSubscriptionItemStream) Map2Int32(fn func(NullableSubscriptionItem, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSubscriptionItemStream) Map2Int64(fn func(NullableSubscriptionItem, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSubscriptionItemStream) Map2Float32(fn func(NullableSubscriptionItem, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSubscriptionItemStream) Map2Float64(fn func(NullableSubscriptionItem, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSubscriptionItemStream) Map2Bool(fn func(NullableSubscriptionItem, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSubscriptionItemStream) Map2Bytes(fn func(NullableSubscriptionItem, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSubscriptionItemStream) Map2String(fn func(NullableSubscriptionItem, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSubscriptionItemStream) Max(fn func(NullableSubscriptionItem, int) float64) *NullableSubscriptionItem {
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
func (self *NullableSubscriptionItemStream) Min(fn func(NullableSubscriptionItem, int) float64) *NullableSubscriptionItem {
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
func (self *NullableSubscriptionItemStream) NoneMatch(fn func(NullableSubscriptionItem, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableSubscriptionItemStream) Get(index int) *NullableSubscriptionItem {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableSubscriptionItemStream) GetOr(index int, arg NullableSubscriptionItem) NullableSubscriptionItem {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSubscriptionItemStream) Peek(fn func(*NullableSubscriptionItem, int)) *NullableSubscriptionItemStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableSubscriptionItemStream) Reduce(fn func(NullableSubscriptionItem, NullableSubscriptionItem, int) NullableSubscriptionItem) *NullableSubscriptionItemStream {
	return self.ReduceInit(fn, NullableSubscriptionItem{})
}
func (self *NullableSubscriptionItemStream) ReduceInit(fn func(NullableSubscriptionItem, NullableSubscriptionItem, int) NullableSubscriptionItem, initialValue NullableSubscriptionItem) *NullableSubscriptionItemStream {
	result := NullableSubscriptionItemStreamOf()
	self.ForEach(func(v NullableSubscriptionItem, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableSubscriptionItemStream) ReduceInterface(fn func(interface{}, NullableSubscriptionItem, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableSubscriptionItem{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableSubscriptionItemStream) ReduceString(fn func(string, NullableSubscriptionItem, int) string) []string {
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
func (self *NullableSubscriptionItemStream) ReduceInt(fn func(int, NullableSubscriptionItem, int) int) []int {
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
func (self *NullableSubscriptionItemStream) ReduceInt32(fn func(int32, NullableSubscriptionItem, int) int32) []int32 {
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
func (self *NullableSubscriptionItemStream) ReduceInt64(fn func(int64, NullableSubscriptionItem, int) int64) []int64 {
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
func (self *NullableSubscriptionItemStream) ReduceFloat32(fn func(float32, NullableSubscriptionItem, int) float32) []float32 {
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
func (self *NullableSubscriptionItemStream) ReduceFloat64(fn func(float64, NullableSubscriptionItem, int) float64) []float64 {
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
func (self *NullableSubscriptionItemStream) ReduceBool(fn func(bool, NullableSubscriptionItem, int) bool) []bool {
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
func (self *NullableSubscriptionItemStream) Reverse() *NullableSubscriptionItemStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableSubscriptionItemStream) Replace(fn func(NullableSubscriptionItem, int) NullableSubscriptionItem) *NullableSubscriptionItemStream {
	return self.ForEach(func(v NullableSubscriptionItem, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableSubscriptionItemStream) Select(fn func(NullableSubscriptionItem) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableSubscriptionItemStream) Set(index int, val NullableSubscriptionItem) *NullableSubscriptionItemStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableSubscriptionItemStream) Skip(skip int) *NullableSubscriptionItemStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableSubscriptionItemStream) SkippingEach(fn func(NullableSubscriptionItem, int) int) *NullableSubscriptionItemStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableSubscriptionItemStream) Slice(startIndex, n int) *NullableSubscriptionItemStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableSubscriptionItem{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableSubscriptionItemStream) Sort(fn func(i, j int) bool) *NullableSubscriptionItemStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableSubscriptionItemStream) Tail() *NullableSubscriptionItem {
	return self.Last()
}
func (self *NullableSubscriptionItemStream) TailOr(arg NullableSubscriptionItem) NullableSubscriptionItem {
	return self.LastOr(arg)
}
func (self *NullableSubscriptionItemStream) ToList() []NullableSubscriptionItem {
	return self.Val()
}
func (self *NullableSubscriptionItemStream) Unique() *NullableSubscriptionItemStream {
	return self.Distinct()
}
func (self *NullableSubscriptionItemStream) Val() []NullableSubscriptionItem {
	if self == nil {
		return []NullableSubscriptionItem{}
	}
	return *self.Copy()
}
func (self *NullableSubscriptionItemStream) While(fn func(NullableSubscriptionItem, int) bool) *NullableSubscriptionItemStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableSubscriptionItemStream) Where(fn func(NullableSubscriptionItem) bool) *NullableSubscriptionItemStream {
	result := NullableSubscriptionItemStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSubscriptionItemStream) WhereSlim(fn func(NullableSubscriptionItem) bool) *NullableSubscriptionItemStream {
	result := NullableSubscriptionItemStreamOf()
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

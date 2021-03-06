/*
 * Collection utility of NullablePackerFormatJson Struct
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

type NullablePackerFormatJsonStream []NullablePackerFormatJson

func NullablePackerFormatJsonStreamOf(arg ...NullablePackerFormatJson) NullablePackerFormatJsonStream {
	return arg
}
func NullablePackerFormatJsonStreamFrom(arg []NullablePackerFormatJson) NullablePackerFormatJsonStream {
	return arg
}
func CreateNullablePackerFormatJsonStream(arg ...NullablePackerFormatJson) *NullablePackerFormatJsonStream {
	tmp := NullablePackerFormatJsonStreamOf(arg...)
	return &tmp
}
func GenerateNullablePackerFormatJsonStream(arg []NullablePackerFormatJson) *NullablePackerFormatJsonStream {
	tmp := NullablePackerFormatJsonStreamFrom(arg)
	return &tmp
}

func (self *NullablePackerFormatJsonStream) Add(arg NullablePackerFormatJson) *NullablePackerFormatJsonStream {
	return self.AddAll(arg)
}
func (self *NullablePackerFormatJsonStream) AddAll(arg ...NullablePackerFormatJson) *NullablePackerFormatJsonStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullablePackerFormatJsonStream) AddSafe(arg *NullablePackerFormatJson) *NullablePackerFormatJsonStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullablePackerFormatJsonStream) Aggregate(fn func(NullablePackerFormatJson, NullablePackerFormatJson) NullablePackerFormatJson) *NullablePackerFormatJsonStream {
	result := NullablePackerFormatJsonStreamOf()
	self.ForEach(func(v NullablePackerFormatJson, i int) {
		if i == 0 {
			result.Add(fn(NullablePackerFormatJson{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullablePackerFormatJsonStream) AllMatch(fn func(NullablePackerFormatJson, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullablePackerFormatJsonStream) AnyMatch(fn func(NullablePackerFormatJson, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullablePackerFormatJsonStream) Clone() *NullablePackerFormatJsonStream {
	temp := make([]NullablePackerFormatJson, self.Len())
	copy(temp, *self)
	return (*NullablePackerFormatJsonStream)(&temp)
}
func (self *NullablePackerFormatJsonStream) Copy() *NullablePackerFormatJsonStream {
	return self.Clone()
}
func (self *NullablePackerFormatJsonStream) Concat(arg []NullablePackerFormatJson) *NullablePackerFormatJsonStream {
	return self.AddAll(arg...)
}
func (self *NullablePackerFormatJsonStream) Contains(arg NullablePackerFormatJson) bool {
	return self.FindIndex(func(_arg NullablePackerFormatJson, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullablePackerFormatJsonStream) Clean() *NullablePackerFormatJsonStream {
	*self = NullablePackerFormatJsonStreamOf()
	return self
}
func (self *NullablePackerFormatJsonStream) Delete(index int) *NullablePackerFormatJsonStream {
	return self.DeleteRange(index, index)
}
func (self *NullablePackerFormatJsonStream) DeleteRange(startIndex, endIndex int) *NullablePackerFormatJsonStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullablePackerFormatJsonStream) Distinct() *NullablePackerFormatJsonStream {
	caches := map[string]bool{}
	result := NullablePackerFormatJsonStreamOf()
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
func (self *NullablePackerFormatJsonStream) Each(fn func(NullablePackerFormatJson)) *NullablePackerFormatJsonStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullablePackerFormatJsonStream) EachRight(fn func(NullablePackerFormatJson)) *NullablePackerFormatJsonStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullablePackerFormatJsonStream) Equals(arr []NullablePackerFormatJson) bool {
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
func (self *NullablePackerFormatJsonStream) Filter(fn func(NullablePackerFormatJson, int) bool) *NullablePackerFormatJsonStream {
	result := NullablePackerFormatJsonStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullablePackerFormatJsonStream) FilterSlim(fn func(NullablePackerFormatJson, int) bool) *NullablePackerFormatJsonStream {
	result := NullablePackerFormatJsonStreamOf()
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
func (self *NullablePackerFormatJsonStream) Find(fn func(NullablePackerFormatJson, int) bool) *NullablePackerFormatJson {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullablePackerFormatJsonStream) FindOr(fn func(NullablePackerFormatJson, int) bool, or NullablePackerFormatJson) NullablePackerFormatJson {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullablePackerFormatJsonStream) FindIndex(fn func(NullablePackerFormatJson, int) bool) int {
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
func (self *NullablePackerFormatJsonStream) First() *NullablePackerFormatJson {
	return self.Get(0)
}
func (self *NullablePackerFormatJsonStream) FirstOr(arg NullablePackerFormatJson) NullablePackerFormatJson {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullablePackerFormatJsonStream) ForEach(fn func(NullablePackerFormatJson, int)) *NullablePackerFormatJsonStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullablePackerFormatJsonStream) ForEachRight(fn func(NullablePackerFormatJson, int)) *NullablePackerFormatJsonStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullablePackerFormatJsonStream) GroupBy(fn func(NullablePackerFormatJson, int) string) map[string][]NullablePackerFormatJson {
	m := map[string][]NullablePackerFormatJson{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullablePackerFormatJsonStream) GroupByValues(fn func(NullablePackerFormatJson, int) string) [][]NullablePackerFormatJson {
	var tmp [][]NullablePackerFormatJson
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullablePackerFormatJsonStream) IndexOf(arg NullablePackerFormatJson) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullablePackerFormatJsonStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullablePackerFormatJsonStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullablePackerFormatJsonStream) Last() *NullablePackerFormatJson {
	return self.Get(self.Len() - 1)
}
func (self *NullablePackerFormatJsonStream) LastOr(arg NullablePackerFormatJson) NullablePackerFormatJson {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullablePackerFormatJsonStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullablePackerFormatJsonStream) Limit(limit int) *NullablePackerFormatJsonStream {
	self.Slice(0, limit)
	return self
}

func (self *NullablePackerFormatJsonStream) Map(fn func(NullablePackerFormatJson, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullablePackerFormatJsonStream) Map2Int(fn func(NullablePackerFormatJson, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullablePackerFormatJsonStream) Map2Int32(fn func(NullablePackerFormatJson, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullablePackerFormatJsonStream) Map2Int64(fn func(NullablePackerFormatJson, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullablePackerFormatJsonStream) Map2Float32(fn func(NullablePackerFormatJson, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullablePackerFormatJsonStream) Map2Float64(fn func(NullablePackerFormatJson, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullablePackerFormatJsonStream) Map2Bool(fn func(NullablePackerFormatJson, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullablePackerFormatJsonStream) Map2Bytes(fn func(NullablePackerFormatJson, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullablePackerFormatJsonStream) Map2String(fn func(NullablePackerFormatJson, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullablePackerFormatJsonStream) Max(fn func(NullablePackerFormatJson, int) float64) *NullablePackerFormatJson {
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
func (self *NullablePackerFormatJsonStream) Min(fn func(NullablePackerFormatJson, int) float64) *NullablePackerFormatJson {
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
func (self *NullablePackerFormatJsonStream) NoneMatch(fn func(NullablePackerFormatJson, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullablePackerFormatJsonStream) Get(index int) *NullablePackerFormatJson {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullablePackerFormatJsonStream) GetOr(index int, arg NullablePackerFormatJson) NullablePackerFormatJson {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullablePackerFormatJsonStream) Peek(fn func(*NullablePackerFormatJson, int)) *NullablePackerFormatJsonStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullablePackerFormatJsonStream) Reduce(fn func(NullablePackerFormatJson, NullablePackerFormatJson, int) NullablePackerFormatJson) *NullablePackerFormatJsonStream {
	return self.ReduceInit(fn, NullablePackerFormatJson{})
}
func (self *NullablePackerFormatJsonStream) ReduceInit(fn func(NullablePackerFormatJson, NullablePackerFormatJson, int) NullablePackerFormatJson, initialValue NullablePackerFormatJson) *NullablePackerFormatJsonStream {
	result := NullablePackerFormatJsonStreamOf()
	self.ForEach(func(v NullablePackerFormatJson, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullablePackerFormatJsonStream) ReduceInterface(fn func(interface{}, NullablePackerFormatJson, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullablePackerFormatJson{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullablePackerFormatJsonStream) ReduceString(fn func(string, NullablePackerFormatJson, int) string) []string {
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
func (self *NullablePackerFormatJsonStream) ReduceInt(fn func(int, NullablePackerFormatJson, int) int) []int {
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
func (self *NullablePackerFormatJsonStream) ReduceInt32(fn func(int32, NullablePackerFormatJson, int) int32) []int32 {
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
func (self *NullablePackerFormatJsonStream) ReduceInt64(fn func(int64, NullablePackerFormatJson, int) int64) []int64 {
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
func (self *NullablePackerFormatJsonStream) ReduceFloat32(fn func(float32, NullablePackerFormatJson, int) float32) []float32 {
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
func (self *NullablePackerFormatJsonStream) ReduceFloat64(fn func(float64, NullablePackerFormatJson, int) float64) []float64 {
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
func (self *NullablePackerFormatJsonStream) ReduceBool(fn func(bool, NullablePackerFormatJson, int) bool) []bool {
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
func (self *NullablePackerFormatJsonStream) Reverse() *NullablePackerFormatJsonStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullablePackerFormatJsonStream) Replace(fn func(NullablePackerFormatJson, int) NullablePackerFormatJson) *NullablePackerFormatJsonStream {
	return self.ForEach(func(v NullablePackerFormatJson, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullablePackerFormatJsonStream) Select(fn func(NullablePackerFormatJson) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullablePackerFormatJsonStream) Set(index int, val NullablePackerFormatJson) *NullablePackerFormatJsonStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullablePackerFormatJsonStream) Skip(skip int) *NullablePackerFormatJsonStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullablePackerFormatJsonStream) SkippingEach(fn func(NullablePackerFormatJson, int) int) *NullablePackerFormatJsonStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullablePackerFormatJsonStream) Slice(startIndex, n int) *NullablePackerFormatJsonStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullablePackerFormatJson{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullablePackerFormatJsonStream) Sort(fn func(i, j int) bool) *NullablePackerFormatJsonStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullablePackerFormatJsonStream) Tail() *NullablePackerFormatJson {
	return self.Last()
}
func (self *NullablePackerFormatJsonStream) TailOr(arg NullablePackerFormatJson) NullablePackerFormatJson {
	return self.LastOr(arg)
}
func (self *NullablePackerFormatJsonStream) ToList() []NullablePackerFormatJson {
	return self.Val()
}
func (self *NullablePackerFormatJsonStream) Unique() *NullablePackerFormatJsonStream {
	return self.Distinct()
}
func (self *NullablePackerFormatJsonStream) Val() []NullablePackerFormatJson {
	if self == nil {
		return []NullablePackerFormatJson{}
	}
	return *self.Copy()
}
func (self *NullablePackerFormatJsonStream) While(fn func(NullablePackerFormatJson, int) bool) *NullablePackerFormatJsonStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullablePackerFormatJsonStream) Where(fn func(NullablePackerFormatJson) bool) *NullablePackerFormatJsonStream {
	result := NullablePackerFormatJsonStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullablePackerFormatJsonStream) WhereSlim(fn func(NullablePackerFormatJson) bool) *NullablePackerFormatJsonStream {
	result := NullablePackerFormatJsonStreamOf()
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

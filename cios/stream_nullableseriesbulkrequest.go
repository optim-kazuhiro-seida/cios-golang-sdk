/*
 * Collection utility of NullableSeriesBulkRequest Struct
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

type NullableSeriesBulkRequestStream []NullableSeriesBulkRequest

func NullableSeriesBulkRequestStreamOf(arg ...NullableSeriesBulkRequest) NullableSeriesBulkRequestStream {
	return arg
}
func NullableSeriesBulkRequestStreamFrom(arg []NullableSeriesBulkRequest) NullableSeriesBulkRequestStream {
	return arg
}
func CreateNullableSeriesBulkRequestStream(arg ...NullableSeriesBulkRequest) *NullableSeriesBulkRequestStream {
	tmp := NullableSeriesBulkRequestStreamOf(arg...)
	return &tmp
}
func GenerateNullableSeriesBulkRequestStream(arg []NullableSeriesBulkRequest) *NullableSeriesBulkRequestStream {
	tmp := NullableSeriesBulkRequestStreamFrom(arg)
	return &tmp
}

func (self *NullableSeriesBulkRequestStream) Add(arg NullableSeriesBulkRequest) *NullableSeriesBulkRequestStream {
	return self.AddAll(arg)
}
func (self *NullableSeriesBulkRequestStream) AddAll(arg ...NullableSeriesBulkRequest) *NullableSeriesBulkRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableSeriesBulkRequestStream) AddSafe(arg *NullableSeriesBulkRequest) *NullableSeriesBulkRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableSeriesBulkRequestStream) Aggregate(fn func(NullableSeriesBulkRequest, NullableSeriesBulkRequest) NullableSeriesBulkRequest) *NullableSeriesBulkRequestStream {
	result := NullableSeriesBulkRequestStreamOf()
	self.ForEach(func(v NullableSeriesBulkRequest, i int) {
		if i == 0 {
			result.Add(fn(NullableSeriesBulkRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableSeriesBulkRequestStream) AllMatch(fn func(NullableSeriesBulkRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableSeriesBulkRequestStream) AnyMatch(fn func(NullableSeriesBulkRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableSeriesBulkRequestStream) Clone() *NullableSeriesBulkRequestStream {
	temp := make([]NullableSeriesBulkRequest, self.Len())
	copy(temp, *self)
	return (*NullableSeriesBulkRequestStream)(&temp)
}
func (self *NullableSeriesBulkRequestStream) Copy() *NullableSeriesBulkRequestStream {
	return self.Clone()
}
func (self *NullableSeriesBulkRequestStream) Concat(arg []NullableSeriesBulkRequest) *NullableSeriesBulkRequestStream {
	return self.AddAll(arg...)
}
func (self *NullableSeriesBulkRequestStream) Contains(arg NullableSeriesBulkRequest) bool {
	return self.FindIndex(func(_arg NullableSeriesBulkRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableSeriesBulkRequestStream) Clean() *NullableSeriesBulkRequestStream {
	*self = NullableSeriesBulkRequestStreamOf()
	return self
}
func (self *NullableSeriesBulkRequestStream) Delete(index int) *NullableSeriesBulkRequestStream {
	return self.DeleteRange(index, index)
}
func (self *NullableSeriesBulkRequestStream) DeleteRange(startIndex, endIndex int) *NullableSeriesBulkRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableSeriesBulkRequestStream) Distinct() *NullableSeriesBulkRequestStream {
	caches := map[string]bool{}
	result := NullableSeriesBulkRequestStreamOf()
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
func (self *NullableSeriesBulkRequestStream) Each(fn func(NullableSeriesBulkRequest)) *NullableSeriesBulkRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableSeriesBulkRequestStream) EachRight(fn func(NullableSeriesBulkRequest)) *NullableSeriesBulkRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableSeriesBulkRequestStream) Equals(arr []NullableSeriesBulkRequest) bool {
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
func (self *NullableSeriesBulkRequestStream) Filter(fn func(NullableSeriesBulkRequest, int) bool) *NullableSeriesBulkRequestStream {
	result := NullableSeriesBulkRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSeriesBulkRequestStream) FilterSlim(fn func(NullableSeriesBulkRequest, int) bool) *NullableSeriesBulkRequestStream {
	result := NullableSeriesBulkRequestStreamOf()
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
func (self *NullableSeriesBulkRequestStream) Find(fn func(NullableSeriesBulkRequest, int) bool) *NullableSeriesBulkRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableSeriesBulkRequestStream) FindOr(fn func(NullableSeriesBulkRequest, int) bool, or NullableSeriesBulkRequest) NullableSeriesBulkRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableSeriesBulkRequestStream) FindIndex(fn func(NullableSeriesBulkRequest, int) bool) int {
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
func (self *NullableSeriesBulkRequestStream) First() *NullableSeriesBulkRequest {
	return self.Get(0)
}
func (self *NullableSeriesBulkRequestStream) FirstOr(arg NullableSeriesBulkRequest) NullableSeriesBulkRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSeriesBulkRequestStream) ForEach(fn func(NullableSeriesBulkRequest, int)) *NullableSeriesBulkRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableSeriesBulkRequestStream) ForEachRight(fn func(NullableSeriesBulkRequest, int)) *NullableSeriesBulkRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableSeriesBulkRequestStream) GroupBy(fn func(NullableSeriesBulkRequest, int) string) map[string][]NullableSeriesBulkRequest {
	m := map[string][]NullableSeriesBulkRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableSeriesBulkRequestStream) GroupByValues(fn func(NullableSeriesBulkRequest, int) string) [][]NullableSeriesBulkRequest {
	var tmp [][]NullableSeriesBulkRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableSeriesBulkRequestStream) IndexOf(arg NullableSeriesBulkRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableSeriesBulkRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableSeriesBulkRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableSeriesBulkRequestStream) Last() *NullableSeriesBulkRequest {
	return self.Get(self.Len() - 1)
}
func (self *NullableSeriesBulkRequestStream) LastOr(arg NullableSeriesBulkRequest) NullableSeriesBulkRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSeriesBulkRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableSeriesBulkRequestStream) Limit(limit int) *NullableSeriesBulkRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableSeriesBulkRequestStream) Map(fn func(NullableSeriesBulkRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesBulkRequestStream) Map2Int(fn func(NullableSeriesBulkRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesBulkRequestStream) Map2Int32(fn func(NullableSeriesBulkRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesBulkRequestStream) Map2Int64(fn func(NullableSeriesBulkRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesBulkRequestStream) Map2Float32(fn func(NullableSeriesBulkRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesBulkRequestStream) Map2Float64(fn func(NullableSeriesBulkRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesBulkRequestStream) Map2Bool(fn func(NullableSeriesBulkRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesBulkRequestStream) Map2Bytes(fn func(NullableSeriesBulkRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesBulkRequestStream) Map2String(fn func(NullableSeriesBulkRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesBulkRequestStream) Max(fn func(NullableSeriesBulkRequest, int) float64) *NullableSeriesBulkRequest {
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
func (self *NullableSeriesBulkRequestStream) Min(fn func(NullableSeriesBulkRequest, int) float64) *NullableSeriesBulkRequest {
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
func (self *NullableSeriesBulkRequestStream) NoneMatch(fn func(NullableSeriesBulkRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableSeriesBulkRequestStream) Get(index int) *NullableSeriesBulkRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableSeriesBulkRequestStream) GetOr(index int, arg NullableSeriesBulkRequest) NullableSeriesBulkRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSeriesBulkRequestStream) Peek(fn func(*NullableSeriesBulkRequest, int)) *NullableSeriesBulkRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *NullableSeriesBulkRequestStream) Reduce(fn func(NullableSeriesBulkRequest, NullableSeriesBulkRequest, int) NullableSeriesBulkRequest) *NullableSeriesBulkRequestStream {
	return self.ReduceInit(fn, NullableSeriesBulkRequest{})
}
func (self *NullableSeriesBulkRequestStream) ReduceInit(fn func(NullableSeriesBulkRequest, NullableSeriesBulkRequest, int) NullableSeriesBulkRequest, initialValue NullableSeriesBulkRequest) *NullableSeriesBulkRequestStream {
	result := NullableSeriesBulkRequestStreamOf()
	self.ForEach(func(v NullableSeriesBulkRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableSeriesBulkRequestStream) ReduceInterface(fn func(interface{}, NullableSeriesBulkRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableSeriesBulkRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableSeriesBulkRequestStream) ReduceString(fn func(string, NullableSeriesBulkRequest, int) string) []string {
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
func (self *NullableSeriesBulkRequestStream) ReduceInt(fn func(int, NullableSeriesBulkRequest, int) int) []int {
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
func (self *NullableSeriesBulkRequestStream) ReduceInt32(fn func(int32, NullableSeriesBulkRequest, int) int32) []int32 {
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
func (self *NullableSeriesBulkRequestStream) ReduceInt64(fn func(int64, NullableSeriesBulkRequest, int) int64) []int64 {
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
func (self *NullableSeriesBulkRequestStream) ReduceFloat32(fn func(float32, NullableSeriesBulkRequest, int) float32) []float32 {
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
func (self *NullableSeriesBulkRequestStream) ReduceFloat64(fn func(float64, NullableSeriesBulkRequest, int) float64) []float64 {
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
func (self *NullableSeriesBulkRequestStream) ReduceBool(fn func(bool, NullableSeriesBulkRequest, int) bool) []bool {
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
func (self *NullableSeriesBulkRequestStream) Reverse() *NullableSeriesBulkRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableSeriesBulkRequestStream) Replace(fn func(NullableSeriesBulkRequest, int) NullableSeriesBulkRequest) *NullableSeriesBulkRequestStream {
	return self.ForEach(func(v NullableSeriesBulkRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableSeriesBulkRequestStream) Select(fn func(NullableSeriesBulkRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableSeriesBulkRequestStream) Set(index int, val NullableSeriesBulkRequest) *NullableSeriesBulkRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableSeriesBulkRequestStream) Skip(skip int) *NullableSeriesBulkRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableSeriesBulkRequestStream) SkippingEach(fn func(NullableSeriesBulkRequest, int) int) *NullableSeriesBulkRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableSeriesBulkRequestStream) Slice(startIndex, n int) *NullableSeriesBulkRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableSeriesBulkRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableSeriesBulkRequestStream) Sort(fn func(i, j int) bool) *NullableSeriesBulkRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableSeriesBulkRequestStream) Tail() *NullableSeriesBulkRequest {
	return self.Last()
}
func (self *NullableSeriesBulkRequestStream) TailOr(arg NullableSeriesBulkRequest) NullableSeriesBulkRequest {
	return self.LastOr(arg)
}
func (self *NullableSeriesBulkRequestStream) ToList() []NullableSeriesBulkRequest {
	return self.Val()
}
func (self *NullableSeriesBulkRequestStream) Unique() *NullableSeriesBulkRequestStream {
	return self.Distinct()
}
func (self *NullableSeriesBulkRequestStream) Val() []NullableSeriesBulkRequest {
	if self == nil {
		return []NullableSeriesBulkRequest{}
	}
	return *self.Copy()
}
func (self *NullableSeriesBulkRequestStream) While(fn func(NullableSeriesBulkRequest, int) bool) *NullableSeriesBulkRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableSeriesBulkRequestStream) Where(fn func(NullableSeriesBulkRequest) bool) *NullableSeriesBulkRequestStream {
	result := NullableSeriesBulkRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSeriesBulkRequestStream) WhereSlim(fn func(NullableSeriesBulkRequest) bool) *NullableSeriesBulkRequestStream {
	result := NullableSeriesBulkRequestStreamOf()
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
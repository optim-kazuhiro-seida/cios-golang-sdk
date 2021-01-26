/*
 * Collection utility of ApiDeleteCircleRequest Struct
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

type ApiDeleteCircleRequestStream []ApiDeleteCircleRequest

func ApiDeleteCircleRequestStreamOf(arg ...ApiDeleteCircleRequest) ApiDeleteCircleRequestStream {
	return arg
}
func ApiDeleteCircleRequestStreamFrom(arg []ApiDeleteCircleRequest) ApiDeleteCircleRequestStream {
	return arg
}
func CreateApiDeleteCircleRequestStream(arg ...ApiDeleteCircleRequest) *ApiDeleteCircleRequestStream {
	tmp := ApiDeleteCircleRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiDeleteCircleRequestStream(arg []ApiDeleteCircleRequest) *ApiDeleteCircleRequestStream {
	tmp := ApiDeleteCircleRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiDeleteCircleRequestStream) Add(arg ApiDeleteCircleRequest) *ApiDeleteCircleRequestStream {
	return self.AddAll(arg)
}
func (self *ApiDeleteCircleRequestStream) AddAll(arg ...ApiDeleteCircleRequest) *ApiDeleteCircleRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiDeleteCircleRequestStream) AddSafe(arg *ApiDeleteCircleRequest) *ApiDeleteCircleRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiDeleteCircleRequestStream) Aggregate(fn func(ApiDeleteCircleRequest, ApiDeleteCircleRequest) ApiDeleteCircleRequest) *ApiDeleteCircleRequestStream {
	result := ApiDeleteCircleRequestStreamOf()
	self.ForEach(func(v ApiDeleteCircleRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiDeleteCircleRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiDeleteCircleRequestStream) AllMatch(fn func(ApiDeleteCircleRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiDeleteCircleRequestStream) AnyMatch(fn func(ApiDeleteCircleRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiDeleteCircleRequestStream) Clone() *ApiDeleteCircleRequestStream {
	temp := make([]ApiDeleteCircleRequest, self.Len())
	copy(temp, *self)
	return (*ApiDeleteCircleRequestStream)(&temp)
}
func (self *ApiDeleteCircleRequestStream) Copy() *ApiDeleteCircleRequestStream {
	return self.Clone()
}
func (self *ApiDeleteCircleRequestStream) Concat(arg []ApiDeleteCircleRequest) *ApiDeleteCircleRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiDeleteCircleRequestStream) Contains(arg ApiDeleteCircleRequest) bool {
	return self.FindIndex(func(_arg ApiDeleteCircleRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiDeleteCircleRequestStream) Clean() *ApiDeleteCircleRequestStream {
	*self = ApiDeleteCircleRequestStreamOf()
	return self
}
func (self *ApiDeleteCircleRequestStream) Delete(index int) *ApiDeleteCircleRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiDeleteCircleRequestStream) DeleteRange(startIndex, endIndex int) *ApiDeleteCircleRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiDeleteCircleRequestStream) Distinct() *ApiDeleteCircleRequestStream {
	caches := map[string]bool{}
	result := ApiDeleteCircleRequestStreamOf()
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
func (self *ApiDeleteCircleRequestStream) Each(fn func(ApiDeleteCircleRequest)) *ApiDeleteCircleRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiDeleteCircleRequestStream) EachRight(fn func(ApiDeleteCircleRequest)) *ApiDeleteCircleRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiDeleteCircleRequestStream) Equals(arr []ApiDeleteCircleRequest) bool {
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
func (self *ApiDeleteCircleRequestStream) Filter(fn func(ApiDeleteCircleRequest, int) bool) *ApiDeleteCircleRequestStream {
	result := ApiDeleteCircleRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiDeleteCircleRequestStream) FilterSlim(fn func(ApiDeleteCircleRequest, int) bool) *ApiDeleteCircleRequestStream {
	result := ApiDeleteCircleRequestStreamOf()
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
func (self *ApiDeleteCircleRequestStream) Find(fn func(ApiDeleteCircleRequest, int) bool) *ApiDeleteCircleRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiDeleteCircleRequestStream) FindOr(fn func(ApiDeleteCircleRequest, int) bool, or ApiDeleteCircleRequest) ApiDeleteCircleRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiDeleteCircleRequestStream) FindIndex(fn func(ApiDeleteCircleRequest, int) bool) int {
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
func (self *ApiDeleteCircleRequestStream) First() *ApiDeleteCircleRequest {
	return self.Get(0)
}
func (self *ApiDeleteCircleRequestStream) FirstOr(arg ApiDeleteCircleRequest) ApiDeleteCircleRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiDeleteCircleRequestStream) ForEach(fn func(ApiDeleteCircleRequest, int)) *ApiDeleteCircleRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiDeleteCircleRequestStream) ForEachRight(fn func(ApiDeleteCircleRequest, int)) *ApiDeleteCircleRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiDeleteCircleRequestStream) GroupBy(fn func(ApiDeleteCircleRequest, int) string) map[string][]ApiDeleteCircleRequest {
	m := map[string][]ApiDeleteCircleRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiDeleteCircleRequestStream) GroupByValues(fn func(ApiDeleteCircleRequest, int) string) [][]ApiDeleteCircleRequest {
	var tmp [][]ApiDeleteCircleRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiDeleteCircleRequestStream) IndexOf(arg ApiDeleteCircleRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiDeleteCircleRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiDeleteCircleRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiDeleteCircleRequestStream) Last() *ApiDeleteCircleRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiDeleteCircleRequestStream) LastOr(arg ApiDeleteCircleRequest) ApiDeleteCircleRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiDeleteCircleRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiDeleteCircleRequestStream) Limit(limit int) *ApiDeleteCircleRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiDeleteCircleRequestStream) Map(fn func(ApiDeleteCircleRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteCircleRequestStream) Map2Int(fn func(ApiDeleteCircleRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteCircleRequestStream) Map2Int32(fn func(ApiDeleteCircleRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteCircleRequestStream) Map2Int64(fn func(ApiDeleteCircleRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteCircleRequestStream) Map2Float32(fn func(ApiDeleteCircleRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteCircleRequestStream) Map2Float64(fn func(ApiDeleteCircleRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteCircleRequestStream) Map2Bool(fn func(ApiDeleteCircleRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteCircleRequestStream) Map2Bytes(fn func(ApiDeleteCircleRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteCircleRequestStream) Map2String(fn func(ApiDeleteCircleRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteCircleRequestStream) Max(fn func(ApiDeleteCircleRequest, int) float64) *ApiDeleteCircleRequest {
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
func (self *ApiDeleteCircleRequestStream) Min(fn func(ApiDeleteCircleRequest, int) float64) *ApiDeleteCircleRequest {
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
func (self *ApiDeleteCircleRequestStream) NoneMatch(fn func(ApiDeleteCircleRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiDeleteCircleRequestStream) Get(index int) *ApiDeleteCircleRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiDeleteCircleRequestStream) GetOr(index int, arg ApiDeleteCircleRequest) ApiDeleteCircleRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiDeleteCircleRequestStream) Peek(fn func(*ApiDeleteCircleRequest, int)) *ApiDeleteCircleRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *ApiDeleteCircleRequestStream) Reduce(fn func(ApiDeleteCircleRequest, ApiDeleteCircleRequest, int) ApiDeleteCircleRequest) *ApiDeleteCircleRequestStream {
	return self.ReduceInit(fn, ApiDeleteCircleRequest{})
}
func (self *ApiDeleteCircleRequestStream) ReduceInit(fn func(ApiDeleteCircleRequest, ApiDeleteCircleRequest, int) ApiDeleteCircleRequest, initialValue ApiDeleteCircleRequest) *ApiDeleteCircleRequestStream {
	result := ApiDeleteCircleRequestStreamOf()
	self.ForEach(func(v ApiDeleteCircleRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiDeleteCircleRequestStream) ReduceInterface(fn func(interface{}, ApiDeleteCircleRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiDeleteCircleRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiDeleteCircleRequestStream) ReduceString(fn func(string, ApiDeleteCircleRequest, int) string) []string {
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
func (self *ApiDeleteCircleRequestStream) ReduceInt(fn func(int, ApiDeleteCircleRequest, int) int) []int {
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
func (self *ApiDeleteCircleRequestStream) ReduceInt32(fn func(int32, ApiDeleteCircleRequest, int) int32) []int32 {
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
func (self *ApiDeleteCircleRequestStream) ReduceInt64(fn func(int64, ApiDeleteCircleRequest, int) int64) []int64 {
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
func (self *ApiDeleteCircleRequestStream) ReduceFloat32(fn func(float32, ApiDeleteCircleRequest, int) float32) []float32 {
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
func (self *ApiDeleteCircleRequestStream) ReduceFloat64(fn func(float64, ApiDeleteCircleRequest, int) float64) []float64 {
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
func (self *ApiDeleteCircleRequestStream) ReduceBool(fn func(bool, ApiDeleteCircleRequest, int) bool) []bool {
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
func (self *ApiDeleteCircleRequestStream) Reverse() *ApiDeleteCircleRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiDeleteCircleRequestStream) Replace(fn func(ApiDeleteCircleRequest, int) ApiDeleteCircleRequest) *ApiDeleteCircleRequestStream {
	return self.ForEach(func(v ApiDeleteCircleRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiDeleteCircleRequestStream) Select(fn func(ApiDeleteCircleRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiDeleteCircleRequestStream) Set(index int, val ApiDeleteCircleRequest) *ApiDeleteCircleRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiDeleteCircleRequestStream) Skip(skip int) *ApiDeleteCircleRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiDeleteCircleRequestStream) SkippingEach(fn func(ApiDeleteCircleRequest, int) int) *ApiDeleteCircleRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiDeleteCircleRequestStream) Slice(startIndex, n int) *ApiDeleteCircleRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiDeleteCircleRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiDeleteCircleRequestStream) Sort(fn func(i, j int) bool) *ApiDeleteCircleRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiDeleteCircleRequestStream) Tail() *ApiDeleteCircleRequest {
	return self.Last()
}
func (self *ApiDeleteCircleRequestStream) TailOr(arg ApiDeleteCircleRequest) ApiDeleteCircleRequest {
	return self.LastOr(arg)
}
func (self *ApiDeleteCircleRequestStream) ToList() []ApiDeleteCircleRequest {
	return self.Val()
}
func (self *ApiDeleteCircleRequestStream) Unique() *ApiDeleteCircleRequestStream {
	return self.Distinct()
}
func (self *ApiDeleteCircleRequestStream) Val() []ApiDeleteCircleRequest {
	if self == nil {
		return []ApiDeleteCircleRequest{}
	}
	return *self.Copy()
}
func (self *ApiDeleteCircleRequestStream) While(fn func(ApiDeleteCircleRequest, int) bool) *ApiDeleteCircleRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiDeleteCircleRequestStream) Where(fn func(ApiDeleteCircleRequest) bool) *ApiDeleteCircleRequestStream {
	result := ApiDeleteCircleRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiDeleteCircleRequestStream) WhereSlim(fn func(ApiDeleteCircleRequest) bool) *ApiDeleteCircleRequestStream {
	result := ApiDeleteCircleRequestStreamOf()
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
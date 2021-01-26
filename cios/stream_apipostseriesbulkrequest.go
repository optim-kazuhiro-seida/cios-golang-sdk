/*
 * Collection utility of ApiPostSeriesBulkRequest Struct
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

type ApiPostSeriesBulkRequestStream []ApiPostSeriesBulkRequest

func ApiPostSeriesBulkRequestStreamOf(arg ...ApiPostSeriesBulkRequest) ApiPostSeriesBulkRequestStream {
	return arg
}
func ApiPostSeriesBulkRequestStreamFrom(arg []ApiPostSeriesBulkRequest) ApiPostSeriesBulkRequestStream {
	return arg
}
func CreateApiPostSeriesBulkRequestStream(arg ...ApiPostSeriesBulkRequest) *ApiPostSeriesBulkRequestStream {
	tmp := ApiPostSeriesBulkRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiPostSeriesBulkRequestStream(arg []ApiPostSeriesBulkRequest) *ApiPostSeriesBulkRequestStream {
	tmp := ApiPostSeriesBulkRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiPostSeriesBulkRequestStream) Add(arg ApiPostSeriesBulkRequest) *ApiPostSeriesBulkRequestStream {
	return self.AddAll(arg)
}
func (self *ApiPostSeriesBulkRequestStream) AddAll(arg ...ApiPostSeriesBulkRequest) *ApiPostSeriesBulkRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiPostSeriesBulkRequestStream) AddSafe(arg *ApiPostSeriesBulkRequest) *ApiPostSeriesBulkRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiPostSeriesBulkRequestStream) Aggregate(fn func(ApiPostSeriesBulkRequest, ApiPostSeriesBulkRequest) ApiPostSeriesBulkRequest) *ApiPostSeriesBulkRequestStream {
	result := ApiPostSeriesBulkRequestStreamOf()
	self.ForEach(func(v ApiPostSeriesBulkRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiPostSeriesBulkRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiPostSeriesBulkRequestStream) AllMatch(fn func(ApiPostSeriesBulkRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiPostSeriesBulkRequestStream) AnyMatch(fn func(ApiPostSeriesBulkRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiPostSeriesBulkRequestStream) Clone() *ApiPostSeriesBulkRequestStream {
	temp := make([]ApiPostSeriesBulkRequest, self.Len())
	copy(temp, *self)
	return (*ApiPostSeriesBulkRequestStream)(&temp)
}
func (self *ApiPostSeriesBulkRequestStream) Copy() *ApiPostSeriesBulkRequestStream {
	return self.Clone()
}
func (self *ApiPostSeriesBulkRequestStream) Concat(arg []ApiPostSeriesBulkRequest) *ApiPostSeriesBulkRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiPostSeriesBulkRequestStream) Contains(arg ApiPostSeriesBulkRequest) bool {
	return self.FindIndex(func(_arg ApiPostSeriesBulkRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiPostSeriesBulkRequestStream) Clean() *ApiPostSeriesBulkRequestStream {
	*self = ApiPostSeriesBulkRequestStreamOf()
	return self
}
func (self *ApiPostSeriesBulkRequestStream) Delete(index int) *ApiPostSeriesBulkRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiPostSeriesBulkRequestStream) DeleteRange(startIndex, endIndex int) *ApiPostSeriesBulkRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiPostSeriesBulkRequestStream) Distinct() *ApiPostSeriesBulkRequestStream {
	caches := map[string]bool{}
	result := ApiPostSeriesBulkRequestStreamOf()
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
func (self *ApiPostSeriesBulkRequestStream) Each(fn func(ApiPostSeriesBulkRequest)) *ApiPostSeriesBulkRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiPostSeriesBulkRequestStream) EachRight(fn func(ApiPostSeriesBulkRequest)) *ApiPostSeriesBulkRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiPostSeriesBulkRequestStream) Equals(arr []ApiPostSeriesBulkRequest) bool {
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
func (self *ApiPostSeriesBulkRequestStream) Filter(fn func(ApiPostSeriesBulkRequest, int) bool) *ApiPostSeriesBulkRequestStream {
	result := ApiPostSeriesBulkRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiPostSeriesBulkRequestStream) FilterSlim(fn func(ApiPostSeriesBulkRequest, int) bool) *ApiPostSeriesBulkRequestStream {
	result := ApiPostSeriesBulkRequestStreamOf()
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
func (self *ApiPostSeriesBulkRequestStream) Find(fn func(ApiPostSeriesBulkRequest, int) bool) *ApiPostSeriesBulkRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiPostSeriesBulkRequestStream) FindOr(fn func(ApiPostSeriesBulkRequest, int) bool, or ApiPostSeriesBulkRequest) ApiPostSeriesBulkRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiPostSeriesBulkRequestStream) FindIndex(fn func(ApiPostSeriesBulkRequest, int) bool) int {
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
func (self *ApiPostSeriesBulkRequestStream) First() *ApiPostSeriesBulkRequest {
	return self.Get(0)
}
func (self *ApiPostSeriesBulkRequestStream) FirstOr(arg ApiPostSeriesBulkRequest) ApiPostSeriesBulkRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiPostSeriesBulkRequestStream) ForEach(fn func(ApiPostSeriesBulkRequest, int)) *ApiPostSeriesBulkRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiPostSeriesBulkRequestStream) ForEachRight(fn func(ApiPostSeriesBulkRequest, int)) *ApiPostSeriesBulkRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiPostSeriesBulkRequestStream) GroupBy(fn func(ApiPostSeriesBulkRequest, int) string) map[string][]ApiPostSeriesBulkRequest {
	m := map[string][]ApiPostSeriesBulkRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiPostSeriesBulkRequestStream) GroupByValues(fn func(ApiPostSeriesBulkRequest, int) string) [][]ApiPostSeriesBulkRequest {
	var tmp [][]ApiPostSeriesBulkRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiPostSeriesBulkRequestStream) IndexOf(arg ApiPostSeriesBulkRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiPostSeriesBulkRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiPostSeriesBulkRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiPostSeriesBulkRequestStream) Last() *ApiPostSeriesBulkRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiPostSeriesBulkRequestStream) LastOr(arg ApiPostSeriesBulkRequest) ApiPostSeriesBulkRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiPostSeriesBulkRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiPostSeriesBulkRequestStream) Limit(limit int) *ApiPostSeriesBulkRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiPostSeriesBulkRequestStream) Map(fn func(ApiPostSeriesBulkRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiPostSeriesBulkRequestStream) Map2Int(fn func(ApiPostSeriesBulkRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiPostSeriesBulkRequestStream) Map2Int32(fn func(ApiPostSeriesBulkRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiPostSeriesBulkRequestStream) Map2Int64(fn func(ApiPostSeriesBulkRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiPostSeriesBulkRequestStream) Map2Float32(fn func(ApiPostSeriesBulkRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiPostSeriesBulkRequestStream) Map2Float64(fn func(ApiPostSeriesBulkRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiPostSeriesBulkRequestStream) Map2Bool(fn func(ApiPostSeriesBulkRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiPostSeriesBulkRequestStream) Map2Bytes(fn func(ApiPostSeriesBulkRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiPostSeriesBulkRequestStream) Map2String(fn func(ApiPostSeriesBulkRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiPostSeriesBulkRequestStream) Max(fn func(ApiPostSeriesBulkRequest, int) float64) *ApiPostSeriesBulkRequest {
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
func (self *ApiPostSeriesBulkRequestStream) Min(fn func(ApiPostSeriesBulkRequest, int) float64) *ApiPostSeriesBulkRequest {
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
func (self *ApiPostSeriesBulkRequestStream) NoneMatch(fn func(ApiPostSeriesBulkRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiPostSeriesBulkRequestStream) Get(index int) *ApiPostSeriesBulkRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiPostSeriesBulkRequestStream) GetOr(index int, arg ApiPostSeriesBulkRequest) ApiPostSeriesBulkRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiPostSeriesBulkRequestStream) Peek(fn func(*ApiPostSeriesBulkRequest, int)) *ApiPostSeriesBulkRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *ApiPostSeriesBulkRequestStream) Reduce(fn func(ApiPostSeriesBulkRequest, ApiPostSeriesBulkRequest, int) ApiPostSeriesBulkRequest) *ApiPostSeriesBulkRequestStream {
	return self.ReduceInit(fn, ApiPostSeriesBulkRequest{})
}
func (self *ApiPostSeriesBulkRequestStream) ReduceInit(fn func(ApiPostSeriesBulkRequest, ApiPostSeriesBulkRequest, int) ApiPostSeriesBulkRequest, initialValue ApiPostSeriesBulkRequest) *ApiPostSeriesBulkRequestStream {
	result := ApiPostSeriesBulkRequestStreamOf()
	self.ForEach(func(v ApiPostSeriesBulkRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiPostSeriesBulkRequestStream) ReduceInterface(fn func(interface{}, ApiPostSeriesBulkRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiPostSeriesBulkRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiPostSeriesBulkRequestStream) ReduceString(fn func(string, ApiPostSeriesBulkRequest, int) string) []string {
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
func (self *ApiPostSeriesBulkRequestStream) ReduceInt(fn func(int, ApiPostSeriesBulkRequest, int) int) []int {
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
func (self *ApiPostSeriesBulkRequestStream) ReduceInt32(fn func(int32, ApiPostSeriesBulkRequest, int) int32) []int32 {
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
func (self *ApiPostSeriesBulkRequestStream) ReduceInt64(fn func(int64, ApiPostSeriesBulkRequest, int) int64) []int64 {
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
func (self *ApiPostSeriesBulkRequestStream) ReduceFloat32(fn func(float32, ApiPostSeriesBulkRequest, int) float32) []float32 {
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
func (self *ApiPostSeriesBulkRequestStream) ReduceFloat64(fn func(float64, ApiPostSeriesBulkRequest, int) float64) []float64 {
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
func (self *ApiPostSeriesBulkRequestStream) ReduceBool(fn func(bool, ApiPostSeriesBulkRequest, int) bool) []bool {
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
func (self *ApiPostSeriesBulkRequestStream) Reverse() *ApiPostSeriesBulkRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiPostSeriesBulkRequestStream) Replace(fn func(ApiPostSeriesBulkRequest, int) ApiPostSeriesBulkRequest) *ApiPostSeriesBulkRequestStream {
	return self.ForEach(func(v ApiPostSeriesBulkRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiPostSeriesBulkRequestStream) Select(fn func(ApiPostSeriesBulkRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiPostSeriesBulkRequestStream) Set(index int, val ApiPostSeriesBulkRequest) *ApiPostSeriesBulkRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiPostSeriesBulkRequestStream) Skip(skip int) *ApiPostSeriesBulkRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiPostSeriesBulkRequestStream) SkippingEach(fn func(ApiPostSeriesBulkRequest, int) int) *ApiPostSeriesBulkRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiPostSeriesBulkRequestStream) Slice(startIndex, n int) *ApiPostSeriesBulkRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiPostSeriesBulkRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiPostSeriesBulkRequestStream) Sort(fn func(i, j int) bool) *ApiPostSeriesBulkRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiPostSeriesBulkRequestStream) Tail() *ApiPostSeriesBulkRequest {
	return self.Last()
}
func (self *ApiPostSeriesBulkRequestStream) TailOr(arg ApiPostSeriesBulkRequest) ApiPostSeriesBulkRequest {
	return self.LastOr(arg)
}
func (self *ApiPostSeriesBulkRequestStream) ToList() []ApiPostSeriesBulkRequest {
	return self.Val()
}
func (self *ApiPostSeriesBulkRequestStream) Unique() *ApiPostSeriesBulkRequestStream {
	return self.Distinct()
}
func (self *ApiPostSeriesBulkRequestStream) Val() []ApiPostSeriesBulkRequest {
	if self == nil {
		return []ApiPostSeriesBulkRequest{}
	}
	return *self.Copy()
}
func (self *ApiPostSeriesBulkRequestStream) While(fn func(ApiPostSeriesBulkRequest, int) bool) *ApiPostSeriesBulkRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiPostSeriesBulkRequestStream) Where(fn func(ApiPostSeriesBulkRequest) bool) *ApiPostSeriesBulkRequestStream {
	result := ApiPostSeriesBulkRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiPostSeriesBulkRequestStream) WhereSlim(fn func(ApiPostSeriesBulkRequest) bool) *ApiPostSeriesBulkRequestStream {
	result := ApiPostSeriesBulkRequestStreamOf()
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
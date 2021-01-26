/*
 * Collection utility of ApiGetGroupsRequest Struct
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

type ApiGetGroupsRequestStream []ApiGetGroupsRequest

func ApiGetGroupsRequestStreamOf(arg ...ApiGetGroupsRequest) ApiGetGroupsRequestStream {
	return arg
}
func ApiGetGroupsRequestStreamFrom(arg []ApiGetGroupsRequest) ApiGetGroupsRequestStream {
	return arg
}
func CreateApiGetGroupsRequestStream(arg ...ApiGetGroupsRequest) *ApiGetGroupsRequestStream {
	tmp := ApiGetGroupsRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiGetGroupsRequestStream(arg []ApiGetGroupsRequest) *ApiGetGroupsRequestStream {
	tmp := ApiGetGroupsRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiGetGroupsRequestStream) Add(arg ApiGetGroupsRequest) *ApiGetGroupsRequestStream {
	return self.AddAll(arg)
}
func (self *ApiGetGroupsRequestStream) AddAll(arg ...ApiGetGroupsRequest) *ApiGetGroupsRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiGetGroupsRequestStream) AddSafe(arg *ApiGetGroupsRequest) *ApiGetGroupsRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiGetGroupsRequestStream) Aggregate(fn func(ApiGetGroupsRequest, ApiGetGroupsRequest) ApiGetGroupsRequest) *ApiGetGroupsRequestStream {
	result := ApiGetGroupsRequestStreamOf()
	self.ForEach(func(v ApiGetGroupsRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiGetGroupsRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiGetGroupsRequestStream) AllMatch(fn func(ApiGetGroupsRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiGetGroupsRequestStream) AnyMatch(fn func(ApiGetGroupsRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiGetGroupsRequestStream) Clone() *ApiGetGroupsRequestStream {
	temp := make([]ApiGetGroupsRequest, self.Len())
	copy(temp, *self)
	return (*ApiGetGroupsRequestStream)(&temp)
}
func (self *ApiGetGroupsRequestStream) Copy() *ApiGetGroupsRequestStream {
	return self.Clone()
}
func (self *ApiGetGroupsRequestStream) Concat(arg []ApiGetGroupsRequest) *ApiGetGroupsRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiGetGroupsRequestStream) Contains(arg ApiGetGroupsRequest) bool {
	return self.FindIndex(func(_arg ApiGetGroupsRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiGetGroupsRequestStream) Clean() *ApiGetGroupsRequestStream {
	*self = ApiGetGroupsRequestStreamOf()
	return self
}
func (self *ApiGetGroupsRequestStream) Delete(index int) *ApiGetGroupsRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiGetGroupsRequestStream) DeleteRange(startIndex, endIndex int) *ApiGetGroupsRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiGetGroupsRequestStream) Distinct() *ApiGetGroupsRequestStream {
	caches := map[string]bool{}
	result := ApiGetGroupsRequestStreamOf()
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
func (self *ApiGetGroupsRequestStream) Each(fn func(ApiGetGroupsRequest)) *ApiGetGroupsRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiGetGroupsRequestStream) EachRight(fn func(ApiGetGroupsRequest)) *ApiGetGroupsRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiGetGroupsRequestStream) Equals(arr []ApiGetGroupsRequest) bool {
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
func (self *ApiGetGroupsRequestStream) Filter(fn func(ApiGetGroupsRequest, int) bool) *ApiGetGroupsRequestStream {
	result := ApiGetGroupsRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetGroupsRequestStream) FilterSlim(fn func(ApiGetGroupsRequest, int) bool) *ApiGetGroupsRequestStream {
	result := ApiGetGroupsRequestStreamOf()
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
func (self *ApiGetGroupsRequestStream) Find(fn func(ApiGetGroupsRequest, int) bool) *ApiGetGroupsRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiGetGroupsRequestStream) FindOr(fn func(ApiGetGroupsRequest, int) bool, or ApiGetGroupsRequest) ApiGetGroupsRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiGetGroupsRequestStream) FindIndex(fn func(ApiGetGroupsRequest, int) bool) int {
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
func (self *ApiGetGroupsRequestStream) First() *ApiGetGroupsRequest {
	return self.Get(0)
}
func (self *ApiGetGroupsRequestStream) FirstOr(arg ApiGetGroupsRequest) ApiGetGroupsRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetGroupsRequestStream) ForEach(fn func(ApiGetGroupsRequest, int)) *ApiGetGroupsRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiGetGroupsRequestStream) ForEachRight(fn func(ApiGetGroupsRequest, int)) *ApiGetGroupsRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiGetGroupsRequestStream) GroupBy(fn func(ApiGetGroupsRequest, int) string) map[string][]ApiGetGroupsRequest {
	m := map[string][]ApiGetGroupsRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiGetGroupsRequestStream) GroupByValues(fn func(ApiGetGroupsRequest, int) string) [][]ApiGetGroupsRequest {
	var tmp [][]ApiGetGroupsRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiGetGroupsRequestStream) IndexOf(arg ApiGetGroupsRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiGetGroupsRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiGetGroupsRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiGetGroupsRequestStream) Last() *ApiGetGroupsRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiGetGroupsRequestStream) LastOr(arg ApiGetGroupsRequest) ApiGetGroupsRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetGroupsRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiGetGroupsRequestStream) Limit(limit int) *ApiGetGroupsRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiGetGroupsRequestStream) Map(fn func(ApiGetGroupsRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetGroupsRequestStream) Map2Int(fn func(ApiGetGroupsRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetGroupsRequestStream) Map2Int32(fn func(ApiGetGroupsRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetGroupsRequestStream) Map2Int64(fn func(ApiGetGroupsRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetGroupsRequestStream) Map2Float32(fn func(ApiGetGroupsRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetGroupsRequestStream) Map2Float64(fn func(ApiGetGroupsRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetGroupsRequestStream) Map2Bool(fn func(ApiGetGroupsRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetGroupsRequestStream) Map2Bytes(fn func(ApiGetGroupsRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetGroupsRequestStream) Map2String(fn func(ApiGetGroupsRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetGroupsRequestStream) Max(fn func(ApiGetGroupsRequest, int) float64) *ApiGetGroupsRequest {
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
func (self *ApiGetGroupsRequestStream) Min(fn func(ApiGetGroupsRequest, int) float64) *ApiGetGroupsRequest {
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
func (self *ApiGetGroupsRequestStream) NoneMatch(fn func(ApiGetGroupsRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiGetGroupsRequestStream) Get(index int) *ApiGetGroupsRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiGetGroupsRequestStream) GetOr(index int, arg ApiGetGroupsRequest) ApiGetGroupsRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetGroupsRequestStream) Peek(fn func(*ApiGetGroupsRequest, int)) *ApiGetGroupsRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *ApiGetGroupsRequestStream) Reduce(fn func(ApiGetGroupsRequest, ApiGetGroupsRequest, int) ApiGetGroupsRequest) *ApiGetGroupsRequestStream {
	return self.ReduceInit(fn, ApiGetGroupsRequest{})
}
func (self *ApiGetGroupsRequestStream) ReduceInit(fn func(ApiGetGroupsRequest, ApiGetGroupsRequest, int) ApiGetGroupsRequest, initialValue ApiGetGroupsRequest) *ApiGetGroupsRequestStream {
	result := ApiGetGroupsRequestStreamOf()
	self.ForEach(func(v ApiGetGroupsRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiGetGroupsRequestStream) ReduceInterface(fn func(interface{}, ApiGetGroupsRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiGetGroupsRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetGroupsRequestStream) ReduceString(fn func(string, ApiGetGroupsRequest, int) string) []string {
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
func (self *ApiGetGroupsRequestStream) ReduceInt(fn func(int, ApiGetGroupsRequest, int) int) []int {
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
func (self *ApiGetGroupsRequestStream) ReduceInt32(fn func(int32, ApiGetGroupsRequest, int) int32) []int32 {
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
func (self *ApiGetGroupsRequestStream) ReduceInt64(fn func(int64, ApiGetGroupsRequest, int) int64) []int64 {
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
func (self *ApiGetGroupsRequestStream) ReduceFloat32(fn func(float32, ApiGetGroupsRequest, int) float32) []float32 {
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
func (self *ApiGetGroupsRequestStream) ReduceFloat64(fn func(float64, ApiGetGroupsRequest, int) float64) []float64 {
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
func (self *ApiGetGroupsRequestStream) ReduceBool(fn func(bool, ApiGetGroupsRequest, int) bool) []bool {
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
func (self *ApiGetGroupsRequestStream) Reverse() *ApiGetGroupsRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiGetGroupsRequestStream) Replace(fn func(ApiGetGroupsRequest, int) ApiGetGroupsRequest) *ApiGetGroupsRequestStream {
	return self.ForEach(func(v ApiGetGroupsRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiGetGroupsRequestStream) Select(fn func(ApiGetGroupsRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiGetGroupsRequestStream) Set(index int, val ApiGetGroupsRequest) *ApiGetGroupsRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiGetGroupsRequestStream) Skip(skip int) *ApiGetGroupsRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiGetGroupsRequestStream) SkippingEach(fn func(ApiGetGroupsRequest, int) int) *ApiGetGroupsRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiGetGroupsRequestStream) Slice(startIndex, n int) *ApiGetGroupsRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiGetGroupsRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiGetGroupsRequestStream) Sort(fn func(i, j int) bool) *ApiGetGroupsRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiGetGroupsRequestStream) Tail() *ApiGetGroupsRequest {
	return self.Last()
}
func (self *ApiGetGroupsRequestStream) TailOr(arg ApiGetGroupsRequest) ApiGetGroupsRequest {
	return self.LastOr(arg)
}
func (self *ApiGetGroupsRequestStream) ToList() []ApiGetGroupsRequest {
	return self.Val()
}
func (self *ApiGetGroupsRequestStream) Unique() *ApiGetGroupsRequestStream {
	return self.Distinct()
}
func (self *ApiGetGroupsRequestStream) Val() []ApiGetGroupsRequest {
	if self == nil {
		return []ApiGetGroupsRequest{}
	}
	return *self.Copy()
}
func (self *ApiGetGroupsRequestStream) While(fn func(ApiGetGroupsRequest, int) bool) *ApiGetGroupsRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiGetGroupsRequestStream) Where(fn func(ApiGetGroupsRequest) bool) *ApiGetGroupsRequestStream {
	result := ApiGetGroupsRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetGroupsRequestStream) WhereSlim(fn func(ApiGetGroupsRequest) bool) *ApiGetGroupsRequestStream {
	result := ApiGetGroupsRequestStreamOf()
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
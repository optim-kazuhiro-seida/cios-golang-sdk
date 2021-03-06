/*
 * Collection utility of ApiSetMemberRequest Struct
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

type ApiSetMemberRequestStream []ApiSetMemberRequest

func ApiSetMemberRequestStreamOf(arg ...ApiSetMemberRequest) ApiSetMemberRequestStream {
	return arg
}
func ApiSetMemberRequestStreamFrom(arg []ApiSetMemberRequest) ApiSetMemberRequestStream {
	return arg
}
func CreateApiSetMemberRequestStream(arg ...ApiSetMemberRequest) *ApiSetMemberRequestStream {
	tmp := ApiSetMemberRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiSetMemberRequestStream(arg []ApiSetMemberRequest) *ApiSetMemberRequestStream {
	tmp := ApiSetMemberRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiSetMemberRequestStream) Add(arg ApiSetMemberRequest) *ApiSetMemberRequestStream {
	return self.AddAll(arg)
}
func (self *ApiSetMemberRequestStream) AddAll(arg ...ApiSetMemberRequest) *ApiSetMemberRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiSetMemberRequestStream) AddSafe(arg *ApiSetMemberRequest) *ApiSetMemberRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiSetMemberRequestStream) Aggregate(fn func(ApiSetMemberRequest, ApiSetMemberRequest) ApiSetMemberRequest) *ApiSetMemberRequestStream {
	result := ApiSetMemberRequestStreamOf()
	self.ForEach(func(v ApiSetMemberRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiSetMemberRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiSetMemberRequestStream) AllMatch(fn func(ApiSetMemberRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiSetMemberRequestStream) AnyMatch(fn func(ApiSetMemberRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiSetMemberRequestStream) Clone() *ApiSetMemberRequestStream {
	temp := make([]ApiSetMemberRequest, self.Len())
	copy(temp, *self)
	return (*ApiSetMemberRequestStream)(&temp)
}
func (self *ApiSetMemberRequestStream) Copy() *ApiSetMemberRequestStream {
	return self.Clone()
}
func (self *ApiSetMemberRequestStream) Concat(arg []ApiSetMemberRequest) *ApiSetMemberRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiSetMemberRequestStream) Contains(arg ApiSetMemberRequest) bool {
	return self.FindIndex(func(_arg ApiSetMemberRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiSetMemberRequestStream) Clean() *ApiSetMemberRequestStream {
	*self = ApiSetMemberRequestStreamOf()
	return self
}
func (self *ApiSetMemberRequestStream) Delete(index int) *ApiSetMemberRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiSetMemberRequestStream) DeleteRange(startIndex, endIndex int) *ApiSetMemberRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiSetMemberRequestStream) Distinct() *ApiSetMemberRequestStream {
	caches := map[string]bool{}
	result := ApiSetMemberRequestStreamOf()
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
func (self *ApiSetMemberRequestStream) Each(fn func(ApiSetMemberRequest)) *ApiSetMemberRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiSetMemberRequestStream) EachRight(fn func(ApiSetMemberRequest)) *ApiSetMemberRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiSetMemberRequestStream) Equals(arr []ApiSetMemberRequest) bool {
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
func (self *ApiSetMemberRequestStream) Filter(fn func(ApiSetMemberRequest, int) bool) *ApiSetMemberRequestStream {
	result := ApiSetMemberRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiSetMemberRequestStream) FilterSlim(fn func(ApiSetMemberRequest, int) bool) *ApiSetMemberRequestStream {
	result := ApiSetMemberRequestStreamOf()
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
func (self *ApiSetMemberRequestStream) Find(fn func(ApiSetMemberRequest, int) bool) *ApiSetMemberRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiSetMemberRequestStream) FindOr(fn func(ApiSetMemberRequest, int) bool, or ApiSetMemberRequest) ApiSetMemberRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiSetMemberRequestStream) FindIndex(fn func(ApiSetMemberRequest, int) bool) int {
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
func (self *ApiSetMemberRequestStream) First() *ApiSetMemberRequest {
	return self.Get(0)
}
func (self *ApiSetMemberRequestStream) FirstOr(arg ApiSetMemberRequest) ApiSetMemberRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiSetMemberRequestStream) ForEach(fn func(ApiSetMemberRequest, int)) *ApiSetMemberRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiSetMemberRequestStream) ForEachRight(fn func(ApiSetMemberRequest, int)) *ApiSetMemberRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiSetMemberRequestStream) GroupBy(fn func(ApiSetMemberRequest, int) string) map[string][]ApiSetMemberRequest {
	m := map[string][]ApiSetMemberRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiSetMemberRequestStream) GroupByValues(fn func(ApiSetMemberRequest, int) string) [][]ApiSetMemberRequest {
	var tmp [][]ApiSetMemberRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiSetMemberRequestStream) IndexOf(arg ApiSetMemberRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiSetMemberRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiSetMemberRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiSetMemberRequestStream) Last() *ApiSetMemberRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiSetMemberRequestStream) LastOr(arg ApiSetMemberRequest) ApiSetMemberRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiSetMemberRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiSetMemberRequestStream) Limit(limit int) *ApiSetMemberRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiSetMemberRequestStream) Map(fn func(ApiSetMemberRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiSetMemberRequestStream) Map2Int(fn func(ApiSetMemberRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiSetMemberRequestStream) Map2Int32(fn func(ApiSetMemberRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiSetMemberRequestStream) Map2Int64(fn func(ApiSetMemberRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiSetMemberRequestStream) Map2Float32(fn func(ApiSetMemberRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiSetMemberRequestStream) Map2Float64(fn func(ApiSetMemberRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiSetMemberRequestStream) Map2Bool(fn func(ApiSetMemberRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiSetMemberRequestStream) Map2Bytes(fn func(ApiSetMemberRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiSetMemberRequestStream) Map2String(fn func(ApiSetMemberRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiSetMemberRequestStream) Max(fn func(ApiSetMemberRequest, int) float64) *ApiSetMemberRequest {
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
func (self *ApiSetMemberRequestStream) Min(fn func(ApiSetMemberRequest, int) float64) *ApiSetMemberRequest {
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
func (self *ApiSetMemberRequestStream) NoneMatch(fn func(ApiSetMemberRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiSetMemberRequestStream) Get(index int) *ApiSetMemberRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiSetMemberRequestStream) GetOr(index int, arg ApiSetMemberRequest) ApiSetMemberRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiSetMemberRequestStream) Peek(fn func(*ApiSetMemberRequest, int)) *ApiSetMemberRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiSetMemberRequestStream) Reduce(fn func(ApiSetMemberRequest, ApiSetMemberRequest, int) ApiSetMemberRequest) *ApiSetMemberRequestStream {
	return self.ReduceInit(fn, ApiSetMemberRequest{})
}
func (self *ApiSetMemberRequestStream) ReduceInit(fn func(ApiSetMemberRequest, ApiSetMemberRequest, int) ApiSetMemberRequest, initialValue ApiSetMemberRequest) *ApiSetMemberRequestStream {
	result := ApiSetMemberRequestStreamOf()
	self.ForEach(func(v ApiSetMemberRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiSetMemberRequestStream) ReduceInterface(fn func(interface{}, ApiSetMemberRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiSetMemberRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiSetMemberRequestStream) ReduceString(fn func(string, ApiSetMemberRequest, int) string) []string {
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
func (self *ApiSetMemberRequestStream) ReduceInt(fn func(int, ApiSetMemberRequest, int) int) []int {
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
func (self *ApiSetMemberRequestStream) ReduceInt32(fn func(int32, ApiSetMemberRequest, int) int32) []int32 {
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
func (self *ApiSetMemberRequestStream) ReduceInt64(fn func(int64, ApiSetMemberRequest, int) int64) []int64 {
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
func (self *ApiSetMemberRequestStream) ReduceFloat32(fn func(float32, ApiSetMemberRequest, int) float32) []float32 {
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
func (self *ApiSetMemberRequestStream) ReduceFloat64(fn func(float64, ApiSetMemberRequest, int) float64) []float64 {
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
func (self *ApiSetMemberRequestStream) ReduceBool(fn func(bool, ApiSetMemberRequest, int) bool) []bool {
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
func (self *ApiSetMemberRequestStream) Reverse() *ApiSetMemberRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiSetMemberRequestStream) Replace(fn func(ApiSetMemberRequest, int) ApiSetMemberRequest) *ApiSetMemberRequestStream {
	return self.ForEach(func(v ApiSetMemberRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiSetMemberRequestStream) Select(fn func(ApiSetMemberRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiSetMemberRequestStream) Set(index int, val ApiSetMemberRequest) *ApiSetMemberRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiSetMemberRequestStream) Skip(skip int) *ApiSetMemberRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiSetMemberRequestStream) SkippingEach(fn func(ApiSetMemberRequest, int) int) *ApiSetMemberRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiSetMemberRequestStream) Slice(startIndex, n int) *ApiSetMemberRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiSetMemberRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiSetMemberRequestStream) Sort(fn func(i, j int) bool) *ApiSetMemberRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiSetMemberRequestStream) Tail() *ApiSetMemberRequest {
	return self.Last()
}
func (self *ApiSetMemberRequestStream) TailOr(arg ApiSetMemberRequest) ApiSetMemberRequest {
	return self.LastOr(arg)
}
func (self *ApiSetMemberRequestStream) ToList() []ApiSetMemberRequest {
	return self.Val()
}
func (self *ApiSetMemberRequestStream) Unique() *ApiSetMemberRequestStream {
	return self.Distinct()
}
func (self *ApiSetMemberRequestStream) Val() []ApiSetMemberRequest {
	if self == nil {
		return []ApiSetMemberRequest{}
	}
	return *self.Copy()
}
func (self *ApiSetMemberRequestStream) While(fn func(ApiSetMemberRequest, int) bool) *ApiSetMemberRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiSetMemberRequestStream) Where(fn func(ApiSetMemberRequest) bool) *ApiSetMemberRequestStream {
	result := ApiSetMemberRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiSetMemberRequestStream) WhereSlim(fn func(ApiSetMemberRequest) bool) *ApiSetMemberRequestStream {
	result := ApiSetMemberRequestStreamOf()
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

/*
 * Collection utility of VideoUpdateRequest Struct
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

type VideoUpdateRequestStream []VideoUpdateRequest

func VideoUpdateRequestStreamOf(arg ...VideoUpdateRequest) VideoUpdateRequestStream {
	return arg
}
func VideoUpdateRequestStreamFrom(arg []VideoUpdateRequest) VideoUpdateRequestStream {
	return arg
}
func CreateVideoUpdateRequestStream(arg ...VideoUpdateRequest) *VideoUpdateRequestStream {
	tmp := VideoUpdateRequestStreamOf(arg...)
	return &tmp
}
func GenerateVideoUpdateRequestStream(arg []VideoUpdateRequest) *VideoUpdateRequestStream {
	tmp := VideoUpdateRequestStreamFrom(arg)
	return &tmp
}

func (self *VideoUpdateRequestStream) Add(arg VideoUpdateRequest) *VideoUpdateRequestStream {
	return self.AddAll(arg)
}
func (self *VideoUpdateRequestStream) AddAll(arg ...VideoUpdateRequest) *VideoUpdateRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *VideoUpdateRequestStream) AddSafe(arg *VideoUpdateRequest) *VideoUpdateRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *VideoUpdateRequestStream) Aggregate(fn func(VideoUpdateRequest, VideoUpdateRequest) VideoUpdateRequest) *VideoUpdateRequestStream {
	result := VideoUpdateRequestStreamOf()
	self.ForEach(func(v VideoUpdateRequest, i int) {
		if i == 0 {
			result.Add(fn(VideoUpdateRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *VideoUpdateRequestStream) AllMatch(fn func(VideoUpdateRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *VideoUpdateRequestStream) AnyMatch(fn func(VideoUpdateRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *VideoUpdateRequestStream) Clone() *VideoUpdateRequestStream {
	temp := make([]VideoUpdateRequest, self.Len())
	copy(temp, *self)
	return (*VideoUpdateRequestStream)(&temp)
}
func (self *VideoUpdateRequestStream) Copy() *VideoUpdateRequestStream {
	return self.Clone()
}
func (self *VideoUpdateRequestStream) Concat(arg []VideoUpdateRequest) *VideoUpdateRequestStream {
	return self.AddAll(arg...)
}
func (self *VideoUpdateRequestStream) Contains(arg VideoUpdateRequest) bool {
	return self.FindIndex(func(_arg VideoUpdateRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *VideoUpdateRequestStream) Clean() *VideoUpdateRequestStream {
	*self = VideoUpdateRequestStreamOf()
	return self
}
func (self *VideoUpdateRequestStream) Delete(index int) *VideoUpdateRequestStream {
	return self.DeleteRange(index, index)
}
func (self *VideoUpdateRequestStream) DeleteRange(startIndex, endIndex int) *VideoUpdateRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *VideoUpdateRequestStream) Distinct() *VideoUpdateRequestStream {
	caches := map[string]bool{}
	result := VideoUpdateRequestStreamOf()
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
func (self *VideoUpdateRequestStream) Each(fn func(VideoUpdateRequest)) *VideoUpdateRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *VideoUpdateRequestStream) EachRight(fn func(VideoUpdateRequest)) *VideoUpdateRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *VideoUpdateRequestStream) Equals(arr []VideoUpdateRequest) bool {
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
func (self *VideoUpdateRequestStream) Filter(fn func(VideoUpdateRequest, int) bool) *VideoUpdateRequestStream {
	result := VideoUpdateRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *VideoUpdateRequestStream) FilterSlim(fn func(VideoUpdateRequest, int) bool) *VideoUpdateRequestStream {
	result := VideoUpdateRequestStreamOf()
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
func (self *VideoUpdateRequestStream) Find(fn func(VideoUpdateRequest, int) bool) *VideoUpdateRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *VideoUpdateRequestStream) FindOr(fn func(VideoUpdateRequest, int) bool, or VideoUpdateRequest) VideoUpdateRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *VideoUpdateRequestStream) FindIndex(fn func(VideoUpdateRequest, int) bool) int {
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
func (self *VideoUpdateRequestStream) First() *VideoUpdateRequest {
	return self.Get(0)
}
func (self *VideoUpdateRequestStream) FirstOr(arg VideoUpdateRequest) VideoUpdateRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *VideoUpdateRequestStream) ForEach(fn func(VideoUpdateRequest, int)) *VideoUpdateRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *VideoUpdateRequestStream) ForEachRight(fn func(VideoUpdateRequest, int)) *VideoUpdateRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *VideoUpdateRequestStream) GroupBy(fn func(VideoUpdateRequest, int) string) map[string][]VideoUpdateRequest {
	m := map[string][]VideoUpdateRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *VideoUpdateRequestStream) GroupByValues(fn func(VideoUpdateRequest, int) string) [][]VideoUpdateRequest {
	var tmp [][]VideoUpdateRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *VideoUpdateRequestStream) IndexOf(arg VideoUpdateRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *VideoUpdateRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *VideoUpdateRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *VideoUpdateRequestStream) Last() *VideoUpdateRequest {
	return self.Get(self.Len() - 1)
}
func (self *VideoUpdateRequestStream) LastOr(arg VideoUpdateRequest) VideoUpdateRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *VideoUpdateRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *VideoUpdateRequestStream) Limit(limit int) *VideoUpdateRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *VideoUpdateRequestStream) Map(fn func(VideoUpdateRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *VideoUpdateRequestStream) Map2Int(fn func(VideoUpdateRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *VideoUpdateRequestStream) Map2Int32(fn func(VideoUpdateRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *VideoUpdateRequestStream) Map2Int64(fn func(VideoUpdateRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *VideoUpdateRequestStream) Map2Float32(fn func(VideoUpdateRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *VideoUpdateRequestStream) Map2Float64(fn func(VideoUpdateRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *VideoUpdateRequestStream) Map2Bool(fn func(VideoUpdateRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *VideoUpdateRequestStream) Map2Bytes(fn func(VideoUpdateRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *VideoUpdateRequestStream) Map2String(fn func(VideoUpdateRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *VideoUpdateRequestStream) Max(fn func(VideoUpdateRequest, int) float64) *VideoUpdateRequest {
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
func (self *VideoUpdateRequestStream) Min(fn func(VideoUpdateRequest, int) float64) *VideoUpdateRequest {
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
func (self *VideoUpdateRequestStream) NoneMatch(fn func(VideoUpdateRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *VideoUpdateRequestStream) Get(index int) *VideoUpdateRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *VideoUpdateRequestStream) GetOr(index int, arg VideoUpdateRequest) VideoUpdateRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *VideoUpdateRequestStream) Peek(fn func(*VideoUpdateRequest, int)) *VideoUpdateRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *VideoUpdateRequestStream) Reduce(fn func(VideoUpdateRequest, VideoUpdateRequest, int) VideoUpdateRequest) *VideoUpdateRequestStream {
	return self.ReduceInit(fn, VideoUpdateRequest{})
}
func (self *VideoUpdateRequestStream) ReduceInit(fn func(VideoUpdateRequest, VideoUpdateRequest, int) VideoUpdateRequest, initialValue VideoUpdateRequest) *VideoUpdateRequestStream {
	result := VideoUpdateRequestStreamOf()
	self.ForEach(func(v VideoUpdateRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *VideoUpdateRequestStream) ReduceInterface(fn func(interface{}, VideoUpdateRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(VideoUpdateRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *VideoUpdateRequestStream) ReduceString(fn func(string, VideoUpdateRequest, int) string) []string {
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
func (self *VideoUpdateRequestStream) ReduceInt(fn func(int, VideoUpdateRequest, int) int) []int {
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
func (self *VideoUpdateRequestStream) ReduceInt32(fn func(int32, VideoUpdateRequest, int) int32) []int32 {
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
func (self *VideoUpdateRequestStream) ReduceInt64(fn func(int64, VideoUpdateRequest, int) int64) []int64 {
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
func (self *VideoUpdateRequestStream) ReduceFloat32(fn func(float32, VideoUpdateRequest, int) float32) []float32 {
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
func (self *VideoUpdateRequestStream) ReduceFloat64(fn func(float64, VideoUpdateRequest, int) float64) []float64 {
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
func (self *VideoUpdateRequestStream) ReduceBool(fn func(bool, VideoUpdateRequest, int) bool) []bool {
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
func (self *VideoUpdateRequestStream) Reverse() *VideoUpdateRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *VideoUpdateRequestStream) Replace(fn func(VideoUpdateRequest, int) VideoUpdateRequest) *VideoUpdateRequestStream {
	return self.ForEach(func(v VideoUpdateRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *VideoUpdateRequestStream) Select(fn func(VideoUpdateRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *VideoUpdateRequestStream) Set(index int, val VideoUpdateRequest) *VideoUpdateRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *VideoUpdateRequestStream) Skip(skip int) *VideoUpdateRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *VideoUpdateRequestStream) SkippingEach(fn func(VideoUpdateRequest, int) int) *VideoUpdateRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *VideoUpdateRequestStream) Slice(startIndex, n int) *VideoUpdateRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []VideoUpdateRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *VideoUpdateRequestStream) Sort(fn func(i, j int) bool) *VideoUpdateRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *VideoUpdateRequestStream) Tail() *VideoUpdateRequest {
	return self.Last()
}
func (self *VideoUpdateRequestStream) TailOr(arg VideoUpdateRequest) VideoUpdateRequest {
	return self.LastOr(arg)
}
func (self *VideoUpdateRequestStream) ToList() []VideoUpdateRequest {
	return self.Val()
}
func (self *VideoUpdateRequestStream) Unique() *VideoUpdateRequestStream {
	return self.Distinct()
}
func (self *VideoUpdateRequestStream) Val() []VideoUpdateRequest {
	if self == nil {
		return []VideoUpdateRequest{}
	}
	return *self.Copy()
}
func (self *VideoUpdateRequestStream) While(fn func(VideoUpdateRequest, int) bool) *VideoUpdateRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *VideoUpdateRequestStream) Where(fn func(VideoUpdateRequest) bool) *VideoUpdateRequestStream {
	result := VideoUpdateRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *VideoUpdateRequestStream) WhereSlim(fn func(VideoUpdateRequest) bool) *VideoUpdateRequestStream {
	result := VideoUpdateRequestStreamOf()
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

/*
 * Collection utility of DevicePolicyRequest Struct
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

type DevicePolicyRequestStream []DevicePolicyRequest

func DevicePolicyRequestStreamOf(arg ...DevicePolicyRequest) DevicePolicyRequestStream {
	return arg
}
func DevicePolicyRequestStreamFrom(arg []DevicePolicyRequest) DevicePolicyRequestStream {
	return arg
}
func CreateDevicePolicyRequestStream(arg ...DevicePolicyRequest) *DevicePolicyRequestStream {
	tmp := DevicePolicyRequestStreamOf(arg...)
	return &tmp
}
func GenerateDevicePolicyRequestStream(arg []DevicePolicyRequest) *DevicePolicyRequestStream {
	tmp := DevicePolicyRequestStreamFrom(arg)
	return &tmp
}

func (self *DevicePolicyRequestStream) Add(arg DevicePolicyRequest) *DevicePolicyRequestStream {
	return self.AddAll(arg)
}
func (self *DevicePolicyRequestStream) AddAll(arg ...DevicePolicyRequest) *DevicePolicyRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *DevicePolicyRequestStream) AddSafe(arg *DevicePolicyRequest) *DevicePolicyRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *DevicePolicyRequestStream) Aggregate(fn func(DevicePolicyRequest, DevicePolicyRequest) DevicePolicyRequest) *DevicePolicyRequestStream {
	result := DevicePolicyRequestStreamOf()
	self.ForEach(func(v DevicePolicyRequest, i int) {
		if i == 0 {
			result.Add(fn(DevicePolicyRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *DevicePolicyRequestStream) AllMatch(fn func(DevicePolicyRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *DevicePolicyRequestStream) AnyMatch(fn func(DevicePolicyRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *DevicePolicyRequestStream) Clone() *DevicePolicyRequestStream {
	temp := make([]DevicePolicyRequest, self.Len())
	copy(temp, *self)
	return (*DevicePolicyRequestStream)(&temp)
}
func (self *DevicePolicyRequestStream) Copy() *DevicePolicyRequestStream {
	return self.Clone()
}
func (self *DevicePolicyRequestStream) Concat(arg []DevicePolicyRequest) *DevicePolicyRequestStream {
	return self.AddAll(arg...)
}
func (self *DevicePolicyRequestStream) Contains(arg DevicePolicyRequest) bool {
	return self.FindIndex(func(_arg DevicePolicyRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *DevicePolicyRequestStream) Clean() *DevicePolicyRequestStream {
	*self = DevicePolicyRequestStreamOf()
	return self
}
func (self *DevicePolicyRequestStream) Delete(index int) *DevicePolicyRequestStream {
	return self.DeleteRange(index, index)
}
func (self *DevicePolicyRequestStream) DeleteRange(startIndex, endIndex int) *DevicePolicyRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *DevicePolicyRequestStream) Distinct() *DevicePolicyRequestStream {
	caches := map[string]bool{}
	result := DevicePolicyRequestStreamOf()
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
func (self *DevicePolicyRequestStream) Each(fn func(DevicePolicyRequest)) *DevicePolicyRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *DevicePolicyRequestStream) EachRight(fn func(DevicePolicyRequest)) *DevicePolicyRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *DevicePolicyRequestStream) Equals(arr []DevicePolicyRequest) bool {
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
func (self *DevicePolicyRequestStream) Filter(fn func(DevicePolicyRequest, int) bool) *DevicePolicyRequestStream {
	result := DevicePolicyRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *DevicePolicyRequestStream) FilterSlim(fn func(DevicePolicyRequest, int) bool) *DevicePolicyRequestStream {
	result := DevicePolicyRequestStreamOf()
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
func (self *DevicePolicyRequestStream) Find(fn func(DevicePolicyRequest, int) bool) *DevicePolicyRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *DevicePolicyRequestStream) FindOr(fn func(DevicePolicyRequest, int) bool, or DevicePolicyRequest) DevicePolicyRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *DevicePolicyRequestStream) FindIndex(fn func(DevicePolicyRequest, int) bool) int {
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
func (self *DevicePolicyRequestStream) First() *DevicePolicyRequest {
	return self.Get(0)
}
func (self *DevicePolicyRequestStream) FirstOr(arg DevicePolicyRequest) DevicePolicyRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *DevicePolicyRequestStream) ForEach(fn func(DevicePolicyRequest, int)) *DevicePolicyRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *DevicePolicyRequestStream) ForEachRight(fn func(DevicePolicyRequest, int)) *DevicePolicyRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *DevicePolicyRequestStream) GroupBy(fn func(DevicePolicyRequest, int) string) map[string][]DevicePolicyRequest {
	m := map[string][]DevicePolicyRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *DevicePolicyRequestStream) GroupByValues(fn func(DevicePolicyRequest, int) string) [][]DevicePolicyRequest {
	var tmp [][]DevicePolicyRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *DevicePolicyRequestStream) IndexOf(arg DevicePolicyRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *DevicePolicyRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *DevicePolicyRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *DevicePolicyRequestStream) Last() *DevicePolicyRequest {
	return self.Get(self.Len() - 1)
}
func (self *DevicePolicyRequestStream) LastOr(arg DevicePolicyRequest) DevicePolicyRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *DevicePolicyRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *DevicePolicyRequestStream) Limit(limit int) *DevicePolicyRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *DevicePolicyRequestStream) Map(fn func(DevicePolicyRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DevicePolicyRequestStream) Map2Int(fn func(DevicePolicyRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DevicePolicyRequestStream) Map2Int32(fn func(DevicePolicyRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DevicePolicyRequestStream) Map2Int64(fn func(DevicePolicyRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DevicePolicyRequestStream) Map2Float32(fn func(DevicePolicyRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DevicePolicyRequestStream) Map2Float64(fn func(DevicePolicyRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DevicePolicyRequestStream) Map2Bool(fn func(DevicePolicyRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DevicePolicyRequestStream) Map2Bytes(fn func(DevicePolicyRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DevicePolicyRequestStream) Map2String(fn func(DevicePolicyRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DevicePolicyRequestStream) Max(fn func(DevicePolicyRequest, int) float64) *DevicePolicyRequest {
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
func (self *DevicePolicyRequestStream) Min(fn func(DevicePolicyRequest, int) float64) *DevicePolicyRequest {
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
func (self *DevicePolicyRequestStream) NoneMatch(fn func(DevicePolicyRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *DevicePolicyRequestStream) Get(index int) *DevicePolicyRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *DevicePolicyRequestStream) GetOr(index int, arg DevicePolicyRequest) DevicePolicyRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *DevicePolicyRequestStream) Peek(fn func(*DevicePolicyRequest, int)) *DevicePolicyRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *DevicePolicyRequestStream) Reduce(fn func(DevicePolicyRequest, DevicePolicyRequest, int) DevicePolicyRequest) *DevicePolicyRequestStream {
	return self.ReduceInit(fn, DevicePolicyRequest{})
}
func (self *DevicePolicyRequestStream) ReduceInit(fn func(DevicePolicyRequest, DevicePolicyRequest, int) DevicePolicyRequest, initialValue DevicePolicyRequest) *DevicePolicyRequestStream {
	result := DevicePolicyRequestStreamOf()
	self.ForEach(func(v DevicePolicyRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *DevicePolicyRequestStream) ReduceInterface(fn func(interface{}, DevicePolicyRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(DevicePolicyRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *DevicePolicyRequestStream) ReduceString(fn func(string, DevicePolicyRequest, int) string) []string {
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
func (self *DevicePolicyRequestStream) ReduceInt(fn func(int, DevicePolicyRequest, int) int) []int {
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
func (self *DevicePolicyRequestStream) ReduceInt32(fn func(int32, DevicePolicyRequest, int) int32) []int32 {
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
func (self *DevicePolicyRequestStream) ReduceInt64(fn func(int64, DevicePolicyRequest, int) int64) []int64 {
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
func (self *DevicePolicyRequestStream) ReduceFloat32(fn func(float32, DevicePolicyRequest, int) float32) []float32 {
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
func (self *DevicePolicyRequestStream) ReduceFloat64(fn func(float64, DevicePolicyRequest, int) float64) []float64 {
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
func (self *DevicePolicyRequestStream) ReduceBool(fn func(bool, DevicePolicyRequest, int) bool) []bool {
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
func (self *DevicePolicyRequestStream) Reverse() *DevicePolicyRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *DevicePolicyRequestStream) Replace(fn func(DevicePolicyRequest, int) DevicePolicyRequest) *DevicePolicyRequestStream {
	return self.ForEach(func(v DevicePolicyRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *DevicePolicyRequestStream) Select(fn func(DevicePolicyRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *DevicePolicyRequestStream) Set(index int, val DevicePolicyRequest) *DevicePolicyRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *DevicePolicyRequestStream) Skip(skip int) *DevicePolicyRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *DevicePolicyRequestStream) SkippingEach(fn func(DevicePolicyRequest, int) int) *DevicePolicyRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *DevicePolicyRequestStream) Slice(startIndex, n int) *DevicePolicyRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []DevicePolicyRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *DevicePolicyRequestStream) Sort(fn func(i, j int) bool) *DevicePolicyRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *DevicePolicyRequestStream) Tail() *DevicePolicyRequest {
	return self.Last()
}
func (self *DevicePolicyRequestStream) TailOr(arg DevicePolicyRequest) DevicePolicyRequest {
	return self.LastOr(arg)
}
func (self *DevicePolicyRequestStream) ToList() []DevicePolicyRequest {
	return self.Val()
}
func (self *DevicePolicyRequestStream) Unique() *DevicePolicyRequestStream {
	return self.Distinct()
}
func (self *DevicePolicyRequestStream) Val() []DevicePolicyRequest {
	if self == nil {
		return []DevicePolicyRequest{}
	}
	return *self.Copy()
}
func (self *DevicePolicyRequestStream) While(fn func(DevicePolicyRequest, int) bool) *DevicePolicyRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *DevicePolicyRequestStream) Where(fn func(DevicePolicyRequest) bool) *DevicePolicyRequestStream {
	result := DevicePolicyRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *DevicePolicyRequestStream) WhereSlim(fn func(DevicePolicyRequest) bool) *DevicePolicyRequestStream {
	result := DevicePolicyRequestStreamOf()
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

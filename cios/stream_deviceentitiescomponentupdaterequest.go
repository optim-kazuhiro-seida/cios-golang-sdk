/*
 * Collection utility of DeviceEntitiesComponentUpdateRequest Struct
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

type DeviceEntitiesComponentUpdateRequestStream []DeviceEntitiesComponentUpdateRequest

func DeviceEntitiesComponentUpdateRequestStreamOf(arg ...DeviceEntitiesComponentUpdateRequest) DeviceEntitiesComponentUpdateRequestStream {
	return arg
}
func DeviceEntitiesComponentUpdateRequestStreamFrom(arg []DeviceEntitiesComponentUpdateRequest) DeviceEntitiesComponentUpdateRequestStream {
	return arg
}
func CreateDeviceEntitiesComponentUpdateRequestStream(arg ...DeviceEntitiesComponentUpdateRequest) *DeviceEntitiesComponentUpdateRequestStream {
	tmp := DeviceEntitiesComponentUpdateRequestStreamOf(arg...)
	return &tmp
}
func GenerateDeviceEntitiesComponentUpdateRequestStream(arg []DeviceEntitiesComponentUpdateRequest) *DeviceEntitiesComponentUpdateRequestStream {
	tmp := DeviceEntitiesComponentUpdateRequestStreamFrom(arg)
	return &tmp
}

func (self *DeviceEntitiesComponentUpdateRequestStream) Add(arg DeviceEntitiesComponentUpdateRequest) *DeviceEntitiesComponentUpdateRequestStream {
	return self.AddAll(arg)
}
func (self *DeviceEntitiesComponentUpdateRequestStream) AddAll(arg ...DeviceEntitiesComponentUpdateRequest) *DeviceEntitiesComponentUpdateRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *DeviceEntitiesComponentUpdateRequestStream) AddSafe(arg *DeviceEntitiesComponentUpdateRequest) *DeviceEntitiesComponentUpdateRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Aggregate(fn func(DeviceEntitiesComponentUpdateRequest, DeviceEntitiesComponentUpdateRequest) DeviceEntitiesComponentUpdateRequest) *DeviceEntitiesComponentUpdateRequestStream {
	result := DeviceEntitiesComponentUpdateRequestStreamOf()
	self.ForEach(func(v DeviceEntitiesComponentUpdateRequest, i int) {
		if i == 0 {
			result.Add(fn(DeviceEntitiesComponentUpdateRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *DeviceEntitiesComponentUpdateRequestStream) AllMatch(fn func(DeviceEntitiesComponentUpdateRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *DeviceEntitiesComponentUpdateRequestStream) AnyMatch(fn func(DeviceEntitiesComponentUpdateRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Clone() *DeviceEntitiesComponentUpdateRequestStream {
	temp := make([]DeviceEntitiesComponentUpdateRequest, self.Len())
	copy(temp, *self)
	return (*DeviceEntitiesComponentUpdateRequestStream)(&temp)
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Copy() *DeviceEntitiesComponentUpdateRequestStream {
	return self.Clone()
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Concat(arg []DeviceEntitiesComponentUpdateRequest) *DeviceEntitiesComponentUpdateRequestStream {
	return self.AddAll(arg...)
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Contains(arg DeviceEntitiesComponentUpdateRequest) bool {
	return self.FindIndex(func(_arg DeviceEntitiesComponentUpdateRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Clean() *DeviceEntitiesComponentUpdateRequestStream {
	*self = DeviceEntitiesComponentUpdateRequestStreamOf()
	return self
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Delete(index int) *DeviceEntitiesComponentUpdateRequestStream {
	return self.DeleteRange(index, index)
}
func (self *DeviceEntitiesComponentUpdateRequestStream) DeleteRange(startIndex, endIndex int) *DeviceEntitiesComponentUpdateRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Distinct() *DeviceEntitiesComponentUpdateRequestStream {
	caches := map[string]bool{}
	result := DeviceEntitiesComponentUpdateRequestStreamOf()
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
func (self *DeviceEntitiesComponentUpdateRequestStream) Each(fn func(DeviceEntitiesComponentUpdateRequest)) *DeviceEntitiesComponentUpdateRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *DeviceEntitiesComponentUpdateRequestStream) EachRight(fn func(DeviceEntitiesComponentUpdateRequest)) *DeviceEntitiesComponentUpdateRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Equals(arr []DeviceEntitiesComponentUpdateRequest) bool {
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
func (self *DeviceEntitiesComponentUpdateRequestStream) Filter(fn func(DeviceEntitiesComponentUpdateRequest, int) bool) *DeviceEntitiesComponentUpdateRequestStream {
	result := DeviceEntitiesComponentUpdateRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *DeviceEntitiesComponentUpdateRequestStream) FilterSlim(fn func(DeviceEntitiesComponentUpdateRequest, int) bool) *DeviceEntitiesComponentUpdateRequestStream {
	result := DeviceEntitiesComponentUpdateRequestStreamOf()
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
func (self *DeviceEntitiesComponentUpdateRequestStream) Find(fn func(DeviceEntitiesComponentUpdateRequest, int) bool) *DeviceEntitiesComponentUpdateRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *DeviceEntitiesComponentUpdateRequestStream) FindOr(fn func(DeviceEntitiesComponentUpdateRequest, int) bool, or DeviceEntitiesComponentUpdateRequest) DeviceEntitiesComponentUpdateRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *DeviceEntitiesComponentUpdateRequestStream) FindIndex(fn func(DeviceEntitiesComponentUpdateRequest, int) bool) int {
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
func (self *DeviceEntitiesComponentUpdateRequestStream) First() *DeviceEntitiesComponentUpdateRequest {
	return self.Get(0)
}
func (self *DeviceEntitiesComponentUpdateRequestStream) FirstOr(arg DeviceEntitiesComponentUpdateRequest) DeviceEntitiesComponentUpdateRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *DeviceEntitiesComponentUpdateRequestStream) ForEach(fn func(DeviceEntitiesComponentUpdateRequest, int)) *DeviceEntitiesComponentUpdateRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *DeviceEntitiesComponentUpdateRequestStream) ForEachRight(fn func(DeviceEntitiesComponentUpdateRequest, int)) *DeviceEntitiesComponentUpdateRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *DeviceEntitiesComponentUpdateRequestStream) GroupBy(fn func(DeviceEntitiesComponentUpdateRequest, int) string) map[string][]DeviceEntitiesComponentUpdateRequest {
	m := map[string][]DeviceEntitiesComponentUpdateRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *DeviceEntitiesComponentUpdateRequestStream) GroupByValues(fn func(DeviceEntitiesComponentUpdateRequest, int) string) [][]DeviceEntitiesComponentUpdateRequest {
	var tmp [][]DeviceEntitiesComponentUpdateRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *DeviceEntitiesComponentUpdateRequestStream) IndexOf(arg DeviceEntitiesComponentUpdateRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *DeviceEntitiesComponentUpdateRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *DeviceEntitiesComponentUpdateRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Last() *DeviceEntitiesComponentUpdateRequest {
	return self.Get(self.Len() - 1)
}
func (self *DeviceEntitiesComponentUpdateRequestStream) LastOr(arg DeviceEntitiesComponentUpdateRequest) DeviceEntitiesComponentUpdateRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Limit(limit int) *DeviceEntitiesComponentUpdateRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *DeviceEntitiesComponentUpdateRequestStream) Map(fn func(DeviceEntitiesComponentUpdateRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Map2Int(fn func(DeviceEntitiesComponentUpdateRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Map2Int32(fn func(DeviceEntitiesComponentUpdateRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Map2Int64(fn func(DeviceEntitiesComponentUpdateRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Map2Float32(fn func(DeviceEntitiesComponentUpdateRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Map2Float64(fn func(DeviceEntitiesComponentUpdateRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Map2Bool(fn func(DeviceEntitiesComponentUpdateRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Map2Bytes(fn func(DeviceEntitiesComponentUpdateRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Map2String(fn func(DeviceEntitiesComponentUpdateRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Max(fn func(DeviceEntitiesComponentUpdateRequest, int) float64) *DeviceEntitiesComponentUpdateRequest {
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
func (self *DeviceEntitiesComponentUpdateRequestStream) Min(fn func(DeviceEntitiesComponentUpdateRequest, int) float64) *DeviceEntitiesComponentUpdateRequest {
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
func (self *DeviceEntitiesComponentUpdateRequestStream) NoneMatch(fn func(DeviceEntitiesComponentUpdateRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Get(index int) *DeviceEntitiesComponentUpdateRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *DeviceEntitiesComponentUpdateRequestStream) GetOr(index int, arg DeviceEntitiesComponentUpdateRequest) DeviceEntitiesComponentUpdateRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Peek(fn func(*DeviceEntitiesComponentUpdateRequest, int)) *DeviceEntitiesComponentUpdateRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *DeviceEntitiesComponentUpdateRequestStream) Reduce(fn func(DeviceEntitiesComponentUpdateRequest, DeviceEntitiesComponentUpdateRequest, int) DeviceEntitiesComponentUpdateRequest) *DeviceEntitiesComponentUpdateRequestStream {
	return self.ReduceInit(fn, DeviceEntitiesComponentUpdateRequest{})
}
func (self *DeviceEntitiesComponentUpdateRequestStream) ReduceInit(fn func(DeviceEntitiesComponentUpdateRequest, DeviceEntitiesComponentUpdateRequest, int) DeviceEntitiesComponentUpdateRequest, initialValue DeviceEntitiesComponentUpdateRequest) *DeviceEntitiesComponentUpdateRequestStream {
	result := DeviceEntitiesComponentUpdateRequestStreamOf()
	self.ForEach(func(v DeviceEntitiesComponentUpdateRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *DeviceEntitiesComponentUpdateRequestStream) ReduceInterface(fn func(interface{}, DeviceEntitiesComponentUpdateRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(DeviceEntitiesComponentUpdateRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *DeviceEntitiesComponentUpdateRequestStream) ReduceString(fn func(string, DeviceEntitiesComponentUpdateRequest, int) string) []string {
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
func (self *DeviceEntitiesComponentUpdateRequestStream) ReduceInt(fn func(int, DeviceEntitiesComponentUpdateRequest, int) int) []int {
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
func (self *DeviceEntitiesComponentUpdateRequestStream) ReduceInt32(fn func(int32, DeviceEntitiesComponentUpdateRequest, int) int32) []int32 {
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
func (self *DeviceEntitiesComponentUpdateRequestStream) ReduceInt64(fn func(int64, DeviceEntitiesComponentUpdateRequest, int) int64) []int64 {
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
func (self *DeviceEntitiesComponentUpdateRequestStream) ReduceFloat32(fn func(float32, DeviceEntitiesComponentUpdateRequest, int) float32) []float32 {
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
func (self *DeviceEntitiesComponentUpdateRequestStream) ReduceFloat64(fn func(float64, DeviceEntitiesComponentUpdateRequest, int) float64) []float64 {
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
func (self *DeviceEntitiesComponentUpdateRequestStream) ReduceBool(fn func(bool, DeviceEntitiesComponentUpdateRequest, int) bool) []bool {
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
func (self *DeviceEntitiesComponentUpdateRequestStream) Reverse() *DeviceEntitiesComponentUpdateRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Replace(fn func(DeviceEntitiesComponentUpdateRequest, int) DeviceEntitiesComponentUpdateRequest) *DeviceEntitiesComponentUpdateRequestStream {
	return self.ForEach(func(v DeviceEntitiesComponentUpdateRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Select(fn func(DeviceEntitiesComponentUpdateRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Set(index int, val DeviceEntitiesComponentUpdateRequest) *DeviceEntitiesComponentUpdateRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Skip(skip int) *DeviceEntitiesComponentUpdateRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *DeviceEntitiesComponentUpdateRequestStream) SkippingEach(fn func(DeviceEntitiesComponentUpdateRequest, int) int) *DeviceEntitiesComponentUpdateRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Slice(startIndex, n int) *DeviceEntitiesComponentUpdateRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []DeviceEntitiesComponentUpdateRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Sort(fn func(i, j int) bool) *DeviceEntitiesComponentUpdateRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *DeviceEntitiesComponentUpdateRequestStream) Tail() *DeviceEntitiesComponentUpdateRequest {
	return self.Last()
}
func (self *DeviceEntitiesComponentUpdateRequestStream) TailOr(arg DeviceEntitiesComponentUpdateRequest) DeviceEntitiesComponentUpdateRequest {
	return self.LastOr(arg)
}
func (self *DeviceEntitiesComponentUpdateRequestStream) ToList() []DeviceEntitiesComponentUpdateRequest {
	return self.Val()
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Unique() *DeviceEntitiesComponentUpdateRequestStream {
	return self.Distinct()
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Val() []DeviceEntitiesComponentUpdateRequest {
	if self == nil {
		return []DeviceEntitiesComponentUpdateRequest{}
	}
	return *self.Copy()
}
func (self *DeviceEntitiesComponentUpdateRequestStream) While(fn func(DeviceEntitiesComponentUpdateRequest, int) bool) *DeviceEntitiesComponentUpdateRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *DeviceEntitiesComponentUpdateRequestStream) Where(fn func(DeviceEntitiesComponentUpdateRequest) bool) *DeviceEntitiesComponentUpdateRequestStream {
	result := DeviceEntitiesComponentUpdateRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *DeviceEntitiesComponentUpdateRequestStream) WhereSlim(fn func(DeviceEntitiesComponentUpdateRequest) bool) *DeviceEntitiesComponentUpdateRequestStream {
	result := DeviceEntitiesComponentUpdateRequestStreamOf()
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

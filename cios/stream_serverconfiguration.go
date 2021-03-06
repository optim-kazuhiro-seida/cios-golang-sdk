/*
 * Collection utility of ServerConfiguration Struct
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

type ServerConfigurationStream []ServerConfiguration

func ServerConfigurationStreamOf(arg ...ServerConfiguration) ServerConfigurationStream {
	return arg
}
func ServerConfigurationStreamFrom(arg []ServerConfiguration) ServerConfigurationStream {
	return arg
}
func CreateServerConfigurationStream(arg ...ServerConfiguration) *ServerConfigurationStream {
	tmp := ServerConfigurationStreamOf(arg...)
	return &tmp
}
func GenerateServerConfigurationStream(arg []ServerConfiguration) *ServerConfigurationStream {
	tmp := ServerConfigurationStreamFrom(arg)
	return &tmp
}

func (self *ServerConfigurationStream) Add(arg ServerConfiguration) *ServerConfigurationStream {
	return self.AddAll(arg)
}
func (self *ServerConfigurationStream) AddAll(arg ...ServerConfiguration) *ServerConfigurationStream {
	*self = append(*self, arg...)
	return self
}
func (self *ServerConfigurationStream) AddSafe(arg *ServerConfiguration) *ServerConfigurationStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ServerConfigurationStream) Aggregate(fn func(ServerConfiguration, ServerConfiguration) ServerConfiguration) *ServerConfigurationStream {
	result := ServerConfigurationStreamOf()
	self.ForEach(func(v ServerConfiguration, i int) {
		if i == 0 {
			result.Add(fn(ServerConfiguration{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ServerConfigurationStream) AllMatch(fn func(ServerConfiguration, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ServerConfigurationStream) AnyMatch(fn func(ServerConfiguration, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ServerConfigurationStream) Clone() *ServerConfigurationStream {
	temp := make([]ServerConfiguration, self.Len())
	copy(temp, *self)
	return (*ServerConfigurationStream)(&temp)
}
func (self *ServerConfigurationStream) Copy() *ServerConfigurationStream {
	return self.Clone()
}
func (self *ServerConfigurationStream) Concat(arg []ServerConfiguration) *ServerConfigurationStream {
	return self.AddAll(arg...)
}
func (self *ServerConfigurationStream) Contains(arg ServerConfiguration) bool {
	return self.FindIndex(func(_arg ServerConfiguration, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ServerConfigurationStream) Clean() *ServerConfigurationStream {
	*self = ServerConfigurationStreamOf()
	return self
}
func (self *ServerConfigurationStream) Delete(index int) *ServerConfigurationStream {
	return self.DeleteRange(index, index)
}
func (self *ServerConfigurationStream) DeleteRange(startIndex, endIndex int) *ServerConfigurationStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ServerConfigurationStream) Distinct() *ServerConfigurationStream {
	caches := map[string]bool{}
	result := ServerConfigurationStreamOf()
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
func (self *ServerConfigurationStream) Each(fn func(ServerConfiguration)) *ServerConfigurationStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ServerConfigurationStream) EachRight(fn func(ServerConfiguration)) *ServerConfigurationStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ServerConfigurationStream) Equals(arr []ServerConfiguration) bool {
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
func (self *ServerConfigurationStream) Filter(fn func(ServerConfiguration, int) bool) *ServerConfigurationStream {
	result := ServerConfigurationStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ServerConfigurationStream) FilterSlim(fn func(ServerConfiguration, int) bool) *ServerConfigurationStream {
	result := ServerConfigurationStreamOf()
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
func (self *ServerConfigurationStream) Find(fn func(ServerConfiguration, int) bool) *ServerConfiguration {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ServerConfigurationStream) FindOr(fn func(ServerConfiguration, int) bool, or ServerConfiguration) ServerConfiguration {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ServerConfigurationStream) FindIndex(fn func(ServerConfiguration, int) bool) int {
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
func (self *ServerConfigurationStream) First() *ServerConfiguration {
	return self.Get(0)
}
func (self *ServerConfigurationStream) FirstOr(arg ServerConfiguration) ServerConfiguration {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ServerConfigurationStream) ForEach(fn func(ServerConfiguration, int)) *ServerConfigurationStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ServerConfigurationStream) ForEachRight(fn func(ServerConfiguration, int)) *ServerConfigurationStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ServerConfigurationStream) GroupBy(fn func(ServerConfiguration, int) string) map[string][]ServerConfiguration {
	m := map[string][]ServerConfiguration{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ServerConfigurationStream) GroupByValues(fn func(ServerConfiguration, int) string) [][]ServerConfiguration {
	var tmp [][]ServerConfiguration
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ServerConfigurationStream) IndexOf(arg ServerConfiguration) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ServerConfigurationStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ServerConfigurationStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ServerConfigurationStream) Last() *ServerConfiguration {
	return self.Get(self.Len() - 1)
}
func (self *ServerConfigurationStream) LastOr(arg ServerConfiguration) ServerConfiguration {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ServerConfigurationStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ServerConfigurationStream) Limit(limit int) *ServerConfigurationStream {
	self.Slice(0, limit)
	return self
}

func (self *ServerConfigurationStream) Map(fn func(ServerConfiguration, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ServerConfigurationStream) Map2Int(fn func(ServerConfiguration, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ServerConfigurationStream) Map2Int32(fn func(ServerConfiguration, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ServerConfigurationStream) Map2Int64(fn func(ServerConfiguration, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ServerConfigurationStream) Map2Float32(fn func(ServerConfiguration, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ServerConfigurationStream) Map2Float64(fn func(ServerConfiguration, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ServerConfigurationStream) Map2Bool(fn func(ServerConfiguration, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ServerConfigurationStream) Map2Bytes(fn func(ServerConfiguration, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ServerConfigurationStream) Map2String(fn func(ServerConfiguration, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ServerConfigurationStream) Max(fn func(ServerConfiguration, int) float64) *ServerConfiguration {
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
func (self *ServerConfigurationStream) Min(fn func(ServerConfiguration, int) float64) *ServerConfiguration {
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
func (self *ServerConfigurationStream) NoneMatch(fn func(ServerConfiguration, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ServerConfigurationStream) Get(index int) *ServerConfiguration {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ServerConfigurationStream) GetOr(index int, arg ServerConfiguration) ServerConfiguration {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ServerConfigurationStream) Peek(fn func(*ServerConfiguration, int)) *ServerConfigurationStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ServerConfigurationStream) Reduce(fn func(ServerConfiguration, ServerConfiguration, int) ServerConfiguration) *ServerConfigurationStream {
	return self.ReduceInit(fn, ServerConfiguration{})
}
func (self *ServerConfigurationStream) ReduceInit(fn func(ServerConfiguration, ServerConfiguration, int) ServerConfiguration, initialValue ServerConfiguration) *ServerConfigurationStream {
	result := ServerConfigurationStreamOf()
	self.ForEach(func(v ServerConfiguration, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ServerConfigurationStream) ReduceInterface(fn func(interface{}, ServerConfiguration, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ServerConfiguration{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ServerConfigurationStream) ReduceString(fn func(string, ServerConfiguration, int) string) []string {
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
func (self *ServerConfigurationStream) ReduceInt(fn func(int, ServerConfiguration, int) int) []int {
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
func (self *ServerConfigurationStream) ReduceInt32(fn func(int32, ServerConfiguration, int) int32) []int32 {
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
func (self *ServerConfigurationStream) ReduceInt64(fn func(int64, ServerConfiguration, int) int64) []int64 {
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
func (self *ServerConfigurationStream) ReduceFloat32(fn func(float32, ServerConfiguration, int) float32) []float32 {
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
func (self *ServerConfigurationStream) ReduceFloat64(fn func(float64, ServerConfiguration, int) float64) []float64 {
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
func (self *ServerConfigurationStream) ReduceBool(fn func(bool, ServerConfiguration, int) bool) []bool {
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
func (self *ServerConfigurationStream) Reverse() *ServerConfigurationStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ServerConfigurationStream) Replace(fn func(ServerConfiguration, int) ServerConfiguration) *ServerConfigurationStream {
	return self.ForEach(func(v ServerConfiguration, i int) { self.Set(i, fn(v, i)) })
}
func (self *ServerConfigurationStream) Select(fn func(ServerConfiguration) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ServerConfigurationStream) Set(index int, val ServerConfiguration) *ServerConfigurationStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ServerConfigurationStream) Skip(skip int) *ServerConfigurationStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ServerConfigurationStream) SkippingEach(fn func(ServerConfiguration, int) int) *ServerConfigurationStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ServerConfigurationStream) Slice(startIndex, n int) *ServerConfigurationStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ServerConfiguration{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ServerConfigurationStream) Sort(fn func(i, j int) bool) *ServerConfigurationStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ServerConfigurationStream) Tail() *ServerConfiguration {
	return self.Last()
}
func (self *ServerConfigurationStream) TailOr(arg ServerConfiguration) ServerConfiguration {
	return self.LastOr(arg)
}
func (self *ServerConfigurationStream) ToList() []ServerConfiguration {
	return self.Val()
}
func (self *ServerConfigurationStream) Unique() *ServerConfigurationStream {
	return self.Distinct()
}
func (self *ServerConfigurationStream) Val() []ServerConfiguration {
	if self == nil {
		return []ServerConfiguration{}
	}
	return *self.Copy()
}
func (self *ServerConfigurationStream) While(fn func(ServerConfiguration, int) bool) *ServerConfigurationStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ServerConfigurationStream) Where(fn func(ServerConfiguration) bool) *ServerConfigurationStream {
	result := ServerConfigurationStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ServerConfigurationStream) WhereSlim(fn func(ServerConfiguration) bool) *ServerConfigurationStream {
	result := ServerConfigurationStreamOf()
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

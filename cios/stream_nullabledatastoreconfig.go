/*
 * Collection utility of NullableDataStoreConfig Struct
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

type NullableDataStoreConfigStream []NullableDataStoreConfig

func NullableDataStoreConfigStreamOf(arg ...NullableDataStoreConfig) NullableDataStoreConfigStream {
	return arg
}
func NullableDataStoreConfigStreamFrom(arg []NullableDataStoreConfig) NullableDataStoreConfigStream {
	return arg
}
func CreateNullableDataStoreConfigStream(arg ...NullableDataStoreConfig) *NullableDataStoreConfigStream {
	tmp := NullableDataStoreConfigStreamOf(arg...)
	return &tmp
}
func GenerateNullableDataStoreConfigStream(arg []NullableDataStoreConfig) *NullableDataStoreConfigStream {
	tmp := NullableDataStoreConfigStreamFrom(arg)
	return &tmp
}

func (self *NullableDataStoreConfigStream) Add(arg NullableDataStoreConfig) *NullableDataStoreConfigStream {
	return self.AddAll(arg)
}
func (self *NullableDataStoreConfigStream) AddAll(arg ...NullableDataStoreConfig) *NullableDataStoreConfigStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableDataStoreConfigStream) AddSafe(arg *NullableDataStoreConfig) *NullableDataStoreConfigStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableDataStoreConfigStream) Aggregate(fn func(NullableDataStoreConfig, NullableDataStoreConfig) NullableDataStoreConfig) *NullableDataStoreConfigStream {
	result := NullableDataStoreConfigStreamOf()
	self.ForEach(func(v NullableDataStoreConfig, i int) {
		if i == 0 {
			result.Add(fn(NullableDataStoreConfig{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableDataStoreConfigStream) AllMatch(fn func(NullableDataStoreConfig, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableDataStoreConfigStream) AnyMatch(fn func(NullableDataStoreConfig, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableDataStoreConfigStream) Clone() *NullableDataStoreConfigStream {
	temp := make([]NullableDataStoreConfig, self.Len())
	copy(temp, *self)
	return (*NullableDataStoreConfigStream)(&temp)
}
func (self *NullableDataStoreConfigStream) Copy() *NullableDataStoreConfigStream {
	return self.Clone()
}
func (self *NullableDataStoreConfigStream) Concat(arg []NullableDataStoreConfig) *NullableDataStoreConfigStream {
	return self.AddAll(arg...)
}
func (self *NullableDataStoreConfigStream) Contains(arg NullableDataStoreConfig) bool {
	return self.FindIndex(func(_arg NullableDataStoreConfig, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableDataStoreConfigStream) Clean() *NullableDataStoreConfigStream {
	*self = NullableDataStoreConfigStreamOf()
	return self
}
func (self *NullableDataStoreConfigStream) Delete(index int) *NullableDataStoreConfigStream {
	return self.DeleteRange(index, index)
}
func (self *NullableDataStoreConfigStream) DeleteRange(startIndex, endIndex int) *NullableDataStoreConfigStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableDataStoreConfigStream) Distinct() *NullableDataStoreConfigStream {
	caches := map[string]bool{}
	result := NullableDataStoreConfigStreamOf()
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
func (self *NullableDataStoreConfigStream) Each(fn func(NullableDataStoreConfig)) *NullableDataStoreConfigStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableDataStoreConfigStream) EachRight(fn func(NullableDataStoreConfig)) *NullableDataStoreConfigStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableDataStoreConfigStream) Equals(arr []NullableDataStoreConfig) bool {
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
func (self *NullableDataStoreConfigStream) Filter(fn func(NullableDataStoreConfig, int) bool) *NullableDataStoreConfigStream {
	result := NullableDataStoreConfigStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableDataStoreConfigStream) FilterSlim(fn func(NullableDataStoreConfig, int) bool) *NullableDataStoreConfigStream {
	result := NullableDataStoreConfigStreamOf()
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
func (self *NullableDataStoreConfigStream) Find(fn func(NullableDataStoreConfig, int) bool) *NullableDataStoreConfig {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableDataStoreConfigStream) FindOr(fn func(NullableDataStoreConfig, int) bool, or NullableDataStoreConfig) NullableDataStoreConfig {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableDataStoreConfigStream) FindIndex(fn func(NullableDataStoreConfig, int) bool) int {
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
func (self *NullableDataStoreConfigStream) First() *NullableDataStoreConfig {
	return self.Get(0)
}
func (self *NullableDataStoreConfigStream) FirstOr(arg NullableDataStoreConfig) NullableDataStoreConfig {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableDataStoreConfigStream) ForEach(fn func(NullableDataStoreConfig, int)) *NullableDataStoreConfigStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableDataStoreConfigStream) ForEachRight(fn func(NullableDataStoreConfig, int)) *NullableDataStoreConfigStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableDataStoreConfigStream) GroupBy(fn func(NullableDataStoreConfig, int) string) map[string][]NullableDataStoreConfig {
	m := map[string][]NullableDataStoreConfig{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableDataStoreConfigStream) GroupByValues(fn func(NullableDataStoreConfig, int) string) [][]NullableDataStoreConfig {
	var tmp [][]NullableDataStoreConfig
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableDataStoreConfigStream) IndexOf(arg NullableDataStoreConfig) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableDataStoreConfigStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableDataStoreConfigStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableDataStoreConfigStream) Last() *NullableDataStoreConfig {
	return self.Get(self.Len() - 1)
}
func (self *NullableDataStoreConfigStream) LastOr(arg NullableDataStoreConfig) NullableDataStoreConfig {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableDataStoreConfigStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableDataStoreConfigStream) Limit(limit int) *NullableDataStoreConfigStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableDataStoreConfigStream) Map(fn func(NullableDataStoreConfig, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDataStoreConfigStream) Map2Int(fn func(NullableDataStoreConfig, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDataStoreConfigStream) Map2Int32(fn func(NullableDataStoreConfig, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDataStoreConfigStream) Map2Int64(fn func(NullableDataStoreConfig, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDataStoreConfigStream) Map2Float32(fn func(NullableDataStoreConfig, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDataStoreConfigStream) Map2Float64(fn func(NullableDataStoreConfig, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDataStoreConfigStream) Map2Bool(fn func(NullableDataStoreConfig, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDataStoreConfigStream) Map2Bytes(fn func(NullableDataStoreConfig, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDataStoreConfigStream) Map2String(fn func(NullableDataStoreConfig, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDataStoreConfigStream) Max(fn func(NullableDataStoreConfig, int) float64) *NullableDataStoreConfig {
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
func (self *NullableDataStoreConfigStream) Min(fn func(NullableDataStoreConfig, int) float64) *NullableDataStoreConfig {
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
func (self *NullableDataStoreConfigStream) NoneMatch(fn func(NullableDataStoreConfig, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableDataStoreConfigStream) Get(index int) *NullableDataStoreConfig {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableDataStoreConfigStream) GetOr(index int, arg NullableDataStoreConfig) NullableDataStoreConfig {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableDataStoreConfigStream) Peek(fn func(*NullableDataStoreConfig, int)) *NullableDataStoreConfigStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *NullableDataStoreConfigStream) Reduce(fn func(NullableDataStoreConfig, NullableDataStoreConfig, int) NullableDataStoreConfig) *NullableDataStoreConfigStream {
	return self.ReduceInit(fn, NullableDataStoreConfig{})
}
func (self *NullableDataStoreConfigStream) ReduceInit(fn func(NullableDataStoreConfig, NullableDataStoreConfig, int) NullableDataStoreConfig, initialValue NullableDataStoreConfig) *NullableDataStoreConfigStream {
	result := NullableDataStoreConfigStreamOf()
	self.ForEach(func(v NullableDataStoreConfig, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableDataStoreConfigStream) ReduceInterface(fn func(interface{}, NullableDataStoreConfig, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableDataStoreConfig{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableDataStoreConfigStream) ReduceString(fn func(string, NullableDataStoreConfig, int) string) []string {
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
func (self *NullableDataStoreConfigStream) ReduceInt(fn func(int, NullableDataStoreConfig, int) int) []int {
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
func (self *NullableDataStoreConfigStream) ReduceInt32(fn func(int32, NullableDataStoreConfig, int) int32) []int32 {
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
func (self *NullableDataStoreConfigStream) ReduceInt64(fn func(int64, NullableDataStoreConfig, int) int64) []int64 {
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
func (self *NullableDataStoreConfigStream) ReduceFloat32(fn func(float32, NullableDataStoreConfig, int) float32) []float32 {
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
func (self *NullableDataStoreConfigStream) ReduceFloat64(fn func(float64, NullableDataStoreConfig, int) float64) []float64 {
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
func (self *NullableDataStoreConfigStream) ReduceBool(fn func(bool, NullableDataStoreConfig, int) bool) []bool {
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
func (self *NullableDataStoreConfigStream) Reverse() *NullableDataStoreConfigStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableDataStoreConfigStream) Replace(fn func(NullableDataStoreConfig, int) NullableDataStoreConfig) *NullableDataStoreConfigStream {
	return self.ForEach(func(v NullableDataStoreConfig, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableDataStoreConfigStream) Select(fn func(NullableDataStoreConfig) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableDataStoreConfigStream) Set(index int, val NullableDataStoreConfig) *NullableDataStoreConfigStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableDataStoreConfigStream) Skip(skip int) *NullableDataStoreConfigStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableDataStoreConfigStream) SkippingEach(fn func(NullableDataStoreConfig, int) int) *NullableDataStoreConfigStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableDataStoreConfigStream) Slice(startIndex, n int) *NullableDataStoreConfigStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableDataStoreConfig{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableDataStoreConfigStream) Sort(fn func(i, j int) bool) *NullableDataStoreConfigStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableDataStoreConfigStream) Tail() *NullableDataStoreConfig {
	return self.Last()
}
func (self *NullableDataStoreConfigStream) TailOr(arg NullableDataStoreConfig) NullableDataStoreConfig {
	return self.LastOr(arg)
}
func (self *NullableDataStoreConfigStream) ToList() []NullableDataStoreConfig {
	return self.Val()
}
func (self *NullableDataStoreConfigStream) Unique() *NullableDataStoreConfigStream {
	return self.Distinct()
}
func (self *NullableDataStoreConfigStream) Val() []NullableDataStoreConfig {
	if self == nil {
		return []NullableDataStoreConfig{}
	}
	return *self.Copy()
}
func (self *NullableDataStoreConfigStream) While(fn func(NullableDataStoreConfig, int) bool) *NullableDataStoreConfigStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableDataStoreConfigStream) Where(fn func(NullableDataStoreConfig) bool) *NullableDataStoreConfigStream {
	result := NullableDataStoreConfigStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableDataStoreConfigStream) WhereSlim(fn func(NullableDataStoreConfig) bool) *NullableDataStoreConfigStream {
	result := NullableDataStoreConfigStreamOf()
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
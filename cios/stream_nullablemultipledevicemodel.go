/*
 * Collection utility of NullableMultipleDeviceModel Struct
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

type NullableMultipleDeviceModelStream []NullableMultipleDeviceModel

func NullableMultipleDeviceModelStreamOf(arg ...NullableMultipleDeviceModel) NullableMultipleDeviceModelStream {
	return arg
}
func NullableMultipleDeviceModelStreamFrom(arg []NullableMultipleDeviceModel) NullableMultipleDeviceModelStream {
	return arg
}
func CreateNullableMultipleDeviceModelStream(arg ...NullableMultipleDeviceModel) *NullableMultipleDeviceModelStream {
	tmp := NullableMultipleDeviceModelStreamOf(arg...)
	return &tmp
}
func GenerateNullableMultipleDeviceModelStream(arg []NullableMultipleDeviceModel) *NullableMultipleDeviceModelStream {
	tmp := NullableMultipleDeviceModelStreamFrom(arg)
	return &tmp
}

func (self *NullableMultipleDeviceModelStream) Add(arg NullableMultipleDeviceModel) *NullableMultipleDeviceModelStream {
	return self.AddAll(arg)
}
func (self *NullableMultipleDeviceModelStream) AddAll(arg ...NullableMultipleDeviceModel) *NullableMultipleDeviceModelStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableMultipleDeviceModelStream) AddSafe(arg *NullableMultipleDeviceModel) *NullableMultipleDeviceModelStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableMultipleDeviceModelStream) Aggregate(fn func(NullableMultipleDeviceModel, NullableMultipleDeviceModel) NullableMultipleDeviceModel) *NullableMultipleDeviceModelStream {
	result := NullableMultipleDeviceModelStreamOf()
	self.ForEach(func(v NullableMultipleDeviceModel, i int) {
		if i == 0 {
			result.Add(fn(NullableMultipleDeviceModel{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableMultipleDeviceModelStream) AllMatch(fn func(NullableMultipleDeviceModel, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableMultipleDeviceModelStream) AnyMatch(fn func(NullableMultipleDeviceModel, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableMultipleDeviceModelStream) Clone() *NullableMultipleDeviceModelStream {
	temp := make([]NullableMultipleDeviceModel, self.Len())
	copy(temp, *self)
	return (*NullableMultipleDeviceModelStream)(&temp)
}
func (self *NullableMultipleDeviceModelStream) Copy() *NullableMultipleDeviceModelStream {
	return self.Clone()
}
func (self *NullableMultipleDeviceModelStream) Concat(arg []NullableMultipleDeviceModel) *NullableMultipleDeviceModelStream {
	return self.AddAll(arg...)
}
func (self *NullableMultipleDeviceModelStream) Contains(arg NullableMultipleDeviceModel) bool {
	return self.FindIndex(func(_arg NullableMultipleDeviceModel, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableMultipleDeviceModelStream) Clean() *NullableMultipleDeviceModelStream {
	*self = NullableMultipleDeviceModelStreamOf()
	return self
}
func (self *NullableMultipleDeviceModelStream) Delete(index int) *NullableMultipleDeviceModelStream {
	return self.DeleteRange(index, index)
}
func (self *NullableMultipleDeviceModelStream) DeleteRange(startIndex, endIndex int) *NullableMultipleDeviceModelStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableMultipleDeviceModelStream) Distinct() *NullableMultipleDeviceModelStream {
	caches := map[string]bool{}
	result := NullableMultipleDeviceModelStreamOf()
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
func (self *NullableMultipleDeviceModelStream) Each(fn func(NullableMultipleDeviceModel)) *NullableMultipleDeviceModelStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableMultipleDeviceModelStream) EachRight(fn func(NullableMultipleDeviceModel)) *NullableMultipleDeviceModelStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableMultipleDeviceModelStream) Equals(arr []NullableMultipleDeviceModel) bool {
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
func (self *NullableMultipleDeviceModelStream) Filter(fn func(NullableMultipleDeviceModel, int) bool) *NullableMultipleDeviceModelStream {
	result := NullableMultipleDeviceModelStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMultipleDeviceModelStream) FilterSlim(fn func(NullableMultipleDeviceModel, int) bool) *NullableMultipleDeviceModelStream {
	result := NullableMultipleDeviceModelStreamOf()
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
func (self *NullableMultipleDeviceModelStream) Find(fn func(NullableMultipleDeviceModel, int) bool) *NullableMultipleDeviceModel {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableMultipleDeviceModelStream) FindOr(fn func(NullableMultipleDeviceModel, int) bool, or NullableMultipleDeviceModel) NullableMultipleDeviceModel {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableMultipleDeviceModelStream) FindIndex(fn func(NullableMultipleDeviceModel, int) bool) int {
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
func (self *NullableMultipleDeviceModelStream) First() *NullableMultipleDeviceModel {
	return self.Get(0)
}
func (self *NullableMultipleDeviceModelStream) FirstOr(arg NullableMultipleDeviceModel) NullableMultipleDeviceModel {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleDeviceModelStream) ForEach(fn func(NullableMultipleDeviceModel, int)) *NullableMultipleDeviceModelStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableMultipleDeviceModelStream) ForEachRight(fn func(NullableMultipleDeviceModel, int)) *NullableMultipleDeviceModelStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableMultipleDeviceModelStream) GroupBy(fn func(NullableMultipleDeviceModel, int) string) map[string][]NullableMultipleDeviceModel {
	m := map[string][]NullableMultipleDeviceModel{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableMultipleDeviceModelStream) GroupByValues(fn func(NullableMultipleDeviceModel, int) string) [][]NullableMultipleDeviceModel {
	var tmp [][]NullableMultipleDeviceModel
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableMultipleDeviceModelStream) IndexOf(arg NullableMultipleDeviceModel) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableMultipleDeviceModelStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableMultipleDeviceModelStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableMultipleDeviceModelStream) Last() *NullableMultipleDeviceModel {
	return self.Get(self.Len() - 1)
}
func (self *NullableMultipleDeviceModelStream) LastOr(arg NullableMultipleDeviceModel) NullableMultipleDeviceModel {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleDeviceModelStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableMultipleDeviceModelStream) Limit(limit int) *NullableMultipleDeviceModelStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableMultipleDeviceModelStream) Map(fn func(NullableMultipleDeviceModel, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDeviceModelStream) Map2Int(fn func(NullableMultipleDeviceModel, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDeviceModelStream) Map2Int32(fn func(NullableMultipleDeviceModel, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDeviceModelStream) Map2Int64(fn func(NullableMultipleDeviceModel, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDeviceModelStream) Map2Float32(fn func(NullableMultipleDeviceModel, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDeviceModelStream) Map2Float64(fn func(NullableMultipleDeviceModel, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDeviceModelStream) Map2Bool(fn func(NullableMultipleDeviceModel, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDeviceModelStream) Map2Bytes(fn func(NullableMultipleDeviceModel, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDeviceModelStream) Map2String(fn func(NullableMultipleDeviceModel, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDeviceModelStream) Max(fn func(NullableMultipleDeviceModel, int) float64) *NullableMultipleDeviceModel {
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
func (self *NullableMultipleDeviceModelStream) Min(fn func(NullableMultipleDeviceModel, int) float64) *NullableMultipleDeviceModel {
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
func (self *NullableMultipleDeviceModelStream) NoneMatch(fn func(NullableMultipleDeviceModel, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableMultipleDeviceModelStream) Get(index int) *NullableMultipleDeviceModel {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableMultipleDeviceModelStream) GetOr(index int, arg NullableMultipleDeviceModel) NullableMultipleDeviceModel {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleDeviceModelStream) Peek(fn func(*NullableMultipleDeviceModel, int)) *NullableMultipleDeviceModelStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *NullableMultipleDeviceModelStream) Reduce(fn func(NullableMultipleDeviceModel, NullableMultipleDeviceModel, int) NullableMultipleDeviceModel) *NullableMultipleDeviceModelStream {
	return self.ReduceInit(fn, NullableMultipleDeviceModel{})
}
func (self *NullableMultipleDeviceModelStream) ReduceInit(fn func(NullableMultipleDeviceModel, NullableMultipleDeviceModel, int) NullableMultipleDeviceModel, initialValue NullableMultipleDeviceModel) *NullableMultipleDeviceModelStream {
	result := NullableMultipleDeviceModelStreamOf()
	self.ForEach(func(v NullableMultipleDeviceModel, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableMultipleDeviceModelStream) ReduceInterface(fn func(interface{}, NullableMultipleDeviceModel, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableMultipleDeviceModel{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleDeviceModelStream) ReduceString(fn func(string, NullableMultipleDeviceModel, int) string) []string {
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
func (self *NullableMultipleDeviceModelStream) ReduceInt(fn func(int, NullableMultipleDeviceModel, int) int) []int {
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
func (self *NullableMultipleDeviceModelStream) ReduceInt32(fn func(int32, NullableMultipleDeviceModel, int) int32) []int32 {
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
func (self *NullableMultipleDeviceModelStream) ReduceInt64(fn func(int64, NullableMultipleDeviceModel, int) int64) []int64 {
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
func (self *NullableMultipleDeviceModelStream) ReduceFloat32(fn func(float32, NullableMultipleDeviceModel, int) float32) []float32 {
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
func (self *NullableMultipleDeviceModelStream) ReduceFloat64(fn func(float64, NullableMultipleDeviceModel, int) float64) []float64 {
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
func (self *NullableMultipleDeviceModelStream) ReduceBool(fn func(bool, NullableMultipleDeviceModel, int) bool) []bool {
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
func (self *NullableMultipleDeviceModelStream) Reverse() *NullableMultipleDeviceModelStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableMultipleDeviceModelStream) Replace(fn func(NullableMultipleDeviceModel, int) NullableMultipleDeviceModel) *NullableMultipleDeviceModelStream {
	return self.ForEach(func(v NullableMultipleDeviceModel, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableMultipleDeviceModelStream) Select(fn func(NullableMultipleDeviceModel) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableMultipleDeviceModelStream) Set(index int, val NullableMultipleDeviceModel) *NullableMultipleDeviceModelStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableMultipleDeviceModelStream) Skip(skip int) *NullableMultipleDeviceModelStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableMultipleDeviceModelStream) SkippingEach(fn func(NullableMultipleDeviceModel, int) int) *NullableMultipleDeviceModelStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableMultipleDeviceModelStream) Slice(startIndex, n int) *NullableMultipleDeviceModelStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableMultipleDeviceModel{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableMultipleDeviceModelStream) Sort(fn func(i, j int) bool) *NullableMultipleDeviceModelStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableMultipleDeviceModelStream) Tail() *NullableMultipleDeviceModel {
	return self.Last()
}
func (self *NullableMultipleDeviceModelStream) TailOr(arg NullableMultipleDeviceModel) NullableMultipleDeviceModel {
	return self.LastOr(arg)
}
func (self *NullableMultipleDeviceModelStream) ToList() []NullableMultipleDeviceModel {
	return self.Val()
}
func (self *NullableMultipleDeviceModelStream) Unique() *NullableMultipleDeviceModelStream {
	return self.Distinct()
}
func (self *NullableMultipleDeviceModelStream) Val() []NullableMultipleDeviceModel {
	if self == nil {
		return []NullableMultipleDeviceModel{}
	}
	return *self.Copy()
}
func (self *NullableMultipleDeviceModelStream) While(fn func(NullableMultipleDeviceModel, int) bool) *NullableMultipleDeviceModelStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableMultipleDeviceModelStream) Where(fn func(NullableMultipleDeviceModel) bool) *NullableMultipleDeviceModelStream {
	result := NullableMultipleDeviceModelStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMultipleDeviceModelStream) WhereSlim(fn func(NullableMultipleDeviceModel) bool) *NullableMultipleDeviceModelStream {
	result := NullableMultipleDeviceModelStreamOf()
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
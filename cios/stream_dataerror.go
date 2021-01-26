/*
 * Collection utility of DataError Struct
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

type DataErrorStream []DataError

func DataErrorStreamOf(arg ...DataError) DataErrorStream {
	return arg
}
func DataErrorStreamFrom(arg []DataError) DataErrorStream {
	return arg
}
func CreateDataErrorStream(arg ...DataError) *DataErrorStream {
	tmp := DataErrorStreamOf(arg...)
	return &tmp
}
func GenerateDataErrorStream(arg []DataError) *DataErrorStream {
	tmp := DataErrorStreamFrom(arg)
	return &tmp
}

func (self *DataErrorStream) Add(arg DataError) *DataErrorStream {
	return self.AddAll(arg)
}
func (self *DataErrorStream) AddAll(arg ...DataError) *DataErrorStream {
	*self = append(*self, arg...)
	return self
}
func (self *DataErrorStream) AddSafe(arg *DataError) *DataErrorStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *DataErrorStream) Aggregate(fn func(DataError, DataError) DataError) *DataErrorStream {
	result := DataErrorStreamOf()
	self.ForEach(func(v DataError, i int) {
		if i == 0 {
			result.Add(fn(DataError{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *DataErrorStream) AllMatch(fn func(DataError, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *DataErrorStream) AnyMatch(fn func(DataError, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *DataErrorStream) Clone() *DataErrorStream {
	temp := make([]DataError, self.Len())
	copy(temp, *self)
	return (*DataErrorStream)(&temp)
}
func (self *DataErrorStream) Copy() *DataErrorStream {
	return self.Clone()
}
func (self *DataErrorStream) Concat(arg []DataError) *DataErrorStream {
	return self.AddAll(arg...)
}
func (self *DataErrorStream) Contains(arg DataError) bool {
	return self.FindIndex(func(_arg DataError, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *DataErrorStream) Clean() *DataErrorStream {
	*self = DataErrorStreamOf()
	return self
}
func (self *DataErrorStream) Delete(index int) *DataErrorStream {
	return self.DeleteRange(index, index)
}
func (self *DataErrorStream) DeleteRange(startIndex, endIndex int) *DataErrorStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *DataErrorStream) Distinct() *DataErrorStream {
	caches := map[string]bool{}
	result := DataErrorStreamOf()
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
func (self *DataErrorStream) Each(fn func(DataError)) *DataErrorStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *DataErrorStream) EachRight(fn func(DataError)) *DataErrorStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *DataErrorStream) Equals(arr []DataError) bool {
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
func (self *DataErrorStream) Filter(fn func(DataError, int) bool) *DataErrorStream {
	result := DataErrorStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *DataErrorStream) FilterSlim(fn func(DataError, int) bool) *DataErrorStream {
	result := DataErrorStreamOf()
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
func (self *DataErrorStream) Find(fn func(DataError, int) bool) *DataError {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *DataErrorStream) FindOr(fn func(DataError, int) bool, or DataError) DataError {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *DataErrorStream) FindIndex(fn func(DataError, int) bool) int {
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
func (self *DataErrorStream) First() *DataError {
	return self.Get(0)
}
func (self *DataErrorStream) FirstOr(arg DataError) DataError {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *DataErrorStream) ForEach(fn func(DataError, int)) *DataErrorStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *DataErrorStream) ForEachRight(fn func(DataError, int)) *DataErrorStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *DataErrorStream) GroupBy(fn func(DataError, int) string) map[string][]DataError {
	m := map[string][]DataError{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *DataErrorStream) GroupByValues(fn func(DataError, int) string) [][]DataError {
	var tmp [][]DataError
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *DataErrorStream) IndexOf(arg DataError) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *DataErrorStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *DataErrorStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *DataErrorStream) Last() *DataError {
	return self.Get(self.Len() - 1)
}
func (self *DataErrorStream) LastOr(arg DataError) DataError {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *DataErrorStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *DataErrorStream) Limit(limit int) *DataErrorStream {
	self.Slice(0, limit)
	return self
}

func (self *DataErrorStream) Map(fn func(DataError, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DataErrorStream) Map2Int(fn func(DataError, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DataErrorStream) Map2Int32(fn func(DataError, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DataErrorStream) Map2Int64(fn func(DataError, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DataErrorStream) Map2Float32(fn func(DataError, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DataErrorStream) Map2Float64(fn func(DataError, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DataErrorStream) Map2Bool(fn func(DataError, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DataErrorStream) Map2Bytes(fn func(DataError, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DataErrorStream) Map2String(fn func(DataError, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DataErrorStream) Max(fn func(DataError, int) float64) *DataError {
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
func (self *DataErrorStream) Min(fn func(DataError, int) float64) *DataError {
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
func (self *DataErrorStream) NoneMatch(fn func(DataError, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *DataErrorStream) Get(index int) *DataError {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *DataErrorStream) GetOr(index int, arg DataError) DataError {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *DataErrorStream) Peek(fn func(*DataError, int)) *DataErrorStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *DataErrorStream) Reduce(fn func(DataError, DataError, int) DataError) *DataErrorStream {
	return self.ReduceInit(fn, DataError{})
}
func (self *DataErrorStream) ReduceInit(fn func(DataError, DataError, int) DataError, initialValue DataError) *DataErrorStream {
	result := DataErrorStreamOf()
	self.ForEach(func(v DataError, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *DataErrorStream) ReduceInterface(fn func(interface{}, DataError, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(DataError{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *DataErrorStream) ReduceString(fn func(string, DataError, int) string) []string {
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
func (self *DataErrorStream) ReduceInt(fn func(int, DataError, int) int) []int {
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
func (self *DataErrorStream) ReduceInt32(fn func(int32, DataError, int) int32) []int32 {
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
func (self *DataErrorStream) ReduceInt64(fn func(int64, DataError, int) int64) []int64 {
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
func (self *DataErrorStream) ReduceFloat32(fn func(float32, DataError, int) float32) []float32 {
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
func (self *DataErrorStream) ReduceFloat64(fn func(float64, DataError, int) float64) []float64 {
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
func (self *DataErrorStream) ReduceBool(fn func(bool, DataError, int) bool) []bool {
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
func (self *DataErrorStream) Reverse() *DataErrorStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *DataErrorStream) Replace(fn func(DataError, int) DataError) *DataErrorStream {
	return self.ForEach(func(v DataError, i int) { self.Set(i, fn(v, i)) })
}
func (self *DataErrorStream) Select(fn func(DataError) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *DataErrorStream) Set(index int, val DataError) *DataErrorStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *DataErrorStream) Skip(skip int) *DataErrorStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *DataErrorStream) SkippingEach(fn func(DataError, int) int) *DataErrorStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *DataErrorStream) Slice(startIndex, n int) *DataErrorStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []DataError{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *DataErrorStream) Sort(fn func(i, j int) bool) *DataErrorStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *DataErrorStream) Tail() *DataError {
	return self.Last()
}
func (self *DataErrorStream) TailOr(arg DataError) DataError {
	return self.LastOr(arg)
}
func (self *DataErrorStream) ToList() []DataError {
	return self.Val()
}
func (self *DataErrorStream) Unique() *DataErrorStream {
	return self.Distinct()
}
func (self *DataErrorStream) Val() []DataError {
	if self == nil {
		return []DataError{}
	}
	return *self.Copy()
}
func (self *DataErrorStream) While(fn func(DataError, int) bool) *DataErrorStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *DataErrorStream) Where(fn func(DataError) bool) *DataErrorStream {
	result := DataErrorStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *DataErrorStream) WhereSlim(fn func(DataError) bool) *DataErrorStream {
	result := DataErrorStreamOf()
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
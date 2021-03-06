/*
 * Collection utility of NullableSingleSession Struct
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

type NullableSingleSessionStream []NullableSingleSession

func NullableSingleSessionStreamOf(arg ...NullableSingleSession) NullableSingleSessionStream {
	return arg
}
func NullableSingleSessionStreamFrom(arg []NullableSingleSession) NullableSingleSessionStream {
	return arg
}
func CreateNullableSingleSessionStream(arg ...NullableSingleSession) *NullableSingleSessionStream {
	tmp := NullableSingleSessionStreamOf(arg...)
	return &tmp
}
func GenerateNullableSingleSessionStream(arg []NullableSingleSession) *NullableSingleSessionStream {
	tmp := NullableSingleSessionStreamFrom(arg)
	return &tmp
}

func (self *NullableSingleSessionStream) Add(arg NullableSingleSession) *NullableSingleSessionStream {
	return self.AddAll(arg)
}
func (self *NullableSingleSessionStream) AddAll(arg ...NullableSingleSession) *NullableSingleSessionStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableSingleSessionStream) AddSafe(arg *NullableSingleSession) *NullableSingleSessionStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableSingleSessionStream) Aggregate(fn func(NullableSingleSession, NullableSingleSession) NullableSingleSession) *NullableSingleSessionStream {
	result := NullableSingleSessionStreamOf()
	self.ForEach(func(v NullableSingleSession, i int) {
		if i == 0 {
			result.Add(fn(NullableSingleSession{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableSingleSessionStream) AllMatch(fn func(NullableSingleSession, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableSingleSessionStream) AnyMatch(fn func(NullableSingleSession, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableSingleSessionStream) Clone() *NullableSingleSessionStream {
	temp := make([]NullableSingleSession, self.Len())
	copy(temp, *self)
	return (*NullableSingleSessionStream)(&temp)
}
func (self *NullableSingleSessionStream) Copy() *NullableSingleSessionStream {
	return self.Clone()
}
func (self *NullableSingleSessionStream) Concat(arg []NullableSingleSession) *NullableSingleSessionStream {
	return self.AddAll(arg...)
}
func (self *NullableSingleSessionStream) Contains(arg NullableSingleSession) bool {
	return self.FindIndex(func(_arg NullableSingleSession, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableSingleSessionStream) Clean() *NullableSingleSessionStream {
	*self = NullableSingleSessionStreamOf()
	return self
}
func (self *NullableSingleSessionStream) Delete(index int) *NullableSingleSessionStream {
	return self.DeleteRange(index, index)
}
func (self *NullableSingleSessionStream) DeleteRange(startIndex, endIndex int) *NullableSingleSessionStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableSingleSessionStream) Distinct() *NullableSingleSessionStream {
	caches := map[string]bool{}
	result := NullableSingleSessionStreamOf()
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
func (self *NullableSingleSessionStream) Each(fn func(NullableSingleSession)) *NullableSingleSessionStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableSingleSessionStream) EachRight(fn func(NullableSingleSession)) *NullableSingleSessionStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableSingleSessionStream) Equals(arr []NullableSingleSession) bool {
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
func (self *NullableSingleSessionStream) Filter(fn func(NullableSingleSession, int) bool) *NullableSingleSessionStream {
	result := NullableSingleSessionStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSingleSessionStream) FilterSlim(fn func(NullableSingleSession, int) bool) *NullableSingleSessionStream {
	result := NullableSingleSessionStreamOf()
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
func (self *NullableSingleSessionStream) Find(fn func(NullableSingleSession, int) bool) *NullableSingleSession {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableSingleSessionStream) FindOr(fn func(NullableSingleSession, int) bool, or NullableSingleSession) NullableSingleSession {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableSingleSessionStream) FindIndex(fn func(NullableSingleSession, int) bool) int {
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
func (self *NullableSingleSessionStream) First() *NullableSingleSession {
	return self.Get(0)
}
func (self *NullableSingleSessionStream) FirstOr(arg NullableSingleSession) NullableSingleSession {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSingleSessionStream) ForEach(fn func(NullableSingleSession, int)) *NullableSingleSessionStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableSingleSessionStream) ForEachRight(fn func(NullableSingleSession, int)) *NullableSingleSessionStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableSingleSessionStream) GroupBy(fn func(NullableSingleSession, int) string) map[string][]NullableSingleSession {
	m := map[string][]NullableSingleSession{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableSingleSessionStream) GroupByValues(fn func(NullableSingleSession, int) string) [][]NullableSingleSession {
	var tmp [][]NullableSingleSession
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableSingleSessionStream) IndexOf(arg NullableSingleSession) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableSingleSessionStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableSingleSessionStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableSingleSessionStream) Last() *NullableSingleSession {
	return self.Get(self.Len() - 1)
}
func (self *NullableSingleSessionStream) LastOr(arg NullableSingleSession) NullableSingleSession {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSingleSessionStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableSingleSessionStream) Limit(limit int) *NullableSingleSessionStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableSingleSessionStream) Map(fn func(NullableSingleSession, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleSessionStream) Map2Int(fn func(NullableSingleSession, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleSessionStream) Map2Int32(fn func(NullableSingleSession, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleSessionStream) Map2Int64(fn func(NullableSingleSession, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleSessionStream) Map2Float32(fn func(NullableSingleSession, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleSessionStream) Map2Float64(fn func(NullableSingleSession, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleSessionStream) Map2Bool(fn func(NullableSingleSession, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleSessionStream) Map2Bytes(fn func(NullableSingleSession, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleSessionStream) Map2String(fn func(NullableSingleSession, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleSessionStream) Max(fn func(NullableSingleSession, int) float64) *NullableSingleSession {
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
func (self *NullableSingleSessionStream) Min(fn func(NullableSingleSession, int) float64) *NullableSingleSession {
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
func (self *NullableSingleSessionStream) NoneMatch(fn func(NullableSingleSession, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableSingleSessionStream) Get(index int) *NullableSingleSession {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableSingleSessionStream) GetOr(index int, arg NullableSingleSession) NullableSingleSession {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSingleSessionStream) Peek(fn func(*NullableSingleSession, int)) *NullableSingleSessionStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableSingleSessionStream) Reduce(fn func(NullableSingleSession, NullableSingleSession, int) NullableSingleSession) *NullableSingleSessionStream {
	return self.ReduceInit(fn, NullableSingleSession{})
}
func (self *NullableSingleSessionStream) ReduceInit(fn func(NullableSingleSession, NullableSingleSession, int) NullableSingleSession, initialValue NullableSingleSession) *NullableSingleSessionStream {
	result := NullableSingleSessionStreamOf()
	self.ForEach(func(v NullableSingleSession, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableSingleSessionStream) ReduceInterface(fn func(interface{}, NullableSingleSession, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableSingleSession{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableSingleSessionStream) ReduceString(fn func(string, NullableSingleSession, int) string) []string {
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
func (self *NullableSingleSessionStream) ReduceInt(fn func(int, NullableSingleSession, int) int) []int {
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
func (self *NullableSingleSessionStream) ReduceInt32(fn func(int32, NullableSingleSession, int) int32) []int32 {
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
func (self *NullableSingleSessionStream) ReduceInt64(fn func(int64, NullableSingleSession, int) int64) []int64 {
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
func (self *NullableSingleSessionStream) ReduceFloat32(fn func(float32, NullableSingleSession, int) float32) []float32 {
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
func (self *NullableSingleSessionStream) ReduceFloat64(fn func(float64, NullableSingleSession, int) float64) []float64 {
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
func (self *NullableSingleSessionStream) ReduceBool(fn func(bool, NullableSingleSession, int) bool) []bool {
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
func (self *NullableSingleSessionStream) Reverse() *NullableSingleSessionStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableSingleSessionStream) Replace(fn func(NullableSingleSession, int) NullableSingleSession) *NullableSingleSessionStream {
	return self.ForEach(func(v NullableSingleSession, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableSingleSessionStream) Select(fn func(NullableSingleSession) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableSingleSessionStream) Set(index int, val NullableSingleSession) *NullableSingleSessionStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableSingleSessionStream) Skip(skip int) *NullableSingleSessionStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableSingleSessionStream) SkippingEach(fn func(NullableSingleSession, int) int) *NullableSingleSessionStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableSingleSessionStream) Slice(startIndex, n int) *NullableSingleSessionStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableSingleSession{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableSingleSessionStream) Sort(fn func(i, j int) bool) *NullableSingleSessionStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableSingleSessionStream) Tail() *NullableSingleSession {
	return self.Last()
}
func (self *NullableSingleSessionStream) TailOr(arg NullableSingleSession) NullableSingleSession {
	return self.LastOr(arg)
}
func (self *NullableSingleSessionStream) ToList() []NullableSingleSession {
	return self.Val()
}
func (self *NullableSingleSessionStream) Unique() *NullableSingleSessionStream {
	return self.Distinct()
}
func (self *NullableSingleSessionStream) Val() []NullableSingleSession {
	if self == nil {
		return []NullableSingleSession{}
	}
	return *self.Copy()
}
func (self *NullableSingleSessionStream) While(fn func(NullableSingleSession, int) bool) *NullableSingleSessionStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableSingleSessionStream) Where(fn func(NullableSingleSession) bool) *NullableSingleSessionStream {
	result := NullableSingleSessionStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSingleSessionStream) WhereSlim(fn func(NullableSingleSession) bool) *NullableSingleSessionStream {
	result := NullableSingleSessionStreamOf()
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

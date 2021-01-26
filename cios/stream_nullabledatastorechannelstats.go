/*
 * Collection utility of NullableDataStoreChannelStats Struct
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

type NullableDataStoreChannelStatsStream []NullableDataStoreChannelStats

func NullableDataStoreChannelStatsStreamOf(arg ...NullableDataStoreChannelStats) NullableDataStoreChannelStatsStream {
	return arg
}
func NullableDataStoreChannelStatsStreamFrom(arg []NullableDataStoreChannelStats) NullableDataStoreChannelStatsStream {
	return arg
}
func CreateNullableDataStoreChannelStatsStream(arg ...NullableDataStoreChannelStats) *NullableDataStoreChannelStatsStream {
	tmp := NullableDataStoreChannelStatsStreamOf(arg...)
	return &tmp
}
func GenerateNullableDataStoreChannelStatsStream(arg []NullableDataStoreChannelStats) *NullableDataStoreChannelStatsStream {
	tmp := NullableDataStoreChannelStatsStreamFrom(arg)
	return &tmp
}

func (self *NullableDataStoreChannelStatsStream) Add(arg NullableDataStoreChannelStats) *NullableDataStoreChannelStatsStream {
	return self.AddAll(arg)
}
func (self *NullableDataStoreChannelStatsStream) AddAll(arg ...NullableDataStoreChannelStats) *NullableDataStoreChannelStatsStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableDataStoreChannelStatsStream) AddSafe(arg *NullableDataStoreChannelStats) *NullableDataStoreChannelStatsStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableDataStoreChannelStatsStream) Aggregate(fn func(NullableDataStoreChannelStats, NullableDataStoreChannelStats) NullableDataStoreChannelStats) *NullableDataStoreChannelStatsStream {
	result := NullableDataStoreChannelStatsStreamOf()
	self.ForEach(func(v NullableDataStoreChannelStats, i int) {
		if i == 0 {
			result.Add(fn(NullableDataStoreChannelStats{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableDataStoreChannelStatsStream) AllMatch(fn func(NullableDataStoreChannelStats, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableDataStoreChannelStatsStream) AnyMatch(fn func(NullableDataStoreChannelStats, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableDataStoreChannelStatsStream) Clone() *NullableDataStoreChannelStatsStream {
	temp := make([]NullableDataStoreChannelStats, self.Len())
	copy(temp, *self)
	return (*NullableDataStoreChannelStatsStream)(&temp)
}
func (self *NullableDataStoreChannelStatsStream) Copy() *NullableDataStoreChannelStatsStream {
	return self.Clone()
}
func (self *NullableDataStoreChannelStatsStream) Concat(arg []NullableDataStoreChannelStats) *NullableDataStoreChannelStatsStream {
	return self.AddAll(arg...)
}
func (self *NullableDataStoreChannelStatsStream) Contains(arg NullableDataStoreChannelStats) bool {
	return self.FindIndex(func(_arg NullableDataStoreChannelStats, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableDataStoreChannelStatsStream) Clean() *NullableDataStoreChannelStatsStream {
	*self = NullableDataStoreChannelStatsStreamOf()
	return self
}
func (self *NullableDataStoreChannelStatsStream) Delete(index int) *NullableDataStoreChannelStatsStream {
	return self.DeleteRange(index, index)
}
func (self *NullableDataStoreChannelStatsStream) DeleteRange(startIndex, endIndex int) *NullableDataStoreChannelStatsStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableDataStoreChannelStatsStream) Distinct() *NullableDataStoreChannelStatsStream {
	caches := map[string]bool{}
	result := NullableDataStoreChannelStatsStreamOf()
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
func (self *NullableDataStoreChannelStatsStream) Each(fn func(NullableDataStoreChannelStats)) *NullableDataStoreChannelStatsStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableDataStoreChannelStatsStream) EachRight(fn func(NullableDataStoreChannelStats)) *NullableDataStoreChannelStatsStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableDataStoreChannelStatsStream) Equals(arr []NullableDataStoreChannelStats) bool {
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
func (self *NullableDataStoreChannelStatsStream) Filter(fn func(NullableDataStoreChannelStats, int) bool) *NullableDataStoreChannelStatsStream {
	result := NullableDataStoreChannelStatsStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableDataStoreChannelStatsStream) FilterSlim(fn func(NullableDataStoreChannelStats, int) bool) *NullableDataStoreChannelStatsStream {
	result := NullableDataStoreChannelStatsStreamOf()
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
func (self *NullableDataStoreChannelStatsStream) Find(fn func(NullableDataStoreChannelStats, int) bool) *NullableDataStoreChannelStats {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableDataStoreChannelStatsStream) FindOr(fn func(NullableDataStoreChannelStats, int) bool, or NullableDataStoreChannelStats) NullableDataStoreChannelStats {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableDataStoreChannelStatsStream) FindIndex(fn func(NullableDataStoreChannelStats, int) bool) int {
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
func (self *NullableDataStoreChannelStatsStream) First() *NullableDataStoreChannelStats {
	return self.Get(0)
}
func (self *NullableDataStoreChannelStatsStream) FirstOr(arg NullableDataStoreChannelStats) NullableDataStoreChannelStats {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableDataStoreChannelStatsStream) ForEach(fn func(NullableDataStoreChannelStats, int)) *NullableDataStoreChannelStatsStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableDataStoreChannelStatsStream) ForEachRight(fn func(NullableDataStoreChannelStats, int)) *NullableDataStoreChannelStatsStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableDataStoreChannelStatsStream) GroupBy(fn func(NullableDataStoreChannelStats, int) string) map[string][]NullableDataStoreChannelStats {
	m := map[string][]NullableDataStoreChannelStats{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableDataStoreChannelStatsStream) GroupByValues(fn func(NullableDataStoreChannelStats, int) string) [][]NullableDataStoreChannelStats {
	var tmp [][]NullableDataStoreChannelStats
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableDataStoreChannelStatsStream) IndexOf(arg NullableDataStoreChannelStats) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableDataStoreChannelStatsStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableDataStoreChannelStatsStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableDataStoreChannelStatsStream) Last() *NullableDataStoreChannelStats {
	return self.Get(self.Len() - 1)
}
func (self *NullableDataStoreChannelStatsStream) LastOr(arg NullableDataStoreChannelStats) NullableDataStoreChannelStats {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableDataStoreChannelStatsStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableDataStoreChannelStatsStream) Limit(limit int) *NullableDataStoreChannelStatsStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableDataStoreChannelStatsStream) Map(fn func(NullableDataStoreChannelStats, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDataStoreChannelStatsStream) Map2Int(fn func(NullableDataStoreChannelStats, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDataStoreChannelStatsStream) Map2Int32(fn func(NullableDataStoreChannelStats, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDataStoreChannelStatsStream) Map2Int64(fn func(NullableDataStoreChannelStats, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDataStoreChannelStatsStream) Map2Float32(fn func(NullableDataStoreChannelStats, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDataStoreChannelStatsStream) Map2Float64(fn func(NullableDataStoreChannelStats, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDataStoreChannelStatsStream) Map2Bool(fn func(NullableDataStoreChannelStats, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDataStoreChannelStatsStream) Map2Bytes(fn func(NullableDataStoreChannelStats, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDataStoreChannelStatsStream) Map2String(fn func(NullableDataStoreChannelStats, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDataStoreChannelStatsStream) Max(fn func(NullableDataStoreChannelStats, int) float64) *NullableDataStoreChannelStats {
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
func (self *NullableDataStoreChannelStatsStream) Min(fn func(NullableDataStoreChannelStats, int) float64) *NullableDataStoreChannelStats {
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
func (self *NullableDataStoreChannelStatsStream) NoneMatch(fn func(NullableDataStoreChannelStats, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableDataStoreChannelStatsStream) Get(index int) *NullableDataStoreChannelStats {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableDataStoreChannelStatsStream) GetOr(index int, arg NullableDataStoreChannelStats) NullableDataStoreChannelStats {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableDataStoreChannelStatsStream) Peek(fn func(*NullableDataStoreChannelStats, int)) *NullableDataStoreChannelStatsStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *NullableDataStoreChannelStatsStream) Reduce(fn func(NullableDataStoreChannelStats, NullableDataStoreChannelStats, int) NullableDataStoreChannelStats) *NullableDataStoreChannelStatsStream {
	return self.ReduceInit(fn, NullableDataStoreChannelStats{})
}
func (self *NullableDataStoreChannelStatsStream) ReduceInit(fn func(NullableDataStoreChannelStats, NullableDataStoreChannelStats, int) NullableDataStoreChannelStats, initialValue NullableDataStoreChannelStats) *NullableDataStoreChannelStatsStream {
	result := NullableDataStoreChannelStatsStreamOf()
	self.ForEach(func(v NullableDataStoreChannelStats, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableDataStoreChannelStatsStream) ReduceInterface(fn func(interface{}, NullableDataStoreChannelStats, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableDataStoreChannelStats{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableDataStoreChannelStatsStream) ReduceString(fn func(string, NullableDataStoreChannelStats, int) string) []string {
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
func (self *NullableDataStoreChannelStatsStream) ReduceInt(fn func(int, NullableDataStoreChannelStats, int) int) []int {
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
func (self *NullableDataStoreChannelStatsStream) ReduceInt32(fn func(int32, NullableDataStoreChannelStats, int) int32) []int32 {
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
func (self *NullableDataStoreChannelStatsStream) ReduceInt64(fn func(int64, NullableDataStoreChannelStats, int) int64) []int64 {
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
func (self *NullableDataStoreChannelStatsStream) ReduceFloat32(fn func(float32, NullableDataStoreChannelStats, int) float32) []float32 {
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
func (self *NullableDataStoreChannelStatsStream) ReduceFloat64(fn func(float64, NullableDataStoreChannelStats, int) float64) []float64 {
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
func (self *NullableDataStoreChannelStatsStream) ReduceBool(fn func(bool, NullableDataStoreChannelStats, int) bool) []bool {
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
func (self *NullableDataStoreChannelStatsStream) Reverse() *NullableDataStoreChannelStatsStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableDataStoreChannelStatsStream) Replace(fn func(NullableDataStoreChannelStats, int) NullableDataStoreChannelStats) *NullableDataStoreChannelStatsStream {
	return self.ForEach(func(v NullableDataStoreChannelStats, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableDataStoreChannelStatsStream) Select(fn func(NullableDataStoreChannelStats) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableDataStoreChannelStatsStream) Set(index int, val NullableDataStoreChannelStats) *NullableDataStoreChannelStatsStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableDataStoreChannelStatsStream) Skip(skip int) *NullableDataStoreChannelStatsStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableDataStoreChannelStatsStream) SkippingEach(fn func(NullableDataStoreChannelStats, int) int) *NullableDataStoreChannelStatsStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableDataStoreChannelStatsStream) Slice(startIndex, n int) *NullableDataStoreChannelStatsStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableDataStoreChannelStats{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableDataStoreChannelStatsStream) Sort(fn func(i, j int) bool) *NullableDataStoreChannelStatsStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableDataStoreChannelStatsStream) Tail() *NullableDataStoreChannelStats {
	return self.Last()
}
func (self *NullableDataStoreChannelStatsStream) TailOr(arg NullableDataStoreChannelStats) NullableDataStoreChannelStats {
	return self.LastOr(arg)
}
func (self *NullableDataStoreChannelStatsStream) ToList() []NullableDataStoreChannelStats {
	return self.Val()
}
func (self *NullableDataStoreChannelStatsStream) Unique() *NullableDataStoreChannelStatsStream {
	return self.Distinct()
}
func (self *NullableDataStoreChannelStatsStream) Val() []NullableDataStoreChannelStats {
	if self == nil {
		return []NullableDataStoreChannelStats{}
	}
	return *self.Copy()
}
func (self *NullableDataStoreChannelStatsStream) While(fn func(NullableDataStoreChannelStats, int) bool) *NullableDataStoreChannelStatsStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableDataStoreChannelStatsStream) Where(fn func(NullableDataStoreChannelStats) bool) *NullableDataStoreChannelStatsStream {
	result := NullableDataStoreChannelStatsStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableDataStoreChannelStatsStream) WhereSlim(fn func(NullableDataStoreChannelStats) bool) *NullableDataStoreChannelStatsStream {
	result := NullableDataStoreChannelStatsStreamOf()
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
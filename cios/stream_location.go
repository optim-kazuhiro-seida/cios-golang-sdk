/*
 * Collection utility of Location Struct
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

type LocationStream []Location

func LocationStreamOf(arg ...Location) LocationStream {
	return arg
}
func LocationStreamFrom(arg []Location) LocationStream {
	return arg
}
func CreateLocationStream(arg ...Location) *LocationStream {
	tmp := LocationStreamOf(arg...)
	return &tmp
}
func GenerateLocationStream(arg []Location) *LocationStream {
	tmp := LocationStreamFrom(arg)
	return &tmp
}

func (self *LocationStream) Add(arg Location) *LocationStream {
	return self.AddAll(arg)
}
func (self *LocationStream) AddAll(arg ...Location) *LocationStream {
	*self = append(*self, arg...)
	return self
}
func (self *LocationStream) AddSafe(arg *Location) *LocationStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *LocationStream) Aggregate(fn func(Location, Location) Location) *LocationStream {
	result := LocationStreamOf()
	self.ForEach(func(v Location, i int) {
		if i == 0 {
			result.Add(fn(Location{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *LocationStream) AllMatch(fn func(Location, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *LocationStream) AnyMatch(fn func(Location, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *LocationStream) Clone() *LocationStream {
	temp := make([]Location, self.Len())
	copy(temp, *self)
	return (*LocationStream)(&temp)
}
func (self *LocationStream) Copy() *LocationStream {
	return self.Clone()
}
func (self *LocationStream) Concat(arg []Location) *LocationStream {
	return self.AddAll(arg...)
}
func (self *LocationStream) Contains(arg Location) bool {
	return self.FindIndex(func(_arg Location, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *LocationStream) Clean() *LocationStream {
	*self = LocationStreamOf()
	return self
}
func (self *LocationStream) Delete(index int) *LocationStream {
	return self.DeleteRange(index, index)
}
func (self *LocationStream) DeleteRange(startIndex, endIndex int) *LocationStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *LocationStream) Distinct() *LocationStream {
	caches := map[string]bool{}
	result := LocationStreamOf()
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
func (self *LocationStream) Each(fn func(Location)) *LocationStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *LocationStream) EachRight(fn func(Location)) *LocationStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *LocationStream) Equals(arr []Location) bool {
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
func (self *LocationStream) Filter(fn func(Location, int) bool) *LocationStream {
	result := LocationStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *LocationStream) FilterSlim(fn func(Location, int) bool) *LocationStream {
	result := LocationStreamOf()
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
func (self *LocationStream) Find(fn func(Location, int) bool) *Location {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *LocationStream) FindOr(fn func(Location, int) bool, or Location) Location {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *LocationStream) FindIndex(fn func(Location, int) bool) int {
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
func (self *LocationStream) First() *Location {
	return self.Get(0)
}
func (self *LocationStream) FirstOr(arg Location) Location {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *LocationStream) ForEach(fn func(Location, int)) *LocationStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *LocationStream) ForEachRight(fn func(Location, int)) *LocationStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *LocationStream) GroupBy(fn func(Location, int) string) map[string][]Location {
	m := map[string][]Location{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *LocationStream) GroupByValues(fn func(Location, int) string) [][]Location {
	var tmp [][]Location
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *LocationStream) IndexOf(arg Location) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *LocationStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *LocationStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *LocationStream) Last() *Location {
	return self.Get(self.Len() - 1)
}
func (self *LocationStream) LastOr(arg Location) Location {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *LocationStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *LocationStream) Limit(limit int) *LocationStream {
	self.Slice(0, limit)
	return self
}

func (self *LocationStream) Map(fn func(Location, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *LocationStream) Map2Int(fn func(Location, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *LocationStream) Map2Int32(fn func(Location, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *LocationStream) Map2Int64(fn func(Location, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *LocationStream) Map2Float32(fn func(Location, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *LocationStream) Map2Float64(fn func(Location, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *LocationStream) Map2Bool(fn func(Location, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *LocationStream) Map2Bytes(fn func(Location, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *LocationStream) Map2String(fn func(Location, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *LocationStream) Max(fn func(Location, int) float64) *Location {
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
func (self *LocationStream) Min(fn func(Location, int) float64) *Location {
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
func (self *LocationStream) NoneMatch(fn func(Location, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *LocationStream) Get(index int) *Location {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *LocationStream) GetOr(index int, arg Location) Location {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *LocationStream) Peek(fn func(*Location, int)) *LocationStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *LocationStream) Reduce(fn func(Location, Location, int) Location) *LocationStream {
	return self.ReduceInit(fn, Location{})
}
func (self *LocationStream) ReduceInit(fn func(Location, Location, int) Location, initialValue Location) *LocationStream {
	result := LocationStreamOf()
	self.ForEach(func(v Location, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *LocationStream) ReduceInterface(fn func(interface{}, Location, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(Location{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *LocationStream) ReduceString(fn func(string, Location, int) string) []string {
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
func (self *LocationStream) ReduceInt(fn func(int, Location, int) int) []int {
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
func (self *LocationStream) ReduceInt32(fn func(int32, Location, int) int32) []int32 {
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
func (self *LocationStream) ReduceInt64(fn func(int64, Location, int) int64) []int64 {
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
func (self *LocationStream) ReduceFloat32(fn func(float32, Location, int) float32) []float32 {
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
func (self *LocationStream) ReduceFloat64(fn func(float64, Location, int) float64) []float64 {
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
func (self *LocationStream) ReduceBool(fn func(bool, Location, int) bool) []bool {
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
func (self *LocationStream) Reverse() *LocationStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *LocationStream) Replace(fn func(Location, int) Location) *LocationStream {
	return self.ForEach(func(v Location, i int) { self.Set(i, fn(v, i)) })
}
func (self *LocationStream) Select(fn func(Location) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *LocationStream) Set(index int, val Location) *LocationStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *LocationStream) Skip(skip int) *LocationStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *LocationStream) SkippingEach(fn func(Location, int) int) *LocationStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *LocationStream) Slice(startIndex, n int) *LocationStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []Location{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *LocationStream) Sort(fn func(i, j int) bool) *LocationStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *LocationStream) Tail() *Location {
	return self.Last()
}
func (self *LocationStream) TailOr(arg Location) Location {
	return self.LastOr(arg)
}
func (self *LocationStream) ToList() []Location {
	return self.Val()
}
func (self *LocationStream) Unique() *LocationStream {
	return self.Distinct()
}
func (self *LocationStream) Val() []Location {
	if self == nil {
		return []Location{}
	}
	return *self.Copy()
}
func (self *LocationStream) While(fn func(Location, int) bool) *LocationStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *LocationStream) Where(fn func(Location) bool) *LocationStream {
	result := LocationStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *LocationStream) WhereSlim(fn func(Location) bool) *LocationStream {
	result := LocationStreamOf()
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

/*
 * Collection utility of MemberInfo Struct
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

type MemberInfoStream []MemberInfo

func MemberInfoStreamOf(arg ...MemberInfo) MemberInfoStream {
	return arg
}
func MemberInfoStreamFrom(arg []MemberInfo) MemberInfoStream {
	return arg
}
func CreateMemberInfoStream(arg ...MemberInfo) *MemberInfoStream {
	tmp := MemberInfoStreamOf(arg...)
	return &tmp
}
func GenerateMemberInfoStream(arg []MemberInfo) *MemberInfoStream {
	tmp := MemberInfoStreamFrom(arg)
	return &tmp
}

func (self *MemberInfoStream) Add(arg MemberInfo) *MemberInfoStream {
	return self.AddAll(arg)
}
func (self *MemberInfoStream) AddAll(arg ...MemberInfo) *MemberInfoStream {
	*self = append(*self, arg...)
	return self
}
func (self *MemberInfoStream) AddSafe(arg *MemberInfo) *MemberInfoStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *MemberInfoStream) Aggregate(fn func(MemberInfo, MemberInfo) MemberInfo) *MemberInfoStream {
	result := MemberInfoStreamOf()
	self.ForEach(func(v MemberInfo, i int) {
		if i == 0 {
			result.Add(fn(MemberInfo{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *MemberInfoStream) AllMatch(fn func(MemberInfo, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *MemberInfoStream) AnyMatch(fn func(MemberInfo, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *MemberInfoStream) Clone() *MemberInfoStream {
	temp := make([]MemberInfo, self.Len())
	copy(temp, *self)
	return (*MemberInfoStream)(&temp)
}
func (self *MemberInfoStream) Copy() *MemberInfoStream {
	return self.Clone()
}
func (self *MemberInfoStream) Concat(arg []MemberInfo) *MemberInfoStream {
	return self.AddAll(arg...)
}
func (self *MemberInfoStream) Contains(arg MemberInfo) bool {
	return self.FindIndex(func(_arg MemberInfo, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *MemberInfoStream) Clean() *MemberInfoStream {
	*self = MemberInfoStreamOf()
	return self
}
func (self *MemberInfoStream) Delete(index int) *MemberInfoStream {
	return self.DeleteRange(index, index)
}
func (self *MemberInfoStream) DeleteRange(startIndex, endIndex int) *MemberInfoStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *MemberInfoStream) Distinct() *MemberInfoStream {
	caches := map[string]bool{}
	result := MemberInfoStreamOf()
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
func (self *MemberInfoStream) Each(fn func(MemberInfo)) *MemberInfoStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *MemberInfoStream) EachRight(fn func(MemberInfo)) *MemberInfoStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *MemberInfoStream) Equals(arr []MemberInfo) bool {
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
func (self *MemberInfoStream) Filter(fn func(MemberInfo, int) bool) *MemberInfoStream {
	result := MemberInfoStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *MemberInfoStream) FilterSlim(fn func(MemberInfo, int) bool) *MemberInfoStream {
	result := MemberInfoStreamOf()
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
func (self *MemberInfoStream) Find(fn func(MemberInfo, int) bool) *MemberInfo {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *MemberInfoStream) FindOr(fn func(MemberInfo, int) bool, or MemberInfo) MemberInfo {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *MemberInfoStream) FindIndex(fn func(MemberInfo, int) bool) int {
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
func (self *MemberInfoStream) First() *MemberInfo {
	return self.Get(0)
}
func (self *MemberInfoStream) FirstOr(arg MemberInfo) MemberInfo {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *MemberInfoStream) ForEach(fn func(MemberInfo, int)) *MemberInfoStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *MemberInfoStream) ForEachRight(fn func(MemberInfo, int)) *MemberInfoStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *MemberInfoStream) GroupBy(fn func(MemberInfo, int) string) map[string][]MemberInfo {
	m := map[string][]MemberInfo{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *MemberInfoStream) GroupByValues(fn func(MemberInfo, int) string) [][]MemberInfo {
	var tmp [][]MemberInfo
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *MemberInfoStream) IndexOf(arg MemberInfo) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *MemberInfoStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *MemberInfoStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *MemberInfoStream) Last() *MemberInfo {
	return self.Get(self.Len() - 1)
}
func (self *MemberInfoStream) LastOr(arg MemberInfo) MemberInfo {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *MemberInfoStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *MemberInfoStream) Limit(limit int) *MemberInfoStream {
	self.Slice(0, limit)
	return self
}

func (self *MemberInfoStream) Map(fn func(MemberInfo, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MemberInfoStream) Map2Int(fn func(MemberInfo, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MemberInfoStream) Map2Int32(fn func(MemberInfo, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MemberInfoStream) Map2Int64(fn func(MemberInfo, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MemberInfoStream) Map2Float32(fn func(MemberInfo, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MemberInfoStream) Map2Float64(fn func(MemberInfo, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MemberInfoStream) Map2Bool(fn func(MemberInfo, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MemberInfoStream) Map2Bytes(fn func(MemberInfo, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MemberInfoStream) Map2String(fn func(MemberInfo, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MemberInfoStream) Max(fn func(MemberInfo, int) float64) *MemberInfo {
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
func (self *MemberInfoStream) Min(fn func(MemberInfo, int) float64) *MemberInfo {
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
func (self *MemberInfoStream) NoneMatch(fn func(MemberInfo, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *MemberInfoStream) Get(index int) *MemberInfo {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *MemberInfoStream) GetOr(index int, arg MemberInfo) MemberInfo {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *MemberInfoStream) Peek(fn func(*MemberInfo, int)) *MemberInfoStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *MemberInfoStream) Reduce(fn func(MemberInfo, MemberInfo, int) MemberInfo) *MemberInfoStream {
	return self.ReduceInit(fn, MemberInfo{})
}
func (self *MemberInfoStream) ReduceInit(fn func(MemberInfo, MemberInfo, int) MemberInfo, initialValue MemberInfo) *MemberInfoStream {
	result := MemberInfoStreamOf()
	self.ForEach(func(v MemberInfo, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *MemberInfoStream) ReduceInterface(fn func(interface{}, MemberInfo, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(MemberInfo{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *MemberInfoStream) ReduceString(fn func(string, MemberInfo, int) string) []string {
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
func (self *MemberInfoStream) ReduceInt(fn func(int, MemberInfo, int) int) []int {
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
func (self *MemberInfoStream) ReduceInt32(fn func(int32, MemberInfo, int) int32) []int32 {
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
func (self *MemberInfoStream) ReduceInt64(fn func(int64, MemberInfo, int) int64) []int64 {
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
func (self *MemberInfoStream) ReduceFloat32(fn func(float32, MemberInfo, int) float32) []float32 {
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
func (self *MemberInfoStream) ReduceFloat64(fn func(float64, MemberInfo, int) float64) []float64 {
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
func (self *MemberInfoStream) ReduceBool(fn func(bool, MemberInfo, int) bool) []bool {
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
func (self *MemberInfoStream) Reverse() *MemberInfoStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *MemberInfoStream) Replace(fn func(MemberInfo, int) MemberInfo) *MemberInfoStream {
	return self.ForEach(func(v MemberInfo, i int) { self.Set(i, fn(v, i)) })
}
func (self *MemberInfoStream) Select(fn func(MemberInfo) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *MemberInfoStream) Set(index int, val MemberInfo) *MemberInfoStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *MemberInfoStream) Skip(skip int) *MemberInfoStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *MemberInfoStream) SkippingEach(fn func(MemberInfo, int) int) *MemberInfoStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *MemberInfoStream) Slice(startIndex, n int) *MemberInfoStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []MemberInfo{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *MemberInfoStream) Sort(fn func(i, j int) bool) *MemberInfoStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *MemberInfoStream) Tail() *MemberInfo {
	return self.Last()
}
func (self *MemberInfoStream) TailOr(arg MemberInfo) MemberInfo {
	return self.LastOr(arg)
}
func (self *MemberInfoStream) ToList() []MemberInfo {
	return self.Val()
}
func (self *MemberInfoStream) Unique() *MemberInfoStream {
	return self.Distinct()
}
func (self *MemberInfoStream) Val() []MemberInfo {
	if self == nil {
		return []MemberInfo{}
	}
	return *self.Copy()
}
func (self *MemberInfoStream) While(fn func(MemberInfo, int) bool) *MemberInfoStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *MemberInfoStream) Where(fn func(MemberInfo) bool) *MemberInfoStream {
	result := MemberInfoStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *MemberInfoStream) WhereSlim(fn func(MemberInfo) bool) *MemberInfoStream {
	result := MemberInfoStreamOf()
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

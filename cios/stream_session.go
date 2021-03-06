/*
 * Collection utility of Session Struct
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

type SessionStream []Session

func SessionStreamOf(arg ...Session) SessionStream {
	return arg
}
func SessionStreamFrom(arg []Session) SessionStream {
	return arg
}
func CreateSessionStream(arg ...Session) *SessionStream {
	tmp := SessionStreamOf(arg...)
	return &tmp
}
func GenerateSessionStream(arg []Session) *SessionStream {
	tmp := SessionStreamFrom(arg)
	return &tmp
}

func (self *SessionStream) Add(arg Session) *SessionStream {
	return self.AddAll(arg)
}
func (self *SessionStream) AddAll(arg ...Session) *SessionStream {
	*self = append(*self, arg...)
	return self
}
func (self *SessionStream) AddSafe(arg *Session) *SessionStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *SessionStream) Aggregate(fn func(Session, Session) Session) *SessionStream {
	result := SessionStreamOf()
	self.ForEach(func(v Session, i int) {
		if i == 0 {
			result.Add(fn(Session{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *SessionStream) AllMatch(fn func(Session, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *SessionStream) AnyMatch(fn func(Session, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *SessionStream) Clone() *SessionStream {
	temp := make([]Session, self.Len())
	copy(temp, *self)
	return (*SessionStream)(&temp)
}
func (self *SessionStream) Copy() *SessionStream {
	return self.Clone()
}
func (self *SessionStream) Concat(arg []Session) *SessionStream {
	return self.AddAll(arg...)
}
func (self *SessionStream) Contains(arg Session) bool {
	return self.FindIndex(func(_arg Session, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *SessionStream) Clean() *SessionStream {
	*self = SessionStreamOf()
	return self
}
func (self *SessionStream) Delete(index int) *SessionStream {
	return self.DeleteRange(index, index)
}
func (self *SessionStream) DeleteRange(startIndex, endIndex int) *SessionStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *SessionStream) Distinct() *SessionStream {
	caches := map[string]bool{}
	result := SessionStreamOf()
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
func (self *SessionStream) Each(fn func(Session)) *SessionStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *SessionStream) EachRight(fn func(Session)) *SessionStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *SessionStream) Equals(arr []Session) bool {
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
func (self *SessionStream) Filter(fn func(Session, int) bool) *SessionStream {
	result := SessionStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *SessionStream) FilterSlim(fn func(Session, int) bool) *SessionStream {
	result := SessionStreamOf()
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
func (self *SessionStream) Find(fn func(Session, int) bool) *Session {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *SessionStream) FindOr(fn func(Session, int) bool, or Session) Session {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *SessionStream) FindIndex(fn func(Session, int) bool) int {
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
func (self *SessionStream) First() *Session {
	return self.Get(0)
}
func (self *SessionStream) FirstOr(arg Session) Session {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *SessionStream) ForEach(fn func(Session, int)) *SessionStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *SessionStream) ForEachRight(fn func(Session, int)) *SessionStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *SessionStream) GroupBy(fn func(Session, int) string) map[string][]Session {
	m := map[string][]Session{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *SessionStream) GroupByValues(fn func(Session, int) string) [][]Session {
	var tmp [][]Session
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *SessionStream) IndexOf(arg Session) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *SessionStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *SessionStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *SessionStream) Last() *Session {
	return self.Get(self.Len() - 1)
}
func (self *SessionStream) LastOr(arg Session) Session {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *SessionStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *SessionStream) Limit(limit int) *SessionStream {
	self.Slice(0, limit)
	return self
}

func (self *SessionStream) Map(fn func(Session, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SessionStream) Map2Int(fn func(Session, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SessionStream) Map2Int32(fn func(Session, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SessionStream) Map2Int64(fn func(Session, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SessionStream) Map2Float32(fn func(Session, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SessionStream) Map2Float64(fn func(Session, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SessionStream) Map2Bool(fn func(Session, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SessionStream) Map2Bytes(fn func(Session, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SessionStream) Map2String(fn func(Session, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SessionStream) Max(fn func(Session, int) float64) *Session {
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
func (self *SessionStream) Min(fn func(Session, int) float64) *Session {
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
func (self *SessionStream) NoneMatch(fn func(Session, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *SessionStream) Get(index int) *Session {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *SessionStream) GetOr(index int, arg Session) Session {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *SessionStream) Peek(fn func(*Session, int)) *SessionStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *SessionStream) Reduce(fn func(Session, Session, int) Session) *SessionStream {
	return self.ReduceInit(fn, Session{})
}
func (self *SessionStream) ReduceInit(fn func(Session, Session, int) Session, initialValue Session) *SessionStream {
	result := SessionStreamOf()
	self.ForEach(func(v Session, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *SessionStream) ReduceInterface(fn func(interface{}, Session, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(Session{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *SessionStream) ReduceString(fn func(string, Session, int) string) []string {
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
func (self *SessionStream) ReduceInt(fn func(int, Session, int) int) []int {
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
func (self *SessionStream) ReduceInt32(fn func(int32, Session, int) int32) []int32 {
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
func (self *SessionStream) ReduceInt64(fn func(int64, Session, int) int64) []int64 {
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
func (self *SessionStream) ReduceFloat32(fn func(float32, Session, int) float32) []float32 {
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
func (self *SessionStream) ReduceFloat64(fn func(float64, Session, int) float64) []float64 {
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
func (self *SessionStream) ReduceBool(fn func(bool, Session, int) bool) []bool {
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
func (self *SessionStream) Reverse() *SessionStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *SessionStream) Replace(fn func(Session, int) Session) *SessionStream {
	return self.ForEach(func(v Session, i int) { self.Set(i, fn(v, i)) })
}
func (self *SessionStream) Select(fn func(Session) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *SessionStream) Set(index int, val Session) *SessionStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *SessionStream) Skip(skip int) *SessionStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *SessionStream) SkippingEach(fn func(Session, int) int) *SessionStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *SessionStream) Slice(startIndex, n int) *SessionStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []Session{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *SessionStream) Sort(fn func(i, j int) bool) *SessionStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *SessionStream) Tail() *Session {
	return self.Last()
}
func (self *SessionStream) TailOr(arg Session) Session {
	return self.LastOr(arg)
}
func (self *SessionStream) ToList() []Session {
	return self.Val()
}
func (self *SessionStream) Unique() *SessionStream {
	return self.Distinct()
}
func (self *SessionStream) Val() []Session {
	if self == nil {
		return []Session{}
	}
	return *self.Copy()
}
func (self *SessionStream) While(fn func(Session, int) bool) *SessionStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *SessionStream) Where(fn func(Session) bool) *SessionStream {
	result := SessionStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *SessionStream) WhereSlim(fn func(Session) bool) *SessionStream {
	result := SessionStreamOf()
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

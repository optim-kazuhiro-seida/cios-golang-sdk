/*
 * Collection utility of Client Struct
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

type ClientStream []Client

func ClientStreamOf(arg ...Client) ClientStream {
	return arg
}
func ClientStreamFrom(arg []Client) ClientStream {
	return arg
}
func CreateClientStream(arg ...Client) *ClientStream {
	tmp := ClientStreamOf(arg...)
	return &tmp
}
func GenerateClientStream(arg []Client) *ClientStream {
	tmp := ClientStreamFrom(arg)
	return &tmp
}

func (self *ClientStream) Add(arg Client) *ClientStream {
	return self.AddAll(arg)
}
func (self *ClientStream) AddAll(arg ...Client) *ClientStream {
	*self = append(*self, arg...)
	return self
}
func (self *ClientStream) AddSafe(arg *Client) *ClientStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ClientStream) Aggregate(fn func(Client, Client) Client) *ClientStream {
	result := ClientStreamOf()
	self.ForEach(func(v Client, i int) {
		if i == 0 {
			result.Add(fn(Client{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ClientStream) AllMatch(fn func(Client, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ClientStream) AnyMatch(fn func(Client, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ClientStream) Clone() *ClientStream {
	temp := make([]Client, self.Len())
	copy(temp, *self)
	return (*ClientStream)(&temp)
}
func (self *ClientStream) Copy() *ClientStream {
	return self.Clone()
}
func (self *ClientStream) Concat(arg []Client) *ClientStream {
	return self.AddAll(arg...)
}
func (self *ClientStream) Contains(arg Client) bool {
	return self.FindIndex(func(_arg Client, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ClientStream) Clean() *ClientStream {
	*self = ClientStreamOf()
	return self
}
func (self *ClientStream) Delete(index int) *ClientStream {
	return self.DeleteRange(index, index)
}
func (self *ClientStream) DeleteRange(startIndex, endIndex int) *ClientStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ClientStream) Distinct() *ClientStream {
	caches := map[string]bool{}
	result := ClientStreamOf()
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
func (self *ClientStream) Each(fn func(Client)) *ClientStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ClientStream) EachRight(fn func(Client)) *ClientStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ClientStream) Equals(arr []Client) bool {
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
func (self *ClientStream) Filter(fn func(Client, int) bool) *ClientStream {
	result := ClientStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ClientStream) FilterSlim(fn func(Client, int) bool) *ClientStream {
	result := ClientStreamOf()
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
func (self *ClientStream) Find(fn func(Client, int) bool) *Client {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ClientStream) FindOr(fn func(Client, int) bool, or Client) Client {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ClientStream) FindIndex(fn func(Client, int) bool) int {
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
func (self *ClientStream) First() *Client {
	return self.Get(0)
}
func (self *ClientStream) FirstOr(arg Client) Client {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ClientStream) ForEach(fn func(Client, int)) *ClientStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ClientStream) ForEachRight(fn func(Client, int)) *ClientStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ClientStream) GroupBy(fn func(Client, int) string) map[string][]Client {
	m := map[string][]Client{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ClientStream) GroupByValues(fn func(Client, int) string) [][]Client {
	var tmp [][]Client
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ClientStream) IndexOf(arg Client) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ClientStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ClientStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ClientStream) Last() *Client {
	return self.Get(self.Len() - 1)
}
func (self *ClientStream) LastOr(arg Client) Client {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ClientStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ClientStream) Limit(limit int) *ClientStream {
	self.Slice(0, limit)
	return self
}

func (self *ClientStream) Map(fn func(Client, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ClientStream) Map2Int(fn func(Client, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ClientStream) Map2Int32(fn func(Client, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ClientStream) Map2Int64(fn func(Client, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ClientStream) Map2Float32(fn func(Client, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ClientStream) Map2Float64(fn func(Client, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ClientStream) Map2Bool(fn func(Client, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ClientStream) Map2Bytes(fn func(Client, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ClientStream) Map2String(fn func(Client, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ClientStream) Max(fn func(Client, int) float64) *Client {
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
func (self *ClientStream) Min(fn func(Client, int) float64) *Client {
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
func (self *ClientStream) NoneMatch(fn func(Client, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ClientStream) Get(index int) *Client {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ClientStream) GetOr(index int, arg Client) Client {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ClientStream) Peek(fn func(*Client, int)) *ClientStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *ClientStream) Reduce(fn func(Client, Client, int) Client) *ClientStream {
	return self.ReduceInit(fn, Client{})
}
func (self *ClientStream) ReduceInit(fn func(Client, Client, int) Client, initialValue Client) *ClientStream {
	result := ClientStreamOf()
	self.ForEach(func(v Client, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ClientStream) ReduceInterface(fn func(interface{}, Client, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(Client{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ClientStream) ReduceString(fn func(string, Client, int) string) []string {
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
func (self *ClientStream) ReduceInt(fn func(int, Client, int) int) []int {
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
func (self *ClientStream) ReduceInt32(fn func(int32, Client, int) int32) []int32 {
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
func (self *ClientStream) ReduceInt64(fn func(int64, Client, int) int64) []int64 {
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
func (self *ClientStream) ReduceFloat32(fn func(float32, Client, int) float32) []float32 {
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
func (self *ClientStream) ReduceFloat64(fn func(float64, Client, int) float64) []float64 {
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
func (self *ClientStream) ReduceBool(fn func(bool, Client, int) bool) []bool {
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
func (self *ClientStream) Reverse() *ClientStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ClientStream) Replace(fn func(Client, int) Client) *ClientStream {
	return self.ForEach(func(v Client, i int) { self.Set(i, fn(v, i)) })
}
func (self *ClientStream) Select(fn func(Client) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ClientStream) Set(index int, val Client) *ClientStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ClientStream) Skip(skip int) *ClientStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ClientStream) SkippingEach(fn func(Client, int) int) *ClientStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ClientStream) Slice(startIndex, n int) *ClientStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []Client{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ClientStream) Sort(fn func(i, j int) bool) *ClientStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ClientStream) Tail() *Client {
	return self.Last()
}
func (self *ClientStream) TailOr(arg Client) Client {
	return self.LastOr(arg)
}
func (self *ClientStream) ToList() []Client {
	return self.Val()
}
func (self *ClientStream) Unique() *ClientStream {
	return self.Distinct()
}
func (self *ClientStream) Val() []Client {
	if self == nil {
		return []Client{}
	}
	return *self.Copy()
}
func (self *ClientStream) While(fn func(Client, int) bool) *ClientStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ClientStream) Where(fn func(Client) bool) *ClientStream {
	result := ClientStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ClientStream) WhereSlim(fn func(Client) bool) *ClientStream {
	result := ClientStreamOf()
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
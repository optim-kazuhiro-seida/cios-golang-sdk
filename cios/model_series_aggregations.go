/*
 * Cios Openapi
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cios

import (
	"encoding/json"
)

// SeriesAggregations struct for SeriesAggregations
type SeriesAggregations struct {
	// intervalの開始時刻の配列
	Intervals []int64 `json:"intervals"`
	// intervalごとの時系列データ件数の配列
	TotalCounts []int64 `json:"total_counts"`
	// 集計結果の配列。リクエストで指定された集計対象・集計方法の順で並んでいる
	Aggregations *[]interface{} `json:"aggregations,omitempty"`
}

// NewSeriesAggregations instantiates a new SeriesAggregations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeriesAggregations(intervals []int64, totalCounts []int64, ) *SeriesAggregations {
	this := SeriesAggregations{}
	this.Intervals = intervals
	this.TotalCounts = totalCounts
	return &this
}

// NewSeriesAggregationsWithDefaults instantiates a new SeriesAggregations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeriesAggregationsWithDefaults() *SeriesAggregations {
	this := SeriesAggregations{}
	return &this
}

// GetIntervals returns the Intervals field value
func (o *SeriesAggregations) GetIntervals() []int64 {
	if o == nil  {
		var ret []int64
		return ret
	}

	return o.Intervals
}

// GetIntervalsOk returns a tuple with the Intervals field value
// and a boolean to check if the value has been set.
func (o *SeriesAggregations) GetIntervalsOk() (*[]int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Intervals, true
}

// SetIntervals sets field value
func (o *SeriesAggregations) SetIntervals(v []int64) {
	o.Intervals = v
}

// GetTotalCounts returns the TotalCounts field value
func (o *SeriesAggregations) GetTotalCounts() []int64 {
	if o == nil  {
		var ret []int64
		return ret
	}

	return o.TotalCounts
}

// GetTotalCountsOk returns a tuple with the TotalCounts field value
// and a boolean to check if the value has been set.
func (o *SeriesAggregations) GetTotalCountsOk() (*[]int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalCounts, true
}

// SetTotalCounts sets field value
func (o *SeriesAggregations) SetTotalCounts(v []int64) {
	o.TotalCounts = v
}

// GetAggregations returns the Aggregations field value if set, zero value otherwise.
func (o *SeriesAggregations) GetAggregations() []interface{} {
	if o == nil || o.Aggregations == nil {
		var ret []interface{}
		return ret
	}
	return *o.Aggregations
}

// GetAggregationsOk returns a tuple with the Aggregations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesAggregations) GetAggregationsOk() (*[]interface{}, bool) {
	if o == nil || o.Aggregations == nil {
		return nil, false
	}
	return o.Aggregations, true
}

// HasAggregations returns a boolean if a field has been set.
func (o *SeriesAggregations) HasAggregations() bool {
	if o != nil && o.Aggregations != nil {
		return true
	}

	return false
}

// SetAggregations gets a reference to the given []interface{} and assigns it to the Aggregations field.
func (o *SeriesAggregations) SetAggregations(v []interface{}) {
	o.Aggregations = &v
}

func (o SeriesAggregations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["intervals"] = o.Intervals
	}
	if true {
		toSerialize["total_counts"] = o.TotalCounts
	}
	if o.Aggregations != nil {
		toSerialize["aggregations"] = o.Aggregations
	}
	return json.Marshal(toSerialize)
}

type NullableSeriesAggregations struct {
	value *SeriesAggregations
	isSet bool
}

func (v NullableSeriesAggregations) Get() *SeriesAggregations {
	return v.value
}

func (v *NullableSeriesAggregations) Set(val *SeriesAggregations) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesAggregations) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesAggregations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesAggregations(val *SeriesAggregations) *NullableSeriesAggregations {
	return &NullableSeriesAggregations{value: val, isSet: true}
}

func (v NullableSeriesAggregations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesAggregations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



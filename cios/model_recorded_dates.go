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

// RecordedDates struct for RecordedDates
type RecordedDates struct {
	// データが存在する日付のリスト Ex.[1, 20, 22, 23]
	Dates []int64 `json:"dates"`
}

// NewRecordedDates instantiates a new RecordedDates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordedDates(dates []int64, ) *RecordedDates {
	this := RecordedDates{}
	this.Dates = dates
	return &this
}

// NewRecordedDatesWithDefaults instantiates a new RecordedDates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordedDatesWithDefaults() *RecordedDates {
	this := RecordedDates{}
	return &this
}

// GetDates returns the Dates field value
func (o *RecordedDates) GetDates() []int64 {
	if o == nil  {
		var ret []int64
		return ret
	}

	return o.Dates
}

// GetDatesOk returns a tuple with the Dates field value
// and a boolean to check if the value has been set.
func (o *RecordedDates) GetDatesOk() (*[]int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Dates, true
}

// SetDates sets field value
func (o *RecordedDates) SetDates(v []int64) {
	o.Dates = v
}

func (o RecordedDates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dates"] = o.Dates
	}
	return json.Marshal(toSerialize)
}

type NullableRecordedDates struct {
	value *RecordedDates
	isSet bool
}

func (v NullableRecordedDates) Get() *RecordedDates {
	return v.value
}

func (v *NullableRecordedDates) Set(val *RecordedDates) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordedDates) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordedDates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordedDates(val *RecordedDates) *NullableRecordedDates {
	return &NullableRecordedDates{value: val, isSet: true}
}

func (v NullableRecordedDates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordedDates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



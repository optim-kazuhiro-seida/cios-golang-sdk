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

// SingleCircle struct for SingleCircle
type SingleCircle struct {
	Circle *Circle `json:"circle,omitempty"`
}

// NewSingleCircle instantiates a new SingleCircle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleCircle() *SingleCircle {
	this := SingleCircle{}
	return &this
}

// NewSingleCircleWithDefaults instantiates a new SingleCircle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleCircleWithDefaults() *SingleCircle {
	this := SingleCircle{}
	return &this
}

// GetCircle returns the Circle field value if set, zero value otherwise.
func (o *SingleCircle) GetCircle() Circle {
	if o == nil || o.Circle == nil {
		var ret Circle
		return ret
	}
	return *o.Circle
}

// GetCircleOk returns a tuple with the Circle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleCircle) GetCircleOk() (*Circle, bool) {
	if o == nil || o.Circle == nil {
		return nil, false
	}
	return o.Circle, true
}

// HasCircle returns a boolean if a field has been set.
func (o *SingleCircle) HasCircle() bool {
	if o != nil && o.Circle != nil {
		return true
	}

	return false
}

// SetCircle gets a reference to the given Circle and assigns it to the Circle field.
func (o *SingleCircle) SetCircle(v Circle) {
	o.Circle = &v
}

func (o SingleCircle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Circle != nil {
		toSerialize["circle"] = o.Circle
	}
	return json.Marshal(toSerialize)
}

type NullableSingleCircle struct {
	value *SingleCircle
	isSet bool
}

func (v NullableSingleCircle) Get() *SingleCircle {
	return v.value
}

func (v *NullableSingleCircle) Set(val *SingleCircle) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleCircle) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleCircle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleCircle(val *SingleCircle) *NullableSingleCircle {
	return &NullableSingleCircle{value: val, isSet: true}
}

func (v NullableSingleCircle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleCircle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



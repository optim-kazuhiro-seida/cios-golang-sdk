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

// MultipleRoute struct for MultipleRoute
type MultipleRoute struct {
	Total *int64 `json:"total,omitempty"`
	Routes *[]Route `json:"routes,omitempty"`
}

// NewMultipleRoute instantiates a new MultipleRoute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultipleRoute() *MultipleRoute {
	this := MultipleRoute{}
	return &this
}

// NewMultipleRouteWithDefaults instantiates a new MultipleRoute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultipleRouteWithDefaults() *MultipleRoute {
	this := MultipleRoute{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *MultipleRoute) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipleRoute) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *MultipleRoute) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *MultipleRoute) SetTotal(v int64) {
	o.Total = &v
}

// GetRoutes returns the Routes field value if set, zero value otherwise.
func (o *MultipleRoute) GetRoutes() []Route {
	if o == nil || o.Routes == nil {
		var ret []Route
		return ret
	}
	return *o.Routes
}

// GetRoutesOk returns a tuple with the Routes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipleRoute) GetRoutesOk() (*[]Route, bool) {
	if o == nil || o.Routes == nil {
		return nil, false
	}
	return o.Routes, true
}

// HasRoutes returns a boolean if a field has been set.
func (o *MultipleRoute) HasRoutes() bool {
	if o != nil && o.Routes != nil {
		return true
	}

	return false
}

// SetRoutes gets a reference to the given []Route and assigns it to the Routes field.
func (o *MultipleRoute) SetRoutes(v []Route) {
	o.Routes = &v
}

func (o MultipleRoute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.Routes != nil {
		toSerialize["routes"] = o.Routes
	}
	return json.Marshal(toSerialize)
}

type NullableMultipleRoute struct {
	value *MultipleRoute
	isSet bool
}

func (v NullableMultipleRoute) Get() *MultipleRoute {
	return v.value
}

func (v *NullableMultipleRoute) Set(val *MultipleRoute) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipleRoute) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipleRoute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipleRoute(val *MultipleRoute) *NullableMultipleRoute {
	return &NullableMultipleRoute{value: val, isSet: true}
}

func (v NullableMultipleRoute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipleRoute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// CreatedByUser struct for CreatedByUser
type CreatedByUser struct {
	Id *string `json:"id,omitempty"`
}

// NewCreatedByUser instantiates a new CreatedByUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatedByUser() *CreatedByUser {
	this := CreatedByUser{}
	return &this
}

// NewCreatedByUserWithDefaults instantiates a new CreatedByUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatedByUserWithDefaults() *CreatedByUser {
	this := CreatedByUser{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreatedByUser) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedByUser) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreatedByUser) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreatedByUser) SetId(v string) {
	o.Id = &v
}

func (o CreatedByUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableCreatedByUser struct {
	value *CreatedByUser
	isSet bool
}

func (v NullableCreatedByUser) Get() *CreatedByUser {
	return v.value
}

func (v *NullableCreatedByUser) Set(val *CreatedByUser) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatedByUser) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatedByUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatedByUser(val *CreatedByUser) *NullableCreatedByUser {
	return &NullableCreatedByUser{value: val, isSet: true}
}

func (v NullableCreatedByUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatedByUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



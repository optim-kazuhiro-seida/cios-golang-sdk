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

// NodeName struct for NodeName
type NodeName struct {
	Name string `json:"name"`
}

// NewNodeName instantiates a new NodeName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeName(name string, ) *NodeName {
	this := NodeName{}
	this.Name = name
	return &this
}

// NewNodeNameWithDefaults instantiates a new NodeName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeNameWithDefaults() *NodeName {
	this := NodeName{}
	return &this
}

// GetName returns the Name field value
func (o *NodeName) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NodeName) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NodeName) SetName(v string) {
	o.Name = v
}

func (o NodeName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableNodeName struct {
	value *NodeName
	isSet bool
}

func (v NullableNodeName) Get() *NodeName {
	return v.value
}

func (v *NullableNodeName) Set(val *NodeName) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeName) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeName(val *NodeName) *NullableNodeName {
	return &NullableNodeName{value: val, isSet: true}
}

func (v NullableNodeName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



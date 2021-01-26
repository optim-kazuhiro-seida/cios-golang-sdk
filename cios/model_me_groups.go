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

// MeGroups struct for MeGroups
type MeGroups struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	CorporationId NullableString `json:"corporation_id"`
}

// NewMeGroups instantiates a new MeGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeGroups(id string, name string, type_ string, corporationId NullableString, ) *MeGroups {
	this := MeGroups{}
	this.Id = id
	this.Name = name
	this.Type = type_
	this.CorporationId = corporationId
	return &this
}

// NewMeGroupsWithDefaults instantiates a new MeGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeGroupsWithDefaults() *MeGroups {
	this := MeGroups{}
	return &this
}

// GetId returns the Id field value
func (o *MeGroups) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MeGroups) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MeGroups) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *MeGroups) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MeGroups) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MeGroups) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *MeGroups) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MeGroups) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MeGroups) SetType(v string) {
	o.Type = v
}

// GetCorporationId returns the CorporationId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MeGroups) GetCorporationId() string {
	if o == nil || o.CorporationId.Get() == nil {
		var ret string
		return ret
	}

	return *o.CorporationId.Get()
}

// GetCorporationIdOk returns a tuple with the CorporationId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MeGroups) GetCorporationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CorporationId.Get(), o.CorporationId.IsSet()
}

// SetCorporationId sets field value
func (o *MeGroups) SetCorporationId(v string) {
	o.CorporationId.Set(&v)
}

func (o MeGroups) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["corporation_id"] = o.CorporationId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMeGroups struct {
	value *MeGroups
	isSet bool
}

func (v NullableMeGroups) Get() *MeGroups {
	return v.value
}

func (v *NullableMeGroups) Set(val *MeGroups) {
	v.value = val
	v.isSet = true
}

func (v NullableMeGroups) IsSet() bool {
	return v.isSet
}

func (v *NullableMeGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeGroups(val *MeGroups) *NullableMeGroups {
	return &NullableMeGroups{value: val, isSet: true}
}

func (v NullableMeGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


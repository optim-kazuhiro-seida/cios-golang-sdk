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

// Point struct for Point
type Point struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Location Location `json:"location"`
	Altitude *float32 `json:"altitude,omitempty"`
	Labels *[]Label `json:"labels,omitempty"`
	Description *string `json:"description,omitempty"`
	DisplayInfo *[]DisplayInfo `json:"display_info,omitempty"`
	ResourceOwnerId string `json:"resource_owner_id"`
}

// NewPoint instantiates a new Point object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoint(id string, name string, location Location, resourceOwnerId string, ) *Point {
	this := Point{}
	this.Id = id
	this.Name = name
	this.Location = location
	this.ResourceOwnerId = resourceOwnerId
	return &this
}

// NewPointWithDefaults instantiates a new Point object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPointWithDefaults() *Point {
	this := Point{}
	return &this
}

// GetId returns the Id field value
func (o *Point) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Point) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Point) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Point) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Point) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Point) SetName(v string) {
	o.Name = v
}

// GetLocation returns the Location field value
func (o *Point) GetLocation() Location {
	if o == nil  {
		var ret Location
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *Point) GetLocationOk() (*Location, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *Point) SetLocation(v Location) {
	o.Location = v
}

// GetAltitude returns the Altitude field value if set, zero value otherwise.
func (o *Point) GetAltitude() float32 {
	if o == nil || o.Altitude == nil {
		var ret float32
		return ret
	}
	return *o.Altitude
}

// GetAltitudeOk returns a tuple with the Altitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Point) GetAltitudeOk() (*float32, bool) {
	if o == nil || o.Altitude == nil {
		return nil, false
	}
	return o.Altitude, true
}

// HasAltitude returns a boolean if a field has been set.
func (o *Point) HasAltitude() bool {
	if o != nil && o.Altitude != nil {
		return true
	}

	return false
}

// SetAltitude gets a reference to the given float32 and assigns it to the Altitude field.
func (o *Point) SetAltitude(v float32) {
	o.Altitude = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *Point) GetLabels() []Label {
	if o == nil || o.Labels == nil {
		var ret []Label
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Point) GetLabelsOk() (*[]Label, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *Point) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []Label and assigns it to the Labels field.
func (o *Point) SetLabels(v []Label) {
	o.Labels = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Point) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Point) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Point) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Point) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayInfo returns the DisplayInfo field value if set, zero value otherwise.
func (o *Point) GetDisplayInfo() []DisplayInfo {
	if o == nil || o.DisplayInfo == nil {
		var ret []DisplayInfo
		return ret
	}
	return *o.DisplayInfo
}

// GetDisplayInfoOk returns a tuple with the DisplayInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Point) GetDisplayInfoOk() (*[]DisplayInfo, bool) {
	if o == nil || o.DisplayInfo == nil {
		return nil, false
	}
	return o.DisplayInfo, true
}

// HasDisplayInfo returns a boolean if a field has been set.
func (o *Point) HasDisplayInfo() bool {
	if o != nil && o.DisplayInfo != nil {
		return true
	}

	return false
}

// SetDisplayInfo gets a reference to the given []DisplayInfo and assigns it to the DisplayInfo field.
func (o *Point) SetDisplayInfo(v []DisplayInfo) {
	o.DisplayInfo = &v
}

// GetResourceOwnerId returns the ResourceOwnerId field value
func (o *Point) GetResourceOwnerId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ResourceOwnerId
}

// GetResourceOwnerIdOk returns a tuple with the ResourceOwnerId field value
// and a boolean to check if the value has been set.
func (o *Point) GetResourceOwnerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ResourceOwnerId, true
}

// SetResourceOwnerId sets field value
func (o *Point) SetResourceOwnerId(v string) {
	o.ResourceOwnerId = v
}

func (o Point) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["location"] = o.Location
	}
	if o.Altitude != nil {
		toSerialize["altitude"] = o.Altitude
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayInfo != nil {
		toSerialize["display_info"] = o.DisplayInfo
	}
	if true {
		toSerialize["resource_owner_id"] = o.ResourceOwnerId
	}
	return json.Marshal(toSerialize)
}

type NullablePoint struct {
	value *Point
	isSet bool
}

func (v NullablePoint) Get() *Point {
	return v.value
}

func (v *NullablePoint) Set(val *Point) {
	v.value = val
	v.isSet = true
}

func (v NullablePoint) IsSet() bool {
	return v.isSet
}

func (v *NullablePoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoint(val *Point) *NullablePoint {
	return &NullablePoint{value: val, isSet: true}
}

func (v NullablePoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


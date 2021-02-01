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

// LifeCycleRequest struct for LifeCycleRequest
type LifeCycleRequest struct {
	EventKind string `json:"event_kind"`
	EventMode string `json:"event_mode"`
	EventType string `json:"event_type"`
	BeforeId *string `json:"before_id,omitempty"`
	BeforeComponent NullableDeviceEntitiesComponent `json:"before_component,omitempty"`
	AfterId *string `json:"after_id,omitempty"`
	AfterComponent NullableDeviceEntitiesComponent `json:"after_component,omitempty"`
	// ナノ秒
	EventAt string `json:"event_at"`
	Note *string `json:"note,omitempty"`
	ResourceOwnerId *string `json:"resource_owner_id,omitempty"`
}

// NewLifeCycleRequest instantiates a new LifeCycleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLifeCycleRequest(eventKind string, eventMode string, eventType string, eventAt string, ) *LifeCycleRequest {
	this := LifeCycleRequest{}
	this.EventKind = eventKind
	this.EventMode = eventMode
	this.EventType = eventType
	this.EventAt = eventAt
	return &this
}

// NewLifeCycleRequestWithDefaults instantiates a new LifeCycleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLifeCycleRequestWithDefaults() *LifeCycleRequest {
	this := LifeCycleRequest{}
	return &this
}

// GetEventKind returns the EventKind field value
func (o *LifeCycleRequest) GetEventKind() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.EventKind
}

// GetEventKindOk returns a tuple with the EventKind field value
// and a boolean to check if the value has been set.
func (o *LifeCycleRequest) GetEventKindOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventKind, true
}

// SetEventKind sets field value
func (o *LifeCycleRequest) SetEventKind(v string) {
	o.EventKind = v
}

// GetEventMode returns the EventMode field value
func (o *LifeCycleRequest) GetEventMode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.EventMode
}

// GetEventModeOk returns a tuple with the EventMode field value
// and a boolean to check if the value has been set.
func (o *LifeCycleRequest) GetEventModeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventMode, true
}

// SetEventMode sets field value
func (o *LifeCycleRequest) SetEventMode(v string) {
	o.EventMode = v
}

// GetEventType returns the EventType field value
func (o *LifeCycleRequest) GetEventType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *LifeCycleRequest) GetEventTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *LifeCycleRequest) SetEventType(v string) {
	o.EventType = v
}

// GetBeforeId returns the BeforeId field value if set, zero value otherwise.
func (o *LifeCycleRequest) GetBeforeId() string {
	if o == nil || o.BeforeId == nil {
		var ret string
		return ret
	}
	return *o.BeforeId
}

// GetBeforeIdOk returns a tuple with the BeforeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifeCycleRequest) GetBeforeIdOk() (*string, bool) {
	if o == nil || o.BeforeId == nil {
		return nil, false
	}
	return o.BeforeId, true
}

// HasBeforeId returns a boolean if a field has been set.
func (o *LifeCycleRequest) HasBeforeId() bool {
	if o != nil && o.BeforeId != nil {
		return true
	}

	return false
}

// SetBeforeId gets a reference to the given string and assigns it to the BeforeId field.
func (o *LifeCycleRequest) SetBeforeId(v string) {
	o.BeforeId = &v
}

// GetBeforeComponent returns the BeforeComponent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LifeCycleRequest) GetBeforeComponent() DeviceEntitiesComponent {
	if o == nil || o.BeforeComponent.Get() == nil {
		var ret DeviceEntitiesComponent
		return ret
	}
	return *o.BeforeComponent.Get()
}

// GetBeforeComponentOk returns a tuple with the BeforeComponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LifeCycleRequest) GetBeforeComponentOk() (*DeviceEntitiesComponent, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BeforeComponent.Get(), o.BeforeComponent.IsSet()
}

// HasBeforeComponent returns a boolean if a field has been set.
func (o *LifeCycleRequest) HasBeforeComponent() bool {
	if o != nil && o.BeforeComponent.IsSet() {
		return true
	}

	return false
}

// SetBeforeComponent gets a reference to the given NullableDeviceEntitiesComponent and assigns it to the BeforeComponent field.
func (o *LifeCycleRequest) SetBeforeComponent(v DeviceEntitiesComponent) {
	o.BeforeComponent.Set(&v)
}
// SetBeforeComponentNil sets the value for BeforeComponent to be an explicit nil
func (o *LifeCycleRequest) SetBeforeComponentNil() {
	o.BeforeComponent.Set(nil)
}

// UnsetBeforeComponent ensures that no value is present for BeforeComponent, not even an explicit nil
func (o *LifeCycleRequest) UnsetBeforeComponent() {
	o.BeforeComponent.Unset()
}

// GetAfterId returns the AfterId field value if set, zero value otherwise.
func (o *LifeCycleRequest) GetAfterId() string {
	if o == nil || o.AfterId == nil {
		var ret string
		return ret
	}
	return *o.AfterId
}

// GetAfterIdOk returns a tuple with the AfterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifeCycleRequest) GetAfterIdOk() (*string, bool) {
	if o == nil || o.AfterId == nil {
		return nil, false
	}
	return o.AfterId, true
}

// HasAfterId returns a boolean if a field has been set.
func (o *LifeCycleRequest) HasAfterId() bool {
	if o != nil && o.AfterId != nil {
		return true
	}

	return false
}

// SetAfterId gets a reference to the given string and assigns it to the AfterId field.
func (o *LifeCycleRequest) SetAfterId(v string) {
	o.AfterId = &v
}

// GetAfterComponent returns the AfterComponent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LifeCycleRequest) GetAfterComponent() DeviceEntitiesComponent {
	if o == nil || o.AfterComponent.Get() == nil {
		var ret DeviceEntitiesComponent
		return ret
	}
	return *o.AfterComponent.Get()
}

// GetAfterComponentOk returns a tuple with the AfterComponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LifeCycleRequest) GetAfterComponentOk() (*DeviceEntitiesComponent, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AfterComponent.Get(), o.AfterComponent.IsSet()
}

// HasAfterComponent returns a boolean if a field has been set.
func (o *LifeCycleRequest) HasAfterComponent() bool {
	if o != nil && o.AfterComponent.IsSet() {
		return true
	}

	return false
}

// SetAfterComponent gets a reference to the given NullableDeviceEntitiesComponent and assigns it to the AfterComponent field.
func (o *LifeCycleRequest) SetAfterComponent(v DeviceEntitiesComponent) {
	o.AfterComponent.Set(&v)
}
// SetAfterComponentNil sets the value for AfterComponent to be an explicit nil
func (o *LifeCycleRequest) SetAfterComponentNil() {
	o.AfterComponent.Set(nil)
}

// UnsetAfterComponent ensures that no value is present for AfterComponent, not even an explicit nil
func (o *LifeCycleRequest) UnsetAfterComponent() {
	o.AfterComponent.Unset()
}

// GetEventAt returns the EventAt field value
func (o *LifeCycleRequest) GetEventAt() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.EventAt
}

// GetEventAtOk returns a tuple with the EventAt field value
// and a boolean to check if the value has been set.
func (o *LifeCycleRequest) GetEventAtOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventAt, true
}

// SetEventAt sets field value
func (o *LifeCycleRequest) SetEventAt(v string) {
	o.EventAt = v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *LifeCycleRequest) GetNote() string {
	if o == nil || o.Note == nil {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifeCycleRequest) GetNoteOk() (*string, bool) {
	if o == nil || o.Note == nil {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *LifeCycleRequest) HasNote() bool {
	if o != nil && o.Note != nil {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *LifeCycleRequest) SetNote(v string) {
	o.Note = &v
}

// GetResourceOwnerId returns the ResourceOwnerId field value if set, zero value otherwise.
func (o *LifeCycleRequest) GetResourceOwnerId() string {
	if o == nil || o.ResourceOwnerId == nil {
		var ret string
		return ret
	}
	return *o.ResourceOwnerId
}

// GetResourceOwnerIdOk returns a tuple with the ResourceOwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifeCycleRequest) GetResourceOwnerIdOk() (*string, bool) {
	if o == nil || o.ResourceOwnerId == nil {
		return nil, false
	}
	return o.ResourceOwnerId, true
}

// HasResourceOwnerId returns a boolean if a field has been set.
func (o *LifeCycleRequest) HasResourceOwnerId() bool {
	if o != nil && o.ResourceOwnerId != nil {
		return true
	}

	return false
}

// SetResourceOwnerId gets a reference to the given string and assigns it to the ResourceOwnerId field.
func (o *LifeCycleRequest) SetResourceOwnerId(v string) {
	o.ResourceOwnerId = &v
}

func (o LifeCycleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["event_kind"] = o.EventKind
	}
	if true {
		toSerialize["event_mode"] = o.EventMode
	}
	if true {
		toSerialize["event_type"] = o.EventType
	}
	if o.BeforeId != nil {
		toSerialize["before_id"] = o.BeforeId
	}
	if o.BeforeComponent.IsSet() {
		toSerialize["before_component"] = o.BeforeComponent.Get()
	}
	if o.AfterId != nil {
		toSerialize["after_id"] = o.AfterId
	}
	if o.AfterComponent.IsSet() {
		toSerialize["after_component"] = o.AfterComponent.Get()
	}
	if true {
		toSerialize["event_at"] = o.EventAt
	}
	if o.Note != nil {
		toSerialize["note"] = o.Note
	}
	if o.ResourceOwnerId != nil {
		toSerialize["resource_owner_id"] = o.ResourceOwnerId
	}
	return json.Marshal(toSerialize)
}

type NullableLifeCycleRequest struct {
	value *LifeCycleRequest
	isSet bool
}

func (v NullableLifeCycleRequest) Get() *LifeCycleRequest {
	return v.value
}

func (v *NullableLifeCycleRequest) Set(val *LifeCycleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLifeCycleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLifeCycleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLifeCycleRequest(val *LifeCycleRequest) *NullableLifeCycleRequest {
	return &NullableLifeCycleRequest{value: val, isSet: true}
}

func (v NullableLifeCycleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLifeCycleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// Channel struct for Channel
type Channel struct {
	Id string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	ResourceOwnerId string `json:"resource_owner_id"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	ChannelProtocols *[]ChannelProtocol `json:"channel_protocols,omitempty"`
	DisplayInfo *[]DisplayInfo `json:"display_info,omitempty"`
	Labels *[]Label `json:"labels,omitempty"`
	MessagingConfig *MessagingConfig `json:"messaging_config,omitempty"`
	DatastoreConfig *DataStoreConfig `json:"datastore_config,omitempty"`
}

// NewChannel instantiates a new Channel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannel(id string, createdAt string, updatedAt string, resourceOwnerId string, name string, ) *Channel {
	this := Channel{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.ResourceOwnerId = resourceOwnerId
	this.Name = name
	return &this
}

// NewChannelWithDefaults instantiates a new Channel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelWithDefaults() *Channel {
	this := Channel{}
	return &this
}

// GetId returns the Id field value
func (o *Channel) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Channel) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Channel) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Channel) GetCreatedAt() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Channel) GetCreatedAtOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Channel) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Channel) GetUpdatedAt() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Channel) GetUpdatedAtOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Channel) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetResourceOwnerId returns the ResourceOwnerId field value
func (o *Channel) GetResourceOwnerId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ResourceOwnerId
}

// GetResourceOwnerIdOk returns a tuple with the ResourceOwnerId field value
// and a boolean to check if the value has been set.
func (o *Channel) GetResourceOwnerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ResourceOwnerId, true
}

// SetResourceOwnerId sets field value
func (o *Channel) SetResourceOwnerId(v string) {
	o.ResourceOwnerId = v
}

// GetName returns the Name field value
func (o *Channel) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Channel) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Channel) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Channel) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Channel) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Channel) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Channel) SetDescription(v string) {
	o.Description = &v
}

// GetChannelProtocols returns the ChannelProtocols field value if set, zero value otherwise.
func (o *Channel) GetChannelProtocols() []ChannelProtocol {
	if o == nil || o.ChannelProtocols == nil {
		var ret []ChannelProtocol
		return ret
	}
	return *o.ChannelProtocols
}

// GetChannelProtocolsOk returns a tuple with the ChannelProtocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Channel) GetChannelProtocolsOk() (*[]ChannelProtocol, bool) {
	if o == nil || o.ChannelProtocols == nil {
		return nil, false
	}
	return o.ChannelProtocols, true
}

// HasChannelProtocols returns a boolean if a field has been set.
func (o *Channel) HasChannelProtocols() bool {
	if o != nil && o.ChannelProtocols != nil {
		return true
	}

	return false
}

// SetChannelProtocols gets a reference to the given []ChannelProtocol and assigns it to the ChannelProtocols field.
func (o *Channel) SetChannelProtocols(v []ChannelProtocol) {
	o.ChannelProtocols = &v
}

// GetDisplayInfo returns the DisplayInfo field value if set, zero value otherwise.
func (o *Channel) GetDisplayInfo() []DisplayInfo {
	if o == nil || o.DisplayInfo == nil {
		var ret []DisplayInfo
		return ret
	}
	return *o.DisplayInfo
}

// GetDisplayInfoOk returns a tuple with the DisplayInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Channel) GetDisplayInfoOk() (*[]DisplayInfo, bool) {
	if o == nil || o.DisplayInfo == nil {
		return nil, false
	}
	return o.DisplayInfo, true
}

// HasDisplayInfo returns a boolean if a field has been set.
func (o *Channel) HasDisplayInfo() bool {
	if o != nil && o.DisplayInfo != nil {
		return true
	}

	return false
}

// SetDisplayInfo gets a reference to the given []DisplayInfo and assigns it to the DisplayInfo field.
func (o *Channel) SetDisplayInfo(v []DisplayInfo) {
	o.DisplayInfo = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *Channel) GetLabels() []Label {
	if o == nil || o.Labels == nil {
		var ret []Label
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Channel) GetLabelsOk() (*[]Label, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *Channel) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []Label and assigns it to the Labels field.
func (o *Channel) SetLabels(v []Label) {
	o.Labels = &v
}

// GetMessagingConfig returns the MessagingConfig field value if set, zero value otherwise.
func (o *Channel) GetMessagingConfig() MessagingConfig {
	if o == nil || o.MessagingConfig == nil {
		var ret MessagingConfig
		return ret
	}
	return *o.MessagingConfig
}

// GetMessagingConfigOk returns a tuple with the MessagingConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Channel) GetMessagingConfigOk() (*MessagingConfig, bool) {
	if o == nil || o.MessagingConfig == nil {
		return nil, false
	}
	return o.MessagingConfig, true
}

// HasMessagingConfig returns a boolean if a field has been set.
func (o *Channel) HasMessagingConfig() bool {
	if o != nil && o.MessagingConfig != nil {
		return true
	}

	return false
}

// SetMessagingConfig gets a reference to the given MessagingConfig and assigns it to the MessagingConfig field.
func (o *Channel) SetMessagingConfig(v MessagingConfig) {
	o.MessagingConfig = &v
}

// GetDatastoreConfig returns the DatastoreConfig field value if set, zero value otherwise.
func (o *Channel) GetDatastoreConfig() DataStoreConfig {
	if o == nil || o.DatastoreConfig == nil {
		var ret DataStoreConfig
		return ret
	}
	return *o.DatastoreConfig
}

// GetDatastoreConfigOk returns a tuple with the DatastoreConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Channel) GetDatastoreConfigOk() (*DataStoreConfig, bool) {
	if o == nil || o.DatastoreConfig == nil {
		return nil, false
	}
	return o.DatastoreConfig, true
}

// HasDatastoreConfig returns a boolean if a field has been set.
func (o *Channel) HasDatastoreConfig() bool {
	if o != nil && o.DatastoreConfig != nil {
		return true
	}

	return false
}

// SetDatastoreConfig gets a reference to the given DataStoreConfig and assigns it to the DatastoreConfig field.
func (o *Channel) SetDatastoreConfig(v DataStoreConfig) {
	o.DatastoreConfig = &v
}

func (o Channel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["resource_owner_id"] = o.ResourceOwnerId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ChannelProtocols != nil {
		toSerialize["channel_protocols"] = o.ChannelProtocols
	}
	if o.DisplayInfo != nil {
		toSerialize["display_info"] = o.DisplayInfo
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.MessagingConfig != nil {
		toSerialize["messaging_config"] = o.MessagingConfig
	}
	if o.DatastoreConfig != nil {
		toSerialize["datastore_config"] = o.DatastoreConfig
	}
	return json.Marshal(toSerialize)
}

type NullableChannel struct {
	value *Channel
	isSet bool
}

func (v NullableChannel) Get() *Channel {
	return v.value
}

func (v *NullableChannel) Set(val *Channel) {
	v.value = val
	v.isSet = true
}

func (v NullableChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannel(val *Channel) *NullableChannel {
	return &NullableChannel{value: val, isSet: true}
}

func (v NullableChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



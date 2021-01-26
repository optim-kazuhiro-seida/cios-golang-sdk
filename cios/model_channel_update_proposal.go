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

// ChannelUpdateProposal struct for ChannelUpdateProposal
type ChannelUpdateProposal struct {
	ChannelProtocols *[]ChannelProtocol `json:"channel_protocols,omitempty"`
	DisplayInfo []DisplayInfo `json:"display_info"`
	Labels *[]Label `json:"labels,omitempty"`
	MessagingConfig *MessagingConfig `json:"messaging_config,omitempty"`
	DatastoreConfig *DataStoreConfig `json:"datastore_config,omitempty"`
}

// NewChannelUpdateProposal instantiates a new ChannelUpdateProposal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelUpdateProposal(displayInfo []DisplayInfo, ) *ChannelUpdateProposal {
	this := ChannelUpdateProposal{}
	this.DisplayInfo = displayInfo
	return &this
}

// NewChannelUpdateProposalWithDefaults instantiates a new ChannelUpdateProposal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelUpdateProposalWithDefaults() *ChannelUpdateProposal {
	this := ChannelUpdateProposal{}
	return &this
}

// GetChannelProtocols returns the ChannelProtocols field value if set, zero value otherwise.
func (o *ChannelUpdateProposal) GetChannelProtocols() []ChannelProtocol {
	if o == nil || o.ChannelProtocols == nil {
		var ret []ChannelProtocol
		return ret
	}
	return *o.ChannelProtocols
}

// GetChannelProtocolsOk returns a tuple with the ChannelProtocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelUpdateProposal) GetChannelProtocolsOk() (*[]ChannelProtocol, bool) {
	if o == nil || o.ChannelProtocols == nil {
		return nil, false
	}
	return o.ChannelProtocols, true
}

// HasChannelProtocols returns a boolean if a field has been set.
func (o *ChannelUpdateProposal) HasChannelProtocols() bool {
	if o != nil && o.ChannelProtocols != nil {
		return true
	}

	return false
}

// SetChannelProtocols gets a reference to the given []ChannelProtocol and assigns it to the ChannelProtocols field.
func (o *ChannelUpdateProposal) SetChannelProtocols(v []ChannelProtocol) {
	o.ChannelProtocols = &v
}

// GetDisplayInfo returns the DisplayInfo field value
func (o *ChannelUpdateProposal) GetDisplayInfo() []DisplayInfo {
	if o == nil  {
		var ret []DisplayInfo
		return ret
	}

	return o.DisplayInfo
}

// GetDisplayInfoOk returns a tuple with the DisplayInfo field value
// and a boolean to check if the value has been set.
func (o *ChannelUpdateProposal) GetDisplayInfoOk() (*[]DisplayInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DisplayInfo, true
}

// SetDisplayInfo sets field value
func (o *ChannelUpdateProposal) SetDisplayInfo(v []DisplayInfo) {
	o.DisplayInfo = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ChannelUpdateProposal) GetLabels() []Label {
	if o == nil || o.Labels == nil {
		var ret []Label
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelUpdateProposal) GetLabelsOk() (*[]Label, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ChannelUpdateProposal) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []Label and assigns it to the Labels field.
func (o *ChannelUpdateProposal) SetLabels(v []Label) {
	o.Labels = &v
}

// GetMessagingConfig returns the MessagingConfig field value if set, zero value otherwise.
func (o *ChannelUpdateProposal) GetMessagingConfig() MessagingConfig {
	if o == nil || o.MessagingConfig == nil {
		var ret MessagingConfig
		return ret
	}
	return *o.MessagingConfig
}

// GetMessagingConfigOk returns a tuple with the MessagingConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelUpdateProposal) GetMessagingConfigOk() (*MessagingConfig, bool) {
	if o == nil || o.MessagingConfig == nil {
		return nil, false
	}
	return o.MessagingConfig, true
}

// HasMessagingConfig returns a boolean if a field has been set.
func (o *ChannelUpdateProposal) HasMessagingConfig() bool {
	if o != nil && o.MessagingConfig != nil {
		return true
	}

	return false
}

// SetMessagingConfig gets a reference to the given MessagingConfig and assigns it to the MessagingConfig field.
func (o *ChannelUpdateProposal) SetMessagingConfig(v MessagingConfig) {
	o.MessagingConfig = &v
}

// GetDatastoreConfig returns the DatastoreConfig field value if set, zero value otherwise.
func (o *ChannelUpdateProposal) GetDatastoreConfig() DataStoreConfig {
	if o == nil || o.DatastoreConfig == nil {
		var ret DataStoreConfig
		return ret
	}
	return *o.DatastoreConfig
}

// GetDatastoreConfigOk returns a tuple with the DatastoreConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelUpdateProposal) GetDatastoreConfigOk() (*DataStoreConfig, bool) {
	if o == nil || o.DatastoreConfig == nil {
		return nil, false
	}
	return o.DatastoreConfig, true
}

// HasDatastoreConfig returns a boolean if a field has been set.
func (o *ChannelUpdateProposal) HasDatastoreConfig() bool {
	if o != nil && o.DatastoreConfig != nil {
		return true
	}

	return false
}

// SetDatastoreConfig gets a reference to the given DataStoreConfig and assigns it to the DatastoreConfig field.
func (o *ChannelUpdateProposal) SetDatastoreConfig(v DataStoreConfig) {
	o.DatastoreConfig = &v
}

func (o ChannelUpdateProposal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChannelProtocols != nil {
		toSerialize["channel_protocols"] = o.ChannelProtocols
	}
	if true {
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

type NullableChannelUpdateProposal struct {
	value *ChannelUpdateProposal
	isSet bool
}

func (v NullableChannelUpdateProposal) Get() *ChannelUpdateProposal {
	return v.value
}

func (v *NullableChannelUpdateProposal) Set(val *ChannelUpdateProposal) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelUpdateProposal) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelUpdateProposal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelUpdateProposal(val *ChannelUpdateProposal) *NullableChannelUpdateProposal {
	return &NullableChannelUpdateProposal{value: val, isSet: true}
}

func (v NullableChannelUpdateProposal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelUpdateProposal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


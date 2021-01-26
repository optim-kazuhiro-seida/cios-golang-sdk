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

// MultipleVideo struct for MultipleVideo
type MultipleVideo struct {
	Total int64 `json:"total"`
	Videos []Video `json:"videos"`
}

// NewMultipleVideo instantiates a new MultipleVideo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultipleVideo(total int64, videos []Video, ) *MultipleVideo {
	this := MultipleVideo{}
	this.Total = total
	this.Videos = videos
	return &this
}

// NewMultipleVideoWithDefaults instantiates a new MultipleVideo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultipleVideoWithDefaults() *MultipleVideo {
	this := MultipleVideo{}
	return &this
}

// GetTotal returns the Total field value
func (o *MultipleVideo) GetTotal() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *MultipleVideo) GetTotalOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *MultipleVideo) SetTotal(v int64) {
	o.Total = v
}

// GetVideos returns the Videos field value
func (o *MultipleVideo) GetVideos() []Video {
	if o == nil  {
		var ret []Video
		return ret
	}

	return o.Videos
}

// GetVideosOk returns a tuple with the Videos field value
// and a boolean to check if the value has been set.
func (o *MultipleVideo) GetVideosOk() (*[]Video, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Videos, true
}

// SetVideos sets field value
func (o *MultipleVideo) SetVideos(v []Video) {
	o.Videos = v
}

func (o MultipleVideo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["videos"] = o.Videos
	}
	return json.Marshal(toSerialize)
}

type NullableMultipleVideo struct {
	value *MultipleVideo
	isSet bool
}

func (v NullableMultipleVideo) Get() *MultipleVideo {
	return v.value
}

func (v *NullableMultipleVideo) Set(val *MultipleVideo) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipleVideo) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipleVideo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipleVideo(val *MultipleVideo) *NullableMultipleVideo {
	return &NullableMultipleVideo{value: val, isSet: true}
}

func (v NullableMultipleVideo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipleVideo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


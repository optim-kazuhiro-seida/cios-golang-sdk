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

// SingleBucket struct for SingleBucket
type SingleBucket struct {
	Bucket Bucket `json:"bucket"`
}

// NewSingleBucket instantiates a new SingleBucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleBucket(bucket Bucket, ) *SingleBucket {
	this := SingleBucket{}
	this.Bucket = bucket
	return &this
}

// NewSingleBucketWithDefaults instantiates a new SingleBucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleBucketWithDefaults() *SingleBucket {
	this := SingleBucket{}
	return &this
}

// GetBucket returns the Bucket field value
func (o *SingleBucket) GetBucket() Bucket {
	if o == nil  {
		var ret Bucket
		return ret
	}

	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *SingleBucket) GetBucketOk() (*Bucket, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Bucket, true
}

// SetBucket sets field value
func (o *SingleBucket) SetBucket(v Bucket) {
	o.Bucket = v
}

func (o SingleBucket) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bucket"] = o.Bucket
	}
	return json.Marshal(toSerialize)
}

type NullableSingleBucket struct {
	value *SingleBucket
	isSet bool
}

func (v NullableSingleBucket) Get() *SingleBucket {
	return v.value
}

func (v *NullableSingleBucket) Set(val *SingleBucket) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleBucket) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleBucket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleBucket(val *SingleBucket) *NullableSingleBucket {
	return &NullableSingleBucket{value: val, isSet: true}
}

func (v NullableSingleBucket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleBucket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


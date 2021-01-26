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

// ContractUser struct for ContractUser
type ContractUser struct {
	Id *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
	Name *string `json:"name,omitempty"`
	FamilyName *string `json:"family_name,omitempty"`
	GivenName *string `json:"given_name,omitempty"`
	PhoneticFamilyName *string `json:"phonetic_family_name,omitempty"`
	PhoneticGivenName *string `json:"phonetic_given_name,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	PhoneNumber2 *string `json:"phone_number_2,omitempty"`
	// ISO639
	Language *string `json:"language,omitempty"`
	// uri
	Picture *string `json:"picture,omitempty"`
	Email *string `json:"email,omitempty"`
	Emails *[]string `json:"emails,omitempty"`
	Role *string `json:"role,omitempty"`
	Category *string `json:"category,omitempty"`
	Address *ContractOwnerAddress `json:"address,omitempty"`
	License *ContractUserLicense `json:"license,omitempty"`
}

// NewContractUser instantiates a new ContractUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractUser() *ContractUser {
	this := ContractUser{}
	return &this
}

// NewContractUserWithDefaults instantiates a new ContractUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractUserWithDefaults() *ContractUser {
	this := ContractUser{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ContractUser) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractUser) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ContractUser) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ContractUser) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ContractUser) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractUser) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ContractUser) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ContractUser) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContractUser) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractUser) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContractUser) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContractUser) SetName(v string) {
	o.Name = &v
}

// GetFamilyName returns the FamilyName field value if set, zero value otherwise.
func (o *ContractUser) GetFamilyName() string {
	if o == nil || o.FamilyName == nil {
		var ret string
		return ret
	}
	return *o.FamilyName
}

// GetFamilyNameOk returns a tuple with the FamilyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractUser) GetFamilyNameOk() (*string, bool) {
	if o == nil || o.FamilyName == nil {
		return nil, false
	}
	return o.FamilyName, true
}

// HasFamilyName returns a boolean if a field has been set.
func (o *ContractUser) HasFamilyName() bool {
	if o != nil && o.FamilyName != nil {
		return true
	}

	return false
}

// SetFamilyName gets a reference to the given string and assigns it to the FamilyName field.
func (o *ContractUser) SetFamilyName(v string) {
	o.FamilyName = &v
}

// GetGivenName returns the GivenName field value if set, zero value otherwise.
func (o *ContractUser) GetGivenName() string {
	if o == nil || o.GivenName == nil {
		var ret string
		return ret
	}
	return *o.GivenName
}

// GetGivenNameOk returns a tuple with the GivenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractUser) GetGivenNameOk() (*string, bool) {
	if o == nil || o.GivenName == nil {
		return nil, false
	}
	return o.GivenName, true
}

// HasGivenName returns a boolean if a field has been set.
func (o *ContractUser) HasGivenName() bool {
	if o != nil && o.GivenName != nil {
		return true
	}

	return false
}

// SetGivenName gets a reference to the given string and assigns it to the GivenName field.
func (o *ContractUser) SetGivenName(v string) {
	o.GivenName = &v
}

// GetPhoneticFamilyName returns the PhoneticFamilyName field value if set, zero value otherwise.
func (o *ContractUser) GetPhoneticFamilyName() string {
	if o == nil || o.PhoneticFamilyName == nil {
		var ret string
		return ret
	}
	return *o.PhoneticFamilyName
}

// GetPhoneticFamilyNameOk returns a tuple with the PhoneticFamilyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractUser) GetPhoneticFamilyNameOk() (*string, bool) {
	if o == nil || o.PhoneticFamilyName == nil {
		return nil, false
	}
	return o.PhoneticFamilyName, true
}

// HasPhoneticFamilyName returns a boolean if a field has been set.
func (o *ContractUser) HasPhoneticFamilyName() bool {
	if o != nil && o.PhoneticFamilyName != nil {
		return true
	}

	return false
}

// SetPhoneticFamilyName gets a reference to the given string and assigns it to the PhoneticFamilyName field.
func (o *ContractUser) SetPhoneticFamilyName(v string) {
	o.PhoneticFamilyName = &v
}

// GetPhoneticGivenName returns the PhoneticGivenName field value if set, zero value otherwise.
func (o *ContractUser) GetPhoneticGivenName() string {
	if o == nil || o.PhoneticGivenName == nil {
		var ret string
		return ret
	}
	return *o.PhoneticGivenName
}

// GetPhoneticGivenNameOk returns a tuple with the PhoneticGivenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractUser) GetPhoneticGivenNameOk() (*string, bool) {
	if o == nil || o.PhoneticGivenName == nil {
		return nil, false
	}
	return o.PhoneticGivenName, true
}

// HasPhoneticGivenName returns a boolean if a field has been set.
func (o *ContractUser) HasPhoneticGivenName() bool {
	if o != nil && o.PhoneticGivenName != nil {
		return true
	}

	return false
}

// SetPhoneticGivenName gets a reference to the given string and assigns it to the PhoneticGivenName field.
func (o *ContractUser) SetPhoneticGivenName(v string) {
	o.PhoneticGivenName = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *ContractUser) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractUser) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *ContractUser) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *ContractUser) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetPhoneNumber2 returns the PhoneNumber2 field value if set, zero value otherwise.
func (o *ContractUser) GetPhoneNumber2() string {
	if o == nil || o.PhoneNumber2 == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber2
}

// GetPhoneNumber2Ok returns a tuple with the PhoneNumber2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractUser) GetPhoneNumber2Ok() (*string, bool) {
	if o == nil || o.PhoneNumber2 == nil {
		return nil, false
	}
	return o.PhoneNumber2, true
}

// HasPhoneNumber2 returns a boolean if a field has been set.
func (o *ContractUser) HasPhoneNumber2() bool {
	if o != nil && o.PhoneNumber2 != nil {
		return true
	}

	return false
}

// SetPhoneNumber2 gets a reference to the given string and assigns it to the PhoneNumber2 field.
func (o *ContractUser) SetPhoneNumber2(v string) {
	o.PhoneNumber2 = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *ContractUser) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractUser) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *ContractUser) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *ContractUser) SetLanguage(v string) {
	o.Language = &v
}

// GetPicture returns the Picture field value if set, zero value otherwise.
func (o *ContractUser) GetPicture() string {
	if o == nil || o.Picture == nil {
		var ret string
		return ret
	}
	return *o.Picture
}

// GetPictureOk returns a tuple with the Picture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractUser) GetPictureOk() (*string, bool) {
	if o == nil || o.Picture == nil {
		return nil, false
	}
	return o.Picture, true
}

// HasPicture returns a boolean if a field has been set.
func (o *ContractUser) HasPicture() bool {
	if o != nil && o.Picture != nil {
		return true
	}

	return false
}

// SetPicture gets a reference to the given string and assigns it to the Picture field.
func (o *ContractUser) SetPicture(v string) {
	o.Picture = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ContractUser) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractUser) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ContractUser) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ContractUser) SetEmail(v string) {
	o.Email = &v
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *ContractUser) GetEmails() []string {
	if o == nil || o.Emails == nil {
		var ret []string
		return ret
	}
	return *o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractUser) GetEmailsOk() (*[]string, bool) {
	if o == nil || o.Emails == nil {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *ContractUser) HasEmails() bool {
	if o != nil && o.Emails != nil {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []string and assigns it to the Emails field.
func (o *ContractUser) SetEmails(v []string) {
	o.Emails = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *ContractUser) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractUser) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ContractUser) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *ContractUser) SetRole(v string) {
	o.Role = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *ContractUser) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractUser) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *ContractUser) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *ContractUser) SetCategory(v string) {
	o.Category = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ContractUser) GetAddress() ContractOwnerAddress {
	if o == nil || o.Address == nil {
		var ret ContractOwnerAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractUser) GetAddressOk() (*ContractOwnerAddress, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ContractUser) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given ContractOwnerAddress and assigns it to the Address field.
func (o *ContractUser) SetAddress(v ContractOwnerAddress) {
	o.Address = &v
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *ContractUser) GetLicense() ContractUserLicense {
	if o == nil || o.License == nil {
		var ret ContractUserLicense
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractUser) GetLicenseOk() (*ContractUserLicense, bool) {
	if o == nil || o.License == nil {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *ContractUser) HasLicense() bool {
	if o != nil && o.License != nil {
		return true
	}

	return false
}

// SetLicense gets a reference to the given ContractUserLicense and assigns it to the License field.
func (o *ContractUser) SetLicense(v ContractUserLicense) {
	o.License = &v
}

func (o ContractUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.FamilyName != nil {
		toSerialize["family_name"] = o.FamilyName
	}
	if o.GivenName != nil {
		toSerialize["given_name"] = o.GivenName
	}
	if o.PhoneticFamilyName != nil {
		toSerialize["phonetic_family_name"] = o.PhoneticFamilyName
	}
	if o.PhoneticGivenName != nil {
		toSerialize["phonetic_given_name"] = o.PhoneticGivenName
	}
	if o.PhoneNumber != nil {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if o.PhoneNumber2 != nil {
		toSerialize["phone_number_2"] = o.PhoneNumber2
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.Picture != nil {
		toSerialize["picture"] = o.Picture
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Emails != nil {
		toSerialize["emails"] = o.Emails
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.License != nil {
		toSerialize["license"] = o.License
	}
	return json.Marshal(toSerialize)
}

type NullableContractUser struct {
	value *ContractUser
	isSet bool
}

func (v NullableContractUser) Get() *ContractUser {
	return v.value
}

func (v *NullableContractUser) Set(val *ContractUser) {
	v.value = val
	v.isSet = true
}

func (v NullableContractUser) IsSet() bool {
	return v.isSet
}

func (v *NullableContractUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractUser(val *ContractUser) *NullableContractUser {
	return &NullableContractUser{value: val, isSet: true}
}

func (v NullableContractUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * API version: v0.7.3-alpha.5
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SelfServiceLogoutUrl struct for SelfServiceLogoutUrl
type SelfServiceLogoutUrl struct {
	// LogoutURL can be opened in a browser to  format: uri
	LogoutUrl *string `json:"logout_url,omitempty"`
}

// NewSelfServiceLogoutUrl instantiates a new SelfServiceLogoutUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfServiceLogoutUrl() *SelfServiceLogoutUrl {
	this := SelfServiceLogoutUrl{}
	return &this
}

// NewSelfServiceLogoutUrlWithDefaults instantiates a new SelfServiceLogoutUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfServiceLogoutUrlWithDefaults() *SelfServiceLogoutUrl {
	this := SelfServiceLogoutUrl{}
	return &this
}

// GetLogoutUrl returns the LogoutUrl field value if set, zero value otherwise.
func (o *SelfServiceLogoutUrl) GetLogoutUrl() string {
	if o == nil || o.LogoutUrl == nil {
		var ret string
		return ret
	}
	return *o.LogoutUrl
}

// GetLogoutUrlOk returns a tuple with the LogoutUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceLogoutUrl) GetLogoutUrlOk() (*string, bool) {
	if o == nil || o.LogoutUrl == nil {
		return nil, false
	}
	return o.LogoutUrl, true
}

// HasLogoutUrl returns a boolean if a field has been set.
func (o *SelfServiceLogoutUrl) HasLogoutUrl() bool {
	if o != nil && o.LogoutUrl != nil {
		return true
	}

	return false
}

// SetLogoutUrl gets a reference to the given string and assigns it to the LogoutUrl field.
func (o *SelfServiceLogoutUrl) SetLogoutUrl(v string) {
	o.LogoutUrl = &v
}

func (o SelfServiceLogoutUrl) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LogoutUrl != nil {
		toSerialize["logout_url"] = o.LogoutUrl
	}
	return json.Marshal(toSerialize)
}

type NullableSelfServiceLogoutUrl struct {
	value *SelfServiceLogoutUrl
	isSet bool
}

func (v NullableSelfServiceLogoutUrl) Get() *SelfServiceLogoutUrl {
	return v.value
}

func (v *NullableSelfServiceLogoutUrl) Set(val *SelfServiceLogoutUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfServiceLogoutUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfServiceLogoutUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfServiceLogoutUrl(val *SelfServiceLogoutUrl) *NullableSelfServiceLogoutUrl {
	return &NullableSelfServiceLogoutUrl{value: val, isSet: true}
}

func (v NullableSelfServiceLogoutUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfServiceLogoutUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



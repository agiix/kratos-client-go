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

// AdminCreateSelfServiceRecoveryLinkBody struct for AdminCreateSelfServiceRecoveryLinkBody
type AdminCreateSelfServiceRecoveryLinkBody struct {
	// Link Expires In  The recovery link will expire at that point in time. Defaults to the configuration value of `selfservice.flows.recovery.request_lifespan`.
	ExpiresIn *string `json:"expires_in,omitempty"`
	IdentityId string `json:"identity_id"`
}

// NewAdminCreateSelfServiceRecoveryLinkBody instantiates a new AdminCreateSelfServiceRecoveryLinkBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminCreateSelfServiceRecoveryLinkBody(identityId string) *AdminCreateSelfServiceRecoveryLinkBody {
	this := AdminCreateSelfServiceRecoveryLinkBody{}
	this.IdentityId = identityId
	return &this
}

// NewAdminCreateSelfServiceRecoveryLinkBodyWithDefaults instantiates a new AdminCreateSelfServiceRecoveryLinkBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminCreateSelfServiceRecoveryLinkBodyWithDefaults() *AdminCreateSelfServiceRecoveryLinkBody {
	this := AdminCreateSelfServiceRecoveryLinkBody{}
	return &this
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *AdminCreateSelfServiceRecoveryLinkBody) GetExpiresIn() string {
	if o == nil || o.ExpiresIn == nil {
		var ret string
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminCreateSelfServiceRecoveryLinkBody) GetExpiresInOk() (*string, bool) {
	if o == nil || o.ExpiresIn == nil {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *AdminCreateSelfServiceRecoveryLinkBody) HasExpiresIn() bool {
	if o != nil && o.ExpiresIn != nil {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given string and assigns it to the ExpiresIn field.
func (o *AdminCreateSelfServiceRecoveryLinkBody) SetExpiresIn(v string) {
	o.ExpiresIn = &v
}

// GetIdentityId returns the IdentityId field value
func (o *AdminCreateSelfServiceRecoveryLinkBody) GetIdentityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value
// and a boolean to check if the value has been set.
func (o *AdminCreateSelfServiceRecoveryLinkBody) GetIdentityIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IdentityId, true
}

// SetIdentityId sets field value
func (o *AdminCreateSelfServiceRecoveryLinkBody) SetIdentityId(v string) {
	o.IdentityId = v
}

func (o AdminCreateSelfServiceRecoveryLinkBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpiresIn != nil {
		toSerialize["expires_in"] = o.ExpiresIn
	}
	if true {
		toSerialize["identity_id"] = o.IdentityId
	}
	return json.Marshal(toSerialize)
}

type NullableAdminCreateSelfServiceRecoveryLinkBody struct {
	value *AdminCreateSelfServiceRecoveryLinkBody
	isSet bool
}

func (v NullableAdminCreateSelfServiceRecoveryLinkBody) Get() *AdminCreateSelfServiceRecoveryLinkBody {
	return v.value
}

func (v *NullableAdminCreateSelfServiceRecoveryLinkBody) Set(val *AdminCreateSelfServiceRecoveryLinkBody) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminCreateSelfServiceRecoveryLinkBody) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminCreateSelfServiceRecoveryLinkBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminCreateSelfServiceRecoveryLinkBody(val *AdminCreateSelfServiceRecoveryLinkBody) *NullableAdminCreateSelfServiceRecoveryLinkBody {
	return &NullableAdminCreateSelfServiceRecoveryLinkBody{value: val, isSet: true}
}

func (v NullableAdminCreateSelfServiceRecoveryLinkBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminCreateSelfServiceRecoveryLinkBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



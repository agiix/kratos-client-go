/*
Ory Identities API

This is the API specification for Ory Identities with features such as registration, login, recovery, account verification, profile settings, password reset, identity management, session management, email and sms delivery, and more. 

API version: v1.1.0
Contact: office@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the Session type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Session{}

// Session A Session
type Session struct {
	// Active state. If false the session is no longer active.
	Active *bool `json:"active,omitempty"`
	// The Session Authentication Timestamp  When this session was authenticated at. If multi-factor authentication was used this is the time when the last factor was authenticated (e.g. the TOTP code challenge was completed).
	AuthenticatedAt *time.Time `json:"authenticated_at,omitempty"`
	// A list of authenticators which were used to authenticate the session.
	AuthenticationMethods []SessionAuthenticationMethod `json:"authentication_methods,omitempty"`
	AuthenticatorAssuranceLevel *AuthenticatorAssuranceLevel `json:"authenticator_assurance_level,omitempty"`
	// Devices has history of all endpoints where the session was used
	Devices []SessionDevice `json:"devices,omitempty"`
	// The Session Expiry  When this session expires at.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// Session ID
	Id string `json:"id"`
	Identity *Identity `json:"identity,omitempty"`
	// The Session Issuance Timestamp  When this session was issued at. Usually equal or close to `authenticated_at`.
	IssuedAt *time.Time `json:"issued_at,omitempty"`
	// Tokenized is the tokenized (e.g. JWT) version of the session.  It is only set when the `tokenize` query parameter was set to a valid tokenize template during calls to `/session/whoami`.
	Tokenized *string `json:"tokenized,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Session Session

// NewSession instantiates a new Session object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSession(id string) *Session {
	this := Session{}
	this.Id = id
	return &this
}

// NewSessionWithDefaults instantiates a new Session object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionWithDefaults() *Session {
	this := Session{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *Session) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *Session) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *Session) SetActive(v bool) {
	o.Active = &v
}

// GetAuthenticatedAt returns the AuthenticatedAt field value if set, zero value otherwise.
func (o *Session) GetAuthenticatedAt() time.Time {
	if o == nil || IsNil(o.AuthenticatedAt) {
		var ret time.Time
		return ret
	}
	return *o.AuthenticatedAt
}

// GetAuthenticatedAtOk returns a tuple with the AuthenticatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetAuthenticatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AuthenticatedAt) {
		return nil, false
	}
	return o.AuthenticatedAt, true
}

// HasAuthenticatedAt returns a boolean if a field has been set.
func (o *Session) HasAuthenticatedAt() bool {
	if o != nil && !IsNil(o.AuthenticatedAt) {
		return true
	}

	return false
}

// SetAuthenticatedAt gets a reference to the given time.Time and assigns it to the AuthenticatedAt field.
func (o *Session) SetAuthenticatedAt(v time.Time) {
	o.AuthenticatedAt = &v
}

// GetAuthenticationMethods returns the AuthenticationMethods field value if set, zero value otherwise.
func (o *Session) GetAuthenticationMethods() []SessionAuthenticationMethod {
	if o == nil || IsNil(o.AuthenticationMethods) {
		var ret []SessionAuthenticationMethod
		return ret
	}
	return o.AuthenticationMethods
}

// GetAuthenticationMethodsOk returns a tuple with the AuthenticationMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetAuthenticationMethodsOk() ([]SessionAuthenticationMethod, bool) {
	if o == nil || IsNil(o.AuthenticationMethods) {
		return nil, false
	}
	return o.AuthenticationMethods, true
}

// HasAuthenticationMethods returns a boolean if a field has been set.
func (o *Session) HasAuthenticationMethods() bool {
	if o != nil && !IsNil(o.AuthenticationMethods) {
		return true
	}

	return false
}

// SetAuthenticationMethods gets a reference to the given []SessionAuthenticationMethod and assigns it to the AuthenticationMethods field.
func (o *Session) SetAuthenticationMethods(v []SessionAuthenticationMethod) {
	o.AuthenticationMethods = v
}

// GetAuthenticatorAssuranceLevel returns the AuthenticatorAssuranceLevel field value if set, zero value otherwise.
func (o *Session) GetAuthenticatorAssuranceLevel() AuthenticatorAssuranceLevel {
	if o == nil || IsNil(o.AuthenticatorAssuranceLevel) {
		var ret AuthenticatorAssuranceLevel
		return ret
	}
	return *o.AuthenticatorAssuranceLevel
}

// GetAuthenticatorAssuranceLevelOk returns a tuple with the AuthenticatorAssuranceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetAuthenticatorAssuranceLevelOk() (*AuthenticatorAssuranceLevel, bool) {
	if o == nil || IsNil(o.AuthenticatorAssuranceLevel) {
		return nil, false
	}
	return o.AuthenticatorAssuranceLevel, true
}

// HasAuthenticatorAssuranceLevel returns a boolean if a field has been set.
func (o *Session) HasAuthenticatorAssuranceLevel() bool {
	if o != nil && !IsNil(o.AuthenticatorAssuranceLevel) {
		return true
	}

	return false
}

// SetAuthenticatorAssuranceLevel gets a reference to the given AuthenticatorAssuranceLevel and assigns it to the AuthenticatorAssuranceLevel field.
func (o *Session) SetAuthenticatorAssuranceLevel(v AuthenticatorAssuranceLevel) {
	o.AuthenticatorAssuranceLevel = &v
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *Session) GetDevices() []SessionDevice {
	if o == nil || IsNil(o.Devices) {
		var ret []SessionDevice
		return ret
	}
	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetDevicesOk() ([]SessionDevice, bool) {
	if o == nil || IsNil(o.Devices) {
		return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *Session) HasDevices() bool {
	if o != nil && !IsNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []SessionDevice and assigns it to the Devices field.
func (o *Session) SetDevices(v []SessionDevice) {
	o.Devices = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *Session) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *Session) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *Session) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetId returns the Id field value
func (o *Session) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Session) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Session) SetId(v string) {
	o.Id = v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *Session) GetIdentity() Identity {
	if o == nil || IsNil(o.Identity) {
		var ret Identity
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetIdentityOk() (*Identity, bool) {
	if o == nil || IsNil(o.Identity) {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *Session) HasIdentity() bool {
	if o != nil && !IsNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given Identity and assigns it to the Identity field.
func (o *Session) SetIdentity(v Identity) {
	o.Identity = &v
}

// GetIssuedAt returns the IssuedAt field value if set, zero value otherwise.
func (o *Session) GetIssuedAt() time.Time {
	if o == nil || IsNil(o.IssuedAt) {
		var ret time.Time
		return ret
	}
	return *o.IssuedAt
}

// GetIssuedAtOk returns a tuple with the IssuedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetIssuedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.IssuedAt) {
		return nil, false
	}
	return o.IssuedAt, true
}

// HasIssuedAt returns a boolean if a field has been set.
func (o *Session) HasIssuedAt() bool {
	if o != nil && !IsNil(o.IssuedAt) {
		return true
	}

	return false
}

// SetIssuedAt gets a reference to the given time.Time and assigns it to the IssuedAt field.
func (o *Session) SetIssuedAt(v time.Time) {
	o.IssuedAt = &v
}

// GetTokenized returns the Tokenized field value if set, zero value otherwise.
func (o *Session) GetTokenized() string {
	if o == nil || IsNil(o.Tokenized) {
		var ret string
		return ret
	}
	return *o.Tokenized
}

// GetTokenizedOk returns a tuple with the Tokenized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetTokenizedOk() (*string, bool) {
	if o == nil || IsNil(o.Tokenized) {
		return nil, false
	}
	return o.Tokenized, true
}

// HasTokenized returns a boolean if a field has been set.
func (o *Session) HasTokenized() bool {
	if o != nil && !IsNil(o.Tokenized) {
		return true
	}

	return false
}

// SetTokenized gets a reference to the given string and assigns it to the Tokenized field.
func (o *Session) SetTokenized(v string) {
	o.Tokenized = &v
}

func (o Session) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Session) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.AuthenticatedAt) {
		toSerialize["authenticated_at"] = o.AuthenticatedAt
	}
	if !IsNil(o.AuthenticationMethods) {
		toSerialize["authentication_methods"] = o.AuthenticationMethods
	}
	if !IsNil(o.AuthenticatorAssuranceLevel) {
		toSerialize["authenticator_assurance_level"] = o.AuthenticatorAssuranceLevel
	}
	if !IsNil(o.Devices) {
		toSerialize["devices"] = o.Devices
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Identity) {
		toSerialize["identity"] = o.Identity
	}
	if !IsNil(o.IssuedAt) {
		toSerialize["issued_at"] = o.IssuedAt
	}
	if !IsNil(o.Tokenized) {
		toSerialize["tokenized"] = o.Tokenized
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Session) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSession := _Session{}

	err = json.Unmarshal(bytes, &varSession)

	if err != nil {
		return err
	}

	*o = Session(varSession)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "active")
		delete(additionalProperties, "authenticated_at")
		delete(additionalProperties, "authentication_methods")
		delete(additionalProperties, "authenticator_assurance_level")
		delete(additionalProperties, "devices")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "identity")
		delete(additionalProperties, "issued_at")
		delete(additionalProperties, "tokenized")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSession struct {
	value *Session
	isSet bool
}

func (v NullableSession) Get() *Session {
	return v.value
}

func (v *NullableSession) Set(val *Session) {
	v.value = val
	v.isSet = true
}

func (v NullableSession) IsSet() bool {
	return v.isSet
}

func (v *NullableSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSession(val *Session) *NullableSession {
	return &NullableSession{value: val, isSet: true}
}

func (v NullableSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



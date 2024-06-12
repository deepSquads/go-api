/*
@open-sauced/api.opensauced.pizza

 ## Swagger-UI API Documentation  This REST API can be used to create, read, update or delete data from the Open Sauced community platform. The Swagger-UI provides useful information to get started and an overview of all available resources. Each API route is clickable and has their own detailed description on how to use it. The base URL for the API is [api.opensauced.pizza](https://api.opensauced.pizza).  [comment]: # (TODO: add bearer auth information)  ## Rate limiting  Every IP address is allowed to perform 5000 requests per hour. This is measured by saving the date of the initial request and counting all requests in the next hour. When an IP address goes over the limit, HTTP status code 429 is returned. The returned HTTP headers of any API request show the current rate limit status:  header | description --- | --- `X-RateLimit-Limit` | The maximum number of requests allowed per hour `X-RateLimit-Remaining` | The number of requests remaining in the current rate limit window `X-RateLimit-Reset` | The date and time at which the current rate limit window resets in [UTC epoch seconds](https://en.wikipedia.org/wiki/Unix_time)  [comment]: # (TODO: add pagination information)  ## Common response codes  Each route shows for each method which data they expect and which they will respond when the call succeeds. The table below shows most common response codes you can receive from our endpoints.  code | condition --- | --- [`200`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/200) | The [`GET`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/GET) request was handled successfully. The response provides the requested data. [`201`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/201) | The [`POST`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/POST) request was handled successfully. The response provides the created data. [`204`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/204) | The [`PATCH`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/PATCH) or [`DELETE`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/DELETE) request was handled successfully. The response provides no data, generally. [`400`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400) | The server will not process the request due to something that is perceived to be a client error. Check the provided error for mote information. [`401`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401) | The request requires user authentication. Check the provided error for more information. [`403`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403) | The request was valid, but the server is refusing user access. Check the provided error for more information. [`404`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404) | The requested resource could not be found. Check the provided error for more information. [`429`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429) | The current API Key made too many requests in the last hour. Check [Rate limiting](#ratelimiting) for more information.  ## Additional links

API version: 2
Contact: hello@opensauced.pizza
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the DbInsightMember type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DbInsightMember{}

// DbInsightMember struct for DbInsightMember
type DbInsightMember struct {
	// Insight Member identifier
	Id string `json:"id"`
	// Insight ID
	InsightId int32 `json:"insight_id"`
	// User ID
	UserId *int32 `json:"user_id,omitempty"`
	// User's Name
	Name *string `json:"name,omitempty"`
	// Insight Member Access
	Access string `json:"access"`
	// Timestamp representing team member creation
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Timestamp representing team member last updated
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Timestamp representing team member deletion
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Timestamp representing team member invitation email sent date
	InvitationEmailedAt *time.Time `json:"invitation_emailed_at,omitempty"`
	// Team Member Invitation Email
	InvitationEmail *string `json:"invitation_email,omitempty"`
}

// NewDbInsightMember instantiates a new DbInsightMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbInsightMember(id string, insightId int32, access string) *DbInsightMember {
	this := DbInsightMember{}
	this.Id = id
	this.InsightId = insightId
	this.Access = access
	return &this
}

// NewDbInsightMemberWithDefaults instantiates a new DbInsightMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbInsightMemberWithDefaults() *DbInsightMember {
	this := DbInsightMember{}
	return &this
}

// GetId returns the Id field value
func (o *DbInsightMember) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DbInsightMember) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DbInsightMember) SetId(v string) {
	o.Id = v
}

// GetInsightId returns the InsightId field value
func (o *DbInsightMember) GetInsightId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InsightId
}

// GetInsightIdOk returns a tuple with the InsightId field value
// and a boolean to check if the value has been set.
func (o *DbInsightMember) GetInsightIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InsightId, true
}

// SetInsightId sets field value
func (o *DbInsightMember) SetInsightId(v int32) {
	o.InsightId = v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *DbInsightMember) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbInsightMember) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *DbInsightMember) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *DbInsightMember) SetUserId(v int32) {
	o.UserId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DbInsightMember) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbInsightMember) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DbInsightMember) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DbInsightMember) SetName(v string) {
	o.Name = &v
}

// GetAccess returns the Access field value
func (o *DbInsightMember) GetAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *DbInsightMember) GetAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *DbInsightMember) SetAccess(v string) {
	o.Access = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DbInsightMember) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbInsightMember) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DbInsightMember) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DbInsightMember) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DbInsightMember) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbInsightMember) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DbInsightMember) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DbInsightMember) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *DbInsightMember) GetDeletedAt() time.Time {
	if o == nil || IsNil(o.DeletedAt) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbInsightMember) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *DbInsightMember) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *DbInsightMember) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetInvitationEmailedAt returns the InvitationEmailedAt field value if set, zero value otherwise.
func (o *DbInsightMember) GetInvitationEmailedAt() time.Time {
	if o == nil || IsNil(o.InvitationEmailedAt) {
		var ret time.Time
		return ret
	}
	return *o.InvitationEmailedAt
}

// GetInvitationEmailedAtOk returns a tuple with the InvitationEmailedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbInsightMember) GetInvitationEmailedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.InvitationEmailedAt) {
		return nil, false
	}
	return o.InvitationEmailedAt, true
}

// HasInvitationEmailedAt returns a boolean if a field has been set.
func (o *DbInsightMember) HasInvitationEmailedAt() bool {
	if o != nil && !IsNil(o.InvitationEmailedAt) {
		return true
	}

	return false
}

// SetInvitationEmailedAt gets a reference to the given time.Time and assigns it to the InvitationEmailedAt field.
func (o *DbInsightMember) SetInvitationEmailedAt(v time.Time) {
	o.InvitationEmailedAt = &v
}

// GetInvitationEmail returns the InvitationEmail field value if set, zero value otherwise.
func (o *DbInsightMember) GetInvitationEmail() string {
	if o == nil || IsNil(o.InvitationEmail) {
		var ret string
		return ret
	}
	return *o.InvitationEmail
}

// GetInvitationEmailOk returns a tuple with the InvitationEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbInsightMember) GetInvitationEmailOk() (*string, bool) {
	if o == nil || IsNil(o.InvitationEmail) {
		return nil, false
	}
	return o.InvitationEmail, true
}

// HasInvitationEmail returns a boolean if a field has been set.
func (o *DbInsightMember) HasInvitationEmail() bool {
	if o != nil && !IsNil(o.InvitationEmail) {
		return true
	}

	return false
}

// SetInvitationEmail gets a reference to the given string and assigns it to the InvitationEmail field.
func (o *DbInsightMember) SetInvitationEmail(v string) {
	o.InvitationEmail = &v
}

func (o DbInsightMember) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DbInsightMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["insight_id"] = o.InsightId
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["access"] = o.Access
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if !IsNil(o.InvitationEmailedAt) {
		toSerialize["invitation_emailed_at"] = o.InvitationEmailedAt
	}
	if !IsNil(o.InvitationEmail) {
		toSerialize["invitation_email"] = o.InvitationEmail
	}
	return toSerialize, nil
}

type NullableDbInsightMember struct {
	value *DbInsightMember
	isSet bool
}

func (v NullableDbInsightMember) Get() *DbInsightMember {
	return v.value
}

func (v *NullableDbInsightMember) Set(val *DbInsightMember) {
	v.value = val
	v.isSet = true
}

func (v NullableDbInsightMember) IsSet() bool {
	return v.isSet
}

func (v *NullableDbInsightMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbInsightMember(val *DbInsightMember) *NullableDbInsightMember {
	return &NullableDbInsightMember{value: val, isSet: true}
}

func (v NullableDbInsightMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbInsightMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
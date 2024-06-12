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
)

// checks if the CreateWorkspaceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWorkspaceDto{}

// CreateWorkspaceDto struct for CreateWorkspaceDto
type CreateWorkspaceDto struct {
	// Workspace name
	Name string `json:"name"`
	// Workspace description
	Description string `json:"description"`
	// An array of new member objects
	Members []interface{} `json:"members"`
	// An array of repo objects to be added to the workspace
	Repos []interface{} `json:"repos"`
}

// NewCreateWorkspaceDto instantiates a new CreateWorkspaceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWorkspaceDto(name string, description string, members []interface{}, repos []interface{}) *CreateWorkspaceDto {
	this := CreateWorkspaceDto{}
	this.Name = name
	this.Description = description
	this.Members = members
	this.Repos = repos
	return &this
}

// NewCreateWorkspaceDtoWithDefaults instantiates a new CreateWorkspaceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWorkspaceDtoWithDefaults() *CreateWorkspaceDto {
	this := CreateWorkspaceDto{}
	return &this
}

// GetName returns the Name field value
func (o *CreateWorkspaceDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateWorkspaceDto) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *CreateWorkspaceDto) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateWorkspaceDto) SetDescription(v string) {
	o.Description = v
}

// GetMembers returns the Members field value
func (o *CreateWorkspaceDto) GetMembers() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceDto) GetMembersOk() ([]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Members, true
}

// SetMembers sets field value
func (o *CreateWorkspaceDto) SetMembers(v []interface{}) {
	o.Members = v
}

// GetRepos returns the Repos field value
func (o *CreateWorkspaceDto) GetRepos() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.Repos
}

// GetReposOk returns a tuple with the Repos field value
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceDto) GetReposOk() ([]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Repos, true
}

// SetRepos sets field value
func (o *CreateWorkspaceDto) SetRepos(v []interface{}) {
	o.Repos = v
}

func (o CreateWorkspaceDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWorkspaceDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["members"] = o.Members
	toSerialize["repos"] = o.Repos
	return toSerialize, nil
}

type NullableCreateWorkspaceDto struct {
	value *CreateWorkspaceDto
	isSet bool
}

func (v NullableCreateWorkspaceDto) Get() *CreateWorkspaceDto {
	return v.value
}

func (v *NullableCreateWorkspaceDto) Set(val *CreateWorkspaceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWorkspaceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWorkspaceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWorkspaceDto(val *CreateWorkspaceDto) *NullableCreateWorkspaceDto {
	return &NullableCreateWorkspaceDto{value: val, isSet: true}
}

func (v NullableCreateWorkspaceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWorkspaceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
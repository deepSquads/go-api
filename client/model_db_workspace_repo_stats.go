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

// checks if the DbWorkspaceRepoStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DbWorkspaceRepoStats{}

// DbWorkspaceRepoStats struct for DbWorkspaceRepoStats
type DbWorkspaceRepoStats struct {
	// Number of stars for workspace repos
	Stars int32 `json:"stars"`
	// Number of forks for repos
	Forks int32 `json:"forks"`
	// Repository average health
	Health int32 `json:"health"`
}

// NewDbWorkspaceRepoStats instantiates a new DbWorkspaceRepoStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbWorkspaceRepoStats(stars int32, forks int32, health int32) *DbWorkspaceRepoStats {
	this := DbWorkspaceRepoStats{}
	this.Stars = stars
	this.Forks = forks
	this.Health = health
	return &this
}

// NewDbWorkspaceRepoStatsWithDefaults instantiates a new DbWorkspaceRepoStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbWorkspaceRepoStatsWithDefaults() *DbWorkspaceRepoStats {
	this := DbWorkspaceRepoStats{}
	var stars int32 = 0
	this.Stars = stars
	var forks int32 = 0
	this.Forks = forks
	var health int32 = 0
	this.Health = health
	return &this
}

// GetStars returns the Stars field value
func (o *DbWorkspaceRepoStats) GetStars() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Stars
}

// GetStarsOk returns a tuple with the Stars field value
// and a boolean to check if the value has been set.
func (o *DbWorkspaceRepoStats) GetStarsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stars, true
}

// SetStars sets field value
func (o *DbWorkspaceRepoStats) SetStars(v int32) {
	o.Stars = v
}

// GetForks returns the Forks field value
func (o *DbWorkspaceRepoStats) GetForks() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Forks
}

// GetForksOk returns a tuple with the Forks field value
// and a boolean to check if the value has been set.
func (o *DbWorkspaceRepoStats) GetForksOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Forks, true
}

// SetForks sets field value
func (o *DbWorkspaceRepoStats) SetForks(v int32) {
	o.Forks = v
}

// GetHealth returns the Health field value
func (o *DbWorkspaceRepoStats) GetHealth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Health
}

// GetHealthOk returns a tuple with the Health field value
// and a boolean to check if the value has been set.
func (o *DbWorkspaceRepoStats) GetHealthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Health, true
}

// SetHealth sets field value
func (o *DbWorkspaceRepoStats) SetHealth(v int32) {
	o.Health = v
}

func (o DbWorkspaceRepoStats) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DbWorkspaceRepoStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["stars"] = o.Stars
	toSerialize["forks"] = o.Forks
	toSerialize["health"] = o.Health
	return toSerialize, nil
}

type NullableDbWorkspaceRepoStats struct {
	value *DbWorkspaceRepoStats
	isSet bool
}

func (v NullableDbWorkspaceRepoStats) Get() *DbWorkspaceRepoStats {
	return v.value
}

func (v *NullableDbWorkspaceRepoStats) Set(val *DbWorkspaceRepoStats) {
	v.value = val
	v.isSet = true
}

func (v NullableDbWorkspaceRepoStats) IsSet() bool {
	return v.isSet
}

func (v *NullableDbWorkspaceRepoStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbWorkspaceRepoStats(val *DbWorkspaceRepoStats) *NullableDbWorkspaceRepoStats {
	return &NullableDbWorkspaceRepoStats{value: val, isSet: true}
}

func (v NullableDbWorkspaceRepoStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbWorkspaceRepoStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
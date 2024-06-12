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

// checks if the PageMetaDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PageMetaDto{}

// PageMetaDto struct for PageMetaDto
type PageMetaDto struct {
	// The current page
	Page int32 `json:"page"`
	// The number of items per page
	Limit int32 `json:"limit"`
	// The number of items in the collection
	ItemCount int32 `json:"itemCount"`
	// The number of pages in the collection
	PageCount int32 `json:"pageCount"`
	// Flag indicating if there is a previous page
	HasPreviousPage bool `json:"hasPreviousPage"`
	// Flag indicating if there is a next page
	HasNextPage bool `json:"hasNextPage"`
}

// NewPageMetaDto instantiates a new PageMetaDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageMetaDto(page int32, limit int32, itemCount int32, pageCount int32, hasPreviousPage bool, hasNextPage bool) *PageMetaDto {
	this := PageMetaDto{}
	this.Page = page
	this.Limit = limit
	this.ItemCount = itemCount
	this.PageCount = pageCount
	this.HasPreviousPage = hasPreviousPage
	this.HasNextPage = hasNextPage
	return &this
}

// NewPageMetaDtoWithDefaults instantiates a new PageMetaDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageMetaDtoWithDefaults() *PageMetaDto {
	this := PageMetaDto{}
	return &this
}

// GetPage returns the Page field value
func (o *PageMetaDto) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *PageMetaDto) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *PageMetaDto) SetPage(v int32) {
	o.Page = v
}

// GetLimit returns the Limit field value
func (o *PageMetaDto) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *PageMetaDto) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *PageMetaDto) SetLimit(v int32) {
	o.Limit = v
}

// GetItemCount returns the ItemCount field value
func (o *PageMetaDto) GetItemCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ItemCount
}

// GetItemCountOk returns a tuple with the ItemCount field value
// and a boolean to check if the value has been set.
func (o *PageMetaDto) GetItemCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemCount, true
}

// SetItemCount sets field value
func (o *PageMetaDto) SetItemCount(v int32) {
	o.ItemCount = v
}

// GetPageCount returns the PageCount field value
func (o *PageMetaDto) GetPageCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageCount
}

// GetPageCountOk returns a tuple with the PageCount field value
// and a boolean to check if the value has been set.
func (o *PageMetaDto) GetPageCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageCount, true
}

// SetPageCount sets field value
func (o *PageMetaDto) SetPageCount(v int32) {
	o.PageCount = v
}

// GetHasPreviousPage returns the HasPreviousPage field value
func (o *PageMetaDto) GetHasPreviousPage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasPreviousPage
}

// GetHasPreviousPageOk returns a tuple with the HasPreviousPage field value
// and a boolean to check if the value has been set.
func (o *PageMetaDto) GetHasPreviousPageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasPreviousPage, true
}

// SetHasPreviousPage sets field value
func (o *PageMetaDto) SetHasPreviousPage(v bool) {
	o.HasPreviousPage = v
}

// GetHasNextPage returns the HasNextPage field value
func (o *PageMetaDto) GetHasNextPage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasNextPage
}

// GetHasNextPageOk returns a tuple with the HasNextPage field value
// and a boolean to check if the value has been set.
func (o *PageMetaDto) GetHasNextPageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasNextPage, true
}

// SetHasNextPage sets field value
func (o *PageMetaDto) SetHasNextPage(v bool) {
	o.HasNextPage = v
}

func (o PageMetaDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PageMetaDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["page"] = o.Page
	toSerialize["limit"] = o.Limit
	toSerialize["itemCount"] = o.ItemCount
	toSerialize["pageCount"] = o.PageCount
	toSerialize["hasPreviousPage"] = o.HasPreviousPage
	toSerialize["hasNextPage"] = o.HasNextPage
	return toSerialize, nil
}

type NullablePageMetaDto struct {
	value *PageMetaDto
	isSet bool
}

func (v NullablePageMetaDto) Get() *PageMetaDto {
	return v.value
}

func (v *NullablePageMetaDto) Set(val *PageMetaDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePageMetaDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePageMetaDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageMetaDto(val *PageMetaDto) *NullablePageMetaDto {
	return &NullablePageMetaDto{value: val, isSet: true}
}

func (v NullablePageMetaDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageMetaDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
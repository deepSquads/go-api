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

// checks if the DbUserListContributorStat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DbUserListContributorStat{}

// DbUserListContributorStat struct for DbUserListContributorStat
type DbUserListContributorStat struct {
	// User list collaborator's login
	Login *string `json:"login,omitempty"`
	// Number of commits for login within the time range
	Commits int32 `json:"commits"`
	// Number of PRs created for login within the time range
	PrsCreated int32 `json:"prs_created"`
	// Number of PRs reviewed for login within the time range
	PrsReviewed int32 `json:"prs_reviewed"`
	// Number of issues created for login within the time range
	IssuesCreated int32 `json:"issues_created"`
	// Number of commit comments for login within the time range
	CommitComments int32 `json:"commit_comments"`
	// Number of issue comments for login within the time range
	IssueComments int32 `json:"issue_comments"`
	// Number of pr review comments for login within the time range
	PrReviewComments int32 `json:"pr_review_comments"`
	// Number of total comments for login within the time range
	Comments int32 `json:"comments"`
	// Number of total contributions for a login within the time range
	TotalContributions int32 `json:"total_contributions"`
}

// NewDbUserListContributorStat instantiates a new DbUserListContributorStat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbUserListContributorStat(commits int32, prsCreated int32, prsReviewed int32, issuesCreated int32, commitComments int32, issueComments int32, prReviewComments int32, comments int32, totalContributions int32) *DbUserListContributorStat {
	this := DbUserListContributorStat{}
	this.Commits = commits
	this.PrsCreated = prsCreated
	this.PrsReviewed = prsReviewed
	this.IssuesCreated = issuesCreated
	this.CommitComments = commitComments
	this.IssueComments = issueComments
	this.PrReviewComments = prReviewComments
	this.Comments = comments
	this.TotalContributions = totalContributions
	return &this
}

// NewDbUserListContributorStatWithDefaults instantiates a new DbUserListContributorStat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbUserListContributorStatWithDefaults() *DbUserListContributorStat {
	this := DbUserListContributorStat{}
	return &this
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *DbUserListContributorStat) GetLogin() string {
	if o == nil || IsNil(o.Login) {
		var ret string
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbUserListContributorStat) GetLoginOk() (*string, bool) {
	if o == nil || IsNil(o.Login) {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *DbUserListContributorStat) HasLogin() bool {
	if o != nil && !IsNil(o.Login) {
		return true
	}

	return false
}

// SetLogin gets a reference to the given string and assigns it to the Login field.
func (o *DbUserListContributorStat) SetLogin(v string) {
	o.Login = &v
}

// GetCommits returns the Commits field value
func (o *DbUserListContributorStat) GetCommits() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Commits
}

// GetCommitsOk returns a tuple with the Commits field value
// and a boolean to check if the value has been set.
func (o *DbUserListContributorStat) GetCommitsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Commits, true
}

// SetCommits sets field value
func (o *DbUserListContributorStat) SetCommits(v int32) {
	o.Commits = v
}

// GetPrsCreated returns the PrsCreated field value
func (o *DbUserListContributorStat) GetPrsCreated() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PrsCreated
}

// GetPrsCreatedOk returns a tuple with the PrsCreated field value
// and a boolean to check if the value has been set.
func (o *DbUserListContributorStat) GetPrsCreatedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrsCreated, true
}

// SetPrsCreated sets field value
func (o *DbUserListContributorStat) SetPrsCreated(v int32) {
	o.PrsCreated = v
}

// GetPrsReviewed returns the PrsReviewed field value
func (o *DbUserListContributorStat) GetPrsReviewed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PrsReviewed
}

// GetPrsReviewedOk returns a tuple with the PrsReviewed field value
// and a boolean to check if the value has been set.
func (o *DbUserListContributorStat) GetPrsReviewedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrsReviewed, true
}

// SetPrsReviewed sets field value
func (o *DbUserListContributorStat) SetPrsReviewed(v int32) {
	o.PrsReviewed = v
}

// GetIssuesCreated returns the IssuesCreated field value
func (o *DbUserListContributorStat) GetIssuesCreated() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IssuesCreated
}

// GetIssuesCreatedOk returns a tuple with the IssuesCreated field value
// and a boolean to check if the value has been set.
func (o *DbUserListContributorStat) GetIssuesCreatedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuesCreated, true
}

// SetIssuesCreated sets field value
func (o *DbUserListContributorStat) SetIssuesCreated(v int32) {
	o.IssuesCreated = v
}

// GetCommitComments returns the CommitComments field value
func (o *DbUserListContributorStat) GetCommitComments() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CommitComments
}

// GetCommitCommentsOk returns a tuple with the CommitComments field value
// and a boolean to check if the value has been set.
func (o *DbUserListContributorStat) GetCommitCommentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitComments, true
}

// SetCommitComments sets field value
func (o *DbUserListContributorStat) SetCommitComments(v int32) {
	o.CommitComments = v
}

// GetIssueComments returns the IssueComments field value
func (o *DbUserListContributorStat) GetIssueComments() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IssueComments
}

// GetIssueCommentsOk returns a tuple with the IssueComments field value
// and a boolean to check if the value has been set.
func (o *DbUserListContributorStat) GetIssueCommentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueComments, true
}

// SetIssueComments sets field value
func (o *DbUserListContributorStat) SetIssueComments(v int32) {
	o.IssueComments = v
}

// GetPrReviewComments returns the PrReviewComments field value
func (o *DbUserListContributorStat) GetPrReviewComments() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PrReviewComments
}

// GetPrReviewCommentsOk returns a tuple with the PrReviewComments field value
// and a boolean to check if the value has been set.
func (o *DbUserListContributorStat) GetPrReviewCommentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrReviewComments, true
}

// SetPrReviewComments sets field value
func (o *DbUserListContributorStat) SetPrReviewComments(v int32) {
	o.PrReviewComments = v
}

// GetComments returns the Comments field value
func (o *DbUserListContributorStat) GetComments() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value
// and a boolean to check if the value has been set.
func (o *DbUserListContributorStat) GetCommentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comments, true
}

// SetComments sets field value
func (o *DbUserListContributorStat) SetComments(v int32) {
	o.Comments = v
}

// GetTotalContributions returns the TotalContributions field value
func (o *DbUserListContributorStat) GetTotalContributions() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalContributions
}

// GetTotalContributionsOk returns a tuple with the TotalContributions field value
// and a boolean to check if the value has been set.
func (o *DbUserListContributorStat) GetTotalContributionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalContributions, true
}

// SetTotalContributions sets field value
func (o *DbUserListContributorStat) SetTotalContributions(v int32) {
	o.TotalContributions = v
}

func (o DbUserListContributorStat) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DbUserListContributorStat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Login) {
		toSerialize["login"] = o.Login
	}
	toSerialize["commits"] = o.Commits
	toSerialize["prs_created"] = o.PrsCreated
	toSerialize["prs_reviewed"] = o.PrsReviewed
	toSerialize["issues_created"] = o.IssuesCreated
	toSerialize["commit_comments"] = o.CommitComments
	toSerialize["issue_comments"] = o.IssueComments
	toSerialize["pr_review_comments"] = o.PrReviewComments
	toSerialize["comments"] = o.Comments
	toSerialize["total_contributions"] = o.TotalContributions
	return toSerialize, nil
}

type NullableDbUserListContributorStat struct {
	value *DbUserListContributorStat
	isSet bool
}

func (v NullableDbUserListContributorStat) Get() *DbUserListContributorStat {
	return v.value
}

func (v *NullableDbUserListContributorStat) Set(val *DbUserListContributorStat) {
	v.value = val
	v.isSet = true
}

func (v NullableDbUserListContributorStat) IsSet() bool {
	return v.isSet
}

func (v *NullableDbUserListContributorStat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbUserListContributorStat(val *DbUserListContributorStat) *NullableDbUserListContributorStat {
	return &NullableDbUserListContributorStat{value: val, isSet: true}
}

func (v NullableDbUserListContributorStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbUserListContributorStat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
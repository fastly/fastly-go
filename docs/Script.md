# Script

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique script identifier | [optional] 
**PageId** | Pointer to **string** | Parent page ID | [optional] 
**Source** | Pointer to **string** | Script source (inline or external URL) | [optional] 
**Urls** | Pointer to **[]string** | URLs where this script was observed | [optional] 
**FirstSeenAt** | Pointer to **time.Time** |  | [optional] 
**LastSeenAt** | Pointer to **time.Time** |  | [optional] 
**Justification** | Pointer to **string** | Reason for authorization decision | [optional] 
**CurrentHash** | Pointer to **string** | Current script content hash | [optional] 
**AuthorizedHash** | Pointer to **string** | Hash of authorized script content | [optional] 
**AuthorizationStatus** | Pointer to **string** | Script authorization status | [optional] 
**AuthorizedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewScript

`func NewScript() *Script`

NewScript instantiates a new Script object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptWithDefaults

`func NewScriptWithDefaults() *Script`

NewScriptWithDefaults instantiates a new Script object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Script) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Script) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Script) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Script) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPageId

`func (o *Script) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *Script) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *Script) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *Script) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetSource

`func (o *Script) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Script) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Script) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Script) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUrls

`func (o *Script) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *Script) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *Script) SetUrls(v []string)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *Script) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### GetFirstSeenAt

`func (o *Script) GetFirstSeenAt() time.Time`

GetFirstSeenAt returns the FirstSeenAt field if non-nil, zero value otherwise.

### GetFirstSeenAtOk

`func (o *Script) GetFirstSeenAtOk() (*time.Time, bool)`

GetFirstSeenAtOk returns a tuple with the FirstSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenAt

`func (o *Script) SetFirstSeenAt(v time.Time)`

SetFirstSeenAt sets FirstSeenAt field to given value.

### HasFirstSeenAt

`func (o *Script) HasFirstSeenAt() bool`

HasFirstSeenAt returns a boolean if a field has been set.

### GetLastSeenAt

`func (o *Script) GetLastSeenAt() time.Time`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *Script) GetLastSeenAtOk() (*time.Time, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *Script) SetLastSeenAt(v time.Time)`

SetLastSeenAt sets LastSeenAt field to given value.

### HasLastSeenAt

`func (o *Script) HasLastSeenAt() bool`

HasLastSeenAt returns a boolean if a field has been set.

### GetJustification

`func (o *Script) GetJustification() string`

GetJustification returns the Justification field if non-nil, zero value otherwise.

### GetJustificationOk

`func (o *Script) GetJustificationOk() (*string, bool)`

GetJustificationOk returns a tuple with the Justification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJustification

`func (o *Script) SetJustification(v string)`

SetJustification sets Justification field to given value.

### HasJustification

`func (o *Script) HasJustification() bool`

HasJustification returns a boolean if a field has been set.

### GetCurrentHash

`func (o *Script) GetCurrentHash() string`

GetCurrentHash returns the CurrentHash field if non-nil, zero value otherwise.

### GetCurrentHashOk

`func (o *Script) GetCurrentHashOk() (*string, bool)`

GetCurrentHashOk returns a tuple with the CurrentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentHash

`func (o *Script) SetCurrentHash(v string)`

SetCurrentHash sets CurrentHash field to given value.

### HasCurrentHash

`func (o *Script) HasCurrentHash() bool`

HasCurrentHash returns a boolean if a field has been set.

### GetAuthorizedHash

`func (o *Script) GetAuthorizedHash() string`

GetAuthorizedHash returns the AuthorizedHash field if non-nil, zero value otherwise.

### GetAuthorizedHashOk

`func (o *Script) GetAuthorizedHashOk() (*string, bool)`

GetAuthorizedHashOk returns a tuple with the AuthorizedHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedHash

`func (o *Script) SetAuthorizedHash(v string)`

SetAuthorizedHash sets AuthorizedHash field to given value.

### HasAuthorizedHash

`func (o *Script) HasAuthorizedHash() bool`

HasAuthorizedHash returns a boolean if a field has been set.

### GetAuthorizationStatus

`func (o *Script) GetAuthorizationStatus() string`

GetAuthorizationStatus returns the AuthorizationStatus field if non-nil, zero value otherwise.

### GetAuthorizationStatusOk

`func (o *Script) GetAuthorizationStatusOk() (*string, bool)`

GetAuthorizationStatusOk returns a tuple with the AuthorizationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationStatus

`func (o *Script) SetAuthorizationStatus(v string)`

SetAuthorizationStatus sets AuthorizationStatus field to given value.

### HasAuthorizationStatus

`func (o *Script) HasAuthorizationStatus() bool`

HasAuthorizationStatus returns a boolean if a field has been set.

### GetAuthorizedAt

`func (o *Script) GetAuthorizedAt() time.Time`

GetAuthorizedAt returns the AuthorizedAt field if non-nil, zero value otherwise.

### GetAuthorizedAtOk

`func (o *Script) GetAuthorizedAtOk() (*time.Time, bool)`

GetAuthorizedAtOk returns a tuple with the AuthorizedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedAt

`func (o *Script) SetAuthorizedAt(v time.Time)`

SetAuthorizedAt sets AuthorizedAt field to given value.

### HasAuthorizedAt

`func (o *Script) HasAuthorizedAt() bool`

HasAuthorizedAt returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



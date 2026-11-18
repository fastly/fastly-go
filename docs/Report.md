# Report

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** |  | [optional] 
**PolicyID** | Pointer to **string** |  | [optional] 
**BlockedURI** | Pointer to **string** |  | [optional] 
**DocumentURI** | Pointer to **string** |  | [optional] 
**ViolatedDirective** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewReport

`func NewReport() *Report`

NewReport instantiates a new Report object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportWithDefaults

`func NewReportWithDefaults() *Report`

NewReportWithDefaults instantiates a new Report object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *Report) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *Report) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *Report) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *Report) HasID() bool`

HasID returns a boolean if a field has been set.

### GetPolicyID

`func (o *Report) GetPolicyID() string`

GetPolicyID returns the PolicyID field if non-nil, zero value otherwise.

### GetPolicyIDOk

`func (o *Report) GetPolicyIDOk() (*string, bool)`

GetPolicyIDOk returns a tuple with the PolicyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyID

`func (o *Report) SetPolicyID(v string)`

SetPolicyID sets PolicyID field to given value.

### HasPolicyID

`func (o *Report) HasPolicyID() bool`

HasPolicyID returns a boolean if a field has been set.

### GetBlockedURI

`func (o *Report) GetBlockedURI() string`

GetBlockedURI returns the BlockedURI field if non-nil, zero value otherwise.

### GetBlockedURIOk

`func (o *Report) GetBlockedURIOk() (*string, bool)`

GetBlockedURIOk returns a tuple with the BlockedURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedURI

`func (o *Report) SetBlockedURI(v string)`

SetBlockedURI sets BlockedURI field to given value.

### HasBlockedURI

`func (o *Report) HasBlockedURI() bool`

HasBlockedURI returns a boolean if a field has been set.

### GetDocumentURI

`func (o *Report) GetDocumentURI() string`

GetDocumentURI returns the DocumentURI field if non-nil, zero value otherwise.

### GetDocumentURIOk

`func (o *Report) GetDocumentURIOk() (*string, bool)`

GetDocumentURIOk returns a tuple with the DocumentURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentURI

`func (o *Report) SetDocumentURI(v string)`

SetDocumentURI sets DocumentURI field to given value.

### HasDocumentURI

`func (o *Report) HasDocumentURI() bool`

HasDocumentURI returns a boolean if a field has been set.

### GetViolatedDirective

`func (o *Report) GetViolatedDirective() string`

GetViolatedDirective returns the ViolatedDirective field if non-nil, zero value otherwise.

### GetViolatedDirectiveOk

`func (o *Report) GetViolatedDirectiveOk() (*string, bool)`

GetViolatedDirectiveOk returns a tuple with the ViolatedDirective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolatedDirective

`func (o *Report) SetViolatedDirective(v string)`

SetViolatedDirective sets ViolatedDirective field to given value.

### HasViolatedDirective

`func (o *Report) HasViolatedDirective() bool`

HasViolatedDirective returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Report) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Report) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Report) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Report) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

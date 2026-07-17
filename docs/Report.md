# Report

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**PolicyId** | Pointer to **string** |  | [optional] 
**BlockedUri** | Pointer to **string** |  | [optional] 
**DocumentUri** | Pointer to **string** |  | [optional] 
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

### GetId

`func (o *Report) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Report) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Report) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Report) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPolicyId

`func (o *Report) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *Report) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *Report) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *Report) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetBlockedUri

`func (o *Report) GetBlockedUri() string`

GetBlockedUri returns the BlockedUri field if non-nil, zero value otherwise.

### GetBlockedUriOk

`func (o *Report) GetBlockedUriOk() (*string, bool)`

GetBlockedUriOk returns a tuple with the BlockedUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedUri

`func (o *Report) SetBlockedUri(v string)`

SetBlockedUri sets BlockedUri field to given value.

### HasBlockedUri

`func (o *Report) HasBlockedUri() bool`

HasBlockedUri returns a boolean if a field has been set.

### GetDocumentUri

`func (o *Report) GetDocumentUri() string`

GetDocumentUri returns the DocumentUri field if non-nil, zero value otherwise.

### GetDocumentUriOk

`func (o *Report) GetDocumentUriOk() (*string, bool)`

GetDocumentUriOk returns a tuple with the DocumentUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentUri

`func (o *Report) SetDocumentUri(v string)`

SetDocumentUri sets DocumentUri field to given value.

### HasDocumentUri

`func (o *Report) HasDocumentUri() bool`

HasDocumentUri returns a boolean if a field has been set.

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



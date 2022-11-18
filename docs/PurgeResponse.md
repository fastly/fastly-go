# PurgeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**ID** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewPurgeResponse

`func NewPurgeResponse() *PurgeResponse`

NewPurgeResponse instantiates a new PurgeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurgeResponseWithDefaults

`func NewPurgeResponseWithDefaults() *PurgeResponse`

NewPurgeResponseWithDefaults instantiates a new PurgeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PurgeResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PurgeResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PurgeResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PurgeResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetID

`func (o *PurgeResponse) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *PurgeResponse) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *PurgeResponse) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *PurgeResponse) HasID() bool`

HasID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

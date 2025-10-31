# PurgeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 

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

### GetId

`func (o *PurgeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PurgeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PurgeResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PurgeResponse) HasId() bool`

HasId returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



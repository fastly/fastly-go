# VersionCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** |  | [optional] [readonly] 
**ServiceId** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewVersionCreateResponse

`func NewVersionCreateResponse() *VersionCreateResponse`

NewVersionCreateResponse instantiates a new VersionCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionCreateResponseWithDefaults

`func NewVersionCreateResponseWithDefaults() *VersionCreateResponse`

NewVersionCreateResponseWithDefaults instantiates a new VersionCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *VersionCreateResponse) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *VersionCreateResponse) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *VersionCreateResponse) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *VersionCreateResponse) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetServiceId

`func (o *VersionCreateResponse) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *VersionCreateResponse) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *VersionCreateResponse) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *VersionCreateResponse) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



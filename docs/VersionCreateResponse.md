# VersionCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** |  | [optional] [readonly] 
**ServiceID** | Pointer to **string** |  | [optional] [readonly] 

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

### GetServiceID

`func (o *VersionCreateResponse) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *VersionCreateResponse) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *VersionCreateResponse) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *VersionCreateResponse) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

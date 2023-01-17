# ServiceIDAndVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceID** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewServiceIDAndVersion

`func NewServiceIDAndVersion() *ServiceIDAndVersion`

NewServiceIDAndVersion instantiates a new ServiceIDAndVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceIDAndVersionWithDefaults

`func NewServiceIDAndVersionWithDefaults() *ServiceIDAndVersion`

NewServiceIDAndVersionWithDefaults instantiates a new ServiceIDAndVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceID

`func (o *ServiceIDAndVersion) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *ServiceIDAndVersion) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *ServiceIDAndVersion) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *ServiceIDAndVersion) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetVersion

`func (o *ServiceIDAndVersion) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServiceIDAndVersion) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServiceIDAndVersion) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ServiceIDAndVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

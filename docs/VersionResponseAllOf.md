# VersionResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceID** | Pointer to **string** |  | [optional] [readonly] 
**Environments** | Pointer to [**[]Environment**](Environment.md) | A list of environments where the service has been deployed. | [optional] 

## Methods

### NewVersionResponseAllOf

`func NewVersionResponseAllOf() *VersionResponseAllOf`

NewVersionResponseAllOf instantiates a new VersionResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionResponseAllOfWithDefaults

`func NewVersionResponseAllOfWithDefaults() *VersionResponseAllOf`

NewVersionResponseAllOfWithDefaults instantiates a new VersionResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceID

`func (o *VersionResponseAllOf) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *VersionResponseAllOf) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *VersionResponseAllOf) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *VersionResponseAllOf) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetEnvironments

`func (o *VersionResponseAllOf) GetEnvironments() []Environment`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *VersionResponseAllOf) GetEnvironmentsOk() (*[]Environment, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *VersionResponseAllOf) SetEnvironments(v []Environment)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *VersionResponseAllOf) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

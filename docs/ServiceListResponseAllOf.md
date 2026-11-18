# ServiceListResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **int32** | Current [version](https://www.fastly.com/documentation/reference/api/services/version/) of the service. | [optional] 
**Versions** | Pointer to [**[]SchemasVersionResponse**](SchemasVersionResponse.md) | A list of [versions](https://www.fastly.com/documentation/reference/api/services/version/) associated with the service. | [optional] 
**Environments** | Pointer to [**[]Environment**](Environment.md) | A list of environments where the service has been deployed. | [optional] 

## Methods

### NewServiceListResponseAllOf

`func NewServiceListResponseAllOf() *ServiceListResponseAllOf`

NewServiceListResponseAllOf instantiates a new ServiceListResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceListResponseAllOfWithDefaults

`func NewServiceListResponseAllOfWithDefaults() *ServiceListResponseAllOf`

NewServiceListResponseAllOfWithDefaults instantiates a new ServiceListResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *ServiceListResponseAllOf) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *ServiceListResponseAllOf) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *ServiceListResponseAllOf) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *ServiceListResponseAllOf) HasID() bool`

HasID returns a boolean if a field has been set.

### GetVersion

`func (o *ServiceListResponseAllOf) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServiceListResponseAllOf) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServiceListResponseAllOf) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ServiceListResponseAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersions

`func (o *ServiceListResponseAllOf) GetVersions() []SchemasVersionResponse`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ServiceListResponseAllOf) GetVersionsOk() (*[]SchemasVersionResponse, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ServiceListResponseAllOf) SetVersions(v []SchemasVersionResponse)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *ServiceListResponseAllOf) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetEnvironments

`func (o *ServiceListResponseAllOf) GetEnvironments() []Environment`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *ServiceListResponseAllOf) GetEnvironmentsOk() (*[]Environment, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *ServiceListResponseAllOf) SetEnvironments(v []Environment)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *ServiceListResponseAllOf) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

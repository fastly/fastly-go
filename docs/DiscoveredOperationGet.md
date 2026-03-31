# DiscoveredOperationGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | The HTTP method for the operation. | 
**Domain** | **string** | The domain for the operation. | 
**Path** | **string** | The path for the operation, which may include path parameters. | 
**Id** | **string** | The unique identifier of the discovered operation. | [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The timestamp when the operation was last updated. | [optional] [readonly] 
**LastSeenAt** | Pointer to **time.Time** | The timestamp when the operation was last seen in traffic. | [optional] [readonly] 
**Rps** | Pointer to **float32** | Requests per second observed for this operation. | [optional] [readonly] 

## Methods

### NewDiscoveredOperationGet

`func NewDiscoveredOperationGet(method string, domain string, path string, id string, ) *DiscoveredOperationGet`

NewDiscoveredOperationGet instantiates a new DiscoveredOperationGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveredOperationGetWithDefaults

`func NewDiscoveredOperationGetWithDefaults() *DiscoveredOperationGet`

NewDiscoveredOperationGetWithDefaults instantiates a new DiscoveredOperationGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *DiscoveredOperationGet) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *DiscoveredOperationGet) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *DiscoveredOperationGet) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetDomain

`func (o *DiscoveredOperationGet) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DiscoveredOperationGet) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DiscoveredOperationGet) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetPath

`func (o *DiscoveredOperationGet) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DiscoveredOperationGet) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DiscoveredOperationGet) SetPath(v string)`

SetPath sets Path field to given value.


### GetId

`func (o *DiscoveredOperationGet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiscoveredOperationGet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiscoveredOperationGet) SetId(v string)`

SetId sets Id field to given value.


### GetUpdatedAt

`func (o *DiscoveredOperationGet) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DiscoveredOperationGet) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DiscoveredOperationGet) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DiscoveredOperationGet) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLastSeenAt

`func (o *DiscoveredOperationGet) GetLastSeenAt() time.Time`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *DiscoveredOperationGet) GetLastSeenAtOk() (*time.Time, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *DiscoveredOperationGet) SetLastSeenAt(v time.Time)`

SetLastSeenAt sets LastSeenAt field to given value.

### HasLastSeenAt

`func (o *DiscoveredOperationGet) HasLastSeenAt() bool`

HasLastSeenAt returns a boolean if a field has been set.

### GetRps

`func (o *DiscoveredOperationGet) GetRps() float32`

GetRps returns the Rps field if non-nil, zero value otherwise.

### GetRpsOk

`func (o *DiscoveredOperationGet) GetRpsOk() (*float32, bool)`

GetRpsOk returns a tuple with the Rps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRps

`func (o *DiscoveredOperationGet) SetRps(v float32)`

SetRps sets Rps field to given value.

### HasRps

`func (o *DiscoveredOperationGet) HasRps() bool`

HasRps returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



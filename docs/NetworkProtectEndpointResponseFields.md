# NetworkProtectEndpointResponseFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A read-only unique identifier for a Network Protect resource. | [readonly] 
**CreatedAt** | **time.Time** | An ISO 8601 timestamp in UTC. | [readonly] 
**UpdatedAt** | **time.Time** | An ISO 8601 timestamp in UTC. | [readonly] 

## Methods

### NewNetworkProtectEndpointResponseFields

`func NewNetworkProtectEndpointResponseFields(id string, createdAt time.Time, updatedAt time.Time, ) *NetworkProtectEndpointResponseFields`

NewNetworkProtectEndpointResponseFields instantiates a new NetworkProtectEndpointResponseFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectEndpointResponseFieldsWithDefaults

`func NewNetworkProtectEndpointResponseFieldsWithDefaults() *NetworkProtectEndpointResponseFields`

NewNetworkProtectEndpointResponseFieldsWithDefaults instantiates a new NetworkProtectEndpointResponseFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkProtectEndpointResponseFields) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkProtectEndpointResponseFields) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkProtectEndpointResponseFields) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *NetworkProtectEndpointResponseFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NetworkProtectEndpointResponseFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NetworkProtectEndpointResponseFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *NetworkProtectEndpointResponseFields) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NetworkProtectEndpointResponseFields) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NetworkProtectEndpointResponseFields) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



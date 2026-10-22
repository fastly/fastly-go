# NetworkProtectEndpointResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A human-readable name, unique within the customer. | 
**IpAddress** | **string** | The GRE destination address, unique within the customer. It cannot be changed after endpoint creation. | 
**Id** | **string** | A read-only unique identifier for a Network Protect resource. | [readonly] 
**CreatedAt** | **time.Time** | An ISO 8601 timestamp in UTC. | [readonly] 
**UpdatedAt** | **time.Time** | An ISO 8601 timestamp in UTC. | [readonly] 

## Methods

### NewNetworkProtectEndpointResponse

`func NewNetworkProtectEndpointResponse(name string, ipAddress string, id string, createdAt time.Time, updatedAt time.Time, ) *NetworkProtectEndpointResponse`

NewNetworkProtectEndpointResponse instantiates a new NetworkProtectEndpointResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectEndpointResponseWithDefaults

`func NewNetworkProtectEndpointResponseWithDefaults() *NetworkProtectEndpointResponse`

NewNetworkProtectEndpointResponseWithDefaults instantiates a new NetworkProtectEndpointResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkProtectEndpointResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkProtectEndpointResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkProtectEndpointResponse) SetName(v string)`

SetName sets Name field to given value.


### GetIpAddress

`func (o *NetworkProtectEndpointResponse) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *NetworkProtectEndpointResponse) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *NetworkProtectEndpointResponse) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetId

`func (o *NetworkProtectEndpointResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkProtectEndpointResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkProtectEndpointResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *NetworkProtectEndpointResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NetworkProtectEndpointResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NetworkProtectEndpointResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *NetworkProtectEndpointResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NetworkProtectEndpointResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NetworkProtectEndpointResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



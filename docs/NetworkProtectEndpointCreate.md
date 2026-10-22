# NetworkProtectEndpointCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A human-readable name, unique within the customer. | 
**IpAddress** | **string** | The GRE destination address, unique within the customer. It cannot be changed after endpoint creation. | 

## Methods

### NewNetworkProtectEndpointCreate

`func NewNetworkProtectEndpointCreate(name string, ipAddress string, ) *NetworkProtectEndpointCreate`

NewNetworkProtectEndpointCreate instantiates a new NetworkProtectEndpointCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectEndpointCreateWithDefaults

`func NewNetworkProtectEndpointCreateWithDefaults() *NetworkProtectEndpointCreate`

NewNetworkProtectEndpointCreateWithDefaults instantiates a new NetworkProtectEndpointCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkProtectEndpointCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkProtectEndpointCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkProtectEndpointCreate) SetName(v string)`

SetName sets Name field to given value.


### GetIpAddress

`func (o *NetworkProtectEndpointCreate) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *NetworkProtectEndpointCreate) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *NetworkProtectEndpointCreate) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



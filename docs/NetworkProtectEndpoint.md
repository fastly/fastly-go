# NetworkProtectEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A human-readable name, unique within the customer. | 
**IPAddress** | **string** | The GRE destination address, unique within the customer. It cannot be changed after endpoint creation. | 

## Methods

### NewNetworkProtectEndpoint

`func NewNetworkProtectEndpoint(name string, ipAddress string, ) *NetworkProtectEndpoint`

NewNetworkProtectEndpoint instantiates a new NetworkProtectEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectEndpointWithDefaults

`func NewNetworkProtectEndpointWithDefaults() *NetworkProtectEndpoint`

NewNetworkProtectEndpointWithDefaults instantiates a new NetworkProtectEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkProtectEndpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkProtectEndpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkProtectEndpoint) SetName(v string)`

SetName sets Name field to given value.


### GetIPAddress

`func (o *NetworkProtectEndpoint) GetIPAddress() string`

GetIPAddress returns the IPAddress field if non-nil, zero value otherwise.

### GetIPAddressOk

`func (o *NetworkProtectEndpoint) GetIPAddressOk() (*string, bool)`

GetIPAddressOk returns a tuple with the IPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPAddress

`func (o *NetworkProtectEndpoint) SetIPAddress(v string)`

SetIPAddress sets IPAddress field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

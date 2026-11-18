# NetworkProtectEndpointPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to [**NetworkProtectRegion**](NetworkProtectRegion.md) |  | [optional] [default to NETWORKPROTECTREGION_DEFAULT]
**Endpoints** | [**[]NetworkProtectPolicyEndpoint**](NetworkProtectPolicyEndpoint.md) | The endpoints serving a policy region. | 

## Methods

### NewNetworkProtectEndpointPolicy

`func NewNetworkProtectEndpointPolicy(endpoints []NetworkProtectPolicyEndpoint, ) *NetworkProtectEndpointPolicy`

NewNetworkProtectEndpointPolicy instantiates a new NetworkProtectEndpointPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectEndpointPolicyWithDefaults

`func NewNetworkProtectEndpointPolicyWithDefaults() *NetworkProtectEndpointPolicy`

NewNetworkProtectEndpointPolicyWithDefaults instantiates a new NetworkProtectEndpointPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *NetworkProtectEndpointPolicy) GetRegion() NetworkProtectRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NetworkProtectEndpointPolicy) GetRegionOk() (*NetworkProtectRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NetworkProtectEndpointPolicy) SetRegion(v NetworkProtectRegion)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *NetworkProtectEndpointPolicy) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetEndpoints

`func (o *NetworkProtectEndpointPolicy) GetEndpoints() []NetworkProtectPolicyEndpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *NetworkProtectEndpointPolicy) GetEndpointsOk() (*[]NetworkProtectPolicyEndpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *NetworkProtectEndpointPolicy) SetEndpoints(v []NetworkProtectPolicyEndpoint)`

SetEndpoints sets Endpoints field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

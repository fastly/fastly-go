# PublicIpList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to **[]string** | Fastly&#39;s IPv4 ranges. | [optional] 
**Ipv6Addresses** | Pointer to **[]string** | Fastly&#39;s IPv6 ranges. | [optional] 

## Methods

### NewPublicIpList

`func NewPublicIpList() *PublicIpList`

NewPublicIpList instantiates a new PublicIpList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicIpListWithDefaults

`func NewPublicIpListWithDefaults() *PublicIpList`

NewPublicIpListWithDefaults instantiates a new PublicIpList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *PublicIpList) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *PublicIpList) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *PublicIpList) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *PublicIpList) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetIpv6Addresses

`func (o *PublicIpList) GetIpv6Addresses() []string`

GetIpv6Addresses returns the Ipv6Addresses field if non-nil, zero value otherwise.

### GetIpv6AddressesOk

`func (o *PublicIpList) GetIpv6AddressesOk() (*[]string, bool)`

GetIpv6AddressesOk returns a tuple with the Ipv6Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addresses

`func (o *PublicIpList) SetIpv6Addresses(v []string)`

SetIpv6Addresses sets Ipv6Addresses field to given value.

### HasIpv6Addresses

`func (o *PublicIpList) HasIpv6Addresses() bool`

HasIpv6Addresses returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



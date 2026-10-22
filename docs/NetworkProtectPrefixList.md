# NetworkProtectPrefixList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A human-readable name, unique within the customer. | 
**Family** | [**NetworkProtectPrefixListFamily**](NetworkProtectPrefixListFamily.md) |  | 
**Addresses** | **[]string** | Complete canonical list contents. Bare addresses are canonicalized, CIDRs are masked, equivalent entries are deduplicated, and overlaps are retained. | 

## Methods

### NewNetworkProtectPrefixList

`func NewNetworkProtectPrefixList(name string, family NetworkProtectPrefixListFamily, addresses []string, ) *NetworkProtectPrefixList`

NewNetworkProtectPrefixList instantiates a new NetworkProtectPrefixList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectPrefixListWithDefaults

`func NewNetworkProtectPrefixListWithDefaults() *NetworkProtectPrefixList`

NewNetworkProtectPrefixListWithDefaults instantiates a new NetworkProtectPrefixList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkProtectPrefixList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkProtectPrefixList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkProtectPrefixList) SetName(v string)`

SetName sets Name field to given value.


### GetFamily

`func (o *NetworkProtectPrefixList) GetFamily() NetworkProtectPrefixListFamily`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *NetworkProtectPrefixList) GetFamilyOk() (*NetworkProtectPrefixListFamily, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *NetworkProtectPrefixList) SetFamily(v NetworkProtectPrefixListFamily)`

SetFamily sets Family field to given value.


### GetAddresses

`func (o *NetworkProtectPrefixList) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *NetworkProtectPrefixList) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *NetworkProtectPrefixList) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



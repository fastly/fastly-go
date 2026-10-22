# NetworkProtectPrefixListCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A human-readable name, unique within the customer. | 
**Family** | [**NetworkProtectPrefixListFamily**](NetworkProtectPrefixListFamily.md) |  | 
**Addresses** | **[]string** | Complete request contents. At most 1000 entries may be submitted; bare addresses are canonicalized and CIDRs are masked by the backend. | 

## Methods

### NewNetworkProtectPrefixListCreate

`func NewNetworkProtectPrefixListCreate(name string, family NetworkProtectPrefixListFamily, addresses []string, ) *NetworkProtectPrefixListCreate`

NewNetworkProtectPrefixListCreate instantiates a new NetworkProtectPrefixListCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectPrefixListCreateWithDefaults

`func NewNetworkProtectPrefixListCreateWithDefaults() *NetworkProtectPrefixListCreate`

NewNetworkProtectPrefixListCreateWithDefaults instantiates a new NetworkProtectPrefixListCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkProtectPrefixListCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkProtectPrefixListCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkProtectPrefixListCreate) SetName(v string)`

SetName sets Name field to given value.


### GetFamily

`func (o *NetworkProtectPrefixListCreate) GetFamily() NetworkProtectPrefixListFamily`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *NetworkProtectPrefixListCreate) GetFamilyOk() (*NetworkProtectPrefixListFamily, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *NetworkProtectPrefixListCreate) SetFamily(v NetworkProtectPrefixListFamily)`

SetFamily sets Family field to given value.


### GetAddresses

`func (o *NetworkProtectPrefixListCreate) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *NetworkProtectPrefixListCreate) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *NetworkProtectPrefixListCreate) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



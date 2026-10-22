# NetworkProtectPrefixListUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A human-readable name, unique within the customer. | [optional] 
**Addresses** | Pointer to **[]string** | Complete request contents. At most 1000 entries may be submitted; bare addresses are canonicalized and CIDRs are masked by the backend. | [optional] 

## Methods

### NewNetworkProtectPrefixListUpdate

`func NewNetworkProtectPrefixListUpdate() *NetworkProtectPrefixListUpdate`

NewNetworkProtectPrefixListUpdate instantiates a new NetworkProtectPrefixListUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectPrefixListUpdateWithDefaults

`func NewNetworkProtectPrefixListUpdateWithDefaults() *NetworkProtectPrefixListUpdate`

NewNetworkProtectPrefixListUpdateWithDefaults instantiates a new NetworkProtectPrefixListUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkProtectPrefixListUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkProtectPrefixListUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkProtectPrefixListUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkProtectPrefixListUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddresses

`func (o *NetworkProtectPrefixListUpdate) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *NetworkProtectPrefixListUpdate) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *NetworkProtectPrefixListUpdate) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *NetworkProtectPrefixListUpdate) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



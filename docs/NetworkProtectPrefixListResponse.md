# NetworkProtectPrefixListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A human-readable name, unique within the customer. | 
**Family** | [**NetworkProtectPrefixListFamily**](NetworkProtectPrefixListFamily.md) |  | 
**Addresses** | **[]string** | Complete canonical list contents. Bare addresses are canonicalized, CIDRs are masked, equivalent entries are deduplicated, and overlaps are retained. | 
**ID** | **string** | A read-only unique identifier for a Network Protect resource. | [readonly] 
**CreatedAt** | **time.Time** | An ISO 8601 timestamp in UTC. | [readonly] 
**UpdatedAt** | **time.Time** | An ISO 8601 timestamp in UTC. | [readonly] 

## Methods

### NewNetworkProtectPrefixListResponse

`func NewNetworkProtectPrefixListResponse(name string, family NetworkProtectPrefixListFamily, addresses []string, id string, createdAt time.Time, updatedAt time.Time, ) *NetworkProtectPrefixListResponse`

NewNetworkProtectPrefixListResponse instantiates a new NetworkProtectPrefixListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectPrefixListResponseWithDefaults

`func NewNetworkProtectPrefixListResponseWithDefaults() *NetworkProtectPrefixListResponse`

NewNetworkProtectPrefixListResponseWithDefaults instantiates a new NetworkProtectPrefixListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkProtectPrefixListResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkProtectPrefixListResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkProtectPrefixListResponse) SetName(v string)`

SetName sets Name field to given value.


### GetFamily

`func (o *NetworkProtectPrefixListResponse) GetFamily() NetworkProtectPrefixListFamily`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *NetworkProtectPrefixListResponse) GetFamilyOk() (*NetworkProtectPrefixListFamily, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *NetworkProtectPrefixListResponse) SetFamily(v NetworkProtectPrefixListFamily)`

SetFamily sets Family field to given value.


### GetAddresses

`func (o *NetworkProtectPrefixListResponse) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *NetworkProtectPrefixListResponse) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *NetworkProtectPrefixListResponse) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetID

`func (o *NetworkProtectPrefixListResponse) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *NetworkProtectPrefixListResponse) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *NetworkProtectPrefixListResponse) SetID(v string)`

SetID sets ID field to given value.


### GetCreatedAt

`func (o *NetworkProtectPrefixListResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NetworkProtectPrefixListResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NetworkProtectPrefixListResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *NetworkProtectPrefixListResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NetworkProtectPrefixListResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NetworkProtectPrefixListResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

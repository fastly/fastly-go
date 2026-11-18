# NetworkProtectNetworkResponseFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | A read-only unique identifier for a Network Protect resource. | [readonly] 
**RoaStatus** | Pointer to [**NetworkProtectRoaStatus**](NetworkProtectRoaStatus.md) |  | [optional] 
**RoaCheckedAt** | Pointer to **NullableTime** | When Fastly last completed a ROA check. Null while the status is pending. A failed check leaves this unchanged; an old value can indicate delayed or failing checks. Null also means the controller did not report a check time. | [optional] [readonly] 
**RoaDataAsOf** | Pointer to **NullableTime** | When the RPKI dataset used by the last check was built. This is not the same as `roa_checked_at`; a recent check can use an older dataset. Null while the status is pending or when no dataset build time was reported. | [optional] [readonly] 
**LoaStatus** | Pointer to [**NetworkProtectLoaStatus**](NetworkProtectLoaStatus.md) |  | [optional] 
**CreatedAt** | **time.Time** | An ISO 8601 timestamp in UTC. | [readonly] 
**UpdatedAt** | **time.Time** | An ISO 8601 timestamp in UTC. | [readonly] 

## Methods

### NewNetworkProtectNetworkResponseFields

`func NewNetworkProtectNetworkResponseFields(id string, createdAt time.Time, updatedAt time.Time, ) *NetworkProtectNetworkResponseFields`

NewNetworkProtectNetworkResponseFields instantiates a new NetworkProtectNetworkResponseFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectNetworkResponseFieldsWithDefaults

`func NewNetworkProtectNetworkResponseFieldsWithDefaults() *NetworkProtectNetworkResponseFields`

NewNetworkProtectNetworkResponseFieldsWithDefaults instantiates a new NetworkProtectNetworkResponseFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *NetworkProtectNetworkResponseFields) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *NetworkProtectNetworkResponseFields) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *NetworkProtectNetworkResponseFields) SetID(v string)`

SetID sets ID field to given value.


### GetRoaStatus

`func (o *NetworkProtectNetworkResponseFields) GetRoaStatus() NetworkProtectRoaStatus`

GetRoaStatus returns the RoaStatus field if non-nil, zero value otherwise.

### GetRoaStatusOk

`func (o *NetworkProtectNetworkResponseFields) GetRoaStatusOk() (*NetworkProtectRoaStatus, bool)`

GetRoaStatusOk returns a tuple with the RoaStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoaStatus

`func (o *NetworkProtectNetworkResponseFields) SetRoaStatus(v NetworkProtectRoaStatus)`

SetRoaStatus sets RoaStatus field to given value.

### HasRoaStatus

`func (o *NetworkProtectNetworkResponseFields) HasRoaStatus() bool`

HasRoaStatus returns a boolean if a field has been set.

### GetRoaCheckedAt

`func (o *NetworkProtectNetworkResponseFields) GetRoaCheckedAt() time.Time`

GetRoaCheckedAt returns the RoaCheckedAt field if non-nil, zero value otherwise.

### GetRoaCheckedAtOk

`func (o *NetworkProtectNetworkResponseFields) GetRoaCheckedAtOk() (*time.Time, bool)`

GetRoaCheckedAtOk returns a tuple with the RoaCheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoaCheckedAt

`func (o *NetworkProtectNetworkResponseFields) SetRoaCheckedAt(v time.Time)`

SetRoaCheckedAt sets RoaCheckedAt field to given value.

### HasRoaCheckedAt

`func (o *NetworkProtectNetworkResponseFields) HasRoaCheckedAt() bool`

HasRoaCheckedAt returns a boolean if a field has been set.

### SetRoaCheckedAtNil

`func (o *NetworkProtectNetworkResponseFields) SetRoaCheckedAtNil(b bool)`

 SetRoaCheckedAtNil sets the value for RoaCheckedAt to be an explicit nil

### UnsetRoaCheckedAt
`func (o *NetworkProtectNetworkResponseFields) UnsetRoaCheckedAt()`

UnsetRoaCheckedAt ensures that no value is present for RoaCheckedAt, not even an explicit nil
### GetRoaDataAsOf

`func (o *NetworkProtectNetworkResponseFields) GetRoaDataAsOf() time.Time`

GetRoaDataAsOf returns the RoaDataAsOf field if non-nil, zero value otherwise.

### GetRoaDataAsOfOk

`func (o *NetworkProtectNetworkResponseFields) GetRoaDataAsOfOk() (*time.Time, bool)`

GetRoaDataAsOfOk returns a tuple with the RoaDataAsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoaDataAsOf

`func (o *NetworkProtectNetworkResponseFields) SetRoaDataAsOf(v time.Time)`

SetRoaDataAsOf sets RoaDataAsOf field to given value.

### HasRoaDataAsOf

`func (o *NetworkProtectNetworkResponseFields) HasRoaDataAsOf() bool`

HasRoaDataAsOf returns a boolean if a field has been set.

### SetRoaDataAsOfNil

`func (o *NetworkProtectNetworkResponseFields) SetRoaDataAsOfNil(b bool)`

 SetRoaDataAsOfNil sets the value for RoaDataAsOf to be an explicit nil

### UnsetRoaDataAsOf
`func (o *NetworkProtectNetworkResponseFields) UnsetRoaDataAsOf()`

UnsetRoaDataAsOf ensures that no value is present for RoaDataAsOf, not even an explicit nil
### GetLoaStatus

`func (o *NetworkProtectNetworkResponseFields) GetLoaStatus() NetworkProtectLoaStatus`

GetLoaStatus returns the LoaStatus field if non-nil, zero value otherwise.

### GetLoaStatusOk

`func (o *NetworkProtectNetworkResponseFields) GetLoaStatusOk() (*NetworkProtectLoaStatus, bool)`

GetLoaStatusOk returns a tuple with the LoaStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoaStatus

`func (o *NetworkProtectNetworkResponseFields) SetLoaStatus(v NetworkProtectLoaStatus)`

SetLoaStatus sets LoaStatus field to given value.

### HasLoaStatus

`func (o *NetworkProtectNetworkResponseFields) HasLoaStatus() bool`

HasLoaStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NetworkProtectNetworkResponseFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NetworkProtectNetworkResponseFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NetworkProtectNetworkResponseFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *NetworkProtectNetworkResponseFields) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NetworkProtectNetworkResponseFields) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NetworkProtectNetworkResponseFields) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

# NetworkProtectNetworkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Optional descriptive text for the network. This metadata does not change traffic handling. Null clears it on update. | [optional] 
**Name** | **string** | A human-readable name, unique within the customer. | 
**Prefix** | **string** | The CIDR block to protect, masked to its network address. It cannot be changed after creation. | 
**RuleSetId** | Pointer to **NullableString** | The backend-assigned rule set applied to this network. | [optional] 
**EndpointPolicies** | [**[]NetworkProtectEndpointPolicy**](NetworkProtectEndpointPolicy.md) | Complete endpoint policy configuration. A network has at most one policy per region and may have no policies. | 
**Advertise** | **bool** | Desired advertisement state, which is not the same as observed route state. Enabling it requires approved LOA, valid ROA, and a ready default-region delivery policy. A later ROA failure does not turn an existing advertisement off; losing LOA approval or the last ready default endpoint does. | [default to false]
**Id** | **string** | A read-only unique identifier for a Network Protect resource. | [readonly] 
**RoaStatus** | Pointer to [**NetworkProtectRoaStatus**](NetworkProtectRoaStatus.md) |  | [optional] 
**RoaCheckedAt** | Pointer to **NullableTime** | When Fastly last completed a ROA check. Null while the status is pending. A failed check leaves this unchanged; an old value can indicate delayed or failing checks. Null also means the controller did not report a check time. | [optional] [readonly] 
**RoaDataAsOf** | Pointer to **NullableTime** | When the RPKI dataset used by the last check was built. This is not the same as `roa_checked_at`; a recent check can use an older dataset. Null while the status is pending or when no dataset build time was reported. | [optional] [readonly] 
**LoaStatus** | Pointer to [**NetworkProtectLoaStatus**](NetworkProtectLoaStatus.md) |  | [optional] 
**CreatedAt** | **time.Time** | An ISO 8601 timestamp in UTC. | [readonly] 
**UpdatedAt** | **time.Time** | An ISO 8601 timestamp in UTC. | [readonly] 

## Methods

### NewNetworkProtectNetworkResponse

`func NewNetworkProtectNetworkResponse(name string, prefix string, endpointPolicies []NetworkProtectEndpointPolicy, advertise bool, id string, createdAt time.Time, updatedAt time.Time, ) *NetworkProtectNetworkResponse`

NewNetworkProtectNetworkResponse instantiates a new NetworkProtectNetworkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectNetworkResponseWithDefaults

`func NewNetworkProtectNetworkResponseWithDefaults() *NetworkProtectNetworkResponse`

NewNetworkProtectNetworkResponseWithDefaults instantiates a new NetworkProtectNetworkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *NetworkProtectNetworkResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkProtectNetworkResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkProtectNetworkResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkProtectNetworkResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *NetworkProtectNetworkResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *NetworkProtectNetworkResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *NetworkProtectNetworkResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkProtectNetworkResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkProtectNetworkResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPrefix

`func (o *NetworkProtectNetworkResponse) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *NetworkProtectNetworkResponse) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *NetworkProtectNetworkResponse) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetRuleSetId

`func (o *NetworkProtectNetworkResponse) GetRuleSetId() string`

GetRuleSetId returns the RuleSetId field if non-nil, zero value otherwise.

### GetRuleSetIdOk

`func (o *NetworkProtectNetworkResponse) GetRuleSetIdOk() (*string, bool)`

GetRuleSetIdOk returns a tuple with the RuleSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleSetId

`func (o *NetworkProtectNetworkResponse) SetRuleSetId(v string)`

SetRuleSetId sets RuleSetId field to given value.

### HasRuleSetId

`func (o *NetworkProtectNetworkResponse) HasRuleSetId() bool`

HasRuleSetId returns a boolean if a field has been set.

### SetRuleSetIdNil

`func (o *NetworkProtectNetworkResponse) SetRuleSetIdNil(b bool)`

 SetRuleSetIdNil sets the value for RuleSetId to be an explicit nil

### UnsetRuleSetId
`func (o *NetworkProtectNetworkResponse) UnsetRuleSetId()`

UnsetRuleSetId ensures that no value is present for RuleSetId, not even an explicit nil
### GetEndpointPolicies

`func (o *NetworkProtectNetworkResponse) GetEndpointPolicies() []NetworkProtectEndpointPolicy`

GetEndpointPolicies returns the EndpointPolicies field if non-nil, zero value otherwise.

### GetEndpointPoliciesOk

`func (o *NetworkProtectNetworkResponse) GetEndpointPoliciesOk() (*[]NetworkProtectEndpointPolicy, bool)`

GetEndpointPoliciesOk returns a tuple with the EndpointPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointPolicies

`func (o *NetworkProtectNetworkResponse) SetEndpointPolicies(v []NetworkProtectEndpointPolicy)`

SetEndpointPolicies sets EndpointPolicies field to given value.


### GetAdvertise

`func (o *NetworkProtectNetworkResponse) GetAdvertise() bool`

GetAdvertise returns the Advertise field if non-nil, zero value otherwise.

### GetAdvertiseOk

`func (o *NetworkProtectNetworkResponse) GetAdvertiseOk() (*bool, bool)`

GetAdvertiseOk returns a tuple with the Advertise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertise

`func (o *NetworkProtectNetworkResponse) SetAdvertise(v bool)`

SetAdvertise sets Advertise field to given value.


### GetId

`func (o *NetworkProtectNetworkResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkProtectNetworkResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkProtectNetworkResponse) SetId(v string)`

SetId sets Id field to given value.


### GetRoaStatus

`func (o *NetworkProtectNetworkResponse) GetRoaStatus() NetworkProtectRoaStatus`

GetRoaStatus returns the RoaStatus field if non-nil, zero value otherwise.

### GetRoaStatusOk

`func (o *NetworkProtectNetworkResponse) GetRoaStatusOk() (*NetworkProtectRoaStatus, bool)`

GetRoaStatusOk returns a tuple with the RoaStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoaStatus

`func (o *NetworkProtectNetworkResponse) SetRoaStatus(v NetworkProtectRoaStatus)`

SetRoaStatus sets RoaStatus field to given value.

### HasRoaStatus

`func (o *NetworkProtectNetworkResponse) HasRoaStatus() bool`

HasRoaStatus returns a boolean if a field has been set.

### GetRoaCheckedAt

`func (o *NetworkProtectNetworkResponse) GetRoaCheckedAt() time.Time`

GetRoaCheckedAt returns the RoaCheckedAt field if non-nil, zero value otherwise.

### GetRoaCheckedAtOk

`func (o *NetworkProtectNetworkResponse) GetRoaCheckedAtOk() (*time.Time, bool)`

GetRoaCheckedAtOk returns a tuple with the RoaCheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoaCheckedAt

`func (o *NetworkProtectNetworkResponse) SetRoaCheckedAt(v time.Time)`

SetRoaCheckedAt sets RoaCheckedAt field to given value.

### HasRoaCheckedAt

`func (o *NetworkProtectNetworkResponse) HasRoaCheckedAt() bool`

HasRoaCheckedAt returns a boolean if a field has been set.

### SetRoaCheckedAtNil

`func (o *NetworkProtectNetworkResponse) SetRoaCheckedAtNil(b bool)`

 SetRoaCheckedAtNil sets the value for RoaCheckedAt to be an explicit nil

### UnsetRoaCheckedAt
`func (o *NetworkProtectNetworkResponse) UnsetRoaCheckedAt()`

UnsetRoaCheckedAt ensures that no value is present for RoaCheckedAt, not even an explicit nil
### GetRoaDataAsOf

`func (o *NetworkProtectNetworkResponse) GetRoaDataAsOf() time.Time`

GetRoaDataAsOf returns the RoaDataAsOf field if non-nil, zero value otherwise.

### GetRoaDataAsOfOk

`func (o *NetworkProtectNetworkResponse) GetRoaDataAsOfOk() (*time.Time, bool)`

GetRoaDataAsOfOk returns a tuple with the RoaDataAsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoaDataAsOf

`func (o *NetworkProtectNetworkResponse) SetRoaDataAsOf(v time.Time)`

SetRoaDataAsOf sets RoaDataAsOf field to given value.

### HasRoaDataAsOf

`func (o *NetworkProtectNetworkResponse) HasRoaDataAsOf() bool`

HasRoaDataAsOf returns a boolean if a field has been set.

### SetRoaDataAsOfNil

`func (o *NetworkProtectNetworkResponse) SetRoaDataAsOfNil(b bool)`

 SetRoaDataAsOfNil sets the value for RoaDataAsOf to be an explicit nil

### UnsetRoaDataAsOf
`func (o *NetworkProtectNetworkResponse) UnsetRoaDataAsOf()`

UnsetRoaDataAsOf ensures that no value is present for RoaDataAsOf, not even an explicit nil
### GetLoaStatus

`func (o *NetworkProtectNetworkResponse) GetLoaStatus() NetworkProtectLoaStatus`

GetLoaStatus returns the LoaStatus field if non-nil, zero value otherwise.

### GetLoaStatusOk

`func (o *NetworkProtectNetworkResponse) GetLoaStatusOk() (*NetworkProtectLoaStatus, bool)`

GetLoaStatusOk returns a tuple with the LoaStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoaStatus

`func (o *NetworkProtectNetworkResponse) SetLoaStatus(v NetworkProtectLoaStatus)`

SetLoaStatus sets LoaStatus field to given value.

### HasLoaStatus

`func (o *NetworkProtectNetworkResponse) HasLoaStatus() bool`

HasLoaStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NetworkProtectNetworkResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NetworkProtectNetworkResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NetworkProtectNetworkResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *NetworkProtectNetworkResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NetworkProtectNetworkResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NetworkProtectNetworkResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



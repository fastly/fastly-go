# DiscoveredOperationGetExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the discovered operation. | [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The timestamp when the operation was last updated. | [optional] [readonly] 
**LastSeenAt** | Pointer to **time.Time** | The timestamp when the operation was last seen in traffic. | [optional] [readonly] 
**Rps** | Pointer to **float32** | Requests per second observed for this operation. | [optional] [readonly] 

## Methods

### NewDiscoveredOperationGetExtra

`func NewDiscoveredOperationGetExtra(id string, ) *DiscoveredOperationGetExtra`

NewDiscoveredOperationGetExtra instantiates a new DiscoveredOperationGetExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveredOperationGetExtraWithDefaults

`func NewDiscoveredOperationGetExtraWithDefaults() *DiscoveredOperationGetExtra`

NewDiscoveredOperationGetExtraWithDefaults instantiates a new DiscoveredOperationGetExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DiscoveredOperationGetExtra) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiscoveredOperationGetExtra) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiscoveredOperationGetExtra) SetId(v string)`

SetId sets Id field to given value.


### GetUpdatedAt

`func (o *DiscoveredOperationGetExtra) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DiscoveredOperationGetExtra) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DiscoveredOperationGetExtra) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DiscoveredOperationGetExtra) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLastSeenAt

`func (o *DiscoveredOperationGetExtra) GetLastSeenAt() time.Time`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *DiscoveredOperationGetExtra) GetLastSeenAtOk() (*time.Time, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *DiscoveredOperationGetExtra) SetLastSeenAt(v time.Time)`

SetLastSeenAt sets LastSeenAt field to given value.

### HasLastSeenAt

`func (o *DiscoveredOperationGetExtra) HasLastSeenAt() bool`

HasLastSeenAt returns a boolean if a field has been set.

### GetRps

`func (o *DiscoveredOperationGetExtra) GetRps() float32`

GetRps returns the Rps field if non-nil, zero value otherwise.

### GetRpsOk

`func (o *DiscoveredOperationGetExtra) GetRpsOk() (*float32, bool)`

GetRpsOk returns a tuple with the Rps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRps

`func (o *DiscoveredOperationGetExtra) SetRps(v float32)`

SetRps sets Rps field to given value.

### HasRps

`func (o *DiscoveredOperationGetExtra) HasRps() bool`

HasRps returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



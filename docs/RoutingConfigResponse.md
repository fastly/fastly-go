# RoutingConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**Id** | Pointer to **string** | Alphanumeric string identifying the routing config. | [optional] [readonly] 
**Name** | Pointer to **string** | The user-defined name for the routing config. | [optional] 
**State** | Pointer to [**RoutingConfigState**](RoutingConfigState.md) |  | [optional] 
**ActivatedAt** | Pointer to **NullableTime** | Timestamp of when the version was most recently activated. `null` if the version has never been activated. | [optional] [readonly] 
**Links** | Pointer to **map[string]string** | HATEOAS links to related resources. | [optional] [readonly] 

## Methods

### NewRoutingConfigResponse

`func NewRoutingConfigResponse() *RoutingConfigResponse`

NewRoutingConfigResponse instantiates a new RoutingConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingConfigResponseWithDefaults

`func NewRoutingConfigResponseWithDefaults() *RoutingConfigResponse`

NewRoutingConfigResponseWithDefaults instantiates a new RoutingConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *RoutingConfigResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RoutingConfigResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RoutingConfigResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RoutingConfigResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *RoutingConfigResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *RoutingConfigResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *RoutingConfigResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RoutingConfigResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RoutingConfigResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RoutingConfigResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *RoutingConfigResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *RoutingConfigResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetId

`func (o *RoutingConfigResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoutingConfigResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoutingConfigResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RoutingConfigResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RoutingConfigResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoutingConfigResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoutingConfigResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoutingConfigResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *RoutingConfigResponse) GetState() RoutingConfigState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RoutingConfigResponse) GetStateOk() (*RoutingConfigState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RoutingConfigResponse) SetState(v RoutingConfigState)`

SetState sets State field to given value.

### HasState

`func (o *RoutingConfigResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetActivatedAt

`func (o *RoutingConfigResponse) GetActivatedAt() time.Time`

GetActivatedAt returns the ActivatedAt field if non-nil, zero value otherwise.

### GetActivatedAtOk

`func (o *RoutingConfigResponse) GetActivatedAtOk() (*time.Time, bool)`

GetActivatedAtOk returns a tuple with the ActivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatedAt

`func (o *RoutingConfigResponse) SetActivatedAt(v time.Time)`

SetActivatedAt sets ActivatedAt field to given value.

### HasActivatedAt

`func (o *RoutingConfigResponse) HasActivatedAt() bool`

HasActivatedAt returns a boolean if a field has been set.

### SetActivatedAtNil

`func (o *RoutingConfigResponse) SetActivatedAtNil(b bool)`

 SetActivatedAtNil sets the value for ActivatedAt to be an explicit nil

### UnsetActivatedAt
`func (o *RoutingConfigResponse) UnsetActivatedAt()`

UnsetActivatedAt ensures that no value is present for ActivatedAt, not even an explicit nil
### GetLinks

`func (o *RoutingConfigResponse) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RoutingConfigResponse) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RoutingConfigResponse) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RoutingConfigResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



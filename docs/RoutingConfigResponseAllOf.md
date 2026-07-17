# RoutingConfigResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Alphanumeric string identifying the routing config. | [optional] [readonly] 
**Name** | Pointer to **string** | The user-defined name for the routing config. | [optional] 
**State** | Pointer to [**RoutingConfigState**](RoutingConfigState.md) |  | [optional] 
**ActivatedAt** | Pointer to **NullableTime** | Timestamp of when the version was most recently activated. `null` if the version has never been activated. | [optional] [readonly] 
**Links** | Pointer to **map[string]string** | HATEOAS links to related resources. | [optional] [readonly] 

## Methods

### NewRoutingConfigResponseAllOf

`func NewRoutingConfigResponseAllOf() *RoutingConfigResponseAllOf`

NewRoutingConfigResponseAllOf instantiates a new RoutingConfigResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingConfigResponseAllOfWithDefaults

`func NewRoutingConfigResponseAllOfWithDefaults() *RoutingConfigResponseAllOf`

NewRoutingConfigResponseAllOfWithDefaults instantiates a new RoutingConfigResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RoutingConfigResponseAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoutingConfigResponseAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoutingConfigResponseAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RoutingConfigResponseAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RoutingConfigResponseAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoutingConfigResponseAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoutingConfigResponseAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoutingConfigResponseAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *RoutingConfigResponseAllOf) GetState() RoutingConfigState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RoutingConfigResponseAllOf) GetStateOk() (*RoutingConfigState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RoutingConfigResponseAllOf) SetState(v RoutingConfigState)`

SetState sets State field to given value.

### HasState

`func (o *RoutingConfigResponseAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetActivatedAt

`func (o *RoutingConfigResponseAllOf) GetActivatedAt() time.Time`

GetActivatedAt returns the ActivatedAt field if non-nil, zero value otherwise.

### GetActivatedAtOk

`func (o *RoutingConfigResponseAllOf) GetActivatedAtOk() (*time.Time, bool)`

GetActivatedAtOk returns a tuple with the ActivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatedAt

`func (o *RoutingConfigResponseAllOf) SetActivatedAt(v time.Time)`

SetActivatedAt sets ActivatedAt field to given value.

### HasActivatedAt

`func (o *RoutingConfigResponseAllOf) HasActivatedAt() bool`

HasActivatedAt returns a boolean if a field has been set.

### SetActivatedAtNil

`func (o *RoutingConfigResponseAllOf) SetActivatedAtNil(b bool)`

 SetActivatedAtNil sets the value for ActivatedAt to be an explicit nil

### UnsetActivatedAt
`func (o *RoutingConfigResponseAllOf) UnsetActivatedAt()`

UnsetActivatedAt ensures that no value is present for ActivatedAt, not even an explicit nil
### GetLinks

`func (o *RoutingConfigResponseAllOf) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RoutingConfigResponseAllOf) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RoutingConfigResponseAllOf) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RoutingConfigResponseAllOf) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



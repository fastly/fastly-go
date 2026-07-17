# RoutingConfigVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Alphanumeric string identifying the version. | [optional] [readonly] 
**Comment** | Pointer to **string** | A freeform comment describing the version. | [optional] 
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**ActivatedAt** | Pointer to **NullableTime** | Timestamp of when the version was most recently activated. `null` if the version has never been activated. | [optional] [readonly] 

## Methods

### NewRoutingConfigVersionResponse

`func NewRoutingConfigVersionResponse() *RoutingConfigVersionResponse`

NewRoutingConfigVersionResponse instantiates a new RoutingConfigVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingConfigVersionResponseWithDefaults

`func NewRoutingConfigVersionResponseWithDefaults() *RoutingConfigVersionResponse`

NewRoutingConfigVersionResponseWithDefaults instantiates a new RoutingConfigVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RoutingConfigVersionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoutingConfigVersionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoutingConfigVersionResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RoutingConfigVersionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetComment

`func (o *RoutingConfigVersionResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RoutingConfigVersionResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RoutingConfigVersionResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RoutingConfigVersionResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RoutingConfigVersionResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RoutingConfigVersionResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RoutingConfigVersionResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RoutingConfigVersionResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *RoutingConfigVersionResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *RoutingConfigVersionResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetActivatedAt

`func (o *RoutingConfigVersionResponse) GetActivatedAt() time.Time`

GetActivatedAt returns the ActivatedAt field if non-nil, zero value otherwise.

### GetActivatedAtOk

`func (o *RoutingConfigVersionResponse) GetActivatedAtOk() (*time.Time, bool)`

GetActivatedAtOk returns a tuple with the ActivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatedAt

`func (o *RoutingConfigVersionResponse) SetActivatedAt(v time.Time)`

SetActivatedAt sets ActivatedAt field to given value.

### HasActivatedAt

`func (o *RoutingConfigVersionResponse) HasActivatedAt() bool`

HasActivatedAt returns a boolean if a field has been set.

### SetActivatedAtNil

`func (o *RoutingConfigVersionResponse) SetActivatedAtNil(b bool)`

 SetActivatedAtNil sets the value for ActivatedAt to be an explicit nil

### UnsetActivatedAt
`func (o *RoutingConfigVersionResponse) UnsetActivatedAt()`

UnsetActivatedAt ensures that no value is present for ActivatedAt, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



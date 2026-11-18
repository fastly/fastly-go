# ArcProviderConnectionListItemResponseExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletedAt** | Pointer to **time.Time** | The date and time the connection was deleted, in RFC 3339 format. Present only for deleted connections, which are returned only when `include_deleted` is `true`. | [optional] [readonly] 

## Methods

### NewArcProviderConnectionListItemResponseExtra

`func NewArcProviderConnectionListItemResponseExtra() *ArcProviderConnectionListItemResponseExtra`

NewArcProviderConnectionListItemResponseExtra instantiates a new ArcProviderConnectionListItemResponseExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcProviderConnectionListItemResponseExtraWithDefaults

`func NewArcProviderConnectionListItemResponseExtraWithDefaults() *ArcProviderConnectionListItemResponseExtra`

NewArcProviderConnectionListItemResponseExtraWithDefaults instantiates a new ArcProviderConnectionListItemResponseExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletedAt

`func (o *ArcProviderConnectionListItemResponseExtra) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ArcProviderConnectionListItemResponseExtra) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ArcProviderConnectionListItemResponseExtra) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ArcProviderConnectionListItemResponseExtra) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

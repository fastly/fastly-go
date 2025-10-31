# StarData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeStar**](TypeStar.md) |  | [optional] [default to TYPESTAR_STAR]
**Relationships** | Pointer to [**RelationshipsForStar**](RelationshipsForStar.md) |  | [optional] 

## Methods

### NewStarData

`func NewStarData() *StarData`

NewStarData instantiates a new StarData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStarDataWithDefaults

`func NewStarDataWithDefaults() *StarData`

NewStarDataWithDefaults instantiates a new StarData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StarData) GetType() TypeStar`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StarData) GetTypeOk() (*TypeStar, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StarData) SetType(v TypeStar)`

SetType sets Type field to given value.

### HasType

`func (o *StarData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRelationships

`func (o *StarData) GetRelationships() RelationshipsForStar`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *StarData) GetRelationshipsOk() (*RelationshipsForStar, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *StarData) SetRelationships(v RelationshipsForStar)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *StarData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



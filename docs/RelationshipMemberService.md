# RelationshipMemberService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeService**](TypeService.md) |  | [optional] [default to TYPESERVICE_SERVICE]
**ID** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewRelationshipMemberService

`func NewRelationshipMemberService() *RelationshipMemberService`

NewRelationshipMemberService instantiates a new RelationshipMemberService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipMemberServiceWithDefaults

`func NewRelationshipMemberServiceWithDefaults() *RelationshipMemberService`

NewRelationshipMemberServiceWithDefaults instantiates a new RelationshipMemberService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RelationshipMemberService) GetType() TypeService`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationshipMemberService) GetTypeOk() (*TypeService, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationshipMemberService) SetType(v TypeService)`

SetType sets Type field to given value.

### HasType

`func (o *RelationshipMemberService) HasType() bool`

HasType returns a boolean if a field has been set.

### GetID

`func (o *RelationshipMemberService) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *RelationshipMemberService) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *RelationshipMemberService) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *RelationshipMemberService) HasID() bool`

HasID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

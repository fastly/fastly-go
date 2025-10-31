# RelationshipService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to [**RelationshipMemberService**](RelationshipMemberService.md) |  | [optional] 

## Methods

### NewRelationshipService

`func NewRelationshipService() *RelationshipService`

NewRelationshipService instantiates a new RelationshipService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipServiceWithDefaults

`func NewRelationshipServiceWithDefaults() *RelationshipService`

NewRelationshipServiceWithDefaults instantiates a new RelationshipService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *RelationshipService) GetService() RelationshipMemberService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *RelationshipService) GetServiceOk() (*RelationshipMemberService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *RelationshipService) SetService(v RelationshipMemberService)`

SetService sets Service field to given value.

### HasService

`func (o *RelationshipService) HasService() bool`

HasService returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



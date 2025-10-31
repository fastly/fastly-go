# ServiceAuthorizationDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to [**RelationshipMemberService**](RelationshipMemberService.md) |  | [optional] 
**User** | Pointer to [**ServiceAuthorizationDataRelationshipsUser**](ServiceAuthorizationDataRelationshipsUser.md) |  | [optional] 

## Methods

### NewServiceAuthorizationDataRelationships

`func NewServiceAuthorizationDataRelationships() *ServiceAuthorizationDataRelationships`

NewServiceAuthorizationDataRelationships instantiates a new ServiceAuthorizationDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAuthorizationDataRelationshipsWithDefaults

`func NewServiceAuthorizationDataRelationshipsWithDefaults() *ServiceAuthorizationDataRelationships`

NewServiceAuthorizationDataRelationshipsWithDefaults instantiates a new ServiceAuthorizationDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *ServiceAuthorizationDataRelationships) GetService() RelationshipMemberService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ServiceAuthorizationDataRelationships) GetServiceOk() (*RelationshipMemberService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ServiceAuthorizationDataRelationships) SetService(v RelationshipMemberService)`

SetService sets Service field to given value.

### HasService

`func (o *ServiceAuthorizationDataRelationships) HasService() bool`

HasService returns a boolean if a field has been set.

### GetUser

`func (o *ServiceAuthorizationDataRelationships) GetUser() ServiceAuthorizationDataRelationshipsUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ServiceAuthorizationDataRelationships) GetUserOk() (*ServiceAuthorizationDataRelationshipsUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ServiceAuthorizationDataRelationships) SetUser(v ServiceAuthorizationDataRelationshipsUser)`

SetUser sets User field to given value.

### HasUser

`func (o *ServiceAuthorizationDataRelationships) HasUser() bool`

HasUser returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



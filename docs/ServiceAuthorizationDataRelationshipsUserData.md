# ServiceAuthorizationDataRelationshipsUserData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeUser**](TypeUser.md) |  | [optional] [default to TYPEUSER_USER]
**ID** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewServiceAuthorizationDataRelationshipsUserData

`func NewServiceAuthorizationDataRelationshipsUserData() *ServiceAuthorizationDataRelationshipsUserData`

NewServiceAuthorizationDataRelationshipsUserData instantiates a new ServiceAuthorizationDataRelationshipsUserData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAuthorizationDataRelationshipsUserDataWithDefaults

`func NewServiceAuthorizationDataRelationshipsUserDataWithDefaults() *ServiceAuthorizationDataRelationshipsUserData`

NewServiceAuthorizationDataRelationshipsUserDataWithDefaults instantiates a new ServiceAuthorizationDataRelationshipsUserData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ServiceAuthorizationDataRelationshipsUserData) GetType() TypeUser`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceAuthorizationDataRelationshipsUserData) GetTypeOk() (*TypeUser, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceAuthorizationDataRelationshipsUserData) SetType(v TypeUser)`

SetType sets Type field to given value.

### HasType

`func (o *ServiceAuthorizationDataRelationshipsUserData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetID

`func (o *ServiceAuthorizationDataRelationshipsUserData) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *ServiceAuthorizationDataRelationshipsUserData) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *ServiceAuthorizationDataRelationshipsUserData) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *ServiceAuthorizationDataRelationshipsUserData) HasID() bool`

HasID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

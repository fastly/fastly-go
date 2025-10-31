# RelationshipMemberMutualAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeMutualAuthentication**](TypeMutualAuthentication.md) |  | [optional] [default to TYPEMUTUALAUTHENTICATION_MUTUAL_AUTHENTICATION]
**Id** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewRelationshipMemberMutualAuthentication

`func NewRelationshipMemberMutualAuthentication() *RelationshipMemberMutualAuthentication`

NewRelationshipMemberMutualAuthentication instantiates a new RelationshipMemberMutualAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipMemberMutualAuthenticationWithDefaults

`func NewRelationshipMemberMutualAuthenticationWithDefaults() *RelationshipMemberMutualAuthentication`

NewRelationshipMemberMutualAuthenticationWithDefaults instantiates a new RelationshipMemberMutualAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RelationshipMemberMutualAuthentication) GetType() TypeMutualAuthentication`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationshipMemberMutualAuthentication) GetTypeOk() (*TypeMutualAuthentication, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationshipMemberMutualAuthentication) SetType(v TypeMutualAuthentication)`

SetType sets Type field to given value.

### HasType

`func (o *RelationshipMemberMutualAuthentication) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *RelationshipMemberMutualAuthentication) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelationshipMemberMutualAuthentication) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelationshipMemberMutualAuthentication) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RelationshipMemberMutualAuthentication) HasId() bool`

HasId returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



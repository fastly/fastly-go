# MutualAuthenticationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeMutualAuthentication**](TypeMutualAuthentication.md) |  | [optional] [default to TYPEMUTUALAUTHENTICATION_MUTUAL_AUTHENTICATION]
**Attributes** | Pointer to [**MutualAuthenticationDataAttributes**](MutualAuthenticationDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipsForMutualAuthentication**](RelationshipsForMutualAuthentication.md) |  | [optional] 

## Methods

### NewMutualAuthenticationData

`func NewMutualAuthenticationData() *MutualAuthenticationData`

NewMutualAuthenticationData instantiates a new MutualAuthenticationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutualAuthenticationDataWithDefaults

`func NewMutualAuthenticationDataWithDefaults() *MutualAuthenticationData`

NewMutualAuthenticationDataWithDefaults instantiates a new MutualAuthenticationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MutualAuthenticationData) GetType() TypeMutualAuthentication`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MutualAuthenticationData) GetTypeOk() (*TypeMutualAuthentication, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MutualAuthenticationData) SetType(v TypeMutualAuthentication)`

SetType sets Type field to given value.

### HasType

`func (o *MutualAuthenticationData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *MutualAuthenticationData) GetAttributes() MutualAuthenticationDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MutualAuthenticationData) GetAttributesOk() (*MutualAuthenticationDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MutualAuthenticationData) SetAttributes(v MutualAuthenticationDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *MutualAuthenticationData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *MutualAuthenticationData) GetRelationships() RelationshipsForMutualAuthentication`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *MutualAuthenticationData) GetRelationshipsOk() (*RelationshipsForMutualAuthentication, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *MutualAuthenticationData) SetRelationships(v RelationshipsForMutualAuthentication)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *MutualAuthenticationData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

# MutualAuthenticationResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeMutualAuthentication**](TypeMutualAuthentication.md) |  | [optional] [default to TYPEMUTUALAUTHENTICATION_MUTUAL_AUTHENTICATION]
**Attributes** | Pointer to [**MutualAuthenticationResponseAttributes**](MutualAuthenticationResponseAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipsForMutualAuthentication**](RelationshipsForMutualAuthentication.md) |  | [optional] 
**ID** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewMutualAuthenticationResponseData

`func NewMutualAuthenticationResponseData() *MutualAuthenticationResponseData`

NewMutualAuthenticationResponseData instantiates a new MutualAuthenticationResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutualAuthenticationResponseDataWithDefaults

`func NewMutualAuthenticationResponseDataWithDefaults() *MutualAuthenticationResponseData`

NewMutualAuthenticationResponseDataWithDefaults instantiates a new MutualAuthenticationResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MutualAuthenticationResponseData) GetType() TypeMutualAuthentication`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MutualAuthenticationResponseData) GetTypeOk() (*TypeMutualAuthentication, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MutualAuthenticationResponseData) SetType(v TypeMutualAuthentication)`

SetType sets Type field to given value.

### HasType

`func (o *MutualAuthenticationResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *MutualAuthenticationResponseData) GetAttributes() MutualAuthenticationResponseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MutualAuthenticationResponseData) GetAttributesOk() (*MutualAuthenticationResponseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MutualAuthenticationResponseData) SetAttributes(v MutualAuthenticationResponseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *MutualAuthenticationResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *MutualAuthenticationResponseData) GetRelationships() RelationshipsForMutualAuthentication`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *MutualAuthenticationResponseData) GetRelationshipsOk() (*RelationshipsForMutualAuthentication, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *MutualAuthenticationResponseData) SetRelationships(v RelationshipsForMutualAuthentication)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *MutualAuthenticationResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetID

`func (o *MutualAuthenticationResponseData) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *MutualAuthenticationResponseData) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *MutualAuthenticationResponseData) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *MutualAuthenticationResponseData) HasID() bool`

HasID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

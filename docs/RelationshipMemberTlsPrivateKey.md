# RelationshipMemberTLSPrivateKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTLSPrivateKey**](TypeTLSPrivateKey.md) |  | [optional] [default to TYPETLSPRIVATEKEY_TLS_PRIVATE_KEY]
**ID** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewRelationshipMemberTLSPrivateKey

`func NewRelationshipMemberTLSPrivateKey() *RelationshipMemberTLSPrivateKey`

NewRelationshipMemberTLSPrivateKey instantiates a new RelationshipMemberTLSPrivateKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipMemberTLSPrivateKeyWithDefaults

`func NewRelationshipMemberTLSPrivateKeyWithDefaults() *RelationshipMemberTLSPrivateKey`

NewRelationshipMemberTLSPrivateKeyWithDefaults instantiates a new RelationshipMemberTLSPrivateKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RelationshipMemberTLSPrivateKey) GetType() TypeTLSPrivateKey`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationshipMemberTLSPrivateKey) GetTypeOk() (*TypeTLSPrivateKey, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationshipMemberTLSPrivateKey) SetType(v TypeTLSPrivateKey)`

SetType sets Type field to given value.

### HasType

`func (o *RelationshipMemberTLSPrivateKey) HasType() bool`

HasType returns a boolean if a field has been set.

### GetID

`func (o *RelationshipMemberTLSPrivateKey) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *RelationshipMemberTLSPrivateKey) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *RelationshipMemberTLSPrivateKey) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *RelationshipMemberTLSPrivateKey) HasID() bool`

HasID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

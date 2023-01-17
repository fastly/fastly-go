# RelationshipMemberTLSDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTLSDomain**](TypeTLSDomain.md) |  | [optional] [default to TYPETLSDOMAIN_TLS_DOMAIN]
**ID** | Pointer to **string** | The domain name. | [optional] [readonly] 

## Methods

### NewRelationshipMemberTLSDomain

`func NewRelationshipMemberTLSDomain() *RelationshipMemberTLSDomain`

NewRelationshipMemberTLSDomain instantiates a new RelationshipMemberTLSDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipMemberTLSDomainWithDefaults

`func NewRelationshipMemberTLSDomainWithDefaults() *RelationshipMemberTLSDomain`

NewRelationshipMemberTLSDomainWithDefaults instantiates a new RelationshipMemberTLSDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RelationshipMemberTLSDomain) GetType() TypeTLSDomain`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationshipMemberTLSDomain) GetTypeOk() (*TypeTLSDomain, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationshipMemberTLSDomain) SetType(v TypeTLSDomain)`

SetType sets Type field to given value.

### HasType

`func (o *RelationshipMemberTLSDomain) HasType() bool`

HasType returns a boolean if a field has been set.

### GetID

`func (o *RelationshipMemberTLSDomain) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *RelationshipMemberTLSDomain) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *RelationshipMemberTLSDomain) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *RelationshipMemberTLSDomain) HasID() bool`

HasID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

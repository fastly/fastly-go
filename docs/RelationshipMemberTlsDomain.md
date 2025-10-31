# RelationshipMemberTlsDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTlsDomain**](TypeTlsDomain.md) |  | [optional] [default to TYPETLSDOMAIN_TLS_DOMAIN]
**Id** | Pointer to **string** | The domain name. | [optional] [readonly] 

## Methods

### NewRelationshipMemberTlsDomain

`func NewRelationshipMemberTlsDomain() *RelationshipMemberTlsDomain`

NewRelationshipMemberTlsDomain instantiates a new RelationshipMemberTlsDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipMemberTlsDomainWithDefaults

`func NewRelationshipMemberTlsDomainWithDefaults() *RelationshipMemberTlsDomain`

NewRelationshipMemberTlsDomainWithDefaults instantiates a new RelationshipMemberTlsDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RelationshipMemberTlsDomain) GetType() TypeTlsDomain`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationshipMemberTlsDomain) GetTypeOk() (*TypeTlsDomain, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationshipMemberTlsDomain) SetType(v TypeTlsDomain)`

SetType sets Type field to given value.

### HasType

`func (o *RelationshipMemberTlsDomain) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *RelationshipMemberTlsDomain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelationshipMemberTlsDomain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelationshipMemberTlsDomain) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RelationshipMemberTlsDomain) HasId() bool`

HasId returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



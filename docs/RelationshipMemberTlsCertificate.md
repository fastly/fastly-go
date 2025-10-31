# RelationshipMemberTlsCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTlsCertificate**](TypeTlsCertificate.md) |  | [optional] [default to TYPETLSCERTIFICATE_TLS_CERTIFICATE]
**Id** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewRelationshipMemberTlsCertificate

`func NewRelationshipMemberTlsCertificate() *RelationshipMemberTlsCertificate`

NewRelationshipMemberTlsCertificate instantiates a new RelationshipMemberTlsCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipMemberTlsCertificateWithDefaults

`func NewRelationshipMemberTlsCertificateWithDefaults() *RelationshipMemberTlsCertificate`

NewRelationshipMemberTlsCertificateWithDefaults instantiates a new RelationshipMemberTlsCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RelationshipMemberTlsCertificate) GetType() TypeTlsCertificate`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationshipMemberTlsCertificate) GetTypeOk() (*TypeTlsCertificate, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationshipMemberTlsCertificate) SetType(v TypeTlsCertificate)`

SetType sets Type field to given value.

### HasType

`func (o *RelationshipMemberTlsCertificate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *RelationshipMemberTlsCertificate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelationshipMemberTlsCertificate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelationshipMemberTlsCertificate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RelationshipMemberTlsCertificate) HasId() bool`

HasId returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



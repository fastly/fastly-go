# TlsDomainData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The domain name. | [optional] [readonly] 
**Type** | Pointer to [**TypeTlsDomain**](TypeTlsDomain.md) |  | [optional] [default to TYPETLSDOMAIN_TLS_DOMAIN]
**Relationships** | Pointer to [**RelationshipsForTlsDomain**](RelationshipsForTlsDomain.md) |  | [optional] 

## Methods

### NewTlsDomainData

`func NewTlsDomainData() *TlsDomainData`

NewTlsDomainData instantiates a new TlsDomainData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsDomainDataWithDefaults

`func NewTlsDomainDataWithDefaults() *TlsDomainData`

NewTlsDomainDataWithDefaults instantiates a new TlsDomainData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TlsDomainData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TlsDomainData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TlsDomainData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TlsDomainData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *TlsDomainData) GetType() TypeTlsDomain`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TlsDomainData) GetTypeOk() (*TypeTlsDomain, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TlsDomainData) SetType(v TypeTlsDomain)`

SetType sets Type field to given value.

### HasType

`func (o *TlsDomainData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRelationships

`func (o *TlsDomainData) GetRelationships() RelationshipsForTlsDomain`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TlsDomainData) GetRelationshipsOk() (*RelationshipsForTlsDomain, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TlsDomainData) SetRelationships(v RelationshipsForTlsDomain)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TlsDomainData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



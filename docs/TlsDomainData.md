# TLSDomainData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** | The domain name. | [optional] [readonly] 
**Type** | Pointer to [**TypeTLSDomain**](TypeTLSDomain.md) |  | [optional] [default to TYPETLSDOMAIN_TLS_DOMAIN]
**Relationships** | Pointer to [**RelationshipsForTLSDomain**](RelationshipsForTLSDomain.md) |  | [optional] 

## Methods

### NewTLSDomainData

`func NewTLSDomainData() *TLSDomainData`

NewTLSDomainData instantiates a new TLSDomainData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSDomainDataWithDefaults

`func NewTLSDomainDataWithDefaults() *TLSDomainData`

NewTLSDomainDataWithDefaults instantiates a new TLSDomainData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *TLSDomainData) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *TLSDomainData) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *TLSDomainData) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *TLSDomainData) HasID() bool`

HasID returns a boolean if a field has been set.

### GetType

`func (o *TLSDomainData) GetType() TypeTLSDomain`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TLSDomainData) GetTypeOk() (*TypeTLSDomain, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TLSDomainData) SetType(v TypeTLSDomain)`

SetType sets Type field to given value.

### HasType

`func (o *TLSDomainData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRelationships

`func (o *TLSDomainData) GetRelationships() RelationshipsForTLSDomain`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TLSDomainData) GetRelationshipsOk() (*RelationshipsForTLSDomain, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TLSDomainData) SetRelationships(v RelationshipsForTLSDomain)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TLSDomainData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

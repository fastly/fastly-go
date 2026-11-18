# TLSConfigurationResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTLSConfiguration**](TypeTLSConfiguration.md) |  | [optional] [default to TYPETLSCONFIGURATION_TLS_CONFIGURATION]
**Attributes** | Pointer to [**TLSConfigurationResponseAttributes**](TlsConfigurationResponseAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipsForTLSConfiguration**](RelationshipsForTLSConfiguration.md) |  | [optional] 
**ID** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewTLSConfigurationResponseData

`func NewTLSConfigurationResponseData() *TLSConfigurationResponseData`

NewTLSConfigurationResponseData instantiates a new TLSConfigurationResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSConfigurationResponseDataWithDefaults

`func NewTLSConfigurationResponseDataWithDefaults() *TLSConfigurationResponseData`

NewTLSConfigurationResponseDataWithDefaults instantiates a new TLSConfigurationResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TLSConfigurationResponseData) GetType() TypeTLSConfiguration`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TLSConfigurationResponseData) GetTypeOk() (*TypeTLSConfiguration, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TLSConfigurationResponseData) SetType(v TypeTLSConfiguration)`

SetType sets Type field to given value.

### HasType

`func (o *TLSConfigurationResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *TLSConfigurationResponseData) GetAttributes() TLSConfigurationResponseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TLSConfigurationResponseData) GetAttributesOk() (*TLSConfigurationResponseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TLSConfigurationResponseData) SetAttributes(v TLSConfigurationResponseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TLSConfigurationResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *TLSConfigurationResponseData) GetRelationships() RelationshipsForTLSConfiguration`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TLSConfigurationResponseData) GetRelationshipsOk() (*RelationshipsForTLSConfiguration, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TLSConfigurationResponseData) SetRelationships(v RelationshipsForTLSConfiguration)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TLSConfigurationResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetID

`func (o *TLSConfigurationResponseData) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *TLSConfigurationResponseData) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *TLSConfigurationResponseData) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *TLSConfigurationResponseData) HasID() bool`

HasID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

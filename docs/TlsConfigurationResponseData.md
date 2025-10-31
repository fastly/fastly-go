# TlsConfigurationResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTlsConfiguration**](TypeTlsConfiguration.md) |  | [optional] [default to TYPETLSCONFIGURATION_TLS_CONFIGURATION]
**Attributes** | Pointer to [**TlsConfigurationResponseAttributes**](TlsConfigurationResponseAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipsForTlsConfiguration**](RelationshipsForTlsConfiguration.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewTlsConfigurationResponseData

`func NewTlsConfigurationResponseData() *TlsConfigurationResponseData`

NewTlsConfigurationResponseData instantiates a new TlsConfigurationResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsConfigurationResponseDataWithDefaults

`func NewTlsConfigurationResponseDataWithDefaults() *TlsConfigurationResponseData`

NewTlsConfigurationResponseDataWithDefaults instantiates a new TlsConfigurationResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TlsConfigurationResponseData) GetType() TypeTlsConfiguration`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TlsConfigurationResponseData) GetTypeOk() (*TypeTlsConfiguration, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TlsConfigurationResponseData) SetType(v TypeTlsConfiguration)`

SetType sets Type field to given value.

### HasType

`func (o *TlsConfigurationResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *TlsConfigurationResponseData) GetAttributes() TlsConfigurationResponseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TlsConfigurationResponseData) GetAttributesOk() (*TlsConfigurationResponseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TlsConfigurationResponseData) SetAttributes(v TlsConfigurationResponseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TlsConfigurationResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *TlsConfigurationResponseData) GetRelationships() RelationshipsForTlsConfiguration`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TlsConfigurationResponseData) GetRelationshipsOk() (*RelationshipsForTlsConfiguration, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TlsConfigurationResponseData) SetRelationships(v RelationshipsForTlsConfiguration)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TlsConfigurationResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetId

`func (o *TlsConfigurationResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TlsConfigurationResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TlsConfigurationResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TlsConfigurationResponseData) HasId() bool`

HasId returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



# TlsConfigurationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTlsConfiguration**](TypeTlsConfiguration.md) |  | [optional] [default to TYPETLSCONFIGURATION_TLS_CONFIGURATION]
**Attributes** | Pointer to [**TlsConfigurationDataAttributes**](TlsConfigurationDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipsForTlsConfiguration**](RelationshipsForTlsConfiguration.md) |  | [optional] 

## Methods

### NewTlsConfigurationData

`func NewTlsConfigurationData() *TlsConfigurationData`

NewTlsConfigurationData instantiates a new TlsConfigurationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsConfigurationDataWithDefaults

`func NewTlsConfigurationDataWithDefaults() *TlsConfigurationData`

NewTlsConfigurationDataWithDefaults instantiates a new TlsConfigurationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TlsConfigurationData) GetType() TypeTlsConfiguration`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TlsConfigurationData) GetTypeOk() (*TypeTlsConfiguration, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TlsConfigurationData) SetType(v TypeTlsConfiguration)`

SetType sets Type field to given value.

### HasType

`func (o *TlsConfigurationData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *TlsConfigurationData) GetAttributes() TlsConfigurationDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TlsConfigurationData) GetAttributesOk() (*TlsConfigurationDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TlsConfigurationData) SetAttributes(v TlsConfigurationDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TlsConfigurationData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *TlsConfigurationData) GetRelationships() RelationshipsForTlsConfiguration`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TlsConfigurationData) GetRelationshipsOk() (*RelationshipsForTlsConfiguration, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TlsConfigurationData) SetRelationships(v RelationshipsForTlsConfiguration)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TlsConfigurationData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



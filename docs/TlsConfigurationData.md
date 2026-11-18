# TLSConfigurationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTLSConfiguration**](TypeTLSConfiguration.md) |  | [optional] [default to TYPETLSCONFIGURATION_TLS_CONFIGURATION]
**Attributes** | Pointer to [**TLSConfigurationDataAttributes**](TlsConfigurationDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipsForTLSConfiguration**](RelationshipsForTLSConfiguration.md) |  | [optional] 

## Methods

### NewTLSConfigurationData

`func NewTLSConfigurationData() *TLSConfigurationData`

NewTLSConfigurationData instantiates a new TLSConfigurationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSConfigurationDataWithDefaults

`func NewTLSConfigurationDataWithDefaults() *TLSConfigurationData`

NewTLSConfigurationDataWithDefaults instantiates a new TLSConfigurationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TLSConfigurationData) GetType() TypeTLSConfiguration`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TLSConfigurationData) GetTypeOk() (*TypeTLSConfiguration, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TLSConfigurationData) SetType(v TypeTLSConfiguration)`

SetType sets Type field to given value.

### HasType

`func (o *TLSConfigurationData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *TLSConfigurationData) GetAttributes() TLSConfigurationDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TLSConfigurationData) GetAttributesOk() (*TLSConfigurationDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TLSConfigurationData) SetAttributes(v TLSConfigurationDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TLSConfigurationData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *TLSConfigurationData) GetRelationships() RelationshipsForTLSConfiguration`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TLSConfigurationData) GetRelationshipsOk() (*RelationshipsForTLSConfiguration, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TLSConfigurationData) SetRelationships(v RelationshipsForTLSConfiguration)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TLSConfigurationData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

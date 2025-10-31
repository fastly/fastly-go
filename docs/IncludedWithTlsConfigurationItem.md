# IncludedWithTlsConfigurationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The IP address or hostname of the DNS record. | [optional] 
**Type** | Pointer to [**TypeTlsDnsRecord**](TypeTlsDnsRecord.md) |  | [optional] [default to TYPETLSDNSRECORD_DNS_RECORD]
**Attributes** | Pointer to [**TlsDnsRecord**](TlsDnsRecord.md) |  | [optional] 

## Methods

### NewIncludedWithTlsConfigurationItem

`func NewIncludedWithTlsConfigurationItem() *IncludedWithTlsConfigurationItem`

NewIncludedWithTlsConfigurationItem instantiates a new IncludedWithTlsConfigurationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncludedWithTlsConfigurationItemWithDefaults

`func NewIncludedWithTlsConfigurationItemWithDefaults() *IncludedWithTlsConfigurationItem`

NewIncludedWithTlsConfigurationItemWithDefaults instantiates a new IncludedWithTlsConfigurationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IncludedWithTlsConfigurationItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IncludedWithTlsConfigurationItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IncludedWithTlsConfigurationItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IncludedWithTlsConfigurationItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *IncludedWithTlsConfigurationItem) GetType() TypeTlsDnsRecord`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncludedWithTlsConfigurationItem) GetTypeOk() (*TypeTlsDnsRecord, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncludedWithTlsConfigurationItem) SetType(v TypeTlsDnsRecord)`

SetType sets Type field to given value.

### HasType

`func (o *IncludedWithTlsConfigurationItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *IncludedWithTlsConfigurationItem) GetAttributes() TlsDnsRecord`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IncludedWithTlsConfigurationItem) GetAttributesOk() (*TlsDnsRecord, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IncludedWithTlsConfigurationItem) SetAttributes(v TlsDnsRecord)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IncludedWithTlsConfigurationItem) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



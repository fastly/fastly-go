# IncludedWithTLSConfigurationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** | The IP address or hostname of the DNS record. | [optional] 
**Type** | Pointer to [**TypeTLSDNSRecord**](TypeTLSDNSRecord.md) |  | [optional] [default to TYPETLSDNSRECORD_DNS_RECORD]
**Attributes** | Pointer to [**TLSDNSRecord**](TlsDNSRecord.md) |  | [optional] 

## Methods

### NewIncludedWithTLSConfigurationItem

`func NewIncludedWithTLSConfigurationItem() *IncludedWithTLSConfigurationItem`

NewIncludedWithTLSConfigurationItem instantiates a new IncludedWithTLSConfigurationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncludedWithTLSConfigurationItemWithDefaults

`func NewIncludedWithTLSConfigurationItemWithDefaults() *IncludedWithTLSConfigurationItem`

NewIncludedWithTLSConfigurationItemWithDefaults instantiates a new IncludedWithTLSConfigurationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *IncludedWithTLSConfigurationItem) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *IncludedWithTLSConfigurationItem) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *IncludedWithTLSConfigurationItem) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *IncludedWithTLSConfigurationItem) HasID() bool`

HasID returns a boolean if a field has been set.

### GetType

`func (o *IncludedWithTLSConfigurationItem) GetType() TypeTLSDNSRecord`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncludedWithTLSConfigurationItem) GetTypeOk() (*TypeTLSDNSRecord, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncludedWithTLSConfigurationItem) SetType(v TypeTLSDNSRecord)`

SetType sets Type field to given value.

### HasType

`func (o *IncludedWithTLSConfigurationItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *IncludedWithTLSConfigurationItem) GetAttributes() TLSDNSRecord`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IncludedWithTLSConfigurationItem) GetAttributesOk() (*TLSDNSRecord, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IncludedWithTLSConfigurationItem) SetAttributes(v TLSDNSRecord)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IncludedWithTLSConfigurationItem) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

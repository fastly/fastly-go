# LegacyWafConfigurationSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | The active configuration set is the default configuration set when creating a new WAF. When Fastly adds configuration sets, the new versions become the default (active). | [optional] 
**Name** | Pointer to **string** | The name of the configuration set. | [optional] 

## Methods

### NewLegacyWafConfigurationSet

`func NewLegacyWafConfigurationSet() *LegacyWafConfigurationSet`

NewLegacyWafConfigurationSet instantiates a new LegacyWafConfigurationSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegacyWafConfigurationSetWithDefaults

`func NewLegacyWafConfigurationSetWithDefaults() *LegacyWafConfigurationSet`

NewLegacyWafConfigurationSetWithDefaults instantiates a new LegacyWafConfigurationSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *LegacyWafConfigurationSet) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *LegacyWafConfigurationSet) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *LegacyWafConfigurationSet) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *LegacyWafConfigurationSet) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetName

`func (o *LegacyWafConfigurationSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LegacyWafConfigurationSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LegacyWafConfigurationSet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LegacyWafConfigurationSet) HasName() bool`

HasName returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

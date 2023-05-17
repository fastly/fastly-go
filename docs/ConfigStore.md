# ConfigStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the config store. | [optional] 

## Methods

### NewConfigStore

`func NewConfigStore() *ConfigStore`

NewConfigStore instantiates a new ConfigStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigStoreWithDefaults

`func NewConfigStoreWithDefaults() *ConfigStore`

NewConfigStoreWithDefaults instantiates a new ConfigStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConfigStore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigStore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigStore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfigStore) HasName() bool`

HasName returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

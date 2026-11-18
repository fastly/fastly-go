# SecretStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A human-readable name for the store. The value must contain only letters, numbers, dashes (`-`), underscores (`_`), or periods (`.`). | [optional] 

## Methods

### NewSecretStore

`func NewSecretStore() *SecretStore`

NewSecretStore instantiates a new SecretStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretStoreWithDefaults

`func NewSecretStoreWithDefaults() *SecretStore`

NewSecretStoreWithDefaults instantiates a new SecretStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SecretStore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecretStore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecretStore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecretStore) HasName() bool`

HasName returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

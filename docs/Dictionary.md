# Dictionary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name for the Dictionary (must start with an alphabetic character and can contain only alphanumeric characters, underscores, and whitespace). | [optional] 
**WriteOnly** | Pointer to **bool** | Determines if items in the dictionary are readable or not. | [optional] [default to false]

## Methods

### NewDictionary

`func NewDictionary() *Dictionary`

NewDictionary instantiates a new Dictionary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDictionaryWithDefaults

`func NewDictionaryWithDefaults() *Dictionary`

NewDictionaryWithDefaults instantiates a new Dictionary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Dictionary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dictionary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dictionary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Dictionary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWriteOnly

`func (o *Dictionary) GetWriteOnly() bool`

GetWriteOnly returns the WriteOnly field if non-nil, zero value otherwise.

### GetWriteOnlyOk

`func (o *Dictionary) GetWriteOnlyOk() (*bool, bool)`

GetWriteOnlyOk returns a tuple with the WriteOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteOnly

`func (o *Dictionary) SetWriteOnly(v bool)`

SetWriteOnly sets WriteOnly field to given value.

### HasWriteOnly

`func (o *Dictionary) HasWriteOnly() bool`

HasWriteOnly returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



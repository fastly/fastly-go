# DefaultSettingsError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Detail** | Pointer to **string** |  | [optional] 

## Methods

### NewDefaultSettingsError

`func NewDefaultSettingsError() *DefaultSettingsError`

NewDefaultSettingsError instantiates a new DefaultSettingsError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultSettingsErrorWithDefaults

`func NewDefaultSettingsErrorWithDefaults() *DefaultSettingsError`

NewDefaultSettingsErrorWithDefaults instantiates a new DefaultSettingsError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *DefaultSettingsError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DefaultSettingsError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DefaultSettingsError) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DefaultSettingsError) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *DefaultSettingsError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DefaultSettingsError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DefaultSettingsError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DefaultSettingsError) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDetail

`func (o *DefaultSettingsError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *DefaultSettingsError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *DefaultSettingsError) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *DefaultSettingsError) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

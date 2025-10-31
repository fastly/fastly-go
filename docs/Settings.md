# Settings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneralDefaultHost** | Pointer to **string** | The default host name for the version. | [optional] 
**GeneralDefaultTtl** | Pointer to **int32** | The default time-to-live (TTL) for the version. | [optional] 
**GeneralStaleIfError** | Pointer to **bool** | Enables serving a stale object if there is an error. | [optional] [default to false]
**GeneralStaleIfErrorTtl** | Pointer to **int32** | The default time-to-live (TTL) for serving the stale object for the version. | [optional] [default to 43200]

## Methods

### NewSettings

`func NewSettings() *Settings`

NewSettings instantiates a new Settings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsWithDefaults

`func NewSettingsWithDefaults() *Settings`

NewSettingsWithDefaults instantiates a new Settings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneralDefaultHost

`func (o *Settings) GetGeneralDefaultHost() string`

GetGeneralDefaultHost returns the GeneralDefaultHost field if non-nil, zero value otherwise.

### GetGeneralDefaultHostOk

`func (o *Settings) GetGeneralDefaultHostOk() (*string, bool)`

GetGeneralDefaultHostOk returns a tuple with the GeneralDefaultHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralDefaultHost

`func (o *Settings) SetGeneralDefaultHost(v string)`

SetGeneralDefaultHost sets GeneralDefaultHost field to given value.

### HasGeneralDefaultHost

`func (o *Settings) HasGeneralDefaultHost() bool`

HasGeneralDefaultHost returns a boolean if a field has been set.

### GetGeneralDefaultTtl

`func (o *Settings) GetGeneralDefaultTtl() int32`

GetGeneralDefaultTtl returns the GeneralDefaultTtl field if non-nil, zero value otherwise.

### GetGeneralDefaultTtlOk

`func (o *Settings) GetGeneralDefaultTtlOk() (*int32, bool)`

GetGeneralDefaultTtlOk returns a tuple with the GeneralDefaultTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralDefaultTtl

`func (o *Settings) SetGeneralDefaultTtl(v int32)`

SetGeneralDefaultTtl sets GeneralDefaultTtl field to given value.

### HasGeneralDefaultTtl

`func (o *Settings) HasGeneralDefaultTtl() bool`

HasGeneralDefaultTtl returns a boolean if a field has been set.

### GetGeneralStaleIfError

`func (o *Settings) GetGeneralStaleIfError() bool`

GetGeneralStaleIfError returns the GeneralStaleIfError field if non-nil, zero value otherwise.

### GetGeneralStaleIfErrorOk

`func (o *Settings) GetGeneralStaleIfErrorOk() (*bool, bool)`

GetGeneralStaleIfErrorOk returns a tuple with the GeneralStaleIfError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralStaleIfError

`func (o *Settings) SetGeneralStaleIfError(v bool)`

SetGeneralStaleIfError sets GeneralStaleIfError field to given value.

### HasGeneralStaleIfError

`func (o *Settings) HasGeneralStaleIfError() bool`

HasGeneralStaleIfError returns a boolean if a field has been set.

### GetGeneralStaleIfErrorTtl

`func (o *Settings) GetGeneralStaleIfErrorTtl() int32`

GetGeneralStaleIfErrorTtl returns the GeneralStaleIfErrorTtl field if non-nil, zero value otherwise.

### GetGeneralStaleIfErrorTtlOk

`func (o *Settings) GetGeneralStaleIfErrorTtlOk() (*int32, bool)`

GetGeneralStaleIfErrorTtlOk returns a tuple with the GeneralStaleIfErrorTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralStaleIfErrorTtl

`func (o *Settings) SetGeneralStaleIfErrorTtl(v int32)`

SetGeneralStaleIfErrorTtl sets GeneralStaleIfErrorTtl field to given value.

### HasGeneralStaleIfErrorTtl

`func (o *Settings) HasGeneralStaleIfErrorTtl() bool`

HasGeneralStaleIfErrorTtl returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



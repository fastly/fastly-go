# Settings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneralDefaultHost** | Pointer to **string** | The default host name for the version. | [optional] 
**GeneralDefaultTTL** | Pointer to **int32** | The default time-to-live (TTL) for the version. | [optional] 
**GeneralStaleIfError** | Pointer to **bool** | Enables serving a stale object if there is an error. | [optional] [default to false]
**GeneralStaleIfErrorTTL** | Pointer to **int32** | The default time-to-live (TTL) for serving the stale object for the version. | [optional] [default to 43200]

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

### GetGeneralDefaultTTL

`func (o *Settings) GetGeneralDefaultTTL() int32`

GetGeneralDefaultTTL returns the GeneralDefaultTTL field if non-nil, zero value otherwise.

### GetGeneralDefaultTTLOk

`func (o *Settings) GetGeneralDefaultTTLOk() (*int32, bool)`

GetGeneralDefaultTTLOk returns a tuple with the GeneralDefaultTTL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralDefaultTTL

`func (o *Settings) SetGeneralDefaultTTL(v int32)`

SetGeneralDefaultTTL sets GeneralDefaultTTL field to given value.

### HasGeneralDefaultTTL

`func (o *Settings) HasGeneralDefaultTTL() bool`

HasGeneralDefaultTTL returns a boolean if a field has been set.

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

### GetGeneralStaleIfErrorTTL

`func (o *Settings) GetGeneralStaleIfErrorTTL() int32`

GetGeneralStaleIfErrorTTL returns the GeneralStaleIfErrorTTL field if non-nil, zero value otherwise.

### GetGeneralStaleIfErrorTTLOk

`func (o *Settings) GetGeneralStaleIfErrorTTLOk() (*int32, bool)`

GetGeneralStaleIfErrorTTLOk returns a tuple with the GeneralStaleIfErrorTTL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralStaleIfErrorTTL

`func (o *Settings) SetGeneralStaleIfErrorTTL(v int32)`

SetGeneralStaleIfErrorTTL sets GeneralStaleIfErrorTTL field to given value.

### HasGeneralStaleIfErrorTTL

`func (o *Settings) HasGeneralStaleIfErrorTTL() bool`

HasGeneralStaleIfErrorTTL returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

# SettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneralDefaultHost** | Pointer to **string** | The default host name for the version. | [optional] 
**GeneralDefaultTtl** | Pointer to **int32** | The default time-to-live (TTL) for the version. | [optional] 
**GeneralStaleIfError** | Pointer to **bool** | Enables serving a stale object if there is an error. | [optional] [default to false]
**GeneralStaleIfErrorTtl** | Pointer to **int32** | The default time-to-live (TTL) for serving the stale object for the version. | [optional] [default to 43200]
**ServiceId** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewSettingsResponse

`func NewSettingsResponse() *SettingsResponse`

NewSettingsResponse instantiates a new SettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsResponseWithDefaults

`func NewSettingsResponseWithDefaults() *SettingsResponse`

NewSettingsResponseWithDefaults instantiates a new SettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneralDefaultHost

`func (o *SettingsResponse) GetGeneralDefaultHost() string`

GetGeneralDefaultHost returns the GeneralDefaultHost field if non-nil, zero value otherwise.

### GetGeneralDefaultHostOk

`func (o *SettingsResponse) GetGeneralDefaultHostOk() (*string, bool)`

GetGeneralDefaultHostOk returns a tuple with the GeneralDefaultHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralDefaultHost

`func (o *SettingsResponse) SetGeneralDefaultHost(v string)`

SetGeneralDefaultHost sets GeneralDefaultHost field to given value.

### HasGeneralDefaultHost

`func (o *SettingsResponse) HasGeneralDefaultHost() bool`

HasGeneralDefaultHost returns a boolean if a field has been set.

### GetGeneralDefaultTtl

`func (o *SettingsResponse) GetGeneralDefaultTtl() int32`

GetGeneralDefaultTtl returns the GeneralDefaultTtl field if non-nil, zero value otherwise.

### GetGeneralDefaultTtlOk

`func (o *SettingsResponse) GetGeneralDefaultTtlOk() (*int32, bool)`

GetGeneralDefaultTtlOk returns a tuple with the GeneralDefaultTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralDefaultTtl

`func (o *SettingsResponse) SetGeneralDefaultTtl(v int32)`

SetGeneralDefaultTtl sets GeneralDefaultTtl field to given value.

### HasGeneralDefaultTtl

`func (o *SettingsResponse) HasGeneralDefaultTtl() bool`

HasGeneralDefaultTtl returns a boolean if a field has been set.

### GetGeneralStaleIfError

`func (o *SettingsResponse) GetGeneralStaleIfError() bool`

GetGeneralStaleIfError returns the GeneralStaleIfError field if non-nil, zero value otherwise.

### GetGeneralStaleIfErrorOk

`func (o *SettingsResponse) GetGeneralStaleIfErrorOk() (*bool, bool)`

GetGeneralStaleIfErrorOk returns a tuple with the GeneralStaleIfError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralStaleIfError

`func (o *SettingsResponse) SetGeneralStaleIfError(v bool)`

SetGeneralStaleIfError sets GeneralStaleIfError field to given value.

### HasGeneralStaleIfError

`func (o *SettingsResponse) HasGeneralStaleIfError() bool`

HasGeneralStaleIfError returns a boolean if a field has been set.

### GetGeneralStaleIfErrorTtl

`func (o *SettingsResponse) GetGeneralStaleIfErrorTtl() int32`

GetGeneralStaleIfErrorTtl returns the GeneralStaleIfErrorTtl field if non-nil, zero value otherwise.

### GetGeneralStaleIfErrorTtlOk

`func (o *SettingsResponse) GetGeneralStaleIfErrorTtlOk() (*int32, bool)`

GetGeneralStaleIfErrorTtlOk returns a tuple with the GeneralStaleIfErrorTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralStaleIfErrorTtl

`func (o *SettingsResponse) SetGeneralStaleIfErrorTtl(v int32)`

SetGeneralStaleIfErrorTtl sets GeneralStaleIfErrorTtl field to given value.

### HasGeneralStaleIfErrorTtl

`func (o *SettingsResponse) HasGeneralStaleIfErrorTtl() bool`

HasGeneralStaleIfErrorTtl returns a boolean if a field has been set.

### GetServiceId

`func (o *SettingsResponse) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *SettingsResponse) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *SettingsResponse) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *SettingsResponse) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetVersion

`func (o *SettingsResponse) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SettingsResponse) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SettingsResponse) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SettingsResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



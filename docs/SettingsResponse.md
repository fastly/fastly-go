# SettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneralDefaultHost** | Pointer to **string** | The default host name for the version. | [optional] 
**GeneralDefaultTTL** | Pointer to **int32** | The default time-to-live (TTL) for the version. | [optional] 
**GeneralStaleIfError** | Pointer to **bool** | Enables serving a stale object if there is an error. | [optional] [default to false]
**GeneralStaleIfErrorTTL** | Pointer to **int32** | The default time-to-live (TTL) for serving the stale object for the version. | [optional] [default to 43200]
**ServiceID** | Pointer to **string** |  | [optional] [readonly] 
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

### GetGeneralDefaultTTL

`func (o *SettingsResponse) GetGeneralDefaultTTL() int32`

GetGeneralDefaultTTL returns the GeneralDefaultTTL field if non-nil, zero value otherwise.

### GetGeneralDefaultTTLOk

`func (o *SettingsResponse) GetGeneralDefaultTTLOk() (*int32, bool)`

GetGeneralDefaultTTLOk returns a tuple with the GeneralDefaultTTL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralDefaultTTL

`func (o *SettingsResponse) SetGeneralDefaultTTL(v int32)`

SetGeneralDefaultTTL sets GeneralDefaultTTL field to given value.

### HasGeneralDefaultTTL

`func (o *SettingsResponse) HasGeneralDefaultTTL() bool`

HasGeneralDefaultTTL returns a boolean if a field has been set.

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

### GetGeneralStaleIfErrorTTL

`func (o *SettingsResponse) GetGeneralStaleIfErrorTTL() int32`

GetGeneralStaleIfErrorTTL returns the GeneralStaleIfErrorTTL field if non-nil, zero value otherwise.

### GetGeneralStaleIfErrorTTLOk

`func (o *SettingsResponse) GetGeneralStaleIfErrorTTLOk() (*int32, bool)`

GetGeneralStaleIfErrorTTLOk returns a tuple with the GeneralStaleIfErrorTTL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralStaleIfErrorTTL

`func (o *SettingsResponse) SetGeneralStaleIfErrorTTL(v int32)`

SetGeneralStaleIfErrorTTL sets GeneralStaleIfErrorTTL field to given value.

### HasGeneralStaleIfErrorTTL

`func (o *SettingsResponse) HasGeneralStaleIfErrorTTL() bool`

HasGeneralStaleIfErrorTTL returns a boolean if a field has been set.

### GetServiceID

`func (o *SettingsResponse) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *SettingsResponse) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *SettingsResponse) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *SettingsResponse) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

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

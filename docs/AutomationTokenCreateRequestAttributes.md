# AutomationTokenCreateRequestAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | name of the token | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Services** | Pointer to **[]string** | List of service ids to limit the token | [optional] 
**Scope** | Pointer to **string** |  | [optional] [default to "global"]
**ExpiresAt** | Pointer to **NullableTime** | A UTC timestamp of when the token will expire. | [optional] 
**TlsAccess** | Pointer to **bool** | Indicates whether TLS access is enabled for the token. | [optional] 

## Methods

### NewAutomationTokenCreateRequestAttributes

`func NewAutomationTokenCreateRequestAttributes() *AutomationTokenCreateRequestAttributes`

NewAutomationTokenCreateRequestAttributes instantiates a new AutomationTokenCreateRequestAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationTokenCreateRequestAttributesWithDefaults

`func NewAutomationTokenCreateRequestAttributesWithDefaults() *AutomationTokenCreateRequestAttributes`

NewAutomationTokenCreateRequestAttributesWithDefaults instantiates a new AutomationTokenCreateRequestAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AutomationTokenCreateRequestAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutomationTokenCreateRequestAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutomationTokenCreateRequestAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutomationTokenCreateRequestAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRole

`func (o *AutomationTokenCreateRequestAttributes) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AutomationTokenCreateRequestAttributes) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AutomationTokenCreateRequestAttributes) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AutomationTokenCreateRequestAttributes) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetServices

`func (o *AutomationTokenCreateRequestAttributes) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *AutomationTokenCreateRequestAttributes) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *AutomationTokenCreateRequestAttributes) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *AutomationTokenCreateRequestAttributes) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetScope

`func (o *AutomationTokenCreateRequestAttributes) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AutomationTokenCreateRequestAttributes) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AutomationTokenCreateRequestAttributes) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AutomationTokenCreateRequestAttributes) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetExpiresAt

`func (o *AutomationTokenCreateRequestAttributes) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AutomationTokenCreateRequestAttributes) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AutomationTokenCreateRequestAttributes) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *AutomationTokenCreateRequestAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *AutomationTokenCreateRequestAttributes) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *AutomationTokenCreateRequestAttributes) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetTlsAccess

`func (o *AutomationTokenCreateRequestAttributes) GetTlsAccess() bool`

GetTlsAccess returns the TlsAccess field if non-nil, zero value otherwise.

### GetTlsAccessOk

`func (o *AutomationTokenCreateRequestAttributes) GetTlsAccessOk() (*bool, bool)`

GetTlsAccessOk returns a tuple with the TlsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsAccess

`func (o *AutomationTokenCreateRequestAttributes) SetTlsAccess(v bool)`

SetTlsAccess sets TlsAccess field to given value.

### HasTlsAccess

`func (o *AutomationTokenCreateRequestAttributes) HasTlsAccess() bool`

HasTlsAccess returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



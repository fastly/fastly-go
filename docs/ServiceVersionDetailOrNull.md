# ServiceVersionDetailOrNull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Whether this is the active version or not. | [optional] [default to false]
**Comment** | Pointer to **NullableString** | A freeform descriptive note. | [optional] 
**Deployed** | Pointer to **bool** | Unused at this time. | [optional] 
**Locked** | Pointer to **bool** | Whether this version is locked or not. Objects can not be added or edited on locked versions. | [optional] [default to false]
**Number** | Pointer to **int32** | The number of this version. | [optional] [readonly] 
**Staging** | Pointer to **bool** | Unused at this time. | [optional] [default to false]
**Testing** | Pointer to **bool** | Unused at this time. | [optional] [default to false]
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**ServiceID** | Pointer to **string** |  | [optional] [readonly] 
**Backends** | Pointer to [**[]BackendResponse**](BackendResponse.md) | List of backends associated to this service. | [optional] 
**CacheSettings** | Pointer to [**[]CacheSettingResponse**](CacheSettingResponse.md) | List of cache settings associated to this service. | [optional] 
**Conditions** | Pointer to [**[]ConditionResponse**](ConditionResponse.md) | List of conditions associated to this service. | [optional] 
**Directors** | Pointer to [**[]Director**](Director.md) | List of directors associated to this service. | [optional] 
**Domains** | Pointer to [**[]DomainResponse**](DomainResponse.md) | List of domains associated to this service. | [optional] 
**Gzips** | Pointer to [**[]GzipResponse**](GzipResponse.md) | List of gzip rules associated to this service. | [optional] 
**Headers** | Pointer to [**[]HeaderResponse**](HeaderResponse.md) | List of headers associated to this service. | [optional] 
**Healthchecks** | Pointer to [**[]HealthcheckResponse**](HealthcheckResponse.md) | List of healthchecks associated to this service. | [optional] 
**RequestSettings** | Pointer to [**[]RequestSettingsResponse**](RequestSettingsResponse.md) | List of request settings for this service. | [optional] 
**ResponseObjects** | Pointer to [**[]ResponseObjectResponse**](ResponseObjectResponse.md) | List of response objects for this service. | [optional] 
**Settings** | Pointer to [**VersionDetailSettings**](VersionDetailSettings.md) |  | [optional] 
**Snippets** | Pointer to [**[]SchemasSnippetResponse**](SchemasSnippetResponse.md) | List of VCL snippets for this service. | [optional] 
**Vcls** | Pointer to [**[]SchemasVclResponse**](SchemasVclResponse.md) | List of VCL files for this service. | [optional] 
**Wordpress** | Pointer to **[]map[string]any** | A list of Wordpress rules with this service. | [optional] 

## Methods

### NewServiceVersionDetailOrNull

`func NewServiceVersionDetailOrNull() *ServiceVersionDetailOrNull`

NewServiceVersionDetailOrNull instantiates a new ServiceVersionDetailOrNull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceVersionDetailOrNullWithDefaults

`func NewServiceVersionDetailOrNullWithDefaults() *ServiceVersionDetailOrNull`

NewServiceVersionDetailOrNullWithDefaults instantiates a new ServiceVersionDetailOrNull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ServiceVersionDetailOrNull) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ServiceVersionDetailOrNull) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ServiceVersionDetailOrNull) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ServiceVersionDetailOrNull) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetComment

`func (o *ServiceVersionDetailOrNull) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ServiceVersionDetailOrNull) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ServiceVersionDetailOrNull) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ServiceVersionDetailOrNull) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *ServiceVersionDetailOrNull) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *ServiceVersionDetailOrNull) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetDeployed

`func (o *ServiceVersionDetailOrNull) GetDeployed() bool`

GetDeployed returns the Deployed field if non-nil, zero value otherwise.

### GetDeployedOk

`func (o *ServiceVersionDetailOrNull) GetDeployedOk() (*bool, bool)`

GetDeployedOk returns a tuple with the Deployed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployed

`func (o *ServiceVersionDetailOrNull) SetDeployed(v bool)`

SetDeployed sets Deployed field to given value.

### HasDeployed

`func (o *ServiceVersionDetailOrNull) HasDeployed() bool`

HasDeployed returns a boolean if a field has been set.

### GetLocked

`func (o *ServiceVersionDetailOrNull) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *ServiceVersionDetailOrNull) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *ServiceVersionDetailOrNull) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *ServiceVersionDetailOrNull) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetNumber

`func (o *ServiceVersionDetailOrNull) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ServiceVersionDetailOrNull) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ServiceVersionDetailOrNull) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *ServiceVersionDetailOrNull) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetStaging

`func (o *ServiceVersionDetailOrNull) GetStaging() bool`

GetStaging returns the Staging field if non-nil, zero value otherwise.

### GetStagingOk

`func (o *ServiceVersionDetailOrNull) GetStagingOk() (*bool, bool)`

GetStagingOk returns a tuple with the Staging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaging

`func (o *ServiceVersionDetailOrNull) SetStaging(v bool)`

SetStaging sets Staging field to given value.

### HasStaging

`func (o *ServiceVersionDetailOrNull) HasStaging() bool`

HasStaging returns a boolean if a field has been set.

### GetTesting

`func (o *ServiceVersionDetailOrNull) GetTesting() bool`

GetTesting returns the Testing field if non-nil, zero value otherwise.

### GetTestingOk

`func (o *ServiceVersionDetailOrNull) GetTestingOk() (*bool, bool)`

GetTestingOk returns a tuple with the Testing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesting

`func (o *ServiceVersionDetailOrNull) SetTesting(v bool)`

SetTesting sets Testing field to given value.

### HasTesting

`func (o *ServiceVersionDetailOrNull) HasTesting() bool`

HasTesting returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ServiceVersionDetailOrNull) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceVersionDetailOrNull) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceVersionDetailOrNull) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServiceVersionDetailOrNull) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ServiceVersionDetailOrNull) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ServiceVersionDetailOrNull) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *ServiceVersionDetailOrNull) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ServiceVersionDetailOrNull) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ServiceVersionDetailOrNull) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ServiceVersionDetailOrNull) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *ServiceVersionDetailOrNull) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *ServiceVersionDetailOrNull) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *ServiceVersionDetailOrNull) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServiceVersionDetailOrNull) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServiceVersionDetailOrNull) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ServiceVersionDetailOrNull) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *ServiceVersionDetailOrNull) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ServiceVersionDetailOrNull) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetServiceID

`func (o *ServiceVersionDetailOrNull) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *ServiceVersionDetailOrNull) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *ServiceVersionDetailOrNull) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *ServiceVersionDetailOrNull) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetBackends

`func (o *ServiceVersionDetailOrNull) GetBackends() []BackendResponse`

GetBackends returns the Backends field if non-nil, zero value otherwise.

### GetBackendsOk

`func (o *ServiceVersionDetailOrNull) GetBackendsOk() (*[]BackendResponse, bool)`

GetBackendsOk returns a tuple with the Backends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackends

`func (o *ServiceVersionDetailOrNull) SetBackends(v []BackendResponse)`

SetBackends sets Backends field to given value.

### HasBackends

`func (o *ServiceVersionDetailOrNull) HasBackends() bool`

HasBackends returns a boolean if a field has been set.

### GetCacheSettings

`func (o *ServiceVersionDetailOrNull) GetCacheSettings() []CacheSettingResponse`

GetCacheSettings returns the CacheSettings field if non-nil, zero value otherwise.

### GetCacheSettingsOk

`func (o *ServiceVersionDetailOrNull) GetCacheSettingsOk() (*[]CacheSettingResponse, bool)`

GetCacheSettingsOk returns a tuple with the CacheSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheSettings

`func (o *ServiceVersionDetailOrNull) SetCacheSettings(v []CacheSettingResponse)`

SetCacheSettings sets CacheSettings field to given value.

### HasCacheSettings

`func (o *ServiceVersionDetailOrNull) HasCacheSettings() bool`

HasCacheSettings returns a boolean if a field has been set.

### GetConditions

`func (o *ServiceVersionDetailOrNull) GetConditions() []ConditionResponse`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ServiceVersionDetailOrNull) GetConditionsOk() (*[]ConditionResponse, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ServiceVersionDetailOrNull) SetConditions(v []ConditionResponse)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ServiceVersionDetailOrNull) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetDirectors

`func (o *ServiceVersionDetailOrNull) GetDirectors() []Director`

GetDirectors returns the Directors field if non-nil, zero value otherwise.

### GetDirectorsOk

`func (o *ServiceVersionDetailOrNull) GetDirectorsOk() (*[]Director, bool)`

GetDirectorsOk returns a tuple with the Directors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectors

`func (o *ServiceVersionDetailOrNull) SetDirectors(v []Director)`

SetDirectors sets Directors field to given value.

### HasDirectors

`func (o *ServiceVersionDetailOrNull) HasDirectors() bool`

HasDirectors returns a boolean if a field has been set.

### GetDomains

`func (o *ServiceVersionDetailOrNull) GetDomains() []DomainResponse`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *ServiceVersionDetailOrNull) GetDomainsOk() (*[]DomainResponse, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *ServiceVersionDetailOrNull) SetDomains(v []DomainResponse)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *ServiceVersionDetailOrNull) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetGzips

`func (o *ServiceVersionDetailOrNull) GetGzips() []GzipResponse`

GetGzips returns the Gzips field if non-nil, zero value otherwise.

### GetGzipsOk

`func (o *ServiceVersionDetailOrNull) GetGzipsOk() (*[]GzipResponse, bool)`

GetGzipsOk returns a tuple with the Gzips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGzips

`func (o *ServiceVersionDetailOrNull) SetGzips(v []GzipResponse)`

SetGzips sets Gzips field to given value.

### HasGzips

`func (o *ServiceVersionDetailOrNull) HasGzips() bool`

HasGzips returns a boolean if a field has been set.

### GetHeaders

`func (o *ServiceVersionDetailOrNull) GetHeaders() []HeaderResponse`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ServiceVersionDetailOrNull) GetHeadersOk() (*[]HeaderResponse, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ServiceVersionDetailOrNull) SetHeaders(v []HeaderResponse)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ServiceVersionDetailOrNull) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetHealthchecks

`func (o *ServiceVersionDetailOrNull) GetHealthchecks() []HealthcheckResponse`

GetHealthchecks returns the Healthchecks field if non-nil, zero value otherwise.

### GetHealthchecksOk

`func (o *ServiceVersionDetailOrNull) GetHealthchecksOk() (*[]HealthcheckResponse, bool)`

GetHealthchecksOk returns a tuple with the Healthchecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthchecks

`func (o *ServiceVersionDetailOrNull) SetHealthchecks(v []HealthcheckResponse)`

SetHealthchecks sets Healthchecks field to given value.

### HasHealthchecks

`func (o *ServiceVersionDetailOrNull) HasHealthchecks() bool`

HasHealthchecks returns a boolean if a field has been set.

### GetRequestSettings

`func (o *ServiceVersionDetailOrNull) GetRequestSettings() []RequestSettingsResponse`

GetRequestSettings returns the RequestSettings field if non-nil, zero value otherwise.

### GetRequestSettingsOk

`func (o *ServiceVersionDetailOrNull) GetRequestSettingsOk() (*[]RequestSettingsResponse, bool)`

GetRequestSettingsOk returns a tuple with the RequestSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSettings

`func (o *ServiceVersionDetailOrNull) SetRequestSettings(v []RequestSettingsResponse)`

SetRequestSettings sets RequestSettings field to given value.

### HasRequestSettings

`func (o *ServiceVersionDetailOrNull) HasRequestSettings() bool`

HasRequestSettings returns a boolean if a field has been set.

### GetResponseObjects

`func (o *ServiceVersionDetailOrNull) GetResponseObjects() []ResponseObjectResponse`

GetResponseObjects returns the ResponseObjects field if non-nil, zero value otherwise.

### GetResponseObjectsOk

`func (o *ServiceVersionDetailOrNull) GetResponseObjectsOk() (*[]ResponseObjectResponse, bool)`

GetResponseObjectsOk returns a tuple with the ResponseObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseObjects

`func (o *ServiceVersionDetailOrNull) SetResponseObjects(v []ResponseObjectResponse)`

SetResponseObjects sets ResponseObjects field to given value.

### HasResponseObjects

`func (o *ServiceVersionDetailOrNull) HasResponseObjects() bool`

HasResponseObjects returns a boolean if a field has been set.

### GetSettings

`func (o *ServiceVersionDetailOrNull) GetSettings() VersionDetailSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ServiceVersionDetailOrNull) GetSettingsOk() (*VersionDetailSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ServiceVersionDetailOrNull) SetSettings(v VersionDetailSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *ServiceVersionDetailOrNull) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetSnippets

`func (o *ServiceVersionDetailOrNull) GetSnippets() []SchemasSnippetResponse`

GetSnippets returns the Snippets field if non-nil, zero value otherwise.

### GetSnippetsOk

`func (o *ServiceVersionDetailOrNull) GetSnippetsOk() (*[]SchemasSnippetResponse, bool)`

GetSnippetsOk returns a tuple with the Snippets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippets

`func (o *ServiceVersionDetailOrNull) SetSnippets(v []SchemasSnippetResponse)`

SetSnippets sets Snippets field to given value.

### HasSnippets

`func (o *ServiceVersionDetailOrNull) HasSnippets() bool`

HasSnippets returns a boolean if a field has been set.

### GetVcls

`func (o *ServiceVersionDetailOrNull) GetVcls() []SchemasVclResponse`

GetVcls returns the Vcls field if non-nil, zero value otherwise.

### GetVclsOk

`func (o *ServiceVersionDetailOrNull) GetVclsOk() (*[]SchemasVclResponse, bool)`

GetVclsOk returns a tuple with the Vcls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcls

`func (o *ServiceVersionDetailOrNull) SetVcls(v []SchemasVclResponse)`

SetVcls sets Vcls field to given value.

### HasVcls

`func (o *ServiceVersionDetailOrNull) HasVcls() bool`

HasVcls returns a boolean if a field has been set.

### GetWordpress

`func (o *ServiceVersionDetailOrNull) GetWordpress() []*map[string]any`

GetWordpress returns the Wordpress field if non-nil, zero value otherwise.

### GetWordpressOk

`func (o *ServiceVersionDetailOrNull) GetWordpressOk() (*[]*map[string]any, bool)`

GetWordpressOk returns a tuple with the Wordpress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWordpress

`func (o *ServiceVersionDetailOrNull) SetWordpress(v []*map[string]any)`

SetWordpress sets Wordpress field to given value.

### HasWordpress

`func (o *ServiceVersionDetailOrNull) HasWordpress() bool`

HasWordpress returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

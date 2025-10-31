# VersionDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**Wordpress** | Pointer to **[]map[string]interface{}** | A list of Wordpress rules with this service. | [optional] 

## Methods

### NewVersionDetail

`func NewVersionDetail() *VersionDetail`

NewVersionDetail instantiates a new VersionDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionDetailWithDefaults

`func NewVersionDetailWithDefaults() *VersionDetail`

NewVersionDetailWithDefaults instantiates a new VersionDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackends

`func (o *VersionDetail) GetBackends() []BackendResponse`

GetBackends returns the Backends field if non-nil, zero value otherwise.

### GetBackendsOk

`func (o *VersionDetail) GetBackendsOk() (*[]BackendResponse, bool)`

GetBackendsOk returns a tuple with the Backends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackends

`func (o *VersionDetail) SetBackends(v []BackendResponse)`

SetBackends sets Backends field to given value.

### HasBackends

`func (o *VersionDetail) HasBackends() bool`

HasBackends returns a boolean if a field has been set.

### GetCacheSettings

`func (o *VersionDetail) GetCacheSettings() []CacheSettingResponse`

GetCacheSettings returns the CacheSettings field if non-nil, zero value otherwise.

### GetCacheSettingsOk

`func (o *VersionDetail) GetCacheSettingsOk() (*[]CacheSettingResponse, bool)`

GetCacheSettingsOk returns a tuple with the CacheSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheSettings

`func (o *VersionDetail) SetCacheSettings(v []CacheSettingResponse)`

SetCacheSettings sets CacheSettings field to given value.

### HasCacheSettings

`func (o *VersionDetail) HasCacheSettings() bool`

HasCacheSettings returns a boolean if a field has been set.

### GetConditions

`func (o *VersionDetail) GetConditions() []ConditionResponse`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *VersionDetail) GetConditionsOk() (*[]ConditionResponse, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *VersionDetail) SetConditions(v []ConditionResponse)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *VersionDetail) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetDirectors

`func (o *VersionDetail) GetDirectors() []Director`

GetDirectors returns the Directors field if non-nil, zero value otherwise.

### GetDirectorsOk

`func (o *VersionDetail) GetDirectorsOk() (*[]Director, bool)`

GetDirectorsOk returns a tuple with the Directors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectors

`func (o *VersionDetail) SetDirectors(v []Director)`

SetDirectors sets Directors field to given value.

### HasDirectors

`func (o *VersionDetail) HasDirectors() bool`

HasDirectors returns a boolean if a field has been set.

### GetDomains

`func (o *VersionDetail) GetDomains() []DomainResponse`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *VersionDetail) GetDomainsOk() (*[]DomainResponse, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *VersionDetail) SetDomains(v []DomainResponse)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *VersionDetail) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetGzips

`func (o *VersionDetail) GetGzips() []GzipResponse`

GetGzips returns the Gzips field if non-nil, zero value otherwise.

### GetGzipsOk

`func (o *VersionDetail) GetGzipsOk() (*[]GzipResponse, bool)`

GetGzipsOk returns a tuple with the Gzips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGzips

`func (o *VersionDetail) SetGzips(v []GzipResponse)`

SetGzips sets Gzips field to given value.

### HasGzips

`func (o *VersionDetail) HasGzips() bool`

HasGzips returns a boolean if a field has been set.

### GetHeaders

`func (o *VersionDetail) GetHeaders() []HeaderResponse`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *VersionDetail) GetHeadersOk() (*[]HeaderResponse, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *VersionDetail) SetHeaders(v []HeaderResponse)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *VersionDetail) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetHealthchecks

`func (o *VersionDetail) GetHealthchecks() []HealthcheckResponse`

GetHealthchecks returns the Healthchecks field if non-nil, zero value otherwise.

### GetHealthchecksOk

`func (o *VersionDetail) GetHealthchecksOk() (*[]HealthcheckResponse, bool)`

GetHealthchecksOk returns a tuple with the Healthchecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthchecks

`func (o *VersionDetail) SetHealthchecks(v []HealthcheckResponse)`

SetHealthchecks sets Healthchecks field to given value.

### HasHealthchecks

`func (o *VersionDetail) HasHealthchecks() bool`

HasHealthchecks returns a boolean if a field has been set.

### GetRequestSettings

`func (o *VersionDetail) GetRequestSettings() []RequestSettingsResponse`

GetRequestSettings returns the RequestSettings field if non-nil, zero value otherwise.

### GetRequestSettingsOk

`func (o *VersionDetail) GetRequestSettingsOk() (*[]RequestSettingsResponse, bool)`

GetRequestSettingsOk returns a tuple with the RequestSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSettings

`func (o *VersionDetail) SetRequestSettings(v []RequestSettingsResponse)`

SetRequestSettings sets RequestSettings field to given value.

### HasRequestSettings

`func (o *VersionDetail) HasRequestSettings() bool`

HasRequestSettings returns a boolean if a field has been set.

### GetResponseObjects

`func (o *VersionDetail) GetResponseObjects() []ResponseObjectResponse`

GetResponseObjects returns the ResponseObjects field if non-nil, zero value otherwise.

### GetResponseObjectsOk

`func (o *VersionDetail) GetResponseObjectsOk() (*[]ResponseObjectResponse, bool)`

GetResponseObjectsOk returns a tuple with the ResponseObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseObjects

`func (o *VersionDetail) SetResponseObjects(v []ResponseObjectResponse)`

SetResponseObjects sets ResponseObjects field to given value.

### HasResponseObjects

`func (o *VersionDetail) HasResponseObjects() bool`

HasResponseObjects returns a boolean if a field has been set.

### GetSettings

`func (o *VersionDetail) GetSettings() VersionDetailSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *VersionDetail) GetSettingsOk() (*VersionDetailSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *VersionDetail) SetSettings(v VersionDetailSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *VersionDetail) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetSnippets

`func (o *VersionDetail) GetSnippets() []SchemasSnippetResponse`

GetSnippets returns the Snippets field if non-nil, zero value otherwise.

### GetSnippetsOk

`func (o *VersionDetail) GetSnippetsOk() (*[]SchemasSnippetResponse, bool)`

GetSnippetsOk returns a tuple with the Snippets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippets

`func (o *VersionDetail) SetSnippets(v []SchemasSnippetResponse)`

SetSnippets sets Snippets field to given value.

### HasSnippets

`func (o *VersionDetail) HasSnippets() bool`

HasSnippets returns a boolean if a field has been set.

### GetVcls

`func (o *VersionDetail) GetVcls() []SchemasVclResponse`

GetVcls returns the Vcls field if non-nil, zero value otherwise.

### GetVclsOk

`func (o *VersionDetail) GetVclsOk() (*[]SchemasVclResponse, bool)`

GetVclsOk returns a tuple with the Vcls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcls

`func (o *VersionDetail) SetVcls(v []SchemasVclResponse)`

SetVcls sets Vcls field to given value.

### HasVcls

`func (o *VersionDetail) HasVcls() bool`

HasVcls returns a boolean if a field has been set.

### GetWordpress

`func (o *VersionDetail) GetWordpress() []*map[string]interface{}`

GetWordpress returns the Wordpress field if non-nil, zero value otherwise.

### GetWordpressOk

`func (o *VersionDetail) GetWordpressOk() (*[]*map[string]interface{}, bool)`

GetWordpressOk returns a tuple with the Wordpress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWordpress

`func (o *VersionDetail) SetWordpress(v []*map[string]interface{})`

SetWordpress sets Wordpress field to given value.

### HasWordpress

`func (o *VersionDetail) HasWordpress() bool`

HasWordpress returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



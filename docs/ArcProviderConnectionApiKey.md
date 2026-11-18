# ArcProviderConnectionApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A display name for the provider connection. Cannot match a name reserved by a built-in provider. | 
**Models** | **[]string** | The AI model identifiers this connection is allowed to use. | 
**BaseURL** | **string** | The base URL of the provider&#39;s API. Must be an absolute `http` or `https` URL and must not end in `/v1`. | 
**ApiKey** | **string** | The provider&#39;s secret API key. | 
**AuthType** | Pointer to **string** | The authentication type for this connection. Omit for the default `api-key` behavior. | [optional] [default to "api-key"]

## Methods

### NewArcProviderConnectionApiKey

`func NewArcProviderConnectionApiKey(name string, models []string, baseURL string, apiKey string, ) *ArcProviderConnectionApiKey`

NewArcProviderConnectionApiKey instantiates a new ArcProviderConnectionApiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcProviderConnectionApiKeyWithDefaults

`func NewArcProviderConnectionApiKeyWithDefaults() *ArcProviderConnectionApiKey`

NewArcProviderConnectionApiKeyWithDefaults instantiates a new ArcProviderConnectionApiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ArcProviderConnectionApiKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArcProviderConnectionApiKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArcProviderConnectionApiKey) SetName(v string)`

SetName sets Name field to given value.


### GetModels

`func (o *ArcProviderConnectionApiKey) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *ArcProviderConnectionApiKey) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *ArcProviderConnectionApiKey) SetModels(v []string)`

SetModels sets Models field to given value.


### GetBaseURL

`func (o *ArcProviderConnectionApiKey) GetBaseURL() string`

GetBaseURL returns the BaseURL field if non-nil, zero value otherwise.

### GetBaseURLOk

`func (o *ArcProviderConnectionApiKey) GetBaseURLOk() (*string, bool)`

GetBaseURLOk returns a tuple with the BaseURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseURL

`func (o *ArcProviderConnectionApiKey) SetBaseURL(v string)`

SetBaseURL sets BaseURL field to given value.


### GetApiKey

`func (o *ArcProviderConnectionApiKey) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ArcProviderConnectionApiKey) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ArcProviderConnectionApiKey) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetAuthType

`func (o *ArcProviderConnectionApiKey) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *ArcProviderConnectionApiKey) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *ArcProviderConnectionApiKey) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *ArcProviderConnectionApiKey) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

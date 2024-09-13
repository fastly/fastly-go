# ConfiguredProductResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to [**ConfiguredProductResponseProduct**](ConfiguredProductResponseProduct.md) |  | [optional] 
**Service** | Pointer to [**EnabledProductResponseService**](EnabledProductResponseService.md) |  | [optional] 
**Configuration** | Pointer to [**ConfiguredProductResponseConfiguration**](ConfiguredProductResponseConfiguration.md) |  | [optional] 
**Links** | Pointer to [**ConfiguredProductResponseLinks**](ConfiguredProductResponseLinks.md) |  | [optional] 

## Methods

### NewConfiguredProductResponse

`func NewConfiguredProductResponse() *ConfiguredProductResponse`

NewConfiguredProductResponse instantiates a new ConfiguredProductResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfiguredProductResponseWithDefaults

`func NewConfiguredProductResponseWithDefaults() *ConfiguredProductResponse`

NewConfiguredProductResponseWithDefaults instantiates a new ConfiguredProductResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *ConfiguredProductResponse) GetProduct() ConfiguredProductResponseProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *ConfiguredProductResponse) GetProductOk() (*ConfiguredProductResponseProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *ConfiguredProductResponse) SetProduct(v ConfiguredProductResponseProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *ConfiguredProductResponse) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetService

`func (o *ConfiguredProductResponse) GetService() EnabledProductResponseService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ConfiguredProductResponse) GetServiceOk() (*EnabledProductResponseService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ConfiguredProductResponse) SetService(v EnabledProductResponseService)`

SetService sets Service field to given value.

### HasService

`func (o *ConfiguredProductResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### GetConfiguration

`func (o *ConfiguredProductResponse) GetConfiguration() ConfiguredProductResponseConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ConfiguredProductResponse) GetConfigurationOk() (*ConfiguredProductResponseConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ConfiguredProductResponse) SetConfiguration(v ConfiguredProductResponseConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ConfiguredProductResponse) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetLinks

`func (o *ConfiguredProductResponse) GetLinks() ConfiguredProductResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ConfiguredProductResponse) GetLinksOk() (*ConfiguredProductResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ConfiguredProductResponse) SetLinks(v ConfiguredProductResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ConfiguredProductResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

# BotManagementResponseConfigure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to [**BotManagementResponseProductProduct**](BotManagementResponseProductProduct.md) |  | [optional] 
**Service** | Pointer to [**ApiDiscoveryResponseServiceService**](ApiDiscoveryResponseServiceService.md) |  | [optional] 
**Configuration** | Pointer to [**BotManagementResponseConfigurationConfiguration**](BotManagementResponseConfigurationConfiguration.md) |  | [optional] 
**Links** | Pointer to [**BotManagementResponseLinksLinks**](BotManagementResponseLinksLinks.md) |  | [optional] 

## Methods

### NewBotManagementResponseConfigure

`func NewBotManagementResponseConfigure() *BotManagementResponseConfigure`

NewBotManagementResponseConfigure instantiates a new BotManagementResponseConfigure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotManagementResponseConfigureWithDefaults

`func NewBotManagementResponseConfigureWithDefaults() *BotManagementResponseConfigure`

NewBotManagementResponseConfigureWithDefaults instantiates a new BotManagementResponseConfigure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *BotManagementResponseConfigure) GetProduct() BotManagementResponseProductProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *BotManagementResponseConfigure) GetProductOk() (*BotManagementResponseProductProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *BotManagementResponseConfigure) SetProduct(v BotManagementResponseProductProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *BotManagementResponseConfigure) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetService

`func (o *BotManagementResponseConfigure) GetService() ApiDiscoveryResponseServiceService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *BotManagementResponseConfigure) GetServiceOk() (*ApiDiscoveryResponseServiceService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *BotManagementResponseConfigure) SetService(v ApiDiscoveryResponseServiceService)`

SetService sets Service field to given value.

### HasService

`func (o *BotManagementResponseConfigure) HasService() bool`

HasService returns a boolean if a field has been set.

### GetConfiguration

`func (o *BotManagementResponseConfigure) GetConfiguration() BotManagementResponseConfigurationConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *BotManagementResponseConfigure) GetConfigurationOk() (*BotManagementResponseConfigurationConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *BotManagementResponseConfigure) SetConfiguration(v BotManagementResponseConfigurationConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *BotManagementResponseConfigure) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetLinks

`func (o *BotManagementResponseConfigure) GetLinks() BotManagementResponseLinksLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BotManagementResponseConfigure) GetLinksOk() (*BotManagementResponseLinksLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BotManagementResponseConfigure) SetLinks(v BotManagementResponseLinksLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BotManagementResponseConfigure) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



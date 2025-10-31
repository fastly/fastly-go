# NgwafResponseConfigure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to [**NgwafResponseProductProduct**](NgwafResponseProductProduct.md) |  | [optional] 
**Service** | Pointer to [**ApiDiscoveryResponseServiceService**](ApiDiscoveryResponseServiceService.md) |  | [optional] 
**Configuration** | Pointer to [**NgwafResponseConfigurationConfiguration**](NgwafResponseConfigurationConfiguration.md) |  | [optional] 
**Links** | Pointer to [**NgwafResponseLinksLinks**](NgwafResponseLinksLinks.md) |  | [optional] 

## Methods

### NewNgwafResponseConfigure

`func NewNgwafResponseConfigure() *NgwafResponseConfigure`

NewNgwafResponseConfigure instantiates a new NgwafResponseConfigure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNgwafResponseConfigureWithDefaults

`func NewNgwafResponseConfigureWithDefaults() *NgwafResponseConfigure`

NewNgwafResponseConfigureWithDefaults instantiates a new NgwafResponseConfigure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *NgwafResponseConfigure) GetProduct() NgwafResponseProductProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *NgwafResponseConfigure) GetProductOk() (*NgwafResponseProductProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *NgwafResponseConfigure) SetProduct(v NgwafResponseProductProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *NgwafResponseConfigure) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetService

`func (o *NgwafResponseConfigure) GetService() ApiDiscoveryResponseServiceService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *NgwafResponseConfigure) GetServiceOk() (*ApiDiscoveryResponseServiceService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *NgwafResponseConfigure) SetService(v ApiDiscoveryResponseServiceService)`

SetService sets Service field to given value.

### HasService

`func (o *NgwafResponseConfigure) HasService() bool`

HasService returns a boolean if a field has been set.

### GetConfiguration

`func (o *NgwafResponseConfigure) GetConfiguration() NgwafResponseConfigurationConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *NgwafResponseConfigure) GetConfigurationOk() (*NgwafResponseConfigurationConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *NgwafResponseConfigure) SetConfiguration(v NgwafResponseConfigurationConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *NgwafResponseConfigure) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetLinks

`func (o *NgwafResponseConfigure) GetLinks() NgwafResponseLinksLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NgwafResponseConfigure) GetLinksOk() (*NgwafResponseLinksLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NgwafResponseConfigure) SetLinks(v NgwafResponseLinksLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NgwafResponseConfigure) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



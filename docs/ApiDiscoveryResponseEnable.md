# ApiDiscoveryResponseEnable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to [**ApiDiscoveryResponseProductProduct**](ApiDiscoveryResponseProductProduct.md) |  | [optional] 
**Service** | Pointer to [**ApiDiscoveryResponseServiceService**](ApiDiscoveryResponseServiceService.md) |  | [optional] 
**Links** | Pointer to [**ApiDiscoveryResponseLinksLinks**](ApiDiscoveryResponseLinksLinks.md) |  | [optional] 

## Methods

### NewApiDiscoveryResponseEnable

`func NewApiDiscoveryResponseEnable() *ApiDiscoveryResponseEnable`

NewApiDiscoveryResponseEnable instantiates a new ApiDiscoveryResponseEnable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiDiscoveryResponseEnableWithDefaults

`func NewApiDiscoveryResponseEnableWithDefaults() *ApiDiscoveryResponseEnable`

NewApiDiscoveryResponseEnableWithDefaults instantiates a new ApiDiscoveryResponseEnable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *ApiDiscoveryResponseEnable) GetProduct() ApiDiscoveryResponseProductProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *ApiDiscoveryResponseEnable) GetProductOk() (*ApiDiscoveryResponseProductProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *ApiDiscoveryResponseEnable) SetProduct(v ApiDiscoveryResponseProductProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *ApiDiscoveryResponseEnable) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetService

`func (o *ApiDiscoveryResponseEnable) GetService() ApiDiscoveryResponseServiceService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ApiDiscoveryResponseEnable) GetServiceOk() (*ApiDiscoveryResponseServiceService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ApiDiscoveryResponseEnable) SetService(v ApiDiscoveryResponseServiceService)`

SetService sets Service field to given value.

### HasService

`func (o *ApiDiscoveryResponseEnable) HasService() bool`

HasService returns a boolean if a field has been set.

### GetLinks

`func (o *ApiDiscoveryResponseEnable) GetLinks() ApiDiscoveryResponseLinksLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiDiscoveryResponseEnable) GetLinksOk() (*ApiDiscoveryResponseLinksLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiDiscoveryResponseEnable) SetLinks(v ApiDiscoveryResponseLinksLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApiDiscoveryResponseEnable) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



# FanoutResponseBodyEnable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to [**FanoutResponseProductProduct**](FanoutResponseProductProduct.md) |  | [optional] 
**Service** | Pointer to [**ApiDiscoveryResponseServiceService**](ApiDiscoveryResponseServiceService.md) |  | [optional] 
**Links** | Pointer to [**FanoutResponseLinksLinks**](FanoutResponseLinksLinks.md) |  | [optional] 

## Methods

### NewFanoutResponseBodyEnable

`func NewFanoutResponseBodyEnable() *FanoutResponseBodyEnable`

NewFanoutResponseBodyEnable instantiates a new FanoutResponseBodyEnable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFanoutResponseBodyEnableWithDefaults

`func NewFanoutResponseBodyEnableWithDefaults() *FanoutResponseBodyEnable`

NewFanoutResponseBodyEnableWithDefaults instantiates a new FanoutResponseBodyEnable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *FanoutResponseBodyEnable) GetProduct() FanoutResponseProductProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *FanoutResponseBodyEnable) GetProductOk() (*FanoutResponseProductProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *FanoutResponseBodyEnable) SetProduct(v FanoutResponseProductProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *FanoutResponseBodyEnable) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetService

`func (o *FanoutResponseBodyEnable) GetService() ApiDiscoveryResponseServiceService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *FanoutResponseBodyEnable) GetServiceOk() (*ApiDiscoveryResponseServiceService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *FanoutResponseBodyEnable) SetService(v ApiDiscoveryResponseServiceService)`

SetService sets Service field to given value.

### HasService

`func (o *FanoutResponseBodyEnable) HasService() bool`

HasService returns a boolean if a field has been set.

### GetLinks

`func (o *FanoutResponseBodyEnable) GetLinks() FanoutResponseLinksLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FanoutResponseBodyEnable) GetLinksOk() (*FanoutResponseLinksLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FanoutResponseBodyEnable) SetLinks(v FanoutResponseLinksLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FanoutResponseBodyEnable) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



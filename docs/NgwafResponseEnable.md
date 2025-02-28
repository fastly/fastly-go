# NgwafResponseEnable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to [**NgwafResponseProductProduct**](NgwafResponseProductProduct.md) |  | [optional] 
**Service** | Pointer to [**BotManagementResponseServiceService**](BotManagementResponseServiceService.md) |  | [optional] 
**Links** | Pointer to [**NgwafResponseLinksLinks**](NgwafResponseLinksLinks.md) |  | [optional] 

## Methods

### NewNgwafResponseEnable

`func NewNgwafResponseEnable() *NgwafResponseEnable`

NewNgwafResponseEnable instantiates a new NgwafResponseEnable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNgwafResponseEnableWithDefaults

`func NewNgwafResponseEnableWithDefaults() *NgwafResponseEnable`

NewNgwafResponseEnableWithDefaults instantiates a new NgwafResponseEnable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *NgwafResponseEnable) GetProduct() NgwafResponseProductProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *NgwafResponseEnable) GetProductOk() (*NgwafResponseProductProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *NgwafResponseEnable) SetProduct(v NgwafResponseProductProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *NgwafResponseEnable) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetService

`func (o *NgwafResponseEnable) GetService() BotManagementResponseServiceService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *NgwafResponseEnable) GetServiceOk() (*BotManagementResponseServiceService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *NgwafResponseEnable) SetService(v BotManagementResponseServiceService)`

SetService sets Service field to given value.

### HasService

`func (o *NgwafResponseEnable) HasService() bool`

HasService returns a boolean if a field has been set.

### GetLinks

`func (o *NgwafResponseEnable) GetLinks() NgwafResponseLinksLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NgwafResponseEnable) GetLinksOk() (*NgwafResponseLinksLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NgwafResponseEnable) SetLinks(v NgwafResponseLinksLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NgwafResponseEnable) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

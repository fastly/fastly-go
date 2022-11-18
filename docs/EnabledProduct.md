# EnabledProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to [**EnabledProductProduct**](EnabledProductProduct.md) |  | [optional] 
**Service** | Pointer to [**EnabledProductProduct**](EnabledProductProduct.md) |  | [optional] 
**Links** | Pointer to [**EnabledProductLinks**](EnabledProductLinks.md) |  | [optional] 

## Methods

### NewEnabledProduct

`func NewEnabledProduct() *EnabledProduct`

NewEnabledProduct instantiates a new EnabledProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnabledProductWithDefaults

`func NewEnabledProductWithDefaults() *EnabledProduct`

NewEnabledProductWithDefaults instantiates a new EnabledProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *EnabledProduct) GetProduct() EnabledProductProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *EnabledProduct) GetProductOk() (*EnabledProductProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *EnabledProduct) SetProduct(v EnabledProductProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *EnabledProduct) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetService

`func (o *EnabledProduct) GetService() EnabledProductProduct`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *EnabledProduct) GetServiceOk() (*EnabledProductProduct, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *EnabledProduct) SetService(v EnabledProductProduct)`

SetService sets Service field to given value.

### HasService

`func (o *EnabledProduct) HasService() bool`

HasService returns a boolean if a field has been set.

### GetLinks

`func (o *EnabledProduct) GetLinks() EnabledProductLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnabledProduct) GetLinksOk() (*EnabledProductLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnabledProduct) SetLinks(v EnabledProductLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EnabledProduct) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

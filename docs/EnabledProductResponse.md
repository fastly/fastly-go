# EnabledProductResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to [**EnabledProductResponseProduct**](EnabledProductResponseProduct.md) |  | [optional] 
**Service** | Pointer to [**EnabledProductResponseService**](EnabledProductResponseService.md) |  | [optional] 
**Links** | Pointer to [**EnabledProductResponseLinks**](EnabledProductResponseLinks.md) |  | [optional] 

## Methods

### NewEnabledProductResponse

`func NewEnabledProductResponse() *EnabledProductResponse`

NewEnabledProductResponse instantiates a new EnabledProductResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnabledProductResponseWithDefaults

`func NewEnabledProductResponseWithDefaults() *EnabledProductResponse`

NewEnabledProductResponseWithDefaults instantiates a new EnabledProductResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *EnabledProductResponse) GetProduct() EnabledProductResponseProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *EnabledProductResponse) GetProductOk() (*EnabledProductResponseProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *EnabledProductResponse) SetProduct(v EnabledProductResponseProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *EnabledProductResponse) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetService

`func (o *EnabledProductResponse) GetService() EnabledProductResponseService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *EnabledProductResponse) GetServiceOk() (*EnabledProductResponseService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *EnabledProductResponse) SetService(v EnabledProductResponseService)`

SetService sets Service field to given value.

### HasService

`func (o *EnabledProductResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### GetLinks

`func (o *EnabledProductResponse) GetLinks() EnabledProductResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnabledProductResponse) GetLinksOk() (*EnabledProductResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnabledProductResponse) SetLinks(v EnabledProductResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EnabledProductResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

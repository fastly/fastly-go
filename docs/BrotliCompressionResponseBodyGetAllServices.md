# BrotliCompressionResponseBodyGetAllServices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to [**BrotliCompressionResponseProductProduct**](BrotliCompressionResponseProductProduct.md) |  | [optional] 
**Customer** | Pointer to [**BotManagementResponseCustomerCustomer**](BotManagementResponseCustomerCustomer.md) |  | [optional] 
**Services** | Pointer to **[]string** | A list of services with Brotli Compression enabled. | [optional] 
**Links** | Pointer to [**BrotliCompressionResponseLinksGetAllServicesLinks**](BrotliCompressionResponseLinksGetAllServicesLinks.md) |  | [optional] 

## Methods

### NewBrotliCompressionResponseBodyGetAllServices

`func NewBrotliCompressionResponseBodyGetAllServices() *BrotliCompressionResponseBodyGetAllServices`

NewBrotliCompressionResponseBodyGetAllServices instantiates a new BrotliCompressionResponseBodyGetAllServices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrotliCompressionResponseBodyGetAllServicesWithDefaults

`func NewBrotliCompressionResponseBodyGetAllServicesWithDefaults() *BrotliCompressionResponseBodyGetAllServices`

NewBrotliCompressionResponseBodyGetAllServicesWithDefaults instantiates a new BrotliCompressionResponseBodyGetAllServices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *BrotliCompressionResponseBodyGetAllServices) GetProduct() BrotliCompressionResponseProductProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *BrotliCompressionResponseBodyGetAllServices) GetProductOk() (*BrotliCompressionResponseProductProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *BrotliCompressionResponseBodyGetAllServices) SetProduct(v BrotliCompressionResponseProductProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *BrotliCompressionResponseBodyGetAllServices) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetCustomer

`func (o *BrotliCompressionResponseBodyGetAllServices) GetCustomer() BotManagementResponseCustomerCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *BrotliCompressionResponseBodyGetAllServices) GetCustomerOk() (*BotManagementResponseCustomerCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *BrotliCompressionResponseBodyGetAllServices) SetCustomer(v BotManagementResponseCustomerCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *BrotliCompressionResponseBodyGetAllServices) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetServices

`func (o *BrotliCompressionResponseBodyGetAllServices) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *BrotliCompressionResponseBodyGetAllServices) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *BrotliCompressionResponseBodyGetAllServices) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *BrotliCompressionResponseBodyGetAllServices) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetLinks

`func (o *BrotliCompressionResponseBodyGetAllServices) GetLinks() BrotliCompressionResponseLinksGetAllServicesLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BrotliCompressionResponseBodyGetAllServices) GetLinksOk() (*BrotliCompressionResponseLinksGetAllServicesLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BrotliCompressionResponseBodyGetAllServices) SetLinks(v BrotliCompressionResponseLinksGetAllServicesLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BrotliCompressionResponseBodyGetAllServices) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



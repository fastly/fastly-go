# BotManagementResponseBodyGetAllServices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to [**BotManagementResponseProductProduct**](BotManagementResponseProductProduct.md) |  | [optional] 
**Customer** | Pointer to [**BotManagementResponseCustomerCustomer**](BotManagementResponseCustomerCustomer.md) |  | [optional] 
**Services** | Pointer to **[]string** | A list of services with Bot Management enabled. | [optional] 
**Links** | Pointer to [**BotManagementResponseLinksGetAllServicesLinks**](BotManagementResponseLinksGetAllServicesLinks.md) |  | [optional] 

## Methods

### NewBotManagementResponseBodyGetAllServices

`func NewBotManagementResponseBodyGetAllServices() *BotManagementResponseBodyGetAllServices`

NewBotManagementResponseBodyGetAllServices instantiates a new BotManagementResponseBodyGetAllServices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotManagementResponseBodyGetAllServicesWithDefaults

`func NewBotManagementResponseBodyGetAllServicesWithDefaults() *BotManagementResponseBodyGetAllServices`

NewBotManagementResponseBodyGetAllServicesWithDefaults instantiates a new BotManagementResponseBodyGetAllServices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *BotManagementResponseBodyGetAllServices) GetProduct() BotManagementResponseProductProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *BotManagementResponseBodyGetAllServices) GetProductOk() (*BotManagementResponseProductProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *BotManagementResponseBodyGetAllServices) SetProduct(v BotManagementResponseProductProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *BotManagementResponseBodyGetAllServices) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetCustomer

`func (o *BotManagementResponseBodyGetAllServices) GetCustomer() BotManagementResponseCustomerCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *BotManagementResponseBodyGetAllServices) GetCustomerOk() (*BotManagementResponseCustomerCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *BotManagementResponseBodyGetAllServices) SetCustomer(v BotManagementResponseCustomerCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *BotManagementResponseBodyGetAllServices) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetServices

`func (o *BotManagementResponseBodyGetAllServices) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *BotManagementResponseBodyGetAllServices) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *BotManagementResponseBodyGetAllServices) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *BotManagementResponseBodyGetAllServices) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetLinks

`func (o *BotManagementResponseBodyGetAllServices) GetLinks() BotManagementResponseLinksGetAllServicesLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BotManagementResponseBodyGetAllServices) GetLinksOk() (*BotManagementResponseLinksGetAllServicesLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BotManagementResponseBodyGetAllServices) SetLinks(v BotManagementResponseLinksGetAllServicesLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BotManagementResponseBodyGetAllServices) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

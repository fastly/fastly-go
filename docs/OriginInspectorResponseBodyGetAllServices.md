# OriginInspectorResponseBodyGetAllServices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to [**OriginInspectorResponseProductProduct**](OriginInspectorResponseProductProduct.md) |  | [optional] 
**Customer** | Pointer to [**BotManagementResponseCustomerCustomer**](BotManagementResponseCustomerCustomer.md) |  | [optional] 
**Services** | Pointer to **[]string** | A list of services with Origin Inspector enabled. | [optional] 
**Links** | Pointer to [**OriginInspectorResponseLinksGetAllServicesLinks**](OriginInspectorResponseLinksGetAllServicesLinks.md) |  | [optional] 

## Methods

### NewOriginInspectorResponseBodyGetAllServices

`func NewOriginInspectorResponseBodyGetAllServices() *OriginInspectorResponseBodyGetAllServices`

NewOriginInspectorResponseBodyGetAllServices instantiates a new OriginInspectorResponseBodyGetAllServices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginInspectorResponseBodyGetAllServicesWithDefaults

`func NewOriginInspectorResponseBodyGetAllServicesWithDefaults() *OriginInspectorResponseBodyGetAllServices`

NewOriginInspectorResponseBodyGetAllServicesWithDefaults instantiates a new OriginInspectorResponseBodyGetAllServices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *OriginInspectorResponseBodyGetAllServices) GetProduct() OriginInspectorResponseProductProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *OriginInspectorResponseBodyGetAllServices) GetProductOk() (*OriginInspectorResponseProductProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *OriginInspectorResponseBodyGetAllServices) SetProduct(v OriginInspectorResponseProductProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *OriginInspectorResponseBodyGetAllServices) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetCustomer

`func (o *OriginInspectorResponseBodyGetAllServices) GetCustomer() BotManagementResponseCustomerCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *OriginInspectorResponseBodyGetAllServices) GetCustomerOk() (*BotManagementResponseCustomerCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *OriginInspectorResponseBodyGetAllServices) SetCustomer(v BotManagementResponseCustomerCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *OriginInspectorResponseBodyGetAllServices) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetServices

`func (o *OriginInspectorResponseBodyGetAllServices) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *OriginInspectorResponseBodyGetAllServices) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *OriginInspectorResponseBodyGetAllServices) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *OriginInspectorResponseBodyGetAllServices) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetLinks

`func (o *OriginInspectorResponseBodyGetAllServices) GetLinks() OriginInspectorResponseLinksGetAllServicesLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OriginInspectorResponseBodyGetAllServices) GetLinksOk() (*OriginInspectorResponseLinksGetAllServicesLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OriginInspectorResponseBodyGetAllServices) SetLinks(v OriginInspectorResponseLinksGetAllServicesLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OriginInspectorResponseBodyGetAllServices) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



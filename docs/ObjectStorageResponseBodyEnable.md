# ObjectStorageResponseBodyEnable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to [**ObjectStorageResponseProductProduct**](ObjectStorageResponseProductProduct.md) |  | [optional] 
**Customer** | Pointer to [**AiAcceleratorResponseCustomerCustomer**](AiAcceleratorResponseCustomerCustomer.md) |  | [optional] 
**Links** | Pointer to [**ObjectStorageResponseLinksLinks**](ObjectStorageResponseLinksLinks.md) |  | [optional] 

## Methods

### NewObjectStorageResponseBodyEnable

`func NewObjectStorageResponseBodyEnable() *ObjectStorageResponseBodyEnable`

NewObjectStorageResponseBodyEnable instantiates a new ObjectStorageResponseBodyEnable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageResponseBodyEnableWithDefaults

`func NewObjectStorageResponseBodyEnableWithDefaults() *ObjectStorageResponseBodyEnable`

NewObjectStorageResponseBodyEnableWithDefaults instantiates a new ObjectStorageResponseBodyEnable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *ObjectStorageResponseBodyEnable) GetProduct() ObjectStorageResponseProductProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *ObjectStorageResponseBodyEnable) GetProductOk() (*ObjectStorageResponseProductProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *ObjectStorageResponseBodyEnable) SetProduct(v ObjectStorageResponseProductProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *ObjectStorageResponseBodyEnable) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetCustomer

`func (o *ObjectStorageResponseBodyEnable) GetCustomer() AiAcceleratorResponseCustomerCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *ObjectStorageResponseBodyEnable) GetCustomerOk() (*AiAcceleratorResponseCustomerCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *ObjectStorageResponseBodyEnable) SetCustomer(v AiAcceleratorResponseCustomerCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *ObjectStorageResponseBodyEnable) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetLinks

`func (o *ObjectStorageResponseBodyEnable) GetLinks() ObjectStorageResponseLinksLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ObjectStorageResponseBodyEnable) GetLinksOk() (*ObjectStorageResponseLinksLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ObjectStorageResponseBodyEnable) SetLinks(v ObjectStorageResponseLinksLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ObjectStorageResponseBodyEnable) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

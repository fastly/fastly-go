# UpdateBillingAddressRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeBillingAddress**](TypeBillingAddress.md) |  | [optional] [default to TYPEBILLINGADDRESS_BILLING_ADDRESS]
**Id** | Pointer to **string** | Alphanumeric string identifying the billing address. | [optional] [readonly] 
**Attributes** | Pointer to [**BillingAddressAttributes**](BillingAddressAttributes.md) |  | [optional] 

## Methods

### NewUpdateBillingAddressRequestData

`func NewUpdateBillingAddressRequestData() *UpdateBillingAddressRequestData`

NewUpdateBillingAddressRequestData instantiates a new UpdateBillingAddressRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBillingAddressRequestDataWithDefaults

`func NewUpdateBillingAddressRequestDataWithDefaults() *UpdateBillingAddressRequestData`

NewUpdateBillingAddressRequestDataWithDefaults instantiates a new UpdateBillingAddressRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateBillingAddressRequestData) GetType() TypeBillingAddress`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateBillingAddressRequestData) GetTypeOk() (*TypeBillingAddress, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateBillingAddressRequestData) SetType(v TypeBillingAddress)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateBillingAddressRequestData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *UpdateBillingAddressRequestData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateBillingAddressRequestData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateBillingAddressRequestData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateBillingAddressRequestData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAttributes

`func (o *UpdateBillingAddressRequestData) GetAttributes() BillingAddressAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpdateBillingAddressRequestData) GetAttributesOk() (*BillingAddressAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UpdateBillingAddressRequestData) SetAttributes(v BillingAddressAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UpdateBillingAddressRequestData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



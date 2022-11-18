# UpdateBillingAddressRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeBillingAddress**](TypeBillingAddress.md) |  | [optional] [default to TYPEBILLINGADDRESS_BILLING_ADDRESS]
**ID** | Pointer to **string** | Alphanumeric string identifying the billing address. | [optional] [readonly] 
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

### GetID

`func (o *UpdateBillingAddressRequestData) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *UpdateBillingAddressRequestData) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *UpdateBillingAddressRequestData) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *UpdateBillingAddressRequestData) HasID() bool`

HasID returns a boolean if a field has been set.

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

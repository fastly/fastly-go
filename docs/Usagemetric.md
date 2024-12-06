# Usagemetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Month** | Pointer to **string** | The year and month of the usage element. | [optional] 
**UsageType** | Pointer to **string** | The usage type identifier for the usage. This is a single, billable metric for the product. | [optional] 
**Name** | Pointer to **string** | Full name of the product usage type as it might appear on a customer&#39;s invoice. | [optional] 
**Region** | Pointer to **string** | The geographical area applicable for regionally based products. | [optional] 
**Unit** | Pointer to **string** | The unit for the usage as shown on an invoice. If there is no explicit unit, this field will be \&quot;unit\&quot; (e.g., a request with `product_id` of &#39;cdn_usage&#39; and `usage_type` of &#39;North America Requests&#39; has no unit, and will return \&quot;unit\&quot;). | [optional] 
**Quantity** | Pointer to **float32** | The quantity of the usage for the product. | [optional] 
**RawQuantity** | Pointer to **float32** | The raw units measured for the product. | [optional] 
**ProductID** | Pointer to **string** | The product identifier associated with the usage type. This corresponds to a Fastly product offering. | [optional] 
**LastUpdatedAt** | Pointer to **string** | The date when the usage metric was last updated. | [optional] 

## Methods

### NewUsagemetric

`func NewUsagemetric() *Usagemetric`

NewUsagemetric instantiates a new Usagemetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsagemetricWithDefaults

`func NewUsagemetricWithDefaults() *Usagemetric`

NewUsagemetricWithDefaults instantiates a new Usagemetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonth

`func (o *Usagemetric) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *Usagemetric) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *Usagemetric) SetMonth(v string)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *Usagemetric) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetUsageType

`func (o *Usagemetric) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *Usagemetric) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *Usagemetric) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *Usagemetric) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### GetName

`func (o *Usagemetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Usagemetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Usagemetric) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Usagemetric) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *Usagemetric) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Usagemetric) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Usagemetric) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Usagemetric) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetUnit

`func (o *Usagemetric) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Usagemetric) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Usagemetric) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *Usagemetric) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetQuantity

`func (o *Usagemetric) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Usagemetric) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Usagemetric) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *Usagemetric) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetRawQuantity

`func (o *Usagemetric) GetRawQuantity() float32`

GetRawQuantity returns the RawQuantity field if non-nil, zero value otherwise.

### GetRawQuantityOk

`func (o *Usagemetric) GetRawQuantityOk() (*float32, bool)`

GetRawQuantityOk returns a tuple with the RawQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawQuantity

`func (o *Usagemetric) SetRawQuantity(v float32)`

SetRawQuantity sets RawQuantity field to given value.

### HasRawQuantity

`func (o *Usagemetric) HasRawQuantity() bool`

HasRawQuantity returns a boolean if a field has been set.

### GetProductID

`func (o *Usagemetric) GetProductID() string`

GetProductID returns the ProductID field if non-nil, zero value otherwise.

### GetProductIDOk

`func (o *Usagemetric) GetProductIDOk() (*string, bool)`

GetProductIDOk returns a tuple with the ProductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductID

`func (o *Usagemetric) SetProductID(v string)`

SetProductID sets ProductID field to given value.

### HasProductID

`func (o *Usagemetric) HasProductID() bool`

HasProductID returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *Usagemetric) GetLastUpdatedAt() string`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *Usagemetric) GetLastUpdatedAtOk() (*string, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *Usagemetric) SetLastUpdatedAt(v string)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *Usagemetric) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

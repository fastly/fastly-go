# ServiceusagemetricsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **string** |  | [optional] [readonly] 
**StartTime** | Pointer to **time.Time** | Date and time (in ISO 8601 format) for initiation point of a billing cycle, signifying the start of charges for a service or subscription. | [optional] 
**EndTime** | Pointer to **time.Time** | Date and time (in ISO 8601 format) for termination point of a billing cycle, signifying the end of charges for a service or subscription. | [optional] 
**UsageType** | Pointer to **string** | The usage type identifier for the usage. This is a single, billable metric for the product. | [optional] 
**Unit** | Pointer to **string** | The unit for the usage as shown on an invoice. If there is no explicit unit, this field will be \&quot;unit\&quot; (e.g., a request with `product_id` of &#39;cdn_usage&#39; and `usage_type` of &#39;North America Requests&#39; has no unit, and will return \&quot;unit\&quot;). | [optional] 
**Details** | Pointer to [**[]Serviceusagemetric**](Serviceusagemetric.md) |  | [optional] 
**Meta** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 

## Methods

### NewServiceusagemetricsData

`func NewServiceusagemetricsData() *ServiceusagemetricsData`

NewServiceusagemetricsData instantiates a new ServiceusagemetricsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceusagemetricsDataWithDefaults

`func NewServiceusagemetricsDataWithDefaults() *ServiceusagemetricsData`

NewServiceusagemetricsDataWithDefaults instantiates a new ServiceusagemetricsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *ServiceusagemetricsData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ServiceusagemetricsData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ServiceusagemetricsData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *ServiceusagemetricsData) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetStartTime

`func (o *ServiceusagemetricsData) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ServiceusagemetricsData) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ServiceusagemetricsData) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ServiceusagemetricsData) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *ServiceusagemetricsData) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ServiceusagemetricsData) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ServiceusagemetricsData) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ServiceusagemetricsData) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetUsageType

`func (o *ServiceusagemetricsData) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *ServiceusagemetricsData) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *ServiceusagemetricsData) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *ServiceusagemetricsData) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### GetUnit

`func (o *ServiceusagemetricsData) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ServiceusagemetricsData) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ServiceusagemetricsData) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *ServiceusagemetricsData) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetDetails

`func (o *ServiceusagemetricsData) GetDetails() []Serviceusagemetric`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ServiceusagemetricsData) GetDetailsOk() (*[]Serviceusagemetric, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ServiceusagemetricsData) SetDetails(v []Serviceusagemetric)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ServiceusagemetricsData) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetMeta

`func (o *ServiceusagemetricsData) GetMeta() Metadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ServiceusagemetricsData) GetMetaOk() (*Metadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ServiceusagemetricsData) SetMeta(v Metadata)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ServiceusagemetricsData) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



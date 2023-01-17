# WafFirewallResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**WafFirewallResponseData**](WafFirewallResponseData.md) |  | [optional] 
**Included** | Pointer to [**[]SchemasWafFirewallVersion**](SchemasWafFirewallVersion.md) |  | [optional] 

## Methods

### NewWafFirewallResponse

`func NewWafFirewallResponse() *WafFirewallResponse`

NewWafFirewallResponse instantiates a new WafFirewallResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafFirewallResponseWithDefaults

`func NewWafFirewallResponseWithDefaults() *WafFirewallResponse`

NewWafFirewallResponseWithDefaults instantiates a new WafFirewallResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WafFirewallResponse) GetData() WafFirewallResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WafFirewallResponse) GetDataOk() (*WafFirewallResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WafFirewallResponse) SetData(v WafFirewallResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *WafFirewallResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetIncluded

`func (o *WafFirewallResponse) GetIncluded() []SchemasWafFirewallVersion`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *WafFirewallResponse) GetIncludedOk() (*[]SchemasWafFirewallVersion, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *WafFirewallResponse) SetIncluded(v []SchemasWafFirewallVersion)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *WafFirewallResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

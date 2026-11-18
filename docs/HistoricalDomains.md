# HistoricalDomains

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Whether or not we were able to successfully execute the query. | [optional] 
**Meta** | Pointer to [**HistoricalDomainsMeta**](HistoricalDomainsMeta.md) |  | [optional] 
**Msg** | Pointer to **NullableString** | If the query was not successful, this will provide a string that explains why. | [optional] 
**Data** | Pointer to [**[]HistoricalDomainsData**](HistoricalDomainsData.md) | A list of [entries](#entry-data-model), each representing one unique combination of dimensions, such as domain, region, or POP. | [optional] 

## Methods

### NewHistoricalDomains

`func NewHistoricalDomains() *HistoricalDomains`

NewHistoricalDomains instantiates a new HistoricalDomains object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricalDomainsWithDefaults

`func NewHistoricalDomainsWithDefaults() *HistoricalDomains`

NewHistoricalDomainsWithDefaults instantiates a new HistoricalDomains object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *HistoricalDomains) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HistoricalDomains) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HistoricalDomains) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HistoricalDomains) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMeta

`func (o *HistoricalDomains) GetMeta() HistoricalDomainsMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *HistoricalDomains) GetMetaOk() (*HistoricalDomainsMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *HistoricalDomains) SetMeta(v HistoricalDomainsMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *HistoricalDomains) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetMsg

`func (o *HistoricalDomains) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *HistoricalDomains) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *HistoricalDomains) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *HistoricalDomains) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### SetMsgNil

`func (o *HistoricalDomains) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *HistoricalDomains) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetData

`func (o *HistoricalDomains) GetData() []HistoricalDomainsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HistoricalDomains) GetDataOk() (*[]HistoricalDomainsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HistoricalDomains) SetData(v []HistoricalDomainsData)`

SetData sets Data field to given value.

### HasData

`func (o *HistoricalDomains) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

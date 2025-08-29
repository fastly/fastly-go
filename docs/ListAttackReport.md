# ListAttackReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]AttackReport**](AttackReport.md) |  | [optional] 
**Meta** | Pointer to [**ListAttackReportMeta**](ListAttackReportMeta.md) |  | [optional] 

## Methods

### NewListAttackReport

`func NewListAttackReport() *ListAttackReport`

NewListAttackReport instantiates a new ListAttackReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAttackReportWithDefaults

`func NewListAttackReportWithDefaults() *ListAttackReport`

NewListAttackReportWithDefaults instantiates a new ListAttackReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListAttackReport) GetData() []AttackReport`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListAttackReport) GetDataOk() (*[]AttackReport, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListAttackReport) SetData(v []AttackReport)`

SetData sets Data field to given value.

### HasData

`func (o *ListAttackReport) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListAttackReport) GetMeta() ListAttackReportMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListAttackReport) GetMetaOk() (*ListAttackReportMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListAttackReport) SetMeta(v ListAttackReportMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListAttackReport) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)

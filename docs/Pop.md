# Pop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | the three-letter code for the [POP](https://www.fastly.com/documentation/learning/concepts/pop/) | 
**Name** | **string** | the name of the POP | 
**Group** | **string** |  | 
**Region** | **string** |  | 
**StatsRegion** | **string** | the region used for stats reporting | 
**BillingRegion** | **string** | the region used for billing | 
**Coordinates** | Pointer to [**PopCoordinates**](PopCoordinates.md) |  | [optional] 
**Shield** | Pointer to **string** | the name of the [shield code](https://www.fastly.com/documentation/learning/concepts/shielding/#choosing-a-shield-location) if this POP is suitable for shielding | [optional] 

## Methods

### NewPop

`func NewPop(code string, name string, group string, region string, statsRegion string, billingRegion string, ) *Pop`

NewPop instantiates a new Pop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPopWithDefaults

`func NewPopWithDefaults() *Pop`

NewPopWithDefaults instantiates a new Pop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Pop) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Pop) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Pop) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *Pop) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Pop) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Pop) SetName(v string)`

SetName sets Name field to given value.


### GetGroup

`func (o *Pop) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Pop) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Pop) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetRegion

`func (o *Pop) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Pop) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Pop) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetStatsRegion

`func (o *Pop) GetStatsRegion() string`

GetStatsRegion returns the StatsRegion field if non-nil, zero value otherwise.

### GetStatsRegionOk

`func (o *Pop) GetStatsRegionOk() (*string, bool)`

GetStatsRegionOk returns a tuple with the StatsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsRegion

`func (o *Pop) SetStatsRegion(v string)`

SetStatsRegion sets StatsRegion field to given value.


### GetBillingRegion

`func (o *Pop) GetBillingRegion() string`

GetBillingRegion returns the BillingRegion field if non-nil, zero value otherwise.

### GetBillingRegionOk

`func (o *Pop) GetBillingRegionOk() (*string, bool)`

GetBillingRegionOk returns a tuple with the BillingRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingRegion

`func (o *Pop) SetBillingRegion(v string)`

SetBillingRegion sets BillingRegion field to given value.


### GetCoordinates

`func (o *Pop) GetCoordinates() PopCoordinates`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *Pop) GetCoordinatesOk() (*PopCoordinates, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *Pop) SetCoordinates(v PopCoordinates)`

SetCoordinates sets Coordinates field to given value.

### HasCoordinates

`func (o *Pop) HasCoordinates() bool`

HasCoordinates returns a boolean if a field has been set.

### GetShield

`func (o *Pop) GetShield() string`

GetShield returns the Shield field if non-nil, zero value otherwise.

### GetShieldOk

`func (o *Pop) GetShieldOk() (*string, bool)`

GetShieldOk returns a tuple with the Shield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShield

`func (o *Pop) SetShield(v string)`

SetShield sets Shield field to given value.

### HasShield

`func (o *Pop) HasShield() bool`

HasShield returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



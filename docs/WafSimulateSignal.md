# WafSimulateSignal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of signal detected (e.g., `SQLI`, `XSS`, `CMDEXE`, `TRAVERSAL`, `BACKDOOR`, `LOG4J-JNDI`, `BLOCKED`). | 
**Detector** | **string** | The detector engine that identified the signal (e.g., `SQLI`, `LIBINJECTIONV5`, `LIBINJECTIONJS`, or a rule ID). | 
**DetectorScope** | **string** | The scope of the detector that identified the signal. Derived from the signal type and detection type at simulation time. `system` — built-in WAF rule (e.g., `SQLI`, `XSS`). `workspace` — workspace-level custom rule or signal (e.g., `site.*` prefix). `account` — account-level custom signal (e.g., `corp.*` prefix). `unknown` — scope could not be determined (e.g., tags fetch failed or unrecognized type). | 
**Redaction** | **string** | The redaction level applied to the detected value. Clients should handle unexpected string values gracefully, as new redaction types may be added. | 
**Location** | Pointer to **string** | Where in the request the signal was detected (e.g., `QUERYSTRING`, `POSTBODY`, `HEADER`, `HEADEROUT`, `POSTARG`). Present for detection signals; absent for custom and action signals. | [optional] 
**Name** | Pointer to **string** | The parameter or header name that triggered detection. Present when the WAF engine identifies a specific parameter or header. | [optional] 
**Value** | Pointer to **string** | The matched payload value that triggered signal detection. For detection signals, contains the matched content. For `BLOCKED` signals, carries the WAF response code as a string. Absent for custom signals. | [optional] 

## Methods

### NewWafSimulateSignal

`func NewWafSimulateSignal(type_ string, detector string, detectorScope string, redaction string, ) *WafSimulateSignal`

NewWafSimulateSignal instantiates a new WafSimulateSignal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafSimulateSignalWithDefaults

`func NewWafSimulateSignalWithDefaults() *WafSimulateSignal`

NewWafSimulateSignalWithDefaults instantiates a new WafSimulateSignal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WafSimulateSignal) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WafSimulateSignal) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WafSimulateSignal) SetType(v string)`

SetType sets Type field to given value.


### GetDetector

`func (o *WafSimulateSignal) GetDetector() string`

GetDetector returns the Detector field if non-nil, zero value otherwise.

### GetDetectorOk

`func (o *WafSimulateSignal) GetDetectorOk() (*string, bool)`

GetDetectorOk returns a tuple with the Detector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetector

`func (o *WafSimulateSignal) SetDetector(v string)`

SetDetector sets Detector field to given value.


### GetDetectorScope

`func (o *WafSimulateSignal) GetDetectorScope() string`

GetDetectorScope returns the DetectorScope field if non-nil, zero value otherwise.

### GetDetectorScopeOk

`func (o *WafSimulateSignal) GetDetectorScopeOk() (*string, bool)`

GetDetectorScopeOk returns a tuple with the DetectorScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectorScope

`func (o *WafSimulateSignal) SetDetectorScope(v string)`

SetDetectorScope sets DetectorScope field to given value.


### GetRedaction

`func (o *WafSimulateSignal) GetRedaction() string`

GetRedaction returns the Redaction field if non-nil, zero value otherwise.

### GetRedactionOk

`func (o *WafSimulateSignal) GetRedactionOk() (*string, bool)`

GetRedactionOk returns a tuple with the Redaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedaction

`func (o *WafSimulateSignal) SetRedaction(v string)`

SetRedaction sets Redaction field to given value.


### GetLocation

`func (o *WafSimulateSignal) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *WafSimulateSignal) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *WafSimulateSignal) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *WafSimulateSignal) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *WafSimulateSignal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WafSimulateSignal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WafSimulateSignal) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WafSimulateSignal) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *WafSimulateSignal) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WafSimulateSignal) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WafSimulateSignal) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *WafSimulateSignal) HasValue() bool`

HasValue returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)



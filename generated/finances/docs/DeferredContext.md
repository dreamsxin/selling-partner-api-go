# DeferredContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeferralReason** | Pointer to **string** | The deferral policy applied to the transaction.  **Examples:** &#x60;B2B&#x60; (invoiced orders), &#x60;DD7&#x60; (delivery date policy) | [optional] 
**MaturityDate** | Pointer to **time.Time** | A date in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date-time format. | [optional] 

## Methods

### NewDeferredContext

`func NewDeferredContext() *DeferredContext`

NewDeferredContext instantiates a new DeferredContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeferredContextWithDefaults

`func NewDeferredContextWithDefaults() *DeferredContext`

NewDeferredContextWithDefaults instantiates a new DeferredContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeferralReason

`func (o *DeferredContext) GetDeferralReason() string`

GetDeferralReason returns the DeferralReason field if non-nil, zero value otherwise.

### GetDeferralReasonOk

`func (o *DeferredContext) GetDeferralReasonOk() (*string, bool)`

GetDeferralReasonOk returns a tuple with the DeferralReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferralReason

`func (o *DeferredContext) SetDeferralReason(v string)`

SetDeferralReason sets DeferralReason field to given value.

### HasDeferralReason

`func (o *DeferredContext) HasDeferralReason() bool`

HasDeferralReason returns a boolean if a field has been set.

### GetMaturityDate

`func (o *DeferredContext) GetMaturityDate() time.Time`

GetMaturityDate returns the MaturityDate field if non-nil, zero value otherwise.

### GetMaturityDateOk

`func (o *DeferredContext) GetMaturityDateOk() (*time.Time, bool)`

GetMaturityDateOk returns a tuple with the MaturityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaturityDate

`func (o *DeferredContext) SetMaturityDate(v time.Time)`

SetMaturityDate sets MaturityDate field to given value.

### HasMaturityDate

`func (o *DeferredContext) HasMaturityDate() bool`

HasMaturityDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



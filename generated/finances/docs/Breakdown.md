# Breakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BreakdownType** | Pointer to **string** | The type of charge. | [optional] 
**BreakdownAmount** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**Breakdowns** | Pointer to [**[]Breakdown**](Breakdown.md) | A list of breakdowns that detail how the total amount is calculated for the transaction. | [optional] 

## Methods

### NewBreakdown

`func NewBreakdown() *Breakdown`

NewBreakdown instantiates a new Breakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBreakdownWithDefaults

`func NewBreakdownWithDefaults() *Breakdown`

NewBreakdownWithDefaults instantiates a new Breakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBreakdownType

`func (o *Breakdown) GetBreakdownType() string`

GetBreakdownType returns the BreakdownType field if non-nil, zero value otherwise.

### GetBreakdownTypeOk

`func (o *Breakdown) GetBreakdownTypeOk() (*string, bool)`

GetBreakdownTypeOk returns a tuple with the BreakdownType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownType

`func (o *Breakdown) SetBreakdownType(v string)`

SetBreakdownType sets BreakdownType field to given value.

### HasBreakdownType

`func (o *Breakdown) HasBreakdownType() bool`

HasBreakdownType returns a boolean if a field has been set.

### GetBreakdownAmount

`func (o *Breakdown) GetBreakdownAmount() Currency`

GetBreakdownAmount returns the BreakdownAmount field if non-nil, zero value otherwise.

### GetBreakdownAmountOk

`func (o *Breakdown) GetBreakdownAmountOk() (*Currency, bool)`

GetBreakdownAmountOk returns a tuple with the BreakdownAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownAmount

`func (o *Breakdown) SetBreakdownAmount(v Currency)`

SetBreakdownAmount sets BreakdownAmount field to given value.

### HasBreakdownAmount

`func (o *Breakdown) HasBreakdownAmount() bool`

HasBreakdownAmount returns a boolean if a field has been set.

### GetBreakdowns

`func (o *Breakdown) GetBreakdowns() []Breakdown`

GetBreakdowns returns the Breakdowns field if non-nil, zero value otherwise.

### GetBreakdownsOk

`func (o *Breakdown) GetBreakdownsOk() (*[]Breakdown, bool)`

GetBreakdownsOk returns a tuple with the Breakdowns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdowns

`func (o *Breakdown) SetBreakdowns(v []Breakdown)`

SetBreakdowns sets Breakdowns field to given value.

### HasBreakdowns

`func (o *Breakdown) HasBreakdowns() bool`

HasBreakdowns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



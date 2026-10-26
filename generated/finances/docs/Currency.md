# Currency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyCode** | Pointer to **string** | The three-digit currency code in ISO 4217 format. | [optional] 
**CurrencyAmount** | Pointer to **float32** | A signed decimal number. | [optional] 

## Methods

### NewCurrency

`func NewCurrency() *Currency`

NewCurrency instantiates a new Currency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyWithDefaults

`func NewCurrencyWithDefaults() *Currency`

NewCurrencyWithDefaults instantiates a new Currency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyCode

`func (o *Currency) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *Currency) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *Currency) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *Currency) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetCurrencyAmount

`func (o *Currency) GetCurrencyAmount() float32`

GetCurrencyAmount returns the CurrencyAmount field if non-nil, zero value otherwise.

### GetCurrencyAmountOk

`func (o *Currency) GetCurrencyAmountOk() (*float32, bool)`

GetCurrencyAmountOk returns a tuple with the CurrencyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyAmount

`func (o *Currency) SetCurrencyAmount(v float32)`

SetCurrencyAmount sets CurrencyAmount field to given value.

### HasCurrencyAmount

`func (o *Currency) HasCurrencyAmount() bool`

HasCurrencyAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



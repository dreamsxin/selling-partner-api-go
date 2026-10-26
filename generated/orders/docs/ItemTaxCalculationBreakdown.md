# ItemTaxCalculationBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportingScheme** | Pointer to **string** | The tax reporting scheme applied to this order item.  **Possible values**: - &#x60;UOSS&#x60; (Union one stop shop. The item being purchased is held in the EU for shipment) - &#x60;IOSS&#x60; (Import one stop shop. The item being purchased is not held in the EU for shipment) | [optional] 

## Methods

### NewItemTaxCalculationBreakdown

`func NewItemTaxCalculationBreakdown() *ItemTaxCalculationBreakdown`

NewItemTaxCalculationBreakdown instantiates a new ItemTaxCalculationBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemTaxCalculationBreakdownWithDefaults

`func NewItemTaxCalculationBreakdownWithDefaults() *ItemTaxCalculationBreakdown`

NewItemTaxCalculationBreakdownWithDefaults instantiates a new ItemTaxCalculationBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportingScheme

`func (o *ItemTaxCalculationBreakdown) GetReportingScheme() string`

GetReportingScheme returns the ReportingScheme field if non-nil, zero value otherwise.

### GetReportingSchemeOk

`func (o *ItemTaxCalculationBreakdown) GetReportingSchemeOk() (*string, bool)`

GetReportingSchemeOk returns a tuple with the ReportingScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingScheme

`func (o *ItemTaxCalculationBreakdown) SetReportingScheme(v string)`

SetReportingScheme sets ReportingScheme field to given value.

### HasReportingScheme

`func (o *ItemTaxCalculationBreakdown) HasReportingScheme() bool`

HasReportingScheme returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



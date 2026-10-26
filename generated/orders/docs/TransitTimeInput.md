# TransitTimeInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**TransitTimeInputDetails**](TransitTimeInputDetails.md) |  | [optional] 
**DataSource** | Pointer to **string** | The source of the transit time data.  **Possible values**: - &#x60;AMAZON_AUTOMATED&#x60; (Amazon selected the shipping service using settings generated on the seller&#39;s behalf.) - &#x60;SELLER_SET&#x60; (The seller configured the transit time through a shipping template they maintain.) - &#x60;THIRD_PARTY_INTEGRATOR&#x60; (The transit time came from a live shipping quote that Amazon obtained from a third-party integrator for the seller&#39;s order.) | [optional] 

## Methods

### NewTransitTimeInput

`func NewTransitTimeInput() *TransitTimeInput`

NewTransitTimeInput instantiates a new TransitTimeInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitTimeInputWithDefaults

`func NewTransitTimeInputWithDefaults() *TransitTimeInput`

NewTransitTimeInputWithDefaults instantiates a new TransitTimeInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *TransitTimeInput) GetDetails() TransitTimeInputDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *TransitTimeInput) GetDetailsOk() (*TransitTimeInputDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *TransitTimeInput) SetDetails(v TransitTimeInputDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *TransitTimeInput) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetDataSource

`func (o *TransitTimeInput) GetDataSource() string`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *TransitTimeInput) GetDataSourceOk() (*string, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *TransitTimeInput) SetDataSource(v string)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *TransitTimeInput) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



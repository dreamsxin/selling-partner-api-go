# Summary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartnerMetadata** | Pointer to [**PartnerMetadata**](PartnerMetadata.md) |  | [optional] 
**NetProceeds** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**RelatedIdentifiers** | Pointer to [**[]RelatedIdentifier**](RelatedIdentifier.md) | Related business identifiers of the transaction. | [optional] 
**Breakdowns** | Pointer to [**[]Breakdown**](Breakdown.md) | A list of breakdowns that detail how the total amount is calculated for the transaction. | [optional] 
**PeriodStart** | Pointer to **time.Time** | A date in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date-time format. | [optional] 
**PeriodEnd** | Pointer to **time.Time** | A date in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date-time format. | [optional] 

## Methods

### NewSummary

`func NewSummary() *Summary`

NewSummary instantiates a new Summary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryWithDefaults

`func NewSummaryWithDefaults() *Summary`

NewSummaryWithDefaults instantiates a new Summary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartnerMetadata

`func (o *Summary) GetPartnerMetadata() PartnerMetadata`

GetPartnerMetadata returns the PartnerMetadata field if non-nil, zero value otherwise.

### GetPartnerMetadataOk

`func (o *Summary) GetPartnerMetadataOk() (*PartnerMetadata, bool)`

GetPartnerMetadataOk returns a tuple with the PartnerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerMetadata

`func (o *Summary) SetPartnerMetadata(v PartnerMetadata)`

SetPartnerMetadata sets PartnerMetadata field to given value.

### HasPartnerMetadata

`func (o *Summary) HasPartnerMetadata() bool`

HasPartnerMetadata returns a boolean if a field has been set.

### GetNetProceeds

`func (o *Summary) GetNetProceeds() Currency`

GetNetProceeds returns the NetProceeds field if non-nil, zero value otherwise.

### GetNetProceedsOk

`func (o *Summary) GetNetProceedsOk() (*Currency, bool)`

GetNetProceedsOk returns a tuple with the NetProceeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetProceeds

`func (o *Summary) SetNetProceeds(v Currency)`

SetNetProceeds sets NetProceeds field to given value.

### HasNetProceeds

`func (o *Summary) HasNetProceeds() bool`

HasNetProceeds returns a boolean if a field has been set.

### GetRelatedIdentifiers

`func (o *Summary) GetRelatedIdentifiers() []RelatedIdentifier`

GetRelatedIdentifiers returns the RelatedIdentifiers field if non-nil, zero value otherwise.

### GetRelatedIdentifiersOk

`func (o *Summary) GetRelatedIdentifiersOk() (*[]RelatedIdentifier, bool)`

GetRelatedIdentifiersOk returns a tuple with the RelatedIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedIdentifiers

`func (o *Summary) SetRelatedIdentifiers(v []RelatedIdentifier)`

SetRelatedIdentifiers sets RelatedIdentifiers field to given value.

### HasRelatedIdentifiers

`func (o *Summary) HasRelatedIdentifiers() bool`

HasRelatedIdentifiers returns a boolean if a field has been set.

### GetBreakdowns

`func (o *Summary) GetBreakdowns() []Breakdown`

GetBreakdowns returns the Breakdowns field if non-nil, zero value otherwise.

### GetBreakdownsOk

`func (o *Summary) GetBreakdownsOk() (*[]Breakdown, bool)`

GetBreakdownsOk returns a tuple with the Breakdowns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdowns

`func (o *Summary) SetBreakdowns(v []Breakdown)`

SetBreakdowns sets Breakdowns field to given value.

### HasBreakdowns

`func (o *Summary) HasBreakdowns() bool`

HasBreakdowns returns a boolean if a field has been set.

### GetPeriodStart

`func (o *Summary) GetPeriodStart() time.Time`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *Summary) GetPeriodStartOk() (*time.Time, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *Summary) SetPeriodStart(v time.Time)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *Summary) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *Summary) GetPeriodEnd() time.Time`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *Summary) GetPeriodEndOk() (*time.Time, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *Summary) SetPeriodEnd(v time.Time)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *Summary) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SummaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summaries** | Pointer to [**[]Summary**](Summary.md) | A list of financial summaries. | [optional] 
**NextToken** | Pointer to **string** | A token to retrieve the next page of results. If null, no more results are available. | [optional] 

## Methods

### NewSummaryResponse

`func NewSummaryResponse() *SummaryResponse`

NewSummaryResponse instantiates a new SummaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryResponseWithDefaults

`func NewSummaryResponseWithDefaults() *SummaryResponse`

NewSummaryResponseWithDefaults instantiates a new SummaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummaries

`func (o *SummaryResponse) GetSummaries() []Summary`

GetSummaries returns the Summaries field if non-nil, zero value otherwise.

### GetSummariesOk

`func (o *SummaryResponse) GetSummariesOk() (*[]Summary, bool)`

GetSummariesOk returns a tuple with the Summaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaries

`func (o *SummaryResponse) SetSummaries(v []Summary)`

SetSummaries sets Summaries field to given value.

### HasSummaries

`func (o *SummaryResponse) HasSummaries() bool`

HasSummaries returns a boolean if a field has been set.

### GetNextToken

`func (o *SummaryResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *SummaryResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *SummaryResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *SummaryResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



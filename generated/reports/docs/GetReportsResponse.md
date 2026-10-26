# GetReportsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reports** | [**[]Report**](Report.md) | A list of reports. | 
**NextToken** | Pointer to **string** | Returned when the number of results exceeds &#x60;pageSize&#x60;. To get the next page of results, call &#x60;getReports&#x60; with this token as the only parameter. | [optional] 

## Methods

### NewGetReportsResponse

`func NewGetReportsResponse(reports []Report, ) *GetReportsResponse`

NewGetReportsResponse instantiates a new GetReportsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReportsResponseWithDefaults

`func NewGetReportsResponseWithDefaults() *GetReportsResponse`

NewGetReportsResponseWithDefaults instantiates a new GetReportsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReports

`func (o *GetReportsResponse) GetReports() []Report`

GetReports returns the Reports field if non-nil, zero value otherwise.

### GetReportsOk

`func (o *GetReportsResponse) GetReportsOk() (*[]Report, bool)`

GetReportsOk returns a tuple with the Reports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReports

`func (o *GetReportsResponse) SetReports(v []Report)`

SetReports sets Reports field to given value.


### GetNextToken

`func (o *GetReportsResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *GetReportsResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *GetReportsResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *GetReportsResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



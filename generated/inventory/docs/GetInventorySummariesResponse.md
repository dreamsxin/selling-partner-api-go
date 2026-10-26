# GetInventorySummariesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payload** | Pointer to [**GetInventorySummariesResult**](GetInventorySummariesResult.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Errors** | Pointer to [**[]Error**](Error.md) | A list of error responses returned when a request is unsuccessful. | [optional] 

## Methods

### NewGetInventorySummariesResponse

`func NewGetInventorySummariesResponse() *GetInventorySummariesResponse`

NewGetInventorySummariesResponse instantiates a new GetInventorySummariesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInventorySummariesResponseWithDefaults

`func NewGetInventorySummariesResponseWithDefaults() *GetInventorySummariesResponse`

NewGetInventorySummariesResponseWithDefaults instantiates a new GetInventorySummariesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayload

`func (o *GetInventorySummariesResponse) GetPayload() GetInventorySummariesResult`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *GetInventorySummariesResponse) GetPayloadOk() (*GetInventorySummariesResult, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *GetInventorySummariesResponse) SetPayload(v GetInventorySummariesResult)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *GetInventorySummariesResponse) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetPagination

`func (o *GetInventorySummariesResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetInventorySummariesResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetInventorySummariesResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetInventorySummariesResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetErrors

`func (o *GetInventorySummariesResponse) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GetInventorySummariesResponse) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GetInventorySummariesResponse) SetErrors(v []Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GetInventorySummariesResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



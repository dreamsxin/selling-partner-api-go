# Report

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketplaceIds** | Pointer to **[]string** | A list of marketplace identifiers for the report. | [optional] 
**ReportId** | **string** | The identifier for the report. This identifier is unique only in combination with a seller ID. | 
**ReportType** | **string** | The report type. Refer to [Report Type Values](https://developer-docs.amazon.com/sp-api/docs/report-type-values) for more information. | 
**DataStartTime** | Pointer to **time.Time** | The start of a date and time range used for selecting the data to report. | [optional] 
**DataEndTime** | Pointer to **time.Time** | The end of a date and time range used for selecting the data to report. | [optional] 
**ReportScheduleId** | Pointer to **string** | The identifier of the report schedule that created this report (if any). This identifier is unique only in combination with a seller ID. | [optional] 
**CreatedTime** | **time.Time** | The date and time when the report was created. | 
**ProcessingStatus** | **string** | The processing status of the report. | 
**ProcessingStartTime** | Pointer to **time.Time** | The date and time when the report processing started, in &lt;a href&#x3D;&#39;https://developer-docs.amazon.com/sp-api/docs/iso-8601&#39;&gt;ISO 8601&lt;/a&gt; date time format. | [optional] 
**ProcessingEndTime** | Pointer to **time.Time** | The date and time when the report processing completed, in &lt;a href&#x3D;&#39;https://developer-docs.amazon.com/sp-api/docs/iso-8601&#39;&gt;ISO 8601&lt;/a&gt; date time format. | [optional] 
**ReportDocumentId** | Pointer to **string** | The identifier for the report document. Pass this into the &#x60;getReportDocument&#x60; operation to get the information you will need to retrieve the report document&#39;s contents. | [optional] 

## Methods

### NewReport

`func NewReport(reportId string, reportType string, createdTime time.Time, processingStatus string, ) *Report`

NewReport instantiates a new Report object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportWithDefaults

`func NewReportWithDefaults() *Report`

NewReportWithDefaults instantiates a new Report object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketplaceIds

`func (o *Report) GetMarketplaceIds() []string`

GetMarketplaceIds returns the MarketplaceIds field if non-nil, zero value otherwise.

### GetMarketplaceIdsOk

`func (o *Report) GetMarketplaceIdsOk() (*[]string, bool)`

GetMarketplaceIdsOk returns a tuple with the MarketplaceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceIds

`func (o *Report) SetMarketplaceIds(v []string)`

SetMarketplaceIds sets MarketplaceIds field to given value.

### HasMarketplaceIds

`func (o *Report) HasMarketplaceIds() bool`

HasMarketplaceIds returns a boolean if a field has been set.

### GetReportId

`func (o *Report) GetReportId() string`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *Report) GetReportIdOk() (*string, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *Report) SetReportId(v string)`

SetReportId sets ReportId field to given value.


### GetReportType

`func (o *Report) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *Report) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *Report) SetReportType(v string)`

SetReportType sets ReportType field to given value.


### GetDataStartTime

`func (o *Report) GetDataStartTime() time.Time`

GetDataStartTime returns the DataStartTime field if non-nil, zero value otherwise.

### GetDataStartTimeOk

`func (o *Report) GetDataStartTimeOk() (*time.Time, bool)`

GetDataStartTimeOk returns a tuple with the DataStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStartTime

`func (o *Report) SetDataStartTime(v time.Time)`

SetDataStartTime sets DataStartTime field to given value.

### HasDataStartTime

`func (o *Report) HasDataStartTime() bool`

HasDataStartTime returns a boolean if a field has been set.

### GetDataEndTime

`func (o *Report) GetDataEndTime() time.Time`

GetDataEndTime returns the DataEndTime field if non-nil, zero value otherwise.

### GetDataEndTimeOk

`func (o *Report) GetDataEndTimeOk() (*time.Time, bool)`

GetDataEndTimeOk returns a tuple with the DataEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataEndTime

`func (o *Report) SetDataEndTime(v time.Time)`

SetDataEndTime sets DataEndTime field to given value.

### HasDataEndTime

`func (o *Report) HasDataEndTime() bool`

HasDataEndTime returns a boolean if a field has been set.

### GetReportScheduleId

`func (o *Report) GetReportScheduleId() string`

GetReportScheduleId returns the ReportScheduleId field if non-nil, zero value otherwise.

### GetReportScheduleIdOk

`func (o *Report) GetReportScheduleIdOk() (*string, bool)`

GetReportScheduleIdOk returns a tuple with the ReportScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportScheduleId

`func (o *Report) SetReportScheduleId(v string)`

SetReportScheduleId sets ReportScheduleId field to given value.

### HasReportScheduleId

`func (o *Report) HasReportScheduleId() bool`

HasReportScheduleId returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Report) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Report) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Report) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetProcessingStatus

`func (o *Report) GetProcessingStatus() string`

GetProcessingStatus returns the ProcessingStatus field if non-nil, zero value otherwise.

### GetProcessingStatusOk

`func (o *Report) GetProcessingStatusOk() (*string, bool)`

GetProcessingStatusOk returns a tuple with the ProcessingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingStatus

`func (o *Report) SetProcessingStatus(v string)`

SetProcessingStatus sets ProcessingStatus field to given value.


### GetProcessingStartTime

`func (o *Report) GetProcessingStartTime() time.Time`

GetProcessingStartTime returns the ProcessingStartTime field if non-nil, zero value otherwise.

### GetProcessingStartTimeOk

`func (o *Report) GetProcessingStartTimeOk() (*time.Time, bool)`

GetProcessingStartTimeOk returns a tuple with the ProcessingStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingStartTime

`func (o *Report) SetProcessingStartTime(v time.Time)`

SetProcessingStartTime sets ProcessingStartTime field to given value.

### HasProcessingStartTime

`func (o *Report) HasProcessingStartTime() bool`

HasProcessingStartTime returns a boolean if a field has been set.

### GetProcessingEndTime

`func (o *Report) GetProcessingEndTime() time.Time`

GetProcessingEndTime returns the ProcessingEndTime field if non-nil, zero value otherwise.

### GetProcessingEndTimeOk

`func (o *Report) GetProcessingEndTimeOk() (*time.Time, bool)`

GetProcessingEndTimeOk returns a tuple with the ProcessingEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingEndTime

`func (o *Report) SetProcessingEndTime(v time.Time)`

SetProcessingEndTime sets ProcessingEndTime field to given value.

### HasProcessingEndTime

`func (o *Report) HasProcessingEndTime() bool`

HasProcessingEndTime returns a boolean if a field has been set.

### GetReportDocumentId

`func (o *Report) GetReportDocumentId() string`

GetReportDocumentId returns the ReportDocumentId field if non-nil, zero value otherwise.

### GetReportDocumentIdOk

`func (o *Report) GetReportDocumentIdOk() (*string, bool)`

GetReportDocumentIdOk returns a tuple with the ReportDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportDocumentId

`func (o *Report) SetReportDocumentId(v string)`

SetReportDocumentId sets ReportDocumentId field to given value.

### HasReportDocumentId

`func (o *Report) HasReportDocumentId() bool`

HasReportDocumentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



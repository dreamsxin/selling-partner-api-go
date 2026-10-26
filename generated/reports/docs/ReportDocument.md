# ReportDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportDocumentId** | **string** | The identifier for the report document. This identifier is unique only in combination with a seller ID. | 
**Url** | **string** | A presigned URL for the report document. If &#x60;compressionAlgorithm&#x60; is not returned, you can download the report directly from this URL. This URL expires after 5 minutes. | 
**CompressionAlgorithm** | Pointer to **string** | If the report document contents have been compressed, the compression algorithm used is returned in this property and you must decompress the report when you download. Otherwise, you can download the report directly. Refer to [Step 2. Download the report](https://developer-docs.amazon.com/sp-api/docs/reports-api-v2021-06-30-retrieve-a-report#step-2-download-the-report) in the use case guide, where sample code is provided. | [optional] 

## Methods

### NewReportDocument

`func NewReportDocument(reportDocumentId string, url string, ) *ReportDocument`

NewReportDocument instantiates a new ReportDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportDocumentWithDefaults

`func NewReportDocumentWithDefaults() *ReportDocument`

NewReportDocumentWithDefaults instantiates a new ReportDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportDocumentId

`func (o *ReportDocument) GetReportDocumentId() string`

GetReportDocumentId returns the ReportDocumentId field if non-nil, zero value otherwise.

### GetReportDocumentIdOk

`func (o *ReportDocument) GetReportDocumentIdOk() (*string, bool)`

GetReportDocumentIdOk returns a tuple with the ReportDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportDocumentId

`func (o *ReportDocument) SetReportDocumentId(v string)`

SetReportDocumentId sets ReportDocumentId field to given value.


### GetUrl

`func (o *ReportDocument) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ReportDocument) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ReportDocument) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCompressionAlgorithm

`func (o *ReportDocument) GetCompressionAlgorithm() string`

GetCompressionAlgorithm returns the CompressionAlgorithm field if non-nil, zero value otherwise.

### GetCompressionAlgorithmOk

`func (o *ReportDocument) GetCompressionAlgorithmOk() (*string, bool)`

GetCompressionAlgorithmOk returns a tuple with the CompressionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionAlgorithm

`func (o *ReportDocument) SetCompressionAlgorithm(v string)`

SetCompressionAlgorithm sets CompressionAlgorithm field to given value.

### HasCompressionAlgorithm

`func (o *ReportDocument) HasCompressionAlgorithm() bool`

HasCompressionAlgorithm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



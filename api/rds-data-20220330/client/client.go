// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type BatchExecuteStatementRequest struct {
	Database      *string                                        `json:"database,omitempty" xml:"database,omitempty"`
	ParameterSets [][]*BatchExecuteStatementRequestParameterSets `json:"parameterSets,omitempty" xml:"parameterSets,omitempty" type:"Repeated"`
	ResourceArn   *string                                        `json:"resourceArn,omitempty" xml:"resourceArn,omitempty"`
	SecretArn     *string                                        `json:"secretArn,omitempty" xml:"secretArn,omitempty"`
	Sql           *string                                        `json:"sql,omitempty" xml:"sql,omitempty"`
	TransactionId *string                                        `json:"transactionId,omitempty" xml:"transactionId,omitempty"`
}

func (s BatchExecuteStatementRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchExecuteStatementRequest) GoString() string {
	return s.String()
}

func (s *BatchExecuteStatementRequest) SetDatabase(v string) *BatchExecuteStatementRequest {
	s.Database = &v
	return s
}

func (s *BatchExecuteStatementRequest) SetParameterSets(v [][]*BatchExecuteStatementRequestParameterSets) *BatchExecuteStatementRequest {
	s.ParameterSets = v
	return s
}

func (s *BatchExecuteStatementRequest) SetResourceArn(v string) *BatchExecuteStatementRequest {
	s.ResourceArn = &v
	return s
}

func (s *BatchExecuteStatementRequest) SetSecretArn(v string) *BatchExecuteStatementRequest {
	s.SecretArn = &v
	return s
}

func (s *BatchExecuteStatementRequest) SetSql(v string) *BatchExecuteStatementRequest {
	s.Sql = &v
	return s
}

func (s *BatchExecuteStatementRequest) SetTransactionId(v string) *BatchExecuteStatementRequest {
	s.TransactionId = &v
	return s
}

type BatchExecuteStatementRequestParameterSets struct {
	Name     *string                                         `json:"name,omitempty" xml:"name,omitempty"`
	TypeHint *string                                         `json:"typeHint,omitempty" xml:"typeHint,omitempty"`
	Value    *BatchExecuteStatementRequestParameterSetsValue `json:"value,omitempty" xml:"value,omitempty" type:"Struct"`
}

func (s BatchExecuteStatementRequestParameterSets) String() string {
	return tea.Prettify(s)
}

func (s BatchExecuteStatementRequestParameterSets) GoString() string {
	return s.String()
}

func (s *BatchExecuteStatementRequestParameterSets) SetName(v string) *BatchExecuteStatementRequestParameterSets {
	s.Name = &v
	return s
}

func (s *BatchExecuteStatementRequestParameterSets) SetTypeHint(v string) *BatchExecuteStatementRequestParameterSets {
	s.TypeHint = &v
	return s
}

func (s *BatchExecuteStatementRequestParameterSets) SetValue(v *BatchExecuteStatementRequestParameterSetsValue) *BatchExecuteStatementRequestParameterSets {
	s.Value = v
	return s
}

type BatchExecuteStatementRequestParameterSetsValue struct {
	ArrayValue   *string   `json:"arrayValue,omitempty" xml:"arrayValue,omitempty"`
	BlobValue    io.Reader `json:"blobValue,omitempty" xml:"blobValue,omitempty"`
	BooleanValue *bool     `json:"booleanValue,omitempty" xml:"booleanValue,omitempty"`
	DoubleValue  *float64  `json:"doubleValue,omitempty" xml:"doubleValue,omitempty"`
	IsNull       *bool     `json:"isNull,omitempty" xml:"isNull,omitempty"`
	LongValue    *int64    `json:"longValue,omitempty" xml:"longValue,omitempty"`
	StringValue  *string   `json:"stringValue,omitempty" xml:"stringValue,omitempty"`
}

func (s BatchExecuteStatementRequestParameterSetsValue) String() string {
	return tea.Prettify(s)
}

func (s BatchExecuteStatementRequestParameterSetsValue) GoString() string {
	return s.String()
}

func (s *BatchExecuteStatementRequestParameterSetsValue) SetArrayValue(v string) *BatchExecuteStatementRequestParameterSetsValue {
	s.ArrayValue = &v
	return s
}

func (s *BatchExecuteStatementRequestParameterSetsValue) SetBlobValue(v io.Reader) *BatchExecuteStatementRequestParameterSetsValue {
	s.BlobValue = v
	return s
}

func (s *BatchExecuteStatementRequestParameterSetsValue) SetBooleanValue(v bool) *BatchExecuteStatementRequestParameterSetsValue {
	s.BooleanValue = &v
	return s
}

func (s *BatchExecuteStatementRequestParameterSetsValue) SetDoubleValue(v float64) *BatchExecuteStatementRequestParameterSetsValue {
	s.DoubleValue = &v
	return s
}

func (s *BatchExecuteStatementRequestParameterSetsValue) SetIsNull(v bool) *BatchExecuteStatementRequestParameterSetsValue {
	s.IsNull = &v
	return s
}

func (s *BatchExecuteStatementRequestParameterSetsValue) SetLongValue(v int64) *BatchExecuteStatementRequestParameterSetsValue {
	s.LongValue = &v
	return s
}

func (s *BatchExecuteStatementRequestParameterSetsValue) SetStringValue(v string) *BatchExecuteStatementRequestParameterSetsValue {
	s.StringValue = &v
	return s
}

type BatchExecuteStatementShrinkRequest struct {
	Database            *string `json:"database,omitempty" xml:"database,omitempty"`
	ParameterSetsShrink *string `json:"parameterSets,omitempty" xml:"parameterSets,omitempty"`
	ResourceArn         *string `json:"resourceArn,omitempty" xml:"resourceArn,omitempty"`
	SecretArn           *string `json:"secretArn,omitempty" xml:"secretArn,omitempty"`
	Sql                 *string `json:"sql,omitempty" xml:"sql,omitempty"`
	TransactionId       *string `json:"transactionId,omitempty" xml:"transactionId,omitempty"`
}

func (s BatchExecuteStatementShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchExecuteStatementShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchExecuteStatementShrinkRequest) SetDatabase(v string) *BatchExecuteStatementShrinkRequest {
	s.Database = &v
	return s
}

func (s *BatchExecuteStatementShrinkRequest) SetParameterSetsShrink(v string) *BatchExecuteStatementShrinkRequest {
	s.ParameterSetsShrink = &v
	return s
}

func (s *BatchExecuteStatementShrinkRequest) SetResourceArn(v string) *BatchExecuteStatementShrinkRequest {
	s.ResourceArn = &v
	return s
}

func (s *BatchExecuteStatementShrinkRequest) SetSecretArn(v string) *BatchExecuteStatementShrinkRequest {
	s.SecretArn = &v
	return s
}

func (s *BatchExecuteStatementShrinkRequest) SetSql(v string) *BatchExecuteStatementShrinkRequest {
	s.Sql = &v
	return s
}

func (s *BatchExecuteStatementShrinkRequest) SetTransactionId(v string) *BatchExecuteStatementShrinkRequest {
	s.TransactionId = &v
	return s
}

type BatchExecuteStatementResponseBody struct {
	Code      *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Data      *BatchExecuteStatementResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message   *string                                `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchExecuteStatementResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchExecuteStatementResponseBody) GoString() string {
	return s.String()
}

func (s *BatchExecuteStatementResponseBody) SetCode(v string) *BatchExecuteStatementResponseBody {
	s.Code = &v
	return s
}

func (s *BatchExecuteStatementResponseBody) SetData(v *BatchExecuteStatementResponseBodyData) *BatchExecuteStatementResponseBody {
	s.Data = v
	return s
}

func (s *BatchExecuteStatementResponseBody) SetMessage(v string) *BatchExecuteStatementResponseBody {
	s.Message = &v
	return s
}

func (s *BatchExecuteStatementResponseBody) SetRequestId(v string) *BatchExecuteStatementResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchExecuteStatementResponseBody) SetSuccess(v bool) *BatchExecuteStatementResponseBody {
	s.Success = &v
	return s
}

type BatchExecuteStatementResponseBodyData struct {
	UpdateResults []*BatchExecuteStatementResponseBodyDataUpdateResults `json:"updateResults,omitempty" xml:"updateResults,omitempty" type:"Repeated"`
}

func (s BatchExecuteStatementResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BatchExecuteStatementResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchExecuteStatementResponseBodyData) SetUpdateResults(v []*BatchExecuteStatementResponseBodyDataUpdateResults) *BatchExecuteStatementResponseBodyData {
	s.UpdateResults = v
	return s
}

type BatchExecuteStatementResponseBodyDataUpdateResults struct {
	GeneratedFields []*BatchExecuteStatementResponseBodyDataUpdateResultsGeneratedFields `json:"generatedFields,omitempty" xml:"generatedFields,omitempty" type:"Repeated"`
}

func (s BatchExecuteStatementResponseBodyDataUpdateResults) String() string {
	return tea.Prettify(s)
}

func (s BatchExecuteStatementResponseBodyDataUpdateResults) GoString() string {
	return s.String()
}

func (s *BatchExecuteStatementResponseBodyDataUpdateResults) SetGeneratedFields(v []*BatchExecuteStatementResponseBodyDataUpdateResultsGeneratedFields) *BatchExecuteStatementResponseBodyDataUpdateResults {
	s.GeneratedFields = v
	return s
}

type BatchExecuteStatementResponseBodyDataUpdateResultsGeneratedFields struct {
	ArrayValue   *string   `json:"arrayValue,omitempty" xml:"arrayValue,omitempty"`
	BlobValue    io.Reader `json:"blobValue,omitempty" xml:"blobValue,omitempty"`
	BooleanValue *bool     `json:"booleanValue,omitempty" xml:"booleanValue,omitempty"`
	DoubleValue  *float64  `json:"doubleValue,omitempty" xml:"doubleValue,omitempty"`
	IsNull       *bool     `json:"isNull,omitempty" xml:"isNull,omitempty"`
	LongValue    *int64    `json:"longValue,omitempty" xml:"longValue,omitempty"`
	StringValue  *string   `json:"stringValue,omitempty" xml:"stringValue,omitempty"`
}

func (s BatchExecuteStatementResponseBodyDataUpdateResultsGeneratedFields) String() string {
	return tea.Prettify(s)
}

func (s BatchExecuteStatementResponseBodyDataUpdateResultsGeneratedFields) GoString() string {
	return s.String()
}

func (s *BatchExecuteStatementResponseBodyDataUpdateResultsGeneratedFields) SetArrayValue(v string) *BatchExecuteStatementResponseBodyDataUpdateResultsGeneratedFields {
	s.ArrayValue = &v
	return s
}

func (s *BatchExecuteStatementResponseBodyDataUpdateResultsGeneratedFields) SetBlobValue(v io.Reader) *BatchExecuteStatementResponseBodyDataUpdateResultsGeneratedFields {
	s.BlobValue = v
	return s
}

func (s *BatchExecuteStatementResponseBodyDataUpdateResultsGeneratedFields) SetBooleanValue(v bool) *BatchExecuteStatementResponseBodyDataUpdateResultsGeneratedFields {
	s.BooleanValue = &v
	return s
}

func (s *BatchExecuteStatementResponseBodyDataUpdateResultsGeneratedFields) SetDoubleValue(v float64) *BatchExecuteStatementResponseBodyDataUpdateResultsGeneratedFields {
	s.DoubleValue = &v
	return s
}

func (s *BatchExecuteStatementResponseBodyDataUpdateResultsGeneratedFields) SetIsNull(v bool) *BatchExecuteStatementResponseBodyDataUpdateResultsGeneratedFields {
	s.IsNull = &v
	return s
}

func (s *BatchExecuteStatementResponseBodyDataUpdateResultsGeneratedFields) SetLongValue(v int64) *BatchExecuteStatementResponseBodyDataUpdateResultsGeneratedFields {
	s.LongValue = &v
	return s
}

func (s *BatchExecuteStatementResponseBodyDataUpdateResultsGeneratedFields) SetStringValue(v string) *BatchExecuteStatementResponseBodyDataUpdateResultsGeneratedFields {
	s.StringValue = &v
	return s
}

type BatchExecuteStatementResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchExecuteStatementResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchExecuteStatementResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchExecuteStatementResponse) GoString() string {
	return s.String()
}

func (s *BatchExecuteStatementResponse) SetHeaders(v map[string]*string) *BatchExecuteStatementResponse {
	s.Headers = v
	return s
}

func (s *BatchExecuteStatementResponse) SetStatusCode(v int32) *BatchExecuteStatementResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchExecuteStatementResponse) SetBody(v *BatchExecuteStatementResponseBody) *BatchExecuteStatementResponse {
	s.Body = v
	return s
}

type BeginTransactionRequest struct {
	Database    *string `json:"database,omitempty" xml:"database,omitempty"`
	ResourceArn *string `json:"resourceArn,omitempty" xml:"resourceArn,omitempty"`
	SecretArn   *string `json:"secretArn,omitempty" xml:"secretArn,omitempty"`
}

func (s BeginTransactionRequest) String() string {
	return tea.Prettify(s)
}

func (s BeginTransactionRequest) GoString() string {
	return s.String()
}

func (s *BeginTransactionRequest) SetDatabase(v string) *BeginTransactionRequest {
	s.Database = &v
	return s
}

func (s *BeginTransactionRequest) SetResourceArn(v string) *BeginTransactionRequest {
	s.ResourceArn = &v
	return s
}

func (s *BeginTransactionRequest) SetSecretArn(v string) *BeginTransactionRequest {
	s.SecretArn = &v
	return s
}

type BeginTransactionResponseBody struct {
	Code      *string                           `json:"code,omitempty" xml:"code,omitempty"`
	Data      *BeginTransactionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message   *string                           `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BeginTransactionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BeginTransactionResponseBody) GoString() string {
	return s.String()
}

func (s *BeginTransactionResponseBody) SetCode(v string) *BeginTransactionResponseBody {
	s.Code = &v
	return s
}

func (s *BeginTransactionResponseBody) SetData(v *BeginTransactionResponseBodyData) *BeginTransactionResponseBody {
	s.Data = v
	return s
}

func (s *BeginTransactionResponseBody) SetMessage(v string) *BeginTransactionResponseBody {
	s.Message = &v
	return s
}

func (s *BeginTransactionResponseBody) SetRequestId(v string) *BeginTransactionResponseBody {
	s.RequestId = &v
	return s
}

func (s *BeginTransactionResponseBody) SetSuccess(v bool) *BeginTransactionResponseBody {
	s.Success = &v
	return s
}

type BeginTransactionResponseBodyData struct {
	TransactionId *string `json:"transactionId,omitempty" xml:"transactionId,omitempty"`
}

func (s BeginTransactionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BeginTransactionResponseBodyData) GoString() string {
	return s.String()
}

func (s *BeginTransactionResponseBodyData) SetTransactionId(v string) *BeginTransactionResponseBodyData {
	s.TransactionId = &v
	return s
}

type BeginTransactionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BeginTransactionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BeginTransactionResponse) String() string {
	return tea.Prettify(s)
}

func (s BeginTransactionResponse) GoString() string {
	return s.String()
}

func (s *BeginTransactionResponse) SetHeaders(v map[string]*string) *BeginTransactionResponse {
	s.Headers = v
	return s
}

func (s *BeginTransactionResponse) SetStatusCode(v int32) *BeginTransactionResponse {
	s.StatusCode = &v
	return s
}

func (s *BeginTransactionResponse) SetBody(v *BeginTransactionResponseBody) *BeginTransactionResponse {
	s.Body = v
	return s
}

type CommitTransactionRequest struct {
	Database      *string `json:"database,omitempty" xml:"database,omitempty"`
	ResourceArn   *string `json:"resourceArn,omitempty" xml:"resourceArn,omitempty"`
	SecretArn     *string `json:"secretArn,omitempty" xml:"secretArn,omitempty"`
	TransactionId *string `json:"transactionId,omitempty" xml:"transactionId,omitempty"`
}

func (s CommitTransactionRequest) String() string {
	return tea.Prettify(s)
}

func (s CommitTransactionRequest) GoString() string {
	return s.String()
}

func (s *CommitTransactionRequest) SetDatabase(v string) *CommitTransactionRequest {
	s.Database = &v
	return s
}

func (s *CommitTransactionRequest) SetResourceArn(v string) *CommitTransactionRequest {
	s.ResourceArn = &v
	return s
}

func (s *CommitTransactionRequest) SetSecretArn(v string) *CommitTransactionRequest {
	s.SecretArn = &v
	return s
}

func (s *CommitTransactionRequest) SetTransactionId(v string) *CommitTransactionRequest {
	s.TransactionId = &v
	return s
}

type CommitTransactionResponseBody struct {
	Code      *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Data      *CommitTransactionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message   *string                            `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                              `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CommitTransactionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CommitTransactionResponseBody) GoString() string {
	return s.String()
}

func (s *CommitTransactionResponseBody) SetCode(v string) *CommitTransactionResponseBody {
	s.Code = &v
	return s
}

func (s *CommitTransactionResponseBody) SetData(v *CommitTransactionResponseBodyData) *CommitTransactionResponseBody {
	s.Data = v
	return s
}

func (s *CommitTransactionResponseBody) SetMessage(v string) *CommitTransactionResponseBody {
	s.Message = &v
	return s
}

func (s *CommitTransactionResponseBody) SetRequestId(v string) *CommitTransactionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CommitTransactionResponseBody) SetSuccess(v bool) *CommitTransactionResponseBody {
	s.Success = &v
	return s
}

type CommitTransactionResponseBodyData struct {
	TransactionStatus *string `json:"transactionStatus,omitempty" xml:"transactionStatus,omitempty"`
}

func (s CommitTransactionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CommitTransactionResponseBodyData) GoString() string {
	return s.String()
}

func (s *CommitTransactionResponseBodyData) SetTransactionStatus(v string) *CommitTransactionResponseBodyData {
	s.TransactionStatus = &v
	return s
}

type CommitTransactionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CommitTransactionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CommitTransactionResponse) String() string {
	return tea.Prettify(s)
}

func (s CommitTransactionResponse) GoString() string {
	return s.String()
}

func (s *CommitTransactionResponse) SetHeaders(v map[string]*string) *CommitTransactionResponse {
	s.Headers = v
	return s
}

func (s *CommitTransactionResponse) SetStatusCode(v int32) *CommitTransactionResponse {
	s.StatusCode = &v
	return s
}

func (s *CommitTransactionResponse) SetBody(v *CommitTransactionResponseBody) *CommitTransactionResponse {
	s.Body = v
	return s
}

type ExecuteStatementRequest struct {
	ContinueAfterTimeout  *bool                                    `json:"continueAfterTimeout,omitempty" xml:"continueAfterTimeout,omitempty"`
	Database              *string                                  `json:"database,omitempty" xml:"database,omitempty"`
	FormatRecordsAs       *string                                  `json:"formatRecordsAs,omitempty" xml:"formatRecordsAs,omitempty"`
	IncludeResultMetadata *string                                  `json:"includeResultMetadata,omitempty" xml:"includeResultMetadata,omitempty"`
	Parameters            []*ExecuteStatementRequestParameters     `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Repeated"`
	ResourceArn           *string                                  `json:"resourceArn,omitempty" xml:"resourceArn,omitempty"`
	ResultSetOptions      *ExecuteStatementRequestResultSetOptions `json:"resultSetOptions,omitempty" xml:"resultSetOptions,omitempty" type:"Struct"`
	SecretArn             *string                                  `json:"secretArn,omitempty" xml:"secretArn,omitempty"`
	Sql                   *string                                  `json:"sql,omitempty" xml:"sql,omitempty"`
	TransactionId         *string                                  `json:"transactionId,omitempty" xml:"transactionId,omitempty"`
}

func (s ExecuteStatementRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteStatementRequest) GoString() string {
	return s.String()
}

func (s *ExecuteStatementRequest) SetContinueAfterTimeout(v bool) *ExecuteStatementRequest {
	s.ContinueAfterTimeout = &v
	return s
}

func (s *ExecuteStatementRequest) SetDatabase(v string) *ExecuteStatementRequest {
	s.Database = &v
	return s
}

func (s *ExecuteStatementRequest) SetFormatRecordsAs(v string) *ExecuteStatementRequest {
	s.FormatRecordsAs = &v
	return s
}

func (s *ExecuteStatementRequest) SetIncludeResultMetadata(v string) *ExecuteStatementRequest {
	s.IncludeResultMetadata = &v
	return s
}

func (s *ExecuteStatementRequest) SetParameters(v []*ExecuteStatementRequestParameters) *ExecuteStatementRequest {
	s.Parameters = v
	return s
}

func (s *ExecuteStatementRequest) SetResourceArn(v string) *ExecuteStatementRequest {
	s.ResourceArn = &v
	return s
}

func (s *ExecuteStatementRequest) SetResultSetOptions(v *ExecuteStatementRequestResultSetOptions) *ExecuteStatementRequest {
	s.ResultSetOptions = v
	return s
}

func (s *ExecuteStatementRequest) SetSecretArn(v string) *ExecuteStatementRequest {
	s.SecretArn = &v
	return s
}

func (s *ExecuteStatementRequest) SetSql(v string) *ExecuteStatementRequest {
	s.Sql = &v
	return s
}

func (s *ExecuteStatementRequest) SetTransactionId(v string) *ExecuteStatementRequest {
	s.TransactionId = &v
	return s
}

type ExecuteStatementRequestParameters struct {
	Name     *string                                 `json:"name,omitempty" xml:"name,omitempty"`
	TypeHint *string                                 `json:"typeHint,omitempty" xml:"typeHint,omitempty"`
	Value    *ExecuteStatementRequestParametersValue `json:"value,omitempty" xml:"value,omitempty" type:"Struct"`
}

func (s ExecuteStatementRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s ExecuteStatementRequestParameters) GoString() string {
	return s.String()
}

func (s *ExecuteStatementRequestParameters) SetName(v string) *ExecuteStatementRequestParameters {
	s.Name = &v
	return s
}

func (s *ExecuteStatementRequestParameters) SetTypeHint(v string) *ExecuteStatementRequestParameters {
	s.TypeHint = &v
	return s
}

func (s *ExecuteStatementRequestParameters) SetValue(v *ExecuteStatementRequestParametersValue) *ExecuteStatementRequestParameters {
	s.Value = v
	return s
}

type ExecuteStatementRequestParametersValue struct {
	ArrayValue   *string   `json:"arrayValue,omitempty" xml:"arrayValue,omitempty"`
	BlobValue    io.Reader `json:"blobValue,omitempty" xml:"blobValue,omitempty"`
	BooleanValue *bool     `json:"booleanValue,omitempty" xml:"booleanValue,omitempty"`
	DoubleValue  *float64  `json:"doubleValue,omitempty" xml:"doubleValue,omitempty"`
	IsNull       *bool     `json:"isNull,omitempty" xml:"isNull,omitempty"`
	LongValue    *int64    `json:"longValue,omitempty" xml:"longValue,omitempty"`
	StringValue  *string   `json:"stringValue,omitempty" xml:"stringValue,omitempty"`
}

func (s ExecuteStatementRequestParametersValue) String() string {
	return tea.Prettify(s)
}

func (s ExecuteStatementRequestParametersValue) GoString() string {
	return s.String()
}

func (s *ExecuteStatementRequestParametersValue) SetArrayValue(v string) *ExecuteStatementRequestParametersValue {
	s.ArrayValue = &v
	return s
}

func (s *ExecuteStatementRequestParametersValue) SetBlobValue(v io.Reader) *ExecuteStatementRequestParametersValue {
	s.BlobValue = v
	return s
}

func (s *ExecuteStatementRequestParametersValue) SetBooleanValue(v bool) *ExecuteStatementRequestParametersValue {
	s.BooleanValue = &v
	return s
}

func (s *ExecuteStatementRequestParametersValue) SetDoubleValue(v float64) *ExecuteStatementRequestParametersValue {
	s.DoubleValue = &v
	return s
}

func (s *ExecuteStatementRequestParametersValue) SetIsNull(v bool) *ExecuteStatementRequestParametersValue {
	s.IsNull = &v
	return s
}

func (s *ExecuteStatementRequestParametersValue) SetLongValue(v int64) *ExecuteStatementRequestParametersValue {
	s.LongValue = &v
	return s
}

func (s *ExecuteStatementRequestParametersValue) SetStringValue(v string) *ExecuteStatementRequestParametersValue {
	s.StringValue = &v
	return s
}

type ExecuteStatementRequestResultSetOptions struct {
	DecimalReturnType *string `json:"decimalReturnType,omitempty" xml:"decimalReturnType,omitempty"`
	LongReturnType    *string `json:"longReturnType,omitempty" xml:"longReturnType,omitempty"`
}

func (s ExecuteStatementRequestResultSetOptions) String() string {
	return tea.Prettify(s)
}

func (s ExecuteStatementRequestResultSetOptions) GoString() string {
	return s.String()
}

func (s *ExecuteStatementRequestResultSetOptions) SetDecimalReturnType(v string) *ExecuteStatementRequestResultSetOptions {
	s.DecimalReturnType = &v
	return s
}

func (s *ExecuteStatementRequestResultSetOptions) SetLongReturnType(v string) *ExecuteStatementRequestResultSetOptions {
	s.LongReturnType = &v
	return s
}

type ExecuteStatementShrinkRequest struct {
	ContinueAfterTimeout   *bool   `json:"continueAfterTimeout,omitempty" xml:"continueAfterTimeout,omitempty"`
	Database               *string `json:"database,omitempty" xml:"database,omitempty"`
	FormatRecordsAs        *string `json:"formatRecordsAs,omitempty" xml:"formatRecordsAs,omitempty"`
	IncludeResultMetadata  *string `json:"includeResultMetadata,omitempty" xml:"includeResultMetadata,omitempty"`
	ParametersShrink       *string `json:"parameters,omitempty" xml:"parameters,omitempty"`
	ResourceArn            *string `json:"resourceArn,omitempty" xml:"resourceArn,omitempty"`
	ResultSetOptionsShrink *string `json:"resultSetOptions,omitempty" xml:"resultSetOptions,omitempty"`
	SecretArn              *string `json:"secretArn,omitempty" xml:"secretArn,omitempty"`
	Sql                    *string `json:"sql,omitempty" xml:"sql,omitempty"`
	TransactionId          *string `json:"transactionId,omitempty" xml:"transactionId,omitempty"`
}

func (s ExecuteStatementShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteStatementShrinkRequest) GoString() string {
	return s.String()
}

func (s *ExecuteStatementShrinkRequest) SetContinueAfterTimeout(v bool) *ExecuteStatementShrinkRequest {
	s.ContinueAfterTimeout = &v
	return s
}

func (s *ExecuteStatementShrinkRequest) SetDatabase(v string) *ExecuteStatementShrinkRequest {
	s.Database = &v
	return s
}

func (s *ExecuteStatementShrinkRequest) SetFormatRecordsAs(v string) *ExecuteStatementShrinkRequest {
	s.FormatRecordsAs = &v
	return s
}

func (s *ExecuteStatementShrinkRequest) SetIncludeResultMetadata(v string) *ExecuteStatementShrinkRequest {
	s.IncludeResultMetadata = &v
	return s
}

func (s *ExecuteStatementShrinkRequest) SetParametersShrink(v string) *ExecuteStatementShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *ExecuteStatementShrinkRequest) SetResourceArn(v string) *ExecuteStatementShrinkRequest {
	s.ResourceArn = &v
	return s
}

func (s *ExecuteStatementShrinkRequest) SetResultSetOptionsShrink(v string) *ExecuteStatementShrinkRequest {
	s.ResultSetOptionsShrink = &v
	return s
}

func (s *ExecuteStatementShrinkRequest) SetSecretArn(v string) *ExecuteStatementShrinkRequest {
	s.SecretArn = &v
	return s
}

func (s *ExecuteStatementShrinkRequest) SetSql(v string) *ExecuteStatementShrinkRequest {
	s.Sql = &v
	return s
}

func (s *ExecuteStatementShrinkRequest) SetTransactionId(v string) *ExecuteStatementShrinkRequest {
	s.TransactionId = &v
	return s
}

type ExecuteStatementResponseBody struct {
	Code      *string                           `json:"code,omitempty" xml:"code,omitempty"`
	Data      *ExecuteStatementResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message   *string                           `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecuteStatementResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteStatementResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteStatementResponseBody) SetCode(v string) *ExecuteStatementResponseBody {
	s.Code = &v
	return s
}

func (s *ExecuteStatementResponseBody) SetData(v *ExecuteStatementResponseBodyData) *ExecuteStatementResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteStatementResponseBody) SetMessage(v string) *ExecuteStatementResponseBody {
	s.Message = &v
	return s
}

func (s *ExecuteStatementResponseBody) SetRequestId(v string) *ExecuteStatementResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteStatementResponseBody) SetSuccess(v bool) *ExecuteStatementResponseBody {
	s.Success = &v
	return s
}

type ExecuteStatementResponseBodyData struct {
	ColumnMetadata         []*ExecuteStatementResponseBodyDataColumnMetadata  `json:"columnMetadata,omitempty" xml:"columnMetadata,omitempty" type:"Repeated"`
	FormattedRecords       *string                                            `json:"formattedRecords,omitempty" xml:"formattedRecords,omitempty"`
	GeneratedFields        []*ExecuteStatementResponseBodyDataGeneratedFields `json:"generatedFields,omitempty" xml:"generatedFields,omitempty" type:"Repeated"`
	NumberOfRecordsUpdated *int64                                             `json:"numberOfRecordsUpdated,omitempty" xml:"numberOfRecordsUpdated,omitempty"`
	Records                [][]*ExecuteStatementResponseBodyDataRecords       `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
}

func (s ExecuteStatementResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExecuteStatementResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExecuteStatementResponseBodyData) SetColumnMetadata(v []*ExecuteStatementResponseBodyDataColumnMetadata) *ExecuteStatementResponseBodyData {
	s.ColumnMetadata = v
	return s
}

func (s *ExecuteStatementResponseBodyData) SetFormattedRecords(v string) *ExecuteStatementResponseBodyData {
	s.FormattedRecords = &v
	return s
}

func (s *ExecuteStatementResponseBodyData) SetGeneratedFields(v []*ExecuteStatementResponseBodyDataGeneratedFields) *ExecuteStatementResponseBodyData {
	s.GeneratedFields = v
	return s
}

func (s *ExecuteStatementResponseBodyData) SetNumberOfRecordsUpdated(v int64) *ExecuteStatementResponseBodyData {
	s.NumberOfRecordsUpdated = &v
	return s
}

func (s *ExecuteStatementResponseBodyData) SetRecords(v [][]*ExecuteStatementResponseBodyDataRecords) *ExecuteStatementResponseBodyData {
	s.Records = v
	return s
}

type ExecuteStatementResponseBodyDataColumnMetadata struct {
	ArrayBaseColumnType *int32  `json:"arrayBaseColumnType,omitempty" xml:"arrayBaseColumnType,omitempty"`
	IsAutoIncrement     *bool   `json:"isAutoIncrement,omitempty" xml:"isAutoIncrement,omitempty"`
	IsCaseSensitive     *bool   `json:"isCaseSensitive,omitempty" xml:"isCaseSensitive,omitempty"`
	IsCurrency          *bool   `json:"isCurrency,omitempty" xml:"isCurrency,omitempty"`
	IsSigned            *bool   `json:"isSigned,omitempty" xml:"isSigned,omitempty"`
	Label               *string `json:"label,omitempty" xml:"label,omitempty"`
	Name                *string `json:"name,omitempty" xml:"name,omitempty"`
	Nullable            *int32  `json:"nullable,omitempty" xml:"nullable,omitempty"`
	Precision           *int32  `json:"precision,omitempty" xml:"precision,omitempty"`
	Scale               *int32  `json:"scale,omitempty" xml:"scale,omitempty"`
	SchemaName          *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
	TableName           *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
	Type                *int32  `json:"type,omitempty" xml:"type,omitempty"`
	TypeName            *string `json:"typeName,omitempty" xml:"typeName,omitempty"`
}

func (s ExecuteStatementResponseBodyDataColumnMetadata) String() string {
	return tea.Prettify(s)
}

func (s ExecuteStatementResponseBodyDataColumnMetadata) GoString() string {
	return s.String()
}

func (s *ExecuteStatementResponseBodyDataColumnMetadata) SetArrayBaseColumnType(v int32) *ExecuteStatementResponseBodyDataColumnMetadata {
	s.ArrayBaseColumnType = &v
	return s
}

func (s *ExecuteStatementResponseBodyDataColumnMetadata) SetIsAutoIncrement(v bool) *ExecuteStatementResponseBodyDataColumnMetadata {
	s.IsAutoIncrement = &v
	return s
}

func (s *ExecuteStatementResponseBodyDataColumnMetadata) SetIsCaseSensitive(v bool) *ExecuteStatementResponseBodyDataColumnMetadata {
	s.IsCaseSensitive = &v
	return s
}

func (s *ExecuteStatementResponseBodyDataColumnMetadata) SetIsCurrency(v bool) *ExecuteStatementResponseBodyDataColumnMetadata {
	s.IsCurrency = &v
	return s
}

func (s *ExecuteStatementResponseBodyDataColumnMetadata) SetIsSigned(v bool) *ExecuteStatementResponseBodyDataColumnMetadata {
	s.IsSigned = &v
	return s
}

func (s *ExecuteStatementResponseBodyDataColumnMetadata) SetLabel(v string) *ExecuteStatementResponseBodyDataColumnMetadata {
	s.Label = &v
	return s
}

func (s *ExecuteStatementResponseBodyDataColumnMetadata) SetName(v string) *ExecuteStatementResponseBodyDataColumnMetadata {
	s.Name = &v
	return s
}

func (s *ExecuteStatementResponseBodyDataColumnMetadata) SetNullable(v int32) *ExecuteStatementResponseBodyDataColumnMetadata {
	s.Nullable = &v
	return s
}

func (s *ExecuteStatementResponseBodyDataColumnMetadata) SetPrecision(v int32) *ExecuteStatementResponseBodyDataColumnMetadata {
	s.Precision = &v
	return s
}

func (s *ExecuteStatementResponseBodyDataColumnMetadata) SetScale(v int32) *ExecuteStatementResponseBodyDataColumnMetadata {
	s.Scale = &v
	return s
}

func (s *ExecuteStatementResponseBodyDataColumnMetadata) SetSchemaName(v string) *ExecuteStatementResponseBodyDataColumnMetadata {
	s.SchemaName = &v
	return s
}

func (s *ExecuteStatementResponseBodyDataColumnMetadata) SetTableName(v string) *ExecuteStatementResponseBodyDataColumnMetadata {
	s.TableName = &v
	return s
}

func (s *ExecuteStatementResponseBodyDataColumnMetadata) SetType(v int32) *ExecuteStatementResponseBodyDataColumnMetadata {
	s.Type = &v
	return s
}

func (s *ExecuteStatementResponseBodyDataColumnMetadata) SetTypeName(v string) *ExecuteStatementResponseBodyDataColumnMetadata {
	s.TypeName = &v
	return s
}

type ExecuteStatementResponseBodyDataGeneratedFields struct {
	ArrayValue   *ExecuteStatementResponseBodyDataGeneratedFieldsArrayValue `json:"arrayValue,omitempty" xml:"arrayValue,omitempty" type:"Struct"`
	BlobValue    io.Reader                                                  `json:"blobValue,omitempty" xml:"blobValue,omitempty"`
	BooleanValue *bool                                                      `json:"booleanValue,omitempty" xml:"booleanValue,omitempty"`
	DoubleValue  *float64                                                   `json:"doubleValue,omitempty" xml:"doubleValue,omitempty"`
	IsNull       *bool                                                      `json:"isNull,omitempty" xml:"isNull,omitempty"`
	LongValue    *int64                                                     `json:"longValue,omitempty" xml:"longValue,omitempty"`
	StringValue  *string                                                    `json:"stringValue,omitempty" xml:"stringValue,omitempty"`
}

func (s ExecuteStatementResponseBodyDataGeneratedFields) String() string {
	return tea.Prettify(s)
}

func (s ExecuteStatementResponseBodyDataGeneratedFields) GoString() string {
	return s.String()
}

func (s *ExecuteStatementResponseBodyDataGeneratedFields) SetArrayValue(v *ExecuteStatementResponseBodyDataGeneratedFieldsArrayValue) *ExecuteStatementResponseBodyDataGeneratedFields {
	s.ArrayValue = v
	return s
}

func (s *ExecuteStatementResponseBodyDataGeneratedFields) SetBlobValue(v io.Reader) *ExecuteStatementResponseBodyDataGeneratedFields {
	s.BlobValue = v
	return s
}

func (s *ExecuteStatementResponseBodyDataGeneratedFields) SetBooleanValue(v bool) *ExecuteStatementResponseBodyDataGeneratedFields {
	s.BooleanValue = &v
	return s
}

func (s *ExecuteStatementResponseBodyDataGeneratedFields) SetDoubleValue(v float64) *ExecuteStatementResponseBodyDataGeneratedFields {
	s.DoubleValue = &v
	return s
}

func (s *ExecuteStatementResponseBodyDataGeneratedFields) SetIsNull(v bool) *ExecuteStatementResponseBodyDataGeneratedFields {
	s.IsNull = &v
	return s
}

func (s *ExecuteStatementResponseBodyDataGeneratedFields) SetLongValue(v int64) *ExecuteStatementResponseBodyDataGeneratedFields {
	s.LongValue = &v
	return s
}

func (s *ExecuteStatementResponseBodyDataGeneratedFields) SetStringValue(v string) *ExecuteStatementResponseBodyDataGeneratedFields {
	s.StringValue = &v
	return s
}

type ExecuteStatementResponseBodyDataGeneratedFieldsArrayValue struct {
	ArrayValues   []interface{} `json:"arrayValues,omitempty" xml:"arrayValues,omitempty" type:"Repeated"`
	BooleanValues []*bool       `json:"booleanValues,omitempty" xml:"booleanValues,omitempty" type:"Repeated"`
	DoubleValues  []*float64    `json:"doubleValues,omitempty" xml:"doubleValues,omitempty" type:"Repeated"`
	LongValues    []*int64      `json:"longValues,omitempty" xml:"longValues,omitempty" type:"Repeated"`
	StringValues  []*string     `json:"stringValues,omitempty" xml:"stringValues,omitempty" type:"Repeated"`
}

func (s ExecuteStatementResponseBodyDataGeneratedFieldsArrayValue) String() string {
	return tea.Prettify(s)
}

func (s ExecuteStatementResponseBodyDataGeneratedFieldsArrayValue) GoString() string {
	return s.String()
}

func (s *ExecuteStatementResponseBodyDataGeneratedFieldsArrayValue) SetArrayValues(v []interface{}) *ExecuteStatementResponseBodyDataGeneratedFieldsArrayValue {
	s.ArrayValues = v
	return s
}

func (s *ExecuteStatementResponseBodyDataGeneratedFieldsArrayValue) SetBooleanValues(v []*bool) *ExecuteStatementResponseBodyDataGeneratedFieldsArrayValue {
	s.BooleanValues = v
	return s
}

func (s *ExecuteStatementResponseBodyDataGeneratedFieldsArrayValue) SetDoubleValues(v []*float64) *ExecuteStatementResponseBodyDataGeneratedFieldsArrayValue {
	s.DoubleValues = v
	return s
}

func (s *ExecuteStatementResponseBodyDataGeneratedFieldsArrayValue) SetLongValues(v []*int64) *ExecuteStatementResponseBodyDataGeneratedFieldsArrayValue {
	s.LongValues = v
	return s
}

func (s *ExecuteStatementResponseBodyDataGeneratedFieldsArrayValue) SetStringValues(v []*string) *ExecuteStatementResponseBodyDataGeneratedFieldsArrayValue {
	s.StringValues = v
	return s
}

type ExecuteStatementResponseBodyDataRecords struct {
	ArrayValue   *ExecuteStatementResponseBodyDataRecordsArrayValue `json:"arrayValue,omitempty" xml:"arrayValue,omitempty" type:"Struct"`
	BlobValue    io.Reader                                          `json:"blobValue,omitempty" xml:"blobValue,omitempty"`
	BooleanValue *bool                                              `json:"booleanValue,omitempty" xml:"booleanValue,omitempty"`
	DoubleValue  *float64                                           `json:"doubleValue,omitempty" xml:"doubleValue,omitempty"`
	IsNull       *bool                                              `json:"isNull,omitempty" xml:"isNull,omitempty"`
	LongValue    *int64                                             `json:"longValue,omitempty" xml:"longValue,omitempty"`
	StringValue  *string                                            `json:"stringValue,omitempty" xml:"stringValue,omitempty"`
}

func (s ExecuteStatementResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s ExecuteStatementResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ExecuteStatementResponseBodyDataRecords) SetArrayValue(v *ExecuteStatementResponseBodyDataRecordsArrayValue) *ExecuteStatementResponseBodyDataRecords {
	s.ArrayValue = v
	return s
}

func (s *ExecuteStatementResponseBodyDataRecords) SetBlobValue(v io.Reader) *ExecuteStatementResponseBodyDataRecords {
	s.BlobValue = v
	return s
}

func (s *ExecuteStatementResponseBodyDataRecords) SetBooleanValue(v bool) *ExecuteStatementResponseBodyDataRecords {
	s.BooleanValue = &v
	return s
}

func (s *ExecuteStatementResponseBodyDataRecords) SetDoubleValue(v float64) *ExecuteStatementResponseBodyDataRecords {
	s.DoubleValue = &v
	return s
}

func (s *ExecuteStatementResponseBodyDataRecords) SetIsNull(v bool) *ExecuteStatementResponseBodyDataRecords {
	s.IsNull = &v
	return s
}

func (s *ExecuteStatementResponseBodyDataRecords) SetLongValue(v int64) *ExecuteStatementResponseBodyDataRecords {
	s.LongValue = &v
	return s
}

func (s *ExecuteStatementResponseBodyDataRecords) SetStringValue(v string) *ExecuteStatementResponseBodyDataRecords {
	s.StringValue = &v
	return s
}

type ExecuteStatementResponseBodyDataRecordsArrayValue struct {
	ArrayValues   []interface{} `json:"arrayValues,omitempty" xml:"arrayValues,omitempty" type:"Repeated"`
	BooleanValues []*bool       `json:"booleanValues,omitempty" xml:"booleanValues,omitempty" type:"Repeated"`
	DoubleValues  []*float64    `json:"doubleValues,omitempty" xml:"doubleValues,omitempty" type:"Repeated"`
	LongValues    []*int64      `json:"longValues,omitempty" xml:"longValues,omitempty" type:"Repeated"`
	StringValues  []*string     `json:"stringValues,omitempty" xml:"stringValues,omitempty" type:"Repeated"`
}

func (s ExecuteStatementResponseBodyDataRecordsArrayValue) String() string {
	return tea.Prettify(s)
}

func (s ExecuteStatementResponseBodyDataRecordsArrayValue) GoString() string {
	return s.String()
}

func (s *ExecuteStatementResponseBodyDataRecordsArrayValue) SetArrayValues(v []interface{}) *ExecuteStatementResponseBodyDataRecordsArrayValue {
	s.ArrayValues = v
	return s
}

func (s *ExecuteStatementResponseBodyDataRecordsArrayValue) SetBooleanValues(v []*bool) *ExecuteStatementResponseBodyDataRecordsArrayValue {
	s.BooleanValues = v
	return s
}

func (s *ExecuteStatementResponseBodyDataRecordsArrayValue) SetDoubleValues(v []*float64) *ExecuteStatementResponseBodyDataRecordsArrayValue {
	s.DoubleValues = v
	return s
}

func (s *ExecuteStatementResponseBodyDataRecordsArrayValue) SetLongValues(v []*int64) *ExecuteStatementResponseBodyDataRecordsArrayValue {
	s.LongValues = v
	return s
}

func (s *ExecuteStatementResponseBodyDataRecordsArrayValue) SetStringValues(v []*string) *ExecuteStatementResponseBodyDataRecordsArrayValue {
	s.StringValues = v
	return s
}

type ExecuteStatementResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExecuteStatementResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecuteStatementResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteStatementResponse) GoString() string {
	return s.String()
}

func (s *ExecuteStatementResponse) SetHeaders(v map[string]*string) *ExecuteStatementResponse {
	s.Headers = v
	return s
}

func (s *ExecuteStatementResponse) SetStatusCode(v int32) *ExecuteStatementResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteStatementResponse) SetBody(v *ExecuteStatementResponseBody) *ExecuteStatementResponse {
	s.Body = v
	return s
}

type RollbackTransactionRequest struct {
	Database      *string `json:"database,omitempty" xml:"database,omitempty"`
	ResourceArn   *string `json:"resourceArn,omitempty" xml:"resourceArn,omitempty"`
	SecretArn     *string `json:"secretArn,omitempty" xml:"secretArn,omitempty"`
	TransactionId *string `json:"transactionId,omitempty" xml:"transactionId,omitempty"`
}

func (s RollbackTransactionRequest) String() string {
	return tea.Prettify(s)
}

func (s RollbackTransactionRequest) GoString() string {
	return s.String()
}

func (s *RollbackTransactionRequest) SetDatabase(v string) *RollbackTransactionRequest {
	s.Database = &v
	return s
}

func (s *RollbackTransactionRequest) SetResourceArn(v string) *RollbackTransactionRequest {
	s.ResourceArn = &v
	return s
}

func (s *RollbackTransactionRequest) SetSecretArn(v string) *RollbackTransactionRequest {
	s.SecretArn = &v
	return s
}

func (s *RollbackTransactionRequest) SetTransactionId(v string) *RollbackTransactionRequest {
	s.TransactionId = &v
	return s
}

type RollbackTransactionResponseBody struct {
	Code      *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Data      *RollbackTransactionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message   *string                              `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RollbackTransactionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RollbackTransactionResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackTransactionResponseBody) SetCode(v string) *RollbackTransactionResponseBody {
	s.Code = &v
	return s
}

func (s *RollbackTransactionResponseBody) SetData(v *RollbackTransactionResponseBodyData) *RollbackTransactionResponseBody {
	s.Data = v
	return s
}

func (s *RollbackTransactionResponseBody) SetMessage(v string) *RollbackTransactionResponseBody {
	s.Message = &v
	return s
}

func (s *RollbackTransactionResponseBody) SetRequestId(v string) *RollbackTransactionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackTransactionResponseBody) SetSuccess(v bool) *RollbackTransactionResponseBody {
	s.Success = &v
	return s
}

type RollbackTransactionResponseBodyData struct {
	TransactionStatus *string `json:"transactionStatus,omitempty" xml:"transactionStatus,omitempty"`
}

func (s RollbackTransactionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RollbackTransactionResponseBodyData) GoString() string {
	return s.String()
}

func (s *RollbackTransactionResponseBodyData) SetTransactionStatus(v string) *RollbackTransactionResponseBodyData {
	s.TransactionStatus = &v
	return s
}

type RollbackTransactionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RollbackTransactionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RollbackTransactionResponse) String() string {
	return tea.Prettify(s)
}

func (s RollbackTransactionResponse) GoString() string {
	return s.String()
}

func (s *RollbackTransactionResponse) SetHeaders(v map[string]*string) *RollbackTransactionResponse {
	s.Headers = v
	return s
}

func (s *RollbackTransactionResponse) SetStatusCode(v int32) *RollbackTransactionResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackTransactionResponse) SetBody(v *RollbackTransactionResponseBody) *RollbackTransactionResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("rds-data"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchExecuteStatementWithOptions(tmpReq *BatchExecuteStatementRequest, runtime *util.RuntimeOptions) (_result *BatchExecuteStatementResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BatchExecuteStatementShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ParameterSets)) {
		request.ParameterSetsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ParameterSets, tea.String("parameterSets"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Database)) {
		body["database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.ParameterSetsShrink)) {
		body["parameterSets"] = request.ParameterSetsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceArn)) {
		body["resourceArn"] = request.ResourceArn
	}

	if !tea.BoolValue(util.IsUnset(request.SecretArn)) {
		body["secretArn"] = request.SecretArn
	}

	if !tea.BoolValue(util.IsUnset(request.Sql)) {
		body["sql"] = request.Sql
	}

	if !tea.BoolValue(util.IsUnset(request.TransactionId)) {
		body["transactionId"] = request.TransactionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchExecuteStatement"),
		Version:     tea.String("2022-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchExecuteStatementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchExecuteStatement(request *BatchExecuteStatementRequest) (_result *BatchExecuteStatementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchExecuteStatementResponse{}
	_body, _err := client.BatchExecuteStatementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BeginTransactionWithOptions(request *BeginTransactionRequest, runtime *util.RuntimeOptions) (_result *BeginTransactionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Database)) {
		body["database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceArn)) {
		body["resourceArn"] = request.ResourceArn
	}

	if !tea.BoolValue(util.IsUnset(request.SecretArn)) {
		body["secretArn"] = request.SecretArn
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BeginTransaction"),
		Version:     tea.String("2022-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BeginTransactionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BeginTransaction(request *BeginTransactionRequest) (_result *BeginTransactionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BeginTransactionResponse{}
	_body, _err := client.BeginTransactionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CommitTransactionWithOptions(request *CommitTransactionRequest, runtime *util.RuntimeOptions) (_result *CommitTransactionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Database)) {
		body["database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceArn)) {
		body["resourceArn"] = request.ResourceArn
	}

	if !tea.BoolValue(util.IsUnset(request.SecretArn)) {
		body["secretArn"] = request.SecretArn
	}

	if !tea.BoolValue(util.IsUnset(request.TransactionId)) {
		body["transactionId"] = request.TransactionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CommitTransaction"),
		Version:     tea.String("2022-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CommitTransactionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CommitTransaction(request *CommitTransactionRequest) (_result *CommitTransactionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CommitTransactionResponse{}
	_body, _err := client.CommitTransactionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecuteStatementWithOptions(tmpReq *ExecuteStatementRequest, runtime *util.RuntimeOptions) (_result *ExecuteStatementResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ExecuteStatementShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("parameters"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.ResultSetOptions))) {
		request.ResultSetOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.ResultSetOptions), tea.String("resultSetOptions"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContinueAfterTimeout)) {
		body["continueAfterTimeout"] = request.ContinueAfterTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.Database)) {
		body["database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.FormatRecordsAs)) {
		body["formatRecordsAs"] = request.FormatRecordsAs
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeResultMetadata)) {
		body["includeResultMetadata"] = request.IncludeResultMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.ParametersShrink)) {
		body["parameters"] = request.ParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceArn)) {
		body["resourceArn"] = request.ResourceArn
	}

	if !tea.BoolValue(util.IsUnset(request.ResultSetOptionsShrink)) {
		body["resultSetOptions"] = request.ResultSetOptionsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SecretArn)) {
		body["secretArn"] = request.SecretArn
	}

	if !tea.BoolValue(util.IsUnset(request.Sql)) {
		body["sql"] = request.Sql
	}

	if !tea.BoolValue(util.IsUnset(request.TransactionId)) {
		body["transactionId"] = request.TransactionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteStatement"),
		Version:     tea.String("2022-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteStatementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteStatement(request *ExecuteStatementRequest) (_result *ExecuteStatementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteStatementResponse{}
	_body, _err := client.ExecuteStatementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RollbackTransactionWithOptions(request *RollbackTransactionRequest, runtime *util.RuntimeOptions) (_result *RollbackTransactionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Database)) {
		body["database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceArn)) {
		body["resourceArn"] = request.ResourceArn
	}

	if !tea.BoolValue(util.IsUnset(request.SecretArn)) {
		body["secretArn"] = request.SecretArn
	}

	if !tea.BoolValue(util.IsUnset(request.TransactionId)) {
		body["transactionId"] = request.TransactionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RollbackTransaction"),
		Version:     tea.String("2022-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RollbackTransactionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RollbackTransaction(request *RollbackTransactionRequest) (_result *RollbackTransactionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RollbackTransactionResponse{}
	_body, _err := client.RollbackTransactionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

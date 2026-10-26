# Amazon Selling Partner API Go SDK

这个目录是面向 Go 后端的 Amazon Selling Partner API（SP-API）接入层。Amazon 官方 `selling-partner-api-sdk` 当前没有 Go 客户端，因此本项目采用以下分层：

- `auth`：LWA refresh token、grantless token、进程内 access token 缓存和 RDT。
- `client`：统一 `http.Client`、SP-API 请求头、区域 endpoint、429/5xx 重试和请求错误。
- `generated`：由 Amazon 官方 OpenAPI models 生成的 Go API 包。生成代码不承担认证。
- `examples`：可运行的最小 Sellers API 调用示例。

根目录的接口选型和 ERP 数据模型见 [SP_API_ERP接口分析.md](../SP_API_ERP接口分析.md) 与 [GO_SP_API接入指南.md](../GO_SP_API接入指南.md)。新项目先阅读本文件，再按 Seller/Vendor、FBA/FBM/MCF 和 P0/P1/P2 范围生成需要的 API。

## 1. 准备应用和租户凭证

在 Amazon Developer Console 创建 SP-API 应用，并根据实际操作申请角色和权限。每个 Seller 租户至少保存：

- `clientId`、`clientSecret`：应用级凭证，使用 Secrets Manager、Vault 等密钥服务保存。
- `refreshToken`：卖家授权后获得，按租户单独保存并加密；不能放进全局配置后被所有租户共用。
- `sellerId`、`marketplaceId`、`region`：账号和站点配置。

请求订单地址、买家信息等受限字段前，还要根据精确的 HTTP 方法和路径申请 Restricted Data Token（RDT）。日志中不能打印 refresh token、access token、RDT 或 PII。

## 2. 安装和运行示例

```powershell
cd D:\work\amzn\selling-partner-api-go
go test ./...

$env:SPAPI_LWA_CLIENT_ID = "amzn1.application-..."
$env:SPAPI_LWA_CLIENT_SECRET = "amzn1.oa2-cs.v1-..."
$env:SPAPI_LWA_REFRESH_TOKEN = "Atzr|..."
go run ./examples/sellers
```

示例使用北美 endpoint。欧洲和远东区域使用 `client.RegionEU`、`client.RegionFE`，不要把 marketplace ID 当成 endpoint。

## 3. 使用生成 API 包

生成包通过 `Configuration.HTTPClient` 接入统一传输层：

```go
tokens := auth.NewLWAProvider(auth.Credentials{
	ClientID:     os.Getenv("SPAPI_LWA_CLIENT_ID"),
	ClientSecret: os.Getenv("SPAPI_LWA_CLIENT_SECRET"),
	RefreshToken: os.Getenv("SPAPI_LWA_REFRESH_TOKEN"),
})

endpoint, err := client.Endpoint(client.RegionNA)
if err != nil {
	return err
}

httpClient := client.NewHTTPClient(client.HTTPClientOptions{
	TokenSource: tokens,
	UserAgent:   "MyERP/1.0 (Language=Go; Platform=Linux)",
	MaxRetries:  3,
})

cfg := sellers.NewConfiguration()
cfg.HTTPClient = httpClient
cfg.Servers[0].URL = endpoint
api := sellers.NewAPIClient(cfg)

result, response, err := api.SellersAPI.
	GetMarketplaceParticipations(ctx).
	Execute()
```

生成 API 的方法通常返回 `result, *http.Response, error`。成功响应里的 `x-amzn-RequestId` 可从 `response.Header` 读取；错误会保留生成客户端的错误模型和 HTTP response。

## 4. 直接调用尚未生成的 API

新接口还没有生成包时，可以使用 `client.DoJSON`，仍然复用同一个认证和重试层：

```go
var payload struct {
	Reports []map[string]any `json:"reports"`
}

err := client.DoJSON(ctx, httpClient, endpoint,
	"/reports/2021-06-30/reports", http.MethodGet,
	url.Values{"marketplaceIds": {marketplaceID}}, nil, &payload)
```

写操作的 JSON body 传入结构体或 map。需要分页时，把 `nextToken` 原样保存并传回下一次请求；不要在业务层截断或重新编码 token。

## 5. LWA、grantless 和 RDT

普通 Seller/Vendor 请求使用 refresh token 换 LWA access token：

```go
accessToken, err := tokens.AccessToken(ctx)
```

使用 grantless operation 时按 scope 单独获取 token，不能与租户 refresh token 的缓存键混用：

```go
grantlessToken, err := tokens.GrantlessToken(ctx,
	"sellingpartnerapi::notifications")
```

RDT 流程是先用普通 LWA token 调 Tokens API，再把 RDT 仅用于声明的受限路径：

```go
rdtClient := auth.RDTClient{
	HTTP:        httpClient,
	Endpoint:    endpoint,
	UserAgent:   "MyERP/1.0 (Language=Go)",
	TokenSource: tokens,
}
rdt, err := rdtClient.CreateRestrictedDataToken(ctx, []auth.RestrictedResource{{
	Method:       http.MethodGet,
	Path:         "/orders/v0/orders/ORDER_ID/address",
	DataElements: []string{"shippingAddress"},
}})
if err != nil {
	return err
}

ctx = client.WithAccessToken(ctx, rdt)
// 把这个 ctx 传给生成的订单地址 API，或传给 client.DoJSON。
```

RDT 生命周期短，不应写入长期缓存。订单地址、买家信息等 PII 也不应写入普通业务日志。

## 6. 生成 API 包

脚本从本工作目录的 `selling-partner-api-models/models` 读取官方 OpenAPI 文件，并固定使用 OpenAPI Generator `7.9.0`。需要 Docker Desktop：

```powershell
cd D:\work\amzn\selling-partner-api-go
.\scripts\generate.ps1 -Api @("sellers","catalog","product_types","listings","orders","inventory","reports","finances","tokens","notifications")
```

也可以只生成一个 API：

```powershell
.\scripts\generate.ps1 -Api orders
```

生成目录为 `generated/<api>`，包名与目录名一致。Amazon API 版本变更时，更新 `scripts/generate.ps1` 中的 spec 文件，再重新生成并运行测试。不要直接修改 `generated` 下的文件；需要定制时在业务 adapter 层处理。

## 7. ERP 接入顺序

P0 只读同步建议依次接入：

1. Website Authorization/LWA、Sellers：建立租户、站点和区域配置。
2. Catalog Items、Product Type Definitions、Listings Items：建立商品和 Listing 映射。
3. Orders：按 `lastUpdatedAfter` 增量同步，并用 Notifications 触发补偿任务。
4. FBA Inventory 或 Listings/Feeds：建立库存快照。
5. Reports、Finances：做历史补偿和日终对账。

P1 再增加 Notifications、RDT、Feeds、Product Pricing、Product Fees、FBA Inbound、Merchant Fulfillment/Shipping、Tracking 和 MCF。Vendor Central 使用独立 adapter，不要把 Vendor Orders 和 Seller Orders 直接映射成同一个 Amazon DTO。

ERP 内部至少保留 `channelAccountId`、`marketplaceId` 和 `region`。Amazon `orderId`、SKU、ASIN 不能直接作为跨平台全局主键。

## 8. 上线检查

- 每个租户独立保存并加密 refresh token。
- 访问受限数据前申请 RDT，并按精确路径使用。
- 处理所有分页、429、5xx、请求 ID 和幂等键。
- Reports、Feeds、Data Kiosk、FBA 入库使用异步任务状态机。
- Notifications 去重、容忍乱序并支持失败重试。
- 对订单、库存、财务保留定时补偿同步和对账任务。
- 发布前重新核对当前 API 版本、角色、区域和 marketplace 可用性。

param(
  [string[]]$Api = @("sellers")
)

$ErrorActionPreference = "Stop"
$root = Split-Path -Parent $PSScriptRoot
$models = Join-Path (Split-Path -Parent $root) "selling-partner-api-models\models"
$generator = "openapitools/openapi-generator-cli:v7.9.0"

$specs = @{
  sellers = "sellers-api-model\sellers.json"
  catalog = "catalog-items-api-model\catalogItems_2022-04-01.json"
  product_types = "product-type-definitions-api-model\definitionsProductTypes_2020-09-01.json"
  listings = "listings-items-api-model\listingsItems_2021-08-01.json"
  orders = "orders-api-model\orders_2026-01-01.json"
  inventory = "fba-inventory-api-model\fbaInventory.json"
  reports = "reports-api-model\reports_2021-06-30.json"
  finances = "finances-api-model\finances_2024-06-19.json"
  tokens = "tokens-api-model\tokens_2021-03-01.json"
  notifications = "notifications-api-model\notifications.json"
}

foreach ($name in $Api) {
  if (-not $specs.ContainsKey($name)) {
    throw "Unknown API '$name'. Supported values: $($specs.Keys -join ', ')"
  }
  $output = Join-Path $root "generated\$name"
  New-Item -ItemType Directory -Force -Path $output | Out-Null
  docker run --rm -v "${root}:/local" -v "${models}:/models:ro" $generator generate `
    -i "/models/$($specs[$name].Replace('\','/'))" `
    -g go `
    -o "/local/generated/$name" `
    --global-property "apiTests=false,modelTests=false" `
    --additional-properties="packageName=$name,generateInterfaces=true,enumClassPrefix=true"
  Remove-Item -Force -ErrorAction SilentlyContinue (Join-Path $output "go.mod"), (Join-Path $output "go.sum")
  Remove-Item -Recurse -Force -ErrorAction SilentlyContinue (Join-Path $output "test")
  Get-ChildItem -Path $output -Filter *.go -Recurse | ForEach-Object {
    $content = Get-Content -Raw $_.FullName
    if ($content -match 'var defaultValue \[\]string = \["') {
      $content = [regex]::Replace($content, 'var defaultValue \[\]string = \["([^"]*)"\]', 'var defaultValue []string = []string{"$1"}')
      Set-Content -NoNewline -Encoding utf8 $_.FullName $content
    }
  }
}

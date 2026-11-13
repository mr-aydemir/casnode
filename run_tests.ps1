<#
.SYNOPSIS
  MakerSpell Forum - Automated Test Suite Runner
.DESCRIPTION
  Runs both the Go unit/integrity tests and the E2E user scenario tests.
#>

$ErrorActionPreference = "Stop"
$forumRoot = $PSScriptRoot

Write-Host "======================================================" -ForegroundColor Cyan
Write-Host "  MakerSpell Forum - Automated Verification Runner" -ForegroundColor Cyan
Write-Host "======================================================" -ForegroundColor Cyan

# 1. Run Go Unit and Integrity Tests
Write-Host "`n[Step 1/2] Running Go Unit & Catalog Integrity Tests..." -ForegroundColor Yellow
Set-Location -Path $forumRoot
go test -v ./object/ -run "TestPrinterCatalogStructure|TestSlugUniquenessAcrossBrandsAndModels|TestSeederTemplates|TestSortCategoriesForDisplayHierarchy|TestMakerSpellAuthProviderSpec"

if ($LASTEXITCODE -ne 0) {
    Write-Host "`n❌ Go unit tests failed!" -ForegroundColor Red
    exit $LASTEXITCODE
}
Write-Host "✅ Go unit & integrity tests passed!" -ForegroundColor Green

# 2. Run E2E User Scenario Tests
Write-Host "`n[Step 2/2] Running End-to-End User Scenario Tests..." -ForegroundColor Yellow
node "$forumRoot/tests/e2e_user_scenarios.mjs"

if ($LASTEXITCODE -ne 0) {
    Write-Host "`n❌ E2E scenario tests failed!" -ForegroundColor Red
    exit $LASTEXITCODE
}

Write-Host "`n======================================================" -ForegroundColor Green
Write-Host "  🎉 All automated tests passed successfully!" -ForegroundColor Green
Write-Host "======================================================" -ForegroundColor Green

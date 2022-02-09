$ErrorActionPreference = "Stop"
# in the future, this variable should cause PS to exit on non-zero exit codes from commands/exes (as opposed to PS cmdlets)
$PSNativeCommandUseErrorActionPreference = $true
# see https://github.com/PowerShell/PowerShell/issues/3415 and https://github.com/PowerShell/PowerShell-RFC/pull/277

Move-Item out\hello.exe win\cf8.exe

# convert line-endings
Get-Content win\LICENSE-WITH-3RD-PARTY-LICENSES | Set-Content win\LICENSE
Get-Content win\CF_NOTICE | Set-Content win\NOTICE

iscc win\windows-installer-v8-x64.iss
Move-Item win\Output\mysetup.exe win\Output\cfsetup.exe

Get-ChildItem win\Output
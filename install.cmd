@echo off
title Xray Launcher - Installer
echo ============================================
echo    XRAY LAUNCHER - INSTALLER
echo ============================================
echo.

echo [1/2] Checking Go...
where go >nul 2>nul
if %errorlevel% neq 0 (
    echo Go not found. Downloading...
    powershell -Command "Invoke-WebRequest -Uri 'https://go.dev/dl/go1.21.0.windows-amd64.msi' -OutFile 'go.msi'"
    echo Installing Go...
    msiexec /i go.msi /quiet
    del go.msi
    echo Go installed!
) else (
    echo Go OK
)

echo.
echo [2/2] Building project...
go build -o XrayLauncher.exe main.go
echo Done!

echo.
echo Starting Xray Launcher...
start XrayLauncher.exe
exit

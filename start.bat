@echo off
title erikao-ai
color 0F
echo.
echo  erikao :: ai terminal
echo  ========================
echo.

:: Para processos anteriores se existirem
taskkill /f /im erikaoia.exe 2>nul
taskkill /f /im ngrok.exe 2>nul
timeout /t 1 /nobreak > nul

:: Builda o executavel
echo  Compilando servidor...
go build -o erikaoia.exe main.go 2>nul

:: Inicia o servidor Go em janela separada
echo  Iniciando servidor Go...
start "erikao-server" cmd /k "erikaoia.exe"
timeout /t 2 /nobreak > nul

:: Inicia ngrok com dominio fixo
echo  Iniciando tunel ngrok...
echo  Dashboard: http://localhost:4040
echo  URL publica: https://hypnoidal-velia-acrimoniously.ngrok-free.dev
echo.
ngrok http 8080 --domain=hypnoidal-velia-acrimoniously.ngrok-free.dev

pause

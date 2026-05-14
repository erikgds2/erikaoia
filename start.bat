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

:: Builda o executavel (se necessario)
echo  Compilando servidor...
go build -o erikaoia.exe main.go 2>nul

:: Inicia o servidor Go em janela separada
echo  Iniciando servidor Go...
start "erikao-server" cmd /k "erikaoia.exe"
timeout /t 2 /nobreak > nul

:: Inicia ngrok com dominio estatico (altere SEUDOMINIO pelo seu dominio ngrok)
echo  Iniciando tunel ngrok...
echo  Dashboard: http://localhost:4040
echo.

:: Se voce tem dominio estatico, use: ngrok http 8080 --domain=SEU-DOMINIO.ngrok-free.app
:: Se nao tem, use a linha abaixo (URL muda a cada reinicio):
ngrok http 8080

echo.
pause

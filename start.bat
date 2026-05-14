@echo off
title erikao-ai
color 0F
echo.
echo  erikao :: ai terminal
echo  ========================
echo.

:: Inicia o servidor Go em janela separada
start "erikao-server" cmd /k "go run main.go"

timeout /t 2 /nobreak > nul

:: Inicia o tunel ngrok
echo  Iniciando tunel ngrok...
echo  Acesse http://localhost:4040 para ver a URL publica
echo.
ngrok http 8080

echo.
echo  Pressione qualquer tecla para sair...
pause > nul

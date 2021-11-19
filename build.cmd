@echo off

if "%~1"=="" goto NONE_ARG
CALL :%~1

@REM Exception, End
:NONE_ARG
echo Need 1 argument
goto EXIT


:all
cmd /c build.cmd win
cmd /c build.cmd wasm
cmd /c build.cmd server
goto EXIT


:win
set GOOS=windows
set GOARCH=amd64
go build -o ./bin/win/
goto EXIT


:wasm
set GOOS=js
set GOARCH=wasm
go build -o ./bin/wasm/main.wasm

copy wasm\assets\*.html bin\wasm
copy %GOROOT%\misc\wasm\wasm_exec.js bin\wasm
goto EXIT


:server
set GOOS=windows
set GOARCH=amd64

go build -o ./bin/wasm/ ./wasm/server
goto EXIT


@REM Try later
:android
set APP=com.hello-ebiten.practice-golang
go install github.com/hajimehoshi/ebiten/v2/cmd/ebitenmobile@latest
.\bin\ebitenmobile.exe bind -target android -javapkg %APP% -o bin/android/hello-ebiten.aar .
goto EXIT


:clean
del bin\* /s /q
FOR /F "eol=| delims=" %%I in ('dir "bin\*" /AD /B 2^>nul') do rd /Q /S "bin\%%I"
goto EXIT

:EXIT
exit
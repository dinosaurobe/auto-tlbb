@echo off
SET NEWLINE=^& echo.
FIND /C /I "tieudattai.org" %WINDIR%\system32\drivers\etc\hosts
IF %ERRORLEVEL% NEQ 0 ECHO %NEWLINE%>>%WINDIR%\System32\drivers\etc\hosts
IF %ERRORLEVEL% NEQ 0 ECHO 127.0.0.1 tieudattai.org>>%WINDIR%\System32\drivers\etc\hosts

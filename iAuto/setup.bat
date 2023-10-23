@echo off
set "HostsFile=%SystemRoot%\System32\drivers\etc\hosts"
set "TempFile=%TEMP%\%~n0.tmp"
%SystemRoot%\System32\attrib.exe -r "%HostsFile%"
%SystemRoot%\System32\findstr.exe /I /R /V /C:"tieudattai\.org" /C:"^$" "%HostsFile%" > "%TempFile%"
COPY /B /V /Y "%TempFile%" "%HostsFile%"
del "%TempFile%"

FIND /C /I "tieudattai.org" "%HostsFile%"
IF %ERRORLEVEL% NEQ 0 (
    echo.>> "%HostsFile%"
    echo 127.0.0.1 tieudattai.org>> "%HostsFile%"
)
set "HostsFile="
set "TempFile="
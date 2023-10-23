@echo off
set "HostsFile=%SystemRoot%\System32\drivers\etc\hosts"
%SystemRoot%\System32\attrib.exe -r "%HostsFile%"
%SystemRoot%\System32\findstr.exe /I /R /V "tieudattai\.org ^$" "%HostsFile%" >"%TEMP%\%~n0.tmp"
COPY /B /V /Y "%TEMP%\%~n0.tmp" "%HostsFile%"
if errorlevel 1 del "%TEMP%\%~n0.tmp"
set "HostsFile="
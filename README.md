Powershell command to query IIS for sites and bindings. Does not work with 2k3 server

Import-Module WebAdministration
Get-ChildItem -Path IIS:\Sites | ConvertTo-Csv -NoTypeInformation > iis.csv

To get Sites and bindings use appcmd.exe
C:\Windows\System32\inetsrv\appcmd.exe list sites > iis.txt

This go program parse both files and matches and outputs to CSV
go run main.go > complete.csv
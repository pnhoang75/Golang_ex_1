# Overview:
This is a Golang exercise that provides a http server which has 2 functions: GetSystemInfo and Ping.

# Build Instruction:
To build binary, cd to the source directory and run command:
> % go build .

# Installation:
golang_ex_1.pkg is installable package for MacOS.  Simply double click on the file to install.  The executlable binary "golang_ex+1" will be added to /usr/loca/bin

# Run:
Simply execute the binary from Terminal
> % golang_ex_1 &

# Testing:
## GetSystemInfo:
> % curl localhost:8080/execute  
> {"Hostname":"Marys-MacBook-Air.local","IPAddress":"192.168.0.16"}

## Ping:
> % curl -X POST localhost:8080/execute -d "www.cnn.com"  
> {"Successful":true,"Time":17529200}

## Unit tests:
Run the following command on the source code directory:
> % go test -v  
> === RUN   TestGetSystemInfo  
> --- PASS: TestGetSystemInfo (0.00s)  
> === RUN   TestPing  
> testcase: Invalid host  
> testcase: localhost  
> testcase: remote host  
> --- PASS: TestPing (8.28s)  
> PASS  
> ok  	golang_ex_1	8.709s  


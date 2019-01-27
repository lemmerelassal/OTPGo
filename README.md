# OTPGo
Open-source Trading Platform written in Go

## Progress
This is in the early stages and not really functional yet.

## Compiling Protobuf line
protoc -I"C:\users\``yourusername``\go\src" -I".\proto\backend" --go_out=plugins=grpc:".\pkg\proto\" .\proto\backend\backend.proto

## Running GRPCC
grpcc -i -p .\otpgo\proto\backend\backend.proto --address 127.0.0.1:4711
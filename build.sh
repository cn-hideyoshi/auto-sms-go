echo "building gateway.."
go build -o bin/gateway-auto-sms gateway/main.go
echo "building company.."
go build -o bin/company-auto-sms company/main.go
echo "building user.."
go build -o bin/user-auto-sms user/main.go
echo "building msg.."
go build -o bin/msg-auto-sms msg/main.go
echo "building down"
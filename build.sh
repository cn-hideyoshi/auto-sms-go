echo "building gateway.."
go build -o bin/gateway gateway/main.go
echo "building company.."
go build -o bin/company company/main.go
echo "building user.."
go build -o bin/user user/main.go
echo "building msg.."
go build -o bin/msg msg/main.go
echo "building down"
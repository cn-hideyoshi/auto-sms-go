if [ $1 ]; then
  echo "building in ./$1/ service"
  protoc --proto_path=$1  --go_out=../pkg/service/$1.v1/ --go_opt=paths=source_relative --go-grpc_out=../pkg/service/$1.v1/ --go-grpc_opt=paths=source_relative ./$1/*.proto
  else echo "please enter service dirname"
fi

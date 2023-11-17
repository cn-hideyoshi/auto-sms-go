if [ $1 ]; then
  echo "正在生成$1文件下的service"
  protoc --proto_path=$1  --go_out=../pkg/service/$1.v1/ --go_opt=paths=source_relative --go-grpc_out=../pkg/service/$1.v1/ --go-grpc_opt=paths=source_relative ./$1/*.proto
  else echo "请输入待生成文件夹"
fi

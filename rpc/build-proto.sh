SRC_DIR="./"
DST_DIR="./"

protoc -I=$SRC_DIR --go_out=$DST_DIR --go-grpc_out=$DST_DIR $SRC_DIR/ultipa.proto
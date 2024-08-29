# 检查并安装 goimports
if ! command -v goimports &> /dev/null; then
    echo "goimports not found, installing..."
    go install golang.org/x/tools/cmd/goimports@latest
else
    echo "goimports already installed"
fi

# 检查并安装 golint
if ! command -v golint &> /dev/null; then
    echo "golint not found, installing..."
    go install golang.org/x/lint/golint@latest
else
    echo "golint already installed"
fi

# 查找所有 Go 文件并检查未格式化的文件
gofiles=$(find ./ -name '*.go') && [ -z "$gofiles" ] \
    || unformatted=$(goimports -l $gofiles) && [ -z "$unformatted" ] \
    || (echo >&2 "Go files must be formatted with goimports. Following files have problems: $unformatted" && true);

# 比较格式化后的代码差异
echo "diff start"
diff <(echo -n) <(gofmt -s -d .)
echo "diff end"

# 运行 golint 检查代码
echo "golint start"
golint ./...
echo "golint end"
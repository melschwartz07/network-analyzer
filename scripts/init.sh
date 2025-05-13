if [ ! -f go.mod ]; then
    go mod init network-analyzer
fi
go mod tidy
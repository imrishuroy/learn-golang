### package cover 
    go get golang.org/x/tools/cmd/cover

### Now we can run our tests with an extra flag, -cover.
    go test -cover

### This will produce a html file.
    go test -coverprofile fmtcoverage.html

### This will open a browser with information on the coverage.
    go tool cover -html=fmtcoverage.html


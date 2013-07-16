### Installation

#### Ubuntu Linux

1. Install golang, libstfl0,  and libstfl-dev
```bash
# Must be root
apt get install golang
apt get install libstfl0
apt-get install libstfl-dev
```

2. Find $GOPATH
```bash
go env
```

3. For this example, $GOPATH=/home/e/go. Install go-stfl
```bash
cd $GOPATH/src
# Downlaod
git clone https://github.com/alethia7/go-stfl.git
cd go-stlf
# Install stfl.go package
go install
# Install example under $GOPATH/bin
cd example
go install
```

4. Run example
```bash
$GOPATH/bin/example ; reset
```

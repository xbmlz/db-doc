NAME=db-doc
BINDIR=bin
VERSION=$(shell git describe --tags || echo "unknown version")
BUILDTIME=$(shell date -u)
GOBUILD=CGO_ENABLED=0 go build -trimpath -ldflags '-X "github.com/xbmlz/db-doc/main.Version=$(VERSION)" \
		-X "github.com/xbmlz/db-doc/main.BuildTime=$(BUILDTIME)" \
		-w -s -buildid='

PLATFORM_LIST = \
	darwin-amd64 \
	linux-amd64

WINDOWS_ARCH_LIST = \
	windows-amd64

all: linux-amd64 darwin-amd64 windows-amd64 # Most used

darwin-amd64:
	GOARCH=amd64 GOOS=darwin $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

linux-amd64:
	GOARCH=amd64 GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

windows-amd64:
	GOARCH=amd64 GOOS=windows $(GOBUILD) -o $(BINDIR)/$(NAME)-$@.exe

gz_releases=$(addsuffix .gz, $(PLATFORM_LIST))
zip_releases=$(addsuffix .zip, $(WINDOWS_ARCH_LIST))

$(gz_releases): %.gz : %
	chmod +x $(BINDIR)/$(NAME)-$(basename $@)
	gzip -f -S -$(VERSION).gz $(BINDIR)/$(NAME)-$(basename $@)

$(zip_releases): %.zip : %
	zip -m -j $(BINDIR)/$(NAME)-$(basename $@)-$(VERSION).zip $(BINDIR)/$(NAME)-$(basename $@).exe

all-arch: $(PLATFORM_LIST) $(WINDOWS_ARCH_LIST)

releases: $(gz_releases) $(zip_releases)

lint:
	GOOS=darwin golangci-lint run ./...
	GOOS=windows golangci-lint run ./...
	GOOS=linux golangci-lint run ./...
	GOOS=freebsd golangci-lint run ./...
	GOOS=openbsd golangci-lint run ./...

clean:
	rm $(BINDIR)/*
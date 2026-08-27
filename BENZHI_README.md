# floatcell

floatcell 是一套矿物浮选槽工业过程控制系统，用于充气阀组、矿浆液位与过程联锁。

## Requirements

- Go 1.22+ (container image uses golang:1.22)
- `GOTOOLCHAIN=local` recommended on host when using a pinned toolchain

## Build

```bash
export GOTOOLCHAIN=local
go build ./...
```

## Run

```bash
export GOTOOLCHAIN=local
go run ./cmd/floatcell
```

Open http://127.0.0.1:8080/

Default ingest secret: `dev-ingest-secret`.

## Test

```bash
export GOTOOLCHAIN=local
go test ./... -count=1
```

## Docker (benzhi)

Must build **linux/amd64** and **linux/arm64**:

```bash
chmod +x build_benzhi_docker.sh
./build_benzhi_docker.sh floatcell linux/amd64
./build_benzhi_docker.sh floatcell linux/arm64
docker run -it floatcell:latest
# inside container:
export GOTOOLCHAIN=local
go version
go build ./...
go test ./... -count=1
```

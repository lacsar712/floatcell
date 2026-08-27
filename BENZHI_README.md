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
go run ./cmd/graindry
```

## Test

```bash
export GOTOOLCHAIN=local
go test ./... -count=1
```

## Docker (benzhi)

```bash
./build_benzhi_docker.sh
```

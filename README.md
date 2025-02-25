# 🌊 BitFlow

A small bit client for TCP protocol.

## Build

```sh
go build -o main *.go
```

## Configuration

```json
{
    "host": "localhost",
    "port": 9092,
    "messages": [
        "00000031004b00001f40a49e000c6b61666b612d746573746572000212756e6b6e6f776e2d746f7069632d71757a00000001ff00",
        "000000230012674a4f74d28b00096b61666b612d636c69000a6b61666b612d636c6904302e3100"
    ]
}
```

## Run

The program loads the configuration (`config.json`). It connects into provided server (the `host` and the `port`) and all the `messages`, one by one.

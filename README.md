<h1 align="center">WebOS Unclutter 🧼</h1>

> Remove the unwanted clutter from your WebOS TV.

## Why?

Modern TVs are great, but are full of bloatware and telemetry.<br>
This project aims to non-destructively block the telemetry of the webOS services you don't use.

## Supported TVs

This project has been built around the telemetry registered on an European LG C1 on <kbd><b>webOS 6.0</b></kbd>

Other webOS versions (or TVs in different regions) may communicate to different servers, and would probably require the addition of new rules in the blacklist.

If you have a different webOS version, please run this program with `allowed=true` in the `logging` configuration section and open an issue with the logged output.

## Installation

### Docker

Docker is the recommended way to run <b>WebOS Unclutter</b>.<br>
You can use the following command to build and run the container for your architecture:

```bash
make docker
```

### Pre-built Binaries

You can download the latest pre-built binaries for your OS and architecture from the [releases page](https://github.com/sonodima/webos-unclutter/releases)

### From Source

Building <b>WebOS Unclutter</b> from source is easy, but you have to make sure you have the following dependencies installed:

- Go 1.19
- Make

```bash
make build
```

You can also specify the `GOOS` and `GOARCH` environment variables to cross-compile for a different OS and architecture.

```bash
GOOS=windows GOARCH=arm64 make build
```

#### About macOS U2B

The <kbd>mach-o</kbd> binary format allows the creation of binaries that can be run on both Intel and Apple Silicon _(M1)_ processors, however, this is not supported when cross-compiling from a different OS.

For this reason, when building for macOS on a non-macOS machine, the build will only produce a <kbd>THIN</kbd> (single-architecture) binary for the current architecture.

## TV Setup

Setting up your TV is easy, and you don't need developer mode enabled.<br>
Go in the network settings of your TV, enable the manual network configuration and set the DNS address to the IP address of the machine running <b>WebOS Unclutter</b>.

## Configuration

<b>WebOS Unclutter</b> uses a configuration file to know what to remove from your TV.<br>
The configuration file is optional, and if not provided, the default configuration will be used.<br>

### Structure

```yaml
network:
  listen_port: 53
  resolver: 8.8.8.8:53

logging:
  blocked: true
  allowed: false

blocking:
  lg_smart_ad: true
  home_dashboard: true
  sports: true
  app_store: true
  internet_channels: true
  lg_iot: true
  amazon: true
  philips_hue: true
  software_updates: true
```

## Environment Variables

| Name     | Description                         | Default      |
| -------- | ----------------------------------- | ------------ |
| `CONFIG` | Name of the YAML configuration file | `config.yml` |

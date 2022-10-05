<h1 align="center">WebOS Unclutter 🧼</h1>

> Remove the unwanted clutter from your WebOS TV.

## Installation

### Docker

Docker is the easiest and recommended way to run <b>WebOS Unclutter</b>.<br>
You can use the following command to build and run the container:

```bash
make docker
```

### From Source

Building <b>WebOS Unclutter</b> from source is easy, but you have to make sure you have the following dependencies installed:

- Go 1.19
- Make

```bash
make build
```

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

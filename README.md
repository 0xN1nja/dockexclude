# Dockexclude
Manage your Docker Compose stack with the ability to exclude services.

## Why Dockexclude?
When working with large Docker Compose stacks, you often need to manage subsets of services without modifying your `docker-compose.yaml`. Dockexclude makes this simple by letting you exclude specific services from up, down, start, and stop operations.

## Installation
```bash
git clone https://github.com/0xN1nja/dockexclude.git
cd dockexclude
make build
```
After building, you can add the binary to your `/usr/local/bin`.

## Usage

### Basic Syntax

With exclusions
```bash
dockexclude <command> --exclude <service1>,<service2>,...
```

Without exclusions (acts like regular docker compose)
```bash
dockexclude <command>
```

### Commands

Start services in detached mode (excluding specified services)
```bash
dockexclude up --exclude service1,service2
```

Start previously created containers (excluding specified services)
```bash
dockexclude start --exclude service1,service2
```

Stop running containers (excluding specified services)
```bash
dockexclude stop --exclude service1,service2
```

Stop and remove containers (excluding specified services)
```bash
dockexclude down --exclude service1,service2
```

Run without exclusions (normal Docker Compose behavior)
```bash
dockexclude up
dockexclude down
dockexclude start
dockexclude stop
```

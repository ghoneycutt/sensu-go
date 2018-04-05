# Sensu 2.0

[![Build Status](https://travis-ci.org/sensu/sensu-go.svg?branch=master)](https://travis-ci.org/sensu/sensu-go)

Sensu is an open source monitoring tool for ephemeral infrastructure
and distributed applications. It is an agent based monitoring system
with built-in auto-discovery, making it very well-suited for cloud
environments. Sensu uses service checks to monitor service health and
collect telemetry data. It also has a number of well defined APIs for
configuration, external data input, and to provide access to Sensu's
data. Sensu is extremely extensible and is commonly referred to as
"the monitoring router".

## What is Sensu 2.0?

Sensu 2.0 is a complete rewrite of Sensu in Go, with new capabilities
and reduced operational overhead. It eliminates several sources of
friction for new and experienced Sensu users.

## Installation

Sensu 2.0 installer packages are available for a number of computing
platforms (e.g. Debian/Ubuntu, RHEL/Centos, etc), but the easiest way
to get started is with the official Docker image, sensuapp/sensu.

Please note the following installation steps to get Sensu up and
running on your local workstation with Docker.

_NOTE: the following instructions are based on Docker Community
Edition (CE), though they may be easily adapted for other container
platforms. Please download and install Docker CE before proceeding._

1. Start the Sensu 2.0 Backend process

```
$ docker run -d --name sensu-backend \
-p 2380:2380 -p 3000:3000 -p 8080:8080 -p 8081:8081 \
sensuapp/sensu-go:2.0.0-alpha sensu-backend start
```

2. Start the Sensu 2.0 Agent process

```
$ docker run -d --name sensu-agent --link sensu-backend \
sensuapp/sensu-go:2.0.0-alpha sensu-agent start \
--backend-url ws://sensu-backend:8081 \
--subscriptions workstation,docker
```

3. Download and install the Sensu 2.0 CLI tool

On macOS

```
$ latest=$(curl -s https://storage.googleapis.com/sensu-binaries/latest.txt)

$ curl -LO https://storage.googleapis.com/sensu-binaries/$latest/darwin/amd64/sensuctl

$ chmod +x sensuctl

$ sudo mv sensuctl /usr/local/bin/
```

On Debian/Ubuntu Linux

```
$ curl -s
https://packagecloud.io/install/repositories/sensu/nightly/script.deb.sh
| sudo bash

$ sudo apt-get install sensu-cli
```

On RHEL/CentOS Linux

```
$ curl -s
https://packagecloud.io/install/repositories/sensu/nightly/script.rpm.sh
| sudo bash

$ sudo yum install sensu-cli
```

4. Configure the Sensu 2.0 CLI tool

```
$ sensuctl configure
? Sensu Backend URL: http://127.0.0.1:8080
? Username: admin
? Password: P@ssw0rd!
? Organization: default
? Environment: default
? Preferred output format: tabular
```

5. List Sensu 2.0 Entities

```
$ sensuctl entity list
```

Congratulations! You have a local Sensu 2.0 deployment!

## Contributing/Development

To make a good faith effort to ensure the criteria of the MIT License
are met, Sensu Inc. requires the Developer Certificate of Origin (DCO)
process to be followed.

For guidelines on how to contribute to this project and more
information on the DCO, please see [CONTRIBUTING.md](CONTRIBUTING.md).

## Protobuf

### Overview

We are using the version **proto3** of the protocol buffers language. Here are some useful resources:

[To learn more about protocol buffers](https://developers.google.com/protocol-buffers/docs/overview)

[The proto3 language guide](https://developers.google.com/protocol-buffers/docs/proto3)


### Installation

Install the protobuf compiler since we don't use the one that golang uses.
```
brew install protobuf
```
Otherwise, see the **for non-C++ users** [instructions here.](https://github.com/google/protobuf#protocol-compiler-installation)

### Quick Start

Once you make a change to any `*.proto` file within the **types** package, you will need to regenerate the associated `*.pb.go` file. To do so, simply run `go generate` on the package.

## Dependencies

Sensu uses [golang/dep](https://github.com/golang/dep) for managing its
dependencies.

### Usage

Running the following will install `dep` (if it is not already) and pull all
required dependencies.

```shell
./build.sh deps
```

Later, if you would like to add a dependency, run:

```shell
dep ensure -add https://my-repo.com/my/dep
```

If you would like to update a dependency, run:

```shell
dep ensure -update https://my-repo.com/my/dep
```

When you would like to remove a dependency, remove the it from `Gopkg.toml` and
then run:

```shell
dep prune
```

### Further Reading

- [The Saga of Go Dependency Management](https://blog.gopheracademy.com/advent-2016/saga-go-dependency-management/)
- [`dep` Usage](https://github.com/golang/dep#usage)

## Building

### Docker

The simplest way to the build Sensu is with the `sensu-go-build` image. The
image contains all the required tools to build the agent, backend and sensuctl.

```sh
docker pull sensuapp/sensu-go-build
docker run -it -e GOOS=darwin -v `pwd`:/go/src/github.com/sensu/sensu-go --entrypoint='/go/src/github.com/sensu/sensu-go/build.sh' sensuapp/sensu-go-build
```

If you would like to build for different platforms and architectures use GOOS
and GOARCH env variables. See [Optional environment variables](https://golang.org/doc/install/source#environment) for more.

When complete your binaries will be present in the `target` directory.

### Manually

First ensure that you have the required tools installed to build the programs.

* Ensure that you have the Go tools installed and your environment configured.
  If not follow the official
  [Install the Go tools](https://golang.org/doc/install#install) guide.
* When building the Sensu backend you will need NodeJS and Yarn installed so
  that the web UI may be included in the binary. Follow
  [Installing Node.js](https://nodejs.org/en/download/package-manager/) and
  [Yarn Installation](https://yarnpkg.com/en/docs/install) for installation
  instructions for your platform.

Once all the tools are installed you are now ready to use the build script. To
build the Sensu backend, agent and sensuctl, run:

```sh
./build.sh build
```

Each product can built separately, with one of the following:

```sh
./build.sh build_agent
./build.sh build_backend
./build.sh build_cli
```

By default the web UI is built along side and bundled into the backend, as this
can be a time intensive process, we provide an escape hatch. Use the `dev` tag
to avoid building the web UI.

```sh
./build.sh build_backend -tags dev
```

## Testing

Install dependencies:

```shell
./build.sh deps
```

Run test suites:

```shell
./build.sh ci
```

Run end-to-end tests:

```shell
./build.sh e2e

# To run a specific test:

./build.sh e2e -run TestRBAC

# To prevent tests from running in parallel:

./build.sh e2e -parallel 1
```

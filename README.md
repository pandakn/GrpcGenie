# GrpcGenie

`GrpcGenie` is a command-line tool that generates gRPC handler files from .proto definitions. It simplifies the process of creating gRPC server implementations by automatically generating boilerplate code.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [How It Works](#how-it-works)
- [License](#license)

## Installation

To install GrpcGenie, use the following command (Deploying in the future. Stay tuned! 📅🚀):

```bash
git clone https://github.com/pandakn/GrpcGenie.git

cd GrpcGenie

go build -o grpcgenie cmd/grpcgenie/main.go
# or
make build
```

## Usage

GrpcGenie is used via the command line. Here's the basic syntax:

#### grpcgenie [flags]

##### Flags

<!-- - -p, --proto: Path to the .proto file (required)
- -o, --output: Path to the output Go handler file (required)
- -d, --package: Name of the Go package (required)
- -g, --go-package-path: Path to the Go package (required)
- -r, --grpc-package: Name of the gRPC package (required) -->

- `-p, --proto`: Required. Tell GrpcGenie where your .proto file lives.
- `-o, --output`: Required. Where do you want the generated Go handler file to be saved?
- `-d, --package`: Required. Name for your Go package (think of it like a label).
- `-g, --go-package-path`: Required. The path to your Go package (e.g., github.com/your-username/your-project).
- `-r, --grpc-package`: Required. Name of the gRPC package defined in your .proto file.

#### Example

```bash
./grpcgenie --proto example/hello/hellopd/hello.proto \
	--output example/hello/hellohandler/hello_grpc_handler.go \
	--package hellohandler \
	--go-package-path github.com/pandakn/GrpcGenie/example/hello \
	--grpc-package hellopd
```

```
📦
example/
    └── hello/
        ├── hellohandler/
        │   └── hello_grpc_handler.go // from grpcgennie
        └── hellopd/
            ├── hello.pb.go
            ├── hello.proto
            └── hello_grpc.pb.go
```

## How It Works

- GrpcGenie parses the specified .proto file to extract service and method information.
- It uses this information to generate a Go file containing handler stuff for each gRPC method.
- The generated file includes all necessary imports and a struct that implements the gRPC service interface.

## License

Licensed under [MIT License](./LICENSE)

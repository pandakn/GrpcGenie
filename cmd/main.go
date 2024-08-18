package main

import (
	"os"

	"github.com/pandakn/GrpcGenie/internal/generator"
	"github.com/pandakn/GrpcGenie/internal/genielogger"
	"github.com/pandakn/GrpcGenie/internal/protoutil"

	"github.com/spf13/cobra"
)

var (
	protoPath       string
	outputPath      string
	packageName     string
	grpcPackageName string
	goPackagePath   string
	log             genielogger.Logger
)

var rootCmd = &cobra.Command{
	Use:   "grpcgenie",
	Short: "GrpcGenie generates gRPC handler files from .proto definitions",
	Run: func(cmd *cobra.Command, args []string) {
		services, err := protoutil.ParseProtoGetServices(protoPath)
		if err != nil {
			log.Error("Error: %v\n", err)
			return
		}

		// Prepare data for template
		methods := services[0].Methods
		serviceName := services[0].Name

		data := generator.TemplateData{
			PackageName:     packageName,
			GrpcPackageName: grpcPackageName,
			ServiceName:     serviceName,
			GoPackagePath:   goPackagePath,
			Methods:         methods,
		}

		// Generate handler file
		err = generator.GenerateHandler(protoPath, outputPath, data)
		if err != nil {
			log.Error("Error generating handler file: %v\n", err)
			os.Exit(1)
		}

		log.Info("Handler file generated successfully!")
	},
}

func init() {
	// Define flags
	rootCmd.Flags().StringVarP(&protoPath, "proto", "p", "", "Path to the .proto file (required)")
	rootCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Path to the output Go handler file (required)")
	rootCmd.Flags().StringVarP(&packageName, "package", "d", "", "Name of the Go package (required)")
	rootCmd.Flags().StringVarP(&goPackagePath, "go-package-path", "g", "", "Path to the Go package (required)")
	rootCmd.Flags().StringVarP(&grpcPackageName, "grpc-package", "r", "", "Name of the gRPC package (required)")

	rootCmd.MarkFlagRequired("proto")
	rootCmd.MarkFlagRequired("output")
	rootCmd.MarkFlagRequired("package")
	rootCmd.MarkFlagRequired("go-package-path")
	rootCmd.MarkFlagRequired("grpc-package")
}

func main() {
	log = genielogger.NewGenieLogger()

	if err := rootCmd.Execute(); err != nil {
		log.Fatal("Command execution failed: %v\n", err)
		os.Exit(1)
	}
}

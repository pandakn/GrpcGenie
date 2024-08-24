package protoutil

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/pandakn/GrpcGenie/internal/generator"
)

type ServiceInfo struct {
	Name    string             `json:"name"`
	Methods []generator.Method `json:"methods"`
}

func ParseProtoGetServices(protoFilePath string) ([]ServiceInfo, error) {
	protoData, err := readProtoFile(protoFilePath)
	if err != nil {
		return nil, err
	}

	services, err := extractServices(protoData)
	if err != nil {
		return nil, err
	}

	return services, nil
}

// readProtoFile handles file reading.
func readProtoFile(protoFilePath string) (string, error) {
	data, err := os.ReadFile(protoFilePath)
	if err != nil {
		return "", fmt.Errorf("error reading proto file: %v", err)
	}
	return string(data), nil
}

// extractServices handles parsing of services and methods.
func extractServices(protoData string) ([]ServiceInfo, error) {
	if strings.TrimSpace(protoData) == "" {
		return nil, errors.New("empty protobuf data string")
	}

	servicePattern := regexp.MustCompile(`service\s+(\w+)\s*{((?:\s*rpc\s+(\w+)\s*\((\w+)\)\s*returns\s*\((\w+)\)\s*(?:;|\{?\s*\}?)\s*)+)}`)
	matches := servicePattern.FindAllStringSubmatch(protoData, -1)

	services := make([]ServiceInfo, 0)

	for _, match := range matches {
		serviceName := match[1]
		rpcMethods := match[2]

		methods, err := extractMethods(rpcMethods)
		if err != nil {
			return nil, err
		}

		serviceInfo := ServiceInfo{
			Name:    serviceName,
			Methods: methods,
		}
		services = append(services, serviceInfo)
	}

	if len(services) == 0 {
		return nil, errors.New("no valid rpc methods found")
	}

	return services, nil
}

// extractMethods handles parsing of methods within a service.
func extractMethods(rpcMethods string) ([]generator.Method, error) {

	if strings.TrimSpace(rpcMethods) == "" {
		return nil, errors.New("empty rpc methods string")
	}

	methodPattern := regexp.MustCompile(`rpc\s+(\w+)\s*\((\w+)\)\s*returns\s*\((\w+)\)`)
	methodMatches := methodPattern.FindAllStringSubmatch(rpcMethods, -1)

	methods := make([]generator.Method, 0)

	for _, methodMatch := range methodMatches {
		methodName := methodMatch[1]
		inputType := methodMatch[2]
		outputType := methodMatch[3]

		methodInfo := generator.Method{
			Name:       methodName,
			InputType:  inputType,
			OutputType: outputType,
		}
		methods = append(methods, methodInfo)
	}

	if len(methods) == 0 {
		return nil, errors.New("no valid rpc methods found")
	}

	return methods, nil
}

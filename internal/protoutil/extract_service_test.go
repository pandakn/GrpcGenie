package protoutil

import (
	"reflect"
	"testing"

	"github.com/pandakn/GrpcGenie/internal/generator"
)

func TestExtractServices(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []ServiceInfo
		hasError bool
	}{
		{
			name: "Single service with single method",
			input: `
service UserService {
    rpc GetUser(GetUserRequest) returns (GetUserResponse);
}`,
			expected: []ServiceInfo{
				{
					Name: "UserService",
					Methods: []generator.Method{
						{Name: "GetUser", InputType: "GetUserRequest", OutputType: "GetUserResponse"},
					},
				},
			},
			hasError: false,
		},
		{
			name: "Multiple services with multiple methods",
			input: `
service UserService {
    rpc GetUser(GetUserRequest) returns (GetUserResponse);
    rpc CreateUser(CreateUserRequest) returns (CreateUserResponse);
}
service ProductService {
    rpc GetProduct(GetProductRequest) returns (GetProductResponse);
}`,
			expected: []ServiceInfo{
				{
					Name: "UserService",
					Methods: []generator.Method{
						{Name: "GetUser", InputType: "GetUserRequest", OutputType: "GetUserResponse"},
						{Name: "CreateUser", InputType: "CreateUserRequest", OutputType: "CreateUserResponse"},
					},
				},
				{
					Name: "ProductService",
					Methods: []generator.Method{
						{Name: "GetProduct", InputType: "GetProductRequest", OutputType: "GetProductResponse"},
					},
				},
			},
			hasError: false,
		},
		{
			name:     "Empty input",
			input:    "",
			expected: nil,
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := extractServices(tt.input)
			if tt.hasError {
				if err == nil {
					t.Errorf("Expected an error, but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("Expected %+v, but got %+v", tt.expected, result)
				}
			}
		})
	}
}

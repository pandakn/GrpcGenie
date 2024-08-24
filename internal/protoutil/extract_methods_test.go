package protoutil

import (
	"reflect"
	"testing"

	"github.com/pandakn/GrpcGenie/internal/generator"
)

func TestParseProtoGetService_ExtractMethodsSuccess(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []generator.Method
		hasError bool
	}{
		{
			name: "Single method",
			input: `
    rpc GetUser(GetUserRequest) returns (GetUserResponse);`,
			expected: []generator.Method{
				{Name: "GetUser", InputType: "GetUserRequest", OutputType: "GetUserResponse"},
			},
			hasError: false,
		},
		{
			name: "Multiple methods",
			input: `
    rpc GetUser(GetUserRequest) returns (GetUserResponse);
    rpc CreateUser(CreateUserRequest) returns (CreateUserResponse);`,
			expected: []generator.Method{
				{Name: "GetUser", InputType: "GetUserRequest", OutputType: "GetUserResponse"},
				{Name: "CreateUser", InputType: "CreateUserRequest", OutputType: "CreateUserResponse"},
			},
			hasError: false,
		},
		{
			name:     "Empty input",
			input:    "",
			expected: nil,
			hasError: true,
		},
		{
			name:     "Whitespace only input",
			input:    "    \n\t",
			expected: nil,
			hasError: true,
		},
		{
			name:     "Invalid input",
			input:    "not a valid rpc method",
			expected: nil,
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := extractMethods(tt.input)
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

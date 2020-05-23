// Error output from IoT-Agent-JSON
package iotagent

const (
	WrongSyntaxError = "WRONG_SYNTAX"
)

type (
	// IoTError is a struct for error message that IoT Agent JSON returns.
	IotError struct {
		Name    string `json:"name"`
		Message string `json:"message"`
	}
)

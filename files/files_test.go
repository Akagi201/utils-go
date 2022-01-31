package files_test

import (
	"io"
	"os"
	"testing"

	"github.com/Akagi201/utils-go/files"
)

// Test_ReadPropertiesFile creates a test properties file and reads it into memory
func Test_ReadPropertiesFile(t *testing.T) {

	//
	// Create the test properties file
	//
	filename := "test.properties"
	file, err := os.Create(filename)
	if err != nil {
		t.Fatal("Error creating test properties file", err)
	}
	_, err = io.WriteString(file, `
# This is a test properties file
maxRows = 20
somethingEnabled = true
error.message = There was an error
`)
	if err != nil {
		t.Fatal("Error writing test properties file", err)
	}
	file.Close()

	//
	// Read the properties from the file
	//
	var properties map[string]string
	var ok bool
	var value string
	properties, err = files.ReadPropertiesFile("test.properties")
	if err != nil {
		t.Fatal("Error reading test properties file", err)
	}
	if _, ok = properties["maxRows"]; !ok {
		t.Fail()
	}
	if _, ok = properties["error.message"]; !ok {
		t.Fail()
	}
	if value, ok = properties["error.message"]; !ok || value != "There was an error" {
		t.Fatalf("value is wrong for error.message: '%s'", properties["error.message"])
	}

}

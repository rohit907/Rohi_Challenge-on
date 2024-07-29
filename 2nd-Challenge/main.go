package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Input structure to represent the incoming JSON
type Input map[string]interface{}

// Output structure to represent the transformed output
type Output struct {
	Map1    map[string]interface{} `json:"map_1"`
	Number1 float64                `json:"number_1"`
	String1 string                 `json:"string_1"`
	String2 int64                  `json:"string_2"`
}

func main() {
	inputJSON := `{
		"number_1": {
			"N": "1.50"
		},
		"string_1": {
			"S": "784498 "
		},
		"string_2": {
			"S": "2014-07-16T20:55:46Z"
		},
		"map_1": {
			"M": {
				"bool_1": {
					"BOOL": "truthy"
				},
				"null_1": {
					"NULL": "true"
				},
				"list_1": {
					"L": [
						{
							"S": ""
						},
						{
							"N": "011"
						},
						{
							"N": "5215s"
						},
						{
							"BOOL": "f"
						},
						{
							"NULL": "0"
						}
					]
				}
			}
		},
		"list_2": {
			"L": "noop"
		},
		"list_3": {
			"L": [
				"noop"
			]
		},
		"": {
			"S": "noop"
		}
	}`

	// Parse the input JSON
	var input Input
	err := json.Unmarshal([]byte(inputJSON), &input)
	if err != nil {
		log.Fatalf("Failed to parse input JSON: %v", err)
	}

	// Transform the input JSON
	output := transform(input)

	// Marshal output to JSON and print
	outputJSON, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal output JSON: %v", err)
	}
	fmt.Println(string(outputJSON))
}

// transform function processes the input data according to transformation rules
func transform(data map[string]interface{}) []Output {
	var results []Output

	// Initialize the output structure
	result := Output{
		Map1: make(map[string]interface{}),
	}

	for key, value := range data {
		sanitizedKey := sanitize(key)
		if sanitizedKey == "" {
			continue
		}

		switch v := value.(type) {
		case map[string]interface{}:
			//omit fields with empty keys.
			if len(v) == 0 {
				continue
			}
			if num, ok := v["N"]; ok {
				if number, err := transformNumber(num); err == nil {
					if sanitizedKey == "number_1" {
						result.Number1 = number
					}
				}
			} else if str, ok := v["S"]; ok {
				if sanitizedKey == "string_1" {
					result.String1 = sanitize(str)
				} else if sanitizedKey == "string_2" {
					if timestamp, err := parseRFC3339Date(str); err == nil {
						result.String2 = timestamp
					}
				}
			} else if m, ok := v["M"]; ok {
				if mapData, ok := m.(map[string]interface{}); ok {
					for mapKey, mapValue := range mapData {
						sanitizedMapKey := sanitize(mapKey)
						if sanitizedMapKey == "" {
							continue
						}
						if boolValue, ok := mapValue.(map[string]interface{})["BOOL"]; ok {
							if b, err := transformBoolean(boolValue); err == nil {
								result.Map1[sanitizedMapKey] = b
							}
						} else if nullValue, ok := mapValue.(map[string]interface{})["NULL"]; ok {
							if n, err := parseNull(nullValue); err == nil {
								result.Map1[sanitizedMapKey] = n
							}
						} else if listValue, ok := mapValue.(map[string]interface{})["L"]; ok {
							if list, ok := listValue.([]interface{}); ok {
								transformedList := transformList(list)
								if len(transformedList) > 0 {
									result.Map1[sanitizedMapKey] = transformedList
								}
							}
						}
					}
				}
			}
		}
	}

	// Add the result to the output slice
	results = append(results, result)
	return results
}

// sanitize for trailing and leading whitespace.
func sanitize(value interface{}) string {
	str, _ := value.(string)
	return strings.TrimSpace(str)
}

// transformNumber to numeric value, sanitize the value of trailing and leading whitespace before processing
func transformNumber(value interface{}) (float64, error) {
	str := sanitize(value)
	str = strings.TrimLeft(str, "0") // Strip leading zeros
	if str == "" {
		return 0, fmt.Errorf("invalid number")
	}
	return strconv.ParseFloat(str, 64)
}

// transformBoolean parses a boolean value from various string formats
func transformBoolean(value interface{}) (bool, error) {
	str := sanitize(value)
	switch str {
	case "1", "t", "T", "TRUE", "true", "True":
		return true, nil
	case "0", "f", "F", "FALSE", "false", "False":
		return false, nil
	default:
		return false, fmt.Errorf("invalid boolean")
	}
}

// parseNull determines if a null value should be represented as nil or omitted
func parseNull(value interface{}) (interface{}, error) {
	str := sanitize(value)
	if str == "1" || str == "t" || str == "T" || str == "TRUE" || str == "true" || str == "True" {
		return nil, nil
	}
	return nil, fmt.Errorf("invalid null")
}

// parseRFC3339Date converts an RFC3339 date string to a Unix timestamp
func parseRFC3339Date(value interface{}) (int64, error) {
	str := sanitize(value)
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return 0, fmt.Errorf("invalid date format")
	}
	return t.Unix(), nil
}

// transformList processes the list items and transforms them to appropriate data types
func transformList(list []interface{}) []interface{} {
	var result []interface{}
	for _, item := range list {
		switch v := item.(type) {
		case map[string]interface{}:
			if n, ok := v["N"]; ok {
				if number, err := transformNumber(n); err == nil {
					result = append(result, int(number)) // Store as int
				}
			} else if s, ok := v["S"]; ok {
				if str := sanitize(s); str != "" {
					result = append(result, str)
				}
			} else if b, ok := v["BOOL"]; ok {
				if booleanValue, err := transformBoolean(b); err == nil {
					result = append(result, booleanValue)
				}
			} else if n, ok := v["NULL"]; ok {
				if nullValue, err := parseNull(n); err == nil && nullValue == nil {
					result = append(result, nil)
				}
			}
		}
	}
	return result
}
// go run main.go
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Define constants for data types
const (
	StringType  = "S"
	NumberType  = "N"
	BooleanType = "BOOL"
	NullType    = "NULL"
	ListType    = "L"
	MapType     = "M"
)

// JSONInput represents the structure of the input JSON
type JSONInput map[string]map[string]interface{}

// JSONOutput represents the structure of the output JSON
type JSONOutput []map[string]interface{}

// TransformInput transforms the input JSON according to the criteria
func TransformInput(input JSONInput) JSONOutput {
	var output JSONOutput

	for key, value := range input {
		// Skip empty keys
		if key == "" {
			continue
		}

		transformedValue := make(map[string]interface{})

		for typeKey, typeValue := range value {
			switch typeKey {
			case StringType:
				transformedValue[key] = strings.TrimSpace(typeValue.(string))
			case NumberType:
				numberStr := strings.TrimSpace(typeValue.(string))
				// Remove leading zeros
				number, err := strconv.ParseFloat(numberStr, 64)
				if err == nil {
					transformedValue[key] = number
				}
			case BooleanType:
				boolStr := strings.TrimSpace(strings.ToLower(typeValue.(string)))
				switch boolStr {
				case "1", "t", "true":
					transformedValue[key] = true
				case "0", "f", "false":
					transformedValue[key] = false
				}
			case NullType:
				nullStr := strings.TrimSpace(strings.ToLower(typeValue.(string)))
				if nullStr == "1" || nullStr == "t" || nullStr == "true" {
					transformedValue[key] = nil
				}
			case ListType:
				list, ok := typeValue.([]interface{})
				if ok {
					transformedList := make([]interface{}, 0)
					for _, listItem := range list {
						switch listItemType := listItem.(type) {
						case map[string]interface{}:
							transformedList = append(transformedList, TransformMap(listItemType))
						case string:
							if listItemType != "" {
								transformedList = append(transformedList, strings.TrimSpace(listItemType))
							}
						}
					}
					if len(transformedList) > 0 {
						transformedValue[key] = transformedList
					}
				}
			case MapType:
				mapValue, ok := typeValue.(map[string]interface{})
				if ok {
					transformedValue[key] = TransformMap(mapValue)
				}
			}
		}

		if len(transformedValue) > 0 {
			output = append(output, transformedValue)
		}
	}

	// Sort the output lexicographically
	sort.Slice(output, func(i, j int) bool {
		return sortedKeys(output[i]) < sortedKeys(output[j])
	})

	return output
}

// TransformMap recursively transforms a map according to the criteria
func TransformMap(input map[string]interface{}) map[string]interface{} {
	output := make(map[string]interface{})

	for key, value := range input {
		switch valueType := value.(type) {
		case map[string]interface{}:
			output[key] = TransformMap(valueType)
		case string:
			output[key] = strings.TrimSpace(valueType)
		}
	}

	return output
}

// sortedKeys returns a string containing the lexicographically sorted keys of a map
func sortedKeys(input map[string]interface{}) string {
	keys := make([]string, 0, len(input))
	for key := range input {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return strings.Join(keys, ",")
}

func main() {

	inputData := readJSON(os.Args[1])

	outputData := TransformInput(inputData)

	// Convert RFC3339 time strings to Unix epoch
	for _, item := range outputData {
		if strTime, ok := item["string_2"].(string); ok {
			t, err := time.Parse(time.RFC3339, strTime)
			if err == nil {
				item["string_2"] = t.Unix()
			}
		}
	}

	// Print the output JSON
	outputJSON, err := json.MarshalIndent(outputData, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling output JSON:", err)
		return
	}

	fmt.Println(string(outputJSON))
}

func readJSON(filename string) JSONInput {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer file.Close()

	var data JSONInput
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return data
}

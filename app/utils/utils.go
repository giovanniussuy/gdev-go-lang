package utils

import (
	"encoding/json"
)

func ConvertStructToJsonString[T any](genericType T) string {
	json, _ := json.Marshal(genericType)

	return string(json)
}

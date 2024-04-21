package conversion

import (
	"fmt"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	var floats []float64
	for stringIndex, stringValue := range strings {
		floatValue, err := strconv.ParseFloat(stringValue, 64)

		if err != nil {
			return nil, fmt.Errorf("error parsing line %d: %v", stringIndex, err)
		}

		floats = append(floats, floatValue)
	}
	return floats, nil
}

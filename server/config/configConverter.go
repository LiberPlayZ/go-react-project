package config

import "strconv"

func GetIntEnv(value string) (int, error) {

	parseValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return parseValue, nil
}

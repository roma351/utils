package utils

func GetIntOrNil(value int) *int {
	if value == 0 {
		return nil
	} else {
		return &value
	}
}

func GetInt32OrNil(value int32) *int32 {
	if value == 0 {
		return nil
	} else {
		return &value
	}
}

func GetInt64OrNil(value int64) *int64 {
	if value == 0 {
		return nil
	} else {
		return &value
	}
}

func GetStringOrNil(value string) *string {
	if value == "" {
		return nil
	} else {
		return &value
	}
}

func GetIntOrInt(value *int) int {
	if value == nil {
		return 0
	} else {
		return *value
	}
}
func GetInt32OrInt32(value *int32) int32 {
	if value == nil {
		return 0
	} else {
		return *value
	}
}
func GetInt64OrInt64(value *int64) int64 {
	if value == nil {
		return 0
	} else {
		return *value
	}
}

func GetStringOrString(value *string) string {
	if value == nil {
		return ""
	} else {
		return *value
	}
}

package hw09_struct_validator //nolint:golint,stylecheck

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrorNotStruct            = errors.New("тип не является структурой")
	ErrorValue                = errors.New("значение не соответствует правилу")
	ErrorUnknownRule          = errors.New("неизвестное правило")
	ErrorRule                 = errors.New("ошибка в правилах валидации")
	ErrorRuleValueIsNotNumber = errors.New("значение не является числом")
	ErrorRegexp               = func(err error, reg string) error {
		return fmt.Errorf("некорректное регулярное выражение: %v, %w", reg, err)
	}
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var resultString string
	for _, err := range v {
		resultString += fmt.Sprintf("error in field: %v, %v\n", err.Field, err.Err)
	}
	return resultString
}

func Validate(v interface{}) error {
	val := reflect.ValueOf(v)

	// Если тип не структура
	if val.Kind() != reflect.Struct {
		return ErrorNotStruct
	}

	var vErrors ValidationErrors

	for i := 0; i < val.NumField(); i++ {
		tag := val.Type().Field(i).Tag.Get("validate")
		if len(tag) == 0 {
			continue
		}

		valueOf := val.Field(i)
		if !val.Field(i).CanInterface() {
			continue
		}

		fieldName := val.Type().Field(i).Name
		var values []interface{}
		// Собираем значения полей из разных структур
		values = collectValues(values, valueOf)

		// Для всех значений проводим валидацию
		for _, value := range values {
			vErrors = append(vErrors, validate(value, fieldName, tag)...)
		}
	}

	if len(vErrors) == 0 {
		return nil
	}

	return vErrors
}

// Собирает значения полей из разных структур.
func collectValues(values []interface{}, value reflect.Value) []interface{} {
	switch value.Kind() {
	case reflect.Slice, reflect.Array:
		for j := 0; j < value.Len(); j++ {
			values = append(values, value.Index(j).Interface())
		}
	case reflect.Map:
		for _, key := range value.MapKeys() {
			values = append(values, value.MapIndex(key).Interface())
		}
	default:
		values = append(values, value.Interface())
	}

	return values
}

// Выполняет валидацию масива значений по типу структуры и набору правил.
func validate(value interface{}, name string, rules string) ValidationErrors {
	var vErr ValidationErrors

	for _, rule := range strings.Split(rules, "|") {
		// Чтобы обработать несколько правил
		rulesArr := strings.SplitN(rule, ":", 2)

		if len(rulesArr) != 2 {
			vErr = append(vErr, appendErr(name, ErrorRule, rule))
			continue
		}

		rName := rulesArr[0]
		rVal := rulesArr[1]
		ok, err := validateValue(toString(value), rName, rVal)
		if err != nil {
			vErr = append(vErr, appendErr(name, err))
			continue
		}

		if !ok {
			vErr = append(vErr, appendErr(name, ErrorValue, rule, " (value=", toString(value), ")"))
		}
	}

	if len(vErr) == 0 {
		return nil
	}

	return vErr
}

// Валидация конкретного значения по правилу.
func validateValue(val string, rName string, rVal string) (bool, error) {
	switch rName {
	case "len":
		ruleVal, err := strconv.ParseFloat(rVal, 64)
		if err != nil {
			return false, ErrorRuleValueIsNotNumber
		}

		return len([]rune(val)) == int(ruleVal), nil
	case "min":
		val, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return false, ErrorRuleValueIsNotNumber
		}

		ruleVal, err := strconv.ParseFloat(rVal, 64)
		if err != nil {
			return false, ErrorRuleValueIsNotNumber
		}
		return val >= ruleVal, nil
	case "regexp":
		reg, err := regexp.Compile(rVal)
		if err != nil {
			return false, ErrorRegexp(err, rVal)
		}
		return reg.MatchString(val), nil
	case "in":
		for _, s := range strings.Split(rVal, ",") {
			if s == val {
				return true, nil
			}
		}
		return false, nil
	}

	return false, ErrorUnknownRule
}

// Форматирует ошибку.
func appendErr(name string, err error, s ...string) ValidationError {
	return ValidationError{
		Field: name,
		Err:   fmt.Errorf("%w, %s", err, s),
	}
}

func toString(i interface{}) string {
	format := "%s"
	switch i.(type) {
	case float32, float64:
		format = "%f"
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		format = "%d"
	}
	return fmt.Sprintf(format, i)
}

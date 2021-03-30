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
	ErrorNotStruct            = errors.New("Тип не является структурой")
	ErrorValue                = errors.New("Значение не соответствует правилу")
	ErrorUnknownRule          = errors.New("Неизвестное правило")
	ErrorRule                 = errors.New("Ошибка в правилах валидации")
	ErrorRuleValueIsNotNumber = errors.New("Значение не является числом")
)

type ValidationError struct {
	Field string
	Err   error
}

var ruleFunctions = map[string]func(fieldValue, ruleValue string) (bool, error){
	"len": func(fieldValue, ruleValue string) (bool, error) {
		mustLen, err := strconv.ParseFloat(ruleValue, 64)
		if err != nil {
			return false, ErrorRuleValueIsNotNumber
		}
		return len([]rune(fieldValue)) == int(mustLen), nil
	},
	"min": func(fieldValue, ruleValue string) (bool, error) {
		val, err := strconv.ParseFloat(fieldValue, 64)
		if err != nil {
			return false, ErrorRuleValueIsNotNumber
		}
		mustMin, err := strconv.ParseFloat(ruleValue, 64)
		if err != nil {
			return false, ErrorRuleValueIsNotNumber
		}
		return val >= mustMin, nil
	},
	"max": func(fieldValue, ruleValue string) (bool, error) {
		val, err := strconv.ParseFloat(fieldValue, 64)
		if err != nil {
			return false, ErrorRuleValueIsNotNumber
		}
		mustMax, err := strconv.ParseFloat(ruleValue, 64)
		if err != nil {
			return false, ErrorRuleValueIsNotNumber
		}
		return val <= mustMax, nil
	},
	"regexp": func(fieldValue, ruleValue string) (bool, error) {
		reg, err := regexp.Compile(ruleValue)
		if err != nil {
			return false, fmt.Errorf("rule value is not valid regexp: %w", err)
		}
		return reg.MatchString(fieldValue), nil
	},
	"in": func(fieldValue, ruleValue string) (bool, error) { //nolint:unparam
		for _, s := range strings.Split(ruleValue, ",") {
			if s == fieldValue {
				return true, nil
			}
		}
		return false, nil
	},
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var resultString string
	for _, error := range v {
		resultString += fmt.Sprintf("error in field: %v, %v\n", error.Field, error.Err)
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
		values = collectValues(values, valueOf)

		for _, value := range values {
			vErrors = append(vErrors, validate(value, fieldName, tag)...)
		}
	}

	if len(vErrors) == 0 {
		return nil
	}

	return vErrors
}

// Собирает значения полей из разных структур
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

// Выполняет валидацию масива значений по типу структуры и набору правил
func validate(value interface{}, name string, rules string) ValidationErrors {
	var vErr ValidationErrors

	for _, rule := range strings.Split(rules, "|") {
		// Чтобы обработатть несколько правил
		rulesArr := strings.SplitN(rule, ":", 2)

		if len(rulesArr) != 2 {
			vErr = append(vErr, appendErr(name, ErrorRule, rule))
			continue
		}

		rName := rulesArr[0]
		rVal := rulesArr[1]
		//ruleFunction, ok := ruleFunctions[rName]
		_, ok := validateValue(rName, rVal)

		if !ok {
			vErr = append(vErr, appendErr(name, ErrorUnknownRule, rName))
			continue
		}
		ok, err := ruleFunction(toString(value), rVal)
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

func validateValue(name string, val string, rule string) (bool, error) {
	switch name {
	case "len":
		ruleVal, err := strconv.ParseFloat(rule, 64)
		if err != nil {
			return false, ErrorRuleValueIsNotNumber
		}

		return len([]rune(val)) == int(ruleVal), nil
	case "min":
		val, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return false, ErrorRuleValueIsNotNumber
		}

		ruleVal, err := strconv.ParseFloat(rule, 64)
		if err != nil {
			return false, ErrorRuleValueIsNotNumber
		}
		return val >= ruleVal, nil

	}

	"len": func(fieldValue, ruleValue string) (bool, error) {
		mustLen, err := strconv.ParseFloat(ruleValue, 64)
		if err != nil {
			return false, ErrorRuleValueIsNotNumber
		}
		return len([]rune(fieldValue)) == int(mustLen), nil
	},
		"min": func(fieldValue, ruleValue string) (bool, error) {
		val, err := strconv.ParseFloat(fieldValue, 64)
		if err != nil {
			return false, ErrorRuleValueIsNotNumber
		}
		mustMin, err := strconv.ParseFloat(ruleValue, 64)
		if err != nil {
			return false, ErrorRuleValueIsNotNumber
		}
		return val >= mustMin, nil
	},
		"max": func(fieldValue, ruleValue string) (bool, error) {
		val, err := strconv.ParseFloat(fieldValue, 64)
		if err != nil {
			return false, ErrorRuleValueIsNotNumber
		}
		mustMax, err := strconv.ParseFloat(ruleValue, 64)
		if err != nil {
			return false, ErrorRuleValueIsNotNumber
		}
		return val <= mustMax, nil
	},
		"regexp": func(fieldValue, ruleValue string) (bool, error) {
		reg, err := regexp.Compile(ruleValue)
		if err != nil {
			return false, fmt.Errorf("rule value is not valid regexp: %w", err)
		}
		return reg.MatchString(fieldValue), nil
	},
		"in": func(fieldValue, ruleValue string) (bool, error) { //nolint:unparam
		for _, s := range strings.Split(ruleValue, ",") {
			if s == fieldValue {
				return true, nil
			}
		}
		return false, nil
	},
}

// Форматирует ошибку
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

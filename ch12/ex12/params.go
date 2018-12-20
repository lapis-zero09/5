package params

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var Regexp map[string]*regexp.Regexp

func init() {
	// ref: https://www.regular-expressions.info/creditcard.html
	Regexp = map[string]*regexp.Regexp{
		"creditCard":       regexp.MustCompile(`\d{13,16}`),
		"visa":             regexp.MustCompile(`^4[0-9]{12}(?:[0-9]{3})?$ `),
		"mastercard":       regexp.MustCompile(`^(?:5[1-5][0-9]{2}|222[1-9]|22[3-9][0-9]|2[3-6][0-9]{2}|27[01][0-9]|2720)[0-9]{12}$`),
		"american express": regexp.MustCompile(`^3[47][0-9]{13}$`),
		"diners club":      regexp.MustCompile(`^3(?:0[0-5]|[68][0-9])[0-9]{11}$`),
		"discover":         regexp.MustCompile(`^6(?:011|5[0-9]{2})[0-9]{12}$`),
		"jcb":              regexp.MustCompile(`^(?:2131|1800|35\d{3})\d{11}$`),
	}
}

func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	// Build map of fields keyed by effective name.
	fields := make(map[string]reflect.Value)
	validFields := make(map[string]string)

	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)

		if validName := tag.Get("valid"); validName != "" {
			validFields[name] = validName
		}

	}

	for name, values := range req.Form {
		vf := validFields[name]
		for _, value := range values {
			if vf == "creditcard" {
				if !creditCardValid(value) {
					return fmt.Errorf("Invalid %s:%s", vf, value)
				}
			}
		}
	}

	// Update struct field for each parameter in the request.
	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue // ignore unrecognized HTTP parameters
		}
		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

func creditCardValid(v string) bool {
	regex := Regexp["creditCard"]
	if !regex.MatchString(v) {
		return false
	}

	switch {
	case Regexp["visa"].MatchString(v):
		return true
	case Regexp["mastercard"].MatchString(v):
		return true
	case Regexp["american express"].MatchString(v):
		return true
	case Regexp["diners club"].MatchString(v):
		return true
	case Regexp["discover"].MatchString(v):
		return true
	case Regexp["jcb"].MatchString(v):
		return true
	}

	return false

}

//!+populate
func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}

//!-populate

func Pack(i interface{}) string {
	vals := url.Values{}
	v := reflect.ValueOf(i).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		f := v.Field(i)
		switch f.Kind() {
		case reflect.Int:
			vals.Add(name, strconv.FormatInt(f.Int(), 10))
		case reflect.String:
			vals.Add(name, f.String())
		case reflect.Bool:
			vals.Add(name, strconv.FormatBool(f.Bool()))
		case reflect.Array, reflect.Slice:
			for j := 0; j < f.Len(); j++ {
				vals.Add(name, f.Index(j).String())
			}
		}
	}

	return vals.Encode()
}

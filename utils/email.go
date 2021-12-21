package utils

import (
	"bytes"
	"errors"
	"fmt"
	"html"
	"html/template"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// ToISO2022JP -> Convert UTF-8 to ISO2022JP
func ToISO2022JP(str string) ([]byte, error) {
	reader := strings.NewReader(str)
	transformer := japanese.ISO2022JP.NewEncoder()

	return ioutil.ReadAll(transform.NewReader(reader, transformer))
}

//ParseTemplate -> to parse the template with given data
func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	templatePath, err := filepath.Abs(fmt.Sprintf("utils/email_templates/%s", templateFileName))
	if err != nil {
		return "", errors.New("invalid template name")
	}
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	body := buf.String()
	unescapedString := html.UnescapeString(body)
	return unescapedString, nil
}

//IsInterfaceEmpty -> to check if the interface is empty
func IsInterfaceEmpty(x interface{}) bool {
	return x == reflect.Zero(reflect.TypeOf(x)).Interface()
}

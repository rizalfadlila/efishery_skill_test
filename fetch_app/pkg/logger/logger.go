package logger

import "github.com/sirupsen/logrus"

// Field :nodoc:
type Field struct {
	Key   string
	Value interface{}
}

func extract(args ...Field) map[string]interface{} {
	data := map[string]interface{}{}

	if len(args) == 0 {
		return data
	}

	for _, fl := range args {
		data[fl.Key] = fl.Value
	}

	return data
}

// Error :nodoc:
func Error(arg interface{}, fl ...Field) {
	logrus.WithFields(map[string]interface{}{
		"event": extract(fl...),
	}).Error(arg)
}

// Panic :nodoc:
func Panic(arg interface{}, fl ...Field) {
	logrus.WithFields(map[string]interface{}{
		"event": extract(fl...),
	}).Panic(arg)
}

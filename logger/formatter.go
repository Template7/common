package logger

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	// Default log format will output [INFO]: 2006-01-02T15:04:05Z07:00 - Log message
	// defaultLogFormat       = "[$level$]: [$time$] - $msg$\n"
	defaultTimestampFormat = time.RFC3339
	FieldKeyMsg            = "$msg$"
	FieldKeyLevel          = "$level$"
	FieldKeyTime           = "$time$"
	FieldKeyReqID          = "$req_id$"
	FieldKeyFunc           = "$func$"
	FieldKeyFile           = "$file$"
	FieldKeyOthers         = "$others$"
)

// Formatter implements logrus.Formatter interface.
type Formatter struct {
	TimestampFormat  string
	LogFormat        string
	OptionFields     []string
	CallerPrettyfier func(f *runtime.Frame) (string, string)
}

func defaultOptionFields() []string {
	return []string{
		fmt.Sprintf("[%s]", FieldKeyReqID),
		fmt.Sprintf("[%s]", FieldKeyFile),
		fmt.Sprintf("[%s]", FieldKeyFunc),
		fmt.Sprintf("[%s]", FieldKeyOthers),
	}
}

func defaultLogFormat() string {
	return fmt.Sprintf("[%s][%s][%s][%s][%s] - %s [%s]\n",
		FieldKeyLevel,
		FieldKeyTime,
		FieldKeyReqID,
		FieldKeyFile,
		FieldKeyFunc,
		FieldKeyMsg,
		FieldKeyOthers)
}

// Format building log message.
func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	output := f.LogFormat
	if output == "" {
		output = defaultLogFormat()
	}

	optionFields := f.OptionFields
	if len(optionFields) == 0 {
		optionFields = defaultOptionFields()
	}

	output = strings.Replace(output, FieldKeyMsg, entry.Message, 1)
	level := strings.ToUpper(entry.Level.String())
	output = strings.Replace(output, FieldKeyLevel, level, 1)

	output = f.replaceTime(entry, output)
	output = f.replaceFileAndFunc(entry, output)
	output = f.replaceReqID(entry, output)
	output = f.replaceOthers(entry, output)
	for _, k := range optionFields {
		output = strings.Replace(output, k, "", 1)
	}

	return []byte(output), nil
}

func (f *Formatter) replaceTime(entry *logrus.Entry, output string) string {

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	output = strings.Replace(output, FieldKeyTime, entry.Time.Format(timestampFormat), 1)
	return output
}

func (f *Formatter) replaceFileAndFunc(entry *logrus.Entry, output string) string {

	if entry.HasCaller() {
		//var funcVal, fileVal string
		var fileVal string
		if f.CallerPrettyfier != nil {
			_, fileVal = f.CallerPrettyfier(entry.Caller)
		} else {
			//funcVal = entry.Caller.Function
			fileVal = fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)
		}

		output = strings.Replace(output, FieldKeyFile, fileVal, 1)
		//output = strings.Replace(output, FieldKeyFunc, funcVal, 1)
	}

	return output
}

func (f *Formatter) replaceReqID(entry *logrus.Entry, output string) string {
	reqIDKey := strings.ReplaceAll(FieldKeyReqID, "$", "")
	if entry.Data[reqIDKey] != nil {
		output = strings.Replace(output, FieldKeyReqID, entry.Data[reqIDKey].(string), 1)
	}
	delete(entry.Data, reqIDKey)
	return output
}

func (f *Formatter) replaceOthers(entry *logrus.Entry, output string) string {
	others := ""
	if len(entry.Data) > 0 {
		bs, err := json.Marshal(entry.Data)
		if err == nil {
			others = string(bs)
		}
	}

	if len(others) > 0 {
		output = strings.Replace(output, FieldKeyOthers, others, 1)
	}
	return output
}

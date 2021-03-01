package my_errors

const (
	ErrorsConfigYamlNotExists string = "config.yml not exist"
	ErrorsConfigGormNotExists string = "gorm_v2.yml not exist"
	ErrorsConfigInitFail      string = "config initial failed"
	ErrorsGormInitFail        string = "Gorm initial failed"
	ErrorsAuthHeaderFail      string = "Authentication fail for http headers"
)

type MyError struct {
	ErrorString string
}

func (e *MyError) Error() string {
	return e.ErrorString
}

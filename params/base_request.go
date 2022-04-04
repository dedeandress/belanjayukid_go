package params

type RequestParams interface {
	Validate() error
}

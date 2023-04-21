package options

func (o *ApiServerOptions) Validate() []error {
	var errs []error
	errs = append(errs, o.MySQLOptions.Validate()...)

	return errs
}

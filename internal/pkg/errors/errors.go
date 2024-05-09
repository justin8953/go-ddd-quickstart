package errors

type GeneralError string

func (e GeneralError) Name() string {
	return "event.general.error"
}

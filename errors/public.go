package errors 


//Public embrulha o erro original com um novo erro
//que possui o método `Public() string` que vai retornar
//uma mensagem que é aceitável mostrar o erro ao público.
//Esse erro também pode ser desembrulhado usando o
//método tradicional com o package `errors`.
func Public(err error, msg string) error {
	return erroPublico{err,msg} 
}

type erroPublico struct {
		err error
		msg string
}

func(ep erroPublico) Error() string {
		return ep.err.Error()
}

func (ep erroPublico) Public() string {
		return ep.msg
}

func (ep erroPublico) Unwrap() error {
			return ep.err
}

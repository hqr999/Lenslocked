package errors

import "errors"


//Essas variáveis são usadas para nos dar acesso a funções
//existentes na std library errors package. Nós também 
//podemos embrulha-los em métodos com funcionalidades
//customizáveis se quiséssemos, ou imitá-los usando testes
var (
	As = errors.As
	Is = errors.Is
)


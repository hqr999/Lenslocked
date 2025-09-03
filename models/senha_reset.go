package models

import (
	"database/sql"
	"fmt"
	"time"
)

const(
		DuracaoDefault = 1 * time.Hour
)

type SenhaReset struct {
	ID     int
	UserID int

	//Token só é definido quando o SenhaReset é criado
	Token     string
	TokenHash string
	ExpiraEm  time.Time
}

type SenhaResetServico struct {
	BD *sql.DB
	//Bytes por token é usado para determinar quantos bytes
	//devem ser usados quando se fizer a geração do token de sessão
	//Se esse valor não for definido ou for menor que a
	//constante MinBytesPorToken, ele será ignorado e usaremos MinBytesPorToken.
	BytesPorToken int
	// Duracao é a quantidade de tempo em que SenhaReset é válido
	//Padrão é DuracaoDefault que é uma constante
	Duracao time.Duration
}


func(servico *SenhaResetServico) Cria(email string) (*SenhaReset,error){
	return nil,fmt.Errorf("A FAZER:Implementar SenhaResetServico.Cria()")
		 
}

func (servico *SenhaResetServico) Consome(token string) (*User,error) {
		return nil,fmt.Errorf("A FAZER:Implementar SenhaResetServico.Consome()")

}

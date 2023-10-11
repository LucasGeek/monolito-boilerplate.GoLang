package events

// UserCreated representa o evento de um usuário sendo criado.
type UserCreated struct {
	CPF       string // CPF do usuário.
	FirstName string // Nome.
	LastName  string // Sobrenome.
}

// NewUserCreated cria e retorna uma nova instância de UserCreated.
func NewUserCreated(cpf, firstName, lastName string) *UserCreated {
	return &UserCreated{
		CPF:       cpf,
		FirstName: firstName,
		LastName:  lastName,
	}
}

// GetName retorna o nome completo do usuário.
func (e *UserCreated) GetName() string {
	return e.FirstName + " " + e.LastName
}

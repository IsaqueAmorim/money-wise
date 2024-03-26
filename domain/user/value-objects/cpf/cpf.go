package cpf

type CPF struct {
	Number string
	Digit  string
}

func NewCPF(cpf string) *CPF {
	err := validate(cpf)

	if err != nil {

	}
	return &CPF{
		Number: cpf[00:10],
		Digit:  cpf[10:11],
	}
}

func validate(cpf string) error {
	return nil
}

func GetCpfFormatted() {

}

func (cpf CPF) GetCpf() string {
	return cpf.Number + cpf.Digit
}

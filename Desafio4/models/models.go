package models

type Pessoas struct {
	ID        string     `json:"id,omitempty"`
	Nome      string     `json:"nome,omitempty"`
	Sobrenome string     `json:"sobrenome,omitempty"`
	Endereco  *Enderecos `json:"endereco,omitempty"`
}
type Enderecos struct {
	Cidade string `json:"cidade,omitempty"`
	Estado string `json:"estado,omitempty"`
}

var Pessoa []Pessoas

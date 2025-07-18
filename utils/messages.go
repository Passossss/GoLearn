package utils

var Messages = struct {
	PostCreated    string
	PostUpdated    string
	PostDeleted    string
	PostNotFound   string
	InvalidPayload string
	ServerError    string
}{
	PostCreated:    "Post criado com sucesso.",
	PostUpdated:    "Post atualizado com sucesso.",
	PostDeleted:    "Post deletado com sucesso.",
	PostNotFound:   "Post não encontrado.",
	InvalidPayload: "Dados inválidos enviados.",
	ServerError:    "Erro interno no servidor.",
}

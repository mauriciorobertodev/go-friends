package controllers

import "net/http"

func ListUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando usuários..."))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuário..."))
}

func StoreUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário..."))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário..."))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário..."))
}

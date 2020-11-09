package controllers

func GetUsers(w http.responseWriter, r *http.Request) {
	w.Write([]byte("List users"))
}

func CreateUsers(w http.responseWriter, r *http.Request) {
	w.Write([]byte("Create user"))
}

func GetUser(w http.responseWriter, r *http.Request) {
	w.Write([]byte("An user"))
}

func UpdateUser(w http.responseWriter, r *http.Request) {
	w.Write([]byte("Update user"))
}

func DeleteUser(w http.responseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}

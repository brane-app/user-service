package main

type CreateUserBody struct {
	Nick     string `json:"nick"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
}

func (_ CreateUserBody) Types() (values map[string]string) {
	values = map[string]string{
		"nick":     "string",
		"password": "string",
		"email":    "string",
		"bio":      "string",
	}

	return
}

func (_ CreateUserBody) Defaults() (values map[string]interface{}) {
	values = map[string]interface{}{
		"bio": "",
	}

	return
}

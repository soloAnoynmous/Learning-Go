package main

type HTTPRequest struct {
	Method string
}

func main() {
	r:= HTTPRequest{Method:"GET"}

	switch r.Method {
	case "GET": println("GET")
	case "PUT": println("PUT")
	case "POST": println("POST")
	case "DELETE": println("DELETE")
	default:
		println("Invalid Request Type....")

	}
}

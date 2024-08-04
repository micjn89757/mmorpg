package main 

func main() {
	auth := NewAuth("127.0.0.1:3000")
	auth.Run()
}
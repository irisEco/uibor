package main
func main(){
	server := NewServer("172.0.0.0.1",8080)
	server.start()
}
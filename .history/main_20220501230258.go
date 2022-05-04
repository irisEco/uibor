package main
func main(){
	server := NewServer("172.0.0.1",3000)
	server.Start()
}
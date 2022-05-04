package main
func main(){
	server := NewServer("12.0.0.1",3000)
	server.Start()
}
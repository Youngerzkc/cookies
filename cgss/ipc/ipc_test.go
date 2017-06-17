package ipc

import (
	"testing"
)
type EchoServer struct{

}
func (server *EchoServer)Handle(method,params string) *Response{
	return &Response{"OK","ECHO: "+method + "~" +params}
}
func (server *EchoServer)Name() string  {
	return "Echoserver"
}
func TestIpc(t *testing.T){
	server:=NewIpcServer(&EchoServer{})
	client1:=NewIpcClient(server)
	client2:=NewIpcClient(server)
	resp1,_:=client1.Call("foo", "From Client1")
	resp2,_:=client2.Call("foo", "From Client2")
	if resp1.Body !="ECHO:foo ~From Client1" ||
	resp2.Body !="ECHO:foo ~From Client2"{
		t.Error("IpcClient.Call failed.")
	}
	client1.Close()
	client2.Close()
}
package server

type HelloService struct {}

func (s *HelloService) Hello(request string,replay *string) error  {
	*replay = "hello:"+request
	return nil
}




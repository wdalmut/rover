package proxy

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"sync"

	"github.com/gorilla/mux"
	zmq "github.com/pebbe/zmq4"
)

const (
	ClientChannel = "CLIENT"
	InProc        = "inproc://httpclient"
	startFrame    = 2048
)

type ChunkWrapper interface {
	Wrap([]byte) []byte
}

type Server struct {
	Router     *mux.Router
	HttpServer *http.Server
	Wrapper    ChunkWrapper
}

func (s *Server) HomeHandler(writer http.ResponseWriter, req *http.Request) {
	home, _ := ioutil.ReadFile("index.html")
	writer.Write(home)
}

func (s *Server) ReadHandler(writer http.ResponseWriter, req *http.Request) {
	log.Printf("New HTTP client connected")

	receiver, _ := zmq.NewSocket(zmq.SUB)
	defer receiver.Close()

	receiver.Connect(InProc)
	receiver.SetSubscribe(ClientChannel)

	writer.Header().Set("Content-Type", "text/html")

	closer := httputil.NewChunkedWriter(writer)
	defer closer.Close()

	//Start frame for browsers
	buf := make([]byte, startFrame)
	closer.Write(buf)
	if f, ok := writer.(http.Flusher); ok {
		f.Flush()
	}

	for {
		//  Read envelope with address
		receiver.RecvBytes(0)
		//  Read message contents
		data, _ := receiver.RecvBytes(0)

		_, e := closer.Write(s.Wrapper.Wrap(data))
		if f, ok := writer.(http.Flusher); ok {
			f.Flush()
		}

		if e != nil {
			break
		}
	}

	log.Println("HTTP client disconnected")
}

func (s *Server) ListenAndServe() {

	servers := new(sync.WaitGroup)

	s.Router.HandleFunc("/", s.HomeHandler).Methods("GET")
	s.Router.HandleFunc("/board/read", s.ReadHandler).Methods("GET")

	servers.Add(1)
	go s.HttpServer.ListenAndServe()

	servers.Wait()
}

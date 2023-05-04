package main

import (
	"flag"
	"fmt"
	"image"
	"log"
	"net"
	"os"

	engine "github.com/mapepema/yolobeer"
	toml "github.com/pelletier/go-toml"
	"google.golang.org/grpc"
)

var (
	confFile = flag.String("cfg", "conf.toml", "Path to the TOML configuration file")
)

func main() {

	cfgBytes, err := os.ReadFile(*confFile)
	if err != nil {
		fmt.Println(err)
		flag.Usage()
		return
	}

	var conf engine.Configuration
	err = toml.Unmarshal(cfgBytes, &conf)
	if err != nil {
		log.Fatalln(err)
	}

	if conf.YOLOBeers.Threshold <= 0.0 {
		conf.YOLOBeers.Threshold = 0.2
	}

	netw, err := engine.NewYOLONetwork(conf.YOLOBeers.Cfg, conf.YOLOBeers.Weights, conf.YOLOBeers.Threshold)
	if err != nil {
		log.Fatalln(err)
	}

	stdListener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", conf.ServerConf.Host, conf.ServerConf.Port))
	if err != nil {
		log.Fatal(err)
		return
	}

	if conf.ServerConf.QueueLimit < 1 {
		conf.ServerConf.QueueLimit = 1
	}

	// Init servers
	rs := &RecognitionServer{
		netW:        netw,
		framesQueue: make(chan *beerInfo, conf.ServerConf.QueueLimit),
		queueLimit:  conf.ServerConf.QueueLimit,
		resp:        make(chan *ServerResponse, conf.ServerConf.QueueLimit),
	}
	// Init neural network's queue
	rs.WaitFrames()

	// Init gRPC server
	grpcInstance := grpc.NewServer()

	// Register servers
	engine.RegisterServiceYOLOServer(
		grpcInstance,
		rs,
	)

	// Start
	if err := grpcInstance.Serve(stdListener); err != nil {
		log.Fatal(err)
		return
	}
}

// RecognitionServer Wrapper around engine.ServiceYOLOServer
type RecognitionServer struct {
	engine.ServiceYOLOServer
	netW        *engine.YOLONetwork
	framesQueue chan *beerInfo
	queueLimit  int
	resp        chan *ServerResponse
}

// ServerResponse Response from server
type ServerResponse struct {
	Resp  *engine.YOLOResponse
	Error error
}

type beerInfo struct {
	imageInfo *engine.ObjectInformation
	img       *image.NRGBA
}

// WaitFrames Endless loop for waiting frames
func (rs *RecognitionServer) WaitFrames() {
	fmt.Println("YOLO networks waiting for frames now")
	go func() {
		for {
			select {
			case n := <-rs.framesQueue:
				// fmt.Println("img of size", n.Bounds().Dx(), n.Bounds().Dy())
				resp, err := rs.netW.ReadBeers(n.img, true)

				rs.resp <- &ServerResponse{resp, err}
				continue
			}
		}
	}()
}

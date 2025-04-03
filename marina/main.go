package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"time"

	pb "marina/proto/grpc-server/proto" // Importa el paquete generado por protoc

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedComunicacionServer
}

// venderPirataMarina implementa la lógica para vender piratas a la Marina
func (s *server) VenderPirataMarina(ctx context.Context, req *pb.VentaRequest) (*pb.VentaResponse, error) {
	// Genera un número aleatorio para determinar si hay una redada
	rand.Seed(time.Now().UnixNano())
	redada := rand.Intn(100) < 25 // 25% de probabilidad de redada

	if redada {
		return &pb.VentaResponse{
			Exito:   false,
			Pago:    0,
			Mensaje: "Redada de la Marina: el pirata fue confiscado antes de la venta.",
		}, nil
	}

	// Calcula el pago (100% de la recompensa oficial)
	pago := 1000000 // Ejemplo de pago fijo

	return &pb.VentaResponse{
		Exito:   true,
		Pago:    int64(pago), // Convertimos `pago` a `int64`
		Mensaje: "Venta exitosa a la Marina.",
	}, nil
}

func main() {
	// Inicia el servidor gRPC
	listener, err := net.Listen("tcp", ":50053") // Puerto para la Marina
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterComunicacionServer(grpcServer, &server{})

	log.Println("Servidor de la Marina escuchando en el puerto 50053")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}

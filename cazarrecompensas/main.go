package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	pb "cazarrecompensas/proto/grpc-server/proto" // Importa el paquete generado por protoc

	"google.golang.org/grpc"
)

// Cazarecompensa representa un cazarrecompensas en el sistema.
type Cazarecompensa struct {
	ID         int64
	Nombre     string
	Reputacion int64
}

// Lista de piratas capturados
var piratasCapturados = []*pb.Pirata{}

func main() {
	// Solicita el nombre del cazarrecompensas.
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese su nombre de cazarrecompensas: ")
	nombre, _ := reader.ReadString('\n')
	nombre = strings.TrimSpace(nombre)

	// Crea un nuevo cazarrecompensas con los datos ingresados.
	cazador := &Cazarecompensa{
		ID:         1,      // Asigna un ID fijo (puedes cambiarlo si necesitas lógica más compleja).
		Nombre:     nombre, // Usa el nombre ingresado por el usuario.
		Reputacion: 100,    // Reputación inicial.
	}

	log.Printf("Cazarrecompensas creado: %s (ID: %d, Reputación: %d)", cazador.Nombre, cazador.ID, cazador.Reputacion)

	// Conéctate al servidor gRPC del gobierno.
	connGobierno, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) // Cambia el puerto al del gobierno
	if err != nil {
		log.Fatalf("No se pudo conectar al servidor del gobierno: %v", err)
	}
	defer connGobierno.Close()
	clientGobierno := pb.NewComunicacionClient(connGobierno)

	// Conéctate al servidor gRPC del Submundo.
	connSubmundo, err := grpc.Dial("localhost:50052", grpc.WithInsecure()) // Cambia el puerto al del Submundo
	if err != nil {
		log.Fatalf("No se pudo conectar al servidor del Submundo: %v", err)
	}
	defer connSubmundo.Close()
	clientSubmundo := pb.NewComunicacionClient(connSubmundo)

	// Conéctate al servidor gRPC de la Marina.
	connMarina, err := grpc.Dial("localhost:50053", grpc.WithInsecure()) // Cambia el puerto al de la Marina
	if err != nil {
		log.Fatalf("No se pudo conectar al servidor de la Marina: %v", err)
	}
	defer connMarina.Close()
	clientMarina := pb.NewComunicacionClient(connMarina)

	// Menú interactivo
	for {
		fmt.Println("\n--- Menú de opciones ---")
		fmt.Println("1. Solicitar lista de piratas al gobierno")
		fmt.Println("2. Capturar un pirata")
		fmt.Println("3. Vender un pirata al Submundo")
		fmt.Println("4. Vender un pirata a la marina")
		fmt.Println("5. Salir")
		fmt.Print("Seleccione una opción: ")

		// Leer la opción seleccionada
		optionStr, _ := reader.ReadString('\n')
		optionStr = strings.TrimSpace(optionStr)
		option, err := strconv.Atoi(optionStr)
		if err != nil {
			fmt.Println("Por favor, ingrese un número válido.")
			continue
		}

		switch option {
		case 1:
			// Solicitar lista de piratas al gobierno
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			response, err := clientGobierno.ConsultarListaPiratas(ctx, &pb.Empty{})
			if err != nil {
				log.Printf("Error al obtener la lista de piratas: %v", err)
				continue
			}

			// Imprime la lista de piratas recibida
			log.Println("Lista de piratas buscados:")
			for _, pirata := range response.Piratas {
				log.Printf("- ID: %d, Nombre: %s, Recompensa: %d, Peligrosidad: %s, Estado: %s",
					pirata.Id, pirata.Nombre, pirata.Recompensa, pirata.NivelDePeligrosidad, pirata.Estado)
			}

		case 2:
			// Capturar un pirata
			fmt.Print("Ingrese el ID del pirata que desea capturar: ")
			pirataIDStr, _ := reader.ReadString('\n')
			pirataIDStr = strings.TrimSpace(pirataIDStr)
			pirataID, err := strconv.ParseInt(pirataIDStr, 10, 64)
			if err != nil {
				fmt.Println("Por favor, ingrese un ID válido.")
				continue
			}

			// Simula la captura del pirata
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			response, err := clientGobierno.ConsultarListaPiratas(ctx, &pb.Empty{})
			if err != nil {
				log.Printf("Error al obtener la lista de piratas: %v", err)
				continue
			}

			var pirataCapturado *pb.Pirata
			for _, pirata := range response.Piratas {
				if pirata.Id == pirataID && pirata.Estado == "Buscado" {
					pirataCapturado = pirata
					pirata.Estado = "Capturado"
					piratasCapturados = append(piratasCapturados, pirata)
					break
				}
			}

			if pirataCapturado != nil {
				log.Printf("¡Pirata capturado! Nombre: %s, Recompensa: %d", pirataCapturado.Nombre, pirataCapturado.Recompensa)
				cazador.Reputacion += 50 // Incrementa la reputación por capturar un pirata
			} else {
				fmt.Println("No se encontró un pirata con ese ID o ya fue capturado.")
			}

		case 3:
			// Vender un pirata al Submundo
			// Vender un pirata al Submundo
			fmt.Print("Ingrese el ID del pirata que desea vender: ")
			pirataIDStr, _ := reader.ReadString('\n')
			pirataIDStr = strings.TrimSpace(pirataIDStr)
			pirataID, err := strconv.ParseInt(pirataIDStr, 10, 64)
			if err != nil {
				fmt.Println("Por favor, ingrese un ID válido.")
				continue
			}

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			venderReq := &pb.VentaRequest{
				PirataId:  pirataID,
				CazadorId: cazador.ID,
			}

			venderResp, err := clientSubmundo.VenderPirataSubmundo(ctx, venderReq)
			if err != nil {
				log.Printf("Error al intentar vender el pirata: %v", err)
				continue
			}

			fmt.Println(venderResp.Mensaje)
			if venderResp.Exito {
				fmt.Printf("¡Venta exitosa! Has recibido %d berries.\n", venderResp.Pago)
				cazador.Reputacion += 50 // Incrementa la reputación por la venta exitosa
			} else {
				fmt.Println("La venta no fue exitosa.")
			}

		case 4:
			// Vender un pirata a la Marina
			fmt.Print("Ingrese el ID del pirata que desea vender: ")
			pirataIDStr, _ := reader.ReadString('\n')
			pirataIDStr = strings.TrimSpace(pirataIDStr)
			pirataID, err := strconv.ParseInt(pirataIDStr, 10, 64)
			if err != nil {
				fmt.Println("Por favor, ingrese un ID válido.")
				continue
			}

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			venderReq := &pb.VentaRequest{
				PirataId:  pirataID,
				CazadorId: cazador.ID,
			}

			venderResp, err := clientMarina.VenderPirataMarina(ctx, venderReq)
			if err != nil {
				log.Printf("Error al intentar vender el pirata: %v", err)
				continue
			}

			fmt.Println(venderResp.Mensaje)
			if venderResp.Exito {
				fmt.Printf("¡Venta exitosa! Has recibido %d berries.\n", venderResp.Pago)
				cazador.Reputacion += 30 // Incrementa la reputación por la venta exitosa
			} else {
				fmt.Println("La venta no fue exitosa.")
			}

		case 5:
			// Salir del programa
			fmt.Println("Saliendo del programa. ¡Hasta luego!")
			return

		default:
			fmt.Println("Opción no válida. Por favor, intente de nuevo.")
		}
	}
}

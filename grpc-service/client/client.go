package main

import (
	"context"
	"log"
	"time"

	pb "grpc-sensor-service/pb/sensor"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func printSensor(client pb.SensorServiceClient, req *pb.GetSensorRequest) {
	log.Printf("Recherche du capteur avec l'ID : %d", req.Id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := client.GetSensor(ctx, req)
	if err != nil {
		log.Fatalf("Erreur du serveur (ou introuvable) : %v", err)
	}

	log.Printf("Succès ! Capteur trouvé : %v", res.GetSensor())
}

func printAllSensors(client pb.SensorServiceClient) {
	log.Printf("Recherche de tous les capteurs")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := client.ListSensors(ctx, &pb.ListSensorsRequest{})
	if err != nil {
		log.Fatalf("Erreur: %v", err)
	}
	log.Printf("Capteurs : %v", res.Sensors)
}

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient("localhost:50051", opts...)
	if err != nil {
		log.Fatalf("Impossible de se connecter au serveur : %v", err)
	}
	defer conn.Close()

	client := pb.NewSensorServiceClient(conn)

	printSensor(client, &pb.GetSensorRequest{
		Id: 1,
	})
	printAllSensors(client)
}

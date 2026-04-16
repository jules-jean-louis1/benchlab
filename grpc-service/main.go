package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"time"

	"grpc-sensor-service/db"
	pb "grpc-sensor-service/pb/sensor"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// 1. La "struct" est notre contrôleur. Elle réunit les composants dont le serveur a besoin.
type sensorServer struct {
	pb.UnimplementedSensorServiceServer         // (Obligatoire) Fournit une réponse "Non implémenté" par défaut
	database                            *sql.DB // C'est ici qu'on stocke notre connexion à la base de données
}

func (s *sensorServer) CreateSensor(ctx context.Context, req *pb.CreateSensorRequest) (*pb.CreateSensorResponse, error) {
	sensor := req.GetSensor()
	if sensor == nil {
		return nil, fmt.Errorf("Le capteur est requis")
	}
	var id int32
	sqlStatment := `INSERT INTO sensor (name, type, location, unit, status, last_value, last_reading_at) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id`

	// On gère les dates proprement (si le client n'envoie pas de date, on met NULL en base)
	var readingAt interface{}
	if sensor.LastReadingAt != nil {
		readingAt = sensor.LastReadingAt.AsTime()
	} else {
		readingAt = nil
	}

	// On utilise .String() sur Type et Status pour envoyer "TEMPERATURE" et non pas 0
	err := s.database.QueryRowContext(ctx, sqlStatment,
		sensor.Name,
		sensor.Type.String(),
		sensor.Location,
		sensor.Unit,
		sensor.Status.String(),
		sensor.LastValue,
		readingAt,
	).Scan(&id)

	if err != nil {
		return nil, fmt.Errorf("Erreur lors de l'insertion du capteur: %v", err)
	}
	sensor.Id = id
	return &pb.CreateSensorResponse{Sensor: sensor}, nil
}

func (s *sensorServer) GetSensor(ctx context.Context, req *pb.GetSensorRequest) (*pb.GetSensorResponse, error) {
	id := req.GetId()
	sensor := &pb.Sensor{}

	var tempType sql.NullString
	var tempStatus sql.NullString
	var tempLastR sql.NullTime
	var tempCr sql.NullTime
	sqlStatment := `SELECT id,name,type,location,unit,status,last_value,last_reading_at,created_at FROM sensor WHERE id=$1`
	err := s.database.QueryRowContext(ctx, sqlStatment, id).Scan(&sensor.Id, &sensor.Name, &tempType, &sensor.Location, &sensor.Unit, &tempStatus, &sensor.LastValue, &tempLastR, &tempCr)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("Capteur introuvable")
	} else if err != nil {
		return nil, fmt.Errorf("Erreur interne: %v", err)
	}
	if tempType.Valid {
		value := pb.SensorType_value[tempType.String]
		sensor.Type = pb.SensorType(value)
	}
	if tempStatus.Valid {
		value := pb.SensorStatus_value[tempStatus.String]
		sensor.Status = pb.SensorStatus(value)
	}
	if tempLastR.Valid {
		sensor.LastReadingAt = timestamppb.New(tempLastR.Time)
	}
	if tempCr.Valid {
		sensor.CreatedAt = timestamppb.New(tempCr.Time)
	}
	return &pb.GetSensorResponse{Sensor: sensor}, nil
}
func (s *sensorServer) ListSensors(ctx context.Context, req *pb.ListSensorsRequest) (*pb.ListSensorsResponse, error) {
	sqlStatment := `SELECT id, name, type, location, unit, status, last_value, last_reading_at, created_at FROM sensor`
	rows, err := s.database.QueryContext(ctx, sqlStatment)
	if err != nil {
		return nil, fmt.Errorf("Erreur lors de la récupération des capteurs: %v", err)
	}
	defer rows.Close() // Toujours fermer les lignes après lecture !

	// On prépare un tableau (slice) vide qui va contenir tous nos capteurs
	var sensors []*pb.Sensor

	for rows.Next() {
		// 1. On prépare un nouveau capteur à chaque tour de boucle
		sensor := &pb.Sensor{}

		var tempType, tempStatus sql.NullString
		var tempLastR, tempCr sql.NullTime

		// 2. On scanne la ligne actuelle
		err := rows.Scan(
			&sensor.Id, &sensor.Name, &tempType, &sensor.Location, &sensor.Unit,
			&tempStatus, &sensor.LastValue, &tempLastR, &tempCr,
		)
		if err != nil {
			return nil, fmt.Errorf("Erreur de format de ligne: %v", err)
		}

		// 3. On hydrate les champs spéciaux (exactement comme dans GetSensor)
		if tempType.Valid {
			sensor.Type = pb.SensorType(pb.SensorType_value[tempType.String])
		}
		if tempStatus.Valid {
			sensor.Status = pb.SensorStatus(pb.SensorStatus_value[tempStatus.String])
		}
		if tempLastR.Valid {
			sensor.LastReadingAt = timestamppb.New(tempLastR.Time)
		}
		if tempCr.Valid {
			sensor.CreatedAt = timestamppb.New(tempCr.Time)
		}

		// 4. On ajoute ce capteur dans notre tableau
		sensors = append(sensors, sensor)
	}

	return &pb.ListSensorsResponse{Sensors: sensors}, nil
}

func (s *sensorServer) UpdateSensor(ctx context.Context, req *pb.UpdateSensorRequest) (*pb.UpdateSensorResponse, error) {
	sensor := req.GetSensor()
	if sensor == nil {
		return nil, fmt.Errorf("Le capteur est requis")
	}

	// Pour un UPDATE, pas besoin de lire des résultats, on utilise ExecContext !
	sqlStatment := `
		UPDATE sensor 
		SET name=$1, type=$2, location=$3, unit=$4, status=$5, last_value=$6, last_reading_at=$7 
		WHERE id=$8
	`

	// Astuce: sensor.Type.String() renvoie "TEMPERATURE" au lieu de 0 pour la base de données
	res, err := s.database.ExecContext(ctx, sqlStatment,
		sensor.Name,
		sensor.Type.String(),
		sensor.Location,
		sensor.Unit,
		sensor.Status.String(),
		sensor.LastValue,
		sensor.LastReadingAt.AsTime(), // Conversion Inverse: Protobuf -> Go time.Time
		sensor.Id,
	)

	if err != nil {
		return nil, fmt.Errorf("Erreur lors de la mise à jour: %v", err)
	}

	// On vérifie si l'ID existait vraiment et a bien été modifié
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return nil, fmt.Errorf("Capteur avec l'ID %d introuvable", sensor.Id)
	}

	return &pb.UpdateSensorResponse{Sensor: sensor}, nil
}

func (s *sensorServer) DeleteSensor(ctx context.Context, req *pb.DeleteSensorRequest) (*pb.DeleteSensorResponse, error) {
	Id := req.GetId()
	sql := `DELETE FROM sensor WHERE id=$1`

	// Pareil ici, DELETE ne renvoie pas de réponse, on utilise ExecContext
	res, err := s.database.ExecContext(ctx, sql, Id)
	if err != nil {
		return nil, fmt.Errorf("Erreur lors de la suppresion: %v", err)
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return nil, fmt.Errorf("Aucun capteur avec l'ID %d", Id)
	}

	return &pb.DeleteSensorResponse{Success: true}, nil
}

func (s *sensorServer) StreamSensorReadings(req *pb.StreamSensorReadingsRequest, stream grpc.ServerStreamingServer[pb.StreamSensorReadingsResponse]) error {
	id := req.GetId()
	sensor := &pb.Sensor{}

	var tempType sql.NullString
	var tempStatus sql.NullString
	var tempLastR sql.NullTime
	var tempCr sql.NullTime
	sqlStatment := `SELECT id,name,type,location,unit,status,last_value,last_reading_at,created_at FROM sensor WHERE id=$1`
	err := s.database.QueryRowContext(stream.Context(), sqlStatment, id).Scan(&sensor.Id, &sensor.Name, &tempType, &sensor.Location, &sensor.Unit, &tempStatus, &sensor.LastValue, &tempLastR, &tempCr)
	if err == sql.ErrNoRows {
		return fmt.Errorf("Capteur introuvable")
	} else if err != nil {
		return fmt.Errorf("Erreur interne: %v", err)
	}
	if tempType.Valid {
		value := pb.SensorType_value[tempType.String]
		sensor.Type = pb.SensorType(value)
	}
	if tempStatus.Valid {
		value := pb.SensorStatus_value[tempStatus.String]
		sensor.Status = pb.SensorStatus(value)
	}
	if tempLastR.Valid {
		sensor.LastReadingAt = timestamppb.New(tempLastR.Time)
	}
	if tempCr.Valid {
		sensor.CreatedAt = timestamppb.New(tempCr.Time)
	}
	for i := 0; i < 1000; i++ {
		sensor.LastReadingAt = timestamppb.New(time.Now())
		sensor.LastValue = sensor.LastValue + float32(i)
		if err := stream.Send(&pb.StreamSensorReadingsResponse{Sensor: sensor}); err != nil {
			return fmt.Errorf("Erreur lors de l'envoi du capteur: %v", err)
		}
	}
	return nil
}

func main() {
	// 2. On récupère la connexion à la base de données
	dbConn := db.Connect()
	defer dbConn.Close() // Bonne pratique: fermer la DB quand le programme s'arrête

	// 3. On ouvre le port réseau 50051 (le port classique de gRPC)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Impossible d'écouter sur le port 50051: %v", err)
	}

	// 4. On crée le moteur du serveur gRPC
	s := grpc.NewServer()

	// 5. On "enregistre" notre code (la struct) auprès de gRPC
	myServer := &sensorServer{
		database: dbConn,
	}
	pb.RegisterSensorServiceServer(s, myServer)

	// 6. On lance le serveur !
	fmt.Println("Serveur gRPC démarré sur le port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Le serveur a crashé: %v", err)
	}
}

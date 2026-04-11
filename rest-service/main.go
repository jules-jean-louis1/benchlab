package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var DB *sql.DB

type Sensor struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Type          string  `json:"type"`
	Location      string  `json:"location"`
	Unit          string  `json:"unit"`
	Status        string  `json:"status"`
	LastValue     float64 `json:"last_value"`
	LastReadingAt string  `json:"last_reading_at"`
	CreatedAt     string  `json:"created_at"`
}

func main() {
	connString := "postgres://postgres:password@db:5432/benchlab?sslmode=disable"

	DB, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()

	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	// Define a simple GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/sensors", func(ctx *gin.Context) {
		var newSensor Sensor
		if err := ctx.ShouldBindJSON(&newSensor); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		sqlStatment := `INSERT INTO sensor (name, type, location, unit, status, last_value, last_reading_at) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id`
		err := DB.QueryRowContext(ctx, sqlStatment, newSensor.Name, newSensor.Type, newSensor.Location, newSensor.Unit, newSensor.Status, newSensor.LastValue, newSensor.LastReadingAt).Scan(&newSensor.ID)
		if err != nil {
			log.Fatal(err)
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Sensor created successfully"})
	})

	r.GET("/sensor/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		if len(id) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "id not provided"})
		}
		var sensor Sensor
		sqlStatment := `SELECT id, name, type, location, unit, status, last_value, last_reading_at, created_at FROM sensor WHERE id=$1`
		err := DB.QueryRowContext(ctx, sqlStatment, id).Scan(&sensor.ID, &sensor.Name, &sensor.Type, &sensor.Location, &sensor.Unit, &sensor.Status, &sensor.LastValue, &sensor.LastReadingAt, &sensor.CreatedAt)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Sensor not found"})
				return
			}
			log.Printf("erreur lors de la récupération du capteur: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du capteur"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"sensor": gin.H{"id": sensor.ID, "name": sensor.Name, "type": sensor.Type, "location": sensor.Location, "unit": sensor.Unit, "status": sensor.Status, "last_value": sensor.LastValue, "last_reading_at": sensor.LastReadingAt, "created_at": sensor.CreatedAt}})
	})

	r.GET("/sensors", func(ctx *gin.Context) {
		rows, err := DB.QueryContext(ctx, `SELECT id, name, type, location, unit, status, last_value, last_reading_at, created_at FROM sensor`)
		if err != nil {
			log.Printf("erreur lors de la récupération des capteurs: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des capteurs"})
			return
		}
		defer rows.Close()

		var sensors []Sensor
		for rows.Next() {
			var sensor Sensor
			err := rows.Scan(&sensor.ID, &sensor.Name, &sensor.Type, &sensor.Location, &sensor.Unit, &sensor.Status, &sensor.LastValue, &sensor.LastReadingAt, &sensor.CreatedAt)
			if err != nil {
				log.Printf("erreur lors du scan des capteurs: %v", err)
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors du scan des capteurs"})
				return
			}
			sensors = append(sensors, sensor)
		}
		ctx.JSON(http.StatusOK, gin.H{"sensors": sensors})
	})

	r.PUT("/sensor/:id", func(ctx *gin.Context) {
		var updatedSensor Sensor
		if err := ctx.ShouldBindJSON(&updatedSensor); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id := ctx.Param("id")
		if len(id) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "id not provided"})
			return
		}
		sqlStatment := `UPDATE sensor SET name=$1, type=$2, location=$3, unit=$4, status=$5, last_value=$6, last_reading_at=$7 WHERE id=$8`
		result, err := DB.ExecContext(ctx, sqlStatment, updatedSensor.Name, updatedSensor.Type, updatedSensor.Location, updatedSensor.Unit, updatedSensor.Status, updatedSensor.LastValue, updatedSensor.LastReadingAt, id)
		if err != nil {
			log.Printf("erreur lors de la mise à jour du capteur: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la mise à jour du capteur"})
			return
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Printf("erreur lors de la récupération du nombre de lignes affectées: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du nombre de lignes affectées"})
			return
		}
		if rowsAffected == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Sensor not found"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Sensor updated successfully"})
	})

	r.DELETE("/sensor/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		if len(id) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "id not provided"})
			return
		}
		sqlStatment := `DELETE FROM sensor WHERE id=$1`
		result, err := DB.ExecContext(ctx, sqlStatment, id)
		if err != nil {
			log.Printf("erreur lors de la suppression du capteur: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la suppression du capteur"})
			return
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Printf("erreur lors de la récupération du nombre de lignes affectées: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du nombre de lignes affectées"})
			return
		}
		if rowsAffected == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Sensor not found"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Sensor deleted successfully"})
	})
	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run()
}

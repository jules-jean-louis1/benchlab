Apres avoir lancé le serveur gRPC, vous pouvez utiliser `grpcurl` pour interagir avec les services définis dans vos fichiers `.proto`. Voici quelques exemples de commandes `grpcurl` pour tester les différentes méthodes de votre service de capteurs.

Installation de grpcurl (si vous ne l'avez pas déjà) :

```bash
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```

Si la commande `grpcurl` n'est pas reconnue, assurez-vous que le répertoire `$GOPATH/bin` est dans votre variable d'environnement `PATH`.

```bash
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
source ~/.bashrc
```

1. Créer un capteur (CreateSensor)

```bash
grpcurl -plaintext \
  -import-path ./proto \
  -proto proto/sensor/sensor.proto \
  -d '{
    "sensor": {
      "name": "Turbine-A3-Temp",
      "type": "TEMPERATURE",
      "location": "Bâtiment C - Salle 12",
      "unit": "°C",
      "status": "ACTIVE",
      "last_value": 25.5
    }
  }' \
  localhost:50051 sensor.SensorService/CreateSensor
```

2. Récupérer un capteur par ID (GetSensor)

```bash
grpcurl -plaintext \
  -import-path ./proto \
  -proto proto/sensor/sensor.proto \
  -d '{
    "id": "1"
  }' \
  localhost:50051 sensor.SensorService/GetSensor
```

3. Mettre à jour un capteur (UpdateSensor)

```bash
grpcurl -plaintext \
  -import-path ./proto \
  -proto proto/sensor/sensor.proto \
  -d '{
    "id": "1",
    "sensor": {
      "name": "Turbine-A3-Temp",
      "type": "TEMPERATURE",
      "location": "Bâtiment C - Salle 12",
      "unit": "°C",
      "status": "ACTIVE",
      "last_value": 26.0
    }
  }' \
  localhost:50051 sensor.SensorService/UpdateSensor
``` 

4. Supprimer un capteur (DeleteSensor)

```bash
grpcurl -plaintext \
  -import-path ./proto \
  -proto proto/sensor/sensor.proto \
  -d '{
    "id": "1"
  }' \
  localhost:50051 sensor.SensorService/DeleteSensor
``` 

5. Lister tous les capteurs (ListSensors)

```bash
grpcurl -plaintext \
  -import-path ./proto \
  -proto proto/sensor/sensor.proto \
  localhost:50051 sensor.SensorService/ListSensors
```

6. Utiliser le stream pour simuler un flux de données (StreamSensorReadings)

```bash
grpcurl -plaintext \
  -import-path ./proto \
  -proto proto/sensor/sensor.proto \
  -d '{
    "id": "4"
  }' \
  localhost:50051 sensor.SensorService/StreamSensorReadings
```

```bash
grpcurl -plaintext \
  -import-path ./proto \
  -proto proto/sensor/sensor.proto \
  -d '{
    "limit": 100
    "offset": 0
  }' \
  localhost:50051 sensor.SensorService/PageSensors
```

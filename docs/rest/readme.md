Après avoir lancé le serveur REST, vous pouvez utiliser `curl` (ou un outil comme Postman / Insomnia) pour interagir avec les endpoints de votre API. Voici quelques exemples de commandes pour tester les différentes routes de votre service de capteurs (qui écoute sur le port 8080 par défaut).

1. Vérifier si l'API est en ligne (Ping)

```bash
curl http://localhost:8080/ping
```

2. Créer un capteur (POST /sensors)

```bash
curl -X POST http://localhost:8080/sensors \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Turbine-A3-Temp",
    "type": "TEMPERATURE",
    "location": "Bâtiment C - Salle 12",
    "unit": "°C",
    "status": "ACTIVE",
    "last_value": 25.5,
    "last_reading_at": "2026-04-11T10:00:00Z"
  }'
```

3. Récupérer un capteur par ID (GET /sensor/:id)

```bash
curl http://localhost:8080/sensor/1
```

4. Mettre à jour un capteur (PUT /sensor/:id)

```bash
curl -X PUT http://localhost:8080/sensor/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Turbine-A3-Temp",
    "type": "TEMPERATURE",
    "location": "Bâtiment C - Salle 12",
    "unit": "°C",
    "status": "ACTIVE",
    "last_value": 26.0,
    "last_reading_at": "2026-04-11T10:05:00Z"
  }'
```

5. Supprimer un capteur (DELETE /sensor/:id)

```bash
curl -X DELETE http://localhost:8080/sensor/1
```

6. Lister tous les capteurs (GET /sensors)

```bash
curl http://localhost:8080/sensors
```

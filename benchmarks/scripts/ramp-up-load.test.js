import http from 'k6/http';
import grpc from 'k6/net/grpc';
import { check } from 'k6';

// 1. On charge le fichier .proto pour que k6 comprenne les méthodes gRPC
const client = new grpc.Client();
client.load(['../../proto'], 'sensor/sensor.proto');

export const options = {
  scenarios: {
    // ---- Scénario REST ----
    rest_scenario: {
      executor: 'ramping-vus',
      startVUs: 10,
      stages: [
        { duration: '30s', target: 100 }, // Ramp-up de 10 à 100 VUs
        { duration: '1m', target: 100 },  // Maintien de 100 VUs
        { duration: '30s', target: 0 },   // Ramp-down de 100 à 0 VUs
      ],
      exec: 'test_rest',
    },
    // ---- Scénario gRPC ----
    grpc_scenario: {
      executor: 'ramping-vus',
      startVUs: 10,
      stages: [
        { duration: '30s', target: 100 }, // Ramp-up de 10 à 100 VUs
        { duration: '1m', target: 100 },  // Maintien de 100 VUs
        { duration: '30s', target: 0 },   // Ramp-down de 100 à 0 VUs
      ],
      exec: 'test_grpc',
    },
  },
};

// Fonction exécutée pour le REST
export function test_rest() {
  const res = http.get('http://localhost:8080/sensor/1');
  check(res, { 'REST status is 200': (r) => r.status === 200 });
}

// Fonction exécutée pour gRPC (Pas besoin de CLI, k6 gère gRPC nativement)
export function test_grpc() {
  client.connect('localhost:50051', {
    plaintext: true, // équivalent de -plaintext sur grpcurl
  });

  const res = client.invoke('sensor.SensorService/GetSensor', { id: "1" });
  check(res, { 'gRPC status is OK': (r) => r.status === grpc.StatusOK });

  client.close();
}
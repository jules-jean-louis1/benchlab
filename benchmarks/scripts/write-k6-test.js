import http from "k6/http";
import grpc from "k6/net/grpc";
import { check } from "k6";

const client = new grpc.Client();
client.load(["../../proto"], "sensor/sensor.proto");

export const options = {
  scenarios: {
    rest_scenario: {
      executor: "shared-iterations",
      vus: 5,
      iterations: 500,
      exec: "test_rest",
    },
    grpc_scenario: {
      executor: "shared-iterations",
      vus: 5,
      iterations: 500,
      exec: "test_grpc",
    },
  },
};

// Le payload est identique pour les deux protocoles
const payloadContent = {
  name: "Turbine-A3-Temp",
  type: "TEMPERATURE",
  location: "Bâtiment C - Salle 12",
  unit: "°C",
  status: "ACTIVE",
  last_value: 25.5,
  last_reading_at: "2026-04-11T10:00:00Z"
};

export function test_rest() {
  const res = http.post(
    "http://localhost:8080/sensors", 
    JSON.stringify(payloadContent), 
    { headers: { "Content-Type": "application/json" } }
  );
  check(res, { "REST status is 200": (r) => r.status === 200 || r.status === 201 });
}

export function test_grpc() {
  client.connect("localhost:50051", { plaintext: true });
  
  // En gRPC, la structure attendue possède la propriété "sensor"
  const res = client.invoke("sensor.SensorService/CreateSensor", { sensor: payloadContent });
  check(res, { "gRPC status is OK": (r) => r.status === grpc.StatusOK });

  client.close();
}

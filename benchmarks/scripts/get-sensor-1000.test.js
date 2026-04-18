import http from "k6/http";
import grpc from "k6/net/grpc";
import { check } from "k6";

// 1. On charge le fichier .proto pour que k6 comprenne les méthodes gRPC
const client = new grpc.Client();
client.load(["../../proto"], "sensor/sensor.proto");

export const options = {
  scenarios: {
    rest_scenario: {
      executor: "req",
      vus: 10,
      iterations: 1000,
      exec: "test_rest",
    },
    grpc_scenario: {
      executor: "req",
      vus: 10,
      iterations: 1,
      exec: "test_grpc",
    },
  },
};

export function test_rest() {
  const res = http.get("http://localhost:8080/sensor/1");
  check(res, { "REST status is 200": (r) => r.status === 200 });
}

export function test_grpc() {
  client.connect("localhost:50051", {
    plaintext: true,
  });

  const res = client.invoke("sensor.SensorService/StreamSensorReadings", { id: "4" });
  check(res, { "gRPC status is OK": (r) => r.status === grpc.StatusOK });

  client.close();
}

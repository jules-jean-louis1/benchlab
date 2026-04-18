import http from "k6/http";
import grpc from "k6/net/grpc";
import { check } from "k6";

const client = new grpc.Client();
client.load(["../../proto"], "sensor/sensor.proto");

export const options = {
  scenarios: {
    rest: {
      executor: "shared-iterations",
      vus: 1,
      iterations: 1,
      exec: "test_rest",
    },
    grpc: {
      executor: "shared-iterations",
      vus: 1,
      iterations: 1,
      exec: "test_grpc",
    },
  },
};

export function test_rest() {
  for (let step = 0; step <= 10; step++) {
    const res = http.get(
      `http://localhost:8080/page-sensors?limit=100&offset=${step * 100}`,
    );
    check(res, { "REST status is 200": (r) => r.status === 200 });
  }
}

export function test_grpc() {
  client.connect("localhost:50051", { plaintext: true });
  for (let step = 0; step <= 10; step++) {
    const res = client.invoke("sensor.SensorService/PageSensors", {
      limit: "100",
      offset: `${step * 100}`,
    });
    check(res, { "gRPC status is OK": (r) => r && r.status === grpc.StatusOK });
  }
  client.close();
}

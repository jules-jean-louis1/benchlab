DROP TABLE IF EXISTS sensor;

CREATE TYPE sensor_type AS ENUM ('TEMPERATURE', 'PRESSURE', 'VIBRATION');
CREATE TYPE sensor_status AS ENUM ('ACTIVE','INACTIVE','MAINTENANCE');

CREATE TABLE sensor (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NULL,
    type sensor_type NULL,
    location VARCHAR(255) NULL,
    unit VARCHAR(50) NULL,
    status sensor_status NULL,
    last_value DOUBLE PRECISION NULL,
    last_reading_at TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO sensor (name, type, location, unit, status, last_value, last_reading_at) VALUES
('Temperature Sensor 1', 'TEMPERATURE', 'Factory Floor', 'Celsius', 'ACTIVE', 25.0, NOW()),
('Pressure Sensor 1', 'PRESSURE', 'Pipeline', 'Bar', 'ACTIVE', 10.0, NOW()),
('Vibration Sensor 1', 'VIBRATION', 'Machine A', 'mm/s', 'MAINTENANCE', 0.5, NOW());
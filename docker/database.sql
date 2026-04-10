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

INSERT INTO sensor (name, type, location, unit, status) VALUES
('Temperature Sensor 1', 'TEMPERATURE', 'Factory Floor', 'Celsius', 'ACTIVE'),
('Pressure Sensor 1', 'PRESSURE', 'Pipeline', 'Bar', 'ACTIVE'),
('Vibration Sensor 1', 'VIBRATION', 'Machine A', 'mm/s', 'MAINTENANCE');
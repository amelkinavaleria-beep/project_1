CREATE TABLE IF NOT EXISTS measurements (
    id BIGSERIAL PRIMARY KEY,
    device_id BIGINT NOT NULL,
    room_id BIGINT NOT NULL,
    value BIGINT NOT NULL,
    created_date TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_date TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_date TIMESTAMP WITH TIME ZONE,
    
    CONSTRAINT fk_measurement_device 
        FOREIGN KEY(device_id) 
        REFERENCES devices(id) 
        ON DELETE CASCADE,
    CONSTRAINT fk_measurement_room 
        FOREIGN KEY(room_id) 
        REFERENCES rooms(id) 
        ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_measurements_device_id ON measurements(device_id);
CREATE INDEX IF NOT EXISTS idx_measurements_room_id ON measurements(room_id);
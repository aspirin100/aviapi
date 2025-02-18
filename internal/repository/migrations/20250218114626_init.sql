-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tickets(
    order_id UUID PRIMARY KEY,
    from_country VARCHAR NOT NULL,
    to_country VARCHAR NOT NULL,
    carrier VARCHAR NOT NULL,
    departure_date TIMESTAMPTZ NOT NULL,
    arrival_date TIMESTAMPTZ NOT NULL,
    registration_date TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS documents(
    id UUID PRIMARY KEY,
    passenger_id UUID NOT NULL,
    document_type VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS passengers(
    id UUID PRIMARY KEY,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    patronymic VARCHAR
);

ALTER TABLE documents
    ADD CONSTRAINT fk_passenger_id
    FOREIGN KEY (passenger_id) REFERENCES passengers(id)
    ON DELETE CASCADE;

CREATE TABLE IF NOT EXISTS ticket_passengers(
    order_id UUID NOT NULL,
    passenger_id UUID NOT NULL,
    PRIMARY KEY(order_id, passenger_id)
);


ALTER TABLE ticket_passengers
    ADD CONSTRAINT fk_passenger_id
    FOREIGN KEY (passenger_id) REFERENCES passengers(id)
    ON DELETE CASCADE;

ALTER TABLE ticket_passengers
    ADD CONSTRAINT fk_order_id
    FOREIGN KEY (order_id) REFERENCES tickets(order_id)
    ON DELETE CASCADE;
-- +goose StatementEnd



-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS tickets;
DROP TABLE IF EXISTS passengers;
DROP TABLE IF EXISTS documents;

DROP TABLE IF EXISTS ticket_passengers;

-- +goose StatementEnd

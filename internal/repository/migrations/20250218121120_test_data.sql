-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

INSERT INTO tickets (order_id, from_country, to_country, carrier, departure_date, arrival_date, registration_date)
VALUES
    (uuid_generate_v4(), 'USA', 'Germany', 'Lufthansa', '2023-11-01 10:00:00+00', '2023-11-01 18:00:00+00', '2023-10-25 12:00:00+00'),
    (uuid_generate_v4(), 'France', 'Japan', 'Air France', '2023-11-05 08:30:00+00', '2023-11-06 14:00:00+00', '2023-10-28 15:00:00+00'),
    (uuid_generate_v4(), 'Russia', 'China', 'Aeroflot', '2023-11-10 12:00:00+00', '2023-11-10 19:30:00+00', '2023-11-01 09:00:00+00');

INSERT INTO passengers (id, first_name, last_name, patronymic)
VALUES
    (uuid_generate_v4(), 'John', 'Doe', 'Michael'),
    (uuid_generate_v4(), 'Jane', 'Smith', NULL),
    (uuid_generate_v4(), 'Alice', 'Johnson', 'Robert');

INSERT INTO documents (id, passenger_id, document_type)
VALUES
    (uuid_generate_v4(), (SELECT id FROM passengers WHERE first_name = 'John' AND last_name = 'Doe'), 'Passport'),
    (uuid_generate_v4(), (SELECT id FROM passengers WHERE first_name = 'Jane' AND last_name = 'Smith'), 'ID Card'),
    (uuid_generate_v4(), (SELECT id FROM passengers WHERE first_name = 'Alice' AND last_name = 'Johnson'), 'Passport');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM tickets;
DELETE FROM documents;
DELETE FROM passengers;
DELETE FROM ticket_passengers;
-- +goose StatementEnd

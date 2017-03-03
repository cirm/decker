CREATE TABLE decker.rooms (
  room_id SERIAL PRIMARY KEY
);

CREATE TABLE decker.room_details (
  id        SERIAL PRIMARY KEY,
  room_id   INTEGER NOT NULL,
  player_id INTEGER NOT NULL,
  FOREIGN KEY (room_id) REFERENCES decker.rooms (room_id),
  FOREIGN KEY (player_id) REFERENCES decker.players (id),
  UNIQUE (room_id, player_id)
);

CREATE TABLE decker.messages (
  id               SERIAL PRIMARY KEY,
  message_datetime TIMESTAMP NOT NULL DEFAULT NOW(),
  message_text     VARCHAR(500),
  chat_id          INTEGER   NOT NULL,
  player_id        INTEGER   NOT NULL,
  FOREIGN KEY (player_id) REFERENCES decker.players (id),
  FOREIGN KEY (chat_id) REFERENCES decker.rooms (room_id)
);


CREATE TABLE IF NOT EXISTS keyboard(
  id   INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT    NOT NULL
);

CREATE TABLE IF NOT EXISTS keyboard_key(
  keyboard_id INTEGER                                NOT NULL,
  type        TEXT CHECK(type IN ('HOLD', 'TOGGLE')) NOT NULL,
  key_code    INTEGER                                NOT NULL,
  min_speed   REAL                                   NOT NULL,
  max_speed   REAL,
  FOREIGN KEY (keyboard_id) REFERENCES keyboard(id)
);
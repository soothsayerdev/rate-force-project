CREATE TABLE players (
    player_id SERIAL PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
):

CREATE TABLE decks (
    deck_id SERIAL PRIMARY KEY,
    deck_name VARCHAR(150) UNIQUE NOT NULL,
    description TEXT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE matches (
    match_id SERIAL PRIMARY KEY, 
    player_id INT NOT NULL REFERENCES players(player_id) ON DELETE CASCADE,
    player_deck_id INT NOT NULL REFERENCES decks(deck_id),
    opponent_deck_id INT NOT NULL REFERENCES decks(deck_id),
    result VARCHAR(10) NOT NULL CHECK (result IN ('win', 'loss', 'draw')),
    match_date DATE NOT NULL DEFAULT CURRENT_DATE,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- index
CREATE INDEX idx_matches_player_id ON matches(player_id);
CREATE INDEX idx_matches_match_date ON matches(match_date);
CREATE INDEX idx_matches_player_deck_id ON matches(player_deck_id);
CREATE INDEX idx_matches_opponent_deck_id ON matches(opponent_deck_id);
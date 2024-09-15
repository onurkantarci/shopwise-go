CREATE TABLE IF NOT EXISTS stores (
    id SERIAL PRIMARY KEY NOT NULL,
    store_name VARCHAR(45) NOT NULL,
    is_active BOOLEAN NOT NULL,
    banner_image VARCHAR(255),
    logo_image VARCHAR(255),
    icon_image VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS deals (
    id SERIAL PRIMARY KEY NOT NULL,
    store_id INTEGER REFERENCES stores(id),
    game_id INTEGER NOT NULL,
    steam_app_id INTEGER,
    internal_name VARCHAR(100) NOT NULL,
    title VARCHAR(255) NOT NULL,
    metacritic_link VARCHAR(255),
    sales_price DECIMAL(10, 2) NOT NULL,
    normal_price DECIMAL(10, 2) NOT NULL,
    is_on_sale BOOLEAN NOT NULL,
    savings DECIMAL(5, 2) NOT NULL,
    metacritic_score INTEGER NOT NULL,
    steam_rating_text VARCHAR(50),
    steam_rating_percent INTEGER NOT NULL,
    steam_rating_count INTEGER NOT NULL,
    release_date INTEGER NOT NULL,
    last_change INTEGER NOT NULL,
    deal_rating DECIMAL(3, 2),
    thumb VARCHAR(255)
)
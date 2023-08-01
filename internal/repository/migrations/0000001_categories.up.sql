CREATE TABLE categories (
    id UUID NOT NULL PRIMARY KEY ,
    child_of UUID REFERENCES categories(id),

    category_name TEXT NOT NULL DEFAULT '',
    image_url TEXT NOT NULL DEFAULT '',
    is_active BOOL NOT NULL DEFAULT true,
    sort_index BIGINT NOT NULL DEFAULT 100
);

CREATE INDEX categories__child_of__idx on categories (child_of);
CREATE INDEX categories__sort_index__idx on categories (sort_index);
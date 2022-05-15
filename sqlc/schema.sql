CREATE TABLE public.users (
    id SERIAL NOT NULL,
    name VARCHAR(30) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE public.article(
    id SERIAL NOT NULL,
    name VARCHAR(30) NOT NULL,
    user_id INTEGER NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES public.users (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);
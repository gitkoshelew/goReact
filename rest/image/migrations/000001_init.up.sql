CREATE TABLE IF NOT EXISTS IMAGES 
(   image_id              serial PRIMARY key,
    type            CHARACTER VARYING(30) NOT NULL ,
    ownerId         INTEGER 
);

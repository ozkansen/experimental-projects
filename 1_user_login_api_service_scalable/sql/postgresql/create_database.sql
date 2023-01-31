CREATE SCHEMA IF NOT EXISTS "public";

CREATE  TABLE IF NOT EXISTS "public".users (
    id          serial  NOT NULL  ,
    username    varchar(50)  NOT NULL  ,
    passwd      varchar(100) NOT NULL  ,

    CONSTRAINT  pk_users PRIMARY KEY ( id )
);

CREATE UNIQUE INDEX unq_users ON "public".users ( username );

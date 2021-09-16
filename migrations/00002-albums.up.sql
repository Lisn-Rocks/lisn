create table if not exists albums (
    albumid serial primary key,
    album text not null,
    artistid int not null constraint artist references artists on delete cascade,
    coverext text default '.jpg' not null
);
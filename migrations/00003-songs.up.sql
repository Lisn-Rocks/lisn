create table if not exists songs (
    songid serial primary key,
    song text not null,
    albumid int not null constraint album references albums on delete cascade,
    audioext text default '.mp3' not null
);
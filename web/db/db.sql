CREATE DATABASE cc_fileshare;

create table vfs (
userid text,
fileid uuid,
filename text,
creation_date date,
attributes JSONB
);
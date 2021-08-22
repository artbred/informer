create table if not exists chats (
    chat_token varchar unique primary key not null,
    chat_id bigint unique not null
);
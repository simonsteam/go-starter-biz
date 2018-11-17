CREATE TABLE IF NOT EXISTS "user" (
    "id" bigserial, 
    "branch_id" bigint, 
    "description" text, 
    "no" text, 
    "password" text, 
    "real_name" text, 
    "state" text, 
    "phone" text, 
    "username" text, 
    "add_time" timestamptz not null, 
    "upd_time" timestamptz, 
    "del_time" timestamptz, 
    PRIMARY KEY ("id")
) 


--##

CREATE TABLE IF NOT EXISTS "group" (
    "id" text, "name" text, 
    "permissions" text[], 
    "desc" text, 
    PRIMARY KEY ("id")
) 

--##


CREATE TABLE IF NOT EXISTS "user_group" (
    "user_id" bigint, 
    "group_id" text, 
    FOREIGN KEY ("user_id") REFERENCES "user" ("id"), 
    FOREIGN KEY ("group_id") REFERENCES "group" ("id")
)  


--##

CREATE TYPE config_type AS ENUM ('bool', 'checkbox_string', 'float', 'image_url', 'int', 'int_arr', 'radio_string', 'string', 'string_arr', 'time' );

--##

CREATE TABLE public.config
(
    id bigserial NOT NULL,
    add_time timestamptz not null,
    upd_time timestamptz,
    del_time timestamptz,
    key text not null unique, -- unique business key
    name text not null,
    module text not null, -- used for group datas
    description text not null, -- description
    options text[], -- avaliable ≈options for radio_string/checkbox_string
    readonly bool not null,
    type config_type not null,
    bool_val bool,
    float_val double precision,
    int_val int,
    str_val text,
    time_val timestamptz,
    int_array_val int[],
    str_array_val text[],
    upd_permissions text[],
    validate jsonb, -- validate rules for val
    visible bool not null,
    PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
);


--##
create table public.branch 
(
    id bigserial NOT NULL,
    add_time timestamptz not null,
    upd_time timestamptz, 
    del_time timestamptz,
    no text, --编号
    name text not null, --名字
    mgr_user_id bigint, --管理员id
    address text not null, -- 地址
    tel text not null, -- 电话
    admin_desc text not null, -- 备注
    introduction text not null, -- 简介
    state text not null, -- 状态
    lat double precision, --经纬度
    lng double precision,
    PRIMARY KEY (id),
    FOREIGN KEY ("mgr_user_id") REFERENCES "user" ("id")
)
WITH (
    OIDS = FALSE
);

--##

ALTER TABLE public.user ADD CONSTRAINT user_branch_fkey
   FOREIGN KEY (branch_id) REFERENCES branch(id);
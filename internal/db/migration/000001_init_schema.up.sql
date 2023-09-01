DROP TYPE IF EXISTS circle_type;
CREATE TYPE circle_type AS ENUM ('ROOM', 'HALL');

CREATE TABLE "circles"
(
    "id"               uuid PRIMARY KEY,
    "owner_id"         uuid           NOT NULL,
    "title"            varchar UNIQUE NOT NULL,
    "avatar"           text,
    "description"      text,
    "circle_type"      circle_type    NOT NULL,
    "is_private"       boolean        NOT NULL DEFAULT true,
    "is_featured"      boolean        NOT NULL DEFAULT false,
    "display_duration" int            NOT NULL DEFAULT 15,
    "created_at"       timestamptz    NOT NULL DEFAULT (now()),
    "deleted_at"       timestamptz             DEFAULT NULL
);

CREATE TABLE "tags"
(
    "id"   bigserial PRIMARY KEY,
    "name" varchar NOT NULL UNIQUE
);
CREATE INDEX "tags_name_index" ON "tags" ("name");

CREATE TABLE "circle_tag"
(
    "id"        bigserial PRIMARY KEY,
    "circle_id" uuid NOT NULL,
    "tag_id"    uuid NOT NULL,
    UNIQUE ("circle_id", "tag_id")
);
CREATE INDEX "circle_tag_circle_id_index" ON "circle_tag" ("circle_id");
CREATE INDEX "circle_tag_tag_id_index" ON "circle_tag" ("tag_id");

DROP TYPE IF EXISTS membership_type;
CREATE TYPE membership_type AS ENUM ('OWNER', 'ADMIN', 'POSTER', 'VIEWER');

DROP TYPE IF EXISTS acceptance_type;
CREATE TYPE acceptance_type AS ENUM ('ACCEPT_ALL', 'ASK_FIRST', 'REJECT_ALL');

CREATE TABLE "circle_members"
(
    "id"              bigserial PRIMARY KEY,
    "circle_id"       uuid            NOT NULL,
    "member_id"       uuid            NOT NULL,
    "membership_type" membership_type NOT NULL DEFAULT 'VIEWER',
    "acceptance_type" acceptance_type NOT NULL DEFAULT 'ASK_FIRST',
    "created_at"      timestamptz     NOT NULL DEFAULT (now()),
    UNIQUE ("circle_id", "member_id")
);

ALTER TABLE "circle_members"
    ADD FOREIGN KEY ("circle_id") REFERENCES "circles" ("id")
        ON UPDATE CASCADE
        ON DELETE CASCADE;

CREATE TABLE "moods"
(
    "id"             uuid PRIMARY KEY,
    "poster_id"      uuid        NOT NULL,
    "circle_id"      uuid        NOT NULL,
    "description"    text,
    "image"          text        NOT NULL,
    "set_background" boolean     NOT NULL default false,
    "created_at"     timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "moods"
    ADD FOREIGN KEY ("circle_id") REFERENCES "circles" ("id")
        ON UPDATE CASCADE
        ON DELETE CASCADE;

CREATE TABLE "circle_join_requests"
(
    "id"        bigserial PRIMARY KEY,
    "circle_id" uuid NOT NULL,
    "user_id"   uuid NOT NULL,
    UNIQUE ("circle_id", "user_id")
);
CREATE INDEX "circle_join_requests_circle_id_user_id_index" ON "circle_join_requests" ("circle_id", "user_id");
CREATE INDEX "circle_join_requests_user_id_index" ON "circle_join_requests" ("user_id");

CREATE TABLE "circle_invitations"
(
    "id"        bigserial PRIMARY KEY,
    "circle_id" uuid NOT NULL,
    "user_id"   uuid NOT NULL,
    "message"   text,
    UNIQUE ("circle_id", "user_id")
);
ALTER TABLE "circle_invitations"
    ADD FOREIGN KEY ("circle_id") REFERENCES "circles" ("id")
        ON UPDATE CASCADE
        ON DELETE CASCADE;
CREATE INDEX "circle_invitations_circle_id_user_id_index" ON "circle_invitations" ("circle_id", "user_id");
CREATE INDEX "circle_invitations_user_id_index" ON "circle_invitations" ("user_id");

CREATE TABLE "reactions"
(
    "id"                  bigserial PRIMARY KEY,
    "name"                varchar NOT NULL UNIQUE,
    "emoji"               varchar,
    "text_representation" varchar
);
CREATE INDEX "reactions_name_index" ON "reactions" ("name");
CREATE INDEX "reactions_text_representation_index" ON "reactions" ("text_representation");

INSERT INTO "reactions" (name, emoji, text_representation)
VALUES ('Thumbs Up', '👍', ':thumbsup:'),
       ('Thumbs Down', '👎', ':thumbdown:');

CREATE TABLE "mood_reactions"
(
    "id"          bigserial PRIMARY KEY,
    "mood_id"     uuid        NOT NULL,
    "user_id"     uuid        NOT NULL,
    "reaction_id" bigint      NOT NULL,
    "created_at"  timestamptz NOT NULL DEFAULT (now()),
    UNIQUE ("mood_id", "user_id")
);

ALTER TABLE "mood_reactions"
    ADD FOREIGN KEY ("reaction_id") REFERENCES "reactions" ("id")
        ON UPDATE CASCADE
        ON DELETE CASCADE;

ALTER TABLE "mood_reactions"
    ADD FOREIGN KEY ("mood_id") REFERENCES "moods" ("id")
        ON UPDATE CASCADE
        ON DELETE CASCADE;

CREATE INDEX "mood_reactions_mood_id_user_id_index" ON "mood_reactions" ("mood_id", "user_id");
CREATE INDEX "mood_reactions_mood_id_index" ON "mood_reactions" ("mood_id");
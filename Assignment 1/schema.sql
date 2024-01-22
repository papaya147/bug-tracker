CREATE TABLE IF NOT EXISTS profile(
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    verified BOOLEAN NOT NULL DEFAULT false,
    createdAt TIMESTAMPTZ NOT NULL,
    updatedAt TIMESTAMPTZ NOT NULL
);
CREATE TABLE IF NOT EXISTS organisation(
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    owner UUID NOT NULL,
    createdAt TIMESTAMPTZ NOT NULL,
    updatedAt TIMESTAMPTZ NOT NULL
);
ALTER TABLE organisation
ADD FOREIGN KEY (owner) REFERENCES profile(id);
CREATE TABLE IF NOT EXISTS team(
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    organisation UUID NOT NULL,
    createdAt TIMESTAMPTZ NOT NULL,
    updatedAt TIMESTAMPTZ NOT NULL
);
ALTER TABLE team
ADD FOREIGN KEY (organisation) REFERENCES organisation(id);
CREATE TABLE IF NOT EXISTS teamMember(
    team UUID NOT NULL,
    profile UUID NOT NULL,
    createdAt TIMESTAMPTZ NOT NULL,
    updatedAt TIMESTAMPTZ NOT NULL,
    PRIMARY KEY (team, profile)
);
ALTER TABLE teamMember
ADD FOREIGN KEY (team) REFERENCES team(id);
ALTER TABLE teamMember
ADD FOREIGN KEY (profile) REFERENCES profile(id);
CREATE TABLE IF NOT EXISTS bug(
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    status TEXT NOT NULL,
    priority TEXT NOT NULL,
    assignedTo UUID NOT NULL,
    assignedBy UUID NOT NULL,
    closedBy UUID,
    createdAt TIMESTAMP NOT NULL,
    updatedAt TIMESTAMP NOT NULL,
    closedAt TIMESTAMP
);
ALTER TABLE bug
ADD FOREIGN KEY (assignedTo) REFERENCES team(id);
ALTER TABLE bug
ADD FOREIGN KEY (assignedBy) REFERENCES profile(id);
ALTER TABLE bug
ADD FOREIGN KEY (closedBy) REFERENCES profile(id);
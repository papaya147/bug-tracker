CREATE TABLE IF NOT EXISTS profile(
    id UUID PRIMARY KEY,
    tokenId UUID NOT NULL,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    verified BOOLEAN NOT NULL DEFAULT false,
    createdAt TIMESTAMPTZ NOT NULL DEFAULT now(),
    updatedAt TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE TABLE IF NOT EXISTS organisation(
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    owner UUID NOT NULL UNIQUE,
    createdAt TIMESTAMPTZ NOT NULL DEFAULT now(),
    updatedAt TIMESTAMPTZ NOT NULL DEFAULT now()
);
ALTER TABLE organisation
ADD FOREIGN KEY (owner) REFERENCES profile(id);
CREATE TABLE IF NOT EXISTS organisationTransfer(
    id UUID NOT NULL,
    organisation UUID NOT NULL,
    fromProfile UUID NOT NULL,
    toProfile UUID NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT false,
    createdAt TIMESTAMPTZ NOT NULL DEFAULT now()
);
ALTER TABLE organisationTransfer
ADD FOREIGN KEY (organisation) REFERENCES organisation(id);
ALTER TABLE organisationTransfer
ADD FOREIGN KEY (fromProfile) REFERENCES profile(id);
ALTER TABLE organisationTransfer
ADD FOREIGN KEY (toProfile) REFERENCES profile(id);
CREATE TABLE IF NOT EXISTS team(
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    organisation UUID NOT NULL,
    createdAt TIMESTAMPTZ NOT NULL DEFAULT now(),
    updatedAt TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE(name, organisation)
);
ALTER TABLE team
ADD FOREIGN KEY (organisation) REFERENCES organisation(id);
CREATE TABLE IF NOT EXISTS teamMember(
    team UUID NOT NULL,
    profile UUID NOT NULL,
    admin BOOLEAN NOT NULL,
    createdAt TIMESTAMPTZ NOT NULL DEFAULT now(),
    updatedAt TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (team, profile)
);
ALTER TABLE teamMember
ADD FOREIGN KEY (team) REFERENCES team(id);
ALTER TABLE teamMember
ADD FOREIGN KEY (profile) REFERENCES profile(id);
CREATE TYPE bugStatus AS ENUM('PENDING', 'PROCESSING');
CREATE TYPE bugPriority AS ENUM('URGENT', 'HIGH', 'LOW');
CREATE TABLE IF NOT EXISTS bug(
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    status bugStatus NOT NULL,
    priority bugPriority NOT NULL,
    assignedTo UUID NOT NULL,
    assignedBy UUID NOT NULL,
    completed BOOLEAN DEFAULT FALSE,
    createdAt TIMESTAMPTZ NOT NULL DEFAULT now(),
    updatedAt TIMESTAMPTZ NOT NULL DEFAULT now(),
    closedBy UUID,
    remarks TEXT,
    closedAt TIMESTAMPTZ
);
ALTER TABLE bug
ADD FOREIGN KEY (assignedTo) REFERENCES team(id);
ALTER TABLE bug
ADD FOREIGN KEY (assignedBy) REFERENCES profile(id);
ALTER TABLE bug
ADD FOREIGN KEY (closedBy) REFERENCES profile(id);
ALTER TABLE votes
    DROP CONSTRAINT fk_session;
ALTER TABLE flows
    DROP CONSTRAINT fk_session;
    ALTER TABLE voters
    DROP CONSTRAINT fk_session;
ALTER TABLE candidates
    DROP CONSTRAINT fk_session;
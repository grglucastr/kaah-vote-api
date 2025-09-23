ALTER TABLE candidates
    ADD CONSTRAINT fk_session
    FOREIGN KEY(session_id)
    REFERENCES sessions(id)
    ON DELETE CASCADE;

ALTER TABLE voters
    ADD CONSTRAINT fk_session
    FOREIGN KEY(session_id)
    REFERENCES sessions(id)
    ON DELETE CASCADE;

ALTER TABLE votes
    ADD CONSTRAINT fk_session
    FOREIGN KEY(session_id)
    REFERENCES sessions(id)
    ON DELETE CASCADE;

ALTER TABLE flows
    ADD CONSTRAINT fk_session
    FOREIGN KEY(session_id)
    REFERENCES sessions(id)
    ON DELETE CASCADE;
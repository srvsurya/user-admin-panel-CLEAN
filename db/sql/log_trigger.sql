CREATE OR REPLACE FUNCTION log_user_creation()
RETURNS TRIGGER
LANGUAGE plpgsql
AS $$
BEGIN
INSERT INTO user_creation_log(user_id, created_at) VALUES (NEW.user_id, NOW());
RETURN NEW;
END;
$$;
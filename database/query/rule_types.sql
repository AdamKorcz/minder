-- name: CreateRuleType :one
INSERT INTO rule_type (
    name,
    project_id,
    description,
    guidance,
    definition,
    severity_value,
    subscription_id,
    display_name
) VALUES (
    $1,
    $2,
    $3,
    $4,
    sqlc.arg(definition)::jsonb,
    sqlc.arg(severity_value),
    sqlc.narg(subscription_id),
    sqlc.arg(display_name)
) RETURNING *;

-- name: ListRuleTypesByProject :many
SELECT * FROM rule_type WHERE project_id = $1;

-- name: GetRuleTypeByID :one
SELECT * FROM rule_type WHERE id = $1;

-- name: GetRuleTypeByName :one
SELECT * FROM rule_type WHERE  project_id = sqlc.arg(project_id) AND lower(name) = lower(sqlc.arg(name));

-- name: DeleteRuleType :exec
DELETE FROM rule_type WHERE id = $1;

-- name: UpdateRuleType :one
UPDATE rule_type
    SET description = $2, definition = sqlc.arg(definition)::jsonb, severity_value = sqlc.arg(severity_value), display_name = sqlc.arg(display_name)
    WHERE id = $1
    RETURNING *;

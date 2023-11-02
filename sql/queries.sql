-- name: CreateAccount :execresult
INSERT INTO Accounts (
    name,
    email,
    username,
    password
) VALUES (
    ?,
    ?,
    ?,
    ?
);

-- name: GetAccount :one
SELECT * FROM Accounts WHERE username = ? AND password = ? LIMIT 1
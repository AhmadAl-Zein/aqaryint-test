-- name: CreateUser :one
INSERT INTO
    users(name, phone_number)
VALUES ($1, $2) RETURNING id,
    name,
    phone_number,
    otp,
    otp_expiration_time;

-- name: GetUserByPhoneNumber :one
SELECT
    id,
    name,
    phone_number,
    otp,
    otp_expiration_time
FROM users
WHERE phone_number = $1;

-- name: UpdateOTP :exec
UPDATE users
SET
    otp = $1,
    otp_expiration_time = $2
WHERE phone_number = $3;

-- name: VerifyOTP :one
SELECT id
FROM users
WHERE
    phone_number = $1
    AND otp = $2
    AND otp_expiration_time > NOW();
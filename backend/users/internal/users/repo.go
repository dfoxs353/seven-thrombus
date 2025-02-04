package users

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"gitlab.com/volgaIt/packages/postgres"
)

type Repo struct {
	db postgres.PostgresQuerier
}

func NewRepo(db postgres.PostgresQuerier) *Repo {
	return &Repo{db}
}

func (r *Repo) GetUserByUsername(ctx context.Context, username string) (User, error) {
	sql := `
	SELECT
		id,
		username,
		"password",
		roles,
		first_name,
		last_name
	FROM
		"users"	
	WHERE not is_deleted AND username = $1`

	rows, err := r.db.Query(ctx, sql, username)
	if err != nil {
		return User{}, err
	}
	defer rows.Close()

	return pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[User])
}

func (r *Repo) InsertUser(
	ctx context.Context,
	username string,
	password string,
	roles []Role,
	firstname string,
	lastname string,
) (int, error) {
	sql := `
	INSERT INTO users(
		username,
		PASSWORD,
		roles,
		first_name,
		last_name)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id
	`

	rolesData, _ := json.Marshal(roles)

	var id int
	err := r.db.QueryRow(
		ctx,
		sql,
		username,
		password,
		string(rolesData),
		firstname,
		lastname,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repo) InsertUserRefreshToken(ctx context.Context, uid int, token string, expIn int) error {
	sql := `
	INSERT INTO tokens(user_id, refresh_token, expires_in)
	VALUES($1, $2, $3)
	`
	_, err := r.db.Exec(ctx, sql, uid, token, time.Unix(int64(expIn), 0))
	return err
}

type Token struct {
	UserId       int       `db:"user_id"`
	RefreshToken string    `db:"refresh_token"`
	ExpiresIn    time.Time `db:"expires_in"`
}

func (r *Repo) GetUserToken(ctx context.Context, token string) (Token, error) {
	sql := `
	SELECT user_id, refresh_token, expires_in
	FROM tokens
	WHERE refresh_token = $1
	`

	rows, err := r.db.Query(ctx, sql, token)
	if err != nil {
		return Token{}, err
	}
	defer rows.Close()

	return pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Token])
}

func (r *Repo) DeleteUserToken(ctx context.Context, token string) error {
	sql := `
	DELETE FROM tokens WHERE refresh_token = $1
	`

	_, err := r.db.Exec(ctx, sql, token)

	return err
}

func (r *Repo) GetUserById(ctx context.Context, id int) (User, error) {
	sql := `
	SELECT
		id,
		username,
		"password",
		roles,
		first_name,
		last_name
	FROM
		"users"	
	WHERE NOT is_deleted AND id = $1`

	rows, err := r.db.Query(ctx, sql, id)
	if err != nil {
		return User{}, err
	}
	defer rows.Close()

	return pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[User])

}

func (r *Repo) UpdateUser(
	ctx context.Context,
	uid int,
	username string,
	password string,
	firstName string,
	lastName string,
	roles []Role,
) error {
	rolesData, _ := json.Marshal(roles)

	sql := `
	UPDATE users
	SET username = $1,
		password = $2,
		first_name = $3,
		last_name = $4,
		roles = $5
	WHERE id = $6
	`

	_, err := r.db.Exec(ctx, sql, username, password, firstName, lastName, string(rolesData), uid)
	return err
}

// nameFilter - фильтер по имени пользователя.
//
// roles - фильтер по ролям пользователя.
// В выборку попадут только те пользователи у которых есть
// все роли указанные в roles
func (r *Repo) GetUsers(
	ctx context.Context,
	from int,
	count int,
	nameFilter *string,
	roles []Role,
) ([]User, error) {
	sql := `
	SELECT
		id,
		username,
		"password",
		roles,
		first_name,
		last_name
	FROM
		"users"	
	WHERE  
		NOT is_deleted
		AND 
			roles @> $1
		AND 
			(concat(first_name, ' ', last_name) ILIKE $2
			OR
			concat(last_name, ' ', first_name) ILIKE $2)
	ORDER BY id
	LIMIT $3
	OFFSET $4
	`

	rolesData := []byte("[]")
	if len(roles) > 0 {
		rolesData, _ = json.Marshal(roles)
	}

	nameLike := `%%%%`
	if nameFilter != nil {
		nameLike = fmt.Sprintf("%%%s%%", *nameFilter)
	}

	rows, err := r.db.Query(ctx, sql, string(rolesData), nameLike, count, from)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByNameLax[User])
}

// SELECT *
// FROM users
// WHERE roles @> '[]' AND concat(first_name, ' ', last_name) ILIKE '%soldatov%';

func (r *Repo) DeleteUserSoft(ctx context.Context, uid int) error {
	sql := `
	UPDATE users
	SET is_deleted = true
	WHERE id = $1
	`

	_, err := r.db.Exec(ctx, sql, uid)
	return err
}

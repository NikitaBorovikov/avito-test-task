package postgres

type PostgresRepo struct {
	UserRepo        *UserRepo
	TeamRepo        *TeamRepo
	PullRequestRepo *PullRequestRepo
}

func NewPostgresRepo() *PostgresRepo {
	return &PostgresRepo{
		UserRepo:        NewUserRepo(),
		TeamRepo:        NewTeamRepo(),
		PullRequestRepo: NewPullRequestRepo(),
	}
}

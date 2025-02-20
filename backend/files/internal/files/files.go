package files

import (
	"context"
	"io"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioCfg struct {
	Endpoint string `mapstructure:"endpoint"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Bucket   string `mapstructure:"password"`
}

type Service struct {
	repo        *Repo
	minioClient *minio.Client
	db          *pgxpool.Pool
}

func NewService(
	cfg MinioCfg,
	db *pgxpool.Pool,
) *Service {
	client, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(cfg.Username, cfg.Password, ""),
	})
	if err != nil {
		panic(err)
	}

	return &Service{NewRepo(db), client, db}
}

type FileMeta struct {
	Id     int    `json:"id"`
	Target string `json:"-"`
	Size   int    `json:"size"`
	Url    string `json:"url"`
}

func UploadFile(
	ctx context.Context,
	filename string,
	size int,
	fs io.Reader,
) (FileMeta, error) {
	// тут зписываем информацию в базу + аплоадим файл в s3 хранилище
}

package redis

import (
    "context"
    "fmt"
    "time"

    "github.com/redis/go-redis/v9"
    "backend_course/branch_api_gateway/config"
)

// Интерфейс для хранения Redis
type IRedisStorage interface {
    SetX(ctx context.Context, key string, value interface{}, duration time.Duration) error
    Get(ctx context.Context, key string) (string, error)
    Del(ctx context.Context, key string) error
}

type Store struct {
    db *redis.Client
}

func New(cfg config.Config) IRedisStorage {
    fmt.Println("cfg.RedisHost:", cfg.RedisHost)
    fmt.Println("cfg.RedisPort:", cfg.RedisPort)
    fmt.Println("cfg.RedisPassword:", cfg.RedisPassword) // Optional: for debugging purposes

    redisClient := redis.NewClient(&redis.Options{
        Addr:     cfg.RedisHost + ":" + cfg.RedisPort,
        Password: cfg.RedisPassword, // Set the password
    })

    // Проверка соединения
    _, err := redisClient.Ping(context.Background()).Result()
    if err != nil {
        fmt.Printf("Could not connect to Redis: %v\n", err)
        return nil
    }

    return &Store{
        db: redisClient,
    }
}

func (s *Store) SetX(ctx context.Context, key string, value interface{}, duration time.Duration) error {
    statusCmd := s.db.SetXX(ctx, key, value, duration)
    return statusCmd.Err()
}

func (s *Store) Get(ctx context.Context, key string) (string, error) {
    result, err := s.db.Get(ctx, key).Result()
    if err == redis.Nil {
        return "", fmt.Errorf("key does not exist")
    }
    return result, err
}

func (s *Store) Del(ctx context.Context, key string) error {
    statusCmd := s.db.Del(ctx, key)
    return statusCmd.Err()
}

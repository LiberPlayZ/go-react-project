package migrations

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

type AdminInformation struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func retry(attempts int, delay time.Duration, action func() error) error {
	var err error
	for i := 0; i < attempts; i++ {
		err = action()
		if err == nil {
			return nil
		}
		time.Sleep(delay)
	}
	return fmt.Errorf("all retry attempts failed: %w", err)
}

func EnsureEnvironment(admin AdminInformation, dbName, username, password string, tableQueries []string) (*sql.DB, error) {
	ctx := context.Background()

	adminConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		admin.Host, admin.Port, admin.User, admin.Password, admin.Database)
	adminDB, err := sql.Open("postgres", adminConnStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect as admin: %w", err)
	}
	defer adminDB.Close()

	var wg sync.WaitGroup
	var ensureErr error
	var mu sync.Mutex

	wg.Add(2)
	go func() {
		defer wg.Done()
		if err := ensureUser(ctx, adminDB, username, password); err != nil {
			mu.Lock()
			ensureErr = err
			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		if err := ensureDatabase(ctx, adminDB, dbName, username); err != nil {
			mu.Lock()
			ensureErr = err
			mu.Unlock()
		}
	}()

	wg.Wait()
	if ensureErr != nil {
		return nil, ensureErr
	}

	userDB, err := connectUserDB(admin.Host, admin.Port, username, password, dbName)
	if err != nil {
		return nil, err
	}

	schema := fmt.Sprintf("s_%s", username)

	_, err = userDB.ExecContext(ctx, fmt.Sprintf("SET search_path TO \"%s\"", schema))
	if err != nil {
		return nil, fmt.Errorf("failed to set session search_path: %w", err)
	}
	log.Printf("🔁 Session search_path changed to schema %s", schema)

	if err := ensureSchema(ctx, userDB, schema, username); err != nil {
		return nil, err
	}

	if err := ensureTablesSequential(ctx, userDB, tableQueries); err != nil {
		return nil, err
	}

	return userDB, nil
}

func ensureUser(ctx context.Context, db *sql.DB, username, password string) error {
	return retry(3, 2*time.Second, func() error {
		var exists bool
		err := db.QueryRowContext(ctx, "SELECT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = $1)", username).Scan(&exists)
		if err != nil {
			return err
		}

		if !exists {
			escapedPassword := strings.ReplaceAll(password, `'`, `''`)
			query := fmt.Sprintf(`CREATE ROLE "%s" WITH LOGIN PASSWORD '%s' CREATEDB NOINHERIT`, username, escapedPassword)
			_, err := db.ExecContext(ctx, query)
			if err != nil {
				return err
			}
			log.Printf("✅ User %s created successfully", username)
		} else {
			log.Printf("ℹ️ User %s already exists", username)
		}
		return nil
	})
}

func ensureDatabase(ctx context.Context, db *sql.DB, dbName, owner string) error {
	return retry(3, 2*time.Second, func() error {
		var exists bool
		err := db.QueryRowContext(ctx, "SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = $1)", dbName).Scan(&exists)
		if err != nil {
			return err
		}

		if !exists {
			query := fmt.Sprintf("CREATE DATABASE \"%s\" OWNER \"%s\"", dbName, owner)
			_, err := db.ExecContext(ctx, query)
			if err != nil {
				return err
			}
			log.Printf("✅ Database %s created successfully", dbName)
		} else {
			log.Printf("ℹ️ Database %s already exists", dbName)
		}
		return nil
	})
}

func connectUserDB(host string, port int, user, password, dbName string) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
	userDB, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to user DB: %w", err)
	}
	log.Printf("🔌 Connected to DB %s as user %s", dbName, user)
	return userDB, nil
}

func ensureSchema(ctx context.Context, db *sql.DB, schemaName, username string) error {
	return retry(3, 2*time.Second, func() error {
		var exists bool
		err := db.QueryRowContext(ctx, "SELECT EXISTS (SELECT 1 FROM information_schema.schemata WHERE schema_name = $1)", schemaName).Scan(&exists)
		if err != nil {
			return err
		}

		if !exists {
			query := fmt.Sprintf("CREATE SCHEMA \"%s\"", schemaName)
			_, err := db.ExecContext(ctx, query)
			if err != nil {
				return err
			}
			log.Printf("✅ Schema %s created successfully", schemaName)
		}

		alter := fmt.Sprintf("ALTER ROLE \"%s\" SET search_path TO \"%s\"", username, schemaName)
		_, err = db.ExecContext(ctx, alter)
		if err != nil {
			return err
		}
		log.Printf("🔁 Default path changed to schema %s", schemaName)
		return nil
	})
}

func ensureTablesSequential(ctx context.Context, db *sql.DB, tableQueries []string) error {
	return retry(3, 1*time.Second, func() error {
		for _, query := range tableQueries {
			_, err := db.ExecContext(ctx, query)
			if err != nil {
				return err
			}
		}
		log.Printf("✅ Tables created successfully")
		return nil
	})
}

// func EnsureSeedData(db *sql.DB, queries []string) error {
// 	ctx := context.Background()
// 	for _, q := range queries {
// 		if _, err := db.ExecContext(ctx, q); err != nil {
// 			return fmt.Errorf("failed to run seed query: %w", err)
// 		}
// 	}
// 	log.Printf("🌱 Seed data inserted successfully")
// 	return nil
// }

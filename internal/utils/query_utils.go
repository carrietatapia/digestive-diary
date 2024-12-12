package utils

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
	"unicode"

	"github.com/jackc/pgx/v4/pgxpool"
)

// BuildUpdateQuery builds a dynamic update query and parameters
func BuildUpdateQuery(table string, id string, fields map[string]interface{}) (string, []interface{}) {
	setClauses := []string{}
	params := []interface{}{id}
	paramCount := 1

	for field, value := range fields {
		paramCount++
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", field, paramCount))
		params = append(params, value)
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $1", table, strings.Join(setClauses, ", "))
	return query, params
}

// ExecuteUpdate executes the update query
func ExecuteUpdate(ctx context.Context, db *pgxpool.Pool, table string, id string, fields map[string]interface{}) error {
	if len(fields) == 0 {
		return nil
	}

	// Add validation for empty ID
	if id == "" {
		return fmt.Errorf("invalid empty ID provided for update")
	}

	query, params := BuildUpdateQuery(table, id, fields)
	log.Println("fields:", fields)
	log.Println("query:", query)
	log.Println("params:", params) // Add this line to debug parameters

	_, err := db.Exec(ctx, query, params...)
	if err != nil {
		return fmt.Errorf("failed to update %s: %w", table, err)
	}

	return nil
}

// BuildUpdateMap builds a map of fields to be updated using reflection
func BuildUpdateMap(user interface{}) map[string]interface{} {
	fields := make(map[string]interface{})
	val := reflect.ValueOf(user).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Tag.Get("db")

		// If no db tag is present, convert the field name to snake_case
		if fieldName == "" {
			fieldName = toSnakeCase(typ.Field(i).Name)
		}

		// Only add non-zero fields to the map
		if !field.IsZero() {
			fields[fieldName] = field.Interface()
		}
	}

	return fields
}

// Add this helper function
func toSnakeCase(str string) string {
	var result strings.Builder
	for i, r := range str {
		if i > 0 && unicode.IsUpper(r) {
			result.WriteRune('_')
		}
		result.WriteRune(unicode.ToLower(r))
	}
	return result.String()
}

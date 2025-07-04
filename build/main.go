package main

		import (
			"sort"
			"fmt"
			"time"
			"bufio"
			"os"
			"strings"
			"crypto/sha256"
			"math"
			"math/rand"
			"encoding/json"
			"net"
			"net/http"
			"strconv"
			"errors"
			"database/sql"
			"path/filepath"
			"bytes"

			"github.com/golang-jwt/jwt/v5"
			_ "github.com/go-sql-driver/mysql"
		)

		

	func sizeOf(value interface{}) int {
		switch v := value.(type) {
		case []interface{}:
			return len(v)
		case string:
			return len(v)
		default:
			return 0
		}
	}

	func toUppercase(s string) string {
		return strings.ToUpper(s)
	}

	func toLowercase(s string) string {
		return strings.ToLower(s)
	}

	func capitalize(s string) string {
		if len(s) == 0 {
			return s
		}
		return strings.ToUpper(s[:1]) + s[1:]
	}

	func removeWhiteSpaces(s string) string {
		return strings.ReplaceAll(s, " ", "")
	}

	func replace(s, old, new string) string {
		return strings.ReplaceAll(s, old, new)
	}

	func addToArrayStart(arr []interface{}, elem interface{}) []interface{} {
		return append([]interface{}{elem}, arr...)
	}

	func addToArrayEnd(arr []interface{}, elem interface{}) []interface{} {
		return append(arr, elem)
	}

	func removeFromArray(arr []interface{}, index int) []interface{} {
		if index < 0 || index >= len(arr) {
			return arr
		}
		return append(arr[:index], arr[index+1:]...)
	}

	func max(arr []interface{}) interface{} {
		if len(arr) == 0 {
			return nil
		}
		maxVal := arr[0].(int)
		for _, v := range arr[1:] {
			val := v.(int)
			if val > maxVal {
				maxVal = val
			}
		}
		return maxVal
	}

	func min(arr []interface{}) interface{} {
		if len(arr) == 0 {
			return nil
		}
		minVal := arr[0].(int)
		for _, v := range arr[1:] {
			val := v.(int)
			if val < minVal {
				minVal = val
			}
		}
		return minVal
	}


	func first(arr []interface{}) interface{} {
		if len(arr) == 0 {
			return nil
		}
		return arr[0]
	}

	func last(arr []interface{}) interface{} {
		if len(arr) == 0 {
			return nil
		}
		return arr[len(arr)-1]
	}


	func allButFirst(arr []interface{}) []interface{} {
		if len(arr) == 0 {
			return arr
		}
		return arr[1:]
	}


	func indexOf(arr []interface{}, elem interface{}) int {
		for i, v := range arr {
			if v == elem {
				return i
			}
		}
		return -1
	}

	func organize(arr []interface{}, order string) []interface{} {
		intArr := make([]int, len(arr))
		for i, v := range arr {
			intArr[i] = v.(int)
		}
		if order == "desc" {
			sort.Sort(sort.Reverse(sort.IntSlice(intArr)))
		} else {
			sort.Ints(intArr)
		}
		result := make([]interface{}, len(intArr))
		for i, v := range intArr {
			result[i] = v
		}
		return result
	}

	func sum(arr []interface{}) interface{} {
		total := 0.0
		for _, v := range arr {
			switch val := v.(type) {
			case int:
				total += float64(val)
			case float64:
				total += val
			}
		}
		if float64(int(total)) == total {
			return int(total)
		}
		return total
	}

	type ZumbraDate struct {
		fullDate time.Time
		hour     int
		minute   int
		second   int
		day      int
		month    int
		year     int
	}

	func date() ZumbraDate {
		now := time.Now()
		return ZumbraDate{
			fullDate: now,
			hour:     now.Hour(),
			minute:   now.Minute(),
			second:   now.Second(),
			day:      now.Day(),
			month:    int(now.Month()),
			year:     now.Year(),
		}
	}

	func addToDict(dict map[string]interface{}, key string, value interface{}) map[string]interface{} {
		dict[key] = value
		return dict
	}

	func deleteFromDict(dict map[string]interface{}, key string) map[string]interface{} {
		delete(dict, key)
		return dict
	}

	func getFromDict(dict map[string]interface{}, key string) interface{} {
		return dict[key]
	}

	func dictKeys(dict map[string]interface{}) []string {
		keys := make([]string, 0, len(dict))
		for k := range dict {
			keys = append(keys, k)
		}
		return keys
	}

	func dictValues(dict map[string]interface{}) []interface{} {
		values := make([]interface{}, 0, len(dict))
		for _, v := range dict {
			values = append(values, v)
		}
		return values
	}

	var EnvVars = map[string]string{}

	func dotenvLoad(filepath string) {
		file, err := os.Open(filepath)
		if err != nil {
			fmt.Println("failed to open file:", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
				continue
			}
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				EnvVars[key] = value
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("failed to read file:", err)
		}
	}

	func dotenvGet(key string) string {
		return EnvVars[key]
	}

	func hashCode(input string) string {
		hash := sha256.New()
		hash.Write([]byte(input))
		hashInBytes := hash.Sum(nil)
		return fmt.Sprintf("%x", hashInBytes)
	}

	func input(prompt ...string) string {
		if len(prompt) > 0 {
			fmt.Print(prompt[0])
		}
		var value string
		fmt.Scanln(&value)
		return value
	}

	func bhaskara(a, b, c float64) interface{} {
		delta := (b * b) - (4 * a * c)
		if delta < 0 {
			return nil
		}
		if delta == 0 {
			return -b / (2 * a)
		}
		sqrtDelta := math.Sqrt(delta)
		x1 := (-b + sqrtDelta) / (2 * a)
		x2 := (-b - sqrtDelta) / (2 * a)
		return []interface{}{x1, x2}
	}

	func randomInteger(args ...int) int {
		min := 0
		max := 10
		if len(args) == 1 {
			max = args[0]
		} else if len(args) == 2 {
			min = args[0]
			max = args[1]
		}
		if min > max {
			min, max = max, min
		}
		return min + rand.Intn(max-min+1)
	}

	func randomFloat(args ...float64) float64 {
		min := 0.0
		max := 10.0
		if len(args) == 1 {
			max = args[0]
		} else if len(args) == 2 {
			min = args[0]
			max = args[1]
		}
		if min > max {
			min, max = max, min
		}
		return min + rand.Float64()*(max-min)
	}

	func toString(value interface{}) string {
		return fmt.Sprintf("%v", value)
	}

	func toInt(value interface{}) int {
		switch v := value.(type) {
		case string:
			n, err := strconv.Atoi(v)
			if err != nil {
				return 0
			}
			return n
		case float64:
			return int(math.Floor(v))
		case bool:
			if v {
				return 1
			}
			return 0
		case int:
			return v
		default:
			return 0
		}
	}

	func toFloat(value interface{}) float64 {
		switch v := value.(type) {
		case string:
			n, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return 0
			}
			return n
		case float64:
			return v
		case bool:
			if v {
				return 1.0
			}
			return 0.0
		case int:
			return float64(v)
		default:
			return 0.0
		}
	}

	func toBool(value interface{}) bool {
		switch v := value.(type) {
		case string:
			return v != ""
		case float64:
			return v != 0
		case bool:
			return v
		case int:
			return v != 0
		default:
			return false
		}
	}

	func json_parse(input string) map[string]interface{} {
		var result map[string]interface{}
		err := json.Unmarshal([]byte(input), &result)
		if err != nil {
			return map[string]interface{}{}
		}
		return result
	}

	var secretKey string

	func jwtCreateToken(username string, secret string, expirationHours int) (string, error) {
		secretKey = secret

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * time.Duration(expirationHours)).Unix(),
		})

		tokenStr, err := token.SignedString([]byte(secretKey))
		if err != nil {
			return "", fmt.Errorf("failed to create token: %v", err)
		}

		return tokenStr, nil
	}

	func jwtVerifyToken(tokenStr string) (string, error) {
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			return "", fmt.Errorf("failed to parse token: %v", err)
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			username, ok := claims["username"].(string)
			if !ok {
				return "", errors.New("username not found in token")
			}
			return username, nil
		}

		return "", errors.New("invalid token")
	}

	var db_connection *sql.DB

	func mysqlConnection(host, port, user, password, database string) error {
		var err error
		db_connection, err = sql.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+database)
		if err != nil {
			return fmt.Errorf("failed to connect: %v", err)
		}

		err = db_connection.Ping()
		if err != nil {
			return fmt.Errorf("failed to ping: %v", err)
		}

		fmt.Printf("Database '%s' connected successfully\n", database)
		return nil
	}

	func mysqlCreateTable(tableName, fields string) error {
		_, err := db_connection.Exec("CREATE TABLE " + tableName + " (" + fields + ");")
		if err != nil {
			return fmt.Errorf("failed to create table: %v", err)
		}
		fmt.Printf("Table '%s' created successfully\n", tableName)
		return nil
	}

	func mysqlShowTables() ([]string, error) {
		rows, err := db_connection.Query("SHOW TABLES")
		if err != nil {
			return nil, fmt.Errorf("failed to show tables: %v", err)
		}
		defer rows.Close()

		var tables []string
		for rows.Next() {
			var table string
			if err := rows.Scan(&table); err != nil {
				return nil, err
			}
			tables = append(tables, table)
		}
		return tables, nil
	}

	func mysqlShowTableColumns(tableName string) ([]string, error) {
		rows, err := db_connection.Query("SHOW COLUMNS FROM " + tableName)
		if err != nil {
			return nil, fmt.Errorf("failed to show columns: %v", err)
		}
		defer rows.Close()

		var columns []string
		for rows.Next() {
			var field, colType, null, key, extra string
			var defaultValue sql.NullString
			if err := rows.Scan(&field, &colType, &null, &key, &defaultValue, &extra); err != nil {
				return nil, err
			}
			columns = append(columns, field)
		}
		return columns, nil
	}

	func mysqlDeleteTable(tableName string) error {
		query := fmt.Sprintf("DROP TABLE %s", tableName)
		_, err := db_connection.Exec(query)
		if err != nil {
			return fmt.Errorf("failed to drop table: %v", err)
		}
		fmt.Printf("Table '%s' deleted successfully\n", tableName)
		return nil
	}

	func mysqlInsertIntoTable(tableName string, data map[string]interface{}) error {
		keys := []string{}
		placeholders := []string{}
		args := []interface{}{}

		for key, value := range data {
			keys = append(keys, key)
			placeholders = append(placeholders, "?")
			args = append(args, value)
		}

		query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(keys, ","), strings.Join(placeholders, ","))
		_, err := db_connection.Exec(query, args...)
		if err != nil {
			return fmt.Errorf("failed to insert: %v", err)
		}

		fmt.Println("Record inserted successfully")
		return nil
	}

	func mysqlGetFromTable(tableName, fields, condition string) ([]map[string]interface{}, error) {
		if db_connection == nil {
			return nil, errors.New("not connected")
		}

		query := fmt.Sprintf("SELECT %s FROM %s", fields, tableName)
		if condition != "" {
			query += " WHERE " + condition
		}

		rows, err := db_connection.Query(query)
		if err != nil {
			return nil, fmt.Errorf("failed to query: %v", err)
		}
		defer rows.Close()

		columns, _ := rows.Columns()
		result := []map[string]interface{}{}

		for rows.Next() {
			values := make([]interface{}, len(columns))
			valuePtrs := make([]interface{}, len(columns))

			for i := range values {
				valuePtrs[i] = &values[i]
			}

			if err := rows.Scan(valuePtrs...); err != nil {
				return nil, err
			}

			rowMap := map[string]interface{}{}
			for i, col := range columns {
				val := values[i]
				b, ok := val.([]byte)
				if ok {
					rowMap[col] = string(b)
				} else {
					rowMap[col] = val
				}
			}
			result = append(result, rowMap)
		}

		return result, nil
	}

	func mysqlUpdateIntoTable(tableName string, data map[string]interface{}, condition string) error {
		assignments := []string{}
		args := []interface{}{}

		for key, value := range data {
			assignments = append(assignments, fmt.Sprintf("%s = ?", key))
			args = append(args, value)
		}

		query := fmt.Sprintf("UPDATE %s SET %s", tableName, strings.Join(assignments, ", "))
		if condition != "" {
			query += " WHERE " + condition
		}

		_, err := db_connection.Exec(query, args...)
		if err != nil {
			return fmt.Errorf("failed to update: %v", err)
		}
		fmt.Println("Record updated successfully")
		return nil
	}

	func mysqlDeleteFromTable(tableName, condition string) error {
		query := fmt.Sprintf("DELETE FROM %s", tableName)
		if condition != "" {
			query += " WHERE " + condition
		}
		_, err := db_connection.Exec(query)
		if err != nil {
			return fmt.Errorf("failed to delete: %v", err)
		}
		fmt.Println("Record deleted successfully")
		return nil
	}

	type Route struct {
		Method  string
		Path    string
		Handler func(http.ResponseWriter, *http.Request)
	}

	var routes []Route
	var staticRoutes []StaticRoute

	type StaticRoute struct {
		Prefix string
		Dir    string
	}

	func server(port int) {
		for _, sr := range staticRoutes {
			http.Handle(sr.Prefix+"/", http.StripPrefix(sr.Prefix, http.FileServer(http.Dir(sr.Dir))))
		}

		for _, r := range routes {
			http.HandleFunc(r.Path, func(w http.ResponseWriter, req *http.Request) {
				if req.Method != r.Method {
					http.NotFound(w, req)
					return
				}
				r.Handler(w, req)
			})
		}

		addr := fmt.Sprintf(":%d", port)
		ln, err := net.Listen("tcp", addr)
		if err != nil {
			fmt.Printf("Failed to bind to port %s. got %s\n", addr, err)
			return
		}

		fmt.Printf("Zumbra server started on port %d\n", port)
		http.Serve(ln, nil)
	}

	func registerRoute(path string, handler string) {
		routes = append(routes, Route{
			Method:  "GET",
			Path:    path,
			Handler: func(w http.ResponseWriter, req *http.Request) { w.Write([]byte(handler)) },
		})
	}

	func html(content string) string {
		return content
	}

	func serveStatic(prefix, dir string) {
		staticRoutes = append(staticRoutes, StaticRoute{
			Prefix: prefix,
			Dir:    dir,
		})
	}

	func get(url string) map[string]string {
		resp, err := http.Get(url)
		if err != nil {
			return map[string]string{"error": err.Error()}
		}
		defer resp.Body.Close()

		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(resp.Body)
		if err != nil {
			return map[string]string{"error": err.Error()}
		}

		return map[string]string{"body": buf.String()}
	}

	func logger(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Request:", r.Method, r.URL.Path)
			next(w, r)
		}
	}

	func serveFile(args ...interface{}) string {
		if len(args) != 1 && len(args) != 2 {
			return ""
		}

		path, ok := args[0].(string)
		if !ok {
			return ""
		}

		cleanPath := filepath.Clean(path)
		content, err := os.ReadFile(cleanPath)
		if err != nil {
			return ""
		}

		html := string(content)

		if len(args) == 1 {
			return html
		}

		dict, ok := args[1].(map[string]interface{})
		if !ok {
			return html
		}

		for key, val := range dict {
			strVal, ok := val.(string)
			if ok {
				placeholder := "{{" + key + "}}"
				html = strings.ReplaceAll(html, placeholder, strVal)
			}
		}

		return html
	}




		func main() {
			    var a = date()
    fmt.Println(a)
    fmt.Println(a.hour)
    fmt.Println(a.minute)
    fmt.Println(a.second)
    fmt.Println(a.day)
    fmt.Println(a.month)
    fmt.Println(a.year)
		}
	
package redisService

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/RolloCasanova/dispatch-workshop-4/model"
	"github.com/go-redis/redis"
)

// employeeRedisService struct consists of the redis client.
type employeeRedisService struct {
	client *redis.Client
}

// New returns a new employeeRedisService instance.
func New(host, password string, port, db int) employeeRedisService {
	log.Println("In redis service - New")

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		DB:       db,
		Password: password,
	})

	// if err := client.Ping().Err(); err != nil {
	// 	log.Panicf("Unable to connect to redis: %v", err)
	// }

	return employeeRedisService{
		client: client,
	}
}

// UpsertEmployee upserts an employee in the redis cache.
func (rs employeeRedisService) UpsertEmployee(e model.Employee) error {
	log.Println("In redis service - UpsertEmployee")

	employeeEntry, err := json.Marshal(e)
	if err != nil {
		return fmt.Errorf("marshalling employee in redis: %v", err)
	}

	key := fmt.Sprintf("%d", e.ID)
	err = rs.client.Set(key, employeeEntry, 0).Err()
	if err != nil {
		return fmt.Errorf("inserting value in redis: %v", err)
	}

	return nil
}

// UpsertEmployees upserts a list of employees in the redis cache.
func (rs employeeRedisService) UpsertEmployees(employees model.Employees) error {
	log.Println("In redis service - UpsertEmployees")

	// new map to store the employees
	employeeEntries := make(map[string][]byte, len(employees))

	// iterate over all employees and flat them into the map
	for _, e := range employees {
		employeeEntry, err := json.Marshal(e)
		if err != nil {
			return fmt.Errorf("marshalling an employee in multiple upsert redis: %v", err)
		}

		key := fmt.Sprintf("%d", e.ID)
		employeeEntries[key] = employeeEntry
	}

	// iterate over flatten map and upsert each employee
	for key, employeeEntry := range employeeEntries {
		err := rs.client.Set(key, employeeEntry, 0).Err()
		if err != nil {
			return fmt.Errorf("inserting a value in multiple upsert redis: %v", err)
		}
	}

	return nil
}

// ReadEmployee reads an employee from the redis cache by its ID.
func (rs employeeRedisService) ReadEmployee(id int) (*model.Employee, error) {
	log.Println("In redis service - ReadEmployee")

	employee, err := rs.client.Get(fmt.Sprintf("%d", id)).Result()
	if err != nil || err == redis.Nil {
		return nil, fmt.Errorf("getting employee from redis: %v", err)
	}

	var e model.Employee
	err = json.Unmarshal([]byte(employee), &e)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling employee from redis: %v", err)
	}

	return &e, nil
}

// DeleteEmployeeByID deletes an employee from the redis cache by its ID.
func (rs employeeRedisService) DeleteEmployeeByID(id int) error {
	log.Println("In redis service - DeleteEmployeeByID")

	err := rs.client.Del(fmt.Sprintf("%d", id)).Err()
	if err != nil {
		return fmt.Errorf("deleting employee from redis: %v", err)
	}

	return nil
}

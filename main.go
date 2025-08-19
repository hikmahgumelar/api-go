package main

import (
	"net/http"
	"api-go/database"
	"api-go/models"
  "strconv"
	"github.com/gin-gonic/gin"
)

var db = database.SetupDB()

func main() {
	r := gin.Default()

	// Routes
	r.POST("/users", CreateUser)
	r.GET("/users", GetUsers)
	r.GET("/users/:id", GetUserByID)
	r.PUT("/users/:id", UpdateUser)
	r.DELETE("/users/:id", DeleteUser)

	r.Run(":8080")
}

// CREATE User
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	result, err := db.Exec(query, user.Name, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	user.ID = int(id)

	c.JSON(http.StatusOK, user)
}

// GET All Users
func GetUsers(c *gin.Context) {
	var users []models.User

	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		rows.Scan(&user.ID, &user.Name, &user.Email)
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}

// DELETE User by ID
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	// Cek dulu apakah user ada
	var user models.User
	err := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Kalau ada, baru delete
	_, err = db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}



// GET User by ID
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	err := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}


// UPDATE User
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	// Bind input JSON
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cek apakah user ada
	var existingUser models.User
	err := db.QueryRow("SELECT id FROM users WHERE id = ?", id).Scan(&existingUser.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Update data
	_, err = db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.ID, _ = strconv.Atoi(id)
	c.JSON(http.StatusOK, user)
}

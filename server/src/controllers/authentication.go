package controllers

import (
	"MyVPN/models"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Authenticator struct {
	mu    sync.RWMutex
	users map[string]string
}

func NewAuthenticator() (*Authenticator, error) {
	var auth Authenticator
	users, err := models.LoadUsers()
	if err != nil {
		return &auth, err
	}

	// Fill the map with all the registerd users and their passwords
	auth.users = make(map[string]string)

	for _, user := range users {
		auth.users[user.Username] = user.PasswordHash
	}

	return &auth, nil
}

func (auth *Authenticator) SignIn(c *gin.Context) {

	// Lock the authenticator as writer
	auth.mu.Lock()
	defer auth.mu.Unlock()

	// Parse the JSON file in the request as a LoginRequest structure
	var req models.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if the username is already used
	_, ok := auth.users[req.Username]
	if ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already in use"})
		return
	}

	// Hash the password, with salt
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate password hash"})
		return
	}

	// Add the new user to the database of registered users and to the cache
	auth.users[req.Username] = string(passwordHash)
	models.AddNewUser(req.Username, string(passwordHash))

	c.JSON(http.StatusCreated, gin.H{"username": req.Username})

	auth.mu.Unlock()
}

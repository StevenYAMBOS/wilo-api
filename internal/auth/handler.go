package auth

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// Inscription
func Register(c *gin.Context) {
    var req struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Requête invalide"})
        return
    }

    err := RegisterUser(req.Email, req.Password)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Utilisateur inscrit avec succès"})
}

// Connexion
func Login(c *gin.Context) {
    var req struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Requête invalide"})
        return
    }

    token, err := LoginUser(req.Email, req.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

// Page d'accueil
func Dashboard(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Bienvenue sur votre dashboard !"})
}

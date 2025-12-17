package app

import (
	"context"

	"github.com/gin-gonic/gin"

	"gym/ent/user"
	"gym/ent/wallet"
	"gym/models"
)

func (a *App) UserHandler(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := a.Client.User.Create().
		SetName(input.Name).
		SetLastName(input.Lastname).
		SetPassword(input.Password).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, result)
}

func (a *App) CoachHandler(c *gin.Context) {
	var input models.Coach
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := a.Client.Coach.Create().
		SetName(input.Name).
		SetPrice(input.Price).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, result)
}

func (a *App) CourseHandler(c *gin.Context) {
	var input models.Course
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := a.Client.Course.Create().
		SetName(input.Name).
		SetPrice(input.Price).
		SetDay(input.Day).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, result)
}

func (a *App) WalletHandler(c *gin.Context) {
	userID, ok := a.Check(c)
	if !ok {
		return
	}

	var input models.Wallet
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := a.Client.Wallet.Create().
		SetAmount(input.Amount).
		SetUserid(userID).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "wallet updated"})
}

func (a *App) LoginHandler(c *gin.Context) {
	var input models.Login
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	record, err := a.Client.User.Get(context.Background(), input.User)
	if err != nil || record.Password != input.Password {
		c.JSON(401, gin.H{"error": "wrong user or password"})
		return
	}

	token, _ := a.CreateToken(record.ID)
	c.String(200, token)
}

func (a *App) ShowHandler(c *gin.Context) {
	coaches, _ := a.Client.Coach.Query().All(context.Background())
	courses, _ := a.Client.Course.Query().All(context.Background())

	c.JSON(200, gin.H{
		"coaches": coaches,
		"courses": courses,
	})
}

func (a *App) MainHandler(c *gin.Context) {
	userID, ok := a.Check(c)
	if !ok {
		return
	}

	var input models.Selectt
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	w, _ := a.Client.Wallet.
		Query().
		Where(wallet.UseridEQ(userID)).
		Only(context.Background())

	coach, _ := a.Client.Coach.Get(context.Background(), input.Coachid)
	course, _ := a.Client.Course.Get(context.Background(), input.Courseid)

	total := coach.Price + course.Price
	if total > w.Amount {
		c.JSON(401, gin.H{"error": "not enough money"})
		return
	}

	a.Client.Wallet.Update().
		Where(wallet.UseridEQ(userID)).
		SetAmount(w.Amount - total).
		Save(context.Background())

	a.Client.User.Update().
		Where(user.IDEQ(userID)).
		SetCoachID(coach.ID).
		SetCourseID(course.ID).
		Save(context.Background())

	c.JSON(200, gin.H{"message": "success"})
}

func (a *App) ShowUser(c *gin.Context) {
	userID, ok := a.Check(c)
	if !ok {
		return
	}

	result, _ := a.Client.User.Get(context.Background(), userID)
	c.JSON(200, result)
}

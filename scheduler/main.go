package main

import (
	"context"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"time"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"github.com/jordan-wright/email"
	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"google.golang.org/api/option"
)

const (
	prompt = `
		Write an email in form of a newsletter about what happened in the MLB this week.
		The newsletter should be in three sections: A general summary of the MLB, 
		a summary of how the following teams did: %s,
		a summary of how the follwing players did: %s.
		`
)

var (
	senderEmail = os.Getenv("SMTP_EMAIL")
	senderPass  = os.Getenv("SMTP_PASS")
	mongoURI    = os.Getenv("MONGODB_ATLAS_URI")
	smtpHost    = "smtp.example.com"
	smtpPort    = "587"
	subject     = "Weekly Update"
	dbName      = "fbdr"
	collection  = "fbdr"
)

func SetupGeminiClient() (*genai.Client, error) {
	ctx := context.Background()
	apiKey := os.Getenv("GEMINI_API_KEY")

	if apiKey == "" {
		return nil, fmt.Errorf("GEMINI_API_KEY environment variable not set")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create genai client: %v", err)
	}

	return client, nil
}

// Sends a POST request to Gemini API to generate email
func callGeminiAPI(model *genai.GenerativeModel, ctx *context.Context, teams []string, players []string) (string, error) {
	model.GenerateContent(ctx, fmt.Sprintf(prompt, teams, players))

	return resp.Text, nil
}

func getRecipients() ([]string, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection(collection)
	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var recipients []string
	for cursor.Next(ctx) {
		var result struct {
			Email string `bson:"email"`
		}
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		recipients = append(recipients, result.Email)
	}
	return recipients, nil
}

func sendEmail() {
	var err error
	e := email.NewEmail()
	e.From = fmt.Sprintf("Your Name <%s>", senderEmail)
	e.To, err = getRecipients()
	e.Subject = subject
	e.Text = []byte(body)

	if err != nil {
		log.Printf("Failed to get recipients: %v", err)
		return
	}
	auth := smtp.PlainAuth("", senderEmail, senderPass, smtpHost)

	err = e.Send(smtpHost+":"+smtpPort, auth)
	if err != nil {
		log.Printf("Failed to send email: %v", err)
	} else {
		log.Println("Weekly email sent successfully")
	}
}

func main() {
	godotenv.Load()
	// Load credentials from environment variables (recommended for security)
	if envEmail := os.Getenv("SMTP_EMAIL"); envEmail != "" {
		senderEmail = envEmail
	}
	if envPass := os.Getenv("SMTP_PASS"); envPass != "" {
		senderPass = envPass
	}

	c := cron.New()
	c.AddFunc("0 0 * * 6", sendEmail) // Runs at midnight every Saturday
	c.Start()

	log.Println("Email scheduler started. Press Ctrl+C to exit.")
	select {} // Keep the program running
}

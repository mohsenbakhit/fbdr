from google import genai
from dotenv import load_dotenv
from os import getenv
import pymongo
import smtplib, ssl
from email.mime.text import MIMEText
load_dotenv()

class User:
    def __init__(self, email, players, teams):
        self.email = email
        self.players = players
        self.teams = teams
    def getEmail():
        return self.email
    def getPlayers():
        return self.players
    def getTeams():
        return self.teams

MongoClient = pymongo.MongoClient(getenv("ATLAS_URI"))
GeminiClient = genai.Client(api_key=getenv("GENAI_API_KEY"))
db = MongoClient["data"]
cluster = db["users"]
users = []
emails = []
for user in cluster.find():
    users.append(User(user["email"], user["favoritePlayers"], user["favoriteTeams"]))
    emails.append(user["email"])

FROM = 'weekly@fbdr.com'


SUBJECT = "Your Weekly MLB Report"


for user in users:
    s = smtplib.SMTP('smtp.gmail.com', 587)
    s.starttls()
    s.login(getenv("SENDER_EMAIL"), getenv("SENDER_PASSWORD"))
    response = GeminiClient.models.generate_content(
    model="gemini-2.0-flash", contents="Generate me a newsletter of the MLB this week. .It should have three sections: general, favorite teams, and favorite players. These are the players: {user.players} and the teams: {user.teams}. ",)
    s.sendmail(FROM, emails, response.text.encode('utf-8'))
    s.quit()
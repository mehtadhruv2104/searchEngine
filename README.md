The task is to make an efficient search engine that can read into given files and search for the needed keywords

The design for this search engine is simple - We are gonna use the Inverted index Algorithm for fast and efficient search through the whole files.
The program will parse through the files and set the inverted index on startup
The user can send two main ways to search  
- query string : this can contain any combination of strings that the user wants to search for.
- filter : This filter will have the basic mostly prompted fields - namespace, severity, appName and MsgID

To start the server -
1. Clone the Repo
2. Go to cd/backend
3. Add the Parquet Files in backend/data folder
4. Run go mod tidy
5. Run go run main.go


To start the client - 

1. Clone the Repo
2. Go to cd/frontend
3. Add the backend server BaseURL in .env file  
4. Run npm install
5. Run npm run dev

https://github.com/user-attachments/assets/550ef350-964c-4a71-b7ab-f08bcccceb85


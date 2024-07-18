# SecretNote-API-SPA-Rowan
A web application that allows users to securely share self-destructing secret notes using Golang and Vue.js

## Backend (Go Lang)

### Server Commands
1. To Run Server:
   ```
   cd backend && go run cmd/server.go && cd ..
   ```
2. To Format Backend:
   ```
   cd backend && go fmt ./... && cd ..
   ```
## DataBase (Sqlite Using GORM)
1. The database is configured using an ORM (GORM) --> helped in giving a common interface to deal with various SQLs
2. Database used is sqLite
<!-- 3. 
4.  -->

## Frontend (Vue + Vite)
### Frontend Creation Commands
1. To Create Project: (use sudo if needed)
   ```
   sudo npm create vite@latest frontend
   ```
2. To Install dependencies: (use sudo if needed)
   ```
   cd frontend
   sudo npm install
   #might need: sudo npm cache clean --force
   ```
3. To Run frontend
   ```
   sudo npm run dev
   ```

## Things Learnt from this project (Notes)

1. Dealing with JSON data with gin server and golang in general
2. Understanding what JWT is:
   1. JWT is for authorization(making sure it's the same user who logged in at authentication)
3. Dealing with 1 to many relation in Database
4. Learning about Queries
5. Learning about Sqlite in particular and its files/in memory feature
6. Trying out Vue using Vite (JS build tool)
<!-- 7. Learning Typescript -->

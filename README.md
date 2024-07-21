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

## Dockerization
### How Does The Dockerfile Build? And Relative to What?
- When typing:
   ```
   rowan@rowan:~/Desktop/Rowan's Internship/SecretNote-API-SPA-Rowan$ sudo docker build --tag=secretnote-spa-backend . -f backend/Dockerfile-backend
   ```
   This didn't work and resulted in:
   ```
   ERROR: failed to solve: process "/bin/sh -c go build -o /cmd ./cmd/server.go" did not complete successfully: exit code: 1 
   ```
- Because the directory of typing command is the starting point of file system from container POV:
  - so, the RUN command won't run starting from backend on which Dockerfile is located on.
  - It's relative path starts from point of declaring the sudo docker build command
- Thus, we don't have a /cm directory from the root repo position, that's why it led to an error.
- When typing:
  ```
  cd backend && sudo docker run -p 8080:8080 secretnote-spa-backend
  ```
  It worked just fine
- Note that we could've located the dockerfile on thr root of repo, and declare relative paths, however go build doesn't work correctly if current path doesn't contain go mod file
- Edit dockerfile to:
  ```
  RUN cd backend && go build -o /cmd ./cmd/server.go
  ```
  In terminal:
  ```
  sudo docker run -p 8080:8080 secretnote-root-trial
  ```
  Worked just fine too!
  
## Things Learnt from this project (Notes)

1. Dealing with JSON data with gin server and golang in general
2. Understanding what JWT is:
   1. JWT is for authorization(making sure it's the same user who logged in at authentication)
3. Dealing with 1 to many relation in Database
4. Learning about Queries
5. Learning about Sqlite in particular and its files/in memory feature
6. Trying out Vue using Vite (JS build tool)
7. Understanding dockerfile positions and containers realtive to directory
<!-- 7. Learning Typescript -->

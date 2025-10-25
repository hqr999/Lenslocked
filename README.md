# Lenslocked 

## Description

Lenslocked is a photo sharing platform made with Go on the backend and with Tailwind as our front-end. It uses Gooses with Docker for automatic creation of our database tables. We use mailtrap for the password backup process.

### How to use it  
First create and fill a  .env file. Copy the .env.template to know what fields  need to be filled.

After that, run docker-compose that will setup, our db,tailwind and adminer(a simple front-end to see the data stored on our database):
    
    docker compose up 

Open another terminal, while docker runs, and run our server with:

    go run cmd/server/server.go


### Screenshots

* Main Page
![image1](screenshots/screen1.jpg)

After loging in with a valid user, that has galleries created.

* My galleries page
![image6](screenshots/screen6.jpg)

* Books gallery screenshot 1 
![image7](screenshots/screen7.jpg)

* Books gallery screenhsot 2
![image8](screenshots/screen8.jpg)

## References 

* https://github.com/go-chi/chi

* https://github.com/cortesi/modd

* https://github.com/pressly/goose

* https://github.com/joho/godotenv

* https://caddyserver.com/

* https://mailtrap.io/

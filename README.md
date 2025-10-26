# Lenslocked 

## Description

Lenslocked is a photo sharing platform made with Go on the backend and with Tailwind as our front-end. It uses Gooses with Docker for automatic creation of our database tables. We use mailtrap for the password backup process.

### How to use it  
First create and fill a  .env file. Copy the .env.template to know what fields  need to be filled.

After that, run docker-compose that will setup, our db,tailwind and adminer(a simple front-end to see the data stored on our database):
    
    docker compose up 

Open another terminal, while docker runs, and run our server with:

    go run cmd/server/server.go

## Screenshots 

Down below are the pages of our application when it up and running, i also added an example of what a gallery page looks like.

* Main Page

![img1](screenshots/screen1.jpg)

* Contact Page

![img2](screenshots/screen2.jpg)

* FAQ Page 

![img3](screenshots/screen3.jpg)


* Sign in Page 

![img4](screenshots/screen4.jpg)

* Sign Up Page 

![img5](screenshots/screen5.jpg)

* Galleries of the user Page 

![img6](screenshots/screen6.jpg)

* Example of a Gallery Page 

![img7](screenshots/screen7.jpg)

* Example two of a Gallery Page 

![img8](screenshots/screen8.jpg)

## References 

* https://github.com/go-chi/chi

* https://github.com/cortesi/modd

* https://github.com/pressly/goose

* https://github.com/joho/godotenv

* https://caddyserver.com/

* https://mailtrap.io/

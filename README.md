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
Down below you can see some of the screens of the app when it runs. I created a simple user and a gallery with some pics for the example:

* Home Page 

<img src="screenshots/screen1.jpg" alt="Descrição da imagem" width="500" height="300">
        
* Contact page

<img src="screenshots/screen2.jpg" alt="Descrição da imagem" width="500" height="300">

* FAQ page

<img src="screenshots/screen3.jpg" alt="Descrição da imagem" width="500" height="300">

* Sign in page

<img src="screenshots/screen4.jpg" alt="Descrição da imagem" width="500" height="300">

* Sign up page

<img src="screenshots/screen5.jpg" alt="Descrição da imagem" width="500" height="300">

* Your Galleries page

<img src="screenshots/screen6.jpg" alt="Descrição da imagem" width="500" height="300">

* Screenshot 1 of the Books gallery

<img src="screenshots/screen7.jpg" alt="Descrição da imagem" width="500" height="300">

* Screenshot 2 of the Books gallery

<img src="screenshots/screen8.jpg" alt="Descrição da imagem" width="500" height="300">

## References 

* https://github.com/go-chi/chi

* https://github.com/cortesi/modd

* https://github.com/pressly/goose

* https://github.com/joho/godotenv

* https://caddyserver.com/

* https://mailtrap.io/

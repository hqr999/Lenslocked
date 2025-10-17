FROM node:latest 
WORKDIR /tailwind 
RUN npm init -y && \
    npm install tailwindcss @tailwindcss/cli 
   
CMD npx @tailwindcss/cli -i /src/style.css -o /dst/styles.css --watch --poll


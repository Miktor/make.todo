# syntax=docker/dockerfile:1

# pull official base image
FROM node:latest

# set working directory
WORKDIR /app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# install app dependencies
COPY . ./
RUN npm install

# start app
CMD ["npm", "start"]    
# syntax=docker/dockerfile:1

# pull official base image
FROM node:latest

# set working directory
WORKDIR /app

COPY ./package.json ./yarn.lock ./
RUN yarn install --frozen-lockfile

# add `/app/node_modules/.bin` to $PATH
ENV PATH="./node_modules/.bin:$PATH"

# install app dependencies
COPY . ./

# start app
CMD npm start  

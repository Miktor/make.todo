# syntax=docker/dockerfile:1
FROM ruby:3.1.2
RUN apt-get update -qq && apt-get install -y build-essential libpq-dev postgresql-client

RUN mkdir /app
WORKDIR /app

COPY Gemfile /app/Gemfile
# COPY Gemfile.lock /app/Gemfile.lock

RUN bundle install

COPY . /app

CMD ["rails", "server", "-b", "0.0.0.0"]
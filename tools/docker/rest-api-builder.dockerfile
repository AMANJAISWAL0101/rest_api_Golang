FROM golang:latest as Base

RUN apt-get update && apt install -y vim 



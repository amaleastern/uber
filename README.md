# Uber

## Pre-Requisite 
- Docker

## Configuration 
- All of the application specific configuration variable should be set inside .env file of the build or project 
- You have to copy .env.example file to .env file and populate .env file variables as per deployment. 

## Running an application in development env 
`make run`

## Build and Run docker based application
`make build PORT=9090`
>You can change port number according to you deployment need 

> If you change port number here then you have to change port number in .env file as well
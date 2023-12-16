package main

import "github.com/joho/godotenv"


func SetupConfig(){
  godotenv.Load(".env")
}

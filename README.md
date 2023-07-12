# ShortURL

Web page to create shortURL and launch shortURL created. 

# Requirements
- Go 
- Sqlite 3 

# Installation

## Env configuration 
Copy the .env.sample to .env file and edit the value PORT to change the default port to execute the aplication. 
The default port is 1234

## Install the database
Create the install file 
```sh
go build cmd\install.go
install.exe
```

## Install de aplication 
Create the application 
```sh
go build
shorturl.exe
```

# Run the application

Run in the browser http://localhost:1234

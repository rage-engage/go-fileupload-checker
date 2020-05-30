# Description
A simple server used to check PUT file uploads. Currently just keeps the uploaded file in memory and does not write to disc.   
Default port is `8080`  

## Usuage

### Init
If you don't want to use the default port you can pass to the CLI `-p=9090`

### Upload
Currently only `PUT` requests to `/upload` is configured. param it's looking for is `fileName`.  
![](./Screenshot%202020-05-30%20at%2015.35.20.png?raw=true)  
You should see more information in CLI

![](./Screenshot%202020-05-30%20at%2015.37.27.png?raw=true)
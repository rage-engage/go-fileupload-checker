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

### Easy install via [GoBinaries](https://gobinaries.com/)  
Using [GoBinaries](https://gobinaries.com/)  
You can install and use this easily  
`curl -sf https://gobinaries.com/rage-engage/go-fileupload-checker | sh`  
Once installed run using  
`go-fileupload-checker`

### Uninstall
I have not tested this on other OSes but for **MacOS** it's installed to `/usr/local/bin` , so to uninstall you should run  
`rm /usr/local/bin/go-fileupload-checker`  
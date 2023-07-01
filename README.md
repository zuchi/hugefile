# Huge file project

This project is to test memory limit when uploading a huge json file in Golang

### How you can run this project
<p>To run this project you must have the docker environment (docker and docker-compose) installed in our machine. Once you get this installed, what you have to do is to clone this repository locally
and then you can execute the follow command.</p>

<p>`docker-compose up -d`. if you get an error, please verify if you are in the same folder of docker-compose.yml file.</p> 

<p>This command will download the respectives images (mongodb) and it will compile the project in order to generate the docker image.</p>

#### Environment variables
There are 3 environment variables inside the docker compose you may want to change. The variables are:
 * `SERVER_ADDRESS`: this variable is responsible for set the address that the api will start the application. Default value is `:3000`
 * `DATABASE_NAME`: this is the name that defines the new database that will be created into mongodb. Default value is `port_db`
 * `DATABASE_URI`: this is the address that mongodb is running the format. Default value is: `mongodb://mongo-database:27017`. Please note that mongo-database should match with the container_name inside the mongo session in the docker-compose.yml file.

#### Makefile
This project also contains a makefile. The makefile here is to facilitate the execution of some features such as testing the project. The list of commands you can use is described following:
* `make test`: This command will execute all tests that we have in the source code. PS: To run the tests, you must have golang installed in our machine.
* `make start`: This command will execute the `docker-compose up -d` to start the containers
* `make stop`: This command will stop the execution of the `docker-compose`
* `make build-containers`: This will execute the composer and will build the images again. this is useful when you make some change in the code and would like to generate the images again. 

### Sending one file to the application:
To send a file to the application you can use any http client that is able to upload one file to the server.
the name of the parameter you have to send the file is `file`. You have to use the method `POST` and the URI is /upload.

Here is example using cURL you can test

```text
curl --location 'http://localhost:3000/upload' \
--form 'file=@"PATH_TO_FILE"'
```
PATH to file should be replaced by the folder the file lives and the name of the file. Let's imagine that the file is 
in this path '/user/jederson/' and the file name is `upload-me.json`. In this case, you would need to replace the curl command
to be like: 
```text
curl --location 'http://localhost:3000/upload' \
--form 'file=@"/user/jederson/upload-me.json"'
```

### Getting one record from our database:
To get one item from the database you can use the a GET method under the HTTP protocol like the following example:
```text
curl --location 'http://localhost:3000/port?id=AEAJM'
```

This command will bring you (if any) one json struct of our database. If you want to change the id, you can just replace AEAJM for something else
# Note Service

## Brief description

<p align="justify">
	
This is a service dedicated to keep brief memos with a structure "Title, Content". The title is constrained to be shorter than 20 letters, 
whereas the content is bounded to 1000 letters. The service' API accepts gRPC or HTTP requests and converts the received Protobuffer 
request into a simple golang struct, isolated from the outer layer. That struct is then passed to specific method of a service layer
according to the initial request. The service layer in turn redirects the received model to specific method in repository layer, which has an 
interface for communication with PostgreSQL database. This service requires at least [Docker](https://www.docker.com/) and [Goose](https://github.com/pressly/goose/) installed as well as using Linux or
WSL to set up the Note Service app and database in a container.
	
</p>

## Implemented technologies


Protobuffer  |  gRPC | Docker
:-------------------------:|:-------------------------:|:-----------------:
   <img  src="./readme_assets/pluginIcon.png" width="80%"> | <img  src="./readme_assets/grpc-icon-color.png" width="50%"> |  <img  src="./readme_assets/vertical-logo-monochromatic.png" width="80%">          


## Project setup

### Out of the box scenario

<p align="justify">
	
In case you want to just use this service out of the box, you need to verify the installation of goose and docker. If you don't have goose installed,
```
curl -fsSL \
    https://raw.githubusercontent.com/pressly/goose/master/install.sh |\
    GOOSE_INSTALL=/usr/local/bin/.goose sh -s v3.5.0
sudo cp -r /home/$USER/.goose/bin/goose /usr/local/bin
```
Then you need to pull Docker images from my repository on DockerHub.
```
docker pull nikitads9/note-service:app
docker pull postgres:14-alpine3.15
```
When it is done, it's time to run containers using pulled images. If you want to specify your own database connection parameters, you should change the environment `-e` and port `-p` flags in the command featured below:
```
docker network create note-service-network
docker run -d -e POSTGRES_DB='notes_db' \
 -e POSTGRES_PASSWORD='notes_pass'\
 -e POSTGRES_USER='postgres'\
 -e PGDATA='/var/lib/postgresql/data/notification'\
 -p 5432:5432\
 -v postgres-volume:'/var/lib/postgresql/data'\
 --network note-service-network \
 --name postgres\
 postgres:14-alpine3.15
docker run -d --name app\
 -p 50051:50051\
 -p 8000:8000\
 -v 'app-volume:/var/lib/note-app/data'\
 --network note-service-network\
 nikitads9/note-service:app
```
**NB**: If you have changed the database configuration in `docker run` command, you should also edit the connection variables in **migration-local.sh** script file. 
And finally, when both containers are up, run this bash script for migration:
```
bash migration-local.sh
```
Now the database table is created and you can send HTTP and gRPC requests to the server app.
</p>

### Advanced installation

<p align="justify">
	
In case you want to build the service yourself, you will need to have these tools installed:
- makefile
- goose
- protobuffer-compiler (protoc)
- docker
- golang
	
If you are ok with that, be sure to edit database connection parameters in **config.yml** file among with **Dockerfile** and **migration-local.sh**. The commands to launch the server app and database are listed below:
```
git clone https://github.com/nikitads9/note-service-api.git
cd note-service-api/
make deps
make vendor-proto
make generate
docker-compose up -d
curl -fsSL \
    https://raw.githubusercontent.com/pressly/goose/master/install.sh |\
    GOOSE_INSTALL=$HOME/.goose sh -s v3.5.0
sudo cp -r /home/$USER/.goose/bin/goose /usr/local/bin
bash migration-local.sh
```
- The `make deps` command installs dependencies required for this project.
- The `make vendor-proto` command downloads the required tools for protobuf and validate to work. Running this command will create proto folder in the root of the project with all necessary `.proto` files.
- The `make generate` command creates three files: `grpc.pb.go`, `pb.go`, `pb.gw.go` based on API description in **note_v1.proto**. These files contain golang structs, interfaces and golang methods generated on the basis of Protobuffer interface description.
- The `docker-compose up -d` command downloads **alpine3.15** image from Docker Hub (if you don't have it locally), builds a binary and creates two containers: one for server app which is the the API service itself and the second one acts as database server. Both containers are connected to default Docker network which enables the two containers to communicate successfully. 
- The `curl -fsSL...` command downloads goose tool for database migration and initiates the installation process.
- The `sudo cp -r /home/$USER/.goose/bin/goose /usr/local/bin` command copies goose binary file to `usr/local/bin` folder so that your Linux could run goose commands from anywhere.
- The `bash migration-local.sh` command starts the bash script, that completes database migration specified in `.sql` file in **/migrations** folder. The parameters required for database connection to complete migration are specified in **migration-local.sh**.

</justify>

## API use instruction

This service is an API that implements the CRUD concept. It features the ability to create, read, update and delete database entries. The instruction below is for simple HTTP+JSON requests with RESTful API style. If you rather want to write a gRPC client, you would need to look up in the **note_v1.proto** file.
<details>
<summary> 
1. AddNote handle 
</summary>
  
**POST** `host:port/note/v1/add` <br />
The JSON object passed to that handle should look like:
```
{
	"note": {
		"title": "YourTitle",
		"content": "YourContent"
	}
}
```
The handle returns JSON with nested id of added memo.
```
{
	"result": {
		"id": "1"
	}
}
```
</details>
<details>
<summary> 
2. RemoveNote handle 
</summary>
  
**DELETE** `host:port/note/v1/remove/{id}` <br />
This handle does not need JSON. It requires a note id in the request instead.
It returns JSON with id of memo removed.
```
{
	"removed": "1"
}
```
</details>
<details>
<summary> 
3. MultiAdd handle 
</summary>
  
**POST** `host:port/note/v1/multi-add` <br />
The JSON object passed to that handle should look like:
```
{
	"notes": [{ 
		"title": "YourTitle1",
		"content": "YourContent1"
}, 
{
			"title": "YourTitle2",
		"content": "YourContent2"
}]
}
```
The handle returns JSON with nested quantity of memos added.
```
{
	"result": {
		"count": "2"
	}
}
```
</details>
<details>
<summary> 
4. GetNote handle 
</summary>
  
**GET** `host:port/note/v1/get/{id}` <br />
This handle does not need JSON. It requires a note id in the request instead.
The handle responds with requested memo id, title and content the latter two being nested in note struct.
```
{
	"id": "1",
	"note": {
		"title": "YourTitle",
		"content": "YourContent"
	}
}
```
</details>
<details>
<summary> 
5. GetList handle 
</summary>
  
**GET** `host:port/note/v1/get-all-notes` <br />
This handle does not require JSON or number. It's goal is to show all entries in a database.
In response the handle provides an array of JSON objects with all info about memos keeped.
```
{
	"results": [
		{
			"id": "1",
			"note": {
				"title": "YourTitle1",
				"content": "YourContent1"
			}
		},
		{
			"id": "2",
			"note": {
				"title": "YourTitle2",
				"content": "YourContent2"
			}
		}
	]
}
```
</details>
<details>
<summary> 
6. UpdateNote handle 
</summary>
  
**PUT** `host:port/note/v1/update` <br />
The JSON object passed to that handle should look like the one below, although one may omit updating either Title or Content (or both) so both fields can take null value.
```
{
	"id": 1,
	"title": "Updated title",
	"content": "updated content"
}
```
The handle response should look like an epmty JSON object: {}

</details>

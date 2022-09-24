# Note Service
## Brief description
<p align="justify">
This is a service dedicated to keep brief memo with structure "Title, Content". The title is constrained to be shorter than 20 letters, 
whereas the content is bounded to 1000 letters. The service' API accepts gRPC or HTTP requests and converts the received Protobuffer 
request into a simple golang struct, isolated from the outer layer. That struct is then passed to specific method of a service layer
according to initial request. The service layer in turn redirects the received model to specific method in repository layer, which has an 
interface for communication with PostgreSQL database.
</p>

## Implemented technologies


Protobuffer  |  gRPC | Docker
:-------------------------:|:-------------------------:|:-----------------:
   <img  src="./readme_assets/pluginIcon.png" width="80%"> | <img  src="https://grpc.io/img/logos/grpc-icon-color.png" width="50%"> |  <img  src="./readme_assets/vertical-logo-monochromatic.png" width="80%">          


## Project setup

```
git clone https://github.com/nikitads9/note-service-api.git
cd note-service-api/
git checkout task4
make generate
docker-compose up -d
bash migration-local.sh
```
<justify>
  
- The `make generate` command creates three files: `grpc.pb.go`, `pb.go`, `pb.gw.go` based on API description in **note_v1.proto**. These files contain golang structs, interfaces and golang methods generated on the basis of Protobuffer interface description.
- The `docker-compose up -d` command downloads of **alpine3.15** image from Docker Hub, builds a binary and creates two containers: one for server app which is the the API service itself and the second one acts as Database server. Both containers are connected to default Docker network which enables the two containers to communicate successfully. 
- The `bash migration-local.sh` command starts the bash script, that completes database migration specified in `.sql` file in **/migrations** folder. The parameters required for database connection to complete migration are specified in **migration-local.sh**.

</justify>

## API use instruction

This service is an API that implements the CRUD concept. It features the ability to create, read, update and delete database entries.
<details>
<summary> 
1. AddNote handle 
</summary>
  
**POST** `host:port/note/v1/add` <br />
The JSON object passed to that handle should look like:
```
{
	"title": "YourTitle",
	"content": "YourContent"
}
```
</details>
<details>
<summary> 
2. RemoveNote handle 
</summary>
  
**DELETE** `host:port/note/v1/remove/{id}` <br />
This handle does not need JSON. It requires a note id in the request instead.
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
</details>
<details>
<summary> 
4. GetNote handle 
</summary>
  
**GET** `host:port/note/v1/get/{id}` <br />
This handle does not need JSON. It requires a note id in the request instead.
</details>
<details>
<summary> 
5. GetList handle 
</summary>
  
**GET** `host:port/note/v1/get-all-notes` <br />
This handle does not require JSON or number. It's goal is to show all entries in a database.
</details>
<details>
<summary> 
6. UpdateNote handle 
</summary>
  
**PUT** `host:port/note/v1/update` <br />
The JSON object passed to that handle should look like:
```
{
	"id": 10,
	"title": "Updated title",
	"content": "updated content"
}
```
</details>

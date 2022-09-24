
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


Protobuffer             |  gRPC
:-------------------------:|:-------------------------:
   <img  src="./readme_assets/pluginIcon.png" width="80%"> | <img  src="https://grpc.io/img/logos/grpc-icon-color.png" width="50%">             


## Project setup

```
git clone https://github.com/nikitads9/note-service-api.git
cd note-service-api/
git checkout task4
make generate
docker-compose up -d
```

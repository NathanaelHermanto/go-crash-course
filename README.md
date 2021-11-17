# go-crash-course
a golang crash course:

- basic syntax
- pointer
- struct & interface
- web dev hello world & docker

### Docker
build the docker image
```
docker build -t go-dock
```
run the docker image
```
docker run -p 3000:3000 go-dock
```
Check on `http://localhost:3000` to see if the app is running

stop the docker container
```
docker stop boilerplate
```

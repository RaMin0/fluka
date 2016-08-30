# Fluka
Chatbots coming to life!

## Compiling

You can compile using `make` the resulting binary is saved to `_dst`

## Running with Docker

#### First time or after Dockerfile change

If you are building the docker image for the first time or updating the current image:

	$ docker build -t fluka .

#### Running from a docker image

	$ docker run -it --rm -p 3000:3000 --name fluka-instance fluka

The `-it` option stands for interactive (not in the background), `--rm` will delete the container
after it has finished executing (in our case after we terminate the app) and `-p` maps internal
container port 3000 to the localhost (or kitematic docker machine ip).

###### Opening the app in the browser

If you are on a \*nix machine you can go to [localhost:3000](http://localhost:3000) directly, otherwise use docker's
virtual machine ip.

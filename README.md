# go-beanstalkd-example

## Build the binaries

Build the consumer app

    $ cd Consumer && ./build.sh

Build the producer app

    $ cd Producer && ./build.sh

Build the docker images

    $ docker-compose build

Run the application

    $ docker-compose up -d
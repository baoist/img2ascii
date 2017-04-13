# img2ascii

A simple webserver that takes images (png, jpg)
and converts them into beautiful ASCII art.

### Requirements

- Install and set up [Go Version 1.8](https://golang.org/dl/).
- Port `8080` available.

### General Usage and Startup

- Retrieve project from github:
    ```
    go get github.com/baoist/img2ascii
    ```
- Start the web server locally
    ```
    go run $GOPATH/src/github.com/baoist/img2ascii/main.go
    ```
- Make a test request
    ```
    IMAGE_PATH=$GOPATH/src/github.com/baoist/img2ascii/example/images/200.lando.jpg
    curl -X POST -H "Content-Type: multipart/form-data" \
      -F "image=@$IMAGE_PATH" \
      http://localhost:8080/process
    ```
- To view a prettier version of this, install the [jq project](https://stedolan.github.io/jq/) and run:
    ```
    curl -X POST -H "Content-Type: multipart/form-data" \
      -F "image=@$IMAGE_PATH" \
      http://localhost:8080/process | jq -r '.image'
    ```

### Potential Issues

Uploaded images that are smaller than the max width/height settings in `image_processor/processor.go`
are scaled up to the max. Any pixels that didn't exist in the original image's matrix
have a value of the [1]th index of the ascii map.

### External Libraries

The only non-standard package used was [github.com/nfnt/resize](https://github.com/nfnt/resize).
This was chosen because it provided an easy way to efficiently resize
(using the Nearest-Neighbor algorithm).

# ASCII Art Text Generator in Golang 

## Description

This project is a local web server that receives a text input from the user and displays an ASCII Art representation of the text. The project is implemented in Golang and uses the net/http package to handle HTTP requests.

## Authors
This web app was designed by Ahmed Abdeen & Ahmed Alali as part of the Reboot01 program.

## Getting Started

To get started with the project, follow these steps:

1. Clone the repository to your local machine:

```go
    git clone https://learn.reboot01.com/git/aabdeen/ascii-art-web
```

2. Install Golang on your machine if you haven't already done so. You can download it from the official website [↗](https://go.dev/doc/install "Golang Install Page").

3. Navigate to the project directory:

```go
    cd ascii-art-web/webprogram
```

4. Start the web server:
```go
    go run main.go
```
The server will start listening for HTTP requests on port 8080. You can access the web interface by navigating to http://localhost:8080 in your browser.

## Usage

To generate ASCII Art from text, you can enter the requested text into the text field:

Textfield input:
>Hello

Choose one of the available banner types e.g.
>Standard

The server will respond with an ASCII Art representation of the text in the choosen banner type:

```
     _    _          _   _          
    | |  | |        | | | |         
    | |__| |   ___  | | | |   ___   
    |  __  |  / _ \ | | | |  / _ \  
    | |  | | |  __/ | | | | | (_) | 
    |_|  |_|  \___| |_| |_|  \___/  
                                
```

## Testing

The project comes with an HTTP methods test file that can be run using Postman. The file is located in the tests directory and can be imported into Postman.

To run the test file:

    Open Postman and import the test file by clicking Import > Upload Files and selecting the ascii-art-generator.postman_collection.json file.

    Run the collection by clicking the Run button.

    The test file will send a few HTTP requests to the server and verify the responses.

### License
This project is licensed under the 01 Schools License - see the LICENSE file for details.
## GO for web developers
- Web Assembly
    Usefull for mixing it with JavaScript and front-end
- Transpile to JS
    To write your frontend code directly in GO
- Web Server
    Serve fils and HTML including a template service
- Wer Services
    RESTful APIs or Microservices


## GO Templates
- HTML file with GO code
- it's in the html/package package
- The template can include expressions in {{}}
- Trimming spaces available
- Actions and pipelines
- if-else conditions
- range for loops
- You can call functions

## Basic GO template system
we create first a template and then execute it to replace the parameters

```go
import "html/template"

type Person struct {
    Name string
}

func main() {
    t := template.New("my-template").Parse(`
        <!DOCTYPE html>
        <title> My website </title>
        <h1> My Website {{.name}}</h1>
    `)
}

person := Person{name:"John"}
output,err := t.ExecuteToString(person)
```

## Final Steps

### Compiling
- Compiling the project - `go build .`
- Compiling in one specific output folder `go build . -o build/`
- Compiling for other platforms and OSs
    ```shell
    env GOOS=target-OS
        GOARCH=target-architecture go build .
    ```
- Compile and install - `go install .`

## Packaging
- GO just produces a binary
- It doesnt provide any packaging solutions
- If we want to embed assests for an app, we need to use third-pary or OSs tools such as:
    - Creating installers for windows
    - Create a DMG package for macOS
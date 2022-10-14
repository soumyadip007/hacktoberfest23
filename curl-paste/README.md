# Golang powered cURL paste

The `HTTP` `POST` request will send the form data along and respond with a link to the paste. The `HTTP` `GET` will retrieve the paste with the given ID as plain-text.

### Usage
Paste a file named `file.txt` using cURL.
```bash
$ curl -F 'title=Request from cURL' -F 'paste=&lt;file.txt' http://curl-paste.local/
```
Paste from stdin using cURL.
```bash
echo "Hello, world." | curl -F 'title=Request from cURL' -F 'paste=&lt;-' http://curl-paste.local/
```
A shell function that can be added to `.bashrc` or `.bash_profle` for quick pasting from the CLI. The command reads from stdin and outputs the URL of the paste to stdout.
```bash
function curl-paste() {
	curl -F 'title=Request from cURL' -F 'paste=<-' https://curl-paste.local
}
```
```bash
echo "hi" | curl-paste
```

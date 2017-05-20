# Auth0 + Go Web App Sample

## Running the App

To run the app, make sure you have **go** and **goget** installed.

Rename the `.env.example` file to `.env` and provide your Auth0 credentials.

```bash
# .env

AUTH0_CLIENT_ID=2DfLXKE_nNXW0cb47CmtcWNl48Md52eO
AUTH0_DOMAIN=trinkio.auth0.com
AUTH0_CLIENT_SECRET=DWu3-lvh19-jqQVHd9CgBRVvBe82c0NrqVKKybV8DYEm046fNlGZIsgIcOZUAymj
AUTH0_CALLBACK_URL=http://localhost:3000/callback
```

Once you've set your Auth0 credentials in the `.env` file, run `go get .` to install the Go dependencies.

Run `go run main.go server.go` to start the app and navigate to [http://localhost:3000/](http://localhost:3000/)

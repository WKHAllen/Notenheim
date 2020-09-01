# PGTR Sample App

This is a sample app that uses a PostgreSQL-Go-TypeScript-React stack.

## Deployment

This sample app is [deployed here](https://pgtr.herokuapp.com/) on [Heroku](https://heroku.com/).

## Stack

For the frontend, [React](https://reactjs.org/) was the obvious choice, as it is one of the most powerful frontend frameworks out there. The [TypeScript](https://www.typescriptlang.org/) variant of React is used because it is very easy to make mistakes when using JavaScript.

Most modern languages could work for the backend. [Go](https://golang.org/) was chosen because of its simplicity and because of one of its most powerful web frameworks, [Gin](https://github.com/gin-gonic/gin). 

[PostgreSQL](https://www.postgresql.org/) is used for the database because of its popularity and advanced features. In addition, Heroku has a helpful addon for interacting with Postgres databases.

## Setup

This repository can serve as a template for projects designed using the PGTR stack. Simply follow the steps below, skipping the ones that are not relevant to your project.

### Clone

First, clone the sample project.

```console
$ git clone https://github.com/WKHAllen/pgtr-sample-app.git
```

Push the cloned project to your own project's repository in GitHub.

### Initializing Heroku

Create the app on [Heroku](https://heroku.com). Then configure the app to deploy using your project's GitHub repository.

### Buildpacks

Add the Go and Node.js buildpacks to the app. You will need both to build and run the project.

### Connecting the database

Add the Heroku Postgres addon to the app. This will automatically provision a PostgreSQL database and configure your app's DATABASE_URL config variable. The variable will be used by the backend to connect to the database.

### Environment variables

App config variables should be set up automatically where your project will be hosted. In order to run the app locally, however, you'll need to create a file called `.env` in the root directory of the project. Inside this file, you will need the following:

```
PORT=3000
DATABASE_URL=<databaseURL>
```

`<databaseURL>` above should be replaced with the URL that has been automatically configured in your app's remote config variables.

### Database initialization

If your database needs to be initialized programmatically, this can be done in [`src/dbinit.go`](src/dbinit.go) using the InitDB function.

**Note:** this function will run every time the app starts. Be careful how you write your SQL statements in this function. For example, use:

```sql
CREATE TABLE IF NOT EXISTS Person(
    ...
);
```

rather than:

```sql
CREATE TABLE Person (
    ...
);
```

since the table will already exist unless this is the first time starting the app.

## Extending the app

The app can be easily extended. Add frontend components or backend REST endpoints using the instructions below.

### Frontend

The frontend of the application is a series of React components written in TypeScript. Some components are set up to make calls to the REST API.

#### Adding components

To add a component, create a new TypeScript React file in the `app/src/components` directory. Corresponding stylesheets should go in `app/src/css`.

The app is configured to load [`app/src/components/errors/NotFound.tsx`](app/src/components/errors/NotFound.tsx) in the event of a 404. Editing this component is encouraged. Deleting it is not.

#### Using REST endpoints

Use the [Fetch API](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API/Using_Fetch) to make requests to REST endpoints. See [`app/src/components/People.tsx`](app/src/components/People.tsx) for an example.

#### Routing

Routing is done on the frontend, using [`react-router-dom`](https://www.npmjs.com/package/react-router-dom). See [`app/src/components/App.tsx`](app/src/components/App.tsx) or [this quickstart guide](https://reactrouter.com/web/guides/quick-start) for examples.

### Backend

The backend is divided into routes and services. The routes package manages the REST endpoints, which use the services to interact with the database.

#### Adding a REST endpoint

To add an endpoint, create a new Go file in the `src/routes` package or open an existing one. Add a function which follows the same format as the existing ones. If the function needs to access the database, you will need to create a new service.

Create or add to a file in the `src/services` package. Add a function which again follows the format as others in the package. Use the `dbm` object to interact with the database.

After creating an endpoint, you'll need to expose it to the router. Open [`src/routes/init.go`](src/routes/init.go) and append the following line to the `LoadRoutes` function body:

```go
api.GET("<URLPath>", <routeFunction>)
```

Replace `<URLPath>` above with the desired API endpoint path. For path format options (i.e. parameters in path), see [Gin's examples](https://github.com/gin-gonic/gin#api-examples).

Replace `<routeFunction>` above with the name of the function you added to the `src/routes` package.  

The `src/routes/helper` package will handle some common routing problems for you. Use it.

**Do not** delete `init.go` from the `src/routes` and `src/services` packages. They are important, and the REST API will not work without them.

## Running locally

Scripts have been provided for running the app locally. Run the appropriate script, then navigate to [`localhost:3000`](http://localhost:3000/) in your browser.

### Windows

On Windows, use `run.bat`:

```console
run
```

### MacOS and Linux

On MacOS or Linux, use `run.sh`:

```console
./run.sh
```

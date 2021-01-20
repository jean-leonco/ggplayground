package ggplayground

import (
	"log"
	"net/http"
)

const playgroundString = `
  <!DOCTYPE html>
  <html>

  <head>
    <meta charset=utf-8/>
    <meta name="viewport" content="user-scalable=no, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, minimal-ui">
    <title>GraphQL Playground</title>
    <link rel="stylesheet" href="//cdn.jsdelivr.net/npm/graphql-playground-react/build/static/css/index.css" />
    <link rel="shortcut icon" href="//cdn.jsdelivr.net/npm/graphql-playground-react/build/favicon.png" />
    <script src="//cdn.jsdelivr.net/npm/graphql-playground-react/build/static/js/middleware.js"></script>
  </head>

  <body>
    <div id="root">
      <style>
        body {
          background-color: rgb(23, 42, 58);
          font-family: Open Sans, sans-serif;
          height: 90vh;
        }

        #root {
          height: 100%;
          width: 100%;
          display: flex;
          align-items: center;
          justify-content: center;
        }

        .loading {
          font-size: 32px;
          font-weight: 200;
          color: rgba(255, 255, 255, .6);
          margin-left: 20px;
        }

        img {
          width: 78px;
          height: 78px;
        }

        .title {
          font-weight: 400;
        }
      </style>
      <img src='//cdn.jsdelivr.net/npm/graphql-playground-react/build/logo.png' alt=''>
      <div class="loading"> Loading
        <span class="title">GraphQL Playground</span>
      </div>
    </div>
    <script>window.addEventListener('load', function (event) {
        GraphQLPlayground.init(document.getElementById('root'), {
					endpoint: '/graphql',
          setTitle: true,
        })
      })</script>
  </body>

  </html>
`

// playgroundHandler returns the GraphQL Playground. If the method is not allowed
// or if there was a error while writing the response, returns a HTTP error.
func playgroundHandler(w http.ResponseWriter, r *http.Request) {
	// If method is not allowed, returns a 450 error
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	playgroundHTML := []byte(playgroundString)
	_, writeErr := w.Write(playgroundHTML)

	// If it was not able to write response, returns a 500 error
	if writeErr != nil {
		http.Error(w, "Unable to load GraphQL playground", http.StatusInternalServerError)
		log.Printf("failed to return playground, error: %v", writeErr)
		return
	}
}

// Middleware add playgroundHandler to the ServeMux.
func Middleware(mux *http.ServeMux) {
	mux.HandleFunc("/playground", playgroundHandler)
}

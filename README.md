GoShort: A URL shortener written in Go
-
Tools used: Go on App Engine (https://developers.google.com/appengine/docs/go/gettingstarted/introduction) and Mux router (http://www.gorillatoolkit.org/pkg/mux)

To run this app locally, download the App Engine SDK from here (https://developers.google.com/appengine/docs/go/gettingstarted/devenvironment). Then run with the goapp serve command.

Note: This is a App Engine application. You need to run it with the App Engine SDK. In the future, I may add a version that only uses the standard Go http package, but in reality, you will probably be deploying a Go app on App Engine or something similar, so best to build for it from the start.

Todo:
-
Clean up and optimize code
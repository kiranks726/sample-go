a="Using stage: dev
Preparing your SST app
Building function cmd/handlers/movies/list/list.go
Building function cmd/handlers/movies/create/create.go
Building function cmd/handlers/movies/get/get.go
Building function cmd/handlers/movies/update/update.go
Building function cmd/handlers/movies/delete/delete.go
Stack dev-ctx-kitchensink-go-movie-stack
There were no differences"

if [[ $a == *"no differences"* ]]; then
  echo "It's there!"
else
    echo "Not there"
fi
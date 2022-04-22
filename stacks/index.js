
import MovieStack from "./MovieStack";
export default function main(app) {

  // Set default runtime for all functions
  app.setDefaultFunctionProps({
    runtime: "go1.x"
  });

  new MovieStack(app, "movie-stack");

}

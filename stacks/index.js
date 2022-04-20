
import OtherStack from "./OtherStack";
import TodoStack from "./TodoStack";
import MovieStack from "./MovieStack";
export default function main(app) {

  // Set default runtime for all functions
  app.setDefaultFunctionProps({
    runtime: "go1.x"
  });

  new OtherStack(app, "other-stack");
  new TodoStack(app, "todo-stack");
  new MovieStack(app, "movie-stack");

}

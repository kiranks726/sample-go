
import MovieStack from "./MovieStack";
import Configstack from "./ConfigStack";
import * as config from "config";

export default function main(app) {

  // Set default runtime for all functions
  app.setDefaultFunctionProps({
    runtime: "go1.x"
  });

  const project = `${app.stage}${config.get("projectPrefix")}`;
  const projectVersion = config.get("projectVersion");

  new MovieStack(app, "movie-stack", {
    tags: {
      "ctx:project": project,
      "ctx:project-version": projectVersion
    }
  });

  // if (app.stage != "local") {
  //   new Configstack(app, "config-stack", {
  //     tags: {
  //       "ctx:project": project,
  //       "ctx:project-version": projectVersion
  //     }
  //   });
  // }
}

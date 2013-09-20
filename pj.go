package main

import (
    //"syscall"
    "os"
    "flag"
    "fmt"
    "strings"
)

const (
  ENV_ID = "PJ_DIRS"
)

/*
  Exports func addProjectDir {
    export a $SHELL variable project-bootstrap_PATH=/path1/:/path2/:etc/
  }
*/
/* 
  Exports func removeProjectDir
    remove path from $SHELL variable if it exists
*/
/* 
  Exports func goToProjectDir {

    Chdir to location in one of your project paths
    Search dir for project-bootstrap.json
    Parse JSON for array of commands to run in the directory
    Open files in the JSON array of files

  }
*/

func main() {

    //fmt.Println("---> PJ GO!")
    // Allow folders with . in front of them?
    allowHiddenFolders := flag.Bool("allow-hidden", false, "Allow folders with . in front to be parsed as projects by pj.")
    targetDirectory := flag.String("project", "foo", "Target project for pj.")

    flag.Parse();

    // Paths should be array
    var pathToProjects string;

    // Get all environment vars
    env := os.Environ()

    // Parse through them checking for our dependant one PB_DIRS
    // when found, get string from opposite side of the equal sign
    // split that by ":"
    for i := 0; i < len(env); i++ {
      currentEnvVar := strings.Split(env[i], "=");

      if currentEnvVar[0] == ENV_ID {
        pathToProjects = currentEnvVar[1]
      }
    }

    // At this point if you dont have a path, bounce.
    if len(pathToProjects) < 1 {
      fmt.Println("-- WARN --> You have not established your project directory.(export PJ_DIRS=/path/to/projects")
      return;
    }

    // Get environment variables and parse out the SHELL var
    // For each string seperated by : push to paths array
    //paths := strings.Split(pathFromShell, ":")

    // TODO: for each instead of allowing 1 dir
    // For each path in paths array
    dir, err := os.Open(pathToProjects)
    if err != nil {
      panic(err)
    }
    defer dir.Close()

    fileInfos, err := dir.Readdirnames(-1)
    if err != nil {
      panic(err)
    }

    for _, name := range fileInfos {
        // For every dir in path -- does it match option passed into to pj? If so
        // goToProjectDir
        if !*allowHiddenFolders && strings.IndexAny(name, ".") == 0 {
          continue;
        }

        if name == *targetDirectory {
          //fmt.Println("pj going to directory: ", *targetDirectory)
          fmt.Fprint(os.Stdout, pathToProjects + "/" + *targetDirectory)
        }
    }

}

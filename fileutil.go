package gocommons

import "log"
import "os"
import "os/user"
import "path/filepath"


// Produces a path that combines the user's home directory with the
// passed filename. If an error occurs, it's logged and an empty
// string is returned 
func PrependHomeDir(filename string) string {
    user,err := user.Current()
    if err != nil {
        log.Printf("prependHomeDir: unable to access current user; returning empty path")
        return ""
    }
    return filepath.Join(user.HomeDir, filename)
}


// Determines whether the specified file exists
func FileExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil {
        return true, nil
    }

    if os.IsNotExist(err) {
        return false, nil
    }

    return false, err
}

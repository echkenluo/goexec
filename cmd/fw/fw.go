package main

import (
    "os"
    "time"
    // "strconv"
    "fmt"
    "log"
    "github.com/fsnotify/fsnotify"
    "k8s.io/apimachinery/pkg/util/wait"
)

func main() {
    stopchan := make(chan struct{})

    go watchFile(VswitchdUnixSock)
    go watchFile(OvsdbUnixSock)

    stopchan <- struct{}{}
}

var (
    OvsdbUnixSock = "/var/run/openvswitch/db.sock"
    VswitchdUnixSock = "/var/run/openvswitch/ovsbr-mgt.mgmt"
)

func watchFile(filePath string) {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        fmt.Printf("")
    }
    defer watcher.Close()

    if err := addWatchFile(watcher, filePath); err != nil {
        log.Fatalf("Failed to add file to watcher, error: %v", err)
    }

    done := make(chan bool)

    go func() {
        for {
            select {
            case event, ok := <-watcher.Events:
                if !ok {
                    return
                }
                if event.Op&fsnotify.Create == fsnotify.Create {
                    fmt.Println("create event")
                }
                if event.Op&fsnotify.Write == fsnotify.Write {
                    fmt.Println("write event")
                }
                if event.Op&fsnotify.Remove == fsnotify.Remove {
                    fmt.Println("remove event")
                    wait.PollImmediate(200*time.Millisecond, 10 * time.Second, func() (do bool, err error){
                        if _, err := os.Stat(VswitchdUnixSock); os.IsNotExist(err) {
                            log.Printf("vswitchd unix sock not found")
                            return false, nil
                        }

                        if err := addWatchFile(watcher, VswitchdUnixSock); err != nil {
                            log.Printf("failed to watch vswitchd unix sock, error: %v", err)
                            return false, nil
                        }

                        log.Printf("watch vswitchd unix sock created")
                        return true, nil
                    })
                }
                if event.Op&fsnotify.Rename == fsnotify.Rename {
                    fmt.Println("rename event")
                }
                if event.Op&fsnotify.Chmod == fsnotify.Chmod {
                    fmt.Println("chmod event")
                }
            case err := <-watcher.Errors:
                fmt.Println("error", err)
            }
        }
    }()

    <-done
}

func addWatchFile(watcher *fsnotify.Watcher, filepath string) error {
    if err := watcher.Add(filepath); err != nil {
        return err
    }

    return nil
}

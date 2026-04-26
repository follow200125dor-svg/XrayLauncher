package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "os/exec"
    "runtime"
)

const xrayURL = "https://github.com/XTLS/Xray-core/releases/download/v1.8.23/Xray-windows-64.zip"

func main() {
    fmt.Println("=== Xray Launcher ===")

    // 1. Download Xray if not exists
    if _, err := os.Stat("xray.exe"); os.IsNotExist(err) {
        fmt.Println("Downloading Xray...")
        if err := downloadFile("xray.zip", xrayURL); err != nil {
            fmt.Println("Error:", err)
            return
        }
        fmt.Println("Unzipping...")
        exec.Command("tar", "-xf", "xray.zip").Run()
        os.Remove("xray.zip")
    }

    // 2. Run Xray with config
    fmt.Println("Starting Xray...")
    cmd := exec.Command("xray.exe", "run", "-c", "config.json")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()
}

func downloadFile(filepath string, url string) error {
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    out, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer out.Close()

    _, err = io.Copy(out, resp.Body)
    return err
}

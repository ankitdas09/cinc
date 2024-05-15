package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	out := compileAndRun()
	w.Write(out)
}

func compileAndRun() []byte {

	tmpFile, err := os.CreateTemp("", "test-*.c")
	if err != nil {
		log.Fatal(err)
	}

	code := `
	#include <stdio.h>
	int main()
	{
		printf("Hello World\n");
		printf("This is compiled and excecuted in a container inside a container!\n");
		printf("Bye World\n");
		return 0;
	}`

	b_code := []byte(code)

	err = os.WriteFile(tmpFile.Name(), b_code, 0666)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	absPath, err := filepath.Abs(tmpFile.Name())
	if err != nil {
		log.Fatal(err)
	}

	mountPath := fmt.Sprintf("%s:/tmp/main.c", absPath)
	execCmd := "gcc -o run /tmp/main.c && ./run"

	cmd := exec.CommandContext(ctx, "docker", "run", "--rm", "-v", mountPath, "-m", "50m", "--cpus", "0.2", "-i", "runner", "sh", "-c", execCmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	return out
}

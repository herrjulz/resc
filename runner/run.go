package runner

import (
	"os"
	"os/exec"
)

type RunFunc func(cmd *exec.Cmd) ([]byte, error)
type RunnerTypeFunc func(r *Runner)
type Runner struct{ run RunFunc }

func New(option RunnerTypeFunc) *Runner {
	var r Runner
	option(&r)
	return &r
}

func Default(r *Runner) { r.run = run }
func Test(r *Runner)    { r.run = testRun }

func (r *Runner) Run(script string, args ...string) ([]byte, error) {
	arguments := []string{"-c", script, os.Args[0]}
	arguments = append(arguments, args...)
	cmd := exec.Command("bash", arguments...)
	return r.run(cmd)
}

func run(cmd *exec.Cmd) ([]byte, error) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return nil, cmd.Run()
}

func testRun(cmd *exec.Cmd) ([]byte, error) {
	return cmd.CombinedOutput()
}

// 2018-02-15 Dave Trudgian

// Minimal pull and run of things in a container

// Simple tests to pull containers, unpack, and run
// This is mostly Liz Rice's containers-from-scratch code with
// some extra use of containers/image and opencontainers/image-tools

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
	"github.com/containers/image/copy"
	"github.com/containers/image/signature"
	"github.com/containers/image/transports/alltransports"
	"github.com/opencontainers/image-tools/image"
)

// go run main.go run <cmd> <args>
func main() {
	switch os.Args[1] {
	case "run":
		imageDir := fetchImage()
		unpackDir := unpackImage(imageDir)
		run(unpackDir)
	case "child":
		child()
	default:
		panic("help")
	}
}

func run(unpackDir string) {

	fmt.Printf("Running %v \n", os.Args[3:])

	args := append([]string{"child"}, unpackDir)
	args = append(args, os.Args[3:]...)
	cmd := exec.Command("/proc/self/exe",  args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	must(cmd.Run())
}

func child() {
	fmt.Printf("Running %v \n", os.Args[3:])
	fmt.Printf("In container at %s \n", os.Args[2])

	cg()

	cmd := exec.Command(os.Args[3], os.Args[4:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(syscall.Sethostname([]byte("container")))
	must(syscall.Chroot(os.Args[2]))
	must(os.Chdir("/"))
	must(syscall.Mount("proc", "proc", "proc", 0, ""))
	must(cmd.Run())

	must(syscall.Unmount("proc", 0))
}

func cg() {
	cgroups := "/sys/fs/cgroup/"
	pids := filepath.Join(cgroups, "pids")
	os.Mkdir(filepath.Join(pids, "dave"), 0755)
	must(ioutil.WriteFile(filepath.Join(pids, "dave/pids.max"), []byte("20"), 0700))
	// Removes the new cgroup in place after the container exits
	must(ioutil.WriteFile(filepath.Join(pids, "dave/notify_on_release"), []byte("1"), 0700))
	must(ioutil.WriteFile(filepath.Join(pids, "dave/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}

func fetchImage() string {

	policy := &signature.Policy{Default: []signature.PolicyRequirement{signature.NewPRInsecureAcceptAnything()}}
	policyContext, err := signature.NewPolicyContext(policy)
	must(err)

	fmt.Printf("Fetching container %s \n", os.Args[2])

	srcRef, err := alltransports.ParseImageName(os.Args[2])
	must(err)

	containerDir, err := ioutil.TempDir("","dockertest")

	fmt.Printf("Files will be fetched to %s \n", containerDir)

	destRef, err := alltransports.ParseImageName("oci:" + containerDir + ":singularity-build")
	must(err)

	must(copy.Image(policyContext, destRef, srcRef, &copy.Options{
		ReportWriter: os.Stdout,
	} ))

	return containerDir
}

func unpackImage(containerDir string) string{

	refs := []string{"name=singularity-build"}
	unpackDir, err := ioutil.TempDir("","dockertest")
	must(err)
	fmt.Printf("Container will be unpacked to %s \n", unpackDir)
	must(image.UnpackLayout(containerDir, unpackDir, "amd64", refs ))

	return unpackDir
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

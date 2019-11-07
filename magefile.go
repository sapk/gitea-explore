//+build mage

package main

import (
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/mitchellh/go-homedir"
)

var gocmd = "go"

func init() {
	gocmd = mg.GoCmd()
}

//TODO enfore GO111MODULE=on ?

// Generate go generate.
func Generate() error {
	return sh.RunV(gocmd, "generate", "./...")
}

// Build the binary.
func Build() error {
	/* Not needed since code is "vendored"
	if err := Generate(); err != nil {
		return err
	}
	*/
	return sh.RunV(gocmd, "build", "-mod=vendor", "-ldflags", "-s -w", ".")
}

// Run run with dev tags.
func Run() error {
	/* Not needed since code is "vendored"
	if err := Generate(); err != nil {
		return err
	}
	*/
	return sh.RunV(gocmd, "run", "-mod=vendor", "--tags", "dev", ".")
}

type Deps mg.Namespace

// Vendor store deps in vendor/ folder.
func (Deps) Vendor() error {
	return sh.RunV(gocmd, "mod", "vendor")
}

// Get get deps to mod store.
func (Deps) Get() error {
	return sh.RunV(gocmd, "get", "-v", "./...")
}

// Update update deps to last version.
func (Deps) Update() error {
	if err := sh.RunV(gocmd, "get", "-u", "-v", "./..."); err != nil {
		return err
	}
	return sh.RunV(gocmd, "mod", "tidy")
}

type Swagger mg.Namespace

// Update retrieve latest version of Gitea swagger definition.
func (Swagger) Update() error {
	return sh.RunV("wget", "-q", "https://try.gitea.io/swagger.v1.json", "-O", "gitea.swagger.json")
}

type Dev mg.Namespace

// Es start an test elasticsearch container + dejavu
func (Dev) Es() error {
	//TODO var goInstall = sh.RunCmd("go", "install") goInstall("github.com/gohugo/hugo")
	err := sh.RunV("docker", "run", "-d", "--rm", "--name", "elasticsearch", "-p", "9200:9200", "-p", "9300:9300", "-e", "discovery.type=single-node", "-e", "http.cors.enabled=true", "-e", "http.cors.allow-origin=http://localhost:1358,http://127.0.0.1:1358", "-e", "http.cors.allow-headers=X-Requested-With,X-Auth-Token,Content-Type,Content-Length,Authorization", "-e", "http.cors.allow-credentials=true", "docker.elastic.co/elasticsearch/elasticsearch-oss:7.3.1")
	if err != nil {
		return err
	}
	return sh.RunV("docker", "run", "-d", "--rm", "--name", "dejavu", "-p", "1358:1358", "appbaseio/dejavu")
}

// Docker build and start an docker instans
func (Dev) Docker() error {
	var d Docker
	err := d.Build()
	if err != nil {
		return err
	}
	return sh.RunV("docker", "run", "-ti", "--rm", "-p", "3000:3000", "--link", "elasticsearch:elasticsearch", "-e", "ES_URLS=http://elasticsearch:9200", "sapk/explore")
}

type Docker mg.Namespace

// SetupBuildx setup buildx needed build docker cross build
func (Docker) SetupBuildx() error {
	env := map[string]string{
		"DOCKER_BUILDKIT": "1",
	}
	err := sh.RunWith(env, "docker", "build", "--platform", "local", "-o", ".", "git://github.com/docker/buildx")
	if err != nil {
		return err
	}
	h, err := homedir.Dir()
	if err != nil {
		return err
	}
	err = os.MkdirAll(h+"/.docker/cli-plugins", os.ModePerm)
	if err != nil {
		return err
	}
	return sh.RunV("mv", "buildx", h+"/.docker/cli-plugins/docker-buildx")
}

// Build build dockerfile
func (Docker) Build() error {
	return sh.RunV("docker", "build", "-t", "sapk/explore", ".")
}

// Buildx buildx dockerfile
func (Docker) Buildx() error {
	return sh.RunV("docker", "buildx", "build", "-t", "sapk/explore", ".")
}

type Webapp mg.Namespace

// Build compile the web app
func (Webapp) Build() error {
	//TODO var goInstall = sh.RunCmd("go", "install") goInstall("github.com/gohugo/hugo")
	return sh.RunV("yarn", "--cwd", "./assets/webapp", "build", "--emoji")
}

// Serve compile the web app
func (Webapp) Serve() error {
	return sh.RunV("yarn", "--cwd", "./assets/webapp", "serve", "--emoji")
}

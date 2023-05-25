package cmd

import (
	"flag"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func init() {
	cmd := &cobra.Command{
		Use:   "deployer",
		Short: "",
		Run:   runDeployer,
		// Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	}
	// cmd.Flags().String("param", "value", "input params")
	appendCommand(cmd)
}

var (
	host     = flag.String("host", "", "SSH Host")
	port     = flag.Int("port", 22, "SSH Port")
	user     = flag.String("user", "", "SSH User")
	password = flag.String("password", "", "SSH Password")
	file     = flag.String("file", "", "File to deploy")
	rollback = flag.Bool("rollback", false, "Rollback to previous version")
	maxVer   = flag.Int("maxver", 5, "Maximum number of versions to keep")
)

func runDeployer(cmd *cobra.Command, args []string) {
	// param, _ := cmd.Flags().GetString("param")
	flag.Parse()

	if *host == "" || *user == "" || *password == "" {
		flag.Usage()
		os.Exit(1)
	}

	config := &ssh.ClientConfig{
		User: *user,
		Auth: []ssh.AuthMethod{ssh.Password(*password)},
	}

	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port), config)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	defer session.Close()

	if *rollback {
		rollbackVersion(session)
	} else {
		deployVersion(session)
	}
}

func deployVersion(session *ssh.Session) {
	// Upload binary file
	err := uploadFile(session, *file)
	if err != nil {
		log.Fatalf("Failed to upload file: %v", err)
	}

	// Create symbolic link to latest version
	latestLink := filepath.Base(*file)
	latestLink = strings.TrimSuffix(latestLink, filepath.Ext(latestLink))
	latestLink = fmt.Sprintf("%s-latest", latestLink)

	err = createSymbolicLink(session, *file, latestLink)
	if err != nil {
		log.Fatalf("Failed to create symbolic link: %v", err)
	}

	// Restart web service
	err = restartService(session)
	if err != nil {
		log.Fatalf("Failed to restart service: %v", err)
	}

	// Delete old versions if necessary
	err = deleteOldVersions(session, *maxVer)
	if err != nil {
		log.Fatalf("Failed to delete old versions: %v", err)
	}
}

func rollbackVersion(session *ssh.Session) {
	// Get previous version
	prevLink, err := getPreviousLink(session)
	if err != nil {
		log.Fatalf("Failed to get previous version: %v", err)
	}

	// Change symbolic link to previous version
	err = updateSymbolicLink(session, prevLink)
	if err != nil {
		log.Fatalf("Failed to update symbolic link: %v", err)
	}

	// Restart web service
	err = restartService(session)
	if err != nil {
		log.Fatalf("Failed to restart service: %v", err)
	}
}

func uploadFile(session *ssh.Session, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	go func() {
		w, _ := session.StdinPipe()
		defer w.Close()

		fmt.Fprintln(w, "C"+strconv.FormatInt(stat.Size(), 10), "0644", filepath.Base(filename))
		io.Copy(w, file)
		fmt.Fprint(w, "\x00")
	}()

	if err := session.Run("/usr/bin/scp -t " + filepath.Dir(filename)); err != nil {
		return err
	}

	return nil
}

func createSymbolicLink(session *ssh.Session, filename string, linkname string) error {
	cmd := fmt.Sprintf("ln -sf %s %s", filename, linkname)
	err := session.Run(cmd)
	if err != nil {
		return err
	}

	return nil
}

func updateSymbolicLink(session *ssh.Session, linkname string) error {
	cmd := fmt.Sprintf("ln -sf %s %s", linkname, strings.TrimSuffix(linkname, "-latest"))
	err := session.Run(cmd)
	if err != nil {
		return err
	}

	return nil
}

func restartService(session *ssh.Session) error {
	_, err := session.CombinedOutput("systemctl restart web.service")
	if err != nil {
		return err
	}

	return nil
}

func deleteOldVersions(session *ssh.Session, maxVer int) (err error) {

	return
}

func getPreviousLink(session *ssh.Session) (prevLink string, err error) {
	return
}

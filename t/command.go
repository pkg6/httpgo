package t

import (
	"bytes"
	"golang.org/x/crypto/ssh"
	"net"
	"os/exec"
	"time"
)

func CommandLocal(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}

func CommandSSH(session *ssh.Session, command string) (string, error) {
	var b bytes.Buffer
	session.Stdout = &b
	defer session.Close()
	err := session.Run(command)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func SSHLoginPassword(hostAddr string, username string, password string) (*ssh.Session, error) {
	var session *ssh.Session
	client, err := ssh.Dial("tcp", hostAddr, &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		Timeout: time.Second * 5,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	})
	if err != nil {
		return session, err
	}
	return client.NewSession()
}

func SSHLoginKey(hostAddr string, username string, privateKey []byte) (session *ssh.Session, err error) {
	signer, err := ssh.ParsePrivateKey(privateKey)
	if err != nil {
		return session, err
	}
	client, err := ssh.Dial("tcp", hostAddr, &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	})
	if err != nil {
		return session, err
	}
	return client.NewSession()
}

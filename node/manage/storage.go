package manage

import (
	"fmt"
	"github.com/yoneyan/vm_mgr/node/etc"
	pb "github.com/yoneyan/vm_mgr/proto/proto-go"
	"os/exec"
	"strconv"
	"strings"
)

type Storage struct {
	Path   string
	Name   string
	Format string
	Size   int
}

func RunStorageCmd(cmd []string) {
	out, _ := exec.Command("qemu-img", cmd...).Output()
	fmt.Println(string(out))
}

func GetMainStorage(data *pb.VMData) string {
	var basepath, path string
	sp := strings.Split(data.Option.GetStoragePath(), ",")
	if len(sp) < 1 {
		return ""
	}
	mode, _ := strconv.Atoi(sp[0])
	if mode/10 == 0 {
		path = sp[1] + "/" + data.GetVmname() + "-" + "0.img"
	} else {
		basepath = etc.GetDiskPath(mode / 10)
		if basepath == "" {
			fmt.Println("Config DiskPath Error")
			return ""
		}
		fmt.Println("basepath: " + basepath)
		return basepath + "/" + data.GetVmname() + "-" + "0.img"
	}
	return path
}

func StorageProcess(data *pb.VMData) string {
	s := strings.Split(data.GetStorage(), ",")
	sp := strings.Split(data.Option.GetStoragePath(), ",")
	j := 0
	var path string
	var result []string
	var mode int
	for i, a := range sp {
		if i%2 == 0 {
			result = append(result, a)
			mode, _ = strconv.Atoi(a)
		} else {
			if mode/10 == 0 {
				path = a
				fmt.Println("path: " + path)
			} else {
				path = etc.GetDiskPath(mode / 10)
				if path == "" {
					fmt.Println("Config DiskPath Error")
					return ""
				}
				fmt.Println("basepath: " + path)
			}
			path = path + "/" + data.GetVmname() + "-" + strconv.Itoa(j) + ".img"
			fmt.Println("path: " + path)
			if FileExistsCheck(path) == false {
				size, result := strconv.Atoi(s[j])
				if result != nil {
					fmt.Println("Error: string to int")
				}
				CreateStorage(&Storage{Path: path, Format: "qcow2", Size: size})
			}
			if mode/10 == 0 {
				result = append(result, path)
			} else {
				result = append(result, data.GetVmname()+"-"+strconv.Itoa(j)+".img")
			}
			j++
		}
	}
	fmt.Println("StorageProcess Result: ")
	fmt.Println(result)
	return strings.Join(result, ",")
}

//path, name string, format, size int
func CreateStorage(s *Storage) error {
	fmt.Println("----storage create----")
	if s.Size < 0 {
		return fmt.Errorf("Wrong storage size !!")
	}
	if s.Format != "qcow2" && s.Format != "raw" {
		return fmt.Errorf("Wrong storage format !!")
	}

	var cmd []string

	//qemu-img create [-f format] filename [size]

	cmdarray := []string{"create", "-f", s.Format, s.Path, strconv.Itoa(s.Size) + "M"}

	fmt.Println(cmdarray)

	cmd = append(cmd, cmdarray...)

	RunStorageCmd(cmd)

	return nil
}

func DeleteStorage(s *Storage) error {
	var cmd []string

	filepath := etc.GeneratePath(s.Path, s.Name)
	if FileExistsCheck(filepath) {
		cmd = append(cmd, "info")
		cmd = append(cmd, filepath+".img")

		return nil
	}
	RunStorageCmd(cmd)

	return fmt.Errorf("File not exits!!")
}

func ResizeStorage(s *Storage, soon bool) error {
	//qemu-img resize filename size

	var cmd []string

	cmd = append(cmd, "qemu-img")
	cmd = append(cmd, "resize")
	cmd = append(cmd, etc.GeneratePath(s.Path, s.Name)+".img")
	cmd = append(cmd, strconv.Itoa(s.Size)+"M")

	RunStorageCmd(cmd)

	return nil
}

func InformationStorage(s *Storage) error {
	//qemu-img info [-f format] filename
	var cmd []string

	cmd = append(cmd, "qemu-img")
	cmd = append(cmd, "info")
	cmd = append(cmd, etc.GeneratePath(s.Path, s.Name))

	RunStorageCmd(cmd)
	return nil
}

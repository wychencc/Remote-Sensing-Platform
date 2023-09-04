package main

import (
	"log"
	"os"
	"os/exec"
)

func call(model_path string, s1 string, s2 string, s2_cloud string) {

	//设定指令，调用call_model.py
	print("call : model_path:   ", model_path)
	print("\n")
	cmd := exec.Command("python", "./call_model1.py", "model_path", model_path, "s1_path", s1, "s2_path", s2, "s2_cloud_path", s2_cloud)

	// 将Python脚本的输出连接到标准输出和标准错误输出
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 执行命令
	err := cmd.Run()
	if err != nil {
		log.Fatal("执行命令时发生错误:", err)
	}

}
func main() {

	modelPath := "model_database/model1.pth"
	s1 := "E:/GO/code1/image_database/s1.tif"
	s2 := "E:/GO/code1/image_database/s2.tif"
	s2_cloud := "E:/GO/code1/image_database/s2_cloud.tif"
	print("modelPath:   ", modelPath)
	print("\n")
	call(modelPath, s1, s2, s2_cloud)
}

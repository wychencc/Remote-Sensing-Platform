package infer

import (
	"log"
	"os"
	"os/exec"
)

func call(model_path, s1, s2_cloud, save_path string) {

	//设定指令，调用call_model.py
	print("call : model_path:   ", model_path)
	print("\n")
	cmd := exec.Command("python", "./infer/call_model1.py", "model_path", model_path, "s1_path", s1, "s2_cloud_path", s2_cloud, "save_path", save_path)

	// 将Python脚本的输出连接到标准输出和标准错误输出
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 执行命令
	err := cmd.Run()
	if err != nil {
		log.Fatal("执行命令时发生错误:", err)
	}

}
func trans(opt, trans, data string) {
	//设定指令，调用call_model.py
	opt_path := opt
	trans_path := trans
	data_type := data

	cmd := exec.Command("python", "./infer/trans.py", "opt_path", opt_path, "trans_path", trans_path, "data_type", data_type)
	// 将Python脚本的输出连接到标准输出和标准错误输出
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// 执行命令
	err := cmd.Run()
	if err != nil {
		log.Fatal("执行命令时发生错误:", err)
	}
}
func GetInferenceResult(sar_path, cloud_path, save_path string) {
	modelPath := "./infer/model_database/model1.pth"
	s1 := sar_path
	s2_cloud := cloud_path
	print("modelPath:   ", modelPath)
	print("\n")
	call(modelPath, s1, s2_cloud, save_path)
}

func TransImage(opt_path, trans_path, data_type string) {
	print("opt_path", opt_path)
	print("trans_path", trans_path)
	trans(opt_path, trans_path, data_type)
}

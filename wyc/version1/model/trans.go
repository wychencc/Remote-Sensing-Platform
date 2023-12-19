package model

import "version1_copy/infer"

func GetTrans(opt_path, trans_path, data_type string) {
	infer.TransImage(opt_path, trans_path, data_type)
}

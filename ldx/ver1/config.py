# -*- coding: utf-8 -*-
"""
Created on Wed Oct  7 16:43:07 2020

@author: ssk
"""


class config():
    # datasets config 数据集设定
    dataset_dir = "/public/home/sjh_2112103041/data/" # 训练数据集路径 自动分割成验证集和训练集
    predict_dataset_dir = "/public/home/sjh_2112103041/Dataset/test/ROIs1970_fall/sar"# 要预测的图片的路径
    #predict_dataset_dir = '/public/home/sjh_2112103041/IM/sar'
    width = 128  # 图片大小
    height = 128
    #threads =2  # num_of_workers 在windows环境下大于0会有BUG——>broken pipe
    threads =0  # num_of_workers 在windows环境下大于0会有BUG——>broken pipe

    # outputimg dir 输出图片路径
    output_dir = "./output"

    # train options 训练设置
    use_gpu = False  # 要不要用GPU？
    save_frequency = 400  # 多少个iteration保存一次网络呢？

    show_freq = 200  # 多少轮进行测试并展示成果？
    train_size = 0.8  # 训练集占总数据集合的百分比
    batch_size =1# batchsize 越大占用显存越多
    # batch_accu = 256
    # 网络初始化pth路径
    net_init ='./net_state_dict/net_epoch_47_iteration_400.pth'
    epoch = 100  # 训练轮数
    lr = 7e-5 # 学习率
    beta1 = 0.9  # ADAM优化器的beta1
    in_ch = 5  # 输入通道数
    out_ch = 3  # 输出通道数
    alpha = 0.1  # resnet aplha
    num_layers = 16  # resnet层数
    feature_sizes = 256  # resnetfeature

from PIL import Image
import sys
import os
import sys, os
import numpy as np
import json
import torch as t
import torch
from models.DSen2_CR import DSen2_CR
import utils.visualize as visualize
# from sen12mscrdata import SEN12MSDataset
from dataset import SEN12MSDataset
import time
import importlib.util
import numpy as np
import utils.img_operations as imgop
import config
import matplotlib.pyplot as plt
os.environ['CUDA_VISIBLE_DEVICES'] = "0"


def save_state_dict(net, epoch, iteration):
    net_path = os.path.join("./net_state_dict", "net_epoch_{}_iteration_{}.pth".format(epoch, iteration))
    if not os.path.exists("./net_state_dict"):
        os.makedirs("./net_state_dict")
    torch.save(net.state_dict(), net_path)
    print("第{}轮训练结果已经保存".format(epoch))


def train(config,model_path,s1_path,s2_path,s2_cloud_path):
    
    # 读取映射文件
    with open("map_database.json", "r") as f:
        mapping = json.load(f)

    # 获取网络模型地址
    net_file_path = mapping.get(model_path)

    print("model_path->net_file_path:",model_path,"->",net_file_path)

    # 根据文件路径导入模块
    spec = importlib.util.spec_from_file_location("network_module", net_file_path)
    network_module = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(network_module)

    # 网络定义
    # 使用导入的模块中的内容
    model_class = getattr(network_module, "DSen2_CR")
    net = model_class(config.in_ch, config.out_ch, config.alpha, config.num_layers, config.feature_sizes)
    net = net.eval()
    print("网络定义完毕")

    param = torch.load(model_path, map_location='cpu')
    net.load_state_dict(param)
    print("model_path:",model_path)
    print("加载 model_path 作为网络模型")
    
    
    # 将数据装入gpu（在有gpu且使用GPU进行训练的情况下）
    cloud_img = t.FloatTensor(config.batch_size, config.in_ch, config.width, config.height)
    ground_truth = t.FloatTensor(config.batch_size, config.out_ch, config.width, config.height)
    csm_img = t.FloatTensor(config.batch_size, 1, config.width, config.height)

    
    iters = 0
    PSNR_4 = 0
    SSIM_4 = 0
    SAM_4 = 0
    MAE_4 = 0


    # 数据集 dataset 与 dataloader
    print("正在创建数据集...")
    train_dataset = SEN12MSDataset(s1_path,s2_path,s2_cloud_path)

    train_dataloader = t.utils.data.DataLoader(train_dataset, batch_size=config.batch_size, shuffle=True)

    # 优化器
    opt = t.optim.Adam(net.parameters(), lr=config.lr, betas=(config.beta1, 0.999), weight_decay=0.00001)

    # 损失函数
    CARL_Loss = imgop.carl_error

    '''
    # 如果使用GPU 则把这些玩意全部放进显存里面去
    if config.use_gpu:
        net = net.cuda()
        cloud_img = cloud_img.cuda()
        ground_truth = ground_truth.cuda()
        csm_img = csm_img.cuda()

    
    '''
    print("开始训练...")
    # 开始循环
    for epoch in range(0, config.epoch):
        print("****")
        loss_log = []
        epoch_start_time = time.time()

        # 数据集的小循环
        for iteration, batch in enumerate(train_dataloader, start=1):
            # 数据操作    numpy 数据转成tensor
            print("!!!!")
            img_cld, img_csm, img_truth, img_sar,img_cloud= batch
            img_cld = cloud_img.resize_(img_cld.shape).copy_(img_cld)
            img_csm = csm_img.resize_(img_csm.shape).copy_(img_csm)
            img_truth = ground_truth.resize_(img_truth.shape).copy_(img_truth)
            img_fake = net(img_cld)
            opt.zero_grad()
            loss = CARL_Loss(img_truth,img_csm,img_fake)
            loss.backward()
            opt.step()

        print("第{}轮训练完毕，用时{}S".format(epoch, time.time() - epoch_start_time))
    
    t.save(net.state_dict(), './model333.pth')







model_path = sys.argv[2]
s1_path = sys.argv[4]
s2_path = sys.argv[6]
s2_cloud_path = sys.argv[8]


print("python :",model_path)
print("s1_path :",s1_path)
print("s2_path :",s2_path)
print("s2_cloud_path :",s2_cloud_path)
#model = torch.load(model_path,map_location=torch.device('cpu'))

# 调用预测函数
myconfig = config.config()
print("config:",myconfig)
train(myconfig,model_path,s1_path,s2_path,s2_cloud_path)
import torch
from models.DSen2_CR import DSen2_CR
from metrics import SSIM, PSNR, SAM, MAE
import torchvision.models as models
import matplotlib.pyplot as plt
import sys,os
import config
import torch as t
import torchvision.transforms as transforms
import numpy as np
import collections
from PIL import Image
import utils.feature_detectors as feature_detectors
import rasterio
import json
import importlib.util

def get_opt_image_shallow(path):
    src = rasterio.open(path, 'r', driver='GTiff')
    image = src.read()
    # print(image.shape)
    src.close()
    image[np.isnan(image)] = np.nanmean(image)  # fill holes and artifacts

    return image.astype('float32')
def get_opt_image(path):
    src = rasterio.open(path, 'r', driver='GTiff')
    image = src.read()
    if image.shape[0] > 3:
        image = image[1:4, :, :]
    # print(image.shape)
    src.close()
    image[np.isnan(image)] = np.nanmean(image)  # fill holes and artifacts

    return image.astype('float32')
def get_sar_image(path):
    src = rasterio.open(path, 'r', driver='GTiff')
    image = src.read()
    # [2,256,256]
    src.close()
    image[np.isnan(image)] = np.nanmean(image)  # fill holes and artifacts

    return image.astype('float32')
def get_normalized_data(data_image, data_type):
    clip_min = [[-25.0, -32.5], [0, 0, 0, 0], [0, 0, 0, 0]]
    clip_max = [[0, 0], [10000, 10000, 10000, 10000], [10000, 10000, 10000, 10000]]

    max_val = 1
    scale = 10000

    # SAR
    if data_type == 1:
        print("len(data_image):",len(data_image))
        for channel in range(len(data_image)):
            print("channel:",channel," data_type-1:",data_type )
            
            data_image[channel] = np.clip(data_image[channel], clip_min[data_type - 1][channel],clip_max[data_type - 1][channel])
            data_image[channel] -= clip_min[data_type - 1][channel]
            data_image[channel] = max_val * (data_image[channel] / (clip_max[data_type - 1][channel] - clip_min[data_type - 1][channel]))
    # OPT
    elif data_type == 2 or data_type == 3:
        for channel in range(len(data_image)):
            data_image[channel] = np.clip(data_image[channel], clip_min[data_type - 1][channel],
                                          clip_max[data_type - 1][channel])
        data_image /= scale
        # [4,256,256]

    return data_image
def transform( x):
        adaptive_pool = torch.nn.AdaptiveAvgPool2d((128, 128))
        x = torch.tensor(x, requires_grad=False, dtype=torch.float32)
        x = adaptive_pool(x)

        return x

def __getitem__( s1_path, s2_cloud_path):
        # 读取
        sar_path = s1_path
        img_sar = get_sar_image(sar_path)
#         rgb_path = s2_path
        img_cloud_path = s2_cloud_path
#         img_rgb = get_opt_image(rgb_path)
        cloud = get_opt_image_shallow(img_cloud_path)
        img_cloud = get_opt_image(img_cloud_path)
        
        #s2CSMimg = feature_detectors.get_cloud_cloudshadow_mask(cloud, 0.2)
        #s2CSMimg = np.expand_dims(s2CSMimg, axis=0).astype('float32')

        # SAR VV归到[-25,0],VH[-35,0].OPTICAL [0,10000]-[0,1]
#         img_rgb = get_normalized_data(img_rgb, data_type=2)
        img_cloud = get_normalized_data(img_cloud, data_type=3)
        img_sar = get_normalized_data(img_sar, data_type=1)
#         img_rgb = transform(img_rgb)
        img_cloud = transform(img_cloud)
        img_sar = transform(img_sar)
        #s2CSMimg = transform(s2CSMimg)

        inputdata = np.concatenate((img_sar, img_cloud), axis=0)

        # 返回三个数据 输入神经网络的数据（8,256,256） Cloud and shadow mask（csm 1*256*256） 无云图片（target 3*256*256）
        return inputdata, img_sar,img_cloud

def predict(config,model_path,s1_path, s2_cloud_path, save_path):

    # 读取映射文件
    with open("./infer/map_database.json", "r") as f:
        mapping = json.load(f)

    # 获取网络模型地址
    net_file_path = mapping.get("model_pth")
    print(net_file_path)

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

    # 如果使用GPU 则把这些玩意全部放进显存里面去
    # 虽然不需要CSM 但是还是把它输出一下吧

    if config.use_gpu:
        net = net.cuda()
        cloud_img = cloud_img.cuda()
#         ground_truth = ground_truth.cuda()
        csm_img = csm_img.cuda()

    print("*******")
    #adaptive_pool = torch.nn.AdaptiveAvgPool2d((128, 128)) 
    with t.no_grad():       
        img_cld, img_sar,img_cloud= __getitem__(s1_path,s2_cloud_path)
         # 先转换为张量
        img_cld = torch.from_numpy(img_cld)

        img_cld=img_cld.unsqueeze(0)
        #img_csm=img_csm.unsqueeze(0)
#         img_truth=img_truth.unsqueeze(0)
        img_sar=img_sar.unsqueeze(0)
        img_cloud=img_cloud.unsqueeze(0)

        img_cld = cloud_img.resize_(img_cld.shape).copy_(img_cld)
        #img_csm = csm_img.resize_(img_csm.shape).copy_(img_csm)
#         img_truth = ground_truth.resize_(img_truth.shape).copy_(img_truth)
        print("^^^^^^^^^^^^^")
        print("inputdata shape:",img_cld.shape)
        img_fake = net(img_cld)

        plt.figure(figsize=(3, 3))
        plt.title('Pre')
        plt.axis('off')
        pre =img_fake.detach().cpu()[0].permute(1, 2, 0)
        pre = pre - pre.min()
        pre = pre / pre.max()
        plt.imshow(pre)
        plt.xticks([])  # 去掉x轴
        plt.yticks([])  # 去掉y轴
        plt.subplots_adjust(top=1, bottom=0, right=1, left=0, hspace=0, wspace=0)
        plt.margins(0, 0)
        plt.savefig(save_path)
        print('Testing done. ')

model_path = sys.argv[2]
s1_path = sys.argv[4]
s2_cloud_path = sys.argv[6]
save_path = sys.argv[8]

'''
model_path = "model_database/model1.pth"
s1_path = "image_database/s1.tif"
s2_path = "image_database/s2.tif"
s2_cloud_path = "image_database/s2_cloud.tif"
'''

print("python :",model_path)
print("s1_path :",s1_path)
# print("s2_path :",s2_path)
print("s2_cloud_path :",s2_cloud_path)
#model = torch.load(model_path,map_location=torch.device('cpu'))

# 调用预测函数
myconfig = config.config()
print("config:",myconfig)
predict(myconfig,model_path,s1_path,s2_cloud_path, save_path)


import os
import torch.utils.data as data
import torch
import torch.nn.functional as F
import numpy as np

import utils.feature_detectors as feature_detectors
import rasterio


def get_opt_image_shallow(path):
    src = rasterio.open(path, 'r', driver='GTiff')
    image = src.read()
    # print(image.shape)
    src.close()
    image[np.isnan(image)] = np.nanmean(image)  # fill holes and artifacts

    return image.astype('float32')


def get_opt_image(path):
    src = rasterio.open(path, 'r', driver='GTiff')
    image = src.read()[1:4, :, :]
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
        for channel in range(len(data_image)):
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


class SEN12MSDataset(data.Dataset):
    def __init__(self, s1_path,s2_path,s2_cloud_path):
      
        self.sar_path=s1_path
        self.rgb_path=s2_path
        self.rgb_cloud_path=s2_cloud_path
        self.sar_files = os.listdir(s1_path)
        self.rgb_files = os.listdir(s2_path)
        self.rgb_cloud_files = os.listdir(s2_cloud_path)
        
    # 其他初始化操作

    def transform(self,x):
        adaptive_pool = torch.nn.AdaptiveAvgPool2d((128, 128))
        x = torch.tensor(x, requires_grad=False, dtype=torch.float32)
        x = adaptive_pool(x)

        return x


    def __getitem__(self,index):
        # 读取
        rgb_file = self.rgb_files[index]
        rgb_cloud_file = self.rgb_cloud_files[index]
        sar_file = self.sar_files[index]

        # 根据文件路径加载数据
        img_sar = get_sar_image(os.path.join(self.sar_path, sar_file))
        img_rgb = get_opt_image(os.path.join(self.rgb_path, rgb_file))
        cloud = get_opt_image_shallow(os.path.join(self.rgb_cloud_path, rgb_cloud_file))
        img_cloud=get_opt_image(os.path.join(self.rgb_cloud_path, rgb_cloud_file))

        print("")
    
        
        s2CSMimg = feature_detectors.get_cloud_cloudshadow_mask(cloud, 0.2)
        s2CSMimg = np.expand_dims(s2CSMimg, axis=0).astype('float32')

        # SAR VV归到[-25,0],VH[-35,0].OPTICAL [0,10000]-[0,1]
        img_rgb = get_normalized_data(img_rgb, data_type=2)
        img_cloud = get_normalized_data(img_cloud, data_type=3)
        img_sar = get_normalized_data(img_sar, data_type=1)
        img_rgb = self.transform(img_rgb)
        img_cloud = self.transform(img_cloud)
        img_sar = self.transform(img_sar)
        s2CSMimg = self.transform(s2CSMimg)

        inputdata = np.concatenate((img_sar, img_cloud), axis=0)

        # 返回三个数据 输入神经网络的数据（8,256,256） Cloud and shadow mask（csm 1*256*256） 无云图片（target 3*256*256）
        return inputdata, s2CSMimg, img_rgb, img_sar,img_cloud



    def __len__(self):
         return min(len(self.rgb_files), len(self.rgb_cloud_files), len(self.sar_files))

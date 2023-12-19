import rasterio
import numpy as np
import cv2 as cv
from PIL import Image
import torch
import matplotlib.pyplot as plt
import sys, os

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
        print("len(data_image):", len(data_image))
        for channel in range(len(data_image)):
            print("channel:", channel, " data_type-1:", data_type)

            data_image[channel] = np.clip(data_image[channel], clip_min[data_type - 1][channel],
                                          clip_max[data_type - 1][channel])
            data_image[channel] -= clip_min[data_type - 1][channel]
            data_image[channel] = max_val * (
                        data_image[channel] / (clip_max[data_type - 1][channel] - clip_min[data_type - 1][channel]))
    # OPT
    elif data_type == 2 or data_type == 3:
        for channel in range(len(data_image)):
            data_image[channel] = np.clip(data_image[channel], clip_min[data_type - 1][channel],
                                          clip_max[data_type - 1][channel])
        data_image /= scale
        # [4,256,256]
    return data_image
def transform(x):
    adaptive_pool = torch.nn.AdaptiveAvgPool2d((128, 128))
    x = torch.tensor(x, requires_grad=False, dtype=torch.float32)
    x = adaptive_pool(x)
    return x

def transpose(opt_path, trans_path, data_type):

    if data_type == "1":
        image = get_sar_image(opt_path)
        image = get_normalized_data(image, data_type=1)
        image = transform(image)
        image = image.unsqueeze(0)
        plt.figure(figsize=(3, 3))
        plt.axis('off')
        sar = image.cpu()[0][0]
        sar = sar - sar.min()
        sar = sar / sar.max()
        plt.imshow(sar)
        plt.xticks([])  # 去掉x轴
        plt.yticks([])  # 去掉y轴
        plt.subplots_adjust(top=1, bottom=0, right=1, left=0, hspace=0, wspace=0)
        plt.margins(0, 0)
        plt.savefig(trans_path)
    else:
        image = get_opt_image(opt_path)
        image = get_normalized_data(image, data_type=3)
        image = transform(image)
        image = image.unsqueeze(0)
        plt.figure(figsize=(3, 3))
        plt.axis('off')
        rgb = image.cpu()[0].permute(1, 2, 0)
        rgb = rgb - rgb.min()
        rgb = rgb / rgb.max()
        plt.imshow(rgb)
        plt.xticks([])  # 去掉x轴
        plt.yticks([])  # 去掉y轴
        plt.subplots_adjust(top=1, bottom=0, right=1, left=0, hspace=0, wspace=0)
        plt.margins(0, 0)
        plt.savefig(trans_path)

opt_path = sys.argv[2]
trans_path = sys.argv[4]
data_type = sys.argv[6]
transpose(opt_path, trans_path, data_type)
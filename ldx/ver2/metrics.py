import math
import torch
import torch.nn.functional as F
from torch.autograd import Variable
import numpy as np
from math import exp

def gaussian(window_size, sigma):
    gauss = torch.Tensor([exp(-(x - window_size / 2) ** 2 / float(2 * sigma ** 2)) for x in range(window_size)])
    return gauss / gauss.sum()

def create_window(window_size, channel):
    _1D_window = gaussian(window_size, 1.5).unsqueeze(1)
    _2D_window = _1D_window.mm(_1D_window.t()).float().unsqueeze(0).unsqueeze(0)
    window = Variable(_2D_window.expand(channel, 1, window_size, window_size))
    return window

def SSIM(img1, img2):
    (_, channel, _, _) = img1.size()
    window_size = 11
    window = create_window(window_size, channel).cuda()
    mu1 = F.conv2d(img1, window, padding=window_size // 2, groups=channel)
    mu2 = F.conv2d(img2, window, padding=window_size // 2, groups=channel)

    mu1_sq = mu1.pow(2)
    mu2_sq = mu2.pow(2)
    mu1_mu2 = mu1 * mu2

    sigma1_sq = F.conv2d(img1 * img1, window, padding=window_size // 2, groups=channel) - mu1_sq
    sigma2_sq = F.conv2d(img2 * img2, window, padding=window_size // 2, groups=channel) - mu2_sq
    sigma12 = F.conv2d(img1 * img2, window, padding=window_size // 2, groups=channel) - mu1_mu2

    C1 = 0.01 ** 2
    C2 = 0.03 ** 2

    ssim_map = ((2 * mu1_mu2 + C1) * (2 * sigma12 + C2)) / ((mu1_sq + mu2_sq + C1) * (sigma1_sq + sigma2_sq + C2))
    return ssim_map.mean().item()

def PSNR(img1, img2, mask=None):
    if mask is not None:
        mse = (img1 - img2) ** 2
        B, C, H, W = mse.size()
        mse = torch.sum(mse * mask.float()) / (torch.sum(mask.float()) * C)
    else:
        mse = torch.mean((img1 - img2) ** 2)

    if mse == 0:
        return 100
    PIXEL_MAX = 1
    return 20 * math.log10(PIXEL_MAX / math.sqrt(mse))

def MAE(ref, tgt, is_mean = True):
    ref = ref.cpu().detach().numpy()
    tgt = tgt.cpu().detach().numpy()
    pixel_errors = np.absolute(ref - tgt)
    if is_mean:
        return pixel_errors.mean()
    else:
        return pixel_errors


def SAM(ref, tgt, is_mean = True):

    ref = ref.cpu().detach().numpy()
    tgt = tgt.cpu().detach().numpy()

    ref = ref.transpose(0, 2, 3, 1)
    tgt = tgt.transpose(0, 2, 3, 1)

    kernel = np.einsum('...k,...k', ref, tgt)

    square_norm_ref = np.einsum('...k,...k', ref, ref).clip(min=np.finfo(np.float16).eps)
    square_norm_tgt = np.einsum('...k,...k', tgt, tgt).clip(min=np.finfo(np.float16).eps)
    normalized_kernel = kernel / np.sqrt(square_norm_ref * square_norm_tgt)

    angles = np.arccos(normalized_kernel.clip(min=-1, max=1)) / np.pi * 180

    if is_mean:
        return angles.mean()
    else:
        return angles

































# import numpy as np
# import math
#
# # def parse_args():
# #     parser = argparse.ArgumentParser(description='script to compute all statistics')
# #     parser.add_argument('--data-path', help='Path to ground truth data', type=str)
# #     parser.add_argument('--output-path', help='Path to output data', type=str)
# #     parser.add_argument('--debug', default=0, help='Debug', type=int)
# #     args = parser.parse_args()
# #     return args
#
# import math
# import torch
# import torch.nn.functional as F
# from torch.autograd import Variable
# import numpy as np
# from math import exp
#
# def gaussian(window_size, sigma):
#     gauss = torch.Tensor([exp(-(x - window_size / 2) ** 2 / float(2 * sigma ** 2)) for x in range(window_size)])
#     return gauss / gauss.sum()
#
# def create_window(window_size, channel):
#     _1D_window = gaussian(window_size, 1.5).unsqueeze(1)
#     _2D_window = _1D_window.mm(_1D_window.t()).float().unsqueeze(0).unsqueeze(0)
#     window = Variable(_2D_window.expand(channel, 1, window_size, window_size))
#     return window
#
# def compare_ssim(img1, img2):
#     (_, channel, _, _) = img1.size()
#     window_size = 11
#     window = create_window(window_size, channel).cuda()
#     mu1 = F.conv2d(img1, window, padding=window_size // 2, groups=channel)
#     mu2 = F.conv2d(img2, window, padding=window_size // 2, groups=channel)
#
#     mu1_sq = mu1.pow(2)
#     mu2_sq = mu2.pow(2)
#     mu1_mu2 = mu1 * mu2
#
#     sigma1_sq = F.conv2d(img1 * img1, window, padding=window_size // 2, groups=channel) - mu1_sq
#     sigma2_sq = F.conv2d(img2 * img2, window, padding=window_size // 2, groups=channel) - mu2_sq
#     sigma12 = F.conv2d(img1 * img2, window, padding=window_size // 2, groups=channel) - mu1_mu2
#
#     C1 = 0.01 ** 2
#     C2 = 0.03 ** 2
#
#     ssim_map = ((2 * mu1_mu2 + C1) * (2 * sigma12 + C2)) / ((mu1_sq + mu2_sq + C1) * (sigma1_sq + sigma2_sq + C2))
#     return ssim_map.mean().item()
#
# def compare_psnr(img1, img2,mse):
#     #if mask is not None:
#       #  mse = (img1 - img2) ** 2
#        # B, C, H, W = mse.size()
#        # mse = torch.sum(mse * mask.float()) / (torch.sum(mask.float()) * C)
#   #  else:
#     #mse = torch.mean((img1 - img2) ** 2)
#
#     if mse == 0:
#         return 100
#     PIXEL_MAX = 1
#     return 20 * math.log10(PIXEL_MAX / math.sqrt(mse))
#
# def compare_mae(ref, tgt, is_mean = True):
#     #ref = ref.cpu().detach().numpy()
#     #tgt = tgt.cpu().detach().numpy()
#     pixel_errors = np.absolute(ref - tgt)
#     if is_mean:
#         return pixel_errors.mean()
#     else:
#         return pixel_errors
#
#
# def compare_sam(ref, tgt, is_mean = True):
#
#  #   ref = ref.cpu().detach().numpy()
#   #  tgt = tgt.cpu().detach().numpy()
#
#    # ref = ref.transpose(0, 2, 3, 1)
#     #tgt = tgt.transpose(0, 2, 3, 1)
#
#     kernel = np.einsum('...k,...k', ref, tgt)
#
#     square_norm_ref = np.einsum('...k,...k', ref, ref).clip(min=np.finfo(np.float16).eps)
#     square_norm_tgt = np.einsum('...k,...k', tgt, tgt).clip(min=np.finfo(np.float16).eps)
#     normalized_kernel = kernel / np.sqrt(square_norm_ref * square_norm_tgt)
#
#     angles = np.arccos(normalized_kernel.clip(min=-1, max=1)) / np.pi * 180
#
#     if is_mean:
#         return angles.mean()
#     else:
#         return angles
#
#
#
#
#
# #def compare_mae(img_true, img_test):
#  #   img_true = img_true.astype(np.float32)
#     #img_test = img_test.astype(np.float32)
#     #return np.sum(np.abs(img_true - img_test)) / np.sum(img_true + img_test)
#
#
# #def compare_sam(x_true, x_pred):
#     #"calculate method in PSGAN code"
#
#     #assert x_true.ndim == 3 and x_true.shape == x_pred.shape
#     #dot_sum = np.sum(x_true * x_pred, axis=2)
#     #norm_true = np.linalg.norm(x_true, axis=2)
#     #norm_pred = np.linalg.norm(x_pred, axis=2)
#
#  #   res = np.arccos(dot_sum / norm_pred / norm_true)
#   #  is_nan = np.nonzero(np.isnan(res))
#     #for (x, y) in zip(is_nan[0], is_nan[1]):
#     #    res[x, y] = 0
#     #
#  #   sam = np.mean(res)
#    # sam = math.degrees(sam)
#     # sam_rad = np.zeros(x_pred.shape[0,1])
#     # for x in range(x_true.shape[0]):
#     #     for y in range(x_true.shape[1]):
#     #         tem_pred = x_pred[x,y].ravel()
#     #         tem_true = x_true[x, y].ravel()
#     #         sam_rad[x,y] = np.arccos(tem_pred / (np.linalg.norm(tem_pred)*tem_true / np.linalg.norm(tem_true)))
#     # sam = sam_rad.mean() * 180 / np.pi
#
#     #return sam
#
#

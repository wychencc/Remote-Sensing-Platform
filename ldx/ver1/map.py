import json

# 创建映射字典
mapping = {
    "model_database/model1.pth": "models/DSen2_CR.py",
    # 添加其他映射关系...
}

# 指定要保存的文件路径
file_path = "map_database.json"

# 将映射字典写入映射文件
with open(file_path, "w") as f:
    json.dump(mapping, f)
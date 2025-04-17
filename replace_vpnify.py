import os
import base64
import random

def generate_random_8digit_base64():
    """生成 8 位随机数字并进行 base64 编码"""
    digits = ''.join(random.choices('0123456789', k=8))
    b64 = base64.b64encode(digits.encode()).decode('utf-8')
    return b64

def replace_in_file(file_path, replacements, script_path):
    # 跳过当前脚本自身
    if os.path.abspath(file_path) == script_path:
        return

    with open(file_path, 'r', encoding='utf-8') as f:
        content = f.read()
    original_content = content
    for old, new in replacements:
        if old in content:
            print(f"🔍 在 {file_path} 中找到 '{old}'，准备替换为 '{new}'")
        content = content.replace(old, new)
    if content != original_content:
        with open(file_path, 'w', encoding='utf-8') as f:
            f.write(content)
        print(f"✅ 替换完成: {file_path}")
    else:
        print(f"🚫 无需替换: {file_path}")

def replace_in_project(root_dir, replacements, script_path):
    for dirpath, _, filenames in os.walk(root_dir):
        for filename in filenames:
            if filename.endswith(('.go', '.py', '.h', '.txt', '.md', '.json', '.xml', '.yaml', '.yml')):
                file_path = os.path.join(dirpath, filename)
                replace_in_file(file_path, replacements, script_path)

if __name__ == '__main__':
    project_root = '.'
    script_path = os.path.abspath(__file__)

    old_base64 = 'MTg5jNU2ODk='  # 要替换的旧字符串
    new_base64 = generate_random_8digit_base64()

    replacements = [
        ('Hope', 'Dancy'),
        (old_base64, new_base64),
    ]

    print(f"🔁 正在将 '{old_base64}' 替换为 '{new_base64}'（忽略当前脚本）")
    replace_in_project(project_root, replacements, script_path)


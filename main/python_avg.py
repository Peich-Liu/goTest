# python_avg.py
import sys

def main():
    # sys.argv[0] 是脚本名字，真正的数据从 sys.argv[1:] 开始
    nums = list(map(int, sys.argv[1:]))
    if not nums:
        print(0.0)
        return
    avg = sum(nums) / len(nums)
    print(avg)

if __name__ == "__main__":
    main()
